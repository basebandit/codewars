This kata is meant to test your grip on string manipulation hence the strings package.

In the solution test file I have included both tests and benchmarks for both my solution and other solutions that I found elegant for learning purposes.

It would help you also if you could write your own solution or contribute here and add a benchmark for your own solution as well.

To run the bench marks alone. Use the following command:

```Go
go test -run=^$ -cpu=1,2,4 -bench=. ./kata2
```

> Ensure your computer is idle before running any benchmarks for consistent repeated results.
> We are passing the `-run` flag a regex that matches nothing and we are using the `-cpu` flag to pass a list of values to run the benchmark with.
