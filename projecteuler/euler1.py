# multiples of 3 or 5
temp = 3 
total = 0 
while temp < 1000:
    if temp % 3 == 0 or temp % 5 == 0:
        total += temp
    temp += 1
print(total)