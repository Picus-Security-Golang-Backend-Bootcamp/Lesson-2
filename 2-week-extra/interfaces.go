package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

}

// Derste de öğrendiğimiz gibi interface'ler bir takım metot imzalarını (isimleri,parametreleri ve dönüş değerleri) barındıran taslak yapılardır.

// Go'da bir interface'in bir struct ya da tip tarafından implemente edilmesinin tek yolu struct içerisindeki metotları birebir aynı isim,parametre , ve dönüş değeri ile yazmaktır

// aşağıdaki tanım named interface olarak geçmektedir.
// Derste de bir arkadaşımız aşağıdaki tanım ile interface{} tanımının aynı şey olup olmadığını sormuştu. Hayır değil, eğer Go perspektifi ile bakacak olursak.
type Printable interface {
	Print()
}

// Önceki derslerimizde de defalarca söylediğim gibi isimlendirilmiş bir tip başka bir tip ile aynı şeyleri içerse bile Go runtime'ı için bunlar iki farklı tip olarak değerlendirilir.

// interface{} tanımı un-named/empty interface olarak geçer hiçbir metot taslağı olmayan boş anonim interface'lerdir. Yukarıdaki tanım ile farkı da bu. Eğer siz parametre olarak interface{} alıyorsanız aslında Go runtime'a özetle şunu söylemiş oluyorsunuz : "Bana bu interface'i (yani hiçbir şeyini) implement etmeyen her şeyi gönderebilirsin" generic yapılar Go'da bu şekilde sağlanmaktadır. Go 1.18 ile gelen Generic'ler ise bu karmaşıklığı çözerek bizlere yardımcı olmaktadır.

// interface{} Go dili içerisindeki her şey olabilir string,int,struct ya da sizin tanımladığınız özel bir tip.

// interface tanımlarının runtime anında reflection ile tipi belirlenir.
// eğer interface tipinde aldığınız bir değerin gerçek tipini bilmek istiyorsanız type assertion diğer bilinen adıyla downcasting yapabilirsiniz

func f() {
	var w io.Writer = os.Stdin

	// io.Writer bir interface ve biz bu interface'e os.Stdin değişkeninin tipi olan *os.File'ı atadık.

	// aşağıdaki kodu açıklayacak olursak değişken adınız nokta parantez içerisinde kontrol etmek istediğiniz tip
	// bu işlemden sonra c değişkenin üzerine mouse'unuz ile gelip baktığınızda artık onun da *os.File tipinde olduğunu görebilirsiniz.

	// çünkü w (io.Writer) içerisinde *os.File tipini barındırmaktadır.
	c := w.(*os.File) // OK
	b := w.(error)    // program burada panic fırlatacak ve uygulamanız sonlanacak

	// Yukarıdaki panic durumundan kaçınmanın yolu ise aşağıdaki gibidir.

	c, isOk := w.(*os.File)
	b, isOk = w.(error)

	if isOk == false {
		fmt.Println("HATA")
	}

	fmt.Println(c)
	fmt.Println(b)
}
