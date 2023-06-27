fn main() {
    println!("Hello, world!")
}

fn longestCommonPrefix(words: &[&str])-> String{
    words.join(",")
}

fn commonChars<'a>(word1: &'a str, word2: &'a str)-> &'a str{
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