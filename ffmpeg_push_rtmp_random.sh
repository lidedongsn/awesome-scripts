#!/bin/bash
# 获取目录下随机文件（用于ffmpeg推送随机文件到rtmp服务器）（哔哩哔哩 (゜-゜)つロ 干杯~-bilibili）
count=0
count1=0
filelist=/root/source/file
for file in $filelist/*
do
  let count++
done
echo ${count}

function random()
{
  min=$1  
  max=$(($2-$min+1))  
  num=$(($RANDOM+1000000000))
  echo $(($num%$max+$min))  

}

index=$(random 0 ${count})
echo ${index}
for file in $filelist/*
do
  if [[ count1 -eq ${index} ]]; then
    ffmpeg  -re -i $file -vcodec libx264  -acodec aac -f flv -y  "rtmp://****"
  fi  
  let count1++
done
