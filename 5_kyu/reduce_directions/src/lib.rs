#[derive(Clone, Copy, Debug, PartialEq, Eq)]
enum Direction {
    North,
    East,
    West,
    South,
}

fn dir_reduc(arr: &[Direction]) -> Vec<Direction> {
    let mut result: Vec<Direction> = Vec::new();
    for d in arr {
        match d {
            Direction::North => {
                if let Some(Direction::South) = result.last() {
                    result.pop();
                } else {
                    result.push(*d);
                }
            }
            Direction::South => {
                if let Some(Direction::North) = result.last() {
                    result.pop();
                } else {
                    result.push(*d);
                }
            }
            Direction::East => {
                if let Some(Direction::West) = result.last() {
                    result.pop();
                } else {
                    result.push(*d);
                }
            }
            Direction::West => {
                if let Some(Direction::East) = result.last() {
                    result.pop();
                } else {
                    result.push(*d);
                }
            }
        }
    }
    result
}

#[cfg(test)]
mod tests {
    use super::{
        dir_reduc,
        Direction::{self, *},
    };

    #[test]
    fn basic() {
        let a = [North, South, South, East, West, North, West];
        assert_eq!(dir_reduc(&a), [West]);

        let a = [North, West, South, East];
        assert_eq!(dir_reduc(&a), [North, West, South, East]);
    }
}
