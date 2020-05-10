/*
Copyright 2019 The etcd-operator Authors

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

package v1beta2

import (
	internalinterfaces "github.com/coreos/etcd-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// EtcdBackups returns a EtcdBackupInformer.
	EtcdBackups() EtcdBackupInformer
	// EtcdClusters returns a EtcdClusterInformer.
	EtcdClusters() EtcdClusterInformer
	// EtcdRestores returns a EtcdRestoreInformer.
	EtcdRestores() EtcdRestoreInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// EtcdBackups returns a EtcdBackupInformer.
func (v *version) EtcdBackups() EtcdBackupInformer {
	return &etcdBackupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EtcdClusters returns a EtcdClusterInformer.
func (v *version) EtcdClusters() EtcdClusterInformer {
	return &etcdClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EtcdRestores returns a EtcdRestoreInformer.
func (v *version) EtcdRestores() EtcdRestoreInformer {
	return &etcdRestoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
