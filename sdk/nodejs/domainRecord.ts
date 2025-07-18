// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Linode Domain Record resource.  This can be used to create, modify, and delete Linodes Domain Records.
 * For more information, see [DNS Manager](https://www.linode.com/docs/platform/manager/dns-manager/) and the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/post-domain-record).
 *
 * ## Example Usage
 *
 * The following example shows how one might use this resource to configure a Domain Record attached to a Linode Domain.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 *
 * const foobar = new linode.Domain("foobar", {
 *     type: "master",
 *     domain: "foobar.example",
 *     soaEmail: "example@foobar.example",
 * });
 * const foobarDomainRecord = new linode.DomainRecord("foobar", {
 *     domainId: foobar.id,
 *     name: "www",
 *     recordType: "CNAME",
 *     target: "foobar.example",
 * });
 * ```
 *
 * ## Import
 *
 * Linodes Domain Records can be imported using the Linode Domain `id` followed by the Domain Record `id` separated by a comma, e.g.
 *
 * ```sh
 * $ pulumi import linode:index/domainRecord:DomainRecord www-foobar 1234567,7654321
 * ```
 */
export class DomainRecord extends pulumi.CustomResource {
    /**
     * Get an existing DomainRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainRecordState, opts?: pulumi.CustomResourceOptions): DomainRecord {
        return new DomainRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'linode:index/domainRecord:DomainRecord';

    /**
     * Returns true if the given object is an instance of DomainRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainRecord.__pulumiType;
    }

    /**
     * The ID of the Domain to access.  *Changing `domainId` forces the creation of a new Linode Domain Record.*.
     */
    public readonly domainId!: pulumi.Output<number>;
    /**
     * The name of this Record. Setting this is invalid for `SRV` records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The port this Record points to.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The priority of the target host. Lower values are preferred.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The protocol this Record's service communicates with. Only valid for SRV records.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. See all supported record types [here](https://techdocs.akamai.com/linode-api/reference/post-domain-record). *Changing `recordType` forces the creation of a new Linode Domain Record.*.
     */
    public readonly recordType!: pulumi.Output<string>;
    /**
     * The service this Record identified. Only valid for SRV records.
     */
    public readonly service!: pulumi.Output<string | undefined>;
    /**
     * The tag portion of a CAA record. It is invalid to set this on other record types.
     */
    public readonly tag!: pulumi.Output<string | undefined>;
    /**
     * The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
     *
     * - - -
     */
    public readonly target!: pulumi.Output<string>;
    /**
     * 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    public readonly ttlSec!: pulumi.Output<number | undefined>;
    /**
     * The relative weight of this Record. Higher values are preferred.
     */
    public readonly weight!: pulumi.Output<number | undefined>;

    /**
     * Create a DomainRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainRecordArgs | DomainRecordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainRecordState | undefined;
            resourceInputs["domainId"] = state ? state.domainId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["recordType"] = state ? state.recordType : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["tag"] = state ? state.tag : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
            resourceInputs["ttlSec"] = state ? state.ttlSec : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as DomainRecordArgs | undefined;
            if ((!args || args.domainId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainId'");
            }
            if ((!args || args.recordType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recordType'");
            }
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            resourceInputs["domainId"] = args ? args.domainId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["recordType"] = args ? args.recordType : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["ttlSec"] = args ? args.ttlSec : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainRecord.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainRecord resources.
 */
export interface DomainRecordState {
    /**
     * The ID of the Domain to access.  *Changing `domainId` forces the creation of a new Linode Domain Record.*.
     */
    domainId?: pulumi.Input<number>;
    /**
     * The name of this Record. Setting this is invalid for `SRV` records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
     */
    name?: pulumi.Input<string>;
    /**
     * The port this Record points to.
     */
    port?: pulumi.Input<number>;
    /**
     * The priority of the target host. Lower values are preferred.
     */
    priority?: pulumi.Input<number>;
    /**
     * The protocol this Record's service communicates with. Only valid for SRV records.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. See all supported record types [here](https://techdocs.akamai.com/linode-api/reference/post-domain-record). *Changing `recordType` forces the creation of a new Linode Domain Record.*.
     */
    recordType?: pulumi.Input<string>;
    /**
     * The service this Record identified. Only valid for SRV records.
     */
    service?: pulumi.Input<string>;
    /**
     * The tag portion of a CAA record. It is invalid to set this on other record types.
     */
    tag?: pulumi.Input<string>;
    /**
     * The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
     *
     * - - -
     */
    target?: pulumi.Input<string>;
    /**
     * 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    ttlSec?: pulumi.Input<number>;
    /**
     * The relative weight of this Record. Higher values are preferred.
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DomainRecord resource.
 */
export interface DomainRecordArgs {
    /**
     * The ID of the Domain to access.  *Changing `domainId` forces the creation of a new Linode Domain Record.*.
     */
    domainId: pulumi.Input<number>;
    /**
     * The name of this Record. Setting this is invalid for `SRV` records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
     */
    name?: pulumi.Input<string>;
    /**
     * The port this Record points to.
     */
    port?: pulumi.Input<number>;
    /**
     * The priority of the target host. Lower values are preferred.
     */
    priority?: pulumi.Input<number>;
    /**
     * The protocol this Record's service communicates with. Only valid for SRV records.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. See all supported record types [here](https://techdocs.akamai.com/linode-api/reference/post-domain-record). *Changing `recordType` forces the creation of a new Linode Domain Record.*.
     */
    recordType: pulumi.Input<string>;
    /**
     * The service this Record identified. Only valid for SRV records.
     */
    service?: pulumi.Input<string>;
    /**
     * The tag portion of a CAA record. It is invalid to set this on other record types.
     */
    tag?: pulumi.Input<string>;
    /**
     * The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
     *
     * - - -
     */
    target: pulumi.Input<string>;
    /**
     * 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    ttlSec?: pulumi.Input<number>;
    /**
     * The relative weight of this Record. Higher values are preferred.
     */
    weight?: pulumi.Input<number>;
}
