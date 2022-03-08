// https://leetcode.com/problems/subtree-of-another-tree/
// S가 T를 포함하고 있는지 확인
// -> s의 어떤 Node에서 시작하는 tree가 T와 같다면 True
package subtree

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	// s의 모든 노드를 돌면서 compareTree를 호출해서 true가 반환되면 true를 반환
	// BRS, DFS
	return DFSFunc(s, t, compareTree)

}

func DFSFunc(s, t *TreeNode, f func(*TreeNode, *TreeNode) bool) bool {
	if s == nil {
		if t == nil {
			return true
		}
		return false
	}
	if f(s, t) == true {
		return true
	}
	if DFSFunc(s.Left, t, f) == true {
		return true
	}
	return DFSFunc(s.Right, t, f)
}

// 두개의 트리가 완전히 같은지 확인
func compareTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		if t2 == nil {
			return true
		}
		return false
	} else if t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	if !compareTree(t1.Left, t2.Left) {
		return false
	}
	return compareTree(t1.Right, t2.Right)
}
