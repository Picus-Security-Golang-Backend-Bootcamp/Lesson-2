package main

import (
	"fmt"
	"time"
)

func main() {

	// Map tipi Go'da bir hash-table implementasyonu olarak karşımıza çıkar
	// Key olarak verilen bir değer belirli algoritmalardan geçirilerek bir index değeri oluşturulur ve bu index'in işaret ettiği lokasyonda key'e karşılık gelen verimiz yer alır.
	// Eğer farklı bir key aynı algoritmalar sonucunda aynı index'i işaret ederse bu sefer o index'in işaret ettiği lokasyonda bir linked-list oluşturulur.

	// Map key-value şeklinde değer tutan yapılardır ve slice'lara göre kullanım çeşitliliği daha fazladır. Karşılaştırılabilen tüm tipler key olarak kullanılabilir.

	// Elbette bool tipi de karşılaştırılabilen (comperable) bir tiptir ancak sadece iki değer alabilmesinden dolayı(true & false) key olarak kullanılması önerilmez.
	m := map[string]Employee{"my-key1": {"Name", 200, time.Now()}}

	fmt.Println(m)

	// Bir struct map'indeki değerlerden memory adresini alamazsınız. Bunun sebebi siz her veri eklediğinizde ya da çıkardığınızda tekrar bir memory allocation olacağından ötürü memory adresini aldığınız değerin kaybolacak olması (memory'deki yerinin değişecek olması)

	// aşağıdaki kod hata verecektir.
	v := &m["my-key1"]

	fmt.Println(v)

	// Aynı sebepten ötürü map içerisindeki struct'a ait bir değerin değerini de değiştiremezsiniz.

	// aşağıdaki kod yukarıda anlattığım sebepten dolayı hata verecektir.

	m["my-key1"].Salary += 10

	// Map içerisinde bir key'in olup olmadığını kontrol etmek için aşağıdaki kodu kullanabilirsiniz

	isExist := m["my-key1"]

	fmt.Println(isExist)

	// Eğer map içerisinde bir değeri silmek isterseniz aşağıdaki kodu kullanabilirsiniz.
	// delete fonksiyonuna map'te olmayan key değeri de verebilirsiniz eğer bu key map içerisinde yoksa hiçbir işlem yapmayacaktır eğer varsa o key-value değerini siler bundan dolayı herhangi bir hata durumu yoktur ve yine bu sebepten dolayı geriye dönüş değeri mevcut değil

	// ilk parametre silme işlemi yapacağınız map değişkeni
	// ikinci parametre silmek istediğiniz key değeri
	delete(m, "my-key2")

	// Bu anlattıklarımın haricinde map içerisinde struct'ları pointer olarak tutabilirsiniz.
	// Yukarıda map içerisindeki bir struct'ın memory adresinin değişeceğinden bundan dolayı da memory adresini alamayacağımızdan ya da bir alanın değerini değiştiremeyeceğimizden bahsetmiştik

	// ama eğer böyle bir ihtiyacınız varsa pointer struct'ları map ile kullanabilirsiniz

	b := map[string]*Employee{"my-key1": {"Name", 200, time.Now()}}

	// aşağıdaki kod hata vermeden çalışacaktır
	// Çünkü artık map değişkenimizde struct değil struct pointer tutuyoruz ve bundan dolayı map tekrar memory allocation yapsa da pointer değerlerimiz değişmeyecek ve sorunsuz bir şekilde kodumuz çalışacak.
	b["my-key1"].Salary += 10

	fmt.Println(b)
}

type Employee struct {
	Name      string
	Salary    float64
	CreatedAt time.Time
}
