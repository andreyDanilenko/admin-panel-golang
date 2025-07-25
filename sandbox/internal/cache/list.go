package cache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: l.front}

	if l.front != nil {
		l.front.Prev = newItem
	} else {
		l.back = newItem
	}

	l.front = newItem
	l.len++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: l.back}

	if l.back != nil {
		l.back.Next = newItem
	} else {
		l.front = newItem
	}

	l.back = newItem
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	// Проверяем наличие предыдущего у удаляемого элемента
	if i.Prev != nil {
		// Если он есть мы обращаемся к нему и меняем следующий у предыдущего на следующий у удаляемого
		i.Prev.Next = i.Next
	} else {
		// Если его нет то это был первый элемент и мы делаем следующий у удаляемого первым элементом
		l.front = i.Next
	}

	// Проверяем наличие следующего у удаляемого жлемента
	if i.Next != nil {
		// Если следующий есть мы меняем его предыдущий на предыдущий удаляемого
		i.Next.Prev = i.Prev
	} else {
		// Если его нет то это был последний и мы делем предыдущий у кдаляемого последним элементом
		l.back = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	// Add to front
	i.Prev = nil
	i.Next = l.front
	if l.front != nil {
		l.front.Prev = i
	} else {
		l.back = i
	}
	l.front = i
}
