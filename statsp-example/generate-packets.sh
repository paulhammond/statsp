#!/bin/bash

exec 3<> /dev/udp/127.0.0.1/8125

NAMES=( banana apple pear orange lemon lime )

while true; do
	sleep $(echo "scale=5; $RANDOM*0.5/32767" | bc)

	INDEX=0
	let "INDEX=$RANDOM % 6"
	printf "${NAMES[$INDEX]}:$RANDOM|g\n" >&3
done
