package main

import (
	"testing"
)

func TestCheckoutBranchSwitch(t *testing.T) {
	repo := newGit("my-repo")
	repo.Commit("Init")

	if repo.head.name == "master" {
		repo.Checkout("new")

		if repo.head.name == "new" {
			repo.Checkout("master")

			if repo.head.name == "master" {
				repo.Checkout("new")

				if repo.head.name == "new" {
					return
				}
				t.Error("Expected 'new' got ", repo.head.name)
			} else {
				t.Error("Expected 'master' got ", repo.head.name)
			}
		} else {
			t.Error("Expected 'new' got ", repo.head.name)
		}
	} else {
		t.Error("Expected 'master' got ", repo.head.name)
	}
}
