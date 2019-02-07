// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Linode Domain Record resource.  This can be used to create, modify, and delete Linodes Domain Records.
// For more information, see [DNS Manager](https://www.linode.com/docs/platform/manager/dns-manager/) and the [Linode APIv4 docs](https://developers.linode.com/api/v4#operation/createDomainRecord).
// 
// ## Attributes
// 
// This resource exports no additional attributes.
type DomainRecord struct {
	s *pulumi.ResourceState
}

// NewDomainRecord registers a new resource with the given unique name, arguments, and options.
func NewDomainRecord(ctx *pulumi.Context,
	name string, args *DomainRecordArgs, opts ...pulumi.ResourceOpt) (*DomainRecord, error) {
	if args == nil || args.DomainId == nil {
		return nil, errors.New("missing required argument 'DomainId'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.RecordType == nil {
		return nil, errors.New("missing required argument 'RecordType'")
	}
	if args == nil || args.Target == nil {
		return nil, errors.New("missing required argument 'Target'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["domainId"] = nil
		inputs["name"] = nil
		inputs["port"] = nil
		inputs["priority"] = nil
		inputs["protocol"] = nil
		inputs["recordType"] = nil
		inputs["service"] = nil
		inputs["tag"] = nil
		inputs["target"] = nil
		inputs["ttlSec"] = nil
		inputs["weight"] = nil
	} else {
		inputs["domainId"] = args.DomainId
		inputs["name"] = args.Name
		inputs["port"] = args.Port
		inputs["priority"] = args.Priority
		inputs["protocol"] = args.Protocol
		inputs["recordType"] = args.RecordType
		inputs["service"] = args.Service
		inputs["tag"] = args.Tag
		inputs["target"] = args.Target
		inputs["ttlSec"] = args.TtlSec
		inputs["weight"] = args.Weight
	}
	s, err := ctx.RegisterResource("linode:index/domainRecord:DomainRecord", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DomainRecord{s: s}, nil
}

// GetDomainRecord gets an existing DomainRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainRecord(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DomainRecordState, opts ...pulumi.ResourceOpt) (*DomainRecord, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["domainId"] = state.DomainId
		inputs["name"] = state.Name
		inputs["port"] = state.Port
		inputs["priority"] = state.Priority
		inputs["protocol"] = state.Protocol
		inputs["recordType"] = state.RecordType
		inputs["service"] = state.Service
		inputs["tag"] = state.Tag
		inputs["target"] = state.Target
		inputs["ttlSec"] = state.TtlSec
		inputs["weight"] = state.Weight
	}
	s, err := ctx.ReadResource("linode:index/domainRecord:DomainRecord", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DomainRecord{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DomainRecord) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DomainRecord) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the Domain to access.  *Changing `domain_id` forces the creation of a new Linode Domain Record.*.
func (r *DomainRecord) DomainId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["domainId"])
}

// The name of this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
func (r *DomainRecord) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The port this Record points to.
func (r *DomainRecord) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// The priority of the target host. Lower values are preferred.
func (r *DomainRecord) Priority() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["priority"])
}

// The protocol this Record's service communicates with. Only valid for SRV records.
func (r *DomainRecord) Protocol() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["protocol"])
}

// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `record_type` forces the creation of a new Linode Domain Record.*.
func (r *DomainRecord) RecordType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["recordType"])
}

// The service this Record identified. Only valid for SRV records.
func (r *DomainRecord) Service() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["service"])
}

// The tag portion of a CAA record. It is invalid to set this on other record types.
func (r *DomainRecord) Tag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tag"])
}

// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
// - - -
func (r *DomainRecord) Target() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["target"])
}

// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
func (r *DomainRecord) TtlSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ttlSec"])
}

// The relative weight of this Record. Higher values are preferred.
func (r *DomainRecord) Weight() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["weight"])
}

// Input properties used for looking up and filtering DomainRecord resources.
type DomainRecordState struct {
	// The ID of the Domain to access.  *Changing `domain_id` forces the creation of a new Linode Domain Record.*.
	DomainId interface{}
	// The name of this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
	Name interface{}
	// The port this Record points to.
	Port interface{}
	// The priority of the target host. Lower values are preferred.
	Priority interface{}
	// The protocol this Record's service communicates with. Only valid for SRV records.
	Protocol interface{}
	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `record_type` forces the creation of a new Linode Domain Record.*.
	RecordType interface{}
	// The service this Record identified. Only valid for SRV records.
	Service interface{}
	// The tag portion of a CAA record. It is invalid to set this on other record types.
	Tag interface{}
	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	// - - -
	Target interface{}
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec interface{}
	// The relative weight of this Record. Higher values are preferred.
	Weight interface{}
}

// The set of arguments for constructing a DomainRecord resource.
type DomainRecordArgs struct {
	// The ID of the Domain to access.  *Changing `domain_id` forces the creation of a new Linode Domain Record.*.
	DomainId interface{}
	// The name of this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
	Name interface{}
	// The port this Record points to.
	Port interface{}
	// The priority of the target host. Lower values are preferred.
	Priority interface{}
	// The protocol this Record's service communicates with. Only valid for SRV records.
	Protocol interface{}
	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `record_type` forces the creation of a new Linode Domain Record.*.
	RecordType interface{}
	// The service this Record identified. Only valid for SRV records.
	Service interface{}
	// The tag portion of a CAA record. It is invalid to set this on other record types.
	Tag interface{}
	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	// - - -
	Target interface{}
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec interface{}
	// The relative weight of this Record. Higher values are preferred.
	Weight interface{}
}