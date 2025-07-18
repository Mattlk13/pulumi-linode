// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about Linode Child Accounts that match a set of filters.
 * For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-child-accounts).
 *
 * **NOTE: Parent/Child related features may not be generally available.**
 *
 * ## Example Usage
 *
 * The following example shows how one might use this data source to access Child Accounts under the current Account.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 *
 * const all = linode.getChildAccounts({});
 * const filtered = linode.getChildAccounts({
 *     filters: [
 *         {
 *             name: "email",
 *             values: ["example@linode.com"],
 *         },
 *         {
 *             name: "first_name",
 *             values: ["John"],
 *         },
 *         {
 *             name: "last_name",
 *             values: ["Smith"],
 *         },
 *     ],
 * });
 * export const allAccounts = all.then(all => all.childAccounts.map(__item => __item.euuid));
 * export const filteredAccounts = filtered.then(filtered => filtered.childAccounts.map(__item => __item.euuid));
 * ```
 *
 * ## Filterable Fields
 *
 * * `euuid`
 *
 * * `email`
 *
 * * `firstName`
 *
 * * `lastName`
 *
 * * `company`
 *
 * * `address1`
 *
 * * `address2`
 *
 * * `phone`
 *
 * * `city`
 *
 * * `state`
 *
 * * `country`
 *
 * * `zip`
 *
 * * `capabilities`
 *
 * * `activeSince`
 */
export function getChildAccounts(args?: GetChildAccountsArgs, opts?: pulumi.InvokeOptions): Promise<GetChildAccountsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("linode:index/getChildAccounts:getChildAccounts", {
        "childAccounts": args.childAccounts,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getChildAccounts.
 */
export interface GetChildAccountsArgs {
    childAccounts?: inputs.GetChildAccountsChildAccount[];
    filters?: inputs.GetChildAccountsFilter[];
}

/**
 * A collection of values returned by getChildAccounts.
 */
export interface GetChildAccountsResult {
    readonly childAccounts?: outputs.GetChildAccountsChildAccount[];
    readonly filters?: outputs.GetChildAccountsFilter[];
    readonly id: string;
}
/**
 * Provides information about Linode Child Accounts that match a set of filters.
 * For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-child-accounts).
 *
 * **NOTE: Parent/Child related features may not be generally available.**
 *
 * ## Example Usage
 *
 * The following example shows how one might use this data source to access Child Accounts under the current Account.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 *
 * const all = linode.getChildAccounts({});
 * const filtered = linode.getChildAccounts({
 *     filters: [
 *         {
 *             name: "email",
 *             values: ["example@linode.com"],
 *         },
 *         {
 *             name: "first_name",
 *             values: ["John"],
 *         },
 *         {
 *             name: "last_name",
 *             values: ["Smith"],
 *         },
 *     ],
 * });
 * export const allAccounts = all.then(all => all.childAccounts.map(__item => __item.euuid));
 * export const filteredAccounts = filtered.then(filtered => filtered.childAccounts.map(__item => __item.euuid));
 * ```
 *
 * ## Filterable Fields
 *
 * * `euuid`
 *
 * * `email`
 *
 * * `firstName`
 *
 * * `lastName`
 *
 * * `company`
 *
 * * `address1`
 *
 * * `address2`
 *
 * * `phone`
 *
 * * `city`
 *
 * * `state`
 *
 * * `country`
 *
 * * `zip`
 *
 * * `capabilities`
 *
 * * `activeSince`
 */
export function getChildAccountsOutput(args?: GetChildAccountsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetChildAccountsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("linode:index/getChildAccounts:getChildAccounts", {
        "childAccounts": args.childAccounts,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getChildAccounts.
 */
export interface GetChildAccountsOutputArgs {
    childAccounts?: pulumi.Input<pulumi.Input<inputs.GetChildAccountsChildAccountArgs>[]>;
    filters?: pulumi.Input<pulumi.Input<inputs.GetChildAccountsFilterArgs>[]>;
}
