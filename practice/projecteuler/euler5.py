# smallest multiple
product = 2520  # smallest number evenly divisible by 1-10
minNum = int(input())  # greater than 10
maxNum = int(input())  # greater than minNum
while minNum < maxNum:
    if product % minNum != 0:
        checkNum = 2
        while minNum % checkNum != 0:
            checkNum += 1
        product *= checkNum
    minNum += 1
print(product)  # checked it in the range 11-500, answer in the docstring
"""73239622318952846593863874519042298829761338251289259046349190034596307
42080371339432775981989132698526831260664840887571331401331362333709431244
0663659803352061415560955398316253892220738945585450197206138869521568000
"""
