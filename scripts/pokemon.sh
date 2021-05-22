#!/bin/sh

set -e

make
cd ~/go/src/github.com/swiftcarrot/pokemon
git checkout .
git pull
rm -rf *
rm -rf .dockerignore .eslintrc.js .gitignore .storybook
dashi new
dashi g scaffold pokemon number:string name:string weight_min:string weight_max:string height_min:string height_max:string classification:string types:strings resistant:strings attacks:text weaknesses:strings flee_rate:float max_cp:integer evolutions:uuids evolution_requirements_amout evolution_requirements_name max_hp:integer image:string
make
cd packages/dashboard && yarn build && cd ../..
cd packages/app && yarn build && cd ../..
yarn build-storybook
git add .
git commit -m "build with `dashi version`"
git push
