#!/bin/bash
OUTDIR=bin/

go build -o $OUTDIR

GOOS=windows go build -o $OUTDIR
