use num::bigint::{BigUint, ToBigUint};

fn highest_bi_prime_factor(pp1: u32, pp2: u32, n: BigUint) -> (BigUint, u32, u32) {

    (2.to_biguint().unwrap(), 1, 1)
}

// Add your tests here.
// See https://doc.rust-lang.org/stable/rust-by-example/testing/unit_testing.html

#[cfg(test)]
mod tests {
    use super::highest_bi_prime_factor;
    use num::bigint::{BigUint, ToBigUint};

    fn dotest(p1: u32, p2: u32, n: BigUint, expected: (BigUint, u32, u32)) {
        let actual = highest_bi_prime_factor(p1, p2, n.clone());
        assert!(
            actual == expected,
            "With p1 = {p1}, p2 = {p2}, n = {n}\nExpected {expected:?} but got {actual:?}"
        )
    }

    #[test]
    fn sample_tests() {
        dotest(
            2,
            3,
            50.to_biguint().unwrap(),
            (48.to_biguint().unwrap(), 4, 1),
        );
        dotest(
            5,
            11,
            1000.to_biguint().unwrap(),
            (605.to_biguint().unwrap(), 1, 2),
        );
        dotest(
            3,
            13,
            5000.to_biguint().unwrap(),
            (4563.to_biguint().unwrap(), 3, 2),
        );
        dotest(
            5,
            7,
            5000.to_biguint().unwrap(),
            (4375.to_biguint().unwrap(), 4, 1),
        );
    }
}
