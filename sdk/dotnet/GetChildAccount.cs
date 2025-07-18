// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    public static class GetChildAccount
    {
        /// <summary>
        /// Provides information about a Linode Child Account.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-child-account).
        /// 
        /// Due to the sensitive nature of the data exposed by this data source, it should not be used in conjunction with the `LINODE_DEBUG` option.  See the [debugging notes](https://www.terraform.io/providers/linode/linode/latest/docs#debugging) for more details.
        /// 
        /// **NOTE: Parent/Child related features may not be generally available.**
        /// 
        /// ## Example Usage
        /// 
        /// The following example shows how one might use this data source to access child account details.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var account = Linode.GetChildAccount.Invoke(new()
        ///     {
        ///         Euuid = "FFFFFFFF-FFFF-FFFF-FFFFFFFFFFFFFFFF",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetChildAccountResult> InvokeAsync(GetChildAccountArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetChildAccountResult>("linode:index/getChildAccount:getChildAccount", args ?? new GetChildAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about a Linode Child Account.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-child-account).
        /// 
        /// Due to the sensitive nature of the data exposed by this data source, it should not be used in conjunction with the `LINODE_DEBUG` option.  See the [debugging notes](https://www.terraform.io/providers/linode/linode/latest/docs#debugging) for more details.
        /// 
        /// **NOTE: Parent/Child related features may not be generally available.**
        /// 
        /// ## Example Usage
        /// 
        /// The following example shows how one might use this data source to access child account details.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var account = Linode.GetChildAccount.Invoke(new()
        ///     {
        ///         Euuid = "FFFFFFFF-FFFF-FFFF-FFFFFFFFFFFFFFFF",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetChildAccountResult> Invoke(GetChildAccountInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetChildAccountResult>("linode:index/getChildAccount:getChildAccount", args ?? new GetChildAccountInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about a Linode Child Account.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-child-account).
        /// 
        /// Due to the sensitive nature of the data exposed by this data source, it should not be used in conjunction with the `LINODE_DEBUG` option.  See the [debugging notes](https://www.terraform.io/providers/linode/linode/latest/docs#debugging) for more details.
        /// 
        /// **NOTE: Parent/Child related features may not be generally available.**
        /// 
        /// ## Example Usage
        /// 
        /// The following example shows how one might use this data source to access child account details.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var account = Linode.GetChildAccount.Invoke(new()
        ///     {
        ///         Euuid = "FFFFFFFF-FFFF-FFFF-FFFFFFFFFFFFFFFF",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetChildAccountResult> Invoke(GetChildAccountInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetChildAccountResult>("linode:index/getChildAccount:getChildAccount", args ?? new GetChildAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChildAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique EUUID of this Child Account.
        /// </summary>
        [Input("euuid", required: true)]
        public string Euuid { get; set; } = null!;

        public GetChildAccountArgs()
        {
        }
        public static new GetChildAccountArgs Empty => new GetChildAccountArgs();
    }

    public sealed class GetChildAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique EUUID of this Child Account.
        /// </summary>
        [Input("euuid", required: true)]
        public Input<string> Euuid { get; set; } = null!;

        public GetChildAccountInvokeArgs()
        {
        }
        public static new GetChildAccountInvokeArgs Empty => new GetChildAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetChildAccountResult
    {
        /// <summary>
        /// When this account was first activated.
        /// </summary>
        public readonly string ActiveSince;
        /// <summary>
        /// First line of this Account's billing address.
        /// </summary>
        public readonly string Address1;
        /// <summary>
        /// Second line of this Account's billing address.
        /// </summary>
        public readonly string Address2;
        /// <summary>
        /// This Account's balance, in US dollars.
        /// </summary>
        public readonly double Balance;
        /// <summary>
        /// A set containing all the capabilities of this Account.
        /// </summary>
        public readonly ImmutableArray<string> Capabilities;
        /// <summary>
        /// The city for this Account's billing address.
        /// </summary>
        public readonly string City;
        /// <summary>
        /// The company name associated with this Account.
        /// </summary>
        public readonly string Company;
        /// <summary>
        /// The two-letter country code of this Account's billing address.
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// The email address for this Account, for account management communications, and may be used for other communications as configured.
        /// </summary>
        public readonly string Email;
        public readonly string Euuid;
        /// <summary>
        /// The first name of the person associated with this Account.
        /// </summary>
        public readonly string FirstName;
        public readonly string Id;
        /// <summary>
        /// The last name of the person associated with this Account.
        /// </summary>
        public readonly string LastName;
        /// <summary>
        /// The phone number associated with this Account.
        /// </summary>
        public readonly string Phone;
        /// <summary>
        /// If billing address is in the United States, this is the State portion of the Account's billing address. If the address is outside the US, this is the Province associated with the Account's billing address.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The zip code of this Account's billing address.
        /// </summary>
        public readonly string Zip;

        [OutputConstructor]
        private GetChildAccountResult(
            string activeSince,

            string address1,

            string address2,

            double balance,

            ImmutableArray<string> capabilities,

            string city,

            string company,

            string country,

            string email,

            string euuid,

            string firstName,

            string id,

            string lastName,

            string phone,

            string state,

            string zip)
        {
            ActiveSince = activeSince;
            Address1 = address1;
            Address2 = address2;
            Balance = balance;
            Capabilities = capabilities;
            City = city;
            Company = company;
            Country = country;
            Email = email;
            Euuid = euuid;
            FirstName = firstName;
            Id = id;
            LastName = lastName;
            Phone = phone;
            State = state;
            Zip = zip;
        }
    }
}
