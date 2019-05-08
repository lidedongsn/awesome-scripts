#!/bin/sh

# 截取视频部分保存为gif（高清）开始时间及时长，支持坐标及大小, ./video2gif.sh <videofile> <out.gif> (来!斗图啊！)

start_time=00:48
duration=5
palette="/tmp/palette.png"

filters="fps=5,scale=240:-1:flags=lanczos,crop=240:240:0:50"

ffmpeg -v warning -ss $start_time -t $duration -i $1 -vf "$filters,palettegen" -y $palette
ffmpeg -v warning -ss $start_time -t $duration -i $1 -i $palette -lavfi "$filters [x]; [x][1:v] paletteuse" -y $2
