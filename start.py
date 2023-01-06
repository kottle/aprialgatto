#!/usr/bin/env python3
# so that script can be run from Brickman

from ev3dev.ev3 import *
from time   import sleep

# Connect EV3 color and touch sensors to any sensor ports
cl = ColorSensor() 

# Put the color sensor into RGB mode.
cl.mode='RGB-RAW'

while (True):    # Stop program by pressing touch sensor button
    red = cl.value(0)
    green=cl.value(1)
    blue=cl.value(2)
    print("Red: " + str(red) + ", Green: " + str(green) + ", Blue: " + str(blue))
    sleep(1)
