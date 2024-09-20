#!/bin/bash

for file in contracts/abi/*.abi; do
  filename=$(basename -- "$file" .abi)
  abigen --abi "$file" --pkg contracts --type "$filename" --out "contracts/$filename.go"
done
