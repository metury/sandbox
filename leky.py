#!/usr/bin/env python3

from datetime import datetime, timedelta

def print_general(times, rep):
	""" Generaly print for every day the gicen times with checks. """
	today = datetime.now()
	print(f'- {today.strftime("%d/%m/%Y")}')
	for t in times:
		print(f"	- [ ] {t}");
	for i in range(1, rep+1):
		next_day = today + timedelta(days=i)
		print(f'- {next_day.strftime("%d/%m/%Y")}')
		for t in times:
			print(f"	- [ ] {t}");

def print_standardized(r, p, v, rep):
	""" Standardized morning, lunch and dinner. """
	l = []
	if r: l.append("Ráno")
	if p: l.append("Poledne")
	if v: l.append("Večer")
	print_leky2(l, rep)

print_general(["Ráno", "Večer"], 30)
