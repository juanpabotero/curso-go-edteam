#!/bin/sh
# script para testear todos los paquetes del proyecto

for d in $(go list ./...); do
	echo "Testeando el paquete $d"
	go test -v $d
done
