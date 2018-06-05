<<<<<<< HEAD
package main

type git struct {
	name         string
	lastCommitID int
	head         *branch
	branches     []branch
}

type commit struct {
	id      int
	parent  *commit
	message string
}

type branch struct {
	name   string
	commit *commit
}

func newGit(name string) *git {
	master := branch{name: "master", commit: &commit{}}

	branches := []branch{}
	branches = append(branches, master)

	git := git{name: name, lastCommitID: 0, head: &master, branches: branches}
	return &git
}

func (git *git) Commit(message string) *commit {
	id := git.lastCommitID + 1
	git.lastCommitID++

	newCommit := commit{id: id, parent: git.head.commit, message: message}
	git.head.commit = &newCommit
	return &newCommit
}

func (git *git) Log() []commit {
	head := git.head.commit
	history := []commit{}

	for {
		if head.id == 0 {
			break
		}

		history = append(history, *head)
		*head = *head.parent
	}
	return history
}

func (git *git) Checkout(branchName string) *git {
	for i, b := range git.branches {
		if b.name == branchName {
			//fmt.Println("Switched to existing branch: ", branchName)
			git.head = &git.branches[i]
			return git
		}
		break
	}

	newBranch := branch{name: branchName, commit: git.head.commit}
	git.branches = append(git.branches, newBranch)
	git.head = &newBranch
	//fmt.Println("Switched to new branch: ", branchName)
	return git
}

func main() {
}
=======
package main

import (
	"fmt"
)

type git struct {
	name         string
	lastCommitID int
	head         *branch
	branches     []branch
}

type commit struct {
	id      int
	parent  *commit
	message string
}

type branch struct {
	name   string
	commit *commit
}

func newGit(name string) *git {
	master := branch{name: "master", commit: &commit{}}

	branches := []branch{}
	branches = append(branches, master)

	git := git{name: name, lastCommitID: 0, head: &master, branches: branches}
	return &git
}

func (git *git) Commit(message string) *commit {
	id := git.lastCommitID + 1
	git.lastCommitID++

	newCommit := commit{id: id, parent: git.head.commit, message: message}
	git.head.commit = &newCommit
	return &newCommit
}

func (git *git) Log() []commit {
	head := git.head.commit
	history := []commit{}

	for {
		if head.id == 0 {
			break
		}

		history = append(history, *head)
		*head = *head.parent
	}
	return history
}

func (git *git) Checkout(branchName string) *git {
	for i, b := range git.branches {
		if b.name == branchName {
			fmt.Println("Switched to existing branch: ", branchName)
			git.head = &git.branches[i]
			return git
		}
		break
	}

	newBranch := branch{name: branchName, commit: git.head.commit}
	git.branches = append(git.branches, newBranch)
	git.head = &newBranch
	fmt.Println("Switched to new branch: ", branchName)
	return git
}

// Testing functions

func main() {
	fmt.Print("Log() Test: ")
	repo := newGit("my-repo")
	repo.Commit("Init")
	repo.Commit("Change 1")
	log := repo.Log()

	if len(log) == 2 {
		if log[0].id == 2 {
			if log[1].id == 1 {
				fmt.Println("Passed!")
				fmt.Println()
			}
		}
	}

	repo1 := newGit("my-repo")
	repo1.Commit("Init")
	repo1.Commit("Change 1")

	repo1.Checkout("new")
	repo1.Commit("testing")
	fmt.Println(repo1.head)
	fmt.Println(repo1.head.commit)
	repo1.Checkout("master")
	repo1.Commit("testing1")
	fmt.Println(repo1.head)
	fmt.Println(repo1.head.commit)
	repo1.Checkout("new")
	fmt.Println(repo1.head)
	fmt.Println(repo1.head.commit)
}
>>>>>>> 9777655a8509047c1c6ba8236e0825d341615a3e
