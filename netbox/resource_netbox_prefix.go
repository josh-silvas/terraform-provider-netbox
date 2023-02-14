package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/ipam"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxPrefixCreate,
		Read:   resourceNetboxPrefixRead,
		Update: resourceNetboxPrefixUpdate,
		Delete: resourceNetboxPrefixDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#prefixes):

> A prefix is an IPv4 or IPv6 network and mask expressed in CIDR notation (e.g. 192.0.2.0/24). A prefix entails only the "network portion" of an IP address: All bits in the address not covered by the mask must be zero. (In other words, a prefix cannot be a specific IP address.)
>
> Prefixes are automatically organized by their parent aggregates. Additionally, each prefix can be assigned to a particular site and virtual routing and forwarding instance (VRF). Each VRF represents a separate IP space or routing table. All prefixes not assigned to a VRF are considered to be in the "global" table.`,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"active", "reserved", "deprecated", "container"}, false),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_pool": {
				Type:     schema.TypeBool,
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
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxPrefixCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritablePrefix{}

	prefix := d.Get("prefix").(string)
	status := d.Get("status").(string)
	description := d.Get("description").(string)
	isPool := d.Get("is_pool").(bool)

	data.Prefix = &prefix
	data.Status = status

	data.Description = description
	data.IsPool = isPool

	if vrfID, ok := d.GetOk("vrf_id"); ok {
		data.Vrf = int64ToPtr(int64(vrfID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if siteID, ok := d.GetOk("site_id"); ok {
		data.Site = int64ToPtr(int64(siteID.(int)))
	}

	if vlanID, ok := d.GetOk("vlan_id"); ok {
		data.Vlan = int64ToPtr(int64(vlanID.(int)))
	}

	if roleID, ok := d.GetOk("role_id"); ok {
		data.Role = int64ToPtr(int64(roleID.(int)))
	}

	data.Tags = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := ipam.NewIpamPrefixesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamPrefixesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxPrefixRead(d, m)
}

func resourceNetboxPrefixRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamPrefixesReadParams().WithID(id)

	res, err := api.Ipam.IpamPrefixesRead(params, nil)
	if err != nil {
		return err
	}

	if err := d.Set("description", res.GetPayload().Description); err != nil {
		return err
	}
	if err := d.Set("is_pool", res.GetPayload().IsPool); err != nil {
		return err
	}
	if res.GetPayload().Status != nil {
		if err := d.Set("status", res.GetPayload().Status.Value); err != nil {
			return err
		}
	}
	if res.GetPayload().Prefix != nil {
		if err := d.Set("prefix", res.GetPayload().Prefix); err != nil {
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

	if res.GetPayload().Site != nil {
		if err := d.Set("site_id", res.GetPayload().Site.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("site_id", nil); err != nil {
			return err
		}
	}

	if res.GetPayload().Vlan != nil {
		if err := d.Set("vlan_id", res.GetPayload().Vlan.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("vlan_id", nil); err != nil {
			return err
		}
	}

	if res.GetPayload().Role != nil {
		if err := d.Set("role_id", res.GetPayload().Role.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("role_id", nil); err != nil {
			return err
		}
	}

	return d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))
	// FIGURE OUT NESTED VRF AND NESTED VLAN (from maybe interfaces?)
}

func resourceNetboxPrefixUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritablePrefix{}
	prefix := d.Get("prefix").(string)
	status := d.Get("status").(string)
	description := d.Get("description").(string)
	isPool := d.Get("is_pool").(bool)

	data.Prefix = &prefix
	data.Status = status

	data.Description = description
	data.IsPool = isPool

	if vrfID, ok := d.GetOk("vrf_id"); ok {
		data.Vrf = int64ToPtr(int64(vrfID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if siteID, ok := d.GetOk("site_id"); ok {
		data.Site = int64ToPtr(int64(siteID.(int)))
	}

	if vlanID, ok := d.GetOk("vlan_id"); ok {
		data.Vlan = int64ToPtr(int64(vlanID.(int)))
	}

	if roleID, ok := d.GetOk("role_id"); ok {
		data.Role = int64ToPtr(int64(roleID.(int)))
	}

	data.Tags = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := ipam.NewIpamPrefixesUpdateParams().WithID(id).WithData(&data)
	// nolint: errcheck
	if _, err := api.Ipam.IpamPrefixesUpdate(params, nil); err != nil {
		return err
	}
	return resourceNetboxPrefixRead(d, m)
}

func resourceNetboxPrefixDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamPrefixesDeleteParams().WithID(id)
	// nolint: errcheck
	if _, err := api.Ipam.IpamPrefixesDelete(params, nil); err != nil {
		return err
	}
	d.SetId("")
	return nil
}
