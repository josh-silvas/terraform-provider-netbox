package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/dcim"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxDeviceType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceTypeCreate,
		Read:   resourceNetboxDeviceTypeRead,
		Update: resourceNetboxDeviceTypeUpdate,
		Delete: resourceNetboxDeviceTypeDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/device-types/#device-types_1):

> A device type represents a particular make and model of hardware that exists in the real world. Device types define the physical attributes of a device (rack height and depth) and its individual components (console, power, network interfaces, and so on).`,

		Schema: map[string]*schema.Schema{
			"model": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 30),
			},
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"u_height": {
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "1.0",
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableDeviceType{}

	model := d.Get("model").(string)
	data.Model = &model

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(model))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	manufacturerIDValue, ok := d.GetOk("manufacturer_id")
	if ok {
		data.Manufacturer = int64ToPtr(int64(manufacturerIDValue.(int)))
	}

	if partNo, ok := d.GetOk("part_number"); ok {
		data.PartNumber = partNo.(string)
	}

	if uHeightValue, ok := d.GetOk("u_height"); ok {
		data.UHeight = int64ToPtr(uHeightValue.(int64))
	}

	params := dcim.NewDcimDeviceTypesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimDeviceTypesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceTypeRead(d, m)
}

func resourceNetboxDeviceTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimDeviceTypesReadParams().WithID(id)

	res, err := api.Dcim.DcimDeviceTypesRead(params, nil)

	if err != nil {
		return err
	}

	deviceType := res.GetPayload()
	if err := d.Set("model", deviceType.Model); err != nil {
		return err
	}
	if err := d.Set("slug", deviceType.Slug); err != nil {
		return err
	}
	if err := d.Set("manufacturer_id", deviceType.Manufacturer.ID); err != nil {
		return err
	}
	if err := d.Set("part_number", deviceType.PartNumber); err != nil {
		return err
	}
	if err := d.Set("u_height", deviceType.UHeight); err != nil {
		return err
	}

	return nil
}

func resourceNetboxDeviceTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableDeviceType{}

	model := d.Get("model").(string)
	data.Model = &model

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(model))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	manufacturerIDValue, ok := d.GetOk("manufacturer_id")
	if ok {
		data.Manufacturer = int64ToPtr(int64(manufacturerIDValue.(int)))
	}

	if partNo, ok := d.GetOk("part_number"); ok {
		data.PartNumber = partNo.(string)
	}

	if uHeightValue, ok := d.GetOk("u_height"); ok {
		data.UHeight = int64ToPtr(uHeightValue.(int64))
	}

	params := dcim.NewDcimDeviceTypesPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Dcim.DcimDeviceTypesPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxDeviceTypeRead(d, m)
}

func resourceNetboxDeviceTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimDeviceTypesDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Dcim.DcimDeviceTypesDelete(params, nil); err != nil {
		return err
	}
	return nil
}
