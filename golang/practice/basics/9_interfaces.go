package basics

import (
	"fmt"
)

func Main_interfaces() {
	// ======== Big Picture: 1
	p1 := Point{
		x: 0,
		y: 0,
	}

	p1.x, p1.y = p1.Move(2, 3)
	fmt.Println(p1.x, p1.y)
	p1.Shift(6, 7)
	fmt.Println(p1.x, p1.y)

	// ======== Big Picture: 2
	var ll1 *LinkedList
	ll1 = ll1.Insert(3)
	ll1 = ll1.Insert(5)
	ll1 = ll1.Insert(8)
	ll1.Print()

	// ======== Big Picture: 3
	var s1 Server
	var l1 Logger
	l1.Log("hello")
	s1.Log("hi")
	// fmt.Println(s1 == l1) - shows error

	// ======== Big Picture: 4
	ActivateStringer(p1)

	// ======== Big Picture: 5
	p2 := NewPoint(0, -1)
	PrintShape(p2)

	// ======== Big Picture: 6
	var s fmt.Stringer
	fmt.Println(s == nil)
	var p *Point
	s = p
	fmt.Println(s == nil, p == nil)
	// Once we assign uninitialized interface to uninitialized value, it becomes not null.
	// Because interface is {*value: {}, *type: {}}
	// if we assign interface to a value even if not initialized
	// its type field will be assigned and becomes not nil

	// ======== Big Picture: 7
	Dump(5)
	Dump("Hello")
	Dump(p)
	Dump(s)

	// ======== Big Picture: 9
	hf := func(i int) {
		fmt.Printf("Handling %d\n", i)
	}

	ad := &HandlerAdapter{hf: hf}
	ad.Handle(42)

	// ======== Big Picture: 10

	var rdb RealDB
	var tdb TestDB

	HandleRequest(&rdb)
	HandleRequest(&tdb)
}

// ======== Big Picture: 1

type Point struct {
	x, y int
}

func (p Point) Move(dx, dy int) (int, int) {
	p.x = p.x + dx
	p.y = p.y + dy

	return p.x, p.y
}

func (p *Point) Shift(dx, dy int) {
	p.x += dx
	p.y += dy
}

// ======== Big Picture: 2

type LinkedList struct {
	val  int
	next *LinkedList
}

func (ll *LinkedList) Insert(a int) *LinkedList {
	if ll == nil {
		return &LinkedList{val: a} // considering nil condition
	}

	ptr := ll
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = &LinkedList{val: a}

	return ll
}

func (ll *LinkedList) Print() {
	ptr := ll
	for ptr != nil {
		fmt.Printf("%d -> ", ptr.val)
		ptr = ptr.next
	}
	fmt.Println()
}

// ======== Big Picture: 3

type Logger struct{}

type Server struct {
	Logger
}

func (l Logger) Log(msg string) {
	fmt.Println("Logger message:", msg)
}

// ======== Big Picture: 4

type Stringer interface {
	String() string
}

func (p Point) String() string {
	return fmt.Sprintf("Point{x: %d, y: %d}", p.x, p.y)
}

func ActivateStringer(s Stringer) {
	// Accepting interface, Point is implicitly implementing it
	fmt.Println(s.String())
}

// ======== Big Picture: 5

func PrintShape(s fmt.Stringer) {
	fmt.Println(s.String())
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

// ======== Big Picture: 7, 8

// func Dump(v interface{}) {
// 	fmt.Printf("%T\n", v)
// }

// func Dump(v interface{}) error {
// 	i, ok := v.(int)
// 	if !ok {
// 		return fmt.Errorf("unexpected type for %v", v)
// 	}
// 	fmt.Printf("%T\n", i)
// 	return nil
// }

func Dump(v interface{}) {
	switch i := v.(type) {
	case int:
		fmt.Println("Type is int", i)
	case string:
		fmt.Println("Type is string", i)
	}
}

// ======== Big Picture: 9

type HandlerFunc func(int)

type Handler interface {
	Handle(int)
}

type HandlerAdapter struct {
	hf HandlerFunc
}

func (ha *HandlerAdapter) Handle(i int) {
	ha.hf(i)
}

// ======== Big Picture: 10

type Database interface {
	GetUser(id int) string
}

type RealDB struct{}
type TestDB struct{}

func (rdb *RealDB) GetUser(id int) string {
	return "Eshmat"
}

func (tdb *TestDB) GetUser(id int) string {
	return "Gishmat"
}

func HandleRequest(db Database) {
	fmt.Println("Handling database request...")
	username := db.GetUser(4)
	fmt.Println("User got:", username)
}
