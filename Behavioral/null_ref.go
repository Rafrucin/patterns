package patterns

//NULL REF

type student struct {
	id   int
	name string
}

var students = []student{{id: 1, name: "AA"}, {id: 2, name: "BB"}}

func getStudentById(id int) student {
	for k, v := range students {
		if v.id == id {
			return students[k]
		}
	}
	return student{}
}

func NullRef() {
	s1 := getStudentById(1)
	println("s1", s1.id, s1.name)
	s2 := getStudentById(3)
	println("s2", s2.id, s2.name)
}
