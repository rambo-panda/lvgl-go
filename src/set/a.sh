#!/bin/bash


for file in `ls *.go`
do
echo `grep -nw "GetObj" ${file}`
  line=`grep -nw "GetObj" ${file} | awk -F: '{print $1}'`
  echo "$file ${line} ---"
  exit
  gsed -i "${line},$(expr $line + 2)d" $file
done
