// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    public static class GetNodeBalancerConfig
    {
        /// <summary>
        /// Provides details about a Linode NodeBalancer Config.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_config = Output.Create(Linode.GetNodeBalancerConfig.InvokeAsync(new Linode.GetNodeBalancerConfigArgs
        ///         {
        ///             Id = 123,
        ///             NodebalancerId = 456,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNodeBalancerConfigResult> InvokeAsync(GetNodeBalancerConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNodeBalancerConfigResult>("linode:index/getNodeBalancerConfig:getNodeBalancerConfig", args ?? new GetNodeBalancerConfigArgs(), options.WithVersion());
    }


    public sealed class GetNodeBalancerConfigArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The config's ID.
        /// </summary>
        [Input("id", required: true)]
        public int Id { get; set; }

        /// <summary>
        /// The ID of the NodeBalancer that contains the config.
        /// </summary>
        [Input("nodebalancerId", required: true)]
        public int NodebalancerId { get; set; }

        public GetNodeBalancerConfigArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNodeBalancerConfigResult
    {
        /// <summary>
        /// What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source
        /// </summary>
        public readonly string Algorithm;
        /// <summary>
        /// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected.
        /// </summary>
        public readonly string Check;
        /// <summary>
        /// How many times to attempt a check before considering a backend to be down. (1-30)
        /// </summary>
        public readonly int CheckAttempts;
        public readonly string CheckBody;
        /// <summary>
        /// How often, in seconds, to check that backends are up and serving requests.
        /// </summary>
        public readonly int CheckInterval;
        /// <summary>
        /// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
        /// </summary>
        public readonly bool CheckPassive;
        /// <summary>
        /// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
        /// </summary>
        public readonly string CheckPath;
        /// <summary>
        /// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
        /// </summary>
        public readonly int CheckTimeout;
        /// <summary>
        /// What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.
        /// </summary>
        public readonly string CipherSuite;
        public readonly int Id;
        public readonly ImmutableArray<Outputs.GetNodeBalancerConfigNodeStatusResult> NodeStatuses;
        public readonly int NodebalancerId;
        /// <summary>
        /// The TCP port this Config is for.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key. (Defaults to "http")
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The version of ProxyProtocol to use for the underlying NodeBalancer. This requires protocol to be `tcp`. Valid values are `none`, `v1`, and `v2`. (Defaults to `none`)
        /// </summary>
        public readonly string ProxyProtocol;
        /// <summary>
        /// The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
        /// </summary>
        public readonly string SslCommonname;
        /// <summary>
        /// The read-only fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
        /// </summary>
        public readonly string SslFingerprint;
        /// <summary>
        /// Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'
        /// </summary>
        public readonly string Stickiness;

        [OutputConstructor]
        private GetNodeBalancerConfigResult(
            string algorithm,

            string check,

            int checkAttempts,

            string checkBody,

            int checkInterval,

            bool checkPassive,

            string checkPath,

            int checkTimeout,

            string cipherSuite,

            int id,

            ImmutableArray<Outputs.GetNodeBalancerConfigNodeStatusResult> nodeStatuses,

            int nodebalancerId,

            int port,

            string protocol,

            string proxyProtocol,

            string sslCommonname,

            string sslFingerprint,

            string stickiness)
        {
            Algorithm = algorithm;
            Check = check;
            CheckAttempts = checkAttempts;
            CheckBody = checkBody;
            CheckInterval = checkInterval;
            CheckPassive = checkPassive;
            CheckPath = checkPath;
            CheckTimeout = checkTimeout;
            CipherSuite = cipherSuite;
            Id = id;
            NodeStatuses = nodeStatuses;
            NodebalancerId = nodebalancerId;
            Port = port;
            Protocol = protocol;
            ProxyProtocol = proxyProtocol;
            SslCommonname = sslCommonname;
            SslFingerprint = sslFingerprint;
            Stickiness = stickiness;
        }
    }
}