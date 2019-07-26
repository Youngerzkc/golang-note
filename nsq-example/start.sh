#!/bin/sh
nohup ./bin/nsqlookupd > /dev/null 2>&1 &
nohup ./bin/nsqd --lookupd-tcp-address=127.0.0.1:4160 > /dev/null 2>&1 &
nohup ./bin/nsqadmin --lookupd-http-address=127.0.0.1:4161 > /dev/null 2>&1 &
