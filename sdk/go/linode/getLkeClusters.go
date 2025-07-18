// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about a list of current Linode Kubernetes (LKE) clusters on your account that match a set of filters.
// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-lke-clusters).
//
// ## Example Usage
//
// Get information about all LKE clusters with a specific tag:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			specific, err := linode.GetLkeClusters(ctx, &linode.GetLkeClustersArgs{
//				Filters: []linode.GetLkeClustersFilter{
//					{
//						Name: "tags",
//						Values: []string{
//							"test-tag",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("lkeCluster", specific.LkeClusters[0].Id)
//			return nil
//		})
//	}
//
// ```
//
// ## Filterable Fields
//
// * `k8sVersion`
//
// * `label`
//
// * `region`
//
// * `tags`
//
// * `status`
//
// * `created`
//
// * `updated`
func GetLkeClusters(ctx *pulumi.Context, args *GetLkeClustersArgs, opts ...pulumi.InvokeOption) (*GetLkeClustersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLkeClustersResult
	err := ctx.Invoke("linode:index/getLkeClusters:getLkeClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLkeClusters.
type GetLkeClustersArgs struct {
	Filters     []GetLkeClustersFilter     `pulumi:"filters"`
	LkeClusters []GetLkeClustersLkeCluster `pulumi:"lkeClusters"`
	// The order in which results should be returned. (`asc`, `desc`; default `asc`)
	Order *string `pulumi:"order"`
	// The attribute to order the results by. See the Filterable Fields section for a list of valid fields.
	OrderBy *string `pulumi:"orderBy"`
}

// A collection of values returned by getLkeClusters.
type GetLkeClustersResult struct {
	Filters []GetLkeClustersFilter `pulumi:"filters"`
	// The LKE Cluster's ID.
	Id          string                     `pulumi:"id"`
	LkeClusters []GetLkeClustersLkeCluster `pulumi:"lkeClusters"`
	Order       *string                    `pulumi:"order"`
	OrderBy     *string                    `pulumi:"orderBy"`
}

func GetLkeClustersOutput(ctx *pulumi.Context, args GetLkeClustersOutputArgs, opts ...pulumi.InvokeOption) GetLkeClustersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLkeClustersResultOutput, error) {
			args := v.(GetLkeClustersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("linode:index/getLkeClusters:getLkeClusters", args, GetLkeClustersResultOutput{}, options).(GetLkeClustersResultOutput), nil
		}).(GetLkeClustersResultOutput)
}

// A collection of arguments for invoking getLkeClusters.
type GetLkeClustersOutputArgs struct {
	Filters     GetLkeClustersFilterArrayInput     `pulumi:"filters"`
	LkeClusters GetLkeClustersLkeClusterArrayInput `pulumi:"lkeClusters"`
	// The order in which results should be returned. (`asc`, `desc`; default `asc`)
	Order pulumi.StringPtrInput `pulumi:"order"`
	// The attribute to order the results by. See the Filterable Fields section for a list of valid fields.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
}

func (GetLkeClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLkeClustersArgs)(nil)).Elem()
}

// A collection of values returned by getLkeClusters.
type GetLkeClustersResultOutput struct{ *pulumi.OutputState }

func (GetLkeClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLkeClustersResult)(nil)).Elem()
}

func (o GetLkeClustersResultOutput) ToGetLkeClustersResultOutput() GetLkeClustersResultOutput {
	return o
}

func (o GetLkeClustersResultOutput) ToGetLkeClustersResultOutputWithContext(ctx context.Context) GetLkeClustersResultOutput {
	return o
}

func (o GetLkeClustersResultOutput) Filters() GetLkeClustersFilterArrayOutput {
	return o.ApplyT(func(v GetLkeClustersResult) []GetLkeClustersFilter { return v.Filters }).(GetLkeClustersFilterArrayOutput)
}

// The LKE Cluster's ID.
func (o GetLkeClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLkeClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLkeClustersResultOutput) LkeClusters() GetLkeClustersLkeClusterArrayOutput {
	return o.ApplyT(func(v GetLkeClustersResult) []GetLkeClustersLkeCluster { return v.LkeClusters }).(GetLkeClustersLkeClusterArrayOutput)
}

func (o GetLkeClustersResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLkeClustersResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o GetLkeClustersResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLkeClustersResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLkeClustersResultOutput{})
}
