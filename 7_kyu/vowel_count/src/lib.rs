fn get_count(string: &str) -> usize {
    let mut vowels_count: usize = 0;
    let vowels = vec!['a', 'e', 'i', 'o', 'u'];
    for c in string.chars() {
        if vowels.contains(&c) {
            vowels_count += 1;
        }
    }
    vowels_count
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn my_tests() {
        assert_eq!(get_count("abracadabra"), 5);
    }
}
