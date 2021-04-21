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
    public sealed class ObjectStorageBucketLifecycleRule
    {
        /// <summary>
        /// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
        /// </summary>
        public readonly int? AbortIncompleteMultipartUploadDays;
        /// <summary>
        /// Specifies whether the lifecycle rule is active.
        /// </summary>
        public readonly bool Enabled;
        public readonly Outputs.ObjectStorageBucketLifecycleRuleExpiration? Expiration;
        /// <summary>
        /// The unique identifier for the rule.
        /// </summary>
        public readonly string? Id;
        public readonly Outputs.ObjectStorageBucketLifecycleRuleNoncurrentVersionExpiration? NoncurrentVersionExpiration;
        /// <summary>
        /// The object key prefix identifying one or more objects to which the rule applies.
        /// </summary>
        public readonly string? Prefix;

        [OutputConstructor]
        private ObjectStorageBucketLifecycleRule(
            int? abortIncompleteMultipartUploadDays,

            bool enabled,

            Outputs.ObjectStorageBucketLifecycleRuleExpiration? expiration,

            string? id,

            Outputs.ObjectStorageBucketLifecycleRuleNoncurrentVersionExpiration? noncurrentVersionExpiration,

            string? prefix)
        {
            AbortIncompleteMultipartUploadDays = abortIncompleteMultipartUploadDays;
            Enabled = enabled;
            Expiration = expiration;
            Id = id;
            NoncurrentVersionExpiration = noncurrentVersionExpiration;
            Prefix = prefix;
        }
    }
}