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

func TestHash(t *testing.T) {
	t.Parallel()
	t.Run("new", func(t *testing.T) {
		t.Parallel()
		h := ed2k.New()
		sum := h.Sum(nil)
		want := "31d6cfe0d16ae931b73c59d7e0c089c0"
		got := fmt.Sprintf("%x", sum)
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 0)
		h := ed2k.New()
		_, err := h.Write(data)
		if err != nil {
			t.Fatal(err)
		}
		sum := h.Sum(nil)
		want := "31d6cfe0d16ae931b73c59d7e0c089c0"
		got := fmt.Sprintf("%x", sum)
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("smallzero", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 9728000)
		h := ed2k.New()
		_, err := h.Write(data)
		if err != nil {
			t.Fatal(err)
		}
		sum := h.Sum(nil)
		want := "d7def262a127cd79096a108e7a9fc138"
		got := fmt.Sprintf("%x", sum)
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("bigzero", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 19456000)
		h := ed2k.New()
		_, err := h.Write(data)
		if err != nil {
			t.Fatal(err)
		}
		sum := h.Sum(nil)
		want := "194ee9e4fa79b2ee9f8829284c466051"
		got := fmt.Sprintf("%x", sum)
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("smallzero_partial", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 9728000)
		h := ed2k.New()
		_, err := h.Write(data[:1000000])
		if err != nil {
			t.Fatal(err)
		}
		_, err = h.Write(data[1000000:])
		if err != nil {
			t.Fatal(err)
		}
		sum := h.Sum(nil)
		want := "d7def262a127cd79096a108e7a9fc138"
		got := fmt.Sprintf("%x", sum)
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("bigzero_blockwise", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 19456000)
		h := ed2k.New()
		_, err := h.Write(data[:9728000])
		if err != nil {
			t.Fatal(err)
		}
		_, err = h.Write(data[9728000:])
		if err != nil {
			t.Fatal(err)
		}
		sum := h.Sum(nil)
		want := "194ee9e4fa79b2ee9f8829284c466051"
		got := fmt.Sprintf("%x", sum)
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("bigzero_partial", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 19456000)
		h := ed2k.New()
		_, err := h.Write(data[:10728000])
		if err != nil {
			t.Fatal(err)
		}
		_, err = h.Write(data[10728000:])
		if err != nil {
			t.Fatal(err)
		}
		sum := h.Sum(nil)
		want := "194ee9e4fa79b2ee9f8829284c466051"
		got := fmt.Sprintf("%x", sum)
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
}
