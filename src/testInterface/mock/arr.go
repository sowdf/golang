package mock

type Arr struct {
	a []int
}

func (r *Arr) Pop() int {
	head := r.a[0]
	r.a = r.a[1:]
	return head
}

func (r *Arr) Push(num int) {
	r.a = append(r.a, num)
}
