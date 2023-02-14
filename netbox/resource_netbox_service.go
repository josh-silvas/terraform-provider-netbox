package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/ipam"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxService() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxServiceCreate,
		Read:   resourceNetboxServiceRead,
		Update: resourceNetboxServiceUpdate,
		Delete: resourceNetboxServiceDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/services/#services):

> A service represents a layer four TCP or UDP service available on a device or virtual machine. For example, you might want to document that an HTTP service is running on a device. Each service includes a name, protocol, and port number; for example, "SSH (TCP/22)" or "DNS (UDP/53)."
>
> A service may optionally be bound to one or more specific IP addresses belonging to its parent device or VM. (If no IP addresses are bound, the service is assumed to be reachable via any assigned IP address.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"virtual_machine_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"tcp", "udp", "sctp"}, false)),
			},
			"port": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"port", "ports"},
				Deprecated:   "This field is deprecated. Please use the new \"ports\" attribute instead.",
			},
			"ports": {
				Type:         schema.TypeSet,
				Optional:     true,
				ExactlyOneOf: []string{"port", "ports"},
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxServiceCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableService{}

	dataName := d.Get("name").(string)
	data.Name = &dataName

	dataProtocol := d.Get("protocol").(string)
	data.Protocol = &dataProtocol

	// for backwards compatibility, we allow either port or ports
	// the API only supports ports. We give precedence to port, if it exists.
	//dataPort := int64(d.Get("port").(int))
	dataPort, dataPortOk := d.GetOk("port")
	if dataPortOk {
		data.Ports = []int64{int64(dataPort.(int))}
	} else {
		// if port is not set, ports has to be set
		var dataPorts []int64
		if v := d.Get("ports").(*schema.Set); v.Len() > 0 {
			for _, v := range v.List() {
				dataPorts = append(dataPorts, int64(v.(int)))
			}
			data.Ports = dataPorts
		}
	}

	dataVirtualMachineID := int64(d.Get("virtual_machine_id").(int))
	data.VirtualMachine = &dataVirtualMachineID

	data.Tags = []*models.NestedTag{}
	data.Ipaddresses = []int64{}

	params := ipam.NewIpamServicesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamServicesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxServiceUpdate(d, m)
}

func resourceNetboxServiceRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamServicesReadParams().WithID(id)

	res, err := api.Ipam.IpamServicesRead(params, nil)
	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}
	if err := d.Set("protocol", res.GetPayload().Protocol.Value); err != nil {
		return err
	}
	if err := d.Set("ports", res.GetPayload().Ports); err != nil {
		return err
	}
	if err := d.Set("virtual_machine_id", res.GetPayload().VirtualMachine.ID); err != nil {
		return err
	}

	return nil
}

func resourceNetboxServiceUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableService{}

	dataName := d.Get("name").(string)
	data.Name = &dataName

	dataProtocol := d.Get("protocol").(string)
	data.Protocol = &dataProtocol

	dataPort, dataPortOk := d.GetOk("port")
	if dataPortOk {
		data.Ports = []int64{int64(dataPort.(int))}
	} else {
		// if port is not set, ports has to be set
		var dataPorts []int64
		if v := d.Get("ports").(*schema.Set); v.Len() > 0 {
			for _, v := range v.List() {
				dataPorts = append(dataPorts, int64(v.(int)))
			}
			data.Ports = dataPorts
		}
	}

	data.Tags = []*models.NestedTag{}
	data.Ipaddresses = []int64{}

	dataVirtualMachineID := int64(d.Get("virtual_machine_id").(int))
	data.VirtualMachine = &dataVirtualMachineID

	params := ipam.NewIpamServicesUpdateParams().WithID(id).WithData(&data)
	// nolint: errcheck
	if _, err := api.Ipam.IpamServicesUpdate(params, nil); err != nil {
		return err
	}
	return resourceNetboxServiceRead(d, m)
}

func resourceNetboxServiceDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamServicesDeleteParams().WithID(id)
	// nolint: errcheck
	if _, err := api.Ipam.IpamServicesDelete(params, nil); err != nil {
		return err
	}
	return nil
}
