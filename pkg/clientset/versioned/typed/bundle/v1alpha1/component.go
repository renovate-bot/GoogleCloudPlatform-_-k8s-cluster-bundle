// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/apis/bundle/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComponentsGetter has a method to return a ComponentInterface.
// A group's client should implement this interface.
type ComponentsGetter interface {
	Components(namespace string) ComponentInterface
}

// ComponentInterface has methods to work with Component resources.
type ComponentInterface interface {
	Create(*v1alpha1.Component) (*v1alpha1.Component, error)
	Update(*v1alpha1.Component) (*v1alpha1.Component, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Component, error)
	List(opts v1.ListOptions) (*v1alpha1.ComponentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Component, err error)
	ComponentExpansion
}

// components implements ComponentInterface
type components struct {
	client rest.Interface
	ns     string
}

// newComponents returns a Components
func newComponents(c *BundleV1alpha1Client, namespace string) *components {
	return &components{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the component, and returns the corresponding component object, and an error if there is any.
func (c *components) Get(name string, options v1.GetOptions) (result *v1alpha1.Component, err error) {
	result = &v1alpha1.Component{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("components").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Components that match those selectors.
func (c *components) List(opts v1.ListOptions) (result *v1alpha1.ComponentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComponentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("components").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested components.
func (c *components) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("components").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a component and creates it.  Returns the server's representation of the component, and an error, if there is any.
func (c *components) Create(component *v1alpha1.Component) (result *v1alpha1.Component, err error) {
	result = &v1alpha1.Component{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("components").
		Body(component).
		Do().
		Into(result)
	return
}

// Update takes the representation of a component and updates it. Returns the server's representation of the component, and an error, if there is any.
func (c *components) Update(component *v1alpha1.Component) (result *v1alpha1.Component, err error) {
	result = &v1alpha1.Component{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("components").
		Name(component.Name).
		Body(component).
		Do().
		Into(result)
	return
}

// Delete takes name of the component and deletes it. Returns an error if one occurs.
func (c *components) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("components").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *components) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("components").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched component.
func (c *components) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Component, err error) {
	result = &v1alpha1.Component{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("components").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
