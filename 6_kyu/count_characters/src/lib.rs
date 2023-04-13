use std::collections::HashMap;

fn count(input: &str) -> HashMap<char, i32> {
    let mut frequency: HashMap<char, i32> = HashMap::new();
    for c in input.chars() {
        let count = frequency.entry(c).or_insert(0);
        *count += 1;
    }
    frequency
}

// Add your tests here.
// See https://doc.rust-lang.org/stable/rust-by-example/testing/unit_testing.html

#[cfg(test)]
mod tests {
    use super::*;

    const ERR_MSG: &str = "\nYour result (left) did not match the expected output (right)";

    #[test]
    fn test_empty_string() {
        let test_input = "";
        let expected: HashMap<char, i32> = HashMap::new();

        assert_eq!(
            count(test_input),
            expected,
            "{ERR_MSG} with input: \"{test_input}\""
        );
    }

    #[test]
    fn test_string_with_two_equal_letters() {
        let test_input = "aa";
        let mut expected: HashMap<char, i32> = HashMap::new();
        expected.insert('a', 2);

        assert_eq!(
            count(test_input),
            expected,
            "{ERR_MSG} with input: \"{test_input}\""
        );
    }

    #[test]
    fn test_string_with_different_letters() {
        let test_input = "aabb";
        let mut expected: HashMap<char, i32> = HashMap::new();
        expected.insert('a', 2);
        expected.insert('b', 2);

        assert_eq!(
            count(test_input),
            expected,
            "{ERR_MSG} with input: \"{test_input}\""
        );
    }
}
