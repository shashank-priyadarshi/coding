# largest palindrome product
def checkPalin(num):
    num = str(num)
    left = num[:len(num) // 2]
    right = num[len(num) // 2:]
    return left == right[::-1]


maxPalin = 0
for x in range(100, 1000):
    for y in range(100, 1000):
        z = x * y
        if checkPalin(z) and z > maxPalin:
            maxPalin = z
print(maxPalin)
