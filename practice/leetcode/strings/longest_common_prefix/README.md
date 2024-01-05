# Longest Common Prefix

## Approach 1

- If length of first word is 0, return empty string
- Find first character of first word
- Iterate over words array
- If any word found with length 0, return empty string
- If word found with a different first character, return empty string
- If current iteration is zero, set common character prefix as the whole first word and continue
- Find common character prefix between current and last word and save in temporary variable
- If length of temporary variable is zero, return empty string
- If length of temporary variable is less than length of prefix, prefix=temporary variable
- continue till last element

## Approach 2

- Find first character of first word, if word found with a different first character, return nil
- Get string with smallest length in the array
-

<https://leetcode.com/problems/longest-common-prefix/>
