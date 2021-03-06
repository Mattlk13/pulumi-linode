// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    public static class GetVolume
    {
        /// <summary>
        /// Provides information about a Linode Volume.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example shows how one might use this data source to access information about a Linode Volume.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Linode = Pulumi.Linode;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foo = Output.Create(Linode.GetVolume.InvokeAsync(new Linode.GetVolumeArgs
        ///         {
        ///             Id = 1234567,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes
        /// 
        /// The Linode Volume resource exports the following attributes:
        /// 
        /// - `id` - The unique ID of this Volume.
        /// 
        /// - `created` - When this Volume was created.
        /// 
        /// - `status` - The current status of the Volume. Can be one of "creating", "active", "resizing", or "contact_support".
        /// 
        /// - `label` - This Volume's label is for display purposes only.
        /// 
        /// - `tags` - An array of tags applied to this object.
        /// 
        /// - `size` - The Volume's size, in GiB.
        /// 
        /// - `region` - The datacenter in which this Volume is located.
        /// 
        /// - `updated` - When this Volume was last updated.
        /// 
        /// - `linode_id` - If a Volume is attached to a specific Linode, the ID of that Linode will be displayed here. If the Volume is unattached, this value will be null.
        /// 
        /// - `filesystem_path` - The full filesystem path for the Volume based on the Volume's label. Path is /dev/disk/by-id/scsi-0LinodeVolume + Volume label.
        /// </summary>
        public static Task<GetVolumeResult> InvokeAsync(GetVolumeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeResult>("linode:index/getVolume:getVolume", args ?? new GetVolumeArgs(), options.WithVersion());
    }


    public sealed class GetVolumeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique numeric ID of the Volume record to query.
        /// </summary>
        [Input("id", required: true)]
        public int Id { get; set; }

        public GetVolumeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVolumeResult
    {
        public readonly string Created;
        public readonly string FilesystemPath;
        public readonly int Id;
        public readonly string Label;
        public readonly int LinodeId;
        public readonly string Region;
        public readonly int Size;
        public readonly string Status;
        public readonly ImmutableArray<string> Tags;
        public readonly string Updated;

        [OutputConstructor]
        private GetVolumeResult(
            string created,

            string filesystemPath,

            int id,

            string label,

            int linodeId,

            string region,

            int size,

            string status,

            ImmutableArray<string> tags,

            string updated)
        {
            Created = created;
            FilesystemPath = filesystemPath;
            Id = id;
            Label = label;
            LinodeId = linodeId;
            Region = region;
            Size = size;
            Status = status;
            Tags = tags;
            Updated = updated;
        }
    }
}
