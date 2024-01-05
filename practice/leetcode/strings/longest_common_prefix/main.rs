fn main() {
    println!("{}", longest_common_prefix(vec!["abaa".to_string(),"abc".to_string(),"abca".to_string()]));
}

fn longest_common_prefix(words: Vec<String>)->String{
    if words[0].len()==0{
        return "".to_string();
    }
    let first_char = words[0].chars().nth(0).unwrap().to_string();
    let mut common_char: String = "".to_string();
    for (key, value) in words.iter().enumerate(){
        if value.len() == 0{
            return "".to_string();
        }
        if value.chars().nth(0).unwrap().to_string() != first_char{
            return "".to_string();
        }
        if key == 0{
            common_char = value.to_string();
            continue;
        }

        // let mut temp_common_char:&str="";
        // let word1=value;
        // let word2 = &words[key-1];
        // if word1 == word2{
        //     temp_common_char=word1;
        // } else {
        // let len1 = word1.len();
        // let len2 = word2.len();
        // if len1 == 0 || len2 == 0{
        //     temp_common_char = "";
        // } else {
        // let length = if len1 > len2 {len2} else {len1};
        // for num in 1..=length{
        //     if word1[..num] != word2[..num]{
        //         temp_common_char = &word1[..num-1];
        //         break;
        //     } else {
        // temp_common_char=&word1[..length];
        //     }
        // }
        // }
        // }
        let temp_common_char = common_chars(value, &words[key-1]);
        if temp_common_char.len() == 0{
            return "".to_string();
        }
        if temp_common_char.len() < common_char.len(){
            common_char = temp_common_char.to_string();
        }
    }
    common_char
}

fn common_chars<'a>(word1: &'a str, word2: &'a str)-> &'a str{
    if word1 == word2{
        return word1;
    }
    let len1 = word1.len();
    let len2 = word2.len();
    if len1 == 0 || len2 == 0{
        return "";
    }
    let length = if len1 > len2 {len2} else {len1};
    for num in 1..=length{
        if word1[..num] != word2[..num]{
            return &word1[..num-1];
        }
        // for (num, (char1, char2)) in word1.chars().zip(word2.chars()).enumerate() {
        //     if char1 != char2 {
        //         return &word1[..num];
        //     }
        // }
    }
    &word1[..length]
}