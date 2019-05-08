#!/bin/bash
# 扫描目录下所有git库，并执行相应的操作
function getdir(){
  for element in `ls $1` 
  do  
    dir_or_file=$1"/"$element
    if [ -d $dir_or_file ]
    then 
      # echo $dir_or_file
      if [ -d $dir_or_file/.git ]
      then
        echo $dir_or_file
        cd $dir_or_file 
        git reset --hard #可以执行想要的git操作
        cd -
      fi  
      getdir $dir_or_file
    #else
      #echo $dir_or_file
    fi  
  done
}
getdir $1
