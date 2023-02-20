use itertools::Itertools;

struct Sudoku {
    data: Vec<Vec<u32>>,
}

fn get_column_at(sudoku: &Vec<Vec<u32>>, index: usize) -> Vec<u32> {
    let mut column = vec![];
    for row in sudoku.iter() {
        column.push(row[index]);
    }
    column
}

fn all_unique(elements: &Vec<u32>) -> bool {
    elements.iter().filter(|&e| e != &0u32).unique().count() == elements.len()
        && all_less_than(elements, elements.len() as u32)
}
fn all_less_than(elements: &Vec<u32>, compared_to: u32) -> bool {
    elements.iter().all(|&e| e <= compared_to)
}

fn get_part(
    sudoku: &Vec<Vec<u32>>,
    column_start: usize,
    column_end: usize,
    row_start: usize,
    row_end: usize,
) -> Vec<u32> {
    let mut part = vec![];
    for row in sudoku.iter().skip(row_start).take(row_end - row_start) {
        for element in row
            .iter()
            .skip(column_start)
            .take(column_end - column_start)
        {
            part.push(*element);
        }
    }
    part
}

impl Sudoku {
    fn is_valid(&self) -> bool {
        Self::validate_sudoku(&self.data)
    }

    fn validate_sudoku(sudoku: &Vec<Vec<u32>>) -> bool {
        let dimension = sudoku.len();
        let small_square_width = (dimension as f32).sqrt() as u32;
        let small_square_count = (dimension / small_square_width as usize) as u32;
        for row in sudoku.iter() {
            if row.len() != dimension {
                return false;
            }
            if !all_unique(row) {
                return false;
            }
        }
        for i in 0..sudoku.len() {
            let column = get_column_at(&sudoku, i);
            if column.len() != dimension {
                return false;
            }
            if !all_unique(&column) {
                return false;
            }
        }

        for vertical in 0..small_square_count {
            for horizontal in 0..small_square_count {
                let inner_cell = get_part(
                    sudoku,
                    (vertical * small_square_width) as usize,
                    (vertical * small_square_width + small_square_width) as usize,
                    (horizontal * small_square_width) as usize,
                    (horizontal * small_square_width + small_square_width) as usize,
                );
                if inner_cell.len() != dimension {
                    return false;
                }
                if !all_unique(&inner_cell) {
                    return false;
                }
            }
        }

        true
    }
}

#[test]
fn good_sudoku() {
    let good_sudoku_1 = Sudoku {
        data: vec![
            vec![7, 8, 4, 1, 5, 9, 3, 2, 6],
            vec![5, 3, 9, 6, 7, 2, 8, 4, 1],
            vec![6, 1, 2, 4, 3, 8, 7, 5, 9],
            vec![9, 2, 8, 7, 1, 5, 4, 6, 3],
            vec![3, 5, 7, 8, 4, 6, 1, 9, 2],
            vec![4, 6, 1, 9, 2, 3, 5, 8, 7],
            vec![8, 7, 6, 3, 9, 4, 2, 1, 5],
            vec![2, 4, 3, 5, 6, 1, 9, 7, 8],
            vec![1, 9, 5, 2, 8, 7, 6, 3, 4],
        ],
    };

    let good_sudoku_2 = Sudoku {
        data: vec![
            vec![1, 4, 2, 3],
            vec![3, 2, 4, 1],
            vec![4, 1, 3, 2],
            vec![2, 3, 1, 4],
        ],
    };
    assert!(good_sudoku_1.is_valid());
    assert!(good_sudoku_2.is_valid());
}

#[test]
fn bad_sudoku() {
    let bad_sudoku_1 = Sudoku {
        data: vec![
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
        ],
    };

    let bad_sudoku_2 = Sudoku {
        data: vec![
            vec![1, 2, 3, 4, 5],
            vec![1, 2, 3, 4],
            vec![1, 2, 3, 4],
            vec![1],
        ],
    };
    assert!(!bad_sudoku_1.is_valid());
    assert!(!bad_sudoku_2.is_valid());
}
