// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getRegion(args: GetRegionArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionResult> {
    return pulumi.runtime.invoke("linode:index/getRegion:getRegion", {
        "country": args.country,
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegion.
 */
export interface GetRegionArgs {
    readonly country?: string;
    readonly id: string;
}

/**
 * A collection of values returned by getRegion.
 */
export interface GetRegionResult {
    readonly country: string;
}