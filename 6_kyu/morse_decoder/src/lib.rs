use std::{collections::HashMap, result};

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

//mod preloaded;
//use preloaded::MORSE_CODE;
// MORSE_CODE is `HashMap<String, String>`. e.g. ".-" -> "A".

fn decode_morse(encoded: &str) -> String {
    let encoded = encoded.trim();
    let mut MORSE_CODE = HashMap::<String, String>::new();
    let mut result = String::new();
    for word in encoded.split("   ") {
        for letter in word.split(" ").into_iter() {
            if MORSE_CODE.contains_key(letter) {
                result += &MORSE_CODE[letter];
            }
        }
        result += " ";
    }
    result.trim_end().to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
