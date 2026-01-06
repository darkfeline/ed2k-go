// Copyright (C) 2021  Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ed2k_test

import (
	"fmt"
	"testing"

	"go.felesatra.moe/hash/ed2k"
)

// testBytes deduplicates bytes for testing.
var testBytes [2 * 9728000]byte

func TestHash(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		data [][]byte
		want string
	}{
		{
			name: "new",
			data: nil,
			want: "31d6cfe0d16ae931b73c59d7e0c089c0",
		},
		{
			name: "empty write",
			data: [][]byte{testBytes[:0]},
			want: "31d6cfe0d16ae931b73c59d7e0c089c0",
		},
		{
			name: "one block",
			data: [][]byte{testBytes[:9728000]},
			want: "d7def262a127cd79096a108e7a9fc138",
		},
		{
			name: "one block multi writes",
			data: [][]byte{testBytes[:1000000], testBytes[:9728000-1000000]},
			want: "d7def262a127cd79096a108e7a9fc138",
		},
		{
			name: "one and partial blocks",
			data: [][]byte{testBytes[:9728000+1]},
			want: "06329e9dba1373512c06386fe29e3c65",
		},
		{
			name: "two block",
			data: [][]byte{testBytes[:2*9728000]},
			want: "194ee9e4fa79b2ee9f8829284c466051",
		},
		{
			name: "two block blockwise writes",
			data: [][]byte{testBytes[:9728000], testBytes[:9728000]},
			want: "194ee9e4fa79b2ee9f8829284c466051",
		},
		{
			name: "two block multi writes",
			data: [][]byte{testBytes[:10728000], testBytes[:2*9728000-10728000]},
			want: "194ee9e4fa79b2ee9f8829284c466051",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			h := ed2k.New()
			for _, d := range c.data {
				_, err := h.Write(d)
				if err != nil {
					t.Fatal(err)
				}
			}
			sum := h.Sum(nil)
			got := fmt.Sprintf("%x", sum)
			if got != c.want {
				t.Errorf("Got %s, want %s", got, c.want)
			}
			sum = h.Sum(nil)
			got = fmt.Sprintf("%x", sum)
			if got != c.want {
				t.Errorf("Second sum differs, got %s, want %s", got, c.want)
			}
		})

	}
}
