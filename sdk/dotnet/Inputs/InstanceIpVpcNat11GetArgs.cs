// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Inputs
{

    public sealed class InstanceIpVpcNat11GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resulting IPv4 address.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        [Input("subnetId", required: true)]
        public Input<int> SubnetId { get; set; } = null!;

        [Input("vpcId", required: true)]
        public Input<int> VpcId { get; set; } = null!;

        public InstanceIpVpcNat11GetArgs()
        {
        }
        public static new InstanceIpVpcNat11GetArgs Empty => new InstanceIpVpcNat11GetArgs();
    }
}
