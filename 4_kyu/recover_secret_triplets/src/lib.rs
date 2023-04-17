fn recover_secret(triplets: Vec<[char; 3]>) -> String {
    let mut result = String::new();
    let mut triplets = triplets;
    let mut i = 0;
    while i < triplets.len() {
        let mut j = i + 1;
        while j < triplets.len() {
            if triplets[i][0] == triplets[j][0] {
                if triplets[i][1] == triplets[j][1] {
                    triplets.remove(j);
                } else if triplets[i][1] == triplets[j][2] {
                    triplets[j].swap(1, 2);
                } else if triplets[i][2] == triplets[j][1] {
                    triplets[i].swap(1, 2);
                } else if triplets[i][2] == triplets[j][2] {
                    triplets[i].swap(1, 2);
                    triplets[j].swap(1, 2);
                } else {
                    j += 1;
                }
            } else if triplets[i][0] == triplets[j][1] {
                if triplets[i][1] == triplets[j][0] {
                    triplets[i].swap(0, 1);
                } else if triplets[i][1] == triplets[j][2] {
                    triplets[i].swap(0, 1);
                    triplets[j].swap(1, 2);
                } else if triplets[i][2] == triplets[j][0] {
                    triplets[i].swap(0, 1);
                    triplets[i].swap(1, 2);
                } else if triplets[i][2] == triplets[j][2] {
                    triplets[i].swap(0, 1);
                    triplets[i].swap(1, 2);
                    triplets[j].swap(1, 2);
                } else {
                    j += 1;
                }
            } else if triplets[i][0] == triplets[j][2] {
                if triplets[i][1] == triplets[j][0] {
                    triplets[i].swap(0, 2);
                } else if triplets[i][1] == triplets[j][1] {
                    triplets[i].swap(0, 2);
                    triplets[j].swap(0, 1);
                } else if triplets[i][2] == triplets[j][0] {
                    triplets[i].swap(0, 2);
                    triplets[i].
[1, 2]




                                
}

// Rust test example:
// TODO: replace with your own tests (TDD), these are just how-to examples.
// See: https://doc.rust-lang.org/book/testing.html

#[test]
fn example_test() {
    assert_eq!(
        recover_secret(vec![
            ['t', 'u', 'p'],
            ['w', 'h', 'i'],
            ['t', 's', 'u'],
            ['a', 't', 's'],
            ['h', 'a', 'p'],
            ['t', 'i', 's'],
            ['w', 'h', 's']
        ]),
        "whatisup"
    );
}
