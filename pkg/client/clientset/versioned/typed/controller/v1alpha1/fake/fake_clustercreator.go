/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/alejandroEsc/k8s-provisioner-juju-example/pkg/apis/controller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterCreators implements ClusterCreatorInterface
type FakeClusterCreators struct {
	Fake *FakeControllerV1alpha1
	ns   string
}

var clustercreatorsResource = schema.GroupVersionResource{Group: "controller.alejandro.esc.com", Version: "v1alpha1", Resource: "clustercreators"}

var clustercreatorsKind = schema.GroupVersionKind{Group: "controller.alejandro.esc.com", Version: "v1alpha1", Kind: "ClusterCreator"}

// Get takes name of the clusterCreator, and returns the corresponding clusterCreator object, and an error if there is any.
func (c *FakeClusterCreators) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterCreator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustercreatorsResource, c.ns, name), &v1alpha1.ClusterCreator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCreator), err
}

// List takes label and field selectors, and returns the list of ClusterCreators that match those selectors.
func (c *FakeClusterCreators) List(opts v1.ListOptions) (result *v1alpha1.ClusterCreatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustercreatorsResource, clustercreatorsKind, c.ns, opts), &v1alpha1.ClusterCreatorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterCreatorList{}
	for _, item := range obj.(*v1alpha1.ClusterCreatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterCreators.
func (c *FakeClusterCreators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustercreatorsResource, c.ns, opts))

}

// Create takes the representation of a clusterCreator and creates it.  Returns the server's representation of the clusterCreator, and an error, if there is any.
func (c *FakeClusterCreators) Create(clusterCreator *v1alpha1.ClusterCreator) (result *v1alpha1.ClusterCreator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustercreatorsResource, c.ns, clusterCreator), &v1alpha1.ClusterCreator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCreator), err
}

// Update takes the representation of a clusterCreator and updates it. Returns the server's representation of the clusterCreator, and an error, if there is any.
func (c *FakeClusterCreators) Update(clusterCreator *v1alpha1.ClusterCreator) (result *v1alpha1.ClusterCreator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustercreatorsResource, c.ns, clusterCreator), &v1alpha1.ClusterCreator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCreator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterCreators) UpdateStatus(clusterCreator *v1alpha1.ClusterCreator) (*v1alpha1.ClusterCreator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clustercreatorsResource, "status", c.ns, clusterCreator), &v1alpha1.ClusterCreator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCreator), err
}

// Delete takes name of the clusterCreator and deletes it. Returns an error if one occurs.
func (c *FakeClusterCreators) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustercreatorsResource, c.ns, name), &v1alpha1.ClusterCreator{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterCreators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustercreatorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterCreatorList{})
	return err
}

// Patch applies the patch and returns the patched clusterCreator.
func (c *FakeClusterCreators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterCreator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustercreatorsResource, c.ns, name, data, subresources...), &v1alpha1.ClusterCreator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCreator), err
}