package netbox

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/dcim"
)

func dataSourceNetboxDeviceType() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDeviceTypeRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Schema: map[string]*schema.Schema{
			"is_full_depth": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"manufacturer": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"model": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"u_height": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxDeviceTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := dcim.NewDcimDeviceTypesListParams()

	params.Limit = int64ToPtr(2)
	if manufacturer, ok := d.Get("manufacturer").(string); ok && manufacturer != "" {
		params.Manufacturer = &manufacturer
	}
	if model, ok := d.Get("model").(string); ok && model != "" {
		params.Model = &model
	}
	if part, ok := d.Get("part_number").(string); ok && part != "" {
		params.PartNumber = &part
	}
	if slug, ok := d.Get("slug").(string); ok && slug != "" {
		params.Slug = &slug
	}

	res, err := api.Dcim.DcimDeviceTypesList(params, nil)
	if err != nil {
		return err
	}
	if count := *res.GetPayload().Count; count != int64(1) {
		return fmt.Errorf("expected one device type, but got %d", count)
	}

	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	if err := d.Set("is_full_depth", result.IsFullDepth); err != nil {
		return err
	}
	if err := d.Set("manufacturer_id", result.Manufacturer.ID); err != nil {
		return err
	}
	if err := d.Set("model", result.Model); err != nil {
		return err
	}
	if err := d.Set("part_number", result.PartNumber); err != nil {
		return err
	}
	if err := d.Set("slug", result.Slug); err != nil {
		return err
	}
	return d.Set("u_height", result.UHeight)
}
