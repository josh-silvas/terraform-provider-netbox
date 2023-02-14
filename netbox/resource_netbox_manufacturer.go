package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/dcim"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxManufacturer() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxManufacturerCreate,
		Read:   resourceNetboxManufacturerRead,
		Update: resourceNetboxManufacturerUpdate,
		Delete: resourceNetboxManufacturerDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/device-types/#manufacturers):

> A manufacturer represents the "make" of a device; e.g. Cisco or Dell. Each device type must be assigned to a manufacturer. (Inventory items and platforms may also be associated with manufacturers.) Each manufacturer must have a unique name and may have a description assigned to it.`,

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
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxManufacturerCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.Manufacturer{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	params := dcim.NewDcimManufacturersCreateParams().WithData(&data)

	res, err := api.Dcim.DcimManufacturersCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxManufacturerRead(d, m)
}

func resourceNetboxManufacturerRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimManufacturersReadParams().WithID(id)

	res, err := api.Dcim.DcimManufacturersRead(params, nil)

	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}
	return d.Set("slug", res.GetPayload().Slug)
}

func resourceNetboxManufacturerUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.Manufacturer{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	params := dcim.NewDcimManufacturersPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Dcim.DcimManufacturersPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxManufacturerRead(d, m)
}

func resourceNetboxManufacturerDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimManufacturersDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Dcim.DcimManufacturersDelete(params, nil); err != nil {
		return err
	}
	return nil
}
