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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/alejandroEsc/k8s-provisioner-juju-example/pkg/apis/controller/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterCreatorLister helps list ClusterCreators.
type ClusterCreatorLister interface {
	// List lists all ClusterCreators in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterCreator, err error)
	// ClusterCreators returns an object that can list and get ClusterCreators.
	ClusterCreators(namespace string) ClusterCreatorNamespaceLister
	ClusterCreatorListerExpansion
}

// clusterCreatorLister implements the ClusterCreatorLister interface.
type clusterCreatorLister struct {
	indexer cache.Indexer
}

// NewClusterCreatorLister returns a new ClusterCreatorLister.
func NewClusterCreatorLister(indexer cache.Indexer) ClusterCreatorLister {
	return &clusterCreatorLister{indexer: indexer}
}

// List lists all ClusterCreators in the indexer.
func (s *clusterCreatorLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterCreator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterCreator))
	})
	return ret, err
}

// ClusterCreators returns an object that can list and get ClusterCreators.
func (s *clusterCreatorLister) ClusterCreators(namespace string) ClusterCreatorNamespaceLister {
	return clusterCreatorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterCreatorNamespaceLister helps list and get ClusterCreators.
type ClusterCreatorNamespaceLister interface {
	// List lists all ClusterCreators in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterCreator, err error)
	// Get retrieves the ClusterCreator from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ClusterCreator, error)
	ClusterCreatorNamespaceListerExpansion
}

// clusterCreatorNamespaceLister implements the ClusterCreatorNamespaceLister
// interface.
type clusterCreatorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterCreators in the indexer for a given namespace.
func (s clusterCreatorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterCreator, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterCreator))
	})
	return ret, err
}

// Get retrieves the ClusterCreator from the indexer for a given namespace and name.
func (s clusterCreatorNamespaceLister) Get(name string) (*v1alpha1.ClusterCreator, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustercreator"), name)
	}
	return obj.(*v1alpha1.ClusterCreator), nil
}
