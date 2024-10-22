package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CheckQuad()
}

func CheckQuad() {
	degerler, _ := io.ReadAll(os.Stdin) // İnputtan girdiğimiz değerleri okur.
	alinan_komutlar := string(degerler)
	if len(alinan_komutlar) == 0 { // Alınan komutların boş olup olmadığını kontrol etme.
		fmt.Println("Bilgiler bulunamadi !Boş Dosya!")
		return
	}
	if alinan_komutlar[len(alinan_komutlar)-1] != '\n' { // Alınan komutlara yeni satır ekleme.
		alinan_komutlar += "\n"
	}
	x := 0
	y := 0
	for i, v := range alinan_komutlar { // Alınan komutlar ile döngü başlatır.
		if v == '\n' {
			y++
		}
		if y == 0 {
			x = i + 1 // Eğer y 0 ise, x'i geçerli karakter indeksini + 1'e ayarlayın.
		}
		ilk_karakter := alinan_komutlar[0] // Karakter girişi. (input)
		sayac := 0
		if ilk_karakter == 'o' && alinan_komutlar == QuadA(x, y) { // Girişi kontrol etti ve buna göre 'yazdir' fonksiyonunu çağırdı.
			yazdir(x, y, &sayac, 'A')
		}
		if ilk_karakter == '/' && alinan_komutlar == QuadB(x, y) {
			yazdir(x, y, &sayac, 'B')
		}
		if ilk_karakter == 'A' && alinan_komutlar == QuadC(x, y) {
			yazdir(x, y, &sayac, 'C')
		}
		if ilk_karakter == 'A' && alinan_komutlar == QuadD(x, y) {
			yazdir(x, y, &sayac, 'D')
		}
		if ilk_karakter == 'A' && alinan_komutlar == QuadE(x, y) {
			yazdir(x, y, &sayac, 'E')
		}
		if sayac == 0 { // sayac 0 ise bu bir fonksiyon değildir.
			fmt.Println("Not a quad function")
		} else {
			fmt.Println()
		}
	}
}

func yazdir(x, y int, sayac *int, l rune) {
	if *sayac > 0 {
		fmt.Printf(" || ")
	}
	fmt.Printf("[quad%s] [%d] [%d]", string(l), x, y) // Koordinatları yazdır.
	*sayac++
}
