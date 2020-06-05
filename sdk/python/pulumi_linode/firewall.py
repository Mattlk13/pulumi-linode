# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Firewall(pulumi.CustomResource):
    devices: pulumi.Output[list]
    """
    The devices associated with this firewall.

      * `entityId` (`float`) - The ID of the underlying entity this device references (i.e. the Linode's ID).
      * `id` (`float`) - The ID of the Firewall Device.
      * `label` (`str`) - This Firewall's unique label.
      * `type` (`str`) - The type of Firewall Device.
      * `url` (`str`)
    """
    disabled: pulumi.Output[bool]
    """
    If `true`, the Firewall's rules are not enforced (defaults to `false`).
    """
    inbounds: pulumi.Output[list]
    """
    A firewall rule that specifies what inbound network traffic is allowed.

      * `addresses` (`list`) - A list of IP addresses, CIDR blocks, or `0.0.0.0/0` (to whitelist all) this rule applies to.
      * `ports` (`list`) - A list of ports and/or port ranges (i.e. "443" or "80-90").
      * `protocol` (`str`) - The network protocol this rule controls.
    """
    label: pulumi.Output[str]
    """
    This Firewall's unique label.
    """
    linodes: pulumi.Output[list]
    """
    A list of IDs of Linodes this Firewall should govern it's network traffic for.
    """
    outbounds: pulumi.Output[list]
    """
    A firewall rule that specifies what outbound network traffic is allowed.

      * `addresses` (`list`) - A list of IP addresses, CIDR blocks, or `0.0.0.0/0` (to whitelist all) this rule applies to.
      * `ports` (`list`) - A list of ports and/or port ranges (i.e. "443" or "80-90").
      * `protocol` (`str`) - The network protocol this rule controls.
    """
    status: pulumi.Output[str]
    """
    The status of the Firewall.
    """
    tags: pulumi.Output[list]
    """
    A list of tags applied to the Kubernetes cluster. Tags are for organizational purposes only.
    """
    def __init__(__self__, resource_name, opts=None, disabled=None, inbounds=None, label=None, linodes=None, outbounds=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        > **NOTICE:** The Firewall feature is currently available through early access. 

        Manages a Linode Firewall.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_linode as linode

        my_instance = linode.Instance("myInstance",
            label="my_instance",
            image="linode/ubuntu18.04",
            region="us-east",
            type="g6-standard-1",
            root_pass="bogusPassword$",
            swap_size=256)
        my_firewall = linode.Firewall("myFirewall",
            label="my_firewall",
            tags=["test"],
            inbound=[{
                "protocol": "TCP",
                "ports": ["80"],
                "addresses": ["0.0.0.0/0"],
            }],
            outbound=[{
                "protocol": "TCP",
                "ports": ["80"],
                "addresses": ["0.0.0.0/0"],
            }],
            linodes=[my_instance.id])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: If `true`, the Firewall's rules are not enforced (defaults to `false`).
        :param pulumi.Input[list] inbounds: A firewall rule that specifies what inbound network traffic is allowed.
        :param pulumi.Input[str] label: This Firewall's unique label.
        :param pulumi.Input[list] linodes: A list of IDs of Linodes this Firewall should govern it's network traffic for.
        :param pulumi.Input[list] outbounds: A firewall rule that specifies what outbound network traffic is allowed.
        :param pulumi.Input[list] tags: A list of tags applied to the Kubernetes cluster. Tags are for organizational purposes only.

        The **inbounds** object supports the following:

          * `addresses` (`pulumi.Input[list]`) - A list of IP addresses, CIDR blocks, or `0.0.0.0/0` (to whitelist all) this rule applies to.
          * `ports` (`pulumi.Input[list]`) - A list of ports and/or port ranges (i.e. "443" or "80-90").
          * `protocol` (`pulumi.Input[str]`) - The network protocol this rule controls.

        The **outbounds** object supports the following:

          * `addresses` (`pulumi.Input[list]`) - A list of IP addresses, CIDR blocks, or `0.0.0.0/0` (to whitelist all) this rule applies to.
          * `ports` (`pulumi.Input[list]`) - A list of ports and/or port ranges (i.e. "443" or "80-90").
          * `protocol` (`pulumi.Input[str]`) - The network protocol this rule controls.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['disabled'] = disabled
            __props__['inbounds'] = inbounds
            __props__['label'] = label
            if linodes is None:
                raise TypeError("Missing required property 'linodes'")
            __props__['linodes'] = linodes
            __props__['outbounds'] = outbounds
            __props__['tags'] = tags
            __props__['devices'] = None
            __props__['status'] = None
        super(Firewall, __self__).__init__(
            'linode:index/firewall:Firewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, devices=None, disabled=None, inbounds=None, label=None, linodes=None, outbounds=None, status=None, tags=None):
        """
        Get an existing Firewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] devices: The devices associated with this firewall.
        :param pulumi.Input[bool] disabled: If `true`, the Firewall's rules are not enforced (defaults to `false`).
        :param pulumi.Input[list] inbounds: A firewall rule that specifies what inbound network traffic is allowed.
        :param pulumi.Input[str] label: This Firewall's unique label.
        :param pulumi.Input[list] linodes: A list of IDs of Linodes this Firewall should govern it's network traffic for.
        :param pulumi.Input[list] outbounds: A firewall rule that specifies what outbound network traffic is allowed.
        :param pulumi.Input[str] status: The status of the Firewall.
        :param pulumi.Input[list] tags: A list of tags applied to the Kubernetes cluster. Tags are for organizational purposes only.

        The **devices** object supports the following:

          * `entityId` (`pulumi.Input[float]`) - The ID of the underlying entity this device references (i.e. the Linode's ID).
          * `id` (`pulumi.Input[float]`) - The ID of the Firewall Device.
          * `label` (`pulumi.Input[str]`) - This Firewall's unique label.
          * `type` (`pulumi.Input[str]`) - The type of Firewall Device.
          * `url` (`pulumi.Input[str]`)

        The **inbounds** object supports the following:

          * `addresses` (`pulumi.Input[list]`) - A list of IP addresses, CIDR blocks, or `0.0.0.0/0` (to whitelist all) this rule applies to.
          * `ports` (`pulumi.Input[list]`) - A list of ports and/or port ranges (i.e. "443" or "80-90").
          * `protocol` (`pulumi.Input[str]`) - The network protocol this rule controls.

        The **outbounds** object supports the following:

          * `addresses` (`pulumi.Input[list]`) - A list of IP addresses, CIDR blocks, or `0.0.0.0/0` (to whitelist all) this rule applies to.
          * `ports` (`pulumi.Input[list]`) - A list of ports and/or port ranges (i.e. "443" or "80-90").
          * `protocol` (`pulumi.Input[str]`) - The network protocol this rule controls.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["devices"] = devices
        __props__["disabled"] = disabled
        __props__["inbounds"] = inbounds
        __props__["label"] = label
        __props__["linodes"] = linodes
        __props__["outbounds"] = outbounds
        __props__["status"] = status
        __props__["tags"] = tags
        return Firewall(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
