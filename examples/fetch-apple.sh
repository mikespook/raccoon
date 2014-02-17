#!/usr/bin/env bash
ROOT=../

pushd . > /dev/null
cd $ROOT/cmd/raccoon
go build
popd > /dev/null

# Fetch top 100 free apps from itunes
$ROOT/cmd/raccoon/raccoon -url='http://www.apple.com/itunes/charts/free-apps/' -script='apple-apps.lua'

# Fetch top 100 paid apps from itunes
$ROOT/cmd/raccoon/raccoon -url='http://www.apple.com/itunes/charts/paid-apps/' -script='apple-apps.lua'
