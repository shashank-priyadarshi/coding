# sum square difference
sum_of_squares = 385  # sum of squares from 1-10
total = 55  # sum of numbers from 1-10
iterNum = 11
n = int(input())
while iterNum <= n:
    sum_of_squares += iterNum * iterNum
    total += iterNum
    iterNum += 1
square_of_sums = total * total
print(square_of_sums - sum_of_squares)
