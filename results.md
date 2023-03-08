## Results

### Inputs
```shell
 ls -lrth *.txt
-rw-r--r--  1 simar  staff    11B Mar  6 21:52 small.txt
-rw-------  1 simar  staff   100M Mar  6 23:01 big.txt
```


### `io.ReadAll` vs `io.Copy` + `buf.Bytes`

#### Small file
```shell
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
```

#### Big file
```shell
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
```

### `io.Copy` + `buf.Bytes` + `json.Unmarshal` vs `json.NewDecoder`

#### Small file
```shell
goos: darwin
goarch: amd64
pkg: github.com/simar7/copy-test
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
       │      j.old       │                j.new                │
       │      sec/op      │      sec/op       vs base           │
JSON-8   0.00003070n ± 1%   0.00002750n ± 1%  -10.42% (n=10000)

       │   j.old    │               j.new               │
       │    B/op    │    B/op     vs base               │
JSON-8   0.000 ± 0%   0.000 ± 0%  ~ (p=1.000 n=10000) ¹
¹ all samples are equal

       │   j.old    │               j.new               │
       │ allocs/op  │ allocs/op   vs base               │
JSON-8   0.000 ± 0%   0.000 ± 0%  ~ (p=1.000 n=10000) ¹
¹ all samples are equal
```
