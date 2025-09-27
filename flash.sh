tinygo build --target=pico2
picotool uf2 convert rover.elf rover.uf2
cp rover.uf2 /run/media/oprc/RP2350
