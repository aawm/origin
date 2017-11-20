// This file was automatically generated by informer-gen

package internalversion

import (
	apps "github.com/openshift/origin/pkg/apps/apis/apps"
	internalinterfaces "github.com/openshift/origin/pkg/apps/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/apps/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/apps/generated/listers/apps/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DeploymentConfigInformer provides access to a shared informer and lister for
// DeploymentConfigs.
type DeploymentConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.DeploymentConfigLister
}

type deploymentConfigInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewDeploymentConfigInformer constructs a new informer for DeploymentConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeploymentConfigInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Apps().DeploymentConfigs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Apps().DeploymentConfigs(namespace).Watch(options)
			},
		},
		&apps.DeploymentConfig{},
		resyncPeriod,
		indexers,
	)
}

func defaultDeploymentConfigInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewDeploymentConfigInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *deploymentConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apps.DeploymentConfig{}, defaultDeploymentConfigInformer)
}

func (f *deploymentConfigInformer) Lister() internalversion.DeploymentConfigLister {
	return internalversion.NewDeploymentConfigLister(f.Informer().GetIndexer())
}
