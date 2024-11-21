#!/bin/bash

set -ueo pipefail

for file in *.flac; do
	raw=$(echo "$file" | cut -f 1 -d ".")
	wav="wav/$raw.wav"
	flac -fd "$file" -o "$wav"
	flac -f "$wav" -o flac/"$raw".flac
done
