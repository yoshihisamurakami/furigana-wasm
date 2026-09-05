## Overview

This repository contains the source code for building the WebAssembly (WASM) file used by the Chrome extension **Furigana-kun**.

The JavaScript code in [furigana-js](https://github.com/yoshihisamurakami/furigana-js) depends on the WASM file built from this repository.

## Build

To build the WASM file, run:

```bash
GOOS=js GOARCH=wasm go build -o kagome.wasm .
```

# About wasm code
The code for wasm/kagome.wasm is available at https://github.com/yoshihisamurakami/furigana-wasm.

## Libraries Used

This extension uses the following open-source library:

| Library                                        | Purpose                                                     | License |
| ---------------------------------------------- | ----------------------------------------------------------- | ------- |
| [Kagome v2](https://github.com/ikawaha/kagome) | Japanese morphological analysis and word reading extraction | MIT     |

Special thanks to the developers and contributors of Kagome.

## License

This project is licensed under the MIT License.

You are free to use, modify, and distribute this software in accordance with the terms of the MIT License.

See the `LICENSE` file for details.
