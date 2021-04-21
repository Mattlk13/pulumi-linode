// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a Linode NodeBalancer Config.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-linode/sdk/v3/go/linode"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := linode.LookupNodeBalancerConfig(ctx, &linode.LookupNodeBalancerConfigArgs{
// 			Id:             123,
// 			NodebalancerId: 456,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupNodeBalancerConfig(ctx *pulumi.Context, args *LookupNodeBalancerConfigArgs, opts ...pulumi.InvokeOption) (*LookupNodeBalancerConfigResult, error) {
	var rv LookupNodeBalancerConfigResult
	err := ctx.Invoke("linode:index/getNodeBalancerConfig:getNodeBalancerConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNodeBalancerConfig.
type LookupNodeBalancerConfigArgs struct {
	// The config's ID.
	Id int `pulumi:"id"`
	// The ID of the NodeBalancer that contains the config.
	NodebalancerId int `pulumi:"nodebalancerId"`
}

// A collection of values returned by getNodeBalancerConfig.
type LookupNodeBalancerConfigResult struct {
	// What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source
	Algorithm string `pulumi:"algorithm"`
	// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and httpBody rely on the backend serving HTTP, and that the response returned matches what is expected.
	Check string `pulumi:"check"`
	// How many times to attempt a check before considering a backend to be down. (1-30)
	CheckAttempts int    `pulumi:"checkAttempts"`
	CheckBody     string `pulumi:"checkBody"`
	// How often, in seconds, to check that backends are up and serving requests.
	CheckInterval int `pulumi:"checkInterval"`
	// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
	CheckPassive bool `pulumi:"checkPassive"`
	// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
	CheckPath string `pulumi:"checkPath"`
	// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
	CheckTimeout int `pulumi:"checkTimeout"`
	// What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.
	CipherSuite    string                            `pulumi:"cipherSuite"`
	Id             int                               `pulumi:"id"`
	NodeStatuses   []GetNodeBalancerConfigNodeStatus `pulumi:"nodeStatuses"`
	NodebalancerId int                               `pulumi:"nodebalancerId"`
	// The TCP port this Config is for.
	Port int `pulumi:"port"`
	// The protocol this port is configured to serve. If this is set to https you must include an sslCert and an ssl_key. (Defaults to "http")
	Protocol string `pulumi:"protocol"`
	// The version of ProxyProtocol to use for the underlying NodeBalancer. This requires protocol to be `tcp`. Valid values are `none`, `v1`, and `v2`. (Defaults to `none`)
	ProxyProtocol string `pulumi:"proxyProtocol"`
	// The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
	SslCommonname string `pulumi:"sslCommonname"`
	// The read-only fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
	SslFingerprint string `pulumi:"sslFingerprint"`
	// Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'
	Stickiness string `pulumi:"stickiness"`
}