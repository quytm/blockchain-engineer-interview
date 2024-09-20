#!/bin/bash

# Create folder
mkdir -p contracts/abi

# Generate all smart contract
solc --include-path ../genomicdao/node_modules/ --base-path ../ --abi ../genomicdao/contracts/*.sol -o contracts/abi

# Only keep these contracts
contracts=(
  "Controller"
  "GeneNFT"
  "PostCovidStrokePrevention"
)

# Delete contracts not use
for abi_file in contracts/abi/*.abi; do
  filename=$(basename -- "$abi_file" .abi)

  check=false
  for sm in "${contracts[@]}"; do
    if [[ "$filename" == "$sm" ]]; then
      check=true
      break
    fi
  done

  if [[ $check == false ]]; then
    rm -f "$abi_file"
  fi
done
