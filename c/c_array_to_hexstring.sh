#!/bin/sh
# use as "./search_for_array.sh your_file
sed $* -e 's/0x//g' -e 's/,//g' | xargs echo -n | sed 's/ //g' 
