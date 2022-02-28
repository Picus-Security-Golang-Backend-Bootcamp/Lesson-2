package main

import "fmt"

func main() {

	// struct ve map kullanımında struct'ların memory adreslerini alamayacağınızdan bahsetmiştim (struct_map.go dosyasına bakınız)
	// Go struct slice'larda sizin struct'larınızın memory adreslerini almanıza izin veriyor ancak bu oldukça riskli ve tavsiye olarak eğer gerçekten ne yaptığınızı bilmiyorsanız uzak durmanız gereken bir işlem

	// örnek olarak aşağıdaki kodda Comics struct'ı ile bir slice oluşturduk
	v := []Comics{{"Batman - Hush", 0}, {"Batman - Dark Knight Returns", 0}}

	// Burada v isimli Comics struct'ının ilk elemanının memory adresini aldık
	batman := &v[0]

	// Yeni bir değer oluşturduk
	superman := Comics{"Superman - New World", 1}

	// oluşturduğumuz değeri slice'a ekledik
	v = append(v, superman)
	// Memory adresini aldığımız struct'ın ilk elemanının count değerini 1 arttırdık
	increaseCount(batman)

	// Ancak buraya baktığınızda değerini arttırdığımız elemanın hala count değerini 0 göreceksiniz. Bunun sebebi biz struct'ın memory adresini değil slice'ın ilk elemanının memory adresini aldık ve eleman ekledikçe/çıkardıkça değişebilecek bir değer

	// append işleminde ne zaman memory adresi değiştiğine dair konular için ilk derse bakabilirsiniz
	fmt.Println(v)

	// Buradaki sorunu da aynı map'de olduğu gibi slice içerisinde struct'ın pointer'larını tutarak çözebilirsiniz.
}

func increaseCount(c *Comics) {
	c.Count++
}

type Comics struct {
	Name  string
	Count int
}
