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
    public sealed class GetInstancesInstanceConfigResult
    {
        public readonly string Comments;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceResult> Devices;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceConfigHelperResult> Helpers;
        public readonly string Kernel;
        public readonly string Label;
        public readonly int MemoryLimit;
        public readonly string RootDevice;
        public readonly string RunLevel;
        public readonly string VirtMode;

        [OutputConstructor]
        private GetInstancesInstanceConfigResult(
            string comments,

            ImmutableArray<Outputs.GetInstancesInstanceConfigDeviceResult> devices,

            ImmutableArray<Outputs.GetInstancesInstanceConfigHelperResult> helpers,

            string kernel,

            string label,

            int memoryLimit,

            string rootDevice,

            string runLevel,

            string virtMode)
        {
            Comments = comments;
            Devices = devices;
            Helpers = helpers;
            Kernel = kernel;
            Label = label;
            MemoryLimit = memoryLimit;
            RootDevice = rootDevice;
            RunLevel = runLevel;
            VirtMode = virtMode;
        }
    }
}