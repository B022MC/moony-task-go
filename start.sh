#!/bin/bash
set -x
ROOT=$(cd `dirname $0`; pwd)
cd $ROOT

if [ "$BINARY" = "" ];then
    BINARY="moony-task"
fi

while [ true ];do
    $ROOT/bin/$BINARY $@ >> log/run.log 2>&1
    sleep 60
done


