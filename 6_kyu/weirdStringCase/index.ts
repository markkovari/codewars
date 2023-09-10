export function toWeirdCase(input: string) {
    return input.split(' ').map((word) => {
        return word.split('').map((letter, index) => {
            return index % 2 === 0 ? letter.toUpperCase() : letter.toLowerCase();
        }).join('');
    }).join(' ');
}