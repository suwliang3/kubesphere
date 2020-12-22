// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRequestAuthentications implements RequestAuthenticationInterface
type FakeRequestAuthentications struct {
	Fake *FakeSecurityV1beta1
	ns   string
}

var requestauthenticationsResource = schema.GroupVersionResource{Group: "security.istio.io", Version: "v1beta1", Resource: "requestauthentications"}

var requestauthenticationsKind = schema.GroupVersionKind{Group: "security.istio.io", Version: "v1beta1", Kind: "RequestAuthentication"}

// Get takes name of the requestAuthentication, and returns the corresponding requestAuthentication object, and an error if there is any.
func (c *FakeRequestAuthentications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.RequestAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(requestauthenticationsResource, c.ns, name), &v1beta1.RequestAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RequestAuthentication), err
}

// List takes label and field selectors, and returns the list of RequestAuthentications that match those selectors.
func (c *FakeRequestAuthentications) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.RequestAuthenticationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(requestauthenticationsResource, requestauthenticationsKind, c.ns, opts), &v1beta1.RequestAuthenticationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.RequestAuthenticationList{ListMeta: obj.(*v1beta1.RequestAuthenticationList).ListMeta}
	for _, item := range obj.(*v1beta1.RequestAuthenticationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested requestAuthentications.
func (c *FakeRequestAuthentications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(requestauthenticationsResource, c.ns, opts))

}

// Create takes the representation of a requestAuthentication and creates it.  Returns the server's representation of the requestAuthentication, and an error, if there is any.
func (c *FakeRequestAuthentications) Create(ctx context.Context, requestAuthentication *v1beta1.RequestAuthentication, opts v1.CreateOptions) (result *v1beta1.RequestAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(requestauthenticationsResource, c.ns, requestAuthentication), &v1beta1.RequestAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RequestAuthentication), err
}

// Update takes the representation of a requestAuthentication and updates it. Returns the server's representation of the requestAuthentication, and an error, if there is any.
func (c *FakeRequestAuthentications) Update(ctx context.Context, requestAuthentication *v1beta1.RequestAuthentication, opts v1.UpdateOptions) (result *v1beta1.RequestAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(requestauthenticationsResource, c.ns, requestAuthentication), &v1beta1.RequestAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RequestAuthentication), err
}

// Delete takes name of the requestAuthentication and deletes it. Returns an error if one occurs.
func (c *FakeRequestAuthentications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(requestauthenticationsResource, c.ns, name), &v1beta1.RequestAuthentication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRequestAuthentications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(requestauthenticationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.RequestAuthenticationList{})
	return err
}

// Patch applies the patch and returns the patched requestAuthentication.
func (c *FakeRequestAuthentications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.RequestAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(requestauthenticationsResource, c.ns, name, pt, data, subresources...), &v1beta1.RequestAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RequestAuthentication), err
}
