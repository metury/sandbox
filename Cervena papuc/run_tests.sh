#!/bin/bash

set -ueo pipefail

if [ $# -lt 3 ]; then
	echo "Three arguments must be present: number of tests, low limit of cards, high limit of cards."
	exit
fi

folder="tests"
results="RESULTS"
tests_origin=$1
lo_number=$2
hi_number=$3

number=$lo_number
tests=$tests_origin

mkdir -p $folder
printf "" > $results

while [[ $number -lt  $((hi_number)) ]]; do
	echo "Tests for $number cards in pack:"
	printf "["
	while [ $tests != 0 ]; do
		file=$folder/$tests
		go run cervena_papuca.go $number > $file
		printf $number >> $results
		printf " " >> $results
		tail -n 1 $file >> $results
		rm $file
		tests=$((tests - 1))
		printf "="
	done
	echo "]"
	number=$((number + 2))
	tests=$tests_origin
done

./plot.py
