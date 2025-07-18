// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Outputs
{

    [OutputType]
    public sealed class InstanceIpVpcNat11
    {
        /// <summary>
        /// The resulting IPv4 address.
        /// </summary>
        public readonly string Address;
        public readonly int SubnetId;
        public readonly int VpcId;

        [OutputConstructor]
        private InstanceIpVpcNat11(
            string address,

            int subnetId,

            int vpcId)
        {
            Address = address;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
