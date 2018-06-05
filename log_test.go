package main

import (
	"testing"
)

func TestLog(t *testing.T) {
	repo := newGit("my-repo")
	repo.Commit("Init")
	repo.Commit("Change 1")
	log := repo.Log()

	for {
		i := 0

		if log[i].id > len(log) {
			t.Error("Unexpected log length")
			break
		} else if log[i].id == 2 {
			if log[i].id == 1 {
				break
			}
			break
		}

		i++
	}
}
