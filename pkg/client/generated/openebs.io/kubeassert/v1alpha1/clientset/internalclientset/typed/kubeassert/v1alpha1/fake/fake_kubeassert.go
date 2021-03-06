/*
Copyright 2019 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/kubeassert/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeAsserts implements KubeAssertInterface
type FakeKubeAsserts struct {
	Fake *FakeOpenebsV1alpha1
	ns   string
}

var kubeassertsResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "kubeasserts"}

var kubeassertsKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "KubeAssert"}

// Get takes name of the kubeAssert, and returns the corresponding kubeAssert object, and an error if there is any.
func (c *FakeKubeAsserts) Get(name string, options v1.GetOptions) (result *v1alpha1.KubeAssert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubeassertsResource, c.ns, name), &v1alpha1.KubeAssert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeAssert), err
}

// List takes label and field selectors, and returns the list of KubeAsserts that match those selectors.
func (c *FakeKubeAsserts) List(opts v1.ListOptions) (result *v1alpha1.KubeAssertList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubeassertsResource, kubeassertsKind, c.ns, opts), &v1alpha1.KubeAssertList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubeAssertList{ListMeta: obj.(*v1alpha1.KubeAssertList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubeAssertList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeAsserts.
func (c *FakeKubeAsserts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubeassertsResource, c.ns, opts))

}

// Create takes the representation of a kubeAssert and creates it.  Returns the server's representation of the kubeAssert, and an error, if there is any.
func (c *FakeKubeAsserts) Create(kubeAssert *v1alpha1.KubeAssert) (result *v1alpha1.KubeAssert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubeassertsResource, c.ns, kubeAssert), &v1alpha1.KubeAssert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeAssert), err
}

// Update takes the representation of a kubeAssert and updates it. Returns the server's representation of the kubeAssert, and an error, if there is any.
func (c *FakeKubeAsserts) Update(kubeAssert *v1alpha1.KubeAssert) (result *v1alpha1.KubeAssert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubeassertsResource, c.ns, kubeAssert), &v1alpha1.KubeAssert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeAssert), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeAsserts) UpdateStatus(kubeAssert *v1alpha1.KubeAssert) (*v1alpha1.KubeAssert, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubeassertsResource, "status", c.ns, kubeAssert), &v1alpha1.KubeAssert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeAssert), err
}

// Delete takes name of the kubeAssert and deletes it. Returns an error if one occurs.
func (c *FakeKubeAsserts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubeassertsResource, c.ns, name), &v1alpha1.KubeAssert{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeAsserts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubeassertsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubeAssertList{})
	return err
}

// Patch applies the patch and returns the patched kubeAssert.
func (c *FakeKubeAsserts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KubeAssert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubeassertsResource, c.ns, name, data, subresources...), &v1alpha1.KubeAssert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeAssert), err
}
