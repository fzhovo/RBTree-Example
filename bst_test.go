package RBTree

import (
    "fmt"
    "testing"
)

func TestMakeTreeBySequence(t *testing.T) {
    bst := MakeTreeBySequence("5,1,4,null,null,3,6")
    bst.InorderPrint()
    fmt.Println(bst.validate())

    bst2 := MakeTreeBySequence("4,2,5,1,3")
    bst2.InorderPrint()
    fmt.Println(bst2.validate())
}

func TestBST(t *testing.T) {
    bst := MakeBST()
    bst.Insert(&Node{key: 3})
    bst.Insert(&Node{key: 2})
    bst.Insert(&Node{key: 5})
    fmt.Println(bst.root)
    bst.InorderPrint()
    x := bst.SearchKey(3)
    bst.Delete(x)
    bst.InorderPrint()
}

func TestRBTree(t *testing.T) {
    rbt := MakeRBTree()
    rbt.InsertKeys(1, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
    rbt.InorderPrint()
    rbt.DeleteKeys(14, 9, 5)
    rbt.InorderPrint()
}
