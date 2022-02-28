package main

func main() {

}

type Book struct {
	Name     string
	StockQty int
}

// Struct metotlarınız aynı anda hem pointer hem de value reciever'ları ile kullanılamazlar
// Aşağıdaki kod satırları hata verecektir.
func (b Book) Sell(i int) {
	b.StockQty -= i
}

func (b *Book) Sell(i int) {
	b.StockQty -= i
}
