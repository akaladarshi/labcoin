#!/bin/bash

# Ensure the script exits if a command fails
set -e

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Ensure solc is installed
if ! command_exists solc; then
    echo "solc not found, installing..."
    npm install -g solc
else
    echo "solc is already installed"
fi

# Ensure abigen is installed
if ! command_exists abigen; then
    echo "abigen not found, installing..."
    go get -u github.com/ethereum/go-ethereum/cmd/abigen
else
    echo "abigen is already installed"
fi

# Define the paths
CONTRACTS_DIR="contracts"
#ARTIFACTS_DIR="artifacts/contracts"
BUILD_DIR="build"
BINDINGS_DIR="services/bindings"

# Ensure the bindings directory exists
mkdir -p $BINDINGS_DIR

CONTRACT_FILES=("Research.sol")

# Generate Go bindings for each contract
for CONTRACT in "${CONTRACT_FILES[@]}"; do
  # Compile the contract
  solc --abi --bin --overwrite  "$CONTRACTS_DIR"/"$CONTRACT" -o build

  CONTRACT_NAME=$(basename "$CONTRACT" .sol)

  CONTRACT_NAME=$(basename "$CONTRACT" .sol)
  ABI_FILE="$BUILD_DIR/$CONTRACT_NAME.abi"
  BIN_FILE="$BUILD_DIR/$CONTRACT_NAME.bin"
  OUTPUT_FILE="$BINDINGS_DIR/$CONTRACT_NAME.go"

  echo "Generating Go bindings for $CONTRACT_NAME..."
  # shellcheck disable=SC2086
  abigen --abi="$ABI_FILE" --bin=$BIN_FILE --pkg=bindings --type $CONTRACT_NAME --out=$OUTPUT_FILE
done

echo "Go bindings generated successfully."