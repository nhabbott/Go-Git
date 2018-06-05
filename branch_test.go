package main

import (
	"strconv"
	"testing"
)

var mapString string

func arrayToString(arr []commit) string {
	i := 0
	for {
		if i == len(arr) {
			break
		} else if i == len(arr)-1 {
			mapString += strconv.Itoa(arr[i].id)
			break
		}

		mapString += strconv.Itoa(arr[i].id) + "-"
		i++
	}
	defer func() { mapString = "" }()
	return mapString
}

func TestBranches(t *testing.T) {
	repo := newGit("my-repo")
	repo.Commit("Init")
	repo.Commit("Change 1")

	if arrayToString(repo.Log()) == "2-1" {
		repo.Checkout("new")
		repo.Commit("Change 3")

		if arrayToString(repo.Log()) == "3-2-0" {
			repo.Checkout("master")

			if arrayToString(repo.Log()) == "2-1" {
				repo.Commit("Change 3")

				if arrayToString(repo.Log()) == "4-2-1" {
					return
				}
				t.Errorf("Expected '4-2-0' got %s", arrayToString(repo.Log()))
			} else {
				t.Errorf("Expected '2-1' got %s", arrayToString(repo.Log()))
			}
		} else {
			t.Errorf("Expected '3-2-0' got %s", arrayToString(repo.Log()))
		}
	} else {
		t.Errorf("Expected '2-1' got %s", arrayToString(repo.Log()))
	}
}
