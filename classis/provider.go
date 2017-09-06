package classis

import (

	//"github.com/hashicorp/terraform/helper/hashcode"
	//"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	//homedir "github.com/mitchellh/go-homedir"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	// TODO: Move the validation to this, requires conditional schemas
	// TODO: Move the configuration to this, requires validation

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_USER", nil),
				Description: "username for classis",
			},
			"password": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_PASSWORD", nil),
				Description: "password for classis",
			},
		},
		ResourcesMap:  map[string]*schema.Resource{},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	//user := d.Get("username").(string)
	//token := d.Get("password").(string)
	//return SimpleFakeApi.New(user, token) //TODO need to create the golang client for classis
	return nil, nil
}
