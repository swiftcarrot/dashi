#!/bin/sh

set -e

make
cd ~/go/src/github.com/swiftcarrot/pokemon
rm -rf *
dashi new
dashi g scaffold pokemon number:string name:string weight_min:string weight_max:string height_min:string height_max:string classification:string max_cp:integer max_hp:integer image:string
git add .
git commit -m "build with `dashi version`"
git push
