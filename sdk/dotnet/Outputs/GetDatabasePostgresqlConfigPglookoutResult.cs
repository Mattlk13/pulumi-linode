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
    public sealed class GetDatabasePostgresqlConfigPglookoutResult
    {
        public readonly Outputs.GetDatabasePostgresqlConfigPglookoutMaxFailoverReplicationTimeLagResult MaxFailoverReplicationTimeLag;

        [OutputConstructor]
        private GetDatabasePostgresqlConfigPglookoutResult(Outputs.GetDatabasePostgresqlConfigPglookoutMaxFailoverReplicationTimeLagResult maxFailoverReplicationTimeLag)
        {
            MaxFailoverReplicationTimeLag = maxFailoverReplicationTimeLag;
        }
    }
}
