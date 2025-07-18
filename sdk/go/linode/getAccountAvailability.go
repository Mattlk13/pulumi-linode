// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about service availability in a region to an account specifically.
// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-account-availability).
//
// ## Example Usage
//
// The following example shows how one might use this data source to access information about a Linode account availability.
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
//			_, err := linode.GetAccountAvailability(ctx, &linode.GetAccountAvailabilityArgs{
//				Region: "us-east",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAccountAvailability(ctx *pulumi.Context, args *GetAccountAvailabilityArgs, opts ...pulumi.InvokeOption) (*GetAccountAvailabilityResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccountAvailabilityResult
	err := ctx.Invoke("linode:index/getAccountAvailability:getAccountAvailability", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountAvailability.
type GetAccountAvailabilityArgs struct {
	// The region ID.
	Region string `pulumi:"region"`
}

// A collection of values returned by getAccountAvailability.
type GetAccountAvailabilityResult struct {
	// A set of services which are available to the account in a region.
	Availables []string `pulumi:"availables"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The region ID.
	Region string `pulumi:"region"`
	// A set of services which are unavailable to the account in a region.
	Unavailables []string `pulumi:"unavailables"`
}

func GetAccountAvailabilityOutput(ctx *pulumi.Context, args GetAccountAvailabilityOutputArgs, opts ...pulumi.InvokeOption) GetAccountAvailabilityResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetAccountAvailabilityResultOutput, error) {
			args := v.(GetAccountAvailabilityArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("linode:index/getAccountAvailability:getAccountAvailability", args, GetAccountAvailabilityResultOutput{}, options).(GetAccountAvailabilityResultOutput), nil
		}).(GetAccountAvailabilityResultOutput)
}

// A collection of arguments for invoking getAccountAvailability.
type GetAccountAvailabilityOutputArgs struct {
	// The region ID.
	Region pulumi.StringInput `pulumi:"region"`
}

func (GetAccountAvailabilityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountAvailabilityArgs)(nil)).Elem()
}

// A collection of values returned by getAccountAvailability.
type GetAccountAvailabilityResultOutput struct{ *pulumi.OutputState }

func (GetAccountAvailabilityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountAvailabilityResult)(nil)).Elem()
}

func (o GetAccountAvailabilityResultOutput) ToGetAccountAvailabilityResultOutput() GetAccountAvailabilityResultOutput {
	return o
}

func (o GetAccountAvailabilityResultOutput) ToGetAccountAvailabilityResultOutputWithContext(ctx context.Context) GetAccountAvailabilityResultOutput {
	return o
}

// A set of services which are available to the account in a region.
func (o GetAccountAvailabilityResultOutput) Availables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountAvailabilityResult) []string { return v.Availables }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountAvailabilityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountAvailabilityResult) string { return v.Id }).(pulumi.StringOutput)
}

// The region ID.
func (o GetAccountAvailabilityResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountAvailabilityResult) string { return v.Region }).(pulumi.StringOutput)
}

// A set of services which are unavailable to the account in a region.
func (o GetAccountAvailabilityResultOutput) Unavailables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountAvailabilityResult) []string { return v.Unavailables }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountAvailabilityResultOutput{})
}
