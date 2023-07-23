#!/usr/bin/env python3
import matplotlib.pyplot as plt
from scipy import stats
from os.path import exists


result_file = "RESULTS"
results = []
numbers=[]

middle = []
mid_num = []

if exists(result_file):
	with open(result_file) as file:
		my_sum = 0
		current_num = 0
		counter = 0
		for line in file:
			parts = line.split(' ')
			results.append(int(parts[1][:-1]))
			numbers.append(int(parts[0]))
			if numbers[-1] == current_num:
				my_sum += results[-1]
				counter += 1
			elif current_num == 0:
				my_sum += results[-1]
				counter += 1
				current_num = numbers[-1]
			else:
				middle.append(my_sum / counter)
				mid_num.append(current_num)
				counter = 1
				my_sum = results[-1]
				current_num = numbers[-1]


def myfunc(x):
	return slope * x + intercept

if len(numbers) != 0:
	slope, intercept, r, p, std_err = stats.linregress(numbers, results)
	model = list(map(myfunc, numbers))
	plt.title("Červená papuč")
	plt.scatter(numbers, results, color='Red', label="Počet kol")
	plt.plot(numbers, model, color='Blue', label="Lineární regrese")
	plt.plot(mid_num, middle, color='Yellow', label="Průměr")
	plt.legend()
	plt.savefig("plot.png")
