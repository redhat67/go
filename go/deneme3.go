// Paket cümlesi (Package Clause)
/* package main

import "fmt"

func main() { */
/* 3 FARKLI YÖNTEMLE YAZDIRMA /*
/* var studentName string="Atamer"
   var grade int = 77
   var isPassed bool  = true */

/* var studentName ="Atamer"
   var grade  = 77
   var isPassed = true */

/* studentName := "Atamer"
grade := 20
isPassed := true */

/* TEK SATIR HALİNDE YAZDIRMA*/
/* 	var studentName, grade, isPassed = "Atamer", 20, true
 */

/* studentName, grade, isPassed := "Atamer", 20, true */

package main

import "fmt"

func deneme3() {

	/*   "var studentName string"  /*  /* BU KISIMA DECLARATİON DENİR. EĞER  ="İSİM" ŞEKLİNDE ATAMA YAPILIRSA BUNA ASSİGN DENİR. TAMANINA İSE İNİTİALİZATİON DENİR. */

	studentName := "Atamer"
	/*studentName := "Altaş"       TEKRARDAN DEĞER ATANAMAZ ATANMASI İÇİN SADECE = İLE DEĞER ATANMALIDIR. */

	fmt.Println(studentName)

}
