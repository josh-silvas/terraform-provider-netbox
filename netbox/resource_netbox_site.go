package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxSiteCreate,
		Read:   resourceNetboxSiteRead,
		Update: resourceNetboxSiteUpdate,
		Delete: resourceNetboxSiteDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/sites-and-racks/#sites):

> How you choose to employ sites when modeling your network may vary depending on the nature of your organization, but generally a site will equate to a building or campus. For example, a chain of banks might create a site to represent each of its branches, a site for its corporate headquarters, and two additional sites for its presence in two colocation facilities.
>
> Each site must be assigned a unique name and may optionally be assigned to a region and/or tenant.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 100),
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "active",
				ValidateFunc: validation.StringInSlice([]string{"planned", "staging", "active", "decommissioning", "retired"}, false),
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"facility": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"latitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"physical_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"shipping_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"region_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			tagsKey: tagsSchema,
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asn_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxSiteCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableSite{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Status = d.Get("status").(string)

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	}

	if facility, ok := d.GetOk("facility"); ok {
		data.Facility = facility.(string)
	}

	latitudeValue, ok := d.GetOk("latitude")
	if ok {
		data.Latitude = float64ToPtr(latitudeValue.(float64))
	}

	longitudeValue, ok := d.GetOk("longitude")
	if ok {
		data.Longitude = float64ToPtr(longitudeValue.(float64))
	}

	physicalAddressValue, ok := d.GetOk("physical_address")
	if ok {
		data.PhysicalAddress = physicalAddressValue.(string)
	}

	shippingAddressValue, ok := d.GetOk("shipping_address")
	if ok {
		data.ShippingAddress = shippingAddressValue.(string)
	}

	regionIDValue, ok := d.GetOk("region_id")
	if ok {
		data.Region = int64ToPtr(int64(regionIDValue.(int)))
	}

	groupIDValue, ok := d.GetOk("group_id")
	if ok {
		data.Group = int64ToPtr(int64(groupIDValue.(int)))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	if timezone, ok := d.GetOk("timezone"); ok {
		data.TimeZone = timezone.(string)
	}

	data.Asns = make([]int64, 0)
	if asnsValue, ok := d.GetOk("asn_ids"); ok {
		data.Asns = toInt64List(asnsValue)
	}

	data.Tags = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimSitesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimSitesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxSiteRead(d, m)
}

func resourceNetboxSiteRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimSitesReadParams().WithID(id)

	res, err := api.Dcim.DcimSitesRead(params, nil)

	if err != nil {
		return err
	}

	site := res.GetPayload()

	if err := d.Set("name", site.Name); err != nil {
		return err
	}
	if err := d.Set("slug", site.Slug); err != nil {
		return err
	}
	if err := d.Set("status", site.Status.Value); err != nil {
		return err
	}
	if err := d.Set("description", site.Description); err != nil {
		return err
	}
	if err := d.Set("facility", site.Facility); err != nil {
		return err
	}
	if err := d.Set("longitude", site.Longitude); err != nil {
		return err
	}
	if err := d.Set("latitude", site.Latitude); err != nil {
		return err
	}
	if err := d.Set("physical_address", site.PhysicalAddress); err != nil {
		return err
	}
	if err := d.Set("shipping_address", site.ShippingAddress); err != nil {
		return err
	}
	if err := d.Set("timezone", site.TimeZone); err != nil {
		return err
	}
	if err := d.Set("asn_ids", site.Asns); err != nil {
		return err
	}

	if res.GetPayload().Region != nil {
		if err := d.Set("region_id", res.GetPayload().Region.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("region_id", nil); err != nil {
			return err
		}
	}

	if res.GetPayload().Group != nil {
		if err := d.Set("group_id", res.GetPayload().Group.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("group_id", nil); err != nil {
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

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		return d.Set(customFieldsKey, cf)
	}
	return d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))
}

func resourceNetboxSiteUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableSite{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Status = d.Get("status").(string)

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	} else if d.HasChange("description") {
		// If GetOK returned unset description and its value changed, set it as a space string to delete it ...
		data.Description = " "
	}

	if facility, ok := d.GetOk("facility"); ok {
		data.Facility = facility.(string)
	}

	latitudeValue, ok := d.GetOk("latitude")
	if ok {
		data.Latitude = float64ToPtr(latitudeValue.(float64))
	}

	longitudeValue, ok := d.GetOk("longitude")
	if ok {
		data.Longitude = float64ToPtr(longitudeValue.(float64))
	}

	physicalAddressValue, ok := d.GetOk("physical_address")
	if ok {
		data.PhysicalAddress = physicalAddressValue.(string)
	} else if d.HasChange("physical_address") {
		// If GetOK returned unset description and its value changed, set it as a space string to delete it ...
		data.PhysicalAddress = " "
	}

	shippingAddressValue, ok := d.GetOk("shipping_address")
	if ok {
		data.ShippingAddress = shippingAddressValue.(string)
	} else if d.HasChange("shipping_address") {
		// If GetOK returned unset description and its value changed, set it as a space string to delete it ...
		data.ShippingAddress = " "
	}

	regionIDValue, ok := d.GetOk("region_id")
	if ok {
		data.Region = int64ToPtr(int64(regionIDValue.(int)))
	}

	groupIDValue, ok := d.GetOk("group_id")
	if ok {
		data.Group = int64ToPtr(int64(groupIDValue.(int)))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	if timezone, ok := d.GetOk("timezone"); ok {
		data.TimeZone = timezone.(string)
	}

	data.Asns = make([]int64, 0)
	if asnsValue, ok := d.GetOk("asn_ids"); ok {
		data.Asns = toInt64List(asnsValue)
	}

	data.Tags = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := dcim.NewDcimSitesPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Dcim.DcimSitesPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxSiteRead(d, m)
}

func resourceNetboxSiteDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimSitesDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Dcim.DcimSitesDelete(params, nil); err != nil {
		return err
	}
	return nil
}
