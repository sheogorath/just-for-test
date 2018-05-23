package main

import "fmt"
import "strings"
import "strconv"

func leap_year(year int) bool {
    if year % 400 == 0 {
        return true
    } else if year % 100 == 0 {
        return false
    } else if year % 4 == 0 {
        return true
    } else {
        return false
    }
}

func dow(day, month, year int) string {
    DOWS := [7]string{"Morndas", "Tirdas", "Middas", "Turdas",
                      "Fredas", "Loredas", "Sundas"}
    MONTHS := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

    dow := 0
    for y := 1996; y < year; y++ {
        dow += 365
        if leap_year(y) {
            dow += 1
        }
        dow %= 7
    }

    for m := 0; m < month - 1; m++ {
        dow += MONTHS[m]
        if m == 1 && leap_year(year) {
            dow += 1
        }
        dow %= 7
    }

    dow += day - 1
    dow %= 7

    return DOWS[dow]
}

func str2date(str string) (int, int, int) {
    var date [3] int

    nums := strings.Split(str, ".")

    for i:= 0; i < 3; i++ {
        num, err := strconv.Atoi(nums[i])
        if err != nil {
            fmt.Println("Something strange happened :(")
            break
        }
        date[i] = num
    }

    return date[0], date[1], date[2]
}

func main() {
    var input string

    fmt.Scanln(&input)

    fmt.Println(dow(str2date(input)))
}
