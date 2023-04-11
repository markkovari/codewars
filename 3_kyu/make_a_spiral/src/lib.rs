fn check_wall(head: &[i8], direction: char, background: &Vec<Vec<i8>>, size: usize) -> bool {
    let mut current = head.to_vec();
    match direction {
        'E' => current[1] += 2,
        'W' => current[1] -= 2,
        'N' => current[0] -= 2,
        'S' => current[0] += 2,
        _ => (),
    }
    if current.contains(&(size as i8 + 2)) || current.contains(&(-1)) {
        return true;
    } else if background[current[0] as usize][current[1] as usize] == 1 {
        return true;
    }
    false
}

fn move_head(head: &mut [i8], direction: char) {
    /* Increments the head location */
    match direction {
        'E' => head[1] += 1,
        'W' => head[1] -= 1,
        'N' => head[0] -= 1,
        'S' => head[0] += 1,
        _ => (),
    }
}

fn spiralize(size: usize) -> Vec<Vec<i8>> {
    let mut background: Vec<Vec<i8>> = vec![vec![0; (size + 2) as usize]; (size + 2) as usize];
    let mut head = [1, 1];
    let directions = ['E', 'S', 'W', 'N'];
    let mut direction = directions[0];
    let mut moved_in_direction = 0;
    loop {
        background[head[0] as usize][head[1] as usize] = 1;
        let change_direction = check_wall(&head, direction, &background, size);
        if change_direction {
            direction = directions
                [(directions.iter().position(|&x| x == direction).unwrap() + 1) % directions.len()];
            if moved_in_direction == 1 || check_wall(&head, direction, &background, size) {
                break;
            }
            moved_in_direction = 0;
        }
        move_head(&mut head, direction);
        moved_in_direction += 1;
    }
    background
        .iter()
        .skip(1)
        .take(size as usize)
        .map(|row| row.iter().skip(1).take(size as usize).map(|x| *x).collect())
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test5() {
        assert_eq!(
            spiralize(5),
            [
                [1, 1, 1, 1, 1],
                [0, 0, 0, 0, 1],
                [1, 1, 1, 0, 1],
                [1, 0, 0, 0, 1],
                [1, 1, 1, 1, 1],
            ],
        );
    }

    #[test]
    fn test8() {
        assert_eq!(
            spiralize(8),
            [
                [1, 1, 1, 1, 1, 1, 1, 1],
                [0, 0, 0, 0, 0, 0, 0, 1],
                [1, 1, 1, 1, 1, 1, 0, 1],
                [1, 0, 0, 0, 0, 1, 0, 1],
                [1, 0, 1, 0, 0, 1, 0, 1],
                [1, 0, 1, 1, 1, 1, 0, 1],
                [1, 0, 0, 0, 0, 0, 0, 1],
                [1, 1, 1, 1, 1, 1, 1, 1],
            ],
        );
    }
}
