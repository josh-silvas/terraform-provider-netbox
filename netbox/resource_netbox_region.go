package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/dcim"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxRegion() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxRegionCreate,
		Read:   resourceNetboxRegionRead,
		Update: resourceNetboxRegionUpdate,
		Delete: resourceNetboxRegionDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/sites-and-racks/#regions):

> Sites can be arranged geographically using regions. A region might represent a continent, country, city, campus, or other area depending on your use case. Regions can be nested recursively to construct a hierarchy. For example, you might define several country regions, and within each of those several state or city regions to which sites are assigned.
>
> Each region must have a name that is unique within its parent region, if any.`,

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
			"parent_region_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxRegionCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableRegion{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	}

	parentRegionIDValue, ok := d.GetOk("parent_region_id")
	if ok {
		data.Parent = int64ToPtr(int64(parentRegionIDValue.(int)))
	}

	params := dcim.NewDcimRegionsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimRegionsCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxRegionRead(d, m)
}

func resourceNetboxRegionRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimRegionsReadParams().WithID(id)

	res, err := api.Dcim.DcimRegionsRead(params, nil)

	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}
	if err := d.Set("slug", res.GetPayload().Slug); err != nil {
		return err
	}
	if res.GetPayload().Parent != nil {
		if err := d.Set("parent_region_id", res.GetPayload().Parent.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("parent_region_id", nil); err != nil {
			return err
		}
	}
	return d.Set("description", res.GetPayload().Description)
}

func resourceNetboxRegionUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableRegion{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	}

	parentRegionIDValue, ok := d.GetOk("parent_region_id")
	if ok {
		data.Parent = int64ToPtr(int64(parentRegionIDValue.(int)))
	}

	params := dcim.NewDcimRegionsPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Dcim.DcimRegionsPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxRegionRead(d, m)
}

func resourceNetboxRegionDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimRegionsDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Dcim.DcimRegionsDelete(params, nil); err != nil {
		return err
	}
	return nil
}
