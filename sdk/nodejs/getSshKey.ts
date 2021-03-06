// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `linode.SshKey` provides access to a specifically labeled SSH Key in the Profile of the User identified by the access token.
 *
 * ## Example Usage
 *
 * The following example shows how the resource might be used to obtain the name of the SSH Key configured on the Linode user profile.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 *
 * const foo = pulumi.output(linode.getSshKey({
 *     label: "foo",
 * }, { async: true }));
 * ```
 */
export function getSshKey(args: GetSshKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetSshKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("linode:index/getSshKey:getSshKey", {
        "label": args.label,
    }, opts);
}

/**
 * A collection of arguments for invoking getSshKey.
 */
export interface GetSshKeyArgs {
    /**
     * The label of the SSH Key to select.
     */
    readonly label: string;
}

/**
 * A collection of values returned by getSshKey.
 */
export interface GetSshKeyResult {
    /**
     * The date this key was added.
     */
    readonly created: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly label: string;
    /**
     * The public SSH Key, which is used to authenticate to the root user of the Linodes you deploy.
     */
    readonly sshKey: string;
}
