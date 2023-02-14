package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/ipam"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIPAddressCreate,
		Read:   resourceNetboxIPAddressRead,
		Update: resourceNetboxIPAddressUpdate,
		Delete: resourceNetboxIPAddressDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#ip-addresses):

> An IP address comprises a single host address (either IPv4 or IPv6) and its subnet mask. Its mask should match exactly how the IP address is configured on an interface in the real world.
>
> Like a prefix, an IP address can optionally be assigned to a VRF (otherwise, it will appear in the "global" table). IP addresses are automatically arranged under parent prefixes within their respective VRFs according to the IP hierarchy.`,

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
			},
			"interface_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrf_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"active", "reserved", "deprecated", "dhcp"}, false),
			},
			"dns_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey: tagsSchema,
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"loopback", "secondary", "anycast", "vip", "vrrp", "hsrp", "glbp", "carp"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxIPAddressCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableIPAddress{}
	ipAddress := d.Get("ip_address").(string)
	data.Address = &ipAddress
	data.Status = d.Get("status").(string)
	data.Description = d.Get("description").(string)
	data.Role = d.Get("role").(string)

	if dnsName, ok := d.GetOk("dns_name"); ok {
		data.DNSName = dnsName.(string)
	}

	params := ipam.NewIpamIPAddressesCreateParams().WithData(&data)

	res, err := api.Ipam.IpamIPAddressesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxIPAddressUpdate(d, m)
}

func resourceNetboxIPAddressRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamIPAddressesReadParams().WithID(id)

	res, err := api.Ipam.IpamIPAddressesRead(params, nil)
	if err != nil {
		return err
	}

	if res.GetPayload().AssignedObjectID != nil {
		if err := d.Set("interface_id", res.GetPayload().AssignedObjectID); err != nil {
			return err
		}
	} else {
		if err := d.Set("interface_id", nil); err != nil {
			return err
		}
	}

	if res.GetPayload().Vrf != nil {
		if err := d.Set("vrf_id", res.GetPayload().Vrf.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("vrf_id", nil); err != nil {
			return err
		}
	}

	if res.GetPayload().Tenant != nil {
		if err := d.Set("tenant_id", res.GetPayload().Tenant.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("tenant_id", nil); err != nil {
			return err
		}
	}

	if res.GetPayload().DNSName != "" {
		if err := d.Set("dns_name", res.GetPayload().DNSName); err != nil {
			return err
		}
	}

	if res.GetPayload().Role != nil {
		if err := d.Set("role", res.GetPayload().Role.Value); err != nil {
			return err
		}
	} else {
		if err := d.Set("role", nil); err != nil {
			return err
		}
	}

	if err := d.Set("ip_address", res.GetPayload().Address); err != nil {
		return err
	}
	if err := d.Set("description", res.GetPayload().Description); err != nil {
		return err
	}
	if err := d.Set("status", res.GetPayload().Status.Value); err != nil {
		return err
	}
	return nil
}

func resourceNetboxIPAddressUpdate(d *schema.ResourceData, m interface{}) error {

	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableIPAddress{}

	ipAddress := d.Get("ip_address").(string)
	status := d.Get("status").(string)

	descriptionValue, ok := d.GetOk("description")
	if ok {
		description := descriptionValue.(string)
		data.Description = description
	} else {
		description := " "
		data.Description = description
	}

	data.Status = status
	data.Address = &ipAddress

	if d.HasChange("dns_name") {
		// WritableIPAddress omits empty values so set to ' '
		if dnsName := d.Get("dns_name"); dnsName.(string) == "" {
			data.DNSName = " "
		} else {
			data.DNSName = dnsName.(string)
		}
	}

	if interfaceID, ok := d.GetOk("interface_id"); ok {
		// The other possible type is dcim.interface for devices
		data.AssignedObjectType = strToPtr("virtualization.vminterface")
		data.AssignedObjectID = int64ToPtr(int64(interfaceID.(int)))
	}

	if vrfID, ok := d.GetOk("vrf_id"); ok {
		data.Vrf = int64ToPtr(int64(vrfID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if role, ok := d.GetOk("role"); ok {
		data.Role = role.(string)
	}

	params := ipam.NewIpamIPAddressesUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Ipam.IpamIPAddressesUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxIPAddressRead(d, m)
}

func resourceNetboxIPAddressDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamIPAddressesDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Ipam.IpamIPAddressesDelete(params, nil); err != nil {
		return err
	}
	return nil
}
