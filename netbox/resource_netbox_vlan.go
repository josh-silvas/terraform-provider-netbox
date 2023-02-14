package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVlanCreate,
		Read:   resourceNetboxVlanRead,
		Update: resourceNetboxVlanUpdate,
		Delete: resourceNetboxVlanDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/vlans/#vlans):

> A VLAN represents an isolated layer two domain, identified by a name and a numeric ID (1-4094) as defined in IEEE 802.1Q. VLANs are arranged into VLAN groups to define scope and to enforce uniqueness.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "active",
				ValidateFunc: validation.StringInSlice([]string{"active", "reserved", "deprecated"}, false),
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVlanCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableVLAN{}

	name := d.Get("name").(string)
	vid := int64(d.Get("vid").(int))
	status := d.Get("status").(string)
	description := d.Get("description").(string)

	data.Name = &name
	data.Vid = &vid
	data.Status = status
	data.Description = description

	if siteID, ok := d.GetOk("site_id"); ok {
		data.Site = int64ToPtr(int64(siteID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if roleID, ok := d.GetOk("role_id"); ok {
		data.Role = int64ToPtr(int64(roleID.(int)))
	}

	params := ipam.NewIpamVlansCreateParams().WithData(&data)
	res, err := api.Ipam.IpamVlansCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVlanRead(d, m)
}

func resourceNetboxVlanRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamVlansReadParams().WithID(id)

	res, err := api.Ipam.IpamVlansRead(params, nil)
	if err != nil {
		return err
	}

	vlan := res.GetPayload()

	if err := d.Set("name", vlan.Name); err != nil {
		return err
	}
	if err := d.Set("vid", vlan.Vid); err != nil {
		return err
	}
	if err := d.Set("description", vlan.Description); err != nil {
		return err
	}

	if vlan.Status != nil {
		if err := d.Set("status", vlan.Status.Value); err != nil {
			return err
		}
	}
	if vlan.Site != nil {
		if err := d.Set("site_id", vlan.Site.ID); err != nil {
			return err
		}
	}
	if vlan.Tenant != nil {
		if err := d.Set("tenant_id", vlan.Tenant.ID); err != nil {
			return err
		}
	}
	if vlan.Role != nil {
		if err := d.Set("role_id", vlan.Role.ID); err != nil {
			return err
		}
	}

	return nil
}

func resourceNetboxVlanUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableVLAN{}
	name := d.Get("name").(string)
	vid := int64(d.Get("vid").(int))
	status := d.Get("status").(string)
	description := d.Get("description").(string)

	data.Name = &name
	data.Vid = &vid
	data.Status = status
	data.Description = description

	if siteID, ok := d.GetOk("site_id"); ok {
		data.Site = int64ToPtr(int64(siteID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if roleID, ok := d.GetOk("role_id"); ok {
		data.Role = int64ToPtr(int64(roleID.(int)))
	}
	params := ipam.NewIpamVlansUpdateParams().WithID(id).WithData(&data)
	// nolint: errcheck
	if _, err := api.Ipam.IpamVlansUpdate(params, nil); err != nil {
		return err
	}
	return resourceNetboxVlanRead(d, m)
}

func resourceNetboxVlanDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamVlansDeleteParams().WithID(id)
	// nolint: errcheck
	if _, err := api.Ipam.IpamVlansDelete(params, nil); err != nil {
		return err
	}

	return nil
}
