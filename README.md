Purpose:

Explore
https://github.com/adrg/xdg#--------

Motivation:
I want to store some config for some app in some proper place.  xdg tells me what is the proper place.

```
PS C:\Users\Administrator> .\purpleenvy.exe
main.go:48: Home data directory: C:\Users\Administrator\AppData\Local
main.go:49: Data directories: [C:\Users\Administrator\AppData\Roaming C:\ProgramData]
main.go:50: Home config directory: C:\Users\Administrator\AppData\Local
main.go:51: Config directories: [C:\ProgramData C:\Users\Administrator\AppData\Roaming]
main.go:52: Home state directory: C:\Users\Administrator\AppData\Local
main.go:53: Cache directory: C:\Users\Administrator\AppData\Local\cache
main.go:54: Runtime directory: C:\Users\Administrator\AppData\Local
main.go:57: Home directory: C:\Users\Administrator
main.go:58: Application directories: [C:\Users\Administrator\AppData\Roaming\Microsoft\Windows\Start Menu\Programs C:\ProgramData\Microsoft\Windows\Start Menu\Programs]
main.go:59: Font directories: [C:\Windows\Fonts C:\Users\Administrator\AppData\Local\Microsoft\Windows\Fonts]
File created successfully.
main.go:72: Save the config file at: C:\Users\Administrator\AppData\Local\appname\config.yaml
main.go:88: Config file was found at: C:\Users\Administrator\AppData\Local\appname\config.yaml
PS C:\Users\Administrator>
```
