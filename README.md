# ascii-art

A CLI tool for generating ASCII Art from any image you like.

<img src="https://github.com/taga3s/ascii-art/assets/107479598/31ac8872-90d0-4a59-a99a-ae4e9d6587df" alt="ascii-art logo" width="670" style="display: block; margin: 0 auto;">

## Usage

Here is basic command. It supports `.jpg (.jpeg)` and `.png` image.

```
$ ascii-art ./cat.png
```

You can use `--threshold` or `-t` option to use any threshold. The default threshold is determined automatically by [OTSU's method](https://en.wikipedia.org/wiki/Otsu%27s_method).

```
$ ascii-art -t 128 ./cat.png
```

You can use `--magnification` or `-m` option to control the size of the ascii art. The default magnification is 1.0.

```
$ ascii-art -m 2.0 ./cat.png
```

## License

- [MIT](https://github.com/taga3s/ascii-art/blob/main/LICENSE)
