// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about a Linode PostgreSQL Database.
 * For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-databases-postgre-sql-instance-backups).
 *
 * ## Example Usage
 *
 * Get information about a PostgreSQL database:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 *
 * const my_db = linode.getDatabasePostgresqlV2({
 *     id: "12345",
 * });
 * ```
 *
 * ## pendingUpdates
 *
 * The following arguments are exposed by each entry in the `pendingUpdates` attribute:
 *
 * * `deadline` - The time when a mandatory update needs to be applied.
 *
 * * `description` - A description of the update.
 *
 * * `plannedFor` - The date and time a maintenance update will be applied.
 *
 * ## updates
 *
 * The following arguments are supported in the `updates` specification block:
 *
 * * `dayOfWeek` - The day to perform maintenance. (`monday`, `tuesday`, ...)
 *
 * * `duration` - The maximum maintenance window time in hours. (`1`..`3`)
 *
 * * `frequency` - The frequency at which maintenance occurs. (`weekly`)
 *
 * * `hourOfDay` - The hour to begin maintenance based in UTC time. (`0`..`23`)
 */
export function getDatabasePostgresqlV2(args: GetDatabasePostgresqlV2Args, opts?: pulumi.InvokeOptions): Promise<GetDatabasePostgresqlV2Result> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("linode:index/getDatabasePostgresqlV2:getDatabasePostgresqlV2", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabasePostgresqlV2.
 */
export interface GetDatabasePostgresqlV2Args {
    /**
     * The ID of the PostgreSQL database.
     */
    id: string;
}

/**
 * A collection of values returned by getDatabasePostgresqlV2.
 */
export interface GetDatabasePostgresqlV2Result {
    /**
     * A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format. Use `linode.DatabaseAccessControls` to manage your allow list separately.
     */
    readonly allowLists: string[];
    /**
     * The base64-encoded SSL CA certificate for the Managed Database.
     */
    readonly caCert: string;
    /**
     * The number of Linode Instance nodes deployed to the Managed Database. (default `1`)
     */
    readonly clusterSize: number;
    /**
     * When this Managed Database was created.
     */
    readonly created: string;
    /**
     * Whether the Managed Databases is encrypted.
     */
    readonly encrypted: boolean;
    /**
     * The Managed Database engine. (e.g. `postgresql`)
     */
    readonly engine: string;
    /**
     * Specifies a fraction of the table size to add to autovacuumAnalyzeThreshold when deciding whether to trigger an ANALYZE. The default is 0.2 (20% of table size)
     */
    readonly engineConfigPgAutovacuumAnalyzeScaleFactor: number;
    /**
     * Specifies the minimum number of inserted, updated or deleted tuples needed to trigger an ANALYZE in any one table. The default is 50 tuples.
     */
    readonly engineConfigPgAutovacuumAnalyzeThreshold: number;
    /**
     * Specifies the maximum number of autovacuum processes (other than the autovacuum launcher) that may be running at any one time. The default is three. This parameter can only be set at server start.
     */
    readonly engineConfigPgAutovacuumMaxWorkers: number;
    /**
     * Specifies the minimum delay between autovacuum runs on any given database. The delay is measured in seconds, and the default is one minute
     */
    readonly engineConfigPgAutovacuumNaptime: number;
    /**
     * Specifies the cost delay value that will be used in automatic VACUUM operations. If -1 is specified, the regular vacuumCostDelay value will be used. The default value is 20 milliseconds
     */
    readonly engineConfigPgAutovacuumVacuumCostDelay: number;
    /**
     * Specifies the cost limit value that will be used in automatic VACUUM operations. If -1 is specified (which is the default), the regular vacuumCostLimit value will be used.
     */
    readonly engineConfigPgAutovacuumVacuumCostLimit: number;
    /**
     * Specifies a fraction of the table size to add to autovacuumVacuumThreshold when deciding whether to trigger a VACUUM. The default is 0.2 (20% of table size)
     */
    readonly engineConfigPgAutovacuumVacuumScaleFactor: number;
    /**
     * Specifies the minimum number of updated or deleted tuples needed to trigger a VACUUM in any one table. The default is 50 tuples.
     */
    readonly engineConfigPgAutovacuumVacuumThreshold: number;
    /**
     * Specifies the delay between activity rounds for the background writer in milliseconds. Default is 200.
     */
    readonly engineConfigPgBgwriterDelay: number;
    /**
     * Whenever more than bgwriterFlushAfter bytes have been written by the background writer, attempt to force the OS to issue these writes to the underlying storage. Specified in kilobytes, default is 512. Setting of 0 disables forced writeback.
     */
    readonly engineConfigPgBgwriterFlushAfter: number;
    /**
     * In each round, no more than this many buffers will be written by the background writer. Setting this to zero disables background writing. Default is 100.
     */
    readonly engineConfigPgBgwriterLruMaxpages: number;
    /**
     * The average recent need for new buffers is multiplied by bgwriterLruMultiplier to arrive at an estimate of the number that will be needed during the next round, (up to bgwriter_lru_maxpages). 1.0 represents a “just in time” policy of writing exactly the number of buffers predicted to be needed. Larger values provide some cushion against spikes in demand, while smaller values intentionally leave writes to be done by server processes. The default is 2.0.
     */
    readonly engineConfigPgBgwriterLruMultiplier: number;
    /**
     * This is the amount of time, in milliseconds, to wait on a lock before checking to see if there is a deadlock condition.
     */
    readonly engineConfigPgDeadlockTimeout: number;
    /**
     * Specifies the default TOAST compression method for values of compressible columns (the default is lz4).
     */
    readonly engineConfigPgDefaultToastCompression: string;
    /**
     * Time out sessions with open transactions after this number of milliseconds.
     */
    readonly engineConfigPgIdleInTransactionSessionTimeout: number;
    /**
     * Controls system-wide use of Just-in-Time Compilation (JIT).
     */
    readonly engineConfigPgJit: boolean;
    /**
     * PostgreSQL maximum number of files that can be open per process.
     */
    readonly engineConfigPgMaxFilesPerProcess: number;
    /**
     * PostgreSQL maximum locks per transaction.
     */
    readonly engineConfigPgMaxLocksPerTransaction: number;
    /**
     * PostgreSQL maximum logical replication workers (taken from the pool of max_parallel_workers).
     */
    readonly engineConfigPgMaxLogicalReplicationWorkers: number;
    /**
     * Sets the maximum number of workers that the system can support for parallel queries.
     */
    readonly engineConfigPgMaxParallelWorkers: number;
    /**
     * Sets the maximum number of workers that can be started by a single Gather or Gather Merge node.
     */
    readonly engineConfigPgMaxParallelWorkersPerGather: number;
    /**
     * PostgreSQL maximum predicate locks per transaction.
     */
    readonly engineConfigPgMaxPredLocksPerTransaction: number;
    /**
     * PostgreSQL maximum replication slots.
     */
    readonly engineConfigPgMaxReplicationSlots: number;
    /**
     * PostgreSQL maximum WAL size (MB) reserved for replication slots. Default is -1 (unlimited). walKeepSize minimum WAL size setting takes precedence over this.
     */
    readonly engineConfigPgMaxSlotWalKeepSize: number;
    /**
     * Maximum depth of the stack in bytes.
     */
    readonly engineConfigPgMaxStackDepth: number;
    /**
     * Max standby archive delay in milliseconds.
     */
    readonly engineConfigPgMaxStandbyArchiveDelay: number;
    /**
     * Max standby streaming delay in milliseconds.
     */
    readonly engineConfigPgMaxStandbyStreamingDelay: number;
    /**
     * PostgreSQL maximum WAL senders.
     */
    readonly engineConfigPgMaxWalSenders: number;
    /**
     * Sets the maximum number of background processes that the system can support.
     */
    readonly engineConfigPgMaxWorkerProcesses: number;
    /**
     * Chooses the algorithm for encrypting passwords.
     */
    readonly engineConfigPgPasswordEncryption: string;
    /**
     * Sets the time interval to run pg_partman's scheduled tasks.
     */
    readonly engineConfigPgPgPartmanBgwInterval: number;
    /**
     * Controls which role to use for pg_partman's scheduled background tasks.
     */
    readonly engineConfigPgPgPartmanBgwRole: string;
    /**
     * Enables or disables query plan monitoring.
     */
    readonly engineConfigPgPgStatMonitorPgsmEnableQueryPlan: boolean;
    /**
     * Sets the maximum number of buckets.
     */
    readonly engineConfigPgPgStatMonitorPgsmMaxBuckets: number;
    /**
     * Controls which statements are counted. Specify top to track top-level statements (those issued directly by clients), all to also track nested statements (such as statements invoked within functions), or none to disable statement statistics collection. The default value is top.
     */
    readonly engineConfigPgPgStatStatementsTrack: string;
    /**
     * Enable the pgStatMonitor extension. Enabling this extension will cause the cluster to be restarted. When this extension is enabled, pgStatStatements results for utility commands are unreliable.
     */
    readonly engineConfigPgStatMonitorEnable: boolean;
    /**
     * PostgreSQL temporary file limit in KiB, -1 for unlimited.
     */
    readonly engineConfigPgTempFileLimit: number;
    /**
     * PostgreSQL service timezone.
     */
    readonly engineConfigPgTimezone: string;
    /**
     * Specifies the number of bytes reserved to track the currently executing command for each active session.
     */
    readonly engineConfigPgTrackActivityQuerySize: number;
    /**
     * Record commit time of transactions.
     */
    readonly engineConfigPgTrackCommitTimestamp: string;
    /**
     * Enables tracking of function call counts and time used.
     */
    readonly engineConfigPgTrackFunctions: string;
    /**
     * Enables timing of database I/O calls. This parameter is off by default, because it will repeatedly query the operating system for the current time, which may cause significant overhead on some platforms.
     */
    readonly engineConfigPgTrackIoTiming: string;
    /**
     * Terminate replication connections that are inactive for longer than this amount of time, in milliseconds. Setting this value to zero disables the timeout.
     */
    readonly engineConfigPgWalSenderTimeout: number;
    /**
     * WAL flush interval in milliseconds. Note that setting this value to lower than the default 200ms may negatively impact performance.
     */
    readonly engineConfigPgWalWriterDelay: number;
    /**
     * Number of seconds of master unavailability before triggering database failover to standby.
     */
    readonly engineConfigPglookoutMaxFailoverReplicationTimeLag: number;
    /**
     * Percentage of total RAM that the database server uses for shared memory buffers. Valid range is 20-60 (float), which corresponds to 20% - 60%. This setting adjusts the sharedBuffers configuration value.
     */
    readonly engineConfigSharedBuffersPercentage: number;
    /**
     * Sets the maximum amount of memory to be used by a query operation (such as a sort or hash table) before writing to temporary disk files, in MB. Default is 1MB + 0.075% of total RAM (up to 32MB).
     */
    readonly engineConfigWorkMem: number;
    /**
     * The Managed Database engine in engine/version format. (e.g. `postgresql/16`)
     */
    readonly engineId: string;
    /**
     * The database timestamp from which it was restored.
     */
    readonly forkRestoreTime: string;
    /**
     * The ID of the database that was forked from.
     */
    readonly forkSource: number;
    /**
     * The primary host for the Managed Database.
     */
    readonly hostPrimary: string;
    /**
     * The secondary/private host for the managed database.
     */
    readonly hostSecondary: string;
    readonly id: string;
    /**
     * A unique, user-defined string referring to the Managed Database.
     */
    readonly label: string;
    readonly members: {[key: string]: string};
    readonly oldestRestoreTime: string;
    readonly pendingUpdates: outputs.GetDatabasePostgresqlV2PendingUpdate[];
    /**
     * The back-end platform for relational databases used by the service.
     */
    readonly platform: string;
    /**
     * The access port for this Managed Database.
     */
    readonly port: number;
    /**
     * The region to use for the Managed Database.
     */
    readonly region: string;
    /**
     * The randomly-generated root password for the Managed Database instance.
     */
    readonly rootPassword: string;
    /**
     * The root username for the Managed Database instance.
     */
    readonly rootUsername: string;
    /**
     * Whether to require SSL credentials to establish a connection to the Managed Database.
     */
    readonly sslConnection: boolean;
    /**
     * The operating status of the Managed Database.
     */
    readonly status: string;
    /**
     * Whether this Managed Database is suspended.
     */
    readonly suspended: boolean;
    /**
     * The Linode Instance type used for the nodes of the Managed Database.
     */
    readonly type: string;
    /**
     * When this Managed Database was last updated.
     */
    readonly updated: string;
    readonly updates: outputs.GetDatabasePostgresqlV2Updates;
    /**
     * The Managed Database engine version. (e.g. `13.2`)
     */
    readonly version: string;
}
/**
 * Provides information about a Linode PostgreSQL Database.
 * For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-databases-postgre-sql-instance-backups).
 *
 * ## Example Usage
 *
 * Get information about a PostgreSQL database:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 *
 * const my_db = linode.getDatabasePostgresqlV2({
 *     id: "12345",
 * });
 * ```
 *
 * ## pendingUpdates
 *
 * The following arguments are exposed by each entry in the `pendingUpdates` attribute:
 *
 * * `deadline` - The time when a mandatory update needs to be applied.
 *
 * * `description` - A description of the update.
 *
 * * `plannedFor` - The date and time a maintenance update will be applied.
 *
 * ## updates
 *
 * The following arguments are supported in the `updates` specification block:
 *
 * * `dayOfWeek` - The day to perform maintenance. (`monday`, `tuesday`, ...)
 *
 * * `duration` - The maximum maintenance window time in hours. (`1`..`3`)
 *
 * * `frequency` - The frequency at which maintenance occurs. (`weekly`)
 *
 * * `hourOfDay` - The hour to begin maintenance based in UTC time. (`0`..`23`)
 */
export function getDatabasePostgresqlV2Output(args: GetDatabasePostgresqlV2OutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDatabasePostgresqlV2Result> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("linode:index/getDatabasePostgresqlV2:getDatabasePostgresqlV2", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabasePostgresqlV2.
 */
export interface GetDatabasePostgresqlV2OutputArgs {
    /**
     * The ID of the PostgreSQL database.
     */
    id: pulumi.Input<string>;
}
