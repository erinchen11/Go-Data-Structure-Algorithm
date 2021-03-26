package dataStructure

import "errors"

/*
使用node(節點)來記錄、表示、儲存資料(data)，
並利用每個node中的pointer指向下一個node，藉此將多個node串連起來，
形成Linked list，並以NULL來代表Linked list的終點

*/

// struct : Node, LinkedList

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

//以LinkedList為型態 其所支援的幾種操作，就是method啦

/*
新增node在linked list
刪除node在linked list
查詢list的長度
查詢 list上某個node的值
*/

//創建新 linkedlist
func NewLinkedList() *LinkedList {
	list := new(LinkedList)
	list.Length = 0
	return list
}

//創建新 node

func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

// linked list的長度

func (list *LinkedList) Len() int {
	return list.Length
}

// 檢查 linked list 是否存在 node

func (list *LinkedList) IsEmpty() bool {
	return list.Length == 0
}

// linked list 前面新增node

func (list *LinkedList) Prepend(value interface{}) {
	node := NewNode(value)
	if list.Len() == 0 {
		//該node為linked list的第一個node
		list.Head = node
		list.Tail = list.Head
	} else {
		//向前新增一個node時，當前node的 prev會指向新的node
		currentHead := list.Head
		currentHead.Prev = node
		//新的node會指向當前node
		node.Next = currentHead
		// list的head改為新的node
		list.Head = node
	}

	list.Length++
}

// 在 LinkedList向後新增一個node
func (list *LinkedList) Append(value interface{}) {
	node := NewNode(value)
	if list.Len() == 0 {
		list.Head = node
		list.Tail = list.Head
	} else {
		//向後新增一個node時，當前node的next會指向新的node
		currentTail := list.Tail
		currentTail.Next = node
		//新的node會指向當前node
		node.Prev = currentTail
		list.Tail = node
	}
	list.Length++

}

//以index獲取該index之前的node
func (list *LinkedList) GetByIndex(index int) (*Node, error) {
	if index > list.Length {
		return nil, errors.New("index out of range")
	}
	//
	node := list.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node, nil

}

//以node的value刪除linked list上的該node
//考量情況有三種，沒有node可以刪除、刪除第一個node、刪除其他位置的node
func (list *LinkedList) Remove(value interface{}) error {
	//先檢查linkedlist中有沒有node
	if list.Len() == 0 {
		return errors.New("Linked list is empty.")
	}

	//如果head所指的node就是要刪除的node
	if list.Head.Value == value {
		//刪掉該node時，因為是第一個, head要指向下一個
		list.Head = list.Head.Next
		//linked list總長度要-1
		list.Length--
		//完成後回傳error為nil
		return nil
	}
	//刪除其他位置的node

	//以 head為指標, 開始輪詢 linked list，直到head指不到東西，每次head有指到node, head就會再繼續下一個node
	//用for傳統迴圈，初值;限制範圍;變化條件

	target := 0
	for current := list.Head; current != nil; current = current.Next {
		//當head指向的node，該node的值與要刪的node之值一樣時，且target 還是0
		//就是要刪那個node
		if current.Value.(*Node) == value && target == 0 {
			//刪完node後，該node的下一個node會指向 該node的前一個node
			current.Next.Prev, current.Prev.Next = current.Prev, current.Next
			list.Length--
			target++
		}

	}
	if target == 0 {
		return errors.New("Node not found")
	}

	return nil

}

//查詢某個node


