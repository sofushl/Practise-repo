pub fn insertion_sort(mut x: Vec<i32>) {
    let mut o: Vec<i32> = vec![];
    o.push(x.pop().unwrap());

    'm: while let Some(y) = x.pop() {
        println!("{:?}", x);
        println!("{}", y);
        println!("{:?}", o);
        if o.is_empty() {
            o.push(y)
        } else {
            for i in 0..o.len() {
                if y <= o[i] {
                    o.insert(i, y);
                    continue 'm;
                }
            }
            o.push(y);
        }
    }

    println!("{:?}", o);
}

fn main() {
    insertion_sort(vec![23, 42, 12, 35, 1, 15, 7, 9, 49, 2, 2]);
}
