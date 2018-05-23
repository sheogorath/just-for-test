#!/usr/bin/python

def leap_year(year):
    if year % 400 == 0:
        return True
    elif year % 100 == 0:
        return False
    elif year % 4 == 0:
        return True
    else:
        return False


def dow(day, month, year):
    DOWS = ["Morndas", "Tirdas", "Middas", "Turdas", 
            "Fredas", "Loredas", "Sundas"]

    MONTHS = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31] 

    if year < 1996:
        return None

    dow = 0 
    for y in range(1996, year):
        dow += 365 
        if leap_year(y):
            dow += 1
        dow %= 7

    for m in range(month - 1): 
        dow += MONTHS[m]
        if m == 1 and leap_year(year):
            dow += 1
        dow %= 7

    dow += day - 1 
    dow %= 7

    return DOWS[dow]


if __name__ == "__main__":
    d, m, y = (int(i) for i in input().split('.'))
    print(dow(d, m, y))
    print("test")
