// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package elasticsearchkeystoreresource

import (
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func Test_expandModel(t *testing.T) {
	esKeystoreRD := newResourceData(t, resDataParams{
		ID:        "some-random-id",
		Resources: newSampleElasticsearchKeystore(),
	})
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name string
		args args
		want *models.KeystoreContents
	}{
		{
			name: "parses the resource",
			args: args{d: esKeystoreRD},
			want: &models.KeystoreContents{
				Secrets: map[string]models.KeystoreSecret{
					"my_secret": {
						AsFile: ec.Bool(true),
						Value:  "supersecret",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expandModel(tt.args.d)
			assert.Equal(t, tt.want, got)
		})
	}
}