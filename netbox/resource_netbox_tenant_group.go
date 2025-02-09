package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/tenancy"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxTenantGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxTenantGroupCreate,
		Read:   resourceNetboxTenantGroupRead,
		Update: resourceNetboxTenantGroupUpdate,
		Delete: resourceNetboxTenantGroupDelete,

		Description: `:meta:subcategory:Tenancy:From the [official documentation](https://docs.netbox.dev/en/stable/features/tenancy/#tenant-groups):

> Tenants can be organized by custom groups. For instance, you might create one group called "Customers" and one called "Departments." The assignment of a tenant to a group is optional.
>
> Tenant groups may be nested recursively to achieve a multi-level hierarchy. For example, you might have a group called "Customers" containing subgroups of individual tenants grouped by product or account team.`,

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
			"parent_id": {
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

func resourceNetboxTenantGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	parentID := int64(d.Get("parent_id").(int))
	description := d.Get("description").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data := &models.WritableTenantGroup{}
	data.Name = &name
	data.Slug = &slug
	data.Description = description

	if parentID != 0 {
		data.Parent = &parentID
	}

	params := tenancy.NewTenancyTenantGroupsCreateParams().WithData(data)

	res, err := api.Tenancy.TenancyTenantGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxTenantGroupRead(d, m)
}

func resourceNetboxTenantGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	params := tenancy.NewTenancyTenantGroupsReadParams().WithID(id)

	res, err := api.Tenancy.TenancyTenantGroupsRead(params, nil)
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
	if res.GetPayload().Parent != nil {
		return d.Set("parent", res.GetPayload().Parent.ID)
	}
	return nil
}

func resourceNetboxTenantGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableTenantGroup{}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	parentID := int64(d.Get("parent_id").(int))

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

	if parentID != 0 {
		data.Parent = &parentID
	}
	params := tenancy.NewTenancyTenantGroupsPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Tenancy.TenancyTenantGroupsPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxTenantGroupRead(d, m)
}

func resourceNetboxTenantGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := tenancy.NewTenancyTenantGroupsDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Tenancy.TenancyTenantGroupsDelete(params, nil); err != nil {
		return err
	}
	return nil
}
