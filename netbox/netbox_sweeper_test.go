package netbox

import (
	"os"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
)

// sweeperNetboxClients is a shared cache of netbox clients
// This prevents client re-initialization for every resource with no benefit
var sweeperNetboxClients map[string]interface{}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// sharedClientForRegion returns a common provider client configured for the specified region
func sharedClientForRegion(region string) (interface{}, error) {
	if client, ok := sweeperNetboxClients[region]; ok {
		return client, nil
	}

	server := os.Getenv("NETBOX_SERVER")
	apiToken := os.Getenv("NETBOX_API_TOKEN")
	transport := httptransport.New(server, client.DefaultBasePath, []string{"http"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+apiToken)
	c := client.New(transport, nil)

	return c, nil
}
