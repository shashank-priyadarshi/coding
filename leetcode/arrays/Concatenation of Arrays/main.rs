fn main(){
    let input_arr = [1, 2, 3];
    let length = input_arr.len();
    let mut ans_arr = vec!(0; 2 * length);
    let mut i = 0;
//     ans_arr = [1, 2, 3, 1, 2, 3]
//     1 : 0, 3 = 3
//     2 : 1, 4 = 3
//     3 : 2, 5 = 3
//     length = 3
    for _i in length..{
        ans_arr[i] = input_arr[i];
        ans_arr[i+length] = input_arr[i];
        i += 1;
    }
    println!("{:?}", ans_arr);
}
