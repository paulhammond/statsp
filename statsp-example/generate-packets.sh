#!/bin/bash

exec 3<> /dev/udp/127.0.0.1/8125

GUAGES=( banana apple pear orange lemon lime )
SIGNS=( '+' '-' '' )
TIMERS=( throw catch jump skip hop drop)

INDEX=0
SIGNINDEX=0
PACKET=""

while true; do
	sleep $(echo "scale=5; $RANDOM*0.5/32767" | bc)
	if (( RANDOM % 2 )); then
		let "INDEX=$RANDOM % 6"
		let "SIGNINDEX=$RANDOM % 3"
		PACKET="${GUAGES[$INDEX]}:${SIGNS[$SIGNINDEX]}$RANDOM|g\n"
	else
		let "INDEX=$RANDOM % 6"
		PACKET="${TIMERS[$INDEX]}:$RANDOM|ms\n"
	fi
	printf $PACKET
	printf $PACKET >&3
done
