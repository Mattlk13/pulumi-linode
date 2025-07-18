// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Inputs
{

    public sealed class GetFirewallsFirewallInboundInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Controls whether traffic is accepted or dropped by this rule (ACCEPT, DROP).
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("ipv4s", required: true)]
        private InputList<string>? _ipv4s;

        /// <summary>
        /// A list of IPv4 addresses or networks in IP/mask format.
        /// </summary>
        public InputList<string> Ipv4s
        {
            get => _ipv4s ?? (_ipv4s = new InputList<string>());
            set => _ipv4s = value;
        }

        [Input("ipv6s", required: true)]
        private InputList<string>? _ipv6s;

        /// <summary>
        /// A list of IPv6 addresses or networks in IP/mask format.
        /// </summary>
        public InputList<string> Ipv6s
        {
            get => _ipv6s ?? (_ipv6s = new InputList<string>());
            set => _ipv6s = value;
        }

        /// <summary>
        /// The label for the Firewall. For display purposes only. If no label is provided, a default will be assigned.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").
        /// </summary>
        [Input("ports", required: true)]
        public Input<string> Ports { get; set; } = null!;

        /// <summary>
        /// The network protocol this rule controls. (TCP, UDP, ICMP)
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public GetFirewallsFirewallInboundInputArgs()
        {
        }
        public static new GetFirewallsFirewallInboundInputArgs Empty => new GetFirewallsFirewallInboundInputArgs();
    }
}
