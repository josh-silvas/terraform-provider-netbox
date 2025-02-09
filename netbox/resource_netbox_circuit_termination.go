package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/circuits"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxCircuitTermination() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitTerminationCreate,
		Read:   resourceNetboxCircuitTerminationRead,
		Update: resourceNetboxCircuitTerminationUpdate,
		Delete: resourceNetboxCircuitTerminationDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuit-terminations):

> The association of a circuit with a particular site and/or device is modeled separately as a circuit termination. A circuit may have up to two terminations, labeled A and Z. A single-termination circuit can be used when you don't know (or care) about the far end of a circuit (for example, an Internet access circuit which connects to a transit provider). A dual-termination circuit is useful for tracking circuits which connect two sites.
>
> Each circuit termination is attached to either a site or to a provider network. Site terminations may optionally be connected via a cable to a specific device interface or port within that site. Each termination must be assigned a port speed, and can optionally be assigned an upstream speed if it differs from the downstream speed (a common scenario with e.g. DOCSIS cable modems). Fields are also available to track cross-connect and patch panel details.`,

		Schema: map[string]*schema.Schema{
			"circuit_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"port_speed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upstream_speed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"term_side": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"A", "Z"}, false),
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCircuitTerminationCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableCircuitTermination{}

	termside := d.Get("term_side").(string)
	data.TermSide = &termside

	circuitIDValue, ok := d.GetOk("circuit_id")
	if ok {
		data.Circuit = int64ToPtr(int64(circuitIDValue.(int)))
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		data.Site = int64ToPtr(int64(siteIDValue.(int)))
	}

	portspeedValue, ok := d.GetOk("port_speed")
	if ok {
		data.PortSpeed = int64ToPtr(int64(portspeedValue.(int)))
	}

	upstreamspeedValue, ok := d.GetOk("upstream_speed")
	if ok {
		data.UpstreamSpeed = int64ToPtr(int64(upstreamspeedValue.(int)))
	}

	params := circuits.NewCircuitsCircuitTerminationsCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsCircuitTerminationsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCircuitTerminationRead(d, m)
}

func resourceNetboxCircuitTerminationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := circuits.NewCircuitsCircuitTerminationsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsCircuitTerminationsRead(params, nil)

	if err != nil {
		return err
	}

	term := res.GetPayload()

	if err := d.Set("term_side", term.TermSide); err != nil {
		return err
	}

	if term.Circuit != nil {
		if err := d.Set("circuit_id", term.Circuit.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("circuit_id", nil); err != nil {
			return err
		}
	}

	if term.Site != nil {
		if err := d.Set("site_id", term.Site.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("site_id", nil); err != nil {
			return err
		}
	}

	if term.PortSpeed != nil {
		if err := d.Set("port_speed", term.PortSpeed); err != nil {
			return err
		}
	} else {
		if err := d.Set("port_speed", nil); err != nil {
			return err
		}
	}

	if term.UpstreamSpeed != nil {
		if err := d.Set("upstream_speed", term.UpstreamSpeed); err != nil {
			return err
		}
	} else {
		if err := d.Set("upstream_speed", nil); err != nil {
			return err
		}
	}

	return nil
}

func resourceNetboxCircuitTerminationUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableCircuitTermination{}

	termside := d.Get("term_side").(string)
	data.TermSide = &termside

	circuitIDValue, ok := d.GetOk("circuit_id")
	if ok {
		data.Circuit = int64ToPtr(int64(circuitIDValue.(int)))
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		data.Site = int64ToPtr(int64(siteIDValue.(int)))
	}

	portspeedValue, ok := d.GetOk("port_speed")
	if ok {
		data.PortSpeed = int64ToPtr(int64(portspeedValue.(int)))
	}

	upstreamspeedValue, ok := d.GetOk("upstream_speed")
	if ok {
		data.UpstreamSpeed = int64ToPtr(int64(upstreamspeedValue.(int)))
	}
	params := circuits.NewCircuitsCircuitTerminationsPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Circuits.CircuitsCircuitTerminationsPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxCircuitTerminationRead(d, m)
}

func resourceNetboxCircuitTerminationDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := circuits.NewCircuitsCircuitTerminationsDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Circuits.CircuitsCircuitTerminationsDelete(params, nil); err != nil {
		return err
	}
	return nil
}
