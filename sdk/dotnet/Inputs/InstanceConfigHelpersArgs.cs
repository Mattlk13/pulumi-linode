// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Inputs
{

    public sealed class InstanceConfigHelpersArgs : Pulumi.ResourceArgs
    {
        [Input("devtmpfsAutomount")]
        public Input<bool>? DevtmpfsAutomount { get; set; }

        /// <summary>
        /// Controls the behavior of the Linode Config's Distribution Helper setting.
        /// </summary>
        [Input("distro")]
        public Input<bool>? Distro { get; set; }

        /// <summary>
        /// Creates a modules dependency file for the Kernel you run.
        /// </summary>
        [Input("modulesDep")]
        public Input<bool>? ModulesDep { get; set; }

        /// <summary>
        /// Controls the behavior of the Linode Config's Network Helper setting, used to automatically configure additional IP addresses assigned to this instance.
        /// </summary>
        [Input("network")]
        public Input<bool>? Network { get; set; }

        /// <summary>
        /// Disables updatedb cron job to avoid disk thrashing.
        /// </summary>
        [Input("updatedbDisabled")]
        public Input<bool>? UpdatedbDisabled { get; set; }

        public InstanceConfigHelpersArgs()
        {
        }
    }
}