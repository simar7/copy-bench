small-readall:
	go test -bench="BenchmarkReadAllFunc_SmallFile" -benchmem -count=10

big-readall:
	go test -bench="BenchmarkReadAllFunc_BigFile" -benchmem -count=10

small-copy:
	go test -bench="BenchmarkCopyFunc_SmallFile" -benchmem -count=10

big-copy:
	go test -bench="BenchmarkCopyFunc_BigFile" -benchmem -count=10

benchstat:
	benchstat smallfile.readall smallfile.copy
	benchstat bigfile.readall bigfile.copy

all:
	make small-readall > smallfile.readall
	make small-copy > smallfile.copy
	make big-readall > bigfile.readall
	make big-copy > bigfile.copy
	make benchstat 
