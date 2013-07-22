// +build samples

package main

func init() {
	addTestCases(fiximportTests, fiximport)
}

var fiximportTests = []testCase{
	{
		Name: "fiximport.0",
		In: `package main

import (
	_ "example.com/mypkg"
)
`,
		Out: `package main

import (
	_ "example.com/mypkg2"
)
`,
	},
}
