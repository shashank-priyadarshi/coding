# 10001st prime
primes = [2, 3, 5, 7, 11, 13]  # first six primes
primeSum = sum(primes)
foundPrimes = 6  # number of primes found and added to list
checkNum = 15
nthNum = int(input())
while checkNum <= nthNum:
    count = 0
    for item in primes:
        if checkNum % item == 0:
            count += 1
            break
    if count == 0:
        primes.append(checkNum)
        primeSum += checkNum
        foundPrimes += 1
    print(checkNum)
    checkNum += 2
print(primeSum)
# try generating the 100000th prime
