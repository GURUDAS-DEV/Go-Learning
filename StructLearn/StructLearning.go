package StructLearn

import "fmt"

type Student struct {
	name   string
	rollno int
	mark   []int
}

func Main() {
	do()
}

func Structlearning() {
	var s1 Student

	s1.name, s1.rollno, s1.mark = "Gurudas", 1232, []int{12, 234, 43543, 63}

	fmt.Println("Name : ", s1.name, " Roll Number : ", s1.rollno, " Marks : ", s1.mark)
}

func PointerTOStruct() {
	s1 := Student{
		name:   "Gurudas",
		rollno: 1241182106,
		mark:   []int{99, 95, 97, 94},
	}

	p := &s1

	p.rollno = 23
	fmt.Println("roll no : ", p.rollno)
}

func ExtraTesting() {
	s1 := Student{
		name:   "Gurudas",
		rollno: 1241182106,
		mark:   []int{99, 95, 97, 94},
	}
	copy := s1

	fmt.Println("MARKS OF MAIN BEFORE : ", copy.mark)
	fmt.Println("MARKS OF COPY BEFORE : ", copy.mark)

	copy.mark[2] = 100

	fmt.Println("MARKS OF MAIN AFTER : ", copy.mark)
	fmt.Println("MARKS OF COPY AFTER : ", copy.mark)

}

func testing() {
	type test struct {
		p *int
	}

	num := 10

	structCopy := test{
		p: &num,
	}

	fmt.Println("Before : ", *(structCopy.p))
	num = 12
	fmt.Println("After 1st (12) from struct : ", *(structCopy.p))
	fmt.Println("After 1st (12)  from num : ", num)
	*structCopy.p = 34
	fmt.Println("After 2nd (34) from struct : ", *(structCopy.p))
	fmt.Println("After 2nd (34) from num : ", num)

}

func do() {
	type abc struct {
		x, y int
	}

	p := &abc{23, 24}
	fmt.Println(p.x)
}
