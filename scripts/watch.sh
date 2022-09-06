#!/bin/bash

# File: watch.sh
# Date: 2022-09-05
# Author: Thomas vanBommel
# Description: This is a program that watches the project for changes. When
#              something has changed it reloads 'make nobuild' for development.
#              It likely this program will not work as intended if there is more
#              then 1 'go-build' process running.

PID=""

restartApplication() {
    # kill process if it exists
    if test $PID
    then
	kill $PID && echo "Process $PID killed" && PID=""
    fi

    # start application
    make nobuild &

    # await new go-build PID
    while ! test $PID
    do
	PID=$(ps -u | grep go-build | grep -v grep | awk '{ print $2 }')
	sleep 0.1
    done
    
    # print divider
    echo "PID: $PID -----------------------------------------------------------"
}

# start process
restartApplication

# restart process when something is modified
while inotifywait -rqe modify .
do
    restartApplication
done
