/*
Copyright 2022 k0s authors

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

package token

import (
	"testing"

	bootstraptokenv1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/bootstraptoken/v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateKubeconfig(t *testing.T) {
	expected := `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: dGhlIGNlcnQ=
    server: the join URL
  name: k0s
contexts:
- context:
    cluster: k0s
    user: the user
  name: k0s
current-context: k0s
kind: Config
users:
- name: the user
  user:
    token: abcdef.0123456789abcdef
`

	tok := bootstraptokenv1.BootstrapTokenString{ID: "abcdef", Secret: "0123456789abcdef"}
	kubeconfig, err := GenerateKubeconfig("the join URL", []byte("the cert"), "the user", &tok)
	require.NoError(t, err)
	assert.Equal(t, expected, string(kubeconfig))
}
