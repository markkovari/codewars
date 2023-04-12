fn rotate(s: &str) -> Vec<String> {
    let mut result = Vec::new();
    let mut s = s.to_string();
    for _ in 0..s.len() {
        s = s.chars().skip(1).chain(s.chars().take(1)).collect();
        result.push(s.clone());
    }
    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_rotate_empty() {
        assert_eq!(
            rotate(""),
            Vec::<String>::new(),
            "Should return empty string for input: {:?}",
            ""
        );
    }

    #[test]
    fn test_rotate_single() {
        assert_eq!(
            rotate("a"),
            vec!["a"],
            "\n\nYour result (left) did not match the expected output (right) for the input: {:?}",
            "a"
        );
    }

    #[test]
    fn test_rotate_triple() {
        assert_eq!(
            rotate("abc"),
            vec!["bca", "cab", "abc"],
            "\n\nYour result (left) did not match the expected output (right) for the input: {:?}",
            "abc"
        );
    }

    #[test]
    fn test_rotate_long() {
        assert_eq!(
            rotate("hello"),
            vec!["elloh", "llohe", "lohel", "ohell", "hello"],
            "\n\nYour result (left) did not match the expected output (right) for the input: {:?}",
            "hello"
        );
    }
}
