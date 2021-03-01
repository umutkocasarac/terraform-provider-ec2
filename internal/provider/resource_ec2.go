package provider

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaffolding() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceScaffoldingCreate,
		ReadContext:   resourceScaffoldingRead,
		UpdateContext: resourceScaffoldingUpdate,
		DeleteContext: resourceScaffoldingDelete,

		Schema: map[string]*schema.Schema{
			"amount": {
				// This description is used by the documentation generator and the language server.
				Description: "Amount attribute.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"ami": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associate_public_ip_address": {
				Type:     schema.TypeBool,
				ForceNew: true,
				Computed: true,
				Optional: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cpu_threads_per_core": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceScaffoldingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] umut resourceScaffoldingCreate is working")
	amount := d.Get("amount").(int)
	ami := d.Get("ami").(string)
	client := meta.(*apiClient)
	opts := &ec2.RunInstancesInput{
		ImageId:      aws.String(ami),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     int32(amount),
		MaxCount:     int32(amount),
	}
	reservation, err := client.client.RunInstances(context.TODO(), opts)
	if err != nil {
		return diag.FromErr(err)
	}
	instanceID := *reservation.Instances[0].InstanceId
	d.SetId(instanceID)
	return resourceScaffoldingRead(ctx, d, meta)
}

func resourceScaffoldingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] umut resourceScaffoldingRead is working")
	var diags diag.Diagnostics
	client := meta.(*apiClient)
	instanceID := d.Id()
	instance, err := resourceAwsInstanceFindByID(client, instanceID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("ami", instance.ImageId)
	d.Set("instance_type", instance.InstanceType)
	if instance.State != nil {
		d.Set("instance_state", instance.State.Name)
	}
	if instance.Placement != nil {
		d.Set("availability_zone", instance.Placement.AvailabilityZone)
	}
	if instance.CpuOptions != nil {
		d.Set("cpu_core_count", instance.CpuOptions.CoreCount)
		d.Set("cpu_threads_per_core", instance.CpuOptions.ThreadsPerCore)
	}
	return diags
}

func resourceAwsInstanceFindByID(client *apiClient, id string) (types.Instance, error) {
	instanceIDs := make([]string, 1)
	instanceIDs[0] = id
	input := &ec2.DescribeInstancesInput{
		InstanceIds: instanceIDs,
	}
	instances, err := client.client.DescribeInstances(context.TODO(), input)
	if err != nil {
		return types.Instance{}, err
	}
	return instances.Reservations[0].Instances[0], nil
}

func resourceScaffoldingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	log.Printf("[INFO] umut resourceScaffoldingUpdate is working")

	return diag.Errorf("not implemented")
}

func resourceScaffoldingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] umut resourceScaffoldingDelete is working")
	var diags diag.Diagnostics
	client := meta.(*apiClient)
	instanceID := d.Id()
	instanceIDs := make([]string, 1)
	instanceIDs[0] = instanceID
	input := &ec2.TerminateInstancesInput{
		InstanceIds: instanceIDs,
	}
	_, err := client.client.TerminateInstances(context.TODO(), input)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
