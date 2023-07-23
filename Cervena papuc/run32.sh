#!/bin/bash

set -ueo pipefail

if [ $# -lt 1 ]; then
	echo "One argument must be present: number of tests."
	exit
fi

folder="tests"
results="RESULTS32"
tests=$1

mkdir -p $folder
printf "" > $results

echo "Tests for 32/8 cards in pack:"
printf "["
while [ $tests != 0 ]; do
	file=$folder/$tests
	go run cervena_papuca.go > $file
	printf 32 >> $results
	printf " " >> $results
	tail -n 1 $file >> $results
	rm $file
	tests=$((tests - 1))
	printf "="
done
echo "]"

./plot32.py
