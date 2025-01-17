/*
Copyright 2022 The KubeSphere Authors.

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

package jclient

import (
	jcredential "github.com/jenkins-zh/jenkins-client/pkg/credential"
	v1 "k8s.io/api/core/v1"
	"kubesphere.io/devops/pkg/client/devops"
	"kubesphere.io/devops/pkg/client/devops/util"
)

// CreateCredentialInProject creates a credential, then returns the ID
func (j *JenkinsClient) CreateCredentialInProject(projectID string, credential *v1.Secret) (id string, err error) {
	client := j.getClient()

	var cre interface{}
	if cre, err = util.ConvertSecretToCredential(credential); err != nil {
		return "", err
	}
	return "", client.CreateInFolder(projectID, cre)
}

// UpdateCredentialInProject updates a credential
func (j *JenkinsClient) UpdateCredentialInProject(projectID string, credential *v1.Secret) (id string, err error) {
	client := j.getClient()
	err = client.UpdateInFolder(projectID, credential.GetName(), credential)
	return
}

// GetCredentialInProject returns a credential
func (j *JenkinsClient) GetCredentialInProject(projectID, id string) (*devops.Credential, error) {
	return j.jenkins.GetCredentialInProject(projectID, id)
}

// DeleteCredentialInProject deletes a credential
func (j *JenkinsClient) DeleteCredentialInProject(projectID, id string) (string, error) {
	client := j.getClient()
	return id, client.DeleteInFolder(projectID, id)
}

func (j *JenkinsClient) getClient() *jcredential.CredentialsManager {
	return &jcredential.CredentialsManager{JenkinsCore: j.Core}
}
