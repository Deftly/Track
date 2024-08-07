# 10 - Regular Expression Matching
Given an input string `s` and a pattern `p`, implement regular expression matching with support for `.` and `*` where:
- `.` Matches any single character.
- `*` Matches zero or more of the preceding element.
The matching should cover the **entire** input string(not partial).

**Example 1:**
```
Input: s = "aa", p = "a"
Output: false
```

**Example 2:**
```
Input: s = "aa", p = "a*"
Output: true
```

**Example 3:**
```
Input: s = "ab", p = ".*"
Output: true
```

**Constraints:**
- `1 <= s.length <= 20`
- `1 <= p.length <= 20`
- `s` contains only lowercase English letters.
- `p` contains only lowercase English letters, `.`, and `*`.
- It is guaranteed for each appearance of the character `*`, there will be a previous valid character to match.
