# 9to5winawake
Keep your windows machine awake from 9 to 5 (local pc time)

## Installation

1. Download the [binary file](https://github.com/mhewedy/winawake/raw/master/9to5winawake.tgz) and unzip it (or build it from source using [build.sh](https://github.com/mhewedy/winawake/blob/master/build.sh))
2. open startup folder (WIN + R then type "shell:startup")
3. copy the downloaded binary (winawake.exe) to the startup folder.
4. have fun!


## How it works?
It sends the NULL (0x00) key every 1 minute.

---

Based on code from [keybd_event](https://github.com/micmonay/keybd_event)
