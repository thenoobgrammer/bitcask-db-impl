### Bitcask implementation

This is my attempt to understand how bitcask implementation work.

You can find more information about it [here](https://riak.com/assets/bitcask-intro.pdf)

### Initial take

What I intend to do first, is to be able to write some entries to a file, with a capped size at `1MB` for each data file.
The purpose is to emulate an HDD-To perform write operations on a hard drive we are following the _append-only_ process, where we only add entries and run a compactor at random during runtime to clear the deleted and corrupted data.

### Opening a directory

When we `open` a directory, we first read **all files** in order to get some information to build our hash table (key/dir map).

We do this by looping through each file, extract the entries at the same time we retrieve the files information-This will allow us to understand the size of the file and the offset (last byte position).

The **active file** is set on the first file we find with less occupied space than the capped amount, e.g

```bash
bitcask-001.data # 1 MB <-- full
bitcask-002.data # 1 MB <-- full
bitcask-003.data # 320 KB <-- active file
```

###### Offset

The offset is the **byte position** for the last added entry. This encourages fast processing for our next reads.

```
0 key="name",value="elie"
29 key="name",value="karl"
88 key="name",value="ralph" // <-- The offset for this file is 88
```
