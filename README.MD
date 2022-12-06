# Bluetooth Reconnect

This is a simple script for windows that will reconnect your bluetooth device to fix delays if you have walked out of
range.

# Requirements

You need install bluetooth commandline tool installed on your computer. You can get it
from [here](https://bluetoothinstaller.com/bluetooth-command-line-tools)
you will need to add the bluetooth commandline tools to your path.

After you have installed the bluetooth command line tools you will need to find the address of your bluetooth device.
You can do this by running the following command in a command prompt.

    btdiscovery

This will list all the bluetooth devices in range. You will need to find the address of your device. It will look
something like this `00:00:00:00:00:00`

# Usage

I have bound this script to a keyboard shortcut. You can do this by creating a desktop shortcut to the script and
then right clicking on the shortcut and selecting properties. Then you can add a keyboard shortcut to the shortcut.