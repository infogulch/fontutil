# fontutil

fontutil is a small collection of utility functions to make working with fonts easier.

## Functions

### Find

Find returns the path to a font file given a font name. It's (intended to be) an OS-agnostic way to locate font files installed on the system. If you need to include fonts with your application instead of finding them on the system, this isn't for you.

Find currently only supports Windows but contributions for other systems are quite welcome! I'm also open to changing the interface. For example, separately specifying style, variant, or weight or a css-like list of prioritized font or family names. I'm unsure if/how these changes could be implemented on the different systems.

##### Supported platforms and resources for future implementation

- [x] Windows ([REG: Fonts Entries](windows))
- [ ] OS X ([Font locations and their purposes](mac))
- [ ] Linux ([Debian](debian), [Ubuntu](ubuntu), others?)

### FindAndParse

FindAndParse calls Find to find a font file, reads it, parses it, and returns the `*truetype.Font`.

### DrawWidth

DrawWidth returns the width of a string if it were drawn to a Context, as the X coordinate of a `raster.Point`.

### CenterX

CenterX returns the the start point required to center a string within a Rectangle, as the X coordinate of a `raster.Point`.

[windows]: https://support.microsoft.com/en-us/kb/102960
[mac]: https://support.apple.com/en-us/HT201722
[debian]: https://wiki.debian.org/Fonts
[ubuntu]: https://wiki.ubuntu.com/Fonts
