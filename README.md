# GitHub Actions demo for a monorepo Go project

The purpose of this repository is to demonstrate using a GitHub action as a pull request status check
in a monorepo where some of the applications are Golang-based and some are not.

There are 3 folders in this demo repo:

* `mono-repo-folder-1` containing a Golang application
* `mono-repo-folder-2` containing a different Golang application
* `mono-repo-folder-3` NOT containing a Golang application

The demonstrated configuration runs `go build` and `go test` only when in a PR there are changes in the Golang applications' folders,
and only for the folder in which the changes are located.

The open PRs in this repo demonstrate the different GitHub action behaviours.

## Adding the action as a status check

* First you need the action to be executed at least once. For example, opening a sample PR after the workflow file is in the `main` branch will execute the action.
* Go to the repo settings -> Add branch protection rule -> `Require status checks to pass before merging`.
* In the text box type the job name `build-and-test` and click on the suggestion that pops up.
* Fill all other details in the `Branch protection rule` page and click `Create` to save the changes.
