## 1. gc rate && ballast
> https://github.com/golang/go/issues/23044

```
> go run ballast/ballast.go
2021/11/26 16:53:10 NumGC: 562 562/s, STW: max 178.309µs, avg 26.892µs; Heap: 4MB 1GB/s, live objects: 138B, ReadMemStats: 19.466µs
2021/11/26 16:53:11 NumGC: 1132 570/s, STW: max 3.270765ms, avg 44.792µs; Heap: 4MB 1GB/s, live objects: 150B, ReadMemStats: 31.161µs
2021/11/26 16:53:12 NumGC: 1698 566/s, STW: max 2.791701ms, avg 41.026µs; Heap: 3MB 1GB/s, live objects: 151B, ReadMemStats: 26.94µs
2021/11/26 16:53:13 NumGC: 2290 592/s, STW: max 106.698µs, avg 26.127µs; Heap: 2MB 1GB/s, live objects: 150B, ReadMemStats: 31.826µs
2021/11/26 16:53:14 NumGC: 2824 534/s, STW: max 872.024µs, avg 29.711µs; Heap: 3MB 1GB/s, live objects: 153B, ReadMemStats: 23.873µs
```

add 1GB ballast:

```
> go run ballast/ballast.go -s 1000
2021/11/26 16:53:22 NumGC: 2 2/s, STW: max 30.855µs, avg 180ns; Heap: 1GB 2GB/s, live objects: 785B, ReadMemStats: 49.633µs
2021/11/26 16:53:23 NumGC: 4 2/s, STW: max 42.334µs, avg 491ns; Heap: 1GB 1GB/s, live objects: 667B, ReadMemStats: 48.002µs
2021/11/26 16:53:24 NumGC: 6 2/s, STW: max 42.334µs, avg 761ns; Heap: 1GB 1GB/s, live objects: 424B, ReadMemStats: 47.206µs
2021/11/26 16:53:25 NumGC: 8 2/s, STW: max 42.334µs, avg 1.05µs; Heap: 1GB 1GB/s, live objects: 298B, ReadMemStats: 97.083µs
2021/11/26 16:53:26 NumGC: 9 1/s, STW: max 59.814µs, avg 1.284µs; Heap: 1GB 1GB/s, live objects: 1KB, ReadMemStats: 61.611µs
2021/11/26 16:53:27 NumGC: 11 2/s, STW: max 83.663µs, avg 1.808µs; Heap: 1GB 1GB/s, live objects: 885B, ReadMemStats: 53.233µs
```