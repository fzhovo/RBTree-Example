package RBTree

const (
    BLACK Color = 0
    RED   Color = 1
)

type Color int

type Node struct {
    key    int
    left   *Node
    right  *Node
    parent *Node
    color  Color
}

func (node *Node) GetColorName() string {
    if node.color == RED {
        return "红"
    } else {
        return "黑"
    }
}
