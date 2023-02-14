package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/users"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxTokenCreate,
		Read:   resourceNetboxTokenRead,
		Update: resourceNetboxTokenUpdate,
		Delete: resourceNetboxTokenDelete,

		Description: `:meta:subcategory:Authentication:From the [official documentation](https://docs.netbox.dev/en/stable/rest-api/authentication/#tokens):

> A token is a unique identifier mapped to a NetBox user account. Each user may have one or more tokens which he or she can use for authentication when making REST API requests. To create a token, navigate to the API tokens page under your user profile.`,

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"key": {
				Type:         schema.TypeString,
				Sensitive:    true,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(40, 256),
			},
			"allowed_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsCIDR,
				},
			},
			"write_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"last_used": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxTokenCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableToken{}

	userid := int64(d.Get("user_id").(int))

	key := d.Get("key").(string)
	allowedIps := d.Get("allowed_ips").([]interface{})

	data.User = &userid
	data.Key = key

	data.AllowedIps = make([]models.IPNetwork, len(allowedIps))
	for i, v := range allowedIps {
		data.AllowedIps[i] = v
	}

	data.WriteEnabled = d.Get("write_enabled").(bool)

	params := users.NewUsersTokensCreateParams().WithData(&data)
	res, err := api.Users.UsersTokensCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxTokenUpdate(d, m)
}

func resourceNetboxTokenRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := users.NewUsersTokensReadParams().WithID(id)

	res, err := api.Users.UsersTokensRead(params, nil)
	if err != nil {
		// nolint: errorlint
		errorcode := err.(*users.UsersTokensReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}
	token := res.GetPayload()

	if token.User != nil {
		if err := d.Set("user_id", token.User.ID); err != nil {
			return err
		}
	}

	if err := d.Set("key", token.Key); err != nil {
		return err
	}
	if err := d.Set("last_used", token.LastUsed); err != nil {
		return err
	}
	if err := d.Set("expires", token.Expires); err != nil {
		return err
	}
	if err := d.Set("allowed_ips", token.AllowedIps); err != nil {
		return err
	}
	if err := d.Set("write_enabled", token.WriteEnabled); err != nil {
		return err
	}

	return nil
}

func resourceNetboxTokenUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableToken{}

	userid := int64(d.Get("user_id").(int))
	key := d.Get("key").(string)
	allowedIps := d.Get("allowed_ips").([]interface{})

	data.User = &userid
	data.Key = key

	data.AllowedIps = make([]models.IPNetwork, len(allowedIps))
	for i, v := range allowedIps {
		data.AllowedIps[i] = v
	}

	data.WriteEnabled = d.Get("write_enabled").(bool)

	params := users.NewUsersTokensUpdateParams().WithID(id).WithData(&data)
	// nolint: errcheck
	if _, err := api.Users.UsersTokensUpdate(params, nil); err != nil {
		return err
	}
	return resourceNetboxTokenRead(d, m)
}

func resourceNetboxTokenDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := users.NewUsersTokensDeleteParams().WithID(id)
	// nolint: errcheck
	if _, err := api.Users.UsersTokensDelete(params, nil); err != nil {
		return err
	}
	d.SetId("")
	return nil
}
