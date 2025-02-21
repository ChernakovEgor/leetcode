package contaminatedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type FindElements struct {
	seen map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	seen := make(map[int]bool)

	var traverse func(*TreeNode, int)
	traverse = func(curr *TreeNode, v int) {
		curr.Val = v
		seen[v] = true
		if curr.Left != nil {
			traverse(curr.Left, 2*v+1)
		}
		if curr.Right != nil {
			traverse(curr.Right, 2*v+2)
		}
	}

	traverse(root, 0)

	return FindElements{seen}
}

func (this *FindElements) Find(target int) bool {
	return this.seen[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
