// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a Linode NodeBalancer node.
// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-node-balancer-node).
//
// ## Example Usage
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
//			_, err := linode.LookupNodeBalancerNode(ctx, &linode.LookupNodeBalancerNodeArgs{
//				Id:             123,
//				NodebalancerId: 456,
//				ConfigId:       789,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNodeBalancerNode(ctx *pulumi.Context, args *LookupNodeBalancerNodeArgs, opts ...pulumi.InvokeOption) (*LookupNodeBalancerNodeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNodeBalancerNodeResult
	err := ctx.Invoke("linode:index/getNodeBalancerNode:getNodeBalancerNode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNodeBalancerNode.
type LookupNodeBalancerNodeArgs struct {
	// The ID of the config that contains the Node.
	ConfigId int `pulumi:"configId"`
	// The node's ID.
	Id int `pulumi:"id"`
	// The ID of the NodeBalancer that contains the node.
	NodebalancerId int `pulumi:"nodebalancerId"`
}

// A collection of values returned by getNodeBalancerNode.
type LookupNodeBalancerNodeResult struct {
	// The private IP Address where this backend can be reached.
	Address  string `pulumi:"address"`
	ConfigId int    `pulumi:"configId"`
	Id       int    `pulumi:"id"`
	// The label of the Linode NodeBalancer Node. This is for display purposes only.
	Label string `pulumi:"label"`
	// The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. (`accept`, `reject`, `drain`, `backup`)
	Mode           string `pulumi:"mode"`
	NodebalancerId int    `pulumi:"nodebalancerId"`
	// The current status of this node, based on the configured checks of its NodeBalancer Config. (`unknown`, `UP`, `DOWN`).
	Status string `pulumi:"status"`
	// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).
	Weight int `pulumi:"weight"`
}

func LookupNodeBalancerNodeOutput(ctx *pulumi.Context, args LookupNodeBalancerNodeOutputArgs, opts ...pulumi.InvokeOption) LookupNodeBalancerNodeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNodeBalancerNodeResultOutput, error) {
			args := v.(LookupNodeBalancerNodeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("linode:index/getNodeBalancerNode:getNodeBalancerNode", args, LookupNodeBalancerNodeResultOutput{}, options).(LookupNodeBalancerNodeResultOutput), nil
		}).(LookupNodeBalancerNodeResultOutput)
}

// A collection of arguments for invoking getNodeBalancerNode.
type LookupNodeBalancerNodeOutputArgs struct {
	// The ID of the config that contains the Node.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// The node's ID.
	Id pulumi.IntInput `pulumi:"id"`
	// The ID of the NodeBalancer that contains the node.
	NodebalancerId pulumi.IntInput `pulumi:"nodebalancerId"`
}

func (LookupNodeBalancerNodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeBalancerNodeArgs)(nil)).Elem()
}

// A collection of values returned by getNodeBalancerNode.
type LookupNodeBalancerNodeResultOutput struct{ *pulumi.OutputState }

func (LookupNodeBalancerNodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeBalancerNodeResult)(nil)).Elem()
}

func (o LookupNodeBalancerNodeResultOutput) ToLookupNodeBalancerNodeResultOutput() LookupNodeBalancerNodeResultOutput {
	return o
}

func (o LookupNodeBalancerNodeResultOutput) ToLookupNodeBalancerNodeResultOutputWithContext(ctx context.Context) LookupNodeBalancerNodeResultOutput {
	return o
}

// The private IP Address where this backend can be reached.
func (o LookupNodeBalancerNodeResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeBalancerNodeResult) string { return v.Address }).(pulumi.StringOutput)
}

func (o LookupNodeBalancerNodeResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeBalancerNodeResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

func (o LookupNodeBalancerNodeResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeBalancerNodeResult) int { return v.Id }).(pulumi.IntOutput)
}

// The label of the Linode NodeBalancer Node. This is for display purposes only.
func (o LookupNodeBalancerNodeResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeBalancerNodeResult) string { return v.Label }).(pulumi.StringOutput)
}

// The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. (`accept`, `reject`, `drain`, `backup`)
func (o LookupNodeBalancerNodeResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeBalancerNodeResult) string { return v.Mode }).(pulumi.StringOutput)
}

func (o LookupNodeBalancerNodeResultOutput) NodebalancerId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeBalancerNodeResult) int { return v.NodebalancerId }).(pulumi.IntOutput)
}

// The current status of this node, based on the configured checks of its NodeBalancer Config. (`unknown`, `UP`, `DOWN`).
func (o LookupNodeBalancerNodeResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeBalancerNodeResult) string { return v.Status }).(pulumi.StringOutput)
}

// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).
func (o LookupNodeBalancerNodeResultOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeBalancerNodeResult) int { return v.Weight }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNodeBalancerNodeResultOutput{})
}
