#!/bin/sh

set -e

make
cd ~/go/src/github.com/swiftcarrot/pokemon
git checkout .
git pull
rm -rf *
rm -rf .dockerignore .eslintrc.js .gitignore .storybook
dashi new
dashi g scaffold pokemon number:string name:string
make
cd packages/dashboard && yarn build && cd ../..
cd packages/app && yarn build && cd ../..
cd packages/storybook && yarn build-storybook && cd ../..
git add .
git commit -m "build with `dashi version`"
git push
