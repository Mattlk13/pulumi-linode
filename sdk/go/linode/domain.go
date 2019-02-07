// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Linode Domain resource.  This can be used to create, modify, and delete Linode Domains through Linode's managed DNS service.
// For more information, see [DNS Manager](https://www.linode.com/docs/platform/manager/dns-manager/) and the [Linode APIv4 docs](https://developers.linode.com/api/v4#operation/createDomain).
// 
// The Linode Guide, [Deploy a WordPress Site Using Terraform and Linode StackScripts](https://www.linode.com/docs/applications/configuration-management/deploy-a-wordpress-site-using-terraform-and-linode-stackscripts/), demonstrates the management of Linode Domain resources in the context of Linode Instance running WordPress.
// 
// ## Attributes
// 
// This resource exports no additional attributes, however `status` may reflect degraded states.
type Domain struct {
	s *pulumi.ResourceState
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOpt) (*Domain, error) {
	if args == nil || args.Domain == nil {
		return nil, errors.New("missing required argument 'Domain'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["axfrIps"] = nil
		inputs["description"] = nil
		inputs["domain"] = nil
		inputs["expireSec"] = nil
		inputs["group"] = nil
		inputs["masterIps"] = nil
		inputs["refreshSec"] = nil
		inputs["retrySec"] = nil
		inputs["soaEmail"] = nil
		inputs["status"] = nil
		inputs["tags"] = nil
		inputs["ttlSec"] = nil
		inputs["type"] = nil
	} else {
		inputs["axfrIps"] = args.AxfrIps
		inputs["description"] = args.Description
		inputs["domain"] = args.Domain
		inputs["expireSec"] = args.ExpireSec
		inputs["group"] = args.Group
		inputs["masterIps"] = args.MasterIps
		inputs["refreshSec"] = args.RefreshSec
		inputs["retrySec"] = args.RetrySec
		inputs["soaEmail"] = args.SoaEmail
		inputs["status"] = args.Status
		inputs["tags"] = args.Tags
		inputs["ttlSec"] = args.TtlSec
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("linode:index/domain:Domain", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DomainState, opts ...pulumi.ResourceOpt) (*Domain, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["axfrIps"] = state.AxfrIps
		inputs["description"] = state.Description
		inputs["domain"] = state.Domain
		inputs["expireSec"] = state.ExpireSec
		inputs["group"] = state.Group
		inputs["masterIps"] = state.MasterIps
		inputs["refreshSec"] = state.RefreshSec
		inputs["retrySec"] = state.RetrySec
		inputs["soaEmail"] = state.SoaEmail
		inputs["status"] = state.Status
		inputs["tags"] = state.Tags
		inputs["ttlSec"] = state.TtlSec
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("linode:index/domain:Domain", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Domain) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Domain) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
func (r *Domain) AxfrIps() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["axfrIps"])
}

// A description for this Domain. This is for display purposes only.
func (r *Domain) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
func (r *Domain) Domain() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["domain"])
}

// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
func (r *Domain) ExpireSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["expireSec"])
}

// The group this Domain belongs to. This is for display purposes only.
func (r *Domain) Group() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["group"])
}

// The IP addresses representing the master DNS for this Domain.
func (r *Domain) MasterIps() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["masterIps"])
}

// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
func (r *Domain) RefreshSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["refreshSec"])
}

// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
func (r *Domain) RetrySec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["retrySec"])
}

// Start of Authority email address. This is required for master Domains.
func (r *Domain) SoaEmail() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["soaEmail"])
}

// Used to control whether this Domain is currently being rendered (defaults to "active").
func (r *Domain) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// A list of tags applied to this object. Tags are for organizational purposes only.
func (r *Domain) Tags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tags"])
}

// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
func (r *Domain) TtlSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ttlSec"])
}

// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
func (r *Domain) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering Domain resources.
type DomainState struct {
	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
	AxfrIps interface{}
	// A description for this Domain. This is for display purposes only.
	Description interface{}
	// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
	Domain interface{}
	// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	ExpireSec interface{}
	// The group this Domain belongs to. This is for display purposes only.
	Group interface{}
	// The IP addresses representing the master DNS for this Domain.
	MasterIps interface{}
	// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RefreshSec interface{}
	// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RetrySec interface{}
	// Start of Authority email address. This is required for master Domains.
	SoaEmail interface{}
	// Used to control whether this Domain is currently being rendered (defaults to "active").
	Status interface{}
	// A list of tags applied to this object. Tags are for organizational purposes only.
	Tags interface{}
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec interface{}
	// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
	Type interface{}
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
	AxfrIps interface{}
	// A description for this Domain. This is for display purposes only.
	Description interface{}
	// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
	Domain interface{}
	// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	ExpireSec interface{}
	// The group this Domain belongs to. This is for display purposes only.
	Group interface{}
	// The IP addresses representing the master DNS for this Domain.
	MasterIps interface{}
	// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RefreshSec interface{}
	// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RetrySec interface{}
	// Start of Authority email address. This is required for master Domains.
	SoaEmail interface{}
	// Used to control whether this Domain is currently being rendered (defaults to "active").
	Status interface{}
	// A list of tags applied to this object. Tags are for organizational purposes only.
	Tags interface{}
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec interface{}
	// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
	Type interface{}
}