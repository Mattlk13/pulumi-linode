// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Inputs
{

    public sealed class ReservedIpAssignmentVpcNat11GetArgs : global::Pulumi.ResourceArgs
    {
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        [Input("subnetId", required: true)]
        public Input<int> SubnetId { get; set; } = null!;

        [Input("vpcId", required: true)]
        public Input<int> VpcId { get; set; } = null!;

        public ReservedIpAssignmentVpcNat11GetArgs()
        {
        }
        public static new ReservedIpAssignmentVpcNat11GetArgs Empty => new ReservedIpAssignmentVpcNat11GetArgs();
    }
}
