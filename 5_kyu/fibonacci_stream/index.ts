export function* fibonacciSequenceGen(): Iterator<number> {
    let [a, b] = [0, 1]
    while (true) {
        [a, b] = [b, a + b];
        yield a;
    }
}

export function fibonacciSequence(): Iterator<number> {
    let [a, b] = [0, 1]
    return {
        next: function () {
            [a, b] = [b, a + b];
            return {
                value: a,
                done: false
            }
        }
    }
}
