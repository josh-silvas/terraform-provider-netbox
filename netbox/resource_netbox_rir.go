package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxRir() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxRirCreate,
		Read:   resourceNetboxRirRead,
		Update: resourceNetboxRirUpdate,
		Delete: resourceNetboxRirDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#regional-internet-registries-rirs):

> Regional Internet registries are responsible for the allocation of globally-routable address space. The five RIRs are ARIN, RIPE, APNIC, LACNIC, and AFRINIC. However, some address space has been set aside for internal use, such as defined in RFCs 1918 and 6598. NetBox considers these RFCs as a sort of RIR as well; that is, an authority which "owns" certain address space. There also exist lower-tier registries which serve particular geographic areas.`,

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
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxRirCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.RIR{}

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data.Name = &name
	data.Slug = &slug
	data.Tags = []*models.NestedTag{}

	params := ipam.NewIpamRirsCreateParams().WithData(&data)
	res, err := api.Ipam.IpamRirsCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxRirUpdate(d, m)
}

func resourceNetboxRirRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamRirsReadParams().WithID(id)

	res, err := api.Ipam.IpamRirsRead(params, nil)
	if err != nil {
		// nolint: errorlint
		errorcode := err.(*ipam.IpamRirsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
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

	return nil
}

func resourceNetboxRirUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.RIR{}

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data.Name = &name
	data.Slug = &slug
	data.Tags = []*models.NestedTag{}

	params := ipam.NewIpamRirsUpdateParams().WithID(id).WithData(&data)
	// nolint: errcheck
	if _, err := api.Ipam.IpamRirsUpdate(params, nil); err != nil {
		return err
	}
	return resourceNetboxRirRead(d, m)
}

func resourceNetboxRirDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamRirsDeleteParams().WithID(id)
	// nolint: errcheck
	if _, err := api.Ipam.IpamRirsDelete(params, nil); err != nil {
		return err
	}
	d.SetId("")
	return nil
}
