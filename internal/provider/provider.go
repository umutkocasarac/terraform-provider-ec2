package provider

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

// New asd
func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"profile": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("aws_profile", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"scaffolding_data_source": dataSourceScaffolding(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"scaffolding_resource": resourceScaffolding(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	client *ec2.Client
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {

	return func(c context.Context, s *schema.ResourceData) (interface{}, diag.Diagnostics) {
		profile := s.Get("profile").(string)
		log.Printf("[INFO] umut provider configuration is working. profile is %v", profile)
		ctx := context.TODO()
		cfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion("us-west-2"),
			config.WithSharedConfigProfile(profile),
		)
		if err != nil {
			panic("configuration error, " + err.Error())
		}
		client := ec2.NewFromConfig(cfg)
		return &apiClient{client}, nil
	}
}
