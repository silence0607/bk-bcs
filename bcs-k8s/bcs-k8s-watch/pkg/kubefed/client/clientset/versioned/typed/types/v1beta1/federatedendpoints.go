/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"
	scheme "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedEndpointsesGetter has a method to return a FederatedEndpointsInterface.
// A group's client should implement this interface.
type FederatedEndpointsesGetter interface {
	FederatedEndpointses(namespace string) FederatedEndpointsInterface
}

// FederatedEndpointsInterface has methods to work with FederatedEndpoints resources.
type FederatedEndpointsInterface interface {
	Create(*v1beta1.FederatedEndpoints) (*v1beta1.FederatedEndpoints, error)
	Update(*v1beta1.FederatedEndpoints) (*v1beta1.FederatedEndpoints, error)
	UpdateStatus(*v1beta1.FederatedEndpoints) (*v1beta1.FederatedEndpoints, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedEndpoints, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedEndpointsList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedEndpoints, err error)
	FederatedEndpointsExpansion
}

// federatedEndpointses implements FederatedEndpointsInterface
type federatedEndpointses struct {
	client rest.Interface
	ns     string
}

// newFederatedEndpointses returns a FederatedEndpointses
func newFederatedEndpointses(c *TypesV1beta1Client, namespace string) *federatedEndpointses {
	return &federatedEndpointses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedEndpoints, and returns the corresponding federatedEndpoints object, and an error if there is any.
func (c *federatedEndpointses) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedEndpoints, err error) {
	result = &v1beta1.FederatedEndpoints{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedendpointses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedEndpointses that match those selectors.
func (c *federatedEndpointses) List(opts v1.ListOptions) (result *v1beta1.FederatedEndpointsList, err error) {
	result = &v1beta1.FederatedEndpointsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedendpointses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedEndpointses.
func (c *federatedEndpointses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedendpointses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedEndpoints and creates it.  Returns the server's representation of the federatedEndpoints, and an error, if there is any.
func (c *federatedEndpointses) Create(federatedEndpoints *v1beta1.FederatedEndpoints) (result *v1beta1.FederatedEndpoints, err error) {
	result = &v1beta1.FederatedEndpoints{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedendpointses").
		Body(federatedEndpoints).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedEndpoints and updates it. Returns the server's representation of the federatedEndpoints, and an error, if there is any.
func (c *federatedEndpointses) Update(federatedEndpoints *v1beta1.FederatedEndpoints) (result *v1beta1.FederatedEndpoints, err error) {
	result = &v1beta1.FederatedEndpoints{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedendpointses").
		Name(federatedEndpoints.Name).
		Body(federatedEndpoints).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedEndpointses) UpdateStatus(federatedEndpoints *v1beta1.FederatedEndpoints) (result *v1beta1.FederatedEndpoints, err error) {
	result = &v1beta1.FederatedEndpoints{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedendpointses").
		Name(federatedEndpoints.Name).
		SubResource("status").
		Body(federatedEndpoints).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedEndpoints and deletes it. Returns an error if one occurs.
func (c *federatedEndpointses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedendpointses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedEndpointses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedendpointses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedEndpoints.
func (c *federatedEndpointses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedEndpoints, err error) {
	result = &v1beta1.FederatedEndpoints{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedendpointses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
