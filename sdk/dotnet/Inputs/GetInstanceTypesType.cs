// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Inputs
{

    public sealed class GetInstanceTypesTypeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The number of VPUs this Linode Type offers.
        /// </summary>
        [Input("acceleratedDevices", required: true)]
        public int AcceleratedDevices { get; set; }

        [Input("addons", required: true)]
        private List<Inputs.GetInstanceTypesTypeAddonArgs>? _addons;

        /// <summary>
        /// Information about the optional Backup service offered for Linodes.
        /// </summary>
        public List<Inputs.GetInstanceTypesTypeAddonArgs> Addons
        {
            get => _addons ?? (_addons = new List<Inputs.GetInstanceTypesTypeAddonArgs>());
            set => _addons = value;
        }

        /// <summary>
        /// The class of the Linode Type. See all classes [here](https://techdocs.akamai.com/linode-api/reference/get-linode-types).
        /// </summary>
        [Input("class", required: true)]
        public string Class { get; set; } = null!;

        /// <summary>
        /// The Disk size, in MB, of the Linode Type.
        /// </summary>
        [Input("disk", required: true)]
        public int Disk { get; set; }

        /// <summary>
        /// The ID representing the Linode Type.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The Linode Type's label is for display purposes only.
        /// </summary>
        [Input("label", required: true)]
        public string Label { get; set; } = null!;

        /// <summary>
        /// The amount of RAM included in this Linode Type.
        /// </summary>
        [Input("memory", required: true)]
        public int Memory { get; set; }

        /// <summary>
        /// The Mbits outbound bandwidth allocation.
        /// </summary>
        [Input("networkOut", required: true)]
        public int NetworkOut { get; set; }

        [Input("prices", required: true)]
        private List<Inputs.GetInstanceTypesTypePriceArgs>? _prices;

        /// <summary>
        /// Cost in US dollars, broken down into hourly and monthly charges.
        /// </summary>
        public List<Inputs.GetInstanceTypesTypePriceArgs> Prices
        {
            get => _prices ?? (_prices = new List<Inputs.GetInstanceTypesTypePriceArgs>());
            set => _prices = value;
        }

        [Input("regionPrices", required: true)]
        private List<Inputs.GetInstanceTypesTypeRegionPriceArgs>? _regionPrices;

        /// <summary>
        /// A list of region-specific prices for this plan.
        /// </summary>
        public List<Inputs.GetInstanceTypesTypeRegionPriceArgs> RegionPrices
        {
            get => _regionPrices ?? (_regionPrices = new List<Inputs.GetInstanceTypesTypeRegionPriceArgs>());
            set => _regionPrices = value;
        }

        /// <summary>
        /// The monthly outbound transfer amount, in MB.
        /// </summary>
        [Input("transfer", required: true)]
        public int Transfer { get; set; }

        /// <summary>
        /// The number of VCPU cores this Linode Type offers.
        /// </summary>
        [Input("vcpus", required: true)]
        public int Vcpus { get; set; }

        public GetInstanceTypesTypeArgs()
        {
        }
        public static new GetInstanceTypesTypeArgs Empty => new GetInstanceTypesTypeArgs();
    }
}
