package netbox

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/virtualization"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxInterface() *schema.Resource {
	validModes := []string{"access", "tagged", "tagged-all"}

	return &schema.Resource{
		CreateContext: resourceNetboxInterfaceCreate,
		ReadContext:   resourceNetboxInterfaceRead,
		UpdateContext: resourceNetboxInterfaceUpdate,
		DeleteContext: resourceNetboxInterfaceDelete,

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#interfaces):

> Virtual machine interfaces behave similarly to device interfaces, and can be assigned to VRFs, and may have IP addresses, VLANs, and services attached to them. However, given their virtual nature, they lack properties pertaining to physical attributes. For example, VM interfaces do not have a physical type and cannot have cables attached to them.`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_machine_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"mac_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsMACAddress,
			},
			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(validModes, false),
			},
			"mtu": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 65536),
			},
			"type": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "This attribute is not supported by netbox any longer. It will be removed in future versions of this provider.",
			},
			tagsKey: tagsSchema,
			"tagged_vlans": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"untagged_vlan": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxInterfaceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	mode := d.Get("mode").(string)
	taggedVlans := toInt64List(d.Get("tagged_vlans"))
	virtualMachineID := int64(d.Get("virtual_machine_id").(int))

	data := models.WritableVMInterface{
		Name:           &name,
		Description:    description,
		Enabled:        enabled,
		Mode:           mode,
		TaggedVlans:    taggedVlans,
		VirtualMachine: &virtualMachineID,
	}
	if macAddress := d.Get("mac_address").(string); macAddress != "" {
		data.MacAddress = &macAddress
	}
	if mtu, ok := d.Get("mtu").(int); ok && mtu != 0 {
		data.Mtu = int64ToPtr(int64(mtu))
	}
	if untaggedVlan, ok := d.Get("untagged_vlan").(int); ok && untaggedVlan != 0 {
		data.UntaggedVlan = int64ToPtr(int64(untaggedVlan))
	}
	params := virtualization.NewVirtualizationInterfacesCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationInterfacesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxInterfaceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	params := virtualization.NewVirtualizationInterfacesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationInterfacesRead(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	iface := res.GetPayload()

	if err := d.Set("name", iface.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", iface.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", iface.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mac_address", iface.MacAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mtu", iface.Mtu); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tagged_vlans", getIDsFromNestedVLAN(iface.TaggedVlans)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("virtual_machine_id", iface.VirtualMachine.ID); err != nil {
		return diag.FromErr(err)
	}

	if iface.Mode != nil {
		if err := d.Set("mode", iface.Mode.Value); err != nil {
			return diag.FromErr(err)
		}
	}
	if iface.UntaggedVlan != nil {
		if err := d.Set("untagged_vlan", iface.UntaggedVlan.ID); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceNetboxInterfaceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	mode := d.Get("mode").(string)
	taggedVlans := toInt64List(d.Get("tagged_vlans"))
	virtualMachineID := int64(d.Get("virtual_machine_id").(int))

	data := models.WritableVMInterface{
		Name:           &name,
		Description:    description,
		Enabled:        enabled,
		Mode:           mode,
		TaggedVlans:    taggedVlans,
		VirtualMachine: &virtualMachineID,
	}

	if d.HasChange("mac_address") {
		macAddress := d.Get("mac_address").(string)
		data.MacAddress = &macAddress
	}
	if d.HasChange("mtu") {
		mtu := int64(d.Get("mtu").(int))
		data.Mtu = &mtu
	}
	if d.HasChange("untagged_vlan") {
		untaggedvlan := int64(d.Get("untagged_vlan").(int))
		data.UntaggedVlan = &untaggedvlan
	}

	params := virtualization.NewVirtualizationInterfacesPartialUpdateParams().WithID(id).WithData(&data)
	// nolint: errcheck
	if _, err := api.Virtualization.VirtualizationInterfacesPartialUpdate(params, nil); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxInterfaceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}
	params := virtualization.NewVirtualizationInterfacesDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Virtualization.VirtualizationInterfacesDelete(params, nil); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func getIDsFromNestedVLAN(nestedvlans []*models.NestedVLAN) []int64 {
	vlans := make([]int64, 0)
	for _, vlan := range nestedvlans {
		vlans = append(vlans, vlan.ID)
	}
	return vlans
}
