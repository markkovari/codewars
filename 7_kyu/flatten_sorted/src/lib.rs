fn flatten_and_sort(arr: &[Vec<i32>]) -> Vec<i32> {
    let mut flattened: Vec<i32> = arr.iter().flatten().cloned().collect();
    flattened.sort();
    flattened
}


// Add your tests here.
// See https://doc.rust-lang.org/stable/rust-by-example/testing/unit_testing.html

#[cfg(test)]
mod tests {
    use super::flatten_and_sort;
        
    fn dotest(arr: &[Vec<i32>], expected: &[i32]) {
        let actual = flatten_and_sort(arr);
        assert!(actual == expected, 
            "With arr = {arr:?}\nExpected: {expected:?}\n But got: {actual:?}\n")
    }

    #[test]
    fn fixed_tests() {
        dotest(&[], &[]);
        dotest(&[vec![], vec![]], &[]);
        dotest(&[vec![], vec![1]], &[1]);
        dotest(&[vec![3, 2, 1], vec![7, 9, 8], vec![6, 4, 5]], &[1, 2, 3, 4, 5, 6, 7, 8, 9]);
        dotest(&[vec![1, 3, 5], vec![100], vec![2, 4, 6]], &[1, 2, 3, 4, 5, 6, 100]);
    }
}

