#!/bin/bash
set -e

REPO="schraf/literate"
ARTIFACT_NAME="literate-generated.zip"
TEMP_DIR="literate_temp"
GOBIN=$(go env GOPATH)/bin
DOWNLOAD_URL="https://github.com/${REPO}/releases/latest/download/${ARTIFACT_NAME}"

echo "Downloading ${ARTIFACT_NAME}..."
if ! curl -L -O "${DOWNLOAD_URL}"; then
    echo "Error: Failed to download ${ARTIFACT_NAME}. Check if the release and asset exist."
    exit 1
fi

echo "Extracting ${ARTIFACT_NAME}..."
mkdir -p "${TEMP_DIR}"
unzip -q -o "${ARTIFACT_NAME}" -d "${TEMP_DIR}"

cd "${TEMP_DIR}"

echo "Installing literate..."
if go install .; then
	echo "Installed to ${GOBIN}/literate"
else
    echo "Error: 'go install' failed. Ensure Go is installed and your source is valid."
    exit 1
fi

cd ..
rm -rf "${TEMP_DIR}" "${ARTIFACT_NAME}"

