# winawake
Keep your windows machine awake

## Installation

1. Download the [binary file](https://github.com/mhewedy/winawake/blob/master/winawake.exe) (or build it from source using [build.sh](https://github.com/mhewedy/winawake/blob/master/build.sh))
2. open startup folder (WIN + R then type "shell:startup")
3. create a shortcut for this binary in the startup folder.
4. have fun!


## How it works?
It sends the NULL (0x00) key every 1 minute.

Based on code from [keybd_event](https://github.com/micmonay/keybd_event)
