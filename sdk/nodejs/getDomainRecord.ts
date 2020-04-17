// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about a Linode Domain Record.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 * 
 * const myRecord = linode.getDomainRecord({
 *     domainId: 3150401,
 *     id: 14950401,
 * });
 * const myWwwRecord = linode.getDomainRecord({
 *     domainId: 3150401,
 *     name: "www",
 * });
 * ```
 * 
 * ## Attributes
 * 
 * The Linode Volume resource exports the following attributes:
 * 
 * - `id` - The unique ID of the Domain Record.
 * 
 * - `name` - The name of the Record.
 * 
 * - `domainId` - The associated domain's unique ID.
 * 
 * - `type` - The type of Record this is in the DNS system.
 * 
 * - `ttlSec` - The amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers.
 * 
 * - `target` - The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
 * 
 * - `priority` - The priority of the target host. Lower values are preferred.
 * 
 * - `weight` - The relative weight of this Record. Higher values are preferred.
 * 
 * - `port` - The port this Record points to.
 * 
 * - `protocol` - The protocol this Record's service communicates with. Only valid for SRV records.
 * 
 * - `service` - The service this Record identified. Only valid for SRV records.
 * 
 * - `tag` - The tag portion of a CAA record.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/d/domain_record.html.md.
 */
export function getDomainRecord(args: GetDomainRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainRecordResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("linode:index/getDomainRecord:getDomainRecord", {
        "domainId": args.domainId,
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomainRecord.
 */
export interface GetDomainRecordArgs {
    readonly domainId: number;
    readonly id?: number;
    readonly name?: string;
}

/**
 * A collection of values returned by getDomainRecord.
 */
export interface GetDomainRecordResult {
    readonly domainId: number;
    readonly id?: number;
    readonly name?: string;
    readonly port: number;
    readonly priority: number;
    readonly protocol: string;
    readonly service: string;
    readonly tag: string;
    readonly target: string;
    readonly ttlSec: number;
    readonly type: string;
    readonly weight: number;
}