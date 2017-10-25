# vaahan

```
docker build -t vaahan .
docker run --rm -it -d -v $(pwd)/src/vaahan:/workspace/src/vaahan -p 18770:18770 --name=vaahan_1 --env CODE_ENV=dev vaahan
```
