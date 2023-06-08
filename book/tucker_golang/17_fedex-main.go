// package main

// import "golang/book/tucker_golang/fedex"

// func SendBook(name string, sender *fedex.FedexSender) {
// 	sender.Send(name)
// }
// func main() {
// 	sender := &fedex.FedexSender{}
// 	SendBook("어린왕자", sender)
// 	SendBook("신데렐라", sender)
// }

// ========= fedex -> fedexkorea
// 기존 코드를 다 변경해야한다.

package main

func SendBook(name string, sender *koreaPost.PostSender) {
	sender.Send(name)
}
func main() {
	sender := &koreaPost.PostSender{}
	SendBook("어린왕자", sender)
	SendBook("신데렐라", sender)
}
