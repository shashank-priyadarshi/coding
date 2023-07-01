# Longest Substring Without Non-repeating Characters

## Approach 1

- If length of string is == 0, return 0
- Iterate over the string
- Maintain a map of unique strings, and length
- If duplicate element found, empty map
- If length is greater than length of longest substring, swap
- Continue till end
