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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	controller_v1alpha1 "github.com/alejandroEsc/k8s-provisioner-juju-example/pkg/apis/controller/v1alpha1"
	versioned "github.com/alejandroEsc/k8s-provisioner-juju-example/pkg/client/clientset/versioned"
	internalinterfaces "github.com/alejandroEsc/k8s-provisioner-juju-example/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/alejandroEsc/k8s-provisioner-juju-example/pkg/client/listers/controller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterCreatorInformer provides access to a shared informer and lister for
// ClusterCreators.
type ClusterCreatorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterCreatorLister
}

type clusterCreatorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterCreatorInformer constructs a new informer for ClusterCreator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterCreatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterCreatorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterCreatorInformer constructs a new informer for ClusterCreator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterCreatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControllerV1alpha1().ClusterCreators(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControllerV1alpha1().ClusterCreators(namespace).Watch(options)
			},
		},
		&controller_v1alpha1.ClusterCreator{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterCreatorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterCreatorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterCreatorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&controller_v1alpha1.ClusterCreator{}, f.defaultInformer)
}

func (f *clusterCreatorInformer) Lister() v1alpha1.ClusterCreatorLister {
	return v1alpha1.NewClusterCreatorLister(f.Informer().GetIndexer())
}
