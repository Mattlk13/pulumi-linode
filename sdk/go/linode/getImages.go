// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about Linode images that match a set of filters.
//
// ## Attributes
//
// Each Linode image will be stored in the `images` attribute and will export the following attributes:
//
// * `id` - The unique ID of this Image.  The ID of private images begin with `private/` followed by the numeric identifier of the private image, for example `private/12345`.
//
// * `label` - A short description of the Image.
//
// * `created` - When this Image was created.
//
// * `createdBy` - The name of the User who created this Image, or "linode" for official Images.
//
// * `deprecated` - Whether or not this Image is deprecated. Will only be true for deprecated public Images.
//
// * `description` - A detailed description of this Image.
//
// * `isPublic` - True if the Image is public.
//
// * `size` - The minimum size this Image needs to deploy. Size is in MB. example: 2500
//
// * `type` - How the Image was created. Manual Images can be created at any time. "Automatic" Images are created automatically from a deleted Linode.
//
// * `vendor` - The upstream distribution vendor. `None` for private Images.
//
// ## Filterable Fields
//
// * `deprecated`
//
// * `isPublic`
//
// * `label`
//
// * `size`
//
// * `vendor`
func GetImages(ctx *pulumi.Context, args *GetImagesArgs, opts ...pulumi.InvokeOption) (*GetImagesResult, error) {
	var rv GetImagesResult
	err := ctx.Invoke("linode:index/getImages:getImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImages.
type GetImagesArgs struct {
	Filters []GetImagesFilter `pulumi:"filters"`
}

// A collection of values returned by getImages.
type GetImagesResult struct {
	Filters []GetImagesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id     string           `pulumi:"id"`
	Images []GetImagesImage `pulumi:"images"`
}