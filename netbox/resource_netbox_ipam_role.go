package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/ipam"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxIpamRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIpamRoleCreate,
		Read:   resourceNetboxIpamRoleRead,
		Update: resourceNetboxIpamRoleUpdate,
		Delete: resourceNetboxIpamRoleDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#prefixvlan-roles):

> A role indicates the function of a prefix or VLAN. For example, you might define Data, Voice, and Security roles. Generally, a prefix will be assigned the same functional role as the VLAN to which it is assigned (if any).`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"weight": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 32767),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxIpamRoleCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.Role{}

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}
	weight := int64(d.Get("weight").(int))
	description := d.Get("description").(string)

	data.Name = &name
	data.Slug = &slug

	data.Weight = &weight
	data.Description = description

	params := ipam.NewIpamRolesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamRolesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxIpamRoleUpdate(d, m)
}

func resourceNetboxIpamRoleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamRolesReadParams().WithID(id)

	res, err := api.Ipam.IpamRolesRead(params, nil)
	if err != nil {
		return err
	}

	if res.GetPayload().Name != nil {
		if err := d.Set("name", res.GetPayload().Name); err != nil {
			return err
		}
	}

	if res.GetPayload().Slug != nil {
		if err := d.Set("slug", res.GetPayload().Slug); err != nil {
			return err
		}
	}

	if res.GetPayload().Weight != nil {
		if err := d.Set("weight", res.GetPayload().Weight); err != nil {
			return err
		}
	}

	if res.GetPayload().Description != "" {
		return d.Set("description", res.GetPayload().Description)
	}

	return nil
}

func resourceNetboxIpamRoleUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.Role{}

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}
	weight := int64(d.Get("weight").(int))
	description := d.Get("description").(string)

	data.Name = &name
	data.Slug = &slug

	data.Weight = &weight
	data.Description = description

	params := ipam.NewIpamRolesUpdateParams().WithID(id).WithData(&data)
	// nolint: errcheck
	if _, err := api.Ipam.IpamRolesUpdate(params, nil); err != nil {
		return err
	}
	return resourceNetboxIpamRoleRead(d, m)
}

func resourceNetboxIpamRoleDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamRolesDeleteParams().WithID(id)
	// nolint: errcheck
	if _, err := api.Ipam.IpamRolesDelete(params, nil); err != nil {
		return err
	}
	d.SetId("")
	return nil
}
