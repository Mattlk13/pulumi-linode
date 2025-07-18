// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages allocation of reserved IPv4 address in a region and optionally assigning the reserved address to a Linode instance.
//
// For more information, see the corresponding [API documentation](https://techdocs.akamai.com/linode-api/reference/post-allocate-ip).
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
//			_, err := linode.NewNetworkingIp(ctx, "test_ip", &linode.NetworkingIpArgs{
//				Type:     pulumi.String("ipv4"),
//				LinodeId: pulumi.Int(12345),
//				Public:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// IP addresses can be imported using the IP address ID, e.g.
//
// ```sh
// $ pulumi import linode:index/networkingIp:NetworkingIp example_ip 172.104.30.209
// ```
type NetworkingIp struct {
	pulumi.CustomResourceState

	// The IPv4 address that is configured as a 1:1 NAT for this VPC interface.
	Address pulumi.StringOutput `pulumi:"address"`
	// The default gateway for this address.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// The ID of the Linode to which the IP address will be assigned. Updating this field on an ephemeral IP will trigger a recreation. Conflicts with `region`.
	LinodeId pulumi.IntOutput `pulumi:"linodeId"`
	// The number of bits set in the subnet mask.
	Prefix pulumi.IntOutput `pulumi:"prefix"`
	// Whether the IP address is public. Defaults to true.
	Public pulumi.BoolOutput `pulumi:"public"`
	// The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.
	Rdns pulumi.StringOutput `pulumi:"rdns"`
	// The region for the reserved IPv4 address. Required when reserved is true and linodeId is not set.
	Region pulumi.StringOutput `pulumi:"region"`
	// Whether the IPv4 address should be reserved.
	Reserved pulumi.BoolOutput `pulumi:"reserved"`
	// The mask that separates host bits from network bits for this address.
	SubnetMask pulumi.StringOutput `pulumi:"subnetMask"`
	// The type of IP address. (ipv4, ipv6, etc.)
	Type pulumi.StringOutput `pulumi:"type"`
	// Contains information about the NAT 1:1 mapping of a public IP address to a VPC subnet.
	VpcNat11 NetworkingIpVpcNat11Output `pulumi:"vpcNat11"`
}

// NewNetworkingIp registers a new resource with the given unique name, arguments, and options.
func NewNetworkingIp(ctx *pulumi.Context,
	name string, args *NetworkingIpArgs, opts ...pulumi.ResourceOption) (*NetworkingIp, error) {
	if args == nil {
		args = &NetworkingIpArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkingIp
	err := ctx.RegisterResource("linode:index/networkingIp:NetworkingIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkingIp gets an existing NetworkingIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkingIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkingIpState, opts ...pulumi.ResourceOption) (*NetworkingIp, error) {
	var resource NetworkingIp
	err := ctx.ReadResource("linode:index/networkingIp:NetworkingIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkingIp resources.
type networkingIpState struct {
	// The IPv4 address that is configured as a 1:1 NAT for this VPC interface.
	Address *string `pulumi:"address"`
	// The default gateway for this address.
	Gateway *string `pulumi:"gateway"`
	// The ID of the Linode to which the IP address will be assigned. Updating this field on an ephemeral IP will trigger a recreation. Conflicts with `region`.
	LinodeId *int `pulumi:"linodeId"`
	// The number of bits set in the subnet mask.
	Prefix *int `pulumi:"prefix"`
	// Whether the IP address is public. Defaults to true.
	Public *bool `pulumi:"public"`
	// The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.
	Rdns *string `pulumi:"rdns"`
	// The region for the reserved IPv4 address. Required when reserved is true and linodeId is not set.
	Region *string `pulumi:"region"`
	// Whether the IPv4 address should be reserved.
	Reserved *bool `pulumi:"reserved"`
	// The mask that separates host bits from network bits for this address.
	SubnetMask *string `pulumi:"subnetMask"`
	// The type of IP address. (ipv4, ipv6, etc.)
	Type *string `pulumi:"type"`
	// Contains information about the NAT 1:1 mapping of a public IP address to a VPC subnet.
	VpcNat11 *NetworkingIpVpcNat11 `pulumi:"vpcNat11"`
}

type NetworkingIpState struct {
	// The IPv4 address that is configured as a 1:1 NAT for this VPC interface.
	Address pulumi.StringPtrInput
	// The default gateway for this address.
	Gateway pulumi.StringPtrInput
	// The ID of the Linode to which the IP address will be assigned. Updating this field on an ephemeral IP will trigger a recreation. Conflicts with `region`.
	LinodeId pulumi.IntPtrInput
	// The number of bits set in the subnet mask.
	Prefix pulumi.IntPtrInput
	// Whether the IP address is public. Defaults to true.
	Public pulumi.BoolPtrInput
	// The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.
	Rdns pulumi.StringPtrInput
	// The region for the reserved IPv4 address. Required when reserved is true and linodeId is not set.
	Region pulumi.StringPtrInput
	// Whether the IPv4 address should be reserved.
	Reserved pulumi.BoolPtrInput
	// The mask that separates host bits from network bits for this address.
	SubnetMask pulumi.StringPtrInput
	// The type of IP address. (ipv4, ipv6, etc.)
	Type pulumi.StringPtrInput
	// Contains information about the NAT 1:1 mapping of a public IP address to a VPC subnet.
	VpcNat11 NetworkingIpVpcNat11PtrInput
}

func (NetworkingIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkingIpState)(nil)).Elem()
}

type networkingIpArgs struct {
	// The ID of the Linode to which the IP address will be assigned. Updating this field on an ephemeral IP will trigger a recreation. Conflicts with `region`.
	LinodeId *int `pulumi:"linodeId"`
	// Whether the IP address is public. Defaults to true.
	Public *bool `pulumi:"public"`
	// The region for the reserved IPv4 address. Required when reserved is true and linodeId is not set.
	Region *string `pulumi:"region"`
	// Whether the IPv4 address should be reserved.
	Reserved *bool `pulumi:"reserved"`
	// The type of IP address. (ipv4, ipv6, etc.)
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a NetworkingIp resource.
type NetworkingIpArgs struct {
	// The ID of the Linode to which the IP address will be assigned. Updating this field on an ephemeral IP will trigger a recreation. Conflicts with `region`.
	LinodeId pulumi.IntPtrInput
	// Whether the IP address is public. Defaults to true.
	Public pulumi.BoolPtrInput
	// The region for the reserved IPv4 address. Required when reserved is true and linodeId is not set.
	Region pulumi.StringPtrInput
	// Whether the IPv4 address should be reserved.
	Reserved pulumi.BoolPtrInput
	// The type of IP address. (ipv4, ipv6, etc.)
	Type pulumi.StringPtrInput
}

func (NetworkingIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkingIpArgs)(nil)).Elem()
}

type NetworkingIpInput interface {
	pulumi.Input

	ToNetworkingIpOutput() NetworkingIpOutput
	ToNetworkingIpOutputWithContext(ctx context.Context) NetworkingIpOutput
}

func (*NetworkingIp) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkingIp)(nil)).Elem()
}

func (i *NetworkingIp) ToNetworkingIpOutput() NetworkingIpOutput {
	return i.ToNetworkingIpOutputWithContext(context.Background())
}

func (i *NetworkingIp) ToNetworkingIpOutputWithContext(ctx context.Context) NetworkingIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingIpOutput)
}

// NetworkingIpArrayInput is an input type that accepts NetworkingIpArray and NetworkingIpArrayOutput values.
// You can construct a concrete instance of `NetworkingIpArrayInput` via:
//
//	NetworkingIpArray{ NetworkingIpArgs{...} }
type NetworkingIpArrayInput interface {
	pulumi.Input

	ToNetworkingIpArrayOutput() NetworkingIpArrayOutput
	ToNetworkingIpArrayOutputWithContext(context.Context) NetworkingIpArrayOutput
}

type NetworkingIpArray []NetworkingIpInput

func (NetworkingIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkingIp)(nil)).Elem()
}

func (i NetworkingIpArray) ToNetworkingIpArrayOutput() NetworkingIpArrayOutput {
	return i.ToNetworkingIpArrayOutputWithContext(context.Background())
}

func (i NetworkingIpArray) ToNetworkingIpArrayOutputWithContext(ctx context.Context) NetworkingIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingIpArrayOutput)
}

// NetworkingIpMapInput is an input type that accepts NetworkingIpMap and NetworkingIpMapOutput values.
// You can construct a concrete instance of `NetworkingIpMapInput` via:
//
//	NetworkingIpMap{ "key": NetworkingIpArgs{...} }
type NetworkingIpMapInput interface {
	pulumi.Input

	ToNetworkingIpMapOutput() NetworkingIpMapOutput
	ToNetworkingIpMapOutputWithContext(context.Context) NetworkingIpMapOutput
}

type NetworkingIpMap map[string]NetworkingIpInput

func (NetworkingIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkingIp)(nil)).Elem()
}

func (i NetworkingIpMap) ToNetworkingIpMapOutput() NetworkingIpMapOutput {
	return i.ToNetworkingIpMapOutputWithContext(context.Background())
}

func (i NetworkingIpMap) ToNetworkingIpMapOutputWithContext(ctx context.Context) NetworkingIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingIpMapOutput)
}

type NetworkingIpOutput struct{ *pulumi.OutputState }

func (NetworkingIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkingIp)(nil)).Elem()
}

func (o NetworkingIpOutput) ToNetworkingIpOutput() NetworkingIpOutput {
	return o
}

func (o NetworkingIpOutput) ToNetworkingIpOutputWithContext(ctx context.Context) NetworkingIpOutput {
	return o
}

// The IPv4 address that is configured as a 1:1 NAT for this VPC interface.
func (o NetworkingIpOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The default gateway for this address.
func (o NetworkingIpOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// The ID of the Linode to which the IP address will be assigned. Updating this field on an ephemeral IP will trigger a recreation. Conflicts with `region`.
func (o NetworkingIpOutput) LinodeId() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.IntOutput { return v.LinodeId }).(pulumi.IntOutput)
}

// The number of bits set in the subnet mask.
func (o NetworkingIpOutput) Prefix() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.IntOutput { return v.Prefix }).(pulumi.IntOutput)
}

// Whether the IP address is public. Defaults to true.
func (o NetworkingIpOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.BoolOutput { return v.Public }).(pulumi.BoolOutput)
}

// The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.
func (o NetworkingIpOutput) Rdns() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.StringOutput { return v.Rdns }).(pulumi.StringOutput)
}

// The region for the reserved IPv4 address. Required when reserved is true and linodeId is not set.
func (o NetworkingIpOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Whether the IPv4 address should be reserved.
func (o NetworkingIpOutput) Reserved() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.BoolOutput { return v.Reserved }).(pulumi.BoolOutput)
}

// The mask that separates host bits from network bits for this address.
func (o NetworkingIpOutput) SubnetMask() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.StringOutput { return v.SubnetMask }).(pulumi.StringOutput)
}

// The type of IP address. (ipv4, ipv6, etc.)
func (o NetworkingIpOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkingIp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Contains information about the NAT 1:1 mapping of a public IP address to a VPC subnet.
func (o NetworkingIpOutput) VpcNat11() NetworkingIpVpcNat11Output {
	return o.ApplyT(func(v *NetworkingIp) NetworkingIpVpcNat11Output { return v.VpcNat11 }).(NetworkingIpVpcNat11Output)
}

type NetworkingIpArrayOutput struct{ *pulumi.OutputState }

func (NetworkingIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkingIp)(nil)).Elem()
}

func (o NetworkingIpArrayOutput) ToNetworkingIpArrayOutput() NetworkingIpArrayOutput {
	return o
}

func (o NetworkingIpArrayOutput) ToNetworkingIpArrayOutputWithContext(ctx context.Context) NetworkingIpArrayOutput {
	return o
}

func (o NetworkingIpArrayOutput) Index(i pulumi.IntInput) NetworkingIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkingIp {
		return vs[0].([]*NetworkingIp)[vs[1].(int)]
	}).(NetworkingIpOutput)
}

type NetworkingIpMapOutput struct{ *pulumi.OutputState }

func (NetworkingIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkingIp)(nil)).Elem()
}

func (o NetworkingIpMapOutput) ToNetworkingIpMapOutput() NetworkingIpMapOutput {
	return o
}

func (o NetworkingIpMapOutput) ToNetworkingIpMapOutputWithContext(ctx context.Context) NetworkingIpMapOutput {
	return o
}

func (o NetworkingIpMapOutput) MapIndex(k pulumi.StringInput) NetworkingIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkingIp {
		return vs[0].(map[string]*NetworkingIp)[vs[1].(string)]
	}).(NetworkingIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingIpInput)(nil)).Elem(), &NetworkingIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingIpArrayInput)(nil)).Elem(), NetworkingIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingIpMapInput)(nil)).Elem(), NetworkingIpMap{})
	pulumi.RegisterOutputType(NetworkingIpOutput{})
	pulumi.RegisterOutputType(NetworkingIpArrayOutput{})
	pulumi.RegisterOutputType(NetworkingIpMapOutput{})
}
