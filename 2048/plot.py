#!/usr/bin/env python3
import matplotlib.pyplot as plt
from scipy import stats
from os.path import exists


result_file = "RESULTS"
results = []
numbers=[]

if exists(result_file):
	with open(result_file) as file:
		current_num = 0
		counter = 0
		for line in file:
			parts = line.split(' ')
			results.append(int(parts[1][:-1]))
			numbers.append(int(parts[0]))
			if numbers[-1] == current_num:
				counter += 1
			elif current_num == 0:
				counter += 1
				current_num = numbers[-1]
			else:
				counter = 1
				current_num = numbers[-1]

middle = [0, 0, 0, 0]
mid_num = [0, 1, 2, 3]

for i in range(len(results)):
	middle[numbers[i]] += results[i]

counter = len(results) / 4

for i in range(len(middle)):
	middle[i] = middle[i] / counter

if len(numbers) != 0:
	slope, intercept, r, p, std_err = stats.linregress(numbers, results)
	def myfunc(x):
		return slope * x + intercept
	model = list(map(myfunc, numbers))
	plt.title("2048")
	plt.scatter(numbers, results, color='Red', label="Typ")
	plt.scatter(mid_num, middle, color='Blue', label="Průměr")
	plt.legend()
	plt.savefig("plot.png")

