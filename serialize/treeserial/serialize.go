package treeserial

import (
	"strconv"
	"strings"
)

//TreeNode is a Node object
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Codec is a serialize, deserialize class
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serialize a tree to a single string.
func (serial *Codec) Serialize(root *TreeNode) string {
	if root == nil {
		return "X,"
	}
	left := serial.Serialize(root.Left)
	right := serial.Serialize(root.Right)

	return strconv.Itoa(root.Val) + "," + left + right
}

// Deserialize your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	return this.deserializeHelper(&nodes)
}

func (this *Codec) deserializeHelper(nodes *[]string) *TreeNode {
	nodeVal := "X"
	if len(*nodes) > 0 {
		nodeVal = (*nodes)[0]
		*nodes = (*nodes)[1:]
	}
	if nodeVal == "X" {
		return nil
	}
	val, _ := strconv.Atoi(nodeVal)
	root := TreeNode{Val: val}
	root.Left = this.deserializeHelper(nodes)
	root.Right = this.deserializeHelper(nodes)

	return &root

}

// func testAbc(t *testing.T) {
// 	a := new()
// 	a.privateFunc()
// }

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// def serialize(self, root):
// 	""" Encodes a tree to a single string.
// 	:type root: TreeNode
// 	:rtype: str
// 	"""
// 	def rserialize(root, string):
// 		""" a recursive helper function for the serialize() function."""
// 		# check base case
// 		if root is None:
// 			string += 'None,'
// 		else:
// 			string += str(root.val) + ','
// 			string = rserialize(root.left, string)
// 			string = rserialize(root.right, string)
// 		return string

// 	return rserialize(root, '')

// def deserialize(self, data):
// 	"""Decodes your encoded data to tree.
// 	:type data: str
// 	:rtype: TreeNode
// 	"""
// 	def rdeserialize(l):
// 		""" a recursive helper function for deserialization."""
// 		if l[0] == 'None':
// 			l.pop(0)
// 			return None

// 		root = TreeNode(l[0])
// 		l.pop(0)
// 		root.left = rdeserialize(l)
// 		root.right = rdeserialize(l)
// 		return root

// 	data_list = data.split(',')
// 	root = rdeserialize(data_list)
// 	return root
