#!/bin/sh
# use as "./search_for_array.sh your_file
grep -E "(0x[a-zA-Z0-9]{1,2},{0,1}){2,}" $*
