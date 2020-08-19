package dog

import ( "fmt";"testing" )

var htodYears [14]int = [14]int{ 1, 7, 2, 14, 3, 21, 4, 28, 5, 35, 6, 42, 7, 49 }

func TestYears(t *testing.T) {

	dogYears := Years( 7 )

	if dogYears != 49 {
		t.Error("Expected 49, Got:", dogYears)
	}
}

func ExampleYears() {

	dogYears := Years( 7 )
	fmt.Println( dogYears )
	// Output: 49

}

func BenchmarkYears(b *testing.B) {
	for i:=0 ; i<b.N; i++ {
		_ = Years( i )
	}
}

func TestYearsTwo(t *testing.T) {

	dogYears := YearsTwo( 7 )

	if dogYears != 49 {
		t.Error("Expected 49, Got:", dogYears)
	}
}

func ExampleYearsTwo() {

	dogYears := YearsTwo( 7 )
	fmt.Println( dogYears )
	// Output: 49

}

func BenchmarkYearsTwo(b *testing.B) {
	for i:=0 ; i<b.N; i++ {
		_ = YearsTwo( i )
	}
}

