package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/dcim"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxLocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxLocationCreate,
		Read:   resourceNetboxLocationRead,
		Update: resourceNetboxLocationUpdate,
		Delete: resourceNetboxLocationDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/sites-and-racks/#locations):

> Racks and devices can be grouped by location within a site. A location may represent a floor, room, cage, or similar organizational unit. Locations can be nested to form a hierarchy. For example, you may have floors within a site, and rooms within a floor.

Each location must have a name that is unique within its parent site and location, if any.`,

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
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxLocationCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableLocation{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		data.Site = int64ToPtr(int64(siteIDValue.(int)))
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimLocationsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimLocationsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxLocationRead(d, m)
}

func resourceNetboxLocationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimLocationsReadParams().WithID(id)

	res, err := api.Dcim.DcimLocationsRead(params, nil)

	if err != nil {
		return err
	}

	location := res.GetPayload()

	if err := d.Set("name", location.Name); err != nil {
		return err
	}
	if err := d.Set("slug", location.Slug); err != nil {
		return err
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

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		if err := d.Set(customFieldsKey, cf); err != nil {
			return err
		}
	}

	return nil
}

func resourceNetboxLocationUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableLocation{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		data.Site = int64ToPtr(int64(siteIDValue.(int)))
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := dcim.NewDcimLocationsPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Dcim.DcimLocationsPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxLocationRead(d, m)
}

func resourceNetboxLocationDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimLocationsDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Dcim.DcimLocationsDelete(params, nil); err != nil {
		return err
	}
	return nil
}
