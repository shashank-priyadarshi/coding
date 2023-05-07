# largest prime factor
def factorisation(num):
    a = 2
    while num > 1:
        if num % a == 0:
            factors.append(a)
            num /= a
            factorisation(num)
        a += 1
    return max(factors)


factors = []
print(factorisation(600851475143))
