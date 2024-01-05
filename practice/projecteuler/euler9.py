# special pythagorean triplet
a = 1
b = 0
c = 0
s = 1000
found = False
while (a < s / 3):
    b = a
    while (b < s / 2):
        c = s - a - b

        if (a * a + b * b == c * c):
            found = True
            break

        b += 1

    if (found):
        print(a * b * c)
        break

    a += 1
