//package ScheduleCourseTest
package main

import (
	"Project1/Service/ScheduleCourse"
	"Project1/Types"
	"fmt"
)

func TestScheduleCourse() {
	var mp = make(map[string][]string)
	mp["t1"] = append(mp["t1"], "c1")
	mp["t1"] = append(mp["t1"], "c2")
	mp["t2"] = append(mp["t2"], "c1")
	r := &Types.ScheduleCourseRequest{TeacherCourseRelationShip: mp}
	res := ScheduleCourse.Schedule(r)
	fmt.Println(res)
}
func main() {
	TestScheduleCourse()
}
