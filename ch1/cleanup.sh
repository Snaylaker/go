
#!/bin/bash

for file in *.go; do
    # Get the filename without the extension
    dir_name="${file%.*}"

    # Create a directory with the filename
    mkdir -p "$dir_name"

    # Move and rename the file to main.go inside the new directory
    mv "$file" "$dir_name/main.go"
done
