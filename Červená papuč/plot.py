#!/usr/bin/env python3
import matplotlib.pyplot as plt
from scipy import stats
from os.path import exists


result_file = "RESULTS"
results = []
numbers=[]

if exists(result_file):
	with open(result_file) as file:
		for line in file:
			parts = line.split(' ')
			results.append(int(parts[1][:-1]))
			numbers.append(int(parts[0]))

def myfunc(x):
	return slope * x + intercept

if len(numbers) != 0:
	slope, intercept, r, p, std_err = stats.linregress(numbers, results)
	model = list(map(myfunc, numbers))
	plt.title("Červená papuč")
	plt.scatter(numbers, results, color='Red')
	plt.plot(numbers, model, color='Blue')
	plt.savefig("plot.png")
