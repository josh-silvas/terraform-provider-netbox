package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/users"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxUserCreate,
		Read:   resourceNetboxUserRead,
		Update: resourceNetboxUserUpdate,
		Delete: resourceNetboxUserDelete,

		Description: `:meta:subcategory:Authentication:This resource is used to manage users.`,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"staff": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxUserCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableUser{}

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	active := d.Get("active").(bool)
	staff := d.Get("staff").(bool)

	data.Username = &username
	data.Password = &password
	data.IsActive = active
	data.IsStaff = staff

	data.Groups = []int64{}

	params := users.NewUsersUsersCreateParams().WithData(&data)
	res, err := api.Users.UsersUsersCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxUserRead(d, m)
}

func resourceNetboxUserRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := users.NewUsersUsersReadParams().WithID(id)

	res, err := api.Users.UsersUsersRead(params, nil)
	if err != nil {
		return err
	}

	if res.GetPayload().Username != nil {
		if err := d.Set("username", res.GetPayload().Username); err != nil {
			return err
		}
	}

	if err := d.Set("staff", res.GetPayload().IsStaff); err != nil {
		return err
	}

	return d.Set("active", res.GetPayload().IsActive)
}

func resourceNetboxUserUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableUser{}

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	active := d.Get("active").(bool)
	staff := d.Get("staff").(bool)

	data.Username = &username
	data.Password = &password
	data.IsActive = active
	data.IsStaff = staff

	data.Groups = []int64{}

	params := users.NewUsersUsersUpdateParams().WithID(id).WithData(&data)
	// nolint: errcheck
	if _, err := api.Users.UsersUsersUpdate(params, nil); err != nil {
		return err
	}
	return resourceNetboxUserRead(d, m)
}

func resourceNetboxUserDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := users.NewUsersUsersDeleteParams().WithID(id)
	// nolint: errcheck
	if _, err := api.Users.UsersUsersDelete(params, nil); err != nil {
		return err
	}
	d.SetId("")
	return nil
}
