package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/circuits"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxCircuitProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitProviderCreate,
		Read:   resourceNetboxCircuitProviderRead,
		Update: resourceNetboxCircuitProviderUpdate,
		Delete: resourceNetboxCircuitProviderDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#providers):

> A circuit provider is any entity which provides some form of connectivity of among sites or organizations within a site. While this obviously includes carriers which offer Internet and private transit service, it might also include Internet exchange (IX) points and even organizations with whom you peer directly. Each circuit within NetBox must be assigned a provider and a circuit ID which is unique to that provider.
>
> Each provider may be assigned an autonomous system number (ASN), an account number, and contact information.`,

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

func resourceNetboxCircuitProviderCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.Provider{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Tags = []*models.NestedTag{}

	params := circuits.NewCircuitsProvidersCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsProvidersCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCircuitProviderRead(d, m)
}

func resourceNetboxCircuitProviderRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := circuits.NewCircuitsProvidersReadParams().WithID(id)

	res, err := api.Circuits.CircuitsProvidersRead(params, nil)
	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}

	return d.Set("slug", res.GetPayload().Slug)
}

func resourceNetboxCircuitProviderUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.Provider{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Tags = []*models.NestedTag{}

	params := circuits.NewCircuitsProvidersPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Circuits.CircuitsProvidersPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxCircuitProviderRead(d, m)
}

func resourceNetboxCircuitProviderDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := circuits.NewCircuitsProvidersDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Circuits.CircuitsProvidersDelete(params, nil); err != nil {
		return err
	}
	return nil
}
