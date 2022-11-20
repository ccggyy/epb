#!/bin/bash

OS="$(uname)"
if [[ "${OS}" == "Linux" ]]; then
  UNAME_MACHINE="$(uname -m)"
elif [[ "${OS}" == "Darwin" ]]; then
  UNAME_MACHINE="$(/usr/bin/uname -m)"
else
  abort "Only supported on macOS and Linux."
fi

EPB_VERSION='0.0.1'
EPB_FILENAME="epb_${EPB_VERSION}_${OS}_${UNAME_MACHINE}.tar.gz"
EPB_DOWNLOAD_URL="https://github.com/ccggyy/epb/releases/download/${EPB_VERSION}/${EPB_FILENAME}"

curl -L "${EPB_DOWNLOAD_URL}" | tar xzf - -C /usr/local/bin/
