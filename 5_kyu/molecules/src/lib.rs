use std::collections::HashMap;
use thiserror::Error;

#[derive(Error, Debug)]
pub enum ParseError {
    #[error("Mismatched parentheses")]
    MismatchedParentheses,
    #[error("Invalid character: {0}")]
    InvalidCharacter(char),
    #[error("Incomplete element")]
    IncompleteElement,
}

pub type Atom = (String, usize);
pub type Molecule = Vec<Atom>;

pub fn parse_molecule(s: &str) -> Result<Molecule, ParseError> {
    let mut stack: Vec<HashMap<String, usize>> = vec![HashMap::new()];
    let mut element = String::new();
    let mut count = String::new();
    let mut i = 0;
    let n = s.len();
    while i < n {
        let ch = s.chars().nth(i).unwrap();
        if ch.is_uppercase() {
            if !element.is_empty() {
                *stack
                    .last_mut()
                    .unwrap()
                    .entry(element.clone())
                    .or_insert(0) += count.parse::<usize>().unwrap_or(1);
                element.clear();
                count.clear();
            }
            element.push(ch);
            i += 1;
            while i < n && s.chars().nth(i).unwrap().is_lowercase() {
                element.push(s.chars().nth(i).unwrap());
                i += 1;
            }
        } else if ch.is_digit(10) {
            count.push(ch);
            i += 1;
            while i < n && s.chars().nth(i).unwrap().is_digit(10) {
                count.push(s.chars().nth(i).unwrap());
                i += 1;
            }
        } else if ch == '(' || ch == '[' || ch == '{' {
            let last = stack.last().unwrap().clone();
            stack.push(last);
            i += 1;
        } else if ch == ')' || ch == ']' || ch == '}' {
            if let Some(opening_ch) = stack.pop().map(|x| x.clone()) {
                let mut multiplier = String::new();
                i += 1;
                while i < n && s.chars().nth(i).unwrap().is_digit(10) {
                    multiplier.push(s.chars().nth(i).unwrap());
                    i += 1;
                }
                let multiplier = if multiplier.is_empty() {
                    1
                } else {
                    multiplier.parse::<usize>().unwrap_or(1)
                };
                for (element, count) in opening_ch {
                    *stack.last_mut().unwrap().entry(element).or_insert(0) += count * multiplier;
                }
            } else {
                return Err(ParseError::MismatchedParentheses);
            }
        } else {
            return Err(ParseError::InvalidCharacter(ch));
        }
    }
    if !element.is_empty() {
        *stack
            .last_mut()
            .unwrap()
            .entry(element.clone())
            .or_insert(0) += count.parse::<usize>().unwrap_or(1);
    }
    if stack.len() != 1 {
        Err(ParseError::MismatchedParentheses)
    } else {
        Ok(stack.last().unwrap().clone().into_iter().collect())
    }
}

#[cfg(test)]
mod tests {
    use super::{parse_molecule, Molecule};

    macro_rules! assert_parse {
        ($formula:expr, $expected:expr, $name:ident) => {
            #[test]
            fn $name() {
                super::assert_parse($formula, &$expected, "");
            }
        };
    }

    mod molecules {
        assert_parse!("H", [("H", 1)], hydrogen);
        assert_parse!("O2", [("O", 2)], oxygen);
        assert_parse!("H2O", [("H", 2), ("O", 1)], water);
        assert_parse!(
            "Mg(OH)2",
            [("Mg", 1), ("O", 2), ("H", 2)],
            magnesium_hydroxide
        );
        assert_parse!(
            "K4[ON(SO3)2]2",
            [("K", 4), ("O", 14), ("N", 2), ("S", 4)],
            fremys_salt
        );
    }

    #[test]
    fn errors() {
        assert_fail("pie", "Not a valid molecule");
        assert_fail("Mg(OH", "Mismatched parenthesis");
        assert_fail("Mg(OH}2", "Mismatched parenthesis");
    }

    fn assert_fail(formula: &str, msg: &str) {
        let result = parse_molecule(formula);
        assert!(
            result.is_err(),
            "expected {} {:?} to fail, got {:?}",
            msg,
            formula,
            result.unwrap()
        );
    }

    fn assert_parse(formula: &str, expected: &[(&str, usize)], _mst: &str) {
        let mut expected = expected
            .into_iter()
            .map(|&(name, usize)| (name.to_owned(), usize))
            .collect::<Molecule>();
        let result = parse_molecule(formula);
        assert!(
            result.is_ok(),
            "expected {:?} to pass, got {:?}",
            formula,
            result
        );
        let mut actual = result.unwrap();
        actual.sort();
        expected.sort();
        assert_eq!(actual, expected);
    }
}
