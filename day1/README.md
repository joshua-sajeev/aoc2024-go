# Day 1

## Example Input

```text
3   4
4   3
2   5
1   3
3   9
3   3
```
## Part 1

The first part was to find the sum of  absolute difference between
the smallest number in left and smallest number in right with
second smallest number in left and second smallest number in right and so on.

### My Approach - Part 1

First I tried to get left and right inputs using line index with line\[0\] and line\[4\].
It worked for the example but not the actual input as it contained 5 digit inputs.
So with help of chatgpt got the correct way to process the lines

After that got the difference found its absolute value(Again had to ask chatgpt) and returned the sum.

## Part 2

For the second part the answer was the sum of product of number in the left and
it's frequency in the right column.

### My Approach - Part 2

What i did was to loop through the left array
if found same number in right array incremented the counter.
The problem with this is it will take O(n×m) (I know, need to improve)

Later found out i could have used a map to find the frequency and
make it O(n) + O(m)(IDK how to calculate chatgpt told me)

```go

func part2() int {
 rightCount := make(map[int]int)
 for _, num := range right {
  rightCount[num]++
 }

 sum := 0
 for _, num := range left {
  sum += num * rightCount[num]
 }
 return sum
}
```

```text
I think I should write test functions and better function names!
```
