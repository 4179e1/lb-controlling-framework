/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "tkestack.io/lb-controlling-framework/pkg/apis/lbcf.tke.cloud.tencent.com/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackendGroupLister helps list BackendGroups.
type BackendGroupLister interface {
	// List lists all BackendGroups in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.BackendGroup, err error)
	// BackendGroups returns an object that can list and get BackendGroups.
	BackendGroups(namespace string) BackendGroupNamespaceLister
	BackendGroupListerExpansion
}

// backendGroupLister implements the BackendGroupLister interface.
type backendGroupLister struct {
	indexer cache.Indexer
}

// NewBackendGroupLister returns a new BackendGroupLister.
func NewBackendGroupLister(indexer cache.Indexer) BackendGroupLister {
	return &backendGroupLister{indexer: indexer}
}

// List lists all BackendGroups in the indexer.
func (s *backendGroupLister) List(selector labels.Selector) (ret []*v1beta1.BackendGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BackendGroup))
	})
	return ret, err
}

// BackendGroups returns an object that can list and get BackendGroups.
func (s *backendGroupLister) BackendGroups(namespace string) BackendGroupNamespaceLister {
	return backendGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackendGroupNamespaceLister helps list and get BackendGroups.
type BackendGroupNamespaceLister interface {
	// List lists all BackendGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.BackendGroup, err error)
	// Get retrieves the BackendGroup from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.BackendGroup, error)
	BackendGroupNamespaceListerExpansion
}

// backendGroupNamespaceLister implements the BackendGroupNamespaceLister
// interface.
type backendGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackendGroups in the indexer for a given namespace.
func (s backendGroupNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.BackendGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BackendGroup))
	})
	return ret, err
}

// Get retrieves the BackendGroup from the indexer for a given namespace and name.
func (s backendGroupNamespaceLister) Get(name string) (*v1beta1.BackendGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("backendgroup"), name)
	}
	return obj.(*v1beta1.BackendGroup), nil
}
