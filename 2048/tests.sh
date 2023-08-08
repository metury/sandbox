#!/bin/bash

set -ueo pipefail

if [ $# -lt 1 ]; then
	echo "One argument must be present: number of tests."
	exit
fi

dir="tests"
results="RESULTS"

mkdir -p $dir
touch $results

counter=$1

while [[ $counter != 0 ]]; do
	counter=$(( $counter - 1 ))
	
	file=$dir/"random"$counter
	go run 2048.go -r > $file
	printf "0 " >> $results
	tail -n 1 $file >> $results
	
	file=$dir/"best"$counter
	go run 2048.go -b > $file
	printf "1 " >> $results
	tail -n 1 $file >> $results
	
	file=$dir/"lu"$counter
	go run 2048.go -lu > $file
	printf "2 " >> $results
	tail -n 1 $file >> $results
	
	file=$dir/"cycle"$counter
	go run 2048.go -c > $file
	printf "3 " >> $results
	tail -n 1 $file >> $results
done

./plot.py
