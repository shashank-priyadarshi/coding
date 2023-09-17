fn main() {
    let input = 50;
    let mut a = 0;
    let mut b = 1;
    let mut c = 0;
    println!("{}\n{}", a , b);
    for _c in input.. {
        c = a + b;
        if c > 50 {
            break;
        }
        a = b;
        b = c;
        println!("{}", c);
    }
}
