// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/twendt/secret-controller/pkg/apis/secretcontroller/v1alpha1"
	scheme "github.com/twendt/secret-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KeyvaultSecretsGetter has a method to return a KeyvaultSecretInterface.
// A group's client should implement this interface.
type KeyvaultSecretsGetter interface {
	KeyvaultSecrets(namespace string) KeyvaultSecretInterface
}

// KeyvaultSecretInterface has methods to work with KeyvaultSecret resources.
type KeyvaultSecretInterface interface {
	Create(*v1alpha1.KeyvaultSecret) (*v1alpha1.KeyvaultSecret, error)
	Update(*v1alpha1.KeyvaultSecret) (*v1alpha1.KeyvaultSecret, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KeyvaultSecret, error)
	List(opts v1.ListOptions) (*v1alpha1.KeyvaultSecretList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KeyvaultSecret, err error)
	KeyvaultSecretExpansion
}

// keyvaultSecrets implements KeyvaultSecretInterface
type keyvaultSecrets struct {
	client rest.Interface
	ns     string
}

// newKeyvaultSecrets returns a KeyvaultSecrets
func newKeyvaultSecrets(c *SecretcontrollerV1alpha1Client, namespace string) *keyvaultSecrets {
	return &keyvaultSecrets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the keyvaultSecret, and returns the corresponding keyvaultSecret object, and an error if there is any.
func (c *keyvaultSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.KeyvaultSecret, err error) {
	result = &v1alpha1.KeyvaultSecret{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keyvaultsecrets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KeyvaultSecrets that match those selectors.
func (c *keyvaultSecrets) List(opts v1.ListOptions) (result *v1alpha1.KeyvaultSecretList, err error) {
	result = &v1alpha1.KeyvaultSecretList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keyvaultsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested keyvaultSecrets.
func (c *keyvaultSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("keyvaultsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a keyvaultSecret and creates it.  Returns the server's representation of the keyvaultSecret, and an error, if there is any.
func (c *keyvaultSecrets) Create(keyvaultSecret *v1alpha1.KeyvaultSecret) (result *v1alpha1.KeyvaultSecret, err error) {
	result = &v1alpha1.KeyvaultSecret{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("keyvaultsecrets").
		Body(keyvaultSecret).
		Do().
		Into(result)
	return
}

// Update takes the representation of a keyvaultSecret and updates it. Returns the server's representation of the keyvaultSecret, and an error, if there is any.
func (c *keyvaultSecrets) Update(keyvaultSecret *v1alpha1.KeyvaultSecret) (result *v1alpha1.KeyvaultSecret, err error) {
	result = &v1alpha1.KeyvaultSecret{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keyvaultsecrets").
		Name(keyvaultSecret.Name).
		Body(keyvaultSecret).
		Do().
		Into(result)
	return
}

// Delete takes name of the keyvaultSecret and deletes it. Returns an error if one occurs.
func (c *keyvaultSecrets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keyvaultsecrets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *keyvaultSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keyvaultsecrets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched keyvaultSecret.
func (c *keyvaultSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KeyvaultSecret, err error) {
	result = &v1alpha1.KeyvaultSecret{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("keyvaultsecrets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
