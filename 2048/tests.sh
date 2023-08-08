#!/bin/bash

set -ueo pipefail

dir="tests"
results="RESULTS"

mkdir -p $dir
touch $results

counter=$((30))

while [[ $counter != 0 ]]; do
	counter=$(( $counter - 1 ))
	file=$dir/"random"$counter
	go run 2048.go -r > $file
	printf "0 " >> $results
	tail -n 1 $file >> $results
	file=$dir/"random"$counter
	go run 2048.go -b > $dir/"best"$counter
	printf "1 " >> $results
	tail -n 1 $file >> $results
	file=$dir/"random"$counter
	go run 2048.go -lu > $dir/"lu"$counter
	printf "2 " >> $results
	tail -n 1 $file >> $results
	file=$dir/"random"$counter
	go run 2048.go -c > $dir/"cycle"$counter
	printf "3 " >> $results
	tail -n 1 $file >> $results
done
