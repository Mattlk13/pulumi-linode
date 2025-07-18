// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    public static class GetProfile
    {
        /// <summary>
        /// Provides information about a Linode profile.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-profile).
        /// 
        /// ## Example Usage
        /// 
        /// The following example shows how one might use this data source to access profile details.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var profile = Linode.GetProfile.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetProfileResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProfileResult>("linode:index/getProfile:getProfile", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Provides information about a Linode profile.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-profile).
        /// 
        /// ## Example Usage
        /// 
        /// The following example shows how one might use this data source to access profile details.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var profile = Linode.GetProfile.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProfileResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProfileResult>("linode:index/getProfile:getProfile", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Provides information about a Linode profile.
        /// For more information, see the [Linode APIv4 docs](https://techdocs.akamai.com/linode-api/reference/get-profile).
        /// 
        /// ## Example Usage
        /// 
        /// The following example shows how one might use this data source to access profile details.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var profile = Linode.GetProfile.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProfileResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProfileResult>("linode:index/getProfile:getProfile", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetProfileResult
    {
        /// <summary>
        /// The list of SSH Keys authorized to use Lish for this user. This value is ignored if lish_auth_method is 'disabled'.
        /// </summary>
        public readonly ImmutableArray<string> AuthorizedKeys;
        /// <summary>
        /// The profile email address. This address will be used for communication with Linode as necessary.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// If true, email notifications will be sent about account activity. If false, when false business-critical communications may still be sent through email.
        /// </summary>
        public readonly bool EmailNotifications;
        public readonly string Id;
        /// <summary>
        /// If true, logins for the user will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled.
        /// </summary>
        public readonly bool IpWhitelistEnabled;
        /// <summary>
        /// The methods of authentication allowed when connecting via Lish. 'keys_only' is the most secure with the intent to use Lish, and 'disabled' is recommended for users that will not use Lish at all.
        /// </summary>
        public readonly string LishAuthMethod;
        /// <summary>
        /// Credit Card information associated with this Account.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProfileReferralResult> Referrals;
        /// <summary>
        /// If true, the user has restrictions on what can be accessed on the Account.
        /// </summary>
        public readonly bool Restricted;
        /// <summary>
        /// The profile's preferred timezone. This is not used by the API, and is for the benefit of clients only. All times the API returns are in UTC.
        /// </summary>
        public readonly string Timezone;
        /// <summary>
        /// If true, logins from untrusted computers will require Two Factor Authentication.
        /// </summary>
        public readonly bool TwoFactorAuth;
        /// <summary>
        /// The username for logging in to Linode services.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetProfileResult(
            ImmutableArray<string> authorizedKeys,

            string email,

            bool emailNotifications,

            string id,

            bool ipWhitelistEnabled,

            string lishAuthMethod,

            ImmutableArray<Outputs.GetProfileReferralResult> referrals,

            bool restricted,

            string timezone,

            bool twoFactorAuth,

            string username)
        {
            AuthorizedKeys = authorizedKeys;
            Email = email;
            EmailNotifications = emailNotifications;
            Id = id;
            IpWhitelistEnabled = ipWhitelistEnabled;
            LishAuthMethod = lishAuthMethod;
            Referrals = referrals;
            Restricted = restricted;
            Timezone = timezone;
            TwoFactorAuth = twoFactorAuth;
            Username = username;
        }
    }
}
