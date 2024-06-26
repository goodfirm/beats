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

//go:build !integration

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLoopback(t *testing.T) {
	check, err := IsLoopback("127.0.0.1")

	assert.NoError(t, err)
	assert.True(t, check)
}

func TestIsLoopback_false(t *testing.T) {
	check, err := IsLoopback("192.168.1.1")
	assert.NoError(t, err)
	assert.False(t, check)
}

func TestIsLoopback_error(t *testing.T) {
	check, err := IsLoopback("19216811")
	assert.Error(t, err)
	assert.False(t, check)
}
