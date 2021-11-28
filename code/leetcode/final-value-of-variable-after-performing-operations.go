func finalValueAfterOperations(operations []string) int {
    out := 0

    for _, ops := range operations {
        switch ops[1] {
        case '+':
            out++
        case '-':
            out--
        }
    }

    return out
}
