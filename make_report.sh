#!/bin/sh -ue
PXLS1=./pxls1
OUTP=./tmp
INPUT=~/works/ansible/t2/2020-09-patches/var/yum_check_update/

mkdir -p "$OUTP"
"$PXLS1" "$INPUT" "$OUTP/patches.xlsx"
"$PXLS1" -y "$INPUT" "$OUTP/updates_db.yml"

echo "complete. output=$OUTP"
