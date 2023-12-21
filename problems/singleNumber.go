// package main

// import (
// 	"fmt"
// )
// func main(){
// 	a := []int{2,2,1,1,3}
// 	fmt.Println(singleNumber(a))
// }
// func singleNumber(nums []int) (ans int) {
// 	for _, v := range nums {
// 		ans ^= v
// 	}
// 	return
// }
// //  Sizning kod qatoringiz Go tilida yozilgan bo'lib, u jamlanmagan sonlarning (xorijiy) XOR-ayib olinishini amalga oshiradi. 
// //  Bu, oddiy XOR (exclusive or) operatsiyasini ishlatish orqali bajariladi.
// //  Bu tuyg'u o'z ichiga qancha elementning bir xil soni ko'p bo'lsa, shuncha marta o'sha sonni ans o'ziga xor qilib turadi.
// //  Bu qator oddiy XOR qo'llaniladigan ma'lumotlar qo'shish algoritmi misoliga oid bo'lib, bu orqali olish mumkin. 
// //  Agar 'nums' degan sliceda (to'plamda) barcha elementlar ikkala martada takrorlansa, natijada ans o'zini nolga xor qilib olishadi.
