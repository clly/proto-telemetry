#!/usr/bin/env bash

set -euf pipefail

WORK="go.work"
if [[ -e $WORK ]]; then
    S=1
    while IFS= read -r line; do
        if [[ $line == ')' ]]; then
            S=1
        fi
        # echo "$S $line"
        if [ $S -eq 0 ]; then
            echo -e "${line}" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//'
        fi

        if [[ $line == 'use (' ]]; then
            S=0
        fi
    done < $WORK
fi
