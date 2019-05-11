base = []
multiplier = []
threeDigits = []

def importBase():
    temp = None
    with open("base.txt", "r") as f:
        temp = [str(x).strip() for x in f.readlines()]

    global base 
    base = temp

def importMult():
    temp = None
    with open("mult.txt", "r") as f:
        temp = [str(x).strip() for x in f.readlines()]

    global multiplier
    multiplier = temp

def buildThreeDigits():
    for i in range(1, int(1e4)):
        threeDigits.append(parseThreeDigit(i))

def parseThreeDigit(n):
    if n <= len(base):
        return base[n-1]
    else:
        temp = []
        it = 0
        while n:
            temp.append(n%10)
            n //= 10

        temp = temp[::-1]

        ans = []

        for idx, number in enumerate(temp):
            if number == 0:
                continue

            if idx == len(temp) - 2:
                tnumber = "{}{}".format(number, temp[idx+1])
                tnumber = int(tnumber)
                if tnumber <= len(base):
                    ans.append(base[tnumber-1])
                    break

            if number == 1 and idx != len(temp) - 1:
                ans.append("se{}".format(multiplier[len(temp) - (idx+1) - 1]))
                continue

            ans.append(base[number-1])
            if idx < len(temp) - 1:
                ans.append(multiplier[len(temp) - (idx+1) - 1])

        return " ".join(map(str, ans))

def toWords(n):
    if n <= len(threeDigits):
        return threeDigits[n-1]
    else:
        tn = str(n)[::-1]
        length = 3
        temp = [tn[0+i:length+i][::-1] for i in range(0, len(tn), length)][::-1]

        # print(temp)
        ans = []

        for idx, frac in enumerate(temp):
            frac = int(frac)

            if frac == 0:
                continue

            tidx = len(temp) - idx - 1

            if 2 <= tidx <= 3:
               ans.append("{} {}".format(threeDigits[frac-1], multiplier[tidx]))
            elif tidx == 1:
                ans.append("{} {}".format(threeDigits[frac-1], multiplier[tidx+1]))
            else:
                ans.append("{}".format(threeDigits[frac-1]))

        return " ".join(map(str, ans))


        # else:
        #     print("Sini")
        #     tnumber = "{}{}".format(number, temp[idx+1])
        #     tnumber = int(tnumber)
        #     if idx == len(temp) - 2:
        #         if tnumber <= len(base):
        #             ans.append(base[tnumber-1])
        #             break


        #         if tnumber <= len(base) and tidx > 3:
        #             ans.append(base[tnumber - 1])
        #             if tidx < 6:
        #                 ans.append(mult[3])
        #             elif tidx < 9:
        #                 ans.append(mult[4])


        # ans.append(base[number-1])
        # if idx < len(temp) - 1:
        #     ans.append(multiplier[len(temp) - (idx+1) - 1])

    # return " ".join(map(str, ans))


if __name__ == "__main__":
    importBase()
    importMult()
    buildThreeDigits()

    # print(threeDigits)
    # for x in threeDigits:
        # print(x)

    n = int(input())
    # print(toWords(n))

    for i in range(1, n+1):
        print(toWords(i))
