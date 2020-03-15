#!/usr/bin/env bash

function wd() {
  OUTPUT=`/path/to/DirectoryWarp $@`

  if [ $? -eq 2 ]
    then cd "$OUTPUT"
    else echo "$OUTPUT"
  fi
}
