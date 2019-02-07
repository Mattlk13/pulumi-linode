// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a Linode user
// 
// ## Attributes
// 
// The Linode User resource exports the following attributes:
// 
// * `ssh_keys` - A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the authorized_users field of a create Linode, rebuild Linode, or create Disk request.
// 
// * `email` - The email address for this User, for account management communications, and may be used for other communications as configured.
// 
// * `restricted` - If true, this User must be granted access to perform actions or access entities on this Account.
func LookupUser(ctx *pulumi.Context, args *GetUserArgs) (*GetUserResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["username"] = args.Username
	}
	outputs, err := ctx.Invoke("linode:index/getUser:getUser", inputs)
	if err != nil {
		return nil, err
	}
	return &GetUserResult{
		Email: outputs["email"],
		Restricted: outputs["restricted"],
		SshKeys: outputs["sshKeys"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getUser.
type GetUserArgs struct {
	// The unique username of this User.
	Username interface{}
}

// A collection of values returned by getUser.
type GetUserResult struct {
	Email interface{}
	Restricted interface{}
	SshKeys interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}