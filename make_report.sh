#!/bin/sh
PXLS1=/home/niteadmin/works/go/pxls1/pxls1
TIMESTAMP=`date +%Y-%m-%d`
OUTP=/mnt/ad/nexs/基盤T/patches/$TIMESTAMP

mkdir -p "$OUTP"
"$PXLS1" ./var/yum_check_update/ "$OUTP/patches-$TIMESTAMP.xlsx"
"$PXLS1" -y ./var/yum_check_update/ "$OUTP/updates_db-$TIMESTAMP.yml"

echo "complete. output=$OUTP"
