## Slidefy

> The way to create slides when you're in a hurry, but you want them to look good.

### Installation

_Using go get_
```sh
go get -u github.com/Halfeld/slidefy
```

_or by the binary file_

```sh
git clone https://github.com/Halfeld/slidefy.git
cd slidefy
./slidefy [flags]
```

### Why?

If you are a lecturer, you have probably done some lecture in a hurry because of the time, but when we do that, the slides become ugly, even doing the basics.
So the intent of this package is to create the slides just from a json!

### Usage

_My json sctruture_
```json
[
  {
    "title": "Awesome presentation!",
    "desc": "Oh god this is so good.."
  },
  {
    "title": "Second slide.",
    "desc": "So far, so good.."
  },
  {
    "title": "Awesome end!",
    "desc": "Everything fine."
  }
]
```

_Command line usage_

Basic usage
```
slidefy --json path/to/json.json --pdfoutput path/to/file.pdf
```

For more
```
slidefy -h
```

_Final pdf generated..._

![Pdf Example](./screenshots/pdf-example.png)

### TO-DO

- [x] Support to user give a background path, like png images!
- [ ] Support to another font colors.
- [ ] Support to markers.
- [ ] Support to texts.
- [x] Support to give images as well.
- [ ] Support to add images on slide.
- [ ] Support to add a config file on home called `~/.slidefy.json`, where will existir some flags.