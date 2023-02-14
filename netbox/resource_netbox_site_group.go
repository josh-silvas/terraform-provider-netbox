package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxSiteGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxSiteGroupCreate,
		Read:   resourceNetboxSiteGroupRead,
		Update: resourceNetboxSiteGroupUpdate,
		Delete: resourceNetboxSiteGroupDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/facilities/#site-groups):

> Like regions, site groups can be arranged in a recursive hierarchy for grouping sites. However, whereas regions are intended for geographic organization, site groups may be used for functional grouping. For example, you might classify sites as corporate, branch, or customer sites in addition to where they are physically located.
>
> The use of both regions and site groups affords to independent but complementary dimensions across which sites can be organized.`,

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

func resourceNetboxSiteGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	data := &models.WritableSiteGroup{}
	data.Name = &name
	data.Slug = &slug
	data.Description = description

	if parentID != 0 {
		data.Parent = &parentID
	}

	params := dcim.NewDcimSiteGroupsCreateParams().WithData(data)

	res, err := api.Dcim.DcimSiteGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxSiteGroupRead(d, m)
}

func resourceNetboxSiteGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	params := dcim.NewDcimSiteGroupsReadParams().WithID(id)

	res, err := api.Dcim.DcimSiteGroupsRead(params, nil)
	if err != nil {
		return err
	}

	siteGroup := res.GetPayload()
	if err := d.Set("name", siteGroup.Name); err != nil {
		return err
	}
	if err := d.Set("slug", siteGroup.Slug); err != nil {
		return err
	}
	if err := d.Set("description", siteGroup.Description); err != nil {
		return err
	}
	if siteGroup.Parent != nil {
		if err := d.Set("parent_id", siteGroup.Parent.ID); err != nil {
			return err
		}
	}
	return nil
}

func resourceNetboxSiteGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableSiteGroup{}

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
	params := dcim.NewDcimSiteGroupsPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Dcim.DcimSiteGroupsPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxSiteGroupRead(d, m)
}

func resourceNetboxSiteGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimSiteGroupsDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Dcim.DcimSiteGroupsDelete(params, nil); err != nil {
		return err
	}
	return nil
}
