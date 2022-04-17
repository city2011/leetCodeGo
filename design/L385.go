package main

func main() {

}

func deserialize(s string) *NestedInteger {
	_, res := dfs(s, 0, len(s))
	return &res
}

func dfs(s string, idx int, n int) (end int, ni NestedInteger) {
	isDigit := false
	num := 0
	for i := idx; i < n; i++ {
		if s[i] == ']'{
			return idx, ni
		} else if s[i] == ','{
			if isDigit {
				ni.SetInteger(num)
				num = 0
			}
		} else if s[i] == '[' {
			if isDigit {
				ni.SetInteger(num)
				num = 0
			}
			isDigit = false
			end, b := dfs(s, i + 1, n)
			ni.Add(b)
			i = end
		} else {
			isDigit = true
			num = num * 10 + int(s[i])
		}
	}
	return
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []NestedInteger {
	return nil
}
