#!/bin/bash
set -e

REPO="schraf/literate"
ARTIFACT_NAME="literate-generated.zip"
DOWNLOAD_URL="https://github.com/${REPO}/releases/latest/download/${ARTIFACT_NAME}"

echo "Downloading latest release from ${REPO}..."

# Create a temporary directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

# Download the zip file
if ! curl -LSs -O "${DOWNLOAD_URL}"; then
    echo "Error: Failed to download ${ARTIFACT_NAME}. Check if the release and asset exist."
    exit 1
fi

echo "Extracting ${ARTIFACT_NAME}..."
unzip -q -o "${ARTIFACT_NAME}" 

echo "Installing literate..."
if go install ./...; then
	echo "Installed to $(go env GOPATH)/bin/literate"
else
    echo "Error: 'go install' failed. Ensure Go is installed and your source is valid."
    exit 1
fi

echo "Cleaning up temporary files..."
cd ~
rm -rf "${TEMP_DIR}" 

