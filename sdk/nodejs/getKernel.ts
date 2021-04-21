// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides information about a Linode kernel
 *
 * ## Example Usage
 *
 * The following example shows how one might use this data source to access information about a Linode kernel.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 *
 * const latest = pulumi.output(linode.getKernel({
 *     id: "linode/latest-64bit",
 * }, { async: true }));
 * ```
 * ## Attributes
 *
 * The Linode Kernel resource exports the following attributes:
 *
 * * `architecture` - The architecture of this Kernel.
 *
 * * `deprecated` - Whether or not this Kernel is deprecated.
 *
 * * `kvm` - If this Kernel is suitable for KVM Linodes.
 *
 * * `label` - The friendly name of this Kernel.
 *
 * * `pvops` - If this Kernel is suitable for paravirtualized operations.
 *
 * * `version` - Linux Kernel version
 *
 * * `xen` - If this Kernel is suitable for Xen Linodes.
 */
export function getKernel(args: GetKernelArgs, opts?: pulumi.InvokeOptions): Promise<GetKernelResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("linode:index/getKernel:getKernel", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getKernel.
 */
export interface GetKernelArgs {
    /**
     * The unique ID of this Kernel.
     */
    readonly id: string;
}

/**
 * A collection of values returned by getKernel.
 */
export interface GetKernelResult {
    readonly architecture: string;
    readonly deprecated: boolean;
    readonly id: string;
    readonly kvm: boolean;
    readonly label: string;
    readonly pvops: boolean;
    readonly version: string;
    readonly xen: boolean;
}