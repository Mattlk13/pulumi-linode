// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Linode
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("linode");
        /// <summary>
        /// An HTTP User-Agent Prefix to prepend in API requests.
        /// </summary>
        public static string? ApiVersion { get; set; } = __config.Get("apiVersion") ?? Utilities.GetEnv("LINODE_API_VERSION");

        /// <summary>
        /// The token that allows you access to your Linode account
        /// </summary>
        public static string? Token { get; set; } = __config.Get("token") ?? Utilities.GetEnv("LINODE_TOKEN", "LINODE_API_TOKEN");

        /// <summary>
        /// An HTTP User-Agent Prefix to prepend in API requests.
        /// </summary>
        public static string? UaPrefix { get; set; } = __config.Get("uaPrefix") ?? Utilities.GetEnv("LINODE_UA_PREFIX");

        /// <summary>
        /// The HTTP(S) API address of the Linode API to use.
        /// </summary>
        public static string? Url { get; set; } = __config.Get("url") ?? Utilities.GetEnv("LINODE_URL");

    }
}