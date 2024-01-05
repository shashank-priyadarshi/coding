# 10001st prime
primes = [2, 3, 5, 7, 11, 13]  # first six primes
foundPrimes = 6  # number of primes found and added to list
checkNum = 15
nthPrime = int(input())
while foundPrimes <= nthPrime:
    count = 0
    for item in primes:
        if checkNum % item == 0:
            count += 1
            break
    if count == 0:
        primes.append(checkNum)
        foundPrimes += 1
    checkNum += 2
print(primes[nthPrime])  # try generating the 100000th prime
