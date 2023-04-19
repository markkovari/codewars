fn quadrant(x: i32, y: i32) -> i32 {
    if x > 0 && y > 0 {
        1
    } else if x < 0 && y > 0 {
        2
    } else if x < 0 && y < 0 {
        3
    } else if x > 0 && y < 0 {
        4
    } else {
        0
    }
}

// Add your tests here.
// See https://doc.rust-lang.org/stable/rust-by-example/testing/unit_testing.html

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn fixed_tests() {
        assert_eq!(quadrant(1, 2), 1);
        assert_eq!(quadrant(3, 5), 1);
        assert_eq!(quadrant(-10, 100), 2);
        assert_eq!(quadrant(-1, -9), 3);
        assert_eq!(quadrant(19, -56), 4);    }
}

