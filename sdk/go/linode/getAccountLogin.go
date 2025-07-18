// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a specific Linode account login.
// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-account-login).
//
// ## Example Usage
//
// The following example shows how one might use this data source to access information about a Linode account login.
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
//			_, err := linode.GetAccountLogin(ctx, &linode.GetAccountLoginArgs{
//				Id: 123456,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAccountLogin(ctx *pulumi.Context, args *GetAccountLoginArgs, opts ...pulumi.InvokeOption) (*GetAccountLoginResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccountLoginResult
	err := ctx.Invoke("linode:index/getAccountLogin:getAccountLogin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountLogin.
type GetAccountLoginArgs struct {
	// The unique ID of this login object.
	Id int `pulumi:"id"`
}

// A collection of values returned by getAccountLogin.
type GetAccountLoginResult struct {
	// When the login was initiated.
	Datetime string `pulumi:"datetime"`
	// The unique ID of this login object.
	Id int `pulumi:"id"`
	// The remote IP address that requested the login.
	Ip string `pulumi:"ip"`
	// True if the User that was logged into was a restricted User, false otherwise.
	Restricted bool   `pulumi:"restricted"`
	Status     string `pulumi:"status"`
	// The username of the User that was logged into.
	Username string `pulumi:"username"`
}

func GetAccountLoginOutput(ctx *pulumi.Context, args GetAccountLoginOutputArgs, opts ...pulumi.InvokeOption) GetAccountLoginResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetAccountLoginResultOutput, error) {
			args := v.(GetAccountLoginArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("linode:index/getAccountLogin:getAccountLogin", args, GetAccountLoginResultOutput{}, options).(GetAccountLoginResultOutput), nil
		}).(GetAccountLoginResultOutput)
}

// A collection of arguments for invoking getAccountLogin.
type GetAccountLoginOutputArgs struct {
	// The unique ID of this login object.
	Id pulumi.IntInput `pulumi:"id"`
}

func (GetAccountLoginOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountLoginArgs)(nil)).Elem()
}

// A collection of values returned by getAccountLogin.
type GetAccountLoginResultOutput struct{ *pulumi.OutputState }

func (GetAccountLoginResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountLoginResult)(nil)).Elem()
}

func (o GetAccountLoginResultOutput) ToGetAccountLoginResultOutput() GetAccountLoginResultOutput {
	return o
}

func (o GetAccountLoginResultOutput) ToGetAccountLoginResultOutputWithContext(ctx context.Context) GetAccountLoginResultOutput {
	return o
}

// When the login was initiated.
func (o GetAccountLoginResultOutput) Datetime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountLoginResult) string { return v.Datetime }).(pulumi.StringOutput)
}

// The unique ID of this login object.
func (o GetAccountLoginResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetAccountLoginResult) int { return v.Id }).(pulumi.IntOutput)
}

// The remote IP address that requested the login.
func (o GetAccountLoginResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountLoginResult) string { return v.Ip }).(pulumi.StringOutput)
}

// True if the User that was logged into was a restricted User, false otherwise.
func (o GetAccountLoginResultOutput) Restricted() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAccountLoginResult) bool { return v.Restricted }).(pulumi.BoolOutput)
}

func (o GetAccountLoginResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountLoginResult) string { return v.Status }).(pulumi.StringOutput)
}

// The username of the User that was logged into.
func (o GetAccountLoginResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountLoginResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountLoginResultOutput{})
}
