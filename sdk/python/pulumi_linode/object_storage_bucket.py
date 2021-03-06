# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ObjectStorageBucket']


class ObjectStorageBucket(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert: Optional[pulumi.Input[pulumi.InputType['ObjectStorageBucketCertArgs']]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Linode Object Storage Bucket resource. This can be used to create, modify, and delete Linodes Object Storage Buckets.

        ## Example Usage

        The following example shows how one might use this resource to create an Object Storage Bucket.

        ```python
        import pulumi
        import pulumi_linode as linode

        primary = linode.get_object_storage_cluster(id="us-east-1")
        foobar = linode.ObjectStorageBucket("foobar",
            cluster=primary.id,
            label="%s")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: The cluster of the Linode Object Storage Bucket.
        :param pulumi.Input[str] label: The label of the Linode Object Storage Bucket.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['cert'] = cert
            if cluster is None:
                raise TypeError("Missing required property 'cluster'")
            __props__['cluster'] = cluster
            if label is None:
                raise TypeError("Missing required property 'label'")
            __props__['label'] = label
        super(ObjectStorageBucket, __self__).__init__(
            'linode:index/objectStorageBucket:ObjectStorageBucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cert: Optional[pulumi.Input[pulumi.InputType['ObjectStorageBucketCertArgs']]] = None,
            cluster: Optional[pulumi.Input[str]] = None,
            label: Optional[pulumi.Input[str]] = None) -> 'ObjectStorageBucket':
        """
        Get an existing ObjectStorageBucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: The cluster of the Linode Object Storage Bucket.
        :param pulumi.Input[str] label: The label of the Linode Object Storage Bucket.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cert"] = cert
        __props__["cluster"] = cluster
        __props__["label"] = label
        return ObjectStorageBucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cert(self) -> pulumi.Output[Optional['outputs.ObjectStorageBucketCert']]:
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Output[str]:
        """
        The cluster of the Linode Object Storage Bucket.
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[str]:
        """
        The label of the Linode Object Storage Bucket.
        """
        return pulumi.get(self, "label")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

