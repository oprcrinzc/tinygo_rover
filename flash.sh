#!/bin/bash
tinygo build --target=pico2
picotool uf2 convert rover.elf rover.uf2
echo "press bootsel on pico then connect to computer [enter] to continue:"
read enter
picotool load -f rover.uf2
picotool reboot -f
