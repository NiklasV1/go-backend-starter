#!/bin/bash

# Find all .env-template files in the directory
for template in **/*.env-template; do
  # Skip if no .env-template files are found
  [ -e "$template" ] || continue

  # Create corresponding .env filename
  env_file="${template%.env-template}.env"

  # Copy the template to the .env file
  cp "$template" "$env_file"

  echo "Created $env_file from $template"
done
