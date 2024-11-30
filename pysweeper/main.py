#!/usr/bin/env python3

import hra
"""první spuštění hry"""
h0 = hra.Hra()
st = input("Zadejte grafiku (T - terminál, A - advanced): ")
while st != "T" and st != "A":
    st = input("T nebo A: ")
h0.set_grafika(st)
h0.start()
