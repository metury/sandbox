#!/bin/bash

set -ueo pipefail

folder="tests"
results="RESULTS"
tests_origin=40
lo_number=40
hi_number=120

number=$lo_number
tests=$tests_origin
mkdir -p $folder
printf "" > $results

while [ $number != $((hi_number + 1)) ]; do
	echo "Tests for $number"
	while [ $tests != 0 ]; do
		file=$folder/$tests
		go run cervena_papuca.go $number > $file
		printf $number >> $results
		printf " " >> $results
		tail -n 1 $file >> $results
		rm $file
		tests=$((tests - 1))
		printf $((tests + 1))
		printf " "
	done
	echo ""
	number=$((number + 1))
	tests=$tests_origin
done

./plot.py
