# even fibonacci numbers
f1 = 0
f2 = 1
f3 = 0
sum = 0
while f3 < 4000000:
	f3 = f2 + f1
	f1 = f2
	f2 = f3
	if f3 % 2 == 0:
		sum += f3
print(sum)