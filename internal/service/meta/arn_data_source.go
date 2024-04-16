// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by tools/tfsdk2fw/main.go. Manual editing is required.

package meta

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
)

// @FrameworkDataSource
func newDataSourceARN(context.Context) (datasource.DataSourceWithConfigure, error) {
	d := &dataSourceARN{}

	return d, nil
}

type dataSourceARN struct {
	framework.DataSourceWithConfigure
}

// Metadata should return the full name of the data source, such as
// examplecloud_thing.
func (d *dataSourceARN) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) { // nosemgrep:ci.meta-in-func-name
	response.TypeName = "aws_arn"
}

// Schema returns the schema for this data source.
func (d *dataSourceARN) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account": schema.StringAttribute{
				Computed: true,
			},
			"arn": schema.StringAttribute{
				CustomType: fwtypes.ARNType,
				Required:   true,
			},
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"partition": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
			},
			"resource": schema.StringAttribute{
				Computed: true,
			},
			"service": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Read is called when the provider must read data source values in order to update state.
// Config values should be read from the ReadRequest and new state values set on the ReadResponse.
func (d *dataSourceARN) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data dataSourceARNData

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	arn := data.ARN.ValueARN()

	data.Account = types.StringValue(arn.AccountID)
	data.ID = types.StringValue(arn.String())
	data.Partition = types.StringValue(arn.Partition)
	data.Region = types.StringValue(arn.Region)
	data.Resource = types.StringValue(arn.Resource)
	data.Service = types.StringValue(arn.Service)

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

type dataSourceARNData struct {
	Account   types.String `tfsdk:"account"`
	ARN       fwtypes.ARN  `tfsdk:"arn"`
	ID        types.String `tfsdk:"id"`
	Partition types.String `tfsdk:"partition"`
	Region    types.String `tfsdk:"region"`
	Resource  types.String `tfsdk:"resource"`
	Service   types.String `tfsdk:"service"`
}
