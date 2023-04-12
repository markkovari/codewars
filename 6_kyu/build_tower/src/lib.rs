pub fn add(left: usize, right: usize) -> usize {
    left + right
}

fn tower_builder(n_floors: usize) -> Vec<String> {
    let mut tower = Vec::new();
    for i in 1..=n_floors {
        let mut floor = String::new();
        for _ in 0..n_floors - i {
            floor.push(' ');
        }
        for _ in 0..2 * i - 1 {
            floor.push('*');
        }
        for _ in 0..n_floors - i {
            floor.push(' ');
        }
        tower.push(floor);
    }
    tower
}

#[cfg(test)]
mod tests {
    use super::tower_builder;

    #[test]
    fn fixed_tests() {
        assert_eq!(tower_builder(1), vec!["*"]);
        assert_eq!(tower_builder(2), vec![" * ", "***"]);
        assert_eq!(tower_builder(3), vec!["  *  ", " *** ", "*****"]);
    }
}
