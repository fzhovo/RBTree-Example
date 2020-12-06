package RBTree

import (
    "fmt"
    "log"
    "strconv"
    "strings"
)

type BST struct {
    root *Node
    nil  *Node
}

func MakeBST() *BST {
    n := &Node{}
    return &BST{root: n, nil: n}
}

// 插入结点
func (t *BST) Insert(z *Node) {
    // y指向x的父结点
    x, y := t.root, t.nil
    // 向下寻找插入位置
    // 直到x为t.nil，y为叶子结点
    for x != t.nil {
        y = x
        if z.key < x.key {
            x = x.left
        } else {
            x = x.right
        }
    }
    z.parent = y
    if y == t.nil {
        // 原先树为空，新插入结点作为根结点
        t.root = z
    } else if z.key < y.key {
        y.left = z
    } else {
        y.right = z
    }
    z.left, z.right = t.nil, t.nil
}

// 用一棵以v为根的子树来替换一棵以u为根的子树并成为其双亲的孩子结点
func (t *BST) Transplant(u, v *Node) {
    if u.parent == t.nil {
        t.root = v
    } else if u == u.parent.left {
        u.parent.left = v
    } else {
        u.parent.right = v
    }
    if v != t.nil {
        v.parent = u.parent
    }
}

// 删除结点
func (t *BST) Delete(z *Node) {
    if z.left == t.nil {
        t.Transplant(z, z.right)
    } else if z.right == t.nil {
        t.Transplant(z, z.left)
    } else {
        y := t.Minimum(z.right)
        if y.parent != z {
            t.Transplant(y, y.right)
            y.right = z.right
            y.right.parent = y
        }
        t.Transplant(z, y)
        y.left = z.left
        y.left.parent = y
    }
}

// 根据Key创建并插入一个结点
func (t *BST) InsertKey(k int) {
    t.Insert(&Node{key: k, parent: t.nil, left: t.nil, right: t.nil})
}

func (t *BST) Search(node *Node, k int) *Node {
    x := node
    for x != t.nil && k != x.key {
        if k < x.key {
            x = x.left
        } else {
            x = x.right
        }
    }
    return x
}

// 查找指定Key的结点
func (t *BST) SearchKey(k int) *Node {
    return t.Search(t.root, k)
}

// 找到当前子树的最小结点
func (t *BST) Minimum(node *Node) *Node {
    x := node
    for x.left != t.nil {
        x = x.left
    }
    return x
}

// 找到当前子树的最大结点
func (t *BST) Maximum(node *Node) *Node {
    x := node
    for x.right != t.nil {
        x = x.right
    }
    return x
}

// 打印中序遍历结果
func (t *BST) InorderPrint() {
    fmt.Println("中序遍历结果：")
    t.inorderPrintNode(t.root, false)
    fmt.Printf("\n")
}

func (t *BST) inorderPrintNode(node *Node, withColor bool) {
    if node == t.nil {
        return
    }
    t.inorderPrintNode(node.left, withColor)
    if !withColor {
        fmt.Printf("%d ", node.key)
    } else {
        fmt.Printf("%d(%s) ", node.key, node.GetColorName())
    }
    t.inorderPrintNode(node.right, withColor)
}

// 根据序列创建树
func MakeTreeBySequence(seq string) *BST {
    makeNodeByStr := func(keyStr string, parent *Node, t *BST) *Node {
        if keyStr == "null" {
            return t.nil
        } else {
            if key, err := strconv.Atoi(keyStr); err != nil {
                log.Fatal("无效的数字")
                return nil
            } else {
                return &Node{
                    key:    key,
                    left:   t.nil,
                    right:  t.nil,
                    parent: parent,
                    color:  0,
                }
            }
        }
    }
    s := strings.Split(seq, ",")
    var q []*Node
    if len(s) == 0 {
        return nil
    }
    t := MakeBST()
    // 创建根节点
    t.root = makeNodeByStr(s[0], t.nil, t)
    if t.root != t.nil {
        q = append(q, t.root)
    } else {
        return nil
    }
    n := len(s)
    for i := 1; i < n; i++ {
        // 出队
        node := q[0]
        q = q[1:]
        // 添加左孩子
        node.left = makeNodeByStr(s[i], node, t)
        if node.left != t.nil {
            q = append(q, node.left)
        }
        // 添加右孩子
        i++
        if i >= n {
            break
        }
        node.right = makeNodeByStr(s[i], node, t)
        if node.right != t.nil {
            q = append(q, node.right)
        }
    }
    return t
}


func (t *BST) validateRecursive(node *Node) (isValid bool, maxKey, minKey int) {
    if node.left == t.nil && node.right == t.nil {
        return true, node.key, node.key
    }
    key := node.key
    lIsValid, lMaxKey, lMinKey := t.validateRecursive(node.left)
    rIsValid, rMaxKey, rMinKey := t.validateRecursive(node.right)
    isValid = lIsValid && rIsValid && lMaxKey <= key && key <= rMinKey
    maxKey = max(max(lMaxKey, rMaxKey), key)
    minKey = min(min(lMinKey, rMinKey), key)
    return
}

// 验证是否是一个有效的二叉搜索树
func (t *BST) validate() (isValid bool) {
    isValid, _, _ = t.validateRecursive(t.root)
    return
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}
