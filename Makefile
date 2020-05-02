lint:
	golangci-lint run

test:
	go test -v -cover -coverprofile=cover.out

coverout:
	go tool cover -html=cover.out

bench_abs:
	go test -v -bench=. -run=^BenchmarkAbsInt32*$$ -benchtime=2000000000x
