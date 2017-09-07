package classis

import (
	"github.com/classis/terraform-provider-classis/classis/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	// TODO: Move the validation to this, requires conditional schemas
	// TODO: Move the configuration to this, requires validation

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLASSIS_URL", nil),
				Description: "url for classis",
			},
			"email": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLASSIS_EMAIL", nil),
				Description: "email for classis",
			},
			"password": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLASSIS_PASSWORD", nil),
				Description: "password for classis",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"classis_aws_spot_group": resourceAwsSpotGroupObject(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	return classis.NewClientWith(url, email, password)
}
