// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	configv1 "github.com/openshift/api/config/v1"
	applyconfigurationsconfigv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ImageTagMirrorSetsGetter has a method to return a ImageTagMirrorSetInterface.
// A group's client should implement this interface.
type ImageTagMirrorSetsGetter interface {
	ImageTagMirrorSets() ImageTagMirrorSetInterface
}

// ImageTagMirrorSetInterface has methods to work with ImageTagMirrorSet resources.
type ImageTagMirrorSetInterface interface {
	Create(ctx context.Context, imageTagMirrorSet *configv1.ImageTagMirrorSet, opts metav1.CreateOptions) (*configv1.ImageTagMirrorSet, error)
	Update(ctx context.Context, imageTagMirrorSet *configv1.ImageTagMirrorSet, opts metav1.UpdateOptions) (*configv1.ImageTagMirrorSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, imageTagMirrorSet *configv1.ImageTagMirrorSet, opts metav1.UpdateOptions) (*configv1.ImageTagMirrorSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*configv1.ImageTagMirrorSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*configv1.ImageTagMirrorSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *configv1.ImageTagMirrorSet, err error)
	Apply(ctx context.Context, imageTagMirrorSet *applyconfigurationsconfigv1.ImageTagMirrorSetApplyConfiguration, opts metav1.ApplyOptions) (result *configv1.ImageTagMirrorSet, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, imageTagMirrorSet *applyconfigurationsconfigv1.ImageTagMirrorSetApplyConfiguration, opts metav1.ApplyOptions) (result *configv1.ImageTagMirrorSet, err error)
	ImageTagMirrorSetExpansion
}

// imageTagMirrorSets implements ImageTagMirrorSetInterface
type imageTagMirrorSets struct {
	*gentype.ClientWithListAndApply[*configv1.ImageTagMirrorSet, *configv1.ImageTagMirrorSetList, *applyconfigurationsconfigv1.ImageTagMirrorSetApplyConfiguration]
}

// newImageTagMirrorSets returns a ImageTagMirrorSets
func newImageTagMirrorSets(c *ConfigV1Client) *imageTagMirrorSets {
	return &imageTagMirrorSets{
		gentype.NewClientWithListAndApply[*configv1.ImageTagMirrorSet, *configv1.ImageTagMirrorSetList, *applyconfigurationsconfigv1.ImageTagMirrorSetApplyConfiguration](
			"imagetagmirrorsets",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *configv1.ImageTagMirrorSet { return &configv1.ImageTagMirrorSet{} },
			func() *configv1.ImageTagMirrorSetList { return &configv1.ImageTagMirrorSetList{} },
		),
	}
}
