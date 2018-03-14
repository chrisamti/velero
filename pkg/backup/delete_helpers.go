/*
Copyright 2018 the Heptio Ark contributors.

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

package backup

import (
	"fmt"

	"github.com/heptio/ark/pkg/apis/ark/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewDeleteBackupRequest creates a DeleteBackupRequest for the backup identified by name.
func NewDeleteBackupRequest(name string) *v1.DeleteBackupRequest {
	return &v1.DeleteBackupRequest{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: name + "-",
			Labels: map[string]string{
				v1.BackupNameLabel: name,
			},
		},
		Spec: v1.DeleteBackupRequestSpec{
			BackupName: name,
		},
	}
}

// NewDeleteBackupRequestListOptions creates a ListOptions with a label selector configured to
// find DeleteBackupRequests for the backup identified by name.
func NewDeleteBackupRequestListOptions(name string) metav1.ListOptions {
	return metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", v1.BackupNameLabel, name),
	}
}
