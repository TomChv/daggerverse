#!/bin/sh

DIR_ID=$(dagger query --doc ../tests/queries/get-dagger-node-sdk-repository.gql --progress=plain | jq -r .git.branch.tree.directory.id)

dagger query --doc ../tests/queries/node-"$1"-query.gql --var dirID=$DIR_ID