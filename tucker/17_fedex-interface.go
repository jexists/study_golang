package main

import "golang/book/tucker_golang/fedex"

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &fedex.FedexSender{}
	SendBook("어린왕자", sender)
	SendBook("신데렐라", sender)
}

// ========= fedex -> fedexkorea

// func main() {
// 	var sender Sender = &koreaPost.PostSender{}
//  이부분만 변경하면됨
// 	SendBook("어린왕자", sender)
// 	SendBook("신데렐라", sender)
// }
