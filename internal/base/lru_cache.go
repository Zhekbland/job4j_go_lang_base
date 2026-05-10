package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	Size int
	Head *Node
	Tail *Node
	Data map[string]*Node
}

func NewLruCache(size int) *LruCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	data := make(map[string]*Node, size)

	return &LruCache{
		Size: size,
		Head: head,
		Tail: tail,
		Data: data,
	}
}

func (l *LruCache) Put(key string, value string) {
	node, isNotEmpty := l.Data[key]
	if isNotEmpty {
		l.Remove(node)
		l.AddToHead(node)
		node.Value = value
		return
	}

	if len(l.Data) >= l.Size {
		delete(l.Data, l.Tail.Prev.Key)
		l.Remove(l.Tail.Prev)
	}

	newNode := &Node{
		Key:   key,
		Value: value,
	}
	l.AddToHead(newNode)
	l.Data[key] = newNode
}

func (l *LruCache) Get(key string) *string {
	node, isNotEmpty := l.Data[key]
	if !isNotEmpty {
		return nil
	}
	l.Remove(node)
	l.AddToHead(node)

	return &node.Value
}

func (l *LruCache) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *LruCache) AddToHead(node *Node) {
	node.Next = l.Head.Next
	node.Next.Prev = node
	l.Head.Next = node
	node.Prev = l.Head
}
