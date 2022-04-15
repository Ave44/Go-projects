#!/bin/bash
mkfifo p1 p2

go run  guessBot.go > p1 < p2 &
go run  zgadywanieDlaBota.go < p1 > p2 

rm -rf p1 p2