#!/bin/bash

# List of folders
folders=("buyer" "seller" "order" "product")

# Create files in each folder
for folder in "${folders[@]}"; do
    mkdir -p "$folder"
    touch "$folder/entity.go" "$folder/impl.go" "$folder/repository.go"
done
