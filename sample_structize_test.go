// +build samples

package main

func init() {
	addTestCases(mypkgstructizeTests, mypkgstructize)
}

var mypkgstructizeTests = []testCase{
	{
		Name: "mypkgstructize.0",
		In: `package main

import mypkg "example.com/mypkg"

type M struct {
	s  mypkg.Widget
	ss []mypkg.Widget
	m  map[mypkg.Widget]mypkg.Widget
}

type S0 mypkg.Widget

func newM(s mypkg.Widget, ss ...mypkg.Widget) *M {
	return &M{
		s:  s,
		ss: ss,
	}
}

var GlobalM = newM(nil.(mypkg.Widget))

type Widget mypkg.Widget

var S Widget = Widget(M.s)
`,
		Out: `package main

import mypkg "example.com/mypkg"

type M struct {
	s  *mypkg.Widget
	ss []*mypkg.Widget
	m  map[*mypkg.Widget]*mypkg.Widget
}

type S0 *mypkg.Widget

func newM(s *mypkg.Widget, ss ...*mypkg.Widget) *M {
	return &M{
		s:  s,
		ss: ss,
	}
}

var GlobalM = newM(nil.(*mypkg.Widget))

type Widget *mypkg.Widget

var S Widget = Widget(M.s)
`,
	},
	{
		Name: "mypkgstruct.1",
		In: `package main

import "example.com/mypkg"

type Widget *gomypkg.Widget
`,
		Out: `package main

import "example.com/mypkg"

type Widget *gomypkg.Widget
`,
	},
}
