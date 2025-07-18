// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about a list of current Linode Kubernetes (LKE) clusters on your account that match a set of filters.
 * For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-lke-clusters).
 *
 * ## Example Usage
 *
 * Get information about all LKE clusters with a specific tag:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 *
 * const specific = linode.getLkeClusters({
 *     filters: [{
 *         name: "tags",
 *         values: ["test-tag"],
 *     }],
 * });
 * export const lkeCluster = specific.then(specific => specific.lkeClusters?.[0]?.id);
 * ```
 *
 * ## Filterable Fields
 *
 * * `k8sVersion`
 *
 * * `label`
 *
 * * `region`
 *
 * * `tags`
 *
 * * `status`
 *
 * * `created`
 *
 * * `updated`
 */
export function getLkeClusters(args?: GetLkeClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetLkeClustersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("linode:index/getLkeClusters:getLkeClusters", {
        "filters": args.filters,
        "lkeClusters": args.lkeClusters,
        "order": args.order,
        "orderBy": args.orderBy,
    }, opts);
}

/**
 * A collection of arguments for invoking getLkeClusters.
 */
export interface GetLkeClustersArgs {
    filters?: inputs.GetLkeClustersFilter[];
    lkeClusters?: inputs.GetLkeClustersLkeCluster[];
    /**
     * The order in which results should be returned. (`asc`, `desc`; default `asc`)
     */
    order?: string;
    /**
     * The attribute to order the results by. See the Filterable Fields section for a list of valid fields.
     */
    orderBy?: string;
}

/**
 * A collection of values returned by getLkeClusters.
 */
export interface GetLkeClustersResult {
    readonly filters?: outputs.GetLkeClustersFilter[];
    /**
     * The LKE Cluster's ID.
     */
    readonly id: string;
    readonly lkeClusters?: outputs.GetLkeClustersLkeCluster[];
    readonly order?: string;
    readonly orderBy?: string;
}
/**
 * Provides information about a list of current Linode Kubernetes (LKE) clusters on your account that match a set of filters.
 * For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-lke-clusters).
 *
 * ## Example Usage
 *
 * Get information about all LKE clusters with a specific tag:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 *
 * const specific = linode.getLkeClusters({
 *     filters: [{
 *         name: "tags",
 *         values: ["test-tag"],
 *     }],
 * });
 * export const lkeCluster = specific.then(specific => specific.lkeClusters?.[0]?.id);
 * ```
 *
 * ## Filterable Fields
 *
 * * `k8sVersion`
 *
 * * `label`
 *
 * * `region`
 *
 * * `tags`
 *
 * * `status`
 *
 * * `created`
 *
 * * `updated`
 */
export function getLkeClustersOutput(args?: GetLkeClustersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLkeClustersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("linode:index/getLkeClusters:getLkeClusters", {
        "filters": args.filters,
        "lkeClusters": args.lkeClusters,
        "order": args.order,
        "orderBy": args.orderBy,
    }, opts);
}

/**
 * A collection of arguments for invoking getLkeClusters.
 */
export interface GetLkeClustersOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetLkeClustersFilterArgs>[]>;
    lkeClusters?: pulumi.Input<pulumi.Input<inputs.GetLkeClustersLkeClusterArgs>[]>;
    /**
     * The order in which results should be returned. (`asc`, `desc`; default `asc`)
     */
    order?: pulumi.Input<string>;
    /**
     * The attribute to order the results by. See the Filterable Fields section for a list of valid fields.
     */
    orderBy?: pulumi.Input<string>;
}
