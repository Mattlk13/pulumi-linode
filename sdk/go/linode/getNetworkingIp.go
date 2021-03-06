// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides information about a Linode Networking IP Address
//
// ## Example Usage
//
// The following example shows how one might use this data source to access information about a Linode Networking IP Address.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-linode/sdk/v2/go/linode"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := linode.GetNetworkingIp(ctx, &linode.GetNetworkingIpArgs{
// 			Address: "162.159.27.72",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Attributes
//
// The Linode Network IP Address resource exports the following attributes:
//
// * `address` - The IP address.
//
// * `gateway` - The default gateway for this address.
//
// * `subnetMask` - The mask that separates host bits from network bits for this address.
//
// * `prefix` - The number of bits set in the subnet mask.
//
// * `type` - The type of address this is (ipv4, ipv6, ipv6/pool, ipv6/range).
//
// * `public` - Whether this is a public or private IP address.
//
// * `rdns` - The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.
//
// * `linodeId` - The ID of the Linode this address currently belongs to.
//
// * `region` - The Region this IP address resides in.
func GetNetworkingIp(ctx *pulumi.Context, args *GetNetworkingIpArgs, opts ...pulumi.InvokeOption) (*GetNetworkingIpResult, error) {
	var rv GetNetworkingIpResult
	err := ctx.Invoke("linode:index/getNetworkingIp:getNetworkingIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkingIp.
type GetNetworkingIpArgs struct {
	// The IP Address to access.  The address must be associated with the account and a resource that the user has access to view.
	Address string `pulumi:"address"`
}

// A collection of values returned by getNetworkingIp.
type GetNetworkingIpResult struct {
	Address string `pulumi:"address"`
	Gateway string `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	LinodeId   int    `pulumi:"linodeId"`
	Prefix     int    `pulumi:"prefix"`
	Public     bool   `pulumi:"public"`
	Rdns       string `pulumi:"rdns"`
	Region     string `pulumi:"region"`
	SubnetMask string `pulumi:"subnetMask"`
	Type       string `pulumi:"type"`
}
