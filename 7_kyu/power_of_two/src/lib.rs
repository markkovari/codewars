pub fn add(left: usize, right: usize) -> usize {
    left + right
}

fn power_of_two(x: u64) -> bool {
    let as_bintary = format!("{:b}", x);
    as_bintary.matches("1").count() == 1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = power_of_two(2);
        assert!(result);
    }

    #[test]
    fn it_works_cool() {
        let result = power_of_two(3);
        assert!(!result);
    }
}
