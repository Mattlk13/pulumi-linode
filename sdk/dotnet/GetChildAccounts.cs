// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    public static class GetChildAccounts
    {
        /// <summary>
        /// Provides information about Linode Child Accounts that match a set of filters.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-child-accounts).
        /// 
        /// **NOTE: Parent/Child related features may not be generally available.**
        /// 
        /// ## Example Usage
        /// 
        /// The following example shows how one might use this data source to access Child Accounts under the current Account.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Linode.GetChildAccounts.Invoke();
        /// 
        ///     var filtered = Linode.GetChildAccounts.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Linode.Inputs.GetChildAccountsFilterInputArgs
        ///             {
        ///                 Name = "email",
        ///                 Values = new[]
        ///                 {
        ///                     "example@linode.com",
        ///                 },
        ///             },
        ///             new Linode.Inputs.GetChildAccountsFilterInputArgs
        ///             {
        ///                 Name = "first_name",
        ///                 Values = new[]
        ///                 {
        ///                     "John",
        ///                 },
        ///             },
        ///             new Linode.Inputs.GetChildAccountsFilterInputArgs
        ///             {
        ///                 Name = "last_name",
        ///                 Values = new[]
        ///                 {
        ///                     "Smith",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["allAccounts"] = all.Apply(getChildAccountsResult =&gt; getChildAccountsResult.ChildAccounts).Select(__item =&gt; __item.Euuid).ToList(),
        ///         ["filteredAccounts"] = filtered.Apply(getChildAccountsResult =&gt; getChildAccountsResult.ChildAccounts).Select(__item =&gt; __item.Euuid).ToList(),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Filterable Fields
        /// 
        /// * `euuid`
        /// 
        /// * `email`
        /// 
        /// * `first_name`
        /// 
        /// * `last_name`
        /// 
        /// * `company`
        /// 
        /// * `address_1`
        /// 
        /// * `address_2`
        /// 
        /// * `phone`
        /// 
        /// * `city`
        /// 
        /// * `state`
        /// 
        /// * `country`
        /// 
        /// * `zip`
        /// 
        /// * `capabilities`
        /// 
        /// * `active_since`
        /// </summary>
        public static Task<GetChildAccountsResult> InvokeAsync(GetChildAccountsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetChildAccountsResult>("linode:index/getChildAccounts:getChildAccounts", args ?? new GetChildAccountsArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about Linode Child Accounts that match a set of filters.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-child-accounts).
        /// 
        /// **NOTE: Parent/Child related features may not be generally available.**
        /// 
        /// ## Example Usage
        /// 
        /// The following example shows how one might use this data source to access Child Accounts under the current Account.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Linode.GetChildAccounts.Invoke();
        /// 
        ///     var filtered = Linode.GetChildAccounts.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Linode.Inputs.GetChildAccountsFilterInputArgs
        ///             {
        ///                 Name = "email",
        ///                 Values = new[]
        ///                 {
        ///                     "example@linode.com",
        ///                 },
        ///             },
        ///             new Linode.Inputs.GetChildAccountsFilterInputArgs
        ///             {
        ///                 Name = "first_name",
        ///                 Values = new[]
        ///                 {
        ///                     "John",
        ///                 },
        ///             },
        ///             new Linode.Inputs.GetChildAccountsFilterInputArgs
        ///             {
        ///                 Name = "last_name",
        ///                 Values = new[]
        ///                 {
        ///                     "Smith",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["allAccounts"] = all.Apply(getChildAccountsResult =&gt; getChildAccountsResult.ChildAccounts).Select(__item =&gt; __item.Euuid).ToList(),
        ///         ["filteredAccounts"] = filtered.Apply(getChildAccountsResult =&gt; getChildAccountsResult.ChildAccounts).Select(__item =&gt; __item.Euuid).ToList(),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Filterable Fields
        /// 
        /// * `euuid`
        /// 
        /// * `email`
        /// 
        /// * `first_name`
        /// 
        /// * `last_name`
        /// 
        /// * `company`
        /// 
        /// * `address_1`
        /// 
        /// * `address_2`
        /// 
        /// * `phone`
        /// 
        /// * `city`
        /// 
        /// * `state`
        /// 
        /// * `country`
        /// 
        /// * `zip`
        /// 
        /// * `capabilities`
        /// 
        /// * `active_since`
        /// </summary>
        public static Output<GetChildAccountsResult> Invoke(GetChildAccountsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetChildAccountsResult>("linode:index/getChildAccounts:getChildAccounts", args ?? new GetChildAccountsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about Linode Child Accounts that match a set of filters.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-child-accounts).
        /// 
        /// **NOTE: Parent/Child related features may not be generally available.**
        /// 
        /// ## Example Usage
        /// 
        /// The following example shows how one might use this data source to access Child Accounts under the current Account.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Linode.GetChildAccounts.Invoke();
        /// 
        ///     var filtered = Linode.GetChildAccounts.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Linode.Inputs.GetChildAccountsFilterInputArgs
        ///             {
        ///                 Name = "email",
        ///                 Values = new[]
        ///                 {
        ///                     "example@linode.com",
        ///                 },
        ///             },
        ///             new Linode.Inputs.GetChildAccountsFilterInputArgs
        ///             {
        ///                 Name = "first_name",
        ///                 Values = new[]
        ///                 {
        ///                     "John",
        ///                 },
        ///             },
        ///             new Linode.Inputs.GetChildAccountsFilterInputArgs
        ///             {
        ///                 Name = "last_name",
        ///                 Values = new[]
        ///                 {
        ///                     "Smith",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["allAccounts"] = all.Apply(getChildAccountsResult =&gt; getChildAccountsResult.ChildAccounts).Select(__item =&gt; __item.Euuid).ToList(),
        ///         ["filteredAccounts"] = filtered.Apply(getChildAccountsResult =&gt; getChildAccountsResult.ChildAccounts).Select(__item =&gt; __item.Euuid).ToList(),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Filterable Fields
        /// 
        /// * `euuid`
        /// 
        /// * `email`
        /// 
        /// * `first_name`
        /// 
        /// * `last_name`
        /// 
        /// * `company`
        /// 
        /// * `address_1`
        /// 
        /// * `address_2`
        /// 
        /// * `phone`
        /// 
        /// * `city`
        /// 
        /// * `state`
        /// 
        /// * `country`
        /// 
        /// * `zip`
        /// 
        /// * `capabilities`
        /// 
        /// * `active_since`
        /// </summary>
        public static Output<GetChildAccountsResult> Invoke(GetChildAccountsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetChildAccountsResult>("linode:index/getChildAccounts:getChildAccounts", args ?? new GetChildAccountsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChildAccountsArgs : global::Pulumi.InvokeArgs
    {
        [Input("childAccounts")]
        private List<Inputs.GetChildAccountsChildAccountArgs>? _childAccounts;
        public List<Inputs.GetChildAccountsChildAccountArgs> ChildAccounts
        {
            get => _childAccounts ?? (_childAccounts = new List<Inputs.GetChildAccountsChildAccountArgs>());
            set => _childAccounts = value;
        }

        [Input("filters")]
        private List<Inputs.GetChildAccountsFilterArgs>? _filters;
        public List<Inputs.GetChildAccountsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetChildAccountsFilterArgs>());
            set => _filters = value;
        }

        public GetChildAccountsArgs()
        {
        }
        public static new GetChildAccountsArgs Empty => new GetChildAccountsArgs();
    }

    public sealed class GetChildAccountsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("childAccounts")]
        private InputList<Inputs.GetChildAccountsChildAccountInputArgs>? _childAccounts;
        public InputList<Inputs.GetChildAccountsChildAccountInputArgs> ChildAccounts
        {
            get => _childAccounts ?? (_childAccounts = new InputList<Inputs.GetChildAccountsChildAccountInputArgs>());
            set => _childAccounts = value;
        }

        [Input("filters")]
        private InputList<Inputs.GetChildAccountsFilterInputArgs>? _filters;
        public InputList<Inputs.GetChildAccountsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetChildAccountsFilterInputArgs>());
            set => _filters = value;
        }

        public GetChildAccountsInvokeArgs()
        {
        }
        public static new GetChildAccountsInvokeArgs Empty => new GetChildAccountsInvokeArgs();
    }


    [OutputType]
    public sealed class GetChildAccountsResult
    {
        public readonly ImmutableArray<Outputs.GetChildAccountsChildAccountResult> ChildAccounts;
        public readonly ImmutableArray<Outputs.GetChildAccountsFilterResult> Filters;
        public readonly string Id;

        [OutputConstructor]
        private GetChildAccountsResult(
            ImmutableArray<Outputs.GetChildAccountsChildAccountResult> childAccounts,

            ImmutableArray<Outputs.GetChildAccountsFilterResult> filters,

            string id)
        {
            ChildAccounts = childAccounts;
            Filters = filters;
            Id = id;
        }
    }
}
