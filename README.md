# vaahan

```
docker build -t vaahan .
docker run --rm -it -v $(pwd)/src/vaahan:/workspace/src/vaahan -p 18770:18770 --name=vaahan_1 vaahan bash
```
