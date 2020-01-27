# Filezipper

This is a console utility that compresses files into a zip archive.

Entry file can be a directory or file. Upon receiving the directory, 
the program packs each file into a separate archive. All files in the 
subdirectories will be reduced to a flat structure. When used with files 
with the same name, the record will be made for the last file

## Build or test

``` bash
# build filezipper
$ make build
```
``` bash
# test filezipper
$ make test
```

## Used

``` bash
$ ./dist/filezipper -entry ./files -out ./zip
```

## Example

``` bash
$ ./dist/filezipper -entry ./files -out ./zip
filezipper v0.1.0
Start processingArchiving 12.zip ...
Archiving testadsfsa.zip ...
Archiving web.zip ...
Finished
```