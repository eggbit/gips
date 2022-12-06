# gips

## About <a name = "about"></a>

A quick and dirty IPS patcher as an excuse to learn Go.

## Installing

Download the source code:

```
git clone https://github.com/eggbit/gips.git
```

Navigate to the directory and build (make sure you have the latest version of Go installed):

```
go build
```

That's it! There aren't any external dependencies so that's all you should have to do to have yourself a nice little executable for most of your IPS patching needs.

## Usage <a name = "usage"></a>

Gips requires two things: an .ips patch file and the ROM file it's going to be applied to. The command is:

```
gips [ips_path] [rom_path]
```
If the patching was successful it'll save the patched rom with ```_patched``` appended to the file name. Eg. ```rom.ext``` -> ```rom_patched.ext```

## What's Missing?
Right now this is just a general purpose IPS patcher but it could probably be expanded in the future. Apparently IPS32 is a thing and there are a whole bunch of other patching formats I'd like to explore like BPS.
