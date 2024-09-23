# 1249 - Minimum Remove to Make Valid Parentheses
Given a string `s` of `'('`,`')'` and lowercase English characters.

Your task is to remove the minimum number of parentheses (`'('` or `')'`, in any positions) so that the resulting parentheses string is valid and return any valid string.

Formally, a parentheses string is valid if an only if:
- It is the empty string, contains only lowercase characters, or
- It can be written as `AB`(`A` concatenated with `B`), where `A` and `B` are valid strings, or
- It can be written as `(A)`, where `A` is a valid string.

### Example 1:
```
Input: s = "lee(t(c)o)de)"
Output: "lee(t(c)o)de"
Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
```

### Example 2:
```
Input: s = "a)b(c)d"
Output: "ab(c)d"
```

### Example 3:
```
Input: s = "))(("
Output: ""
Explanation: An empty string is also valid.
```

### Constraints:
- <code>1 <= s.length <= 10<sup>5</sup></code>
- `s[i]` is either `'('`,`')'`, or a lowercase English letter.


## Explanation
1. Problem Restatement:
The problem is to remove the minimum number of parentheses from a given string such that the resulting string is valid. A string is valid if every opening parentheses has a corresponding closing parenthesis in the correct order.

2. Clarify the Approach:
The approach I'm taking is a two-pass solution. In the first pass, I'll identify mismatched parentheses--either excess closing parentheses or opening parentheses that don't have a match. In the second pass, I'll remove those mismatches to create a valid string. This method ensures that we only remove the minimum number of parentheses.

3. Step-by-Step:
First pass - Identify Mismatches:
- I'll iterate through the string from left to right using a stack to track the indices of unmatched opening parentheses.  When I encounter a closing parenthesis, I check if there's a corresponding opening parenthesis on the stack.
  - If there is, I pop from the stack
  - If there isn't it means this closing parenthesis is unmatched and must be removed, so I mark it for removal by setting it's position in the array to be a special value, such as zero.

Second Pass - Removing Mismatches
- Once I've processed all the characters, there may still be unmatched opening parentheses left in the stack. These must be removed, so I go through the stack and mark them for removal as well.

Rebuilding the Valid String:
Finally, I iterate through the character array, skipping over any characters that were marked for removal, and rebuild the string.

4. Justify the Correctness:
- **Matching Parentheses:** "By using a stack, I ensure that all matched parentheses are processed correctly in a Last-In-First-Out order. Whenever there’s an unmatched closing parenthesis, it’s marked immediately for removal, and unmatched opening parentheses are handled after the full scan."
- **Minimum Removals:** "Since I'm only marking mismatched parentheses, the algorithm guarantees that no valid parentheses are removed. Therefore, this approach ensures that the minimum number of parentheses are removed to make the string valid."

5. Time and Space Complexity Analysis
The algorithm processes the string in two passes. The first pass scans the string and pushes indices onto the stack or marks invalid characters. The second pass constructs the result string by filtering out invalid characters. Both passes take linear time, so the overall time complexity is O(n), where n is the length of the string.

We use an extra stack to store indices of unmatched opening parentheses, which, in the worst case, could contain all opening parentheses in the string. Additionally, we modify the string in place, so the space complexity is O(n) in the worst case due to the stack.

6. Edge Cases:
Some edge cases to consider include: strings that contain no parentheses, strings that are already valid, and strings with all invalid parentheses, like ))((". The algorithm handles all these cases because it can identify when no parentheses need to be removed or when the entire string needs to be cleared.
