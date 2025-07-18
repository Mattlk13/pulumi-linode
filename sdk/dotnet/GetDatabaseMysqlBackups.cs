// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    public static class GetDatabaseMysqlBackups
    {
        /// <summary>
        /// &gt; **DEPRECATION NOTICE:** This data source has been deprecated.
        /// 
        /// Provides information about Linode MySQL Database Backups that match a set of filters.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-databases-mysql-instance-backups).
        /// 
        /// ## Example Usage
        /// 
        /// Get information about all backups for a MySQL database:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all_backups = Linode.GetDatabaseMysqlBackups.Invoke(new()
        ///     {
        ///         DatabaseId = 12345,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// Get information about all automatic MySQL Database Backups:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var auto_backups = Linode.GetDatabaseMysqlBackups.Invoke(new()
        ///     {
        ///         DatabaseId = 12345,
        ///         Filters = new[]
        ///         {
        ///             new Linode.Inputs.GetDatabaseMysqlBackupsFilterInputArgs
        ///             {
        ///                 Name = "type",
        ///                 Values = new[]
        ///                 {
        ///                     "auto",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatabaseMysqlBackupsResult> InvokeAsync(GetDatabaseMysqlBackupsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseMysqlBackupsResult>("linode:index/getDatabaseMysqlBackups:getDatabaseMysqlBackups", args ?? new GetDatabaseMysqlBackupsArgs(), options.WithDefaults());

        /// <summary>
        /// &gt; **DEPRECATION NOTICE:** This data source has been deprecated.
        /// 
        /// Provides information about Linode MySQL Database Backups that match a set of filters.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-databases-mysql-instance-backups).
        /// 
        /// ## Example Usage
        /// 
        /// Get information about all backups for a MySQL database:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all_backups = Linode.GetDatabaseMysqlBackups.Invoke(new()
        ///     {
        ///         DatabaseId = 12345,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// Get information about all automatic MySQL Database Backups:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var auto_backups = Linode.GetDatabaseMysqlBackups.Invoke(new()
        ///     {
        ///         DatabaseId = 12345,
        ///         Filters = new[]
        ///         {
        ///             new Linode.Inputs.GetDatabaseMysqlBackupsFilterInputArgs
        ///             {
        ///                 Name = "type",
        ///                 Values = new[]
        ///                 {
        ///                     "auto",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabaseMysqlBackupsResult> Invoke(GetDatabaseMysqlBackupsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseMysqlBackupsResult>("linode:index/getDatabaseMysqlBackups:getDatabaseMysqlBackups", args ?? new GetDatabaseMysqlBackupsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// &gt; **DEPRECATION NOTICE:** This data source has been deprecated.
        /// 
        /// Provides information about Linode MySQL Database Backups that match a set of filters.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-databases-mysql-instance-backups).
        /// 
        /// ## Example Usage
        /// 
        /// Get information about all backups for a MySQL database:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all_backups = Linode.GetDatabaseMysqlBackups.Invoke(new()
        ///     {
        ///         DatabaseId = 12345,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// Get information about all automatic MySQL Database Backups:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var auto_backups = Linode.GetDatabaseMysqlBackups.Invoke(new()
        ///     {
        ///         DatabaseId = 12345,
        ///         Filters = new[]
        ///         {
        ///             new Linode.Inputs.GetDatabaseMysqlBackupsFilterInputArgs
        ///             {
        ///                 Name = "type",
        ///                 Values = new[]
        ///                 {
        ///                     "auto",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabaseMysqlBackupsResult> Invoke(GetDatabaseMysqlBackupsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseMysqlBackupsResult>("linode:index/getDatabaseMysqlBackups:getDatabaseMysqlBackups", args ?? new GetDatabaseMysqlBackupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseMysqlBackupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the database to retrieve backups for.
        /// </summary>
        [Input("databaseId", required: true)]
        public int DatabaseId { get; set; }

        [Input("filters")]
        private List<Inputs.GetDatabaseMysqlBackupsFilterArgs>? _filters;
        public List<Inputs.GetDatabaseMysqlBackupsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDatabaseMysqlBackupsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// If true, only the latest backup will be returned.
        /// 
        /// * `filter` - (Optional) A set of filters used to select database backups that meet certain requirements.
        /// </summary>
        [Input("latest")]
        public bool? Latest { get; set; }

        /// <summary>
        /// The order in which results should be returned. (`asc`, `desc`; default `asc`)
        /// </summary>
        [Input("order")]
        public string? Order { get; set; }

        /// <summary>
        /// The attribute to order the results by. (`created`)
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        public GetDatabaseMysqlBackupsArgs()
        {
        }
        public static new GetDatabaseMysqlBackupsArgs Empty => new GetDatabaseMysqlBackupsArgs();
    }

    public sealed class GetDatabaseMysqlBackupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the database to retrieve backups for.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<int> DatabaseId { get; set; } = null!;

        [Input("filters")]
        private InputList<Inputs.GetDatabaseMysqlBackupsFilterInputArgs>? _filters;
        public InputList<Inputs.GetDatabaseMysqlBackupsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetDatabaseMysqlBackupsFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// If true, only the latest backup will be returned.
        /// 
        /// * `filter` - (Optional) A set of filters used to select database backups that meet certain requirements.
        /// </summary>
        [Input("latest")]
        public Input<bool>? Latest { get; set; }

        /// <summary>
        /// The order in which results should be returned. (`asc`, `desc`; default `asc`)
        /// </summary>
        [Input("order")]
        public Input<string>? Order { get; set; }

        /// <summary>
        /// The attribute to order the results by. (`created`)
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        public GetDatabaseMysqlBackupsInvokeArgs()
        {
        }
        public static new GetDatabaseMysqlBackupsInvokeArgs Empty => new GetDatabaseMysqlBackupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseMysqlBackupsResult
    {
        public readonly ImmutableArray<Outputs.GetDatabaseMysqlBackupsBackupResult> Backups;
        public readonly int DatabaseId;
        public readonly ImmutableArray<Outputs.GetDatabaseMysqlBackupsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? Latest;
        public readonly string? Order;
        public readonly string? OrderBy;

        [OutputConstructor]
        private GetDatabaseMysqlBackupsResult(
            ImmutableArray<Outputs.GetDatabaseMysqlBackupsBackupResult> backups,

            int databaseId,

            ImmutableArray<Outputs.GetDatabaseMysqlBackupsFilterResult> filters,

            string id,

            bool? latest,

            string? order,

            string? orderBy)
        {
            Backups = backups;
            DatabaseId = databaseId;
            Filters = filters;
            Id = id;
            Latest = latest;
            Order = order;
            OrderBy = orderBy;
        }
    }
}
