// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about Linode Managed Database engines that match a set of filters.
// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-databases-engines).
//
// ## Example Usage
//
// Get information about all Linode Managed Database engines:
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
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// all, err := linode.GetDatabaseEngines(ctx, &linode.GetDatabaseEnginesArgs{
// }, nil);
// if err != nil {
// return err
// }
// ctx.Export("engineIds", pulumi.StringArray(%!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ example.pp:3,11-28)))
// return nil
// })
// }
// ```
//
// Get information about all Linode MySQL Database engines:
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
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// mysql, err := linode.GetDatabaseEngines(ctx, &linode.GetDatabaseEnginesArgs{
// Filters: []linode.GetDatabaseEnginesFilter{
// {
// Name: "engine",
// Values: []string{
// "mysql",
// },
// },
// },
// }, nil);
// if err != nil {
// return err
// }
// ctx.Export("engineIds", pulumi.StringArray(%!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ example.pp:8,11-30)))
// return nil
// })
// }
// ```
//
// Create a Linode MySQL Database using the latest support MySQL version:
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
//			mysql, err := linode.GetDatabaseEngines(ctx, &linode.GetDatabaseEnginesArgs{
//				Latest: pulumi.BoolRef(true),
//				Filters: []linode.GetDatabaseEnginesFilter{
//					{
//						Name: "engine",
//						Values: []string{
//							"mysql",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = linode.NewDatabaseMysql(ctx, "my_db", &linode.DatabaseMysqlArgs{
//				Label:    pulumi.String("mydb"),
//				EngineId: pulumi.String(mysql.Engines[0].Id),
//				Region:   pulumi.String("us-southeast"),
//				Type:     pulumi.String("g6-nanode-1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDatabaseEngines(ctx *pulumi.Context, args *GetDatabaseEnginesArgs, opts ...pulumi.InvokeOption) (*GetDatabaseEnginesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabaseEnginesResult
	err := ctx.Invoke("linode:index/getDatabaseEngines:getDatabaseEngines", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseEngines.
type GetDatabaseEnginesArgs struct {
	Engines []GetDatabaseEnginesEngine `pulumi:"engines"`
	Filters []GetDatabaseEnginesFilter `pulumi:"filters"`
	// If true, only the latest engine version will be returned.
	//
	// * `filter` - (Optional) A set of filters used to select engines that meet certain requirements.
	Latest *bool `pulumi:"latest"`
	// The order in which results should be returned. (`asc`, `desc`; default `asc`)
	Order *string `pulumi:"order"`
	// The attribute to order the results by. (`version`)
	OrderBy *string `pulumi:"orderBy"`
}

// A collection of values returned by getDatabaseEngines.
type GetDatabaseEnginesResult struct {
	Engines []GetDatabaseEnginesEngine `pulumi:"engines"`
	Filters []GetDatabaseEnginesFilter `pulumi:"filters"`
	// The Managed Database engine ID in engine/version format.
	Id      string  `pulumi:"id"`
	Latest  *bool   `pulumi:"latest"`
	Order   *string `pulumi:"order"`
	OrderBy *string `pulumi:"orderBy"`
}

func GetDatabaseEnginesOutput(ctx *pulumi.Context, args GetDatabaseEnginesOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseEnginesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDatabaseEnginesResultOutput, error) {
			args := v.(GetDatabaseEnginesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("linode:index/getDatabaseEngines:getDatabaseEngines", args, GetDatabaseEnginesResultOutput{}, options).(GetDatabaseEnginesResultOutput), nil
		}).(GetDatabaseEnginesResultOutput)
}

// A collection of arguments for invoking getDatabaseEngines.
type GetDatabaseEnginesOutputArgs struct {
	Engines GetDatabaseEnginesEngineArrayInput `pulumi:"engines"`
	Filters GetDatabaseEnginesFilterArrayInput `pulumi:"filters"`
	// If true, only the latest engine version will be returned.
	//
	// * `filter` - (Optional) A set of filters used to select engines that meet certain requirements.
	Latest pulumi.BoolPtrInput `pulumi:"latest"`
	// The order in which results should be returned. (`asc`, `desc`; default `asc`)
	Order pulumi.StringPtrInput `pulumi:"order"`
	// The attribute to order the results by. (`version`)
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
}

func (GetDatabaseEnginesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseEnginesArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseEngines.
type GetDatabaseEnginesResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseEnginesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseEnginesResult)(nil)).Elem()
}

func (o GetDatabaseEnginesResultOutput) ToGetDatabaseEnginesResultOutput() GetDatabaseEnginesResultOutput {
	return o
}

func (o GetDatabaseEnginesResultOutput) ToGetDatabaseEnginesResultOutputWithContext(ctx context.Context) GetDatabaseEnginesResultOutput {
	return o
}

func (o GetDatabaseEnginesResultOutput) Engines() GetDatabaseEnginesEngineArrayOutput {
	return o.ApplyT(func(v GetDatabaseEnginesResult) []GetDatabaseEnginesEngine { return v.Engines }).(GetDatabaseEnginesEngineArrayOutput)
}

func (o GetDatabaseEnginesResultOutput) Filters() GetDatabaseEnginesFilterArrayOutput {
	return o.ApplyT(func(v GetDatabaseEnginesResult) []GetDatabaseEnginesFilter { return v.Filters }).(GetDatabaseEnginesFilterArrayOutput)
}

// The Managed Database engine ID in engine/version format.
func (o GetDatabaseEnginesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseEnginesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDatabaseEnginesResultOutput) Latest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDatabaseEnginesResult) *bool { return v.Latest }).(pulumi.BoolPtrOutput)
}

func (o GetDatabaseEnginesResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseEnginesResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseEnginesResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseEnginesResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseEnginesResultOutput{})
}
