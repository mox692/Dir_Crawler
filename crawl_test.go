package main

import "testing"

func Test_jumpToDir(t *testing.T) {
	inputs := []*Crawler{
		{
			fileNames: []string{"test1.txt"},
			path:      map[string]string{"test1.txt": "root/child/test1.txt"},
		},
	}

	expectDir := []string{"/Users/kimuramotoyuki/go/src/github.com/mox692/Dir_Crawler/root/child"}

	for i, v := range inputs {
		get := v.jumpToDir()

		if get != expectDir[i] {
			t.Errorf("jumpToDir err: expect %s, got %s\n", get, expectDir[i])
		}
	}
}
