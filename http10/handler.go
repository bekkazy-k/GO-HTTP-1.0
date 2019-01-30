package http10

type HandlersMap map[string]func(r *Request)

func (hm HandlersMap) RegisterHandler(path string, handler func(r *Request)) {
	hm[path] = handler
}

var M = make(HandlersMap, 0)

// func main() {

//   Get("1", myFunc1)
//   Get("2", myFunc2)

//   m["1"]("handler 1 output")
//   m["2"]("handler 2 output")

// }

func Get(path string, handler func(r *Request)) {
	M.RegisterHandler(path, handler)
}

// func myFunc1(s string) {
//   fmt.Println("MyFunc1:", s)
// }

// func myFunc2(s string) {
//   fmt.Println("MyFunc2:", s)
// }
