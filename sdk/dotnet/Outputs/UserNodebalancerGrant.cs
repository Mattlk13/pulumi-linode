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
    public sealed class UserNodebalancerGrant
    {
        public readonly int Id;
        public readonly string Permissions;

        [OutputConstructor]
        private UserNodebalancerGrant(
            int id,

            string permissions)
        {
            Id = id;
            Permissions = permissions;
        }
    }
}