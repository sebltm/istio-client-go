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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OCSPStapleLister helps list OCSPStaples.
// All objects returned here must be treated as read-only.
type OCSPStapleLister interface {
	// List lists all OCSPStaples in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.OCSPStaple, err error)
	// OCSPStaples returns an object that can list and get OCSPStaples.
	OCSPStaples(namespace string) OCSPStapleNamespaceLister
	OCSPStapleListerExpansion
}

// oCSPStapleLister implements the OCSPStapleLister interface.
type oCSPStapleLister struct {
	indexer cache.Indexer
}

// NewOCSPStapleLister returns a new OCSPStapleLister.
func NewOCSPStapleLister(indexer cache.Indexer) OCSPStapleLister {
	return &oCSPStapleLister{indexer: indexer}
}

// List lists all OCSPStaples in the indexer.
func (s *oCSPStapleLister) List(selector labels.Selector) (ret []*v1beta1.OCSPStaple, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.OCSPStaple))
	})
	return ret, err
}

// OCSPStaples returns an object that can list and get OCSPStaples.
func (s *oCSPStapleLister) OCSPStaples(namespace string) OCSPStapleNamespaceLister {
	return oCSPStapleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OCSPStapleNamespaceLister helps list and get OCSPStaples.
// All objects returned here must be treated as read-only.
type OCSPStapleNamespaceLister interface {
	// List lists all OCSPStaples in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.OCSPStaple, err error)
	// Get retrieves the OCSPStaple from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.OCSPStaple, error)
	OCSPStapleNamespaceListerExpansion
}

// oCSPStapleNamespaceLister implements the OCSPStapleNamespaceLister
// interface.
type oCSPStapleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OCSPStaples in the indexer for a given namespace.
func (s oCSPStapleNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.OCSPStaple, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.OCSPStaple))
	})
	return ret, err
}

// Get retrieves the OCSPStaple from the indexer for a given namespace and name.
func (s oCSPStapleNamespaceLister) Get(name string) (*v1beta1.OCSPStaple, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("ocspstaple"), name)
	}
	return obj.(*v1beta1.OCSPStaple), nil
}
