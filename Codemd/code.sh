#!/bin/bash

set -ueo pipefail

# Script for creating markdown file with all the code base in given folders.
# Supports languages: cpp java python cs c hpp h hs pl sh
# Each code block start with two ## at the start of the line. Then there is the file path.

set="cpp java cs c hpp h hs pl sh py"
export="export.md"
reverse=0

# If it is reversed then make files out of code blocks

self_name="./code.sh"

# Don't print myself as well.
# This file needs to be called code.sh

# Read arguments. -r is for reverse and everything else is export (last one).
for var in $@; do
	if [ $var == "-r" ]; then
		reverse=1
	else
		export=$var
	fi
done

if [ $reverse == 0 ]; then
	echo "# Kód" > $export
	echo "" >> $export
	# Put the language type to markdown.
	for type in $set; do
		if [ $type == "py" ]; then
			current="python"
		elif [ $type == "pl" ]; then
			current="prolog"
		elif [ $type == "hpp" ]; then
			current="cpp"
		elif [ $type == "h" ]; then
			current="c"
		else
			current=$type
		fi
		# Cat the file to the block
		for file in $(find . -name "*$type" -type f); do
			if [ "$file" != "$self_name" ]; then
				echo "## $file" >> "$export"
				echo "" >> "$export"
				echo "\`\`\`$current" >> "$export"
				cat "$file" >> "$export"
				echo "\`\`\`" >> "$export"
				echo "" >> "$export"
			fi
		done
	done
else
	current="tmp"
	print=0
	IFS='' # To use all spaces as well.
	while read line; do
		if [ $print == 1 ]; then
			end=$(echo $(echo "$line" | grep -o "\`\`\`"))
			if [ "$end" == "$line" ] && [ "$end" != "" ]; then
				print=0
			else
				echo "$line" >> "$current"
			fi
		else
			found=$(echo $(echo "$line" | grep -o '## .*'))
			start=$(echo $(echo "$line" | grep -o '```[a-z].*'))
			if [ "$found" == "$line" ] && [ "$found" != "" ]; then
				mkdir -p "$(echo "$found" | grep -o "\./.*/")"
				printf "" > "$(echo "$found" | grep -o "[^\ #]*")"
				current=$(echo "$(echo "$found" | grep -o "[^\ #]*")")
			elif [ "$start" == "$line" ] && [ "$start" != "" ]; then
				print=1
			fi
		fi
	done < $export
fi
