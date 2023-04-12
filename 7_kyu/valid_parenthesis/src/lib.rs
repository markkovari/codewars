fn valid_parentheses(parens: &str) -> bool {
    let mut stack: Vec<&str> = vec![];
    for c in parens.chars() {
        if c == '(' {
            stack.push("(");
        } else if c == ')' {
            if stack.len() == 0 {
                return false;
            }
            stack.pop();
        }
    }
    stack.len() == 0
}

// Add your tests here.
// See https://doc.rust-lang.org/stable/rust-by-example/testing/unit_testing.html

#[cfg(test)]
mod tests {
    use super::*;

    fn do_test(expected: bool, input: &str) {
        assert_eq!(valid_parentheses(input), expected, "\nYour result (left) did not match the expected output (right) for the input: {input:?}");
    }

    #[test]
    fn valid_cases() {
        do_test(true, "()");
        do_test(true, "((()))");
        do_test(true, "()()()");
        do_test(true, "(()())()");
        do_test(true, "()(())((()))(())()");
    }

    #[test]
    fn invalid_cases() {
        do_test(false, ")(");
        do_test(false, "()()(");
        do_test(false, "((())");
        do_test(false, "())(()");
        do_test(false, ")()");
        do_test(false, ")");
    }

    #[test]
    fn empty_string() {
        do_test(true, "");
    }
}
