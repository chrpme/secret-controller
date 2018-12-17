// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/twendt/secret-controller/pkg/apis/secretcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKeyvaultSecrets implements KeyvaultSecretInterface
type FakeKeyvaultSecrets struct {
	Fake *FakeSecretcontrollerV1alpha1
	ns   string
}

var keyvaultsecretsResource = schema.GroupVersionResource{Group: "secretcontroller.twendt.de", Version: "v1alpha1", Resource: "keyvaultsecrets"}

var keyvaultsecretsKind = schema.GroupVersionKind{Group: "secretcontroller.twendt.de", Version: "v1alpha1", Kind: "KeyvaultSecret"}

// Get takes name of the keyvaultSecret, and returns the corresponding keyvaultSecret object, and an error if there is any.
func (c *FakeKeyvaultSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.KeyvaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keyvaultsecretsResource, c.ns, name), &v1alpha1.KeyvaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyvaultSecret), err
}

// List takes label and field selectors, and returns the list of KeyvaultSecrets that match those selectors.
func (c *FakeKeyvaultSecrets) List(opts v1.ListOptions) (result *v1alpha1.KeyvaultSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(keyvaultsecretsResource, keyvaultsecretsKind, c.ns, opts), &v1alpha1.KeyvaultSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KeyvaultSecretList{ListMeta: obj.(*v1alpha1.KeyvaultSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.KeyvaultSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keyvaultSecrets.
func (c *FakeKeyvaultSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keyvaultsecretsResource, c.ns, opts))

}

// Create takes the representation of a keyvaultSecret and creates it.  Returns the server's representation of the keyvaultSecret, and an error, if there is any.
func (c *FakeKeyvaultSecrets) Create(keyvaultSecret *v1alpha1.KeyvaultSecret) (result *v1alpha1.KeyvaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keyvaultsecretsResource, c.ns, keyvaultSecret), &v1alpha1.KeyvaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyvaultSecret), err
}

// Update takes the representation of a keyvaultSecret and updates it. Returns the server's representation of the keyvaultSecret, and an error, if there is any.
func (c *FakeKeyvaultSecrets) Update(keyvaultSecret *v1alpha1.KeyvaultSecret) (result *v1alpha1.KeyvaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keyvaultsecretsResource, c.ns, keyvaultSecret), &v1alpha1.KeyvaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyvaultSecret), err
}

// Delete takes name of the keyvaultSecret and deletes it. Returns an error if one occurs.
func (c *FakeKeyvaultSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(keyvaultsecretsResource, c.ns, name), &v1alpha1.KeyvaultSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeyvaultSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(keyvaultsecretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KeyvaultSecretList{})
	return err
}

// Patch applies the patch and returns the patched keyvaultSecret.
func (c *FakeKeyvaultSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KeyvaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keyvaultsecretsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KeyvaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyvaultSecret), err
}
