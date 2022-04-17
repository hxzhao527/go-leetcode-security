package linkedlist

type LRUCache struct {
	cap   int
	cache map[int]int

	order         *DoubleLinkedList
	order_mapping map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:           capacity,
		cache:         make(map[int]int),
		order:         &DoubleLinkedList{},
		order_mapping: make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.cache[key]
	if !ok {
		return -1
	}
	old_node := this.order_mapping[key]
	// delete(this.order_mapping, key)
	this.order.remove(old_node)

	new_node := this.order.push_front(key)
	this.order_mapping[key] = new_node

	return val
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.cache[key]
	if ok {
		this.Get(key)
		this.cache[key] = value
		return
	}

	this.cache[key] = value
	node := this.order.push_front(key)
	this.order_mapping[key] = node

	if len(this.cache) > this.cap {
		last_node := this.order.tail
		this.order.remove(last_node)
		delete(this.cache, last_node.val)
		delete(this.order_mapping, last_node.val)
	}
}

type Node struct {
	val int

	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	len int

	head *Node
	tail *Node
}

func (list *DoubleLinkedList) push_front(val int) *Node {
	node := &Node{
		val:  val,
		next: list.head,
	}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.head.prev = node
		list.head = node
	}
	list.len += 1

	return node
}

func (list *DoubleLinkedList) push_back(val int) *Node {
	node := &Node{
		val:  val,
		prev: list.tail,
	}
	if list.tail == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}

	list.len += 1
	return node
}

func (list *DoubleLinkedList) remove(node *Node) {
	if node == nil {
		return
	}

	if node == list.head {
		// if list.head.next != nil{
		//     list.head.next.prev = nil
		// }
		list.head = list.head.next
	}
	if node == list.tail {
		// if list.tail.prev != nil{
		//     list.tail.prev.next = nil
		// }
		list.tail = list.tail.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev = nil
	node.next = nil
	list.len -= 1
}
