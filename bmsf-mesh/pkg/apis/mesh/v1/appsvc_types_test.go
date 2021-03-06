/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package v1

import (
	"testing"
	"time"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestStorageAppSvc(t *testing.T) {
	key := types.NamespacedName{
		Name:      "foo",
		Namespace: "default",
	}
	created := &AppSvc{
		TypeMeta: metav1.TypeMeta{
			Kind:       "appsvc",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Spec: AppSvcSpec{
			Selector: make(map[string]string),
			ServicePorts: []ServicePort{
				ServicePort{
					Name:        "port",
					Protocol:    "http",
					ServicePort: 8080,
					ProxyPort:   8080,
				},
			},
		},
		Status: AppSvcStatus{
			LastUpdateTime: metav1.NewTime(time.Now()),
		},
	}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &AppSvc{}
	g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(created))

	// Test Updating the Labels
	updated := fetched.DeepCopy()
	updated.Labels = map[string]string{"hello": "world"}
	g.Expect(c.Update(context.TODO(), updated)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(updated))

	// Test Delete
	g.Expect(c.Delete(context.TODO(), fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.HaveOccurred())
}
