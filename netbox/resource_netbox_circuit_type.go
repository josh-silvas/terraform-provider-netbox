package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/circuits"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxCircuitType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitTypeCreate,
		Read:   resourceNetboxCircuitTypeRead,
		Update: resourceNetboxCircuitTypeUpdate,
		Delete: resourceNetboxCircuitTypeDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuit-types):

> Circuits are classified by functional type. These types are completely customizable, and are typically used to convey the type of service being delivered over a circuit.`,

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

func resourceNetboxCircuitTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.CircuitType{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	params := circuits.NewCircuitsCircuitTypesCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsCircuitTypesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCircuitTypeRead(d, m)
}

func resourceNetboxCircuitTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := circuits.NewCircuitsCircuitTypesReadParams().WithID(id)

	res, err := api.Circuits.CircuitsCircuitTypesRead(params, nil)

	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}

	return d.Set("slug", res.GetPayload().Slug)
}

func resourceNetboxCircuitTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.CircuitType{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	params := circuits.NewCircuitsCircuitTypesPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Circuits.CircuitsCircuitTypesPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxCircuitTypeRead(d, m)
}

func resourceNetboxCircuitTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := circuits.NewCircuitsCircuitTypesDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Circuits.CircuitsCircuitTypesDelete(params, nil); err != nil {
		return err
	}
	return nil
}
