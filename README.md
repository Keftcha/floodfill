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



It will generate an output like that:


---

Unlike my cousin, this program can color without overstepping.
