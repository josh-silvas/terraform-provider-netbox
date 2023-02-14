package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/tenancy"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxTenantCreate,
		Read:   resourceNetboxTenantRead,
		Update: resourceNetboxTenantUpdate,
		Delete: resourceNetboxTenantDelete,

		Description: `:meta:subcategory:Tenancy:From the [official documentation](https://docs.netbox.dev/en/stable/features/tenancy/#tenants):

> A tenant represents a discrete grouping of resources used for administrative purposes. Typically, tenants are used to represent individual customers or internal departments within an organization.
>
> Tenant assignment is used to signify the ownership of an object in NetBox. As such, each object may only be owned by a single tenant. For example, if you have a firewall dedicated to a particular customer, you would assign it to the tenant which represents that customer. However, if the firewall serves multiple customers, it doesn't belong to any particular customer, so tenant assignment would not be appropriate.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 30),
			},
			tagsKey: tagsSchema,
			"group_id": {
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceNetboxTenantCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	groupID := int64(d.Get("group_id").(int))
	description := d.Get("description").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}
	data := &models.WritableTenant{}

	data.Name = &name
	data.Slug = &slug
	data.Description = description

	if groupID != 0 {
		data.Group = &groupID
	}

	params := tenancy.NewTenancyTenantsCreateParams().WithData(data)

	res, err := api.Tenancy.TenancyTenantsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxTenantRead(d, m)
}

func resourceNetboxTenantRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := tenancy.NewTenancyTenantsReadParams().WithID(id)

	res, err := api.Tenancy.TenancyTenantsRead(params, nil)
	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}
	if err := d.Set("slug", res.GetPayload().Slug); err != nil {
		return err
	}
	if err := d.Set("description", res.GetPayload().Description); err != nil {
		return err
	}
	if res.GetPayload().Group != nil {
		return d.Set("group_id", res.GetPayload().Group.ID)
	}

	return nil
}

func resourceNetboxTenantUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableTenant{}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	groupID := int64(d.Get("group_id").(int))
	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}
	data.Slug = &slug
	data.Name = &name
	data.Description = description
	if groupID != 0 {
		data.Group = &groupID
	}

	params := tenancy.NewTenancyTenantsPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Tenancy.TenancyTenantsPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxTenantRead(d, m)
}

func resourceNetboxTenantDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := tenancy.NewTenancyTenantsDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Tenancy.TenancyTenantsDelete(params, nil); err != nil {
		return err
	}
	return nil
}
