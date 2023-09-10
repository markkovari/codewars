type Point = (f32, f32);

fn point_in_poly(poly: &[Point], point: Point) -> bool {
    let (p_x, p_y) = point;

    let mut inside = false;
    let mut j = poly.len() - 1;
    for i in 0..poly.len() {
        let (i_x, i_y) = poly[i];
        let (j_x, j_y) = poly[j];
        let intersect =
            (i_y > p_y) != (j_y > p_y) && p_x < (j_x - i_x) * (p_y - i_y) / (j_y - i_y) + i_x;
        if intersect {
            inside = !inside;
        }
        j = i;
    }
    inside
}

#[cfg(test)]
mod tests {
    use super::{point_in_poly, Point};

    #[test]
    fn simple_square() {
        let poly = [(-5.0, -5.0), (5.0, -5.0), (5.0, 5.0), (-5.0, 5.0)];

        show_and_test(&poly, (-6.0, 0.0), false);
        show_and_test(&poly, (-1.0, 1.0), true);
    }

    fn show_and_test(poly: &[Point], point: Point, expected: bool) {
        draw_test(poly, point, expected);

        assert_eq!(point_in_poly(poly, point), expected);
    }

    fn draw_test(poly: &[Point], point: Point, inside: bool) {
        fn transform((x, y): Point) -> Point {
            ((x + 7.) * 10. + 0.5, (y + 7.) * 10. + 0.5)
        }

        let points = poly
            .iter()
            .map(|&p| {
                let p = transform(p);

                format!("{},{}", p.0, p.1)
            })
            .collect::<Vec<String>>()
            .join(" ");

        let (cx, cy) = transform(point);

        println!(
            "\
            <div style='background:white; width:140px;'>\
                <svg width='140' height='140'>\
                    <polygon points='{points}' stroke='blue' fill='white'></polygon>\
                    <circle cx='{cx}' cy='{cy}' r='2' fill='{color}'></circle>\
                </svg>\
            </div>",
            points = points,
            cx = cx,
            cy = cy,
            color = if inside { "green" } else { "red" }
        );
    }
}
