// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Outputs
{

    [OutputType]
    public sealed class FirewallInbound
    {
        /// <summary>
        /// A list of IP addresses, CIDR blocks, or `0.0.0.0/0` (to whitelist all) this rule applies to.
        /// </summary>
        public readonly ImmutableArray<string> Addresses;
        /// <summary>
        /// A list of ports and/or port ranges (i.e. "443" or "80-90").
        /// </summary>
        public readonly ImmutableArray<string> Ports;
        /// <summary>
        /// The network protocol this rule controls.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private FirewallInbound(
            ImmutableArray<string> addresses,

            ImmutableArray<string> ports,

            string protocol)
        {
            Addresses = addresses;
            Ports = ports;
            Protocol = protocol;
        }
    }
}