package datastructure

//stack是 LIFO 或 FILO

//創建一個stack
// 操作有 push, pop, isEmpty

type Stack struct {
	list []interface{}
}

func New() *Stack {
	s := new(Stack)

	//stack本身是固定大小的array或 linked list
	s.list = make([]interface{}, 0, 8)

	return s

}

func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

//刪除最頂端的stack 並且回傳該元素
func (s *Stack) Pop() interface{} {

	//如果 stack為空，則回傳nil
	if s.IsEmpty() {
		return nil
	}

	//取出最末端的元素
	temp := s.list[len(s.list)-1]

	//取出後要更新stack的長度，但是我們不用index的方法
	//就把新的stack指定就好
	s.list = s.list[:len(s.list)-1]

	return temp

}

func (s *Stack) Push(item interface{}) {

	//Ｐush是加入最末端的，用append就可以了

	s.list = append(s.list, item)
}

//取出最頂端元素，但是不刪除該元素
func (s *Stack) Peek() interface{} {
	if len(s.list) == 0 {
		return nil
	}

	return s.list[len(s.list)-1]
}
