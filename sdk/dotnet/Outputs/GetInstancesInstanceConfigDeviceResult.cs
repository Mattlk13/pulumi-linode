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
    public sealed class GetInstancesInstanceConfigDeviceResult
    {
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdaResult> Sdas;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdbResult> Sdbs;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdcResult> Sdcs;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSddResult> Sdds;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdeResult> Sdes;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdfResult> Sdfs;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdgResult> Sdgs;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdhResult> Sdhs;

        [OutputConstructor]
        private GetInstancesInstanceConfigDeviceResult(
            ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdaResult> sdas,

            ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdbResult> sdbs,

            ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdcResult> sdcs,

            ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSddResult> sdds,

            ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdeResult> sdes,

            ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdfResult> sdfs,

            ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdgResult> sdgs,

            ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceSdhResult> sdhs)
        {
            Sdas = sdas;
            Sdbs = sdbs;
            Sdcs = sdcs;
            Sdds = sdds;
            Sdes = sdes;
            Sdfs = sdfs;
            Sdgs = sdgs;
            Sdhs = sdhs;
        }
    }
}