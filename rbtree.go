package RBTree

import "fmt"

type RBTree struct {
    BST
}

func MakeRBTree() *RBTree {
    bst := MakeBST()
    return &RBTree{*bst}
}

func (t *RBTree) LeftRotate(x *Node) {
    // y是x的右孩子，即左旋后的新根结点
    y := x.right
    // 将y的左孩子连到x的右边
    x.right = y.left
    if y.left != t.nil {
        y.left.parent = x
    }
    // 把y旋转到x的位置
    y.parent = x.parent
    if x.parent == t.nil {
        t.root = y
    } else if x == x.parent.left {
        x.parent.left = y
    } else {
        x.parent.right = y
    }
    y.left = x
    x.parent = y
}

func (t *RBTree) RightRotate(x *Node) {
    // y是x的左孩子，即右旋后的新根结点
    y := x.left
    x.left = y.right
    // 将y的右孩子连到x的左边
    if y.right != t.nil {
        y.right.parent = x
    }
    // 把y旋转到x的位置
    y.parent = x.parent
    if x.parent == t.nil {
        t.root = y
    } else if x == x.parent.right {
        x.parent.right = y
    } else {
        x.parent.left = y
    }
    y.right = x
    x.parent = y
}

func (t *RBTree) Insert(z *Node) {
    // BST的插入
    t.BST.Insert(z)
    // 新插入结点标为红色
    z.left, z.right, z.color = t.nil, t.nil, RED
    // 插入修复
    t.insertFixup(z)
}

// 插入时出现双红的修复操作
func (t *RBTree) insertFixup(z *Node) {
    for z.parent.color == RED {
        if z.parent == z.parent.parent.left {
            // y是z的叔结点
            y := z.parent.parent.right
            if y.color == RED {
                // 如果叔结点为红：父结点和叔结点都改为黑，递归到祖父结点
                z.parent.color = BLACK
                y.color = BLACK
                z.parent.parent.color = RED
                z = z.parent.parent
            } else {
                if z == z.parent.right {
                    // 如果是三角形：先转直
                    z = z.parent
                    t.LeftRotate(z)
                }
                // 父和祖父颜色互换再转
                z.parent.color = BLACK
                z.parent.parent.color = RED
                t.RightRotate(z.parent.parent)
            }
        } else {
            // 对称同理
            y := z.parent.parent.left
            if y.color == RED {
                z.parent.color = BLACK
                y.color = BLACK
                z.parent.parent.color = RED
                z = z.parent.parent
            } else {
                if z == z.parent.left {
                    z = z.parent
                    t.RightRotate(z)
                }
                z.parent.color = BLACK
                z.parent.parent.color = RED
                t.LeftRotate(z.parent.parent)
            }
        }
    }
    t.root.color = BLACK
}

// 用一棵以v为根的子树来替换一棵以u为根的子树并成为其双亲的孩子结点
func (t *RBTree) Transplant(u, v *Node) {
    if u.parent == t.nil {
        t.root = v
    } else if u == u.parent.left {
        u.parent.left = v
    } else {
        u.parent.right = v
    }
    v.parent = u.parent
}

// 红黑树的删除
func (t *RBTree) Delete(z *Node) {
    // z为被删除结点，y为实际被删结点，x为拼接（替换）结点
    x, y := t.nil, z
    yOriginColor := y.color
    if z.left == t.nil {
        x = z.right
        t.Transplant(z, z.right)
    } else if z.right == t.nil {
        x = z.left
        t.Transplant(z, z.left)
    } else {
        // 找到中序后继
        y = t.Minimum(z.right)

        yOriginColor = y.color
        x = y.right
        if y.parent == z {
            // 当y的原父结点是z时，我们并不想让x.parent指向y的原始父结点，
            // 因为要在树中删除该结点。由于结点y将在树中向上移动占据z的位置，
            // 将x.parent设成y，使得x.parent指向y父结点的原始位置，
            // 甚至当x=T.nil时也是这样。
            x.parent = y
        } else {
            t.Transplant(y, y.right)
            y.right = z.right
            y.right.parent = y
        }
        t.Transplant(z, y)
        y.left = z.left
        y.left.parent = y
        y.color = z.color
    }
    // 如果实际被删结点为黑色，需要进行修复操作
    if yOriginColor == BLACK {
        t.deleteFixup(x)
    }
}

func (t *RBTree) deleteFixup(x *Node) {
    for x != t.root && x.color == BLACK {
        if x == x.parent.left {
            w := x.parent.right
            if w.color == RED {
                w.color = BLACK
                x.parent.color = RED
                t.LeftRotate(x.parent)
                w = x.parent.right
            }
            if w.left.color == BLACK && w.right.color == BLACK {
                w.color = RED
                x = x.parent
            } else {
                if w.right.color == BLACK {
                    w.left.color = BLACK
                    w.color = RED
                    t.RightRotate(w)
                    w = x.parent.right
                }
                w.color = x.parent.color
                x.parent.color = BLACK
                w.right.color = BLACK
                t.LeftRotate(x.parent)
                x = t.root
            }
        } else {
            w := x.parent.left
            if w.color == RED {
                w.color = BLACK
                x.parent.color = RED
                t.RightRotate(x.parent)
                w = x.parent.left
            }
            if w.right.color == BLACK && w.left.color == BLACK {
                w.color = RED
                x = x.parent
            } else {
                if w.left.color == BLACK {
                    w.right.color = BLACK
                    w.color = RED
                    t.LeftRotate(w)
                    w = x.parent.left
                }
                w.color = x.parent.color
                x.parent.color = BLACK
                w.left.color = BLACK
                t.RightRotate(x.parent)
                x = t.root
            }
        }
    }
    x.color = BLACK
}

// 根据Key创建并插入一个结点
func (t *RBTree) InsertKey(k int) {
    t.Insert(&Node{key: k, parent: t.nil, left: t.nil, right: t.nil})
}

func (t *RBTree) InsertKeys(keys ...int) {
    for _, v := range keys {
        t.InsertKey(v)
    }
}

// 根据Key查找并删除一个结点
func (t *RBTree) DeleteKey(k int) {
    node := t.SearchKey(k)
    t.Delete(node)
}

func (t *RBTree) DeleteKeys(keys ...int) {
    for _, v := range keys {
        t.DeleteKey(v)
    }
}

// 打印中序遍历结果
func (t *RBTree) InorderPrint() {
    fmt.Println("中序遍历结果：")
    t.inorderPrintNode(t.root, true)
    fmt.Printf("\n")
}
