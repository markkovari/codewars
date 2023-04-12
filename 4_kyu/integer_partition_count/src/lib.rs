fn partitions(n: u32) -> u32 {
    let mut p = vec![0; (n + 1) as usize];

    p[0] = 1;

    for i in 1..=n {
        for j in i..=n {
            p[j as usize] += p[(j - i) as usize];
        }
    }

    p[n as usize]
}

#[cfg(test)]
mod tests {
    use super::partitions;

    #[test]
    fn basic_test_01() {
        assert_eq!(partitions(1), 1);
    }

    #[test]
    fn basic_test_05() {
        assert_eq!(partitions(5), 7);
    }

    #[test]
    fn basic_test_10() {
        assert_eq!(partitions(10), 42);
    }

    #[test]
    fn basic_test_25() {
        assert_eq!(partitions(25), 1958);
    }
}
