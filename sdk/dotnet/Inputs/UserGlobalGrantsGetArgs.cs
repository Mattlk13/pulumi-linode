// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Inputs
{

    public sealed class UserGlobalGrantsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The level of access this User has to Account-level actions, like billing information. A restricted User will never be able to manage users.
        /// </summary>
        [Input("accountAccess")]
        public Input<string>? AccountAccess { get; set; }

        /// <summary>
        /// If true, this User may add Databases.
        /// </summary>
        [Input("addDatabases")]
        public Input<bool>? AddDatabases { get; set; }

        /// <summary>
        /// If true, this User may add Domains.
        /// </summary>
        [Input("addDomains")]
        public Input<bool>? AddDomains { get; set; }

        /// <summary>
        /// If true, this User may add Firewalls.
        /// </summary>
        [Input("addFirewalls")]
        public Input<bool>? AddFirewalls { get; set; }

        /// <summary>
        /// If true, this User may add Images.
        /// </summary>
        [Input("addImages")]
        public Input<bool>? AddImages { get; set; }

        /// <summary>
        /// If true, this User may create Linodes.
        /// </summary>
        [Input("addLinodes")]
        public Input<bool>? AddLinodes { get; set; }

        /// <summary>
        /// If true, this User may create Longview clients and view the current plan.
        /// </summary>
        [Input("addLongview")]
        public Input<bool>? AddLongview { get; set; }

        /// <summary>
        /// If true, this User may add NodeBalancers.
        /// </summary>
        [Input("addNodebalancers")]
        public Input<bool>? AddNodebalancers { get; set; }

        /// <summary>
        /// If true, this User may add Placement Groups.
        /// </summary>
        [Input("addPlacementGroups")]
        public Input<bool>? AddPlacementGroups { get; set; }

        /// <summary>
        /// If true, this User may add StackScripts.
        /// </summary>
        [Input("addStackscripts")]
        public Input<bool>? AddStackscripts { get; set; }

        /// <summary>
        /// If true, this User may add Volumes.
        /// </summary>
        [Input("addVolumes")]
        public Input<bool>? AddVolumes { get; set; }

        /// <summary>
        /// If true, this User may add Virtual Private Clouds (VPCs).
        /// </summary>
        [Input("addVpcs")]
        public Input<bool>? AddVpcs { get; set; }

        /// <summary>
        /// If true, this User may cancel the entire Account.
        /// </summary>
        [Input("cancelAccount")]
        public Input<bool>? CancelAccount { get; set; }

        /// <summary>
        /// If true, this User may manage the Account’s Longview subscription.
        /// </summary>
        [Input("longviewSubscription")]
        public Input<bool>? LongviewSubscription { get; set; }

        public UserGlobalGrantsGetArgs()
        {
        }
        public static new UserGlobalGrantsGetArgs Empty => new UserGlobalGrantsGetArgs();
    }
}
