/*
Copyright 2018 k8s provisioner juju example Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/alejandroEsc/k8s-controller-example/pkg/apis/controller/v1alpha1"
	scheme "github.com/alejandroEsc/k8s-controller-example/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SampleResourcesGetter has a method to return a SampleResourceInterface.
// A group's client should implement this interface.
type SampleResourcesGetter interface {
	SampleResources(namespace string) SampleResourceInterface
}

// SampleResourceInterface has methods to work with SampleResource resources.
type SampleResourceInterface interface {
	Create(*v1alpha1.SampleResource) (*v1alpha1.SampleResource, error)
	Update(*v1alpha1.SampleResource) (*v1alpha1.SampleResource, error)
	UpdateStatus(*v1alpha1.SampleResource) (*v1alpha1.SampleResource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SampleResource, error)
	List(opts v1.ListOptions) (*v1alpha1.SampleResourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SampleResource, err error)
	SampleResourceExpansion
}

// sampleResources implements SampleResourceInterface
type sampleResources struct {
	client rest.Interface
	ns     string
}

// newSampleResources returns a SampleResources
func newSampleResources(c *ControllerV1alpha1Client, namespace string) *sampleResources {
	return &sampleResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sampleResource, and returns the corresponding sampleResource object, and an error if there is any.
func (c *sampleResources) Get(name string, options v1.GetOptions) (result *v1alpha1.SampleResource, err error) {
	result = &v1alpha1.SampleResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sampleresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SampleResources that match those selectors.
func (c *sampleResources) List(opts v1.ListOptions) (result *v1alpha1.SampleResourceList, err error) {
	result = &v1alpha1.SampleResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sampleresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sampleResources.
func (c *sampleResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sampleresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sampleResource and creates it.  Returns the server's representation of the sampleResource, and an error, if there is any.
func (c *sampleResources) Create(sampleResource *v1alpha1.SampleResource) (result *v1alpha1.SampleResource, err error) {
	result = &v1alpha1.SampleResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sampleresources").
		Body(sampleResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sampleResource and updates it. Returns the server's representation of the sampleResource, and an error, if there is any.
func (c *sampleResources) Update(sampleResource *v1alpha1.SampleResource) (result *v1alpha1.SampleResource, err error) {
	result = &v1alpha1.SampleResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sampleresources").
		Name(sampleResource.Name).
		Body(sampleResource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sampleResources) UpdateStatus(sampleResource *v1alpha1.SampleResource) (result *v1alpha1.SampleResource, err error) {
	result = &v1alpha1.SampleResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sampleresources").
		Name(sampleResource.Name).
		SubResource("status").
		Body(sampleResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the sampleResource and deletes it. Returns an error if one occurs.
func (c *sampleResources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sampleresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sampleResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sampleresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sampleResource.
func (c *sampleResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SampleResource, err error) {
	result = &v1alpha1.SampleResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sampleresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
