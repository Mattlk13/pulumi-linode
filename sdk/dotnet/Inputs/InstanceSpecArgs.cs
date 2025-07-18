// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Inputs
{

    public sealed class InstanceSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of VPUs this Linode has access to.
        /// </summary>
        [Input("acceleratedDevices")]
        public Input<int>? AcceleratedDevices { get; set; }

        /// <summary>
        /// The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an image through POST /linode/instances.
        /// </summary>
        [Input("disk")]
        public Input<int>? Disk { get; set; }

        /// <summary>
        /// The number of GPUs this Linode has access to.
        /// </summary>
        [Input("gpus")]
        public Input<int>? Gpus { get; set; }

        /// <summary>
        /// The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile.
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        /// <summary>
        /// The amount of network transfer this Linode is allotted each month.
        /// </summary>
        [Input("transfer")]
        public Input<int>? Transfer { get; set; }

        /// <summary>
        /// The number of vcpus this Linode has access to. Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile.
        /// </summary>
        [Input("vcpus")]
        public Input<int>? Vcpus { get; set; }

        public InstanceSpecArgs()
        {
        }
        public static new InstanceSpecArgs Empty => new InstanceSpecArgs();
    }
}
