package main

func deneme12() {

	/* for i := 1; i <= 10; i++ {
		fmt.Println(i)
	} */

	/* i := 1
	for ; i <= 10; i++ {
		fmt.Println(i)   // zor gelirse bunu kullan daha anlaşılır

	} */

	/* i := 0
	for ; i <= 10; i += 5 {
		fmt.Println(i)
	}
	fmt.Println(i) */

	/* i := 10
	for i >= 0 {
		fmt.Println(i) // değerleri alt alta verdiğimiz için noktalıvirgül(;) koymamıza gerek yok
		i--
	} */

	/* for i := 0; i <= 10; i++ {
		if i%4 == 0 { // burada ise i nin 4 e bölüp 0 sonucu çıkarttığı sayıları yazdır demek yani: 10 üzerinde 4 ise 4,8 10 üzerinde 3 ise 3,6,9.
			continue // bu ifade döngünün başına git demek
		}
		fmt.Println(i)
	} */

	/* for i := 0; i <= 10; i++ {
		if i == 3 {
			break // break döngüden çıkar
		}
	} */
}
