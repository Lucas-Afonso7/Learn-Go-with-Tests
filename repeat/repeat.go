package integers

const quantityRepetitions = 5

func Repeat(character string) string {
    var repetitions string
    for i := 0; i < quantityRepetitions; i++ {
        repetitions += character
    }
    return repetitions
}