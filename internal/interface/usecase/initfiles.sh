#!/bin/bash

# List of folders
folders=("onboard" "seller" "buyer")

# Create files in each folder
for folder in "${folders[@]}"; do
    mkdir -p "$folder"
    touch "$folder/model.go" "$folder/service_impl.go" "$folder/service.go"
done
