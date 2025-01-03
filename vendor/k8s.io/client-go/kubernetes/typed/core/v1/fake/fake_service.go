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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/api/core/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	gentype "k8s.io/client-go/gentype"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// fakeServices implements ServiceInterface
type fakeServices struct {
	*gentype.FakeClientWithListAndApply[*v1.Service, *v1.ServiceList, *corev1.ServiceApplyConfiguration]
	Fake *FakeCoreV1
}

func newFakeServices(fake *FakeCoreV1, namespace string) typedcorev1.ServiceInterface {
	return &fakeServices{
		gentype.NewFakeClientWithListAndApply[*v1.Service, *v1.ServiceList, *corev1.ServiceApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("services"),
			v1.SchemeGroupVersion.WithKind("Service"),
			func() *v1.Service { return &v1.Service{} },
			func() *v1.ServiceList { return &v1.ServiceList{} },
			func(dst, src *v1.ServiceList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ServiceList) []*v1.Service { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ServiceList, items []*v1.Service) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
