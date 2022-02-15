package ScheduleCourseTest
// package main

import (
	"Project1/Service/ScheduleCourse"
	"Project1/Types"
	"fmt"
	"math/rand"
	"time"
)

func TestScheduleCourse(mp map[string][]string) {
	fmt.Println("Processing ", mp)
	r := &Types.ScheduleCourseRequest{TeacherCourseRelationShip: mp}
	res := ScheduleCourse.Schedule(r)
	fmt.Println(res)
}
func generateMp(n int, m int) map[string][]string {
	var mp = make(map[string][]string)
	for i := 1; i <= n; i++ {
		var teacher string = "T"
		teacher += fmt.Sprintf("%d", i)
		for j := 1; j <= m; j++ {
			var a = rand.Int()
			var b = rand.Int()
			if a > b {
				course := "C"
				course += fmt.Sprintf("%d", j)
				mp[teacher] = append(mp[teacher], course)
			}
		}
	}
	fmt.Println("generated successfully")
	return mp
}
func main() {
	for true {
		var a = rand.Int()%10 + 1
		var b = rand.Int() % 10 % a
		a = 10
		b = 7
		fmt.Println("Size is ", a, " ", b)
		TestScheduleCourse(generateMp(a, b))
		time.Sleep(1000)
		// break
	}

}
