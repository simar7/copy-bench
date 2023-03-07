## Results

### Small file
goos: darwin
goarch: amd64
pkg: github.com/simar7/copy-test
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
            │  small.readall   │             small.copy             │
            │      sec/op      │      sec/op       vs base          │
SmallFile-8   0.00002690n ± 0%   0.00002440n ± 0%  -9.29% (n=10000)

            │ small.readall │            small.copy             │
            │     B/op      │    B/op     vs base               │
SmallFile-8      0.000 ± 0%   0.000 ± 0%  ~ (p=1.000 n=10000) ¹
¹ all samples are equal

            │ small.readall │            small.copy             │
            │   allocs/op   │ allocs/op   vs base               │
SmallFile-8      0.000 ± 0%   0.000 ± 0%  ~ (p=1.000 n=10000) ¹
¹ all samples are equal


### Big file
goos: darwin
goarch: amd64
pkg: github.com/simar7/copy-test
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
          │  big.readall  │            big.copy            │
          │    sec/op     │    sec/op      vs base         │
BigFile-8   0.08135n ± 3%   0.04829n ± 1%  -40.65% (n=100)

          │ big.readall │            big.copy             │
          │    B/op     │    B/op     vs base             │
BigFile-8    0.000 ± 0%   0.000 ± 0%  ~ (p=1.000 n=100) ¹
¹ all samples are equal

          │ big.readall │            big.copy             │
          │  allocs/op  │ allocs/op   vs base             │
BigFile-8    0.000 ± 0%   0.000 ± 0%  ~ (p=1.000 n=100) ¹
¹ all samples are equal

