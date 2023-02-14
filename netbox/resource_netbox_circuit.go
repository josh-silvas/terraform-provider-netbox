package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/circuits"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxCircuit() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitCreate,
		Read:   resourceNetboxCircuitRead,
		Update: resourceNetboxCircuitUpdate,
		Delete: resourceNetboxCircuitDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuits_1):

> A communications circuit represents a single physical link connecting exactly two endpoints, commonly referred to as its A and Z terminations. A circuit in NetBox may have zero, one, or two terminations defined. It is common to have only one termination defined when you don't necessarily care about the details of the provider side of the circuit, e.g. for Internet access circuits. Both terminations would likely be modeled for circuits which connect one customer site to another.
>
> Each circuit is associated with a provider and a user-defined type. For example, you might have Internet access circuits delivered to each site by one provider, and private MPLS circuits delivered by another. Each circuit must be assigned a circuit ID, each of which must be unique per provider.`,

		Schema: map[string]*schema.Schema{
			"provider_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"planned", "provisioning", "active", "offline", "deprovisioning", "decommissioning"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCircuitCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableCircuit{}

	cid := d.Get("cid").(string)
	data.Cid = &cid

	data.Status = d.Get("status").(string)

	providerIDValue, ok := d.GetOk("provider_id")
	if ok {
		data.Provider = int64ToPtr(int64(providerIDValue.(int)))
	}

	typeIDValue, ok := d.GetOk("type_id")
	if ok {
		data.Type = int64ToPtr(int64(typeIDValue.(int)))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	data.Tags = []*models.NestedTag{}

	params := circuits.NewCircuitsCircuitsCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsCircuitsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCircuitRead(d, m)
}

func resourceNetboxCircuitRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := circuits.NewCircuitsCircuitsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsCircuitsRead(params, nil)

	if err != nil {
		return err
	}

	if err := d.Set("cid", res.GetPayload().Cid); err != nil {
		return err
	}
	if err := d.Set("status", res.GetPayload().Status.Value); err != nil {
		return err
	}

	if res.GetPayload().Provider != nil {
		if err := d.Set("provider_id", res.GetPayload().Provider.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("provider_id", nil); err != nil {
			return err
		}
	}

	if res.GetPayload().Type != nil {
		if err := d.Set("type_id", res.GetPayload().Type.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("type_id", nil); err != nil {
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

	return nil
}

func resourceNetboxCircuitUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableCircuit{}

	cid := d.Get("cid").(string)
	data.Cid = &cid

	data.Status = d.Get("status").(string)

	providerIDValue, ok := d.GetOk("provider_id")
	if ok {
		data.Provider = int64ToPtr(int64(providerIDValue.(int)))
	}

	typeIDValue, ok := d.GetOk("type_id")
	if ok {
		data.Type = int64ToPtr(int64(typeIDValue.(int)))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	data.Tags = []*models.NestedTag{}

	params := circuits.NewCircuitsCircuitsPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Circuits.CircuitsCircuitsPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxCircuitRead(d, m)
}

func resourceNetboxCircuitDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := circuits.NewCircuitsCircuitsDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Circuits.CircuitsCircuitsDelete(params, nil); err != nil {
		return err
	}
	return nil
}
