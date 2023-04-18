fn list_position(word: &str) -> u128 {
    let mut letters = [0; 26];
    for b in word.bytes() {
        letters[(b - b'A') as usize] += 1;
    }
    let len = word.len();
    let mut fact = [1; 26];
    for i in 2..=len {
        fact[i] = fact[i - 1] * i as u128;
    }
    let mut res = 0;
    for (i, b) in word.bytes().enumerate() {
        let mut count = 0;
        for &c in &letters[..(b - b'A') as usize] {
            count += c;
        }
        res += count as u128 * fact[len - i - 1];
        letters[(b - b'A') as usize] -= 1;
    }
    res + 1
}

// Add your tests here.
// See https://doc.rust-lang.org/stable/rust-by-example/testing/unit_testing.html

#[cfg(test)]
mod tests {
    use super::list_position;
    
    const ERR_MSG: &str = "\nYour result (left) did not match the expected output (right)";
    
    #[test]
    fn sample_tests() {
        let test_data = [
            (                  "A", 1),
            (               "ABAB", 2),
            (               "AAAB", 1),
            (               "BAAA", 4),
            (               "YMYM", 5),
            (           "QUESTION", 24572),
            (         "BOOKKEEPER", 10743),
      ("IMMUNOELECTROPHORETICALLY", 718393983731145698173),
        ];
        for (word, expected) in test_data {
            assert_eq!(list_position(word), 
                       expected,
                       "\nYour result (left) did not match the expected output (right) for the input: \"{word}\"");
        }
        
    }
}