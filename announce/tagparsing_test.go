// Copyright (c) 2014-2016 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package announce

import (
	"github.com/bitmark-inc/bitmarkd/fault"
	"testing"
)

func TestValidTag(t *testing.T) {

	type testItem struct {
		txt string
		err error
	}

	testData := []testItem{

		{ // 0
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: nil,
		},
		{ // 1
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: nil,
		},
		{ // 2
			txt: "bitmark=v2 a=300.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidIPAddress,
		},
		{ // 3
			txt: "bitmark=v2 a=118.163.120.178;2001:x030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidIPAddress,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=335669 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidPortNumber,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=0 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidPortNumber,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=-12 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidPortNumber,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=335x669 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidPortNumber,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A761934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidFingerprint,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=461934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED04 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidFingerprint,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CZFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED04 s=32135 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidFingerprint,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=321359 c=32136 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidPortNumber,
		},

		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=321369 p=202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidPortNumber,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=1202c14ec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidPublicKey,
		},
		{
			txt: "bitmark=v2 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c1pec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidPublicKey,
		},
		{
			txt: "bitmark=v0 a=118.163.120.178;2001:b030:2314:0200:4649:583d:0001:0120 r=33566 f=48137A7A76934CAFE7635C9AC05339C20F4C00A724D7FA1DC0DC3875476ED004 s=32135 c=32136 p=202c1pec485c21d0d18e9dfd096bd760a558d5ee1139f8e4b2e15863433e7d51",
			err: fault.ErrInvalidDnsTxtRecord,
		},
		{
			txt: "hello world",
			err: fault.ErrInvalidDnsTxtRecord,
		},
	}

	for i, item := range testData {
		tags, err := parseTag(item.txt)

		if item.err != err {
			t.Fatalf("parseTag[%d]: %q  error: %v  expected: %v", i, item.txt, err, item.err)
		} else if nil == err {
			t.Logf("tags[%d]: %#v", i, tags)
		}
	}
}
