package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, terbesar, terkecil int
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
		terbesar = A[i]
		terkecil =A[i]
	}
	for i = 0; i < n; i++{
		if A[i] > terbesar{
			terbesar = A[i]
		}else if A[i] < terkecil{
			terkecil = A[i]
		}
	}
	for terkecil != 0 {
        temp := terkecil
        terkecil = terbesar % terkecil
        terbesar = temp
    }
	fmt.Print(terbesar)
}
