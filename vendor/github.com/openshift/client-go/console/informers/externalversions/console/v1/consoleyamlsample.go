// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	consolev1 "github.com/openshift/api/console/v1"
	versioned "github.com/openshift/client-go/console/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/console/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/console/listers/console/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConsoleYAMLSampleInformer provides access to a shared informer and lister for
// ConsoleYAMLSamples.
type ConsoleYAMLSampleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ConsoleYAMLSampleLister
}

type consoleYAMLSampleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewConsoleYAMLSampleInformer constructs a new informer for ConsoleYAMLSample type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConsoleYAMLSampleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConsoleYAMLSampleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredConsoleYAMLSampleInformer constructs a new informer for ConsoleYAMLSample type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConsoleYAMLSampleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsoleYAMLSamples().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsoleYAMLSamples().Watch(context.TODO(), options)
			},
		},
		&consolev1.ConsoleYAMLSample{},
		resyncPeriod,
		indexers,
	)
}

func (f *consoleYAMLSampleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConsoleYAMLSampleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *consoleYAMLSampleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&consolev1.ConsoleYAMLSample{}, f.defaultInformer)
}

func (f *consoleYAMLSampleInformer) Lister() v1.ConsoleYAMLSampleLister {
	return v1.NewConsoleYAMLSampleLister(f.Informer().GetIndexer())
}
