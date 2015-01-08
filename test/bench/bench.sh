#!/bin/bash

go test -bench=. -cpuprofile cpu.out
go tool pprof --web bench.test cpu.out

