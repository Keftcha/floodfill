# Floodfill

Generate floodfill gif

## Usage

This program is build to use in command line mode.

```shell
$ floodfill SOURCE DEST
```
SOURCE → png image used for source
DEST → gif output

The flag `--delay <int>` can be used to set delay between each gif frame,
default is 10

```shell
floodfill --delay 3 SOURCE DEST
```

For a source like this:

![](https://raw.githubusercontent.com/Keftcha/floodfill/develop/images/input_800_600.png)


It will generate an output like that:

![](https://raw.githubusercontent.com/Keftcha/floodfill/develop/images/output_800_600.gif)

A white pixel will be colored. Starting points are one pixel colored different form black or white.
And black borders will stop the propagation.

---

Unlike my cousin, this program can color without overstepping.
