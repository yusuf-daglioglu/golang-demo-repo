package examples

import "fmt"

func Receiver_Function_Demo() {
	kedi := Kedi{}
	köpek := Kopek{}

	hayvanSesiniYazdir(kedi)  // Çıktı: Miyav
	hayvanSesiniYazdir(köpek) // Çıktı: Hav

	// Eğer "Hayvan" interface'inde "Run" fonksiyonu olsaydı,
	// o zaman Kedi ve Kopek, "Hayvan" olarak hayvanSesiniYazdir
	// fonksiyonuna parametre yollanamayacaktı (derleyici hata verecekti).
	// Go'da kalıtım olmadığı için, go-derleyicisi kalıtımı otomatik algılamaya çalışıyor.
	// Go yukarıdaki örnekte kalıtımı şuradan anlıyor: "Ses" fonksiyonu "Köpek" tarafından override edilmiş.
	// Demek ki "Kopek", "Hayvan"'ın tüm metotlarını override etmiş. Demek ki kalıtım yapılmak istenmiş.
	// Fakat "Run" fonksiyonu olsaydı, "Kopek" ve "Kedi" bunları override etmediğinden ötürü
	// kalıtım olmadığını düşünecekti.
}

type Hayvan interface {
	Ses() string
}

type Kedi struct{}

func (k Kedi) Ses() string {
	return "Miyav"
}

type Kopek struct{}

func (k Kopek) Ses() string {
	return "Hav"
}

func hayvanSesiniYazdir(h Hayvan) {
	fmt.Println(h.Ses())
}
