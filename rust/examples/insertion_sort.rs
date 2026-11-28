pub fn insertion_sort(a: &mut [i32]) {
    for i in 1..a.len() {
        let key = a[i];
        let mut j = i;
        while j > 0 && a[j - 1] > key {
            a[j] = a[j - 1];
            j -= 1;
        }
        a[j] = key;
    }
}

fn main() {
    let mut a: Vec<i32> = vec![23, 42, 12, 35, 1, 15, 7, 9, 49, 2, 2];
    println!("{:?}", a);
    insertion_sort(&mut a);
    println!("{:?}", a);
}
