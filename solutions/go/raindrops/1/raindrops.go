package raindrops

import "fmt"

func isDivisibleByNum(number, num int) bool {
    return number%num == 0
}

func Convert(number int) string {
    var result string

    if isDivisibleByNum(number, 3) {
        result += "Pling"
    }

    if isDivisibleByNum(number, 5) {
        result += "Plang"
    }

    if isDivisibleByNum(number, 7) {
        result += "Plong"
    }

    if number%3 != 0 && number%5 != 0 && number%7 != 0 {
        return fmt.Sprintf("%d", number)
    }

    return result
}
