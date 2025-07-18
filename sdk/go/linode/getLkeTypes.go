// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about Linode LKE types that match a set of filters.
// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-lke-types).
func GetLkeTypes(ctx *pulumi.Context, args *GetLkeTypesArgs, opts ...pulumi.InvokeOption) (*GetLkeTypesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLkeTypesResult
	err := ctx.Invoke("linode:index/getLkeTypes:getLkeTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLkeTypes.
type GetLkeTypesArgs struct {
	Filters []GetLkeTypesFilter `pulumi:"filters"`
	// The order in which results should be returned. (`asc`, `desc`; default `asc`)
	Order *string `pulumi:"order"`
	// The attribute to order the results by. See the Filterable Fields section for a list of valid fields.
	OrderBy *string           `pulumi:"orderBy"`
	Types   []GetLkeTypesType `pulumi:"types"`
}

// A collection of values returned by getLkeTypes.
type GetLkeTypesResult struct {
	Filters []GetLkeTypesFilter `pulumi:"filters"`
	// The ID representing the Kubernetes type.
	Id      string            `pulumi:"id"`
	Order   *string           `pulumi:"order"`
	OrderBy *string           `pulumi:"orderBy"`
	Types   []GetLkeTypesType `pulumi:"types"`
}

func GetLkeTypesOutput(ctx *pulumi.Context, args GetLkeTypesOutputArgs, opts ...pulumi.InvokeOption) GetLkeTypesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLkeTypesResultOutput, error) {
			args := v.(GetLkeTypesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("linode:index/getLkeTypes:getLkeTypes", args, GetLkeTypesResultOutput{}, options).(GetLkeTypesResultOutput), nil
		}).(GetLkeTypesResultOutput)
}

// A collection of arguments for invoking getLkeTypes.
type GetLkeTypesOutputArgs struct {
	Filters GetLkeTypesFilterArrayInput `pulumi:"filters"`
	// The order in which results should be returned. (`asc`, `desc`; default `asc`)
	Order pulumi.StringPtrInput `pulumi:"order"`
	// The attribute to order the results by. See the Filterable Fields section for a list of valid fields.
	OrderBy pulumi.StringPtrInput     `pulumi:"orderBy"`
	Types   GetLkeTypesTypeArrayInput `pulumi:"types"`
}

func (GetLkeTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLkeTypesArgs)(nil)).Elem()
}

// A collection of values returned by getLkeTypes.
type GetLkeTypesResultOutput struct{ *pulumi.OutputState }

func (GetLkeTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLkeTypesResult)(nil)).Elem()
}

func (o GetLkeTypesResultOutput) ToGetLkeTypesResultOutput() GetLkeTypesResultOutput {
	return o
}

func (o GetLkeTypesResultOutput) ToGetLkeTypesResultOutputWithContext(ctx context.Context) GetLkeTypesResultOutput {
	return o
}

func (o GetLkeTypesResultOutput) Filters() GetLkeTypesFilterArrayOutput {
	return o.ApplyT(func(v GetLkeTypesResult) []GetLkeTypesFilter { return v.Filters }).(GetLkeTypesFilterArrayOutput)
}

// The ID representing the Kubernetes type.
func (o GetLkeTypesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLkeTypesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLkeTypesResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLkeTypesResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o GetLkeTypesResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLkeTypesResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetLkeTypesResultOutput) Types() GetLkeTypesTypeArrayOutput {
	return o.ApplyT(func(v GetLkeTypesResult) []GetLkeTypesType { return v.Types }).(GetLkeTypesTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLkeTypesResultOutput{})
}
