package main

import (
	"errors"
	"fmt"
)

func main() {

	// Go içerisinde yer alan new() fonksiyonu ise sizin struct'ınızı oluşturup içerisindeki tüm alanlara zero-value (varsayılan değer örn : integer için 0)  atayarak geriye bir pointer döner.

	p1 := new(Person) // new metodu sizin struct'ınız için allocation yaparak memory'deki yerini hazırlar.

	/*
		Dönüş değeri : *Person => FirstName : "", LastName : "", Age : 0
	*/

	fmt.Println(p1)

	p, err := New("First name", "Last name", 28)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v\n", p)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Go'da diğer dillerdan farklı olarak bir constructor yapısı bulunmamakta
// Bunun yerine kendiniz özel bir constructor fonksiyonu yazmalısınız.

// Fonksiyon adı New olmak zorunda değil istediğiniz isimle struct için constructor fonksiyon yazabilirsiniz ancak genel olarak New ya da New ile başlayan isimlendirmeler tercih edilir
func New(firstName, lastName string, age int) (*Person, error) {
	// bu fonksiyon içinde projenizde yaratacağınız struct'ın her zaman geçerli olduğunu garanti edebilirsiniz ve istediğiniz kuralı burada kontrol edebilirsiniz
	if age < 18 {
		return nil, errors.New("18 yaşından büyük olmalısınız")
	}
	person := &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	return person, nil
}
