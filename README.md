# gips

## About <a name = "about"></a>

A quick and dirty IPS patcher as an excuse to learn Go.

## Installing

Download the source code:

```
git clone https://github.com/eggbit/gips.git
```

Navigate to the directory and build:

```
go build
```

That's it! You should have a nice little executable for (some) of your IPS patching needs!

## Usage <a name = "usage"></a>

Gips requires two things: an .ips patch file and the ROM file it's going to be applied to. The command is:

```
gips [ips_path] [rom_path]
```
If the patching was successful it'll save the patched rom with ```_patched``` appended to the file name. Eg. ```rom.ext``` -> ```rom_patched.ext```

## What's Missing?
Right now it can't make use of patches that would make the resulting ROM larger than the ROM that's given. There are certain patches for Gameboy ROMs that turn them into Gameboy Color ROMs, for example. This is planned to be fixed in the future.

The comannd line structure could also be updated. And IPS32 is a thing?
