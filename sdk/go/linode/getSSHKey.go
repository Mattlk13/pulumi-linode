// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `linode_sshkey` provides access to a specifically labeled SSH Key in the Profile of the User identified by the access token.
func LookupSSHKey(ctx *pulumi.Context, args *GetSSHKeyArgs) (*GetSSHKeyResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["label"] = args.Label
	}
	outputs, err := ctx.Invoke("linode:index/getSSHKey:getSSHKey", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSSHKeyResult{
		Created: outputs["created"],
		SshKey: outputs["sshKey"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSSHKey.
type GetSSHKeyArgs struct {
	Label interface{}
}

// A collection of values returned by getSSHKey.
type GetSSHKeyResult struct {
	Created interface{}
	SshKey interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}