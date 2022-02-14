package ScheduleCourse

import (
	"Project1/Types"
	"math"
)

type node struct {
	to  int
	net int
	val int
}

var n int = 0
var m int = 0
var head []int
var cur []int
var dep []int
var e []node
var e1 []node
var tot = 1

func dinit() {
	head = make([]int, n+5, n+5)
	cur = make([]int, n+5, n+5)
	dep = make([]int, n+5, n+5)
	e = make([]node, m+5, m+5)
	e1 = make([]node, m+5, m+5)
}
func copy1() {
	copy(e1, e)
}
func addEdge(u int, v int) {
	tot++
	e[tot].val = 1
	e[tot].net = head[u]
	e[tot].to = v
	tot++
	e[tot].val = 0
	e[tot].net = head[v]
	e[tot].to = u
}
func bfs() bool {
	for i := 0; i <= n; i++ {
		dep[i] = 0
	}
	dep[1] = 1
	cur[1] = head[1]
	var queue []int = make([]int, m, m)
	var st = 0
	var en = 1
	queue[0] = st
	for st < en {
		var k = queue[st]
		st++
		for i := head[k]; i > 0; i = e[i].net {
			var v = e[i].to
			if e[i].val > 0 && dep[v] == 0 {
				dep[v] = dep[k] + 1
				if v == n {
					return true
				}
				queue[en] = v
				en++
			}
		}
	}
	return false
}
func min(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(x int, sum int) int {
	if x == n {
		return sum
	}
	var ans = 0
	for i := cur[x]; i > 0; i = e[i].net {
		cur[x] = i
		var v = e[i].to
		if dep[v] == dep[x]+1 && e[i].val > 0 {
			var ans1 = dfs(v, min(sum-ans, e[i].val))
			if ans1 == 0 {
				dep[v] = 0
				continue
			}
			e[i].val -= ans1
			e[i^1].val += ans1
			ans += ans1
			if ans == sum {
				break
			}
		}
	}
	return ans
}
func dinic() {
	for bfs() {
		dfs(1, math.MaxInt32)
	}
}
func Schedule(request Types.ScheduleCourseRequest) Types.ScheduleCourseResponse {
	var res Types.ScheduleCourseResponse
	res.Code = Types.OK
	res.Data = make(map[string]string)
	id1 := 1
	id2 := 1
	var teacher map[int]string = make(map[int]string)
	var teacher1 map[string]int = make(map[string]int)
	var course map[int]string = make(map[int]string)
	var course1 map[string]int = make(map[string]int)
	for i := range request.TeacherCourseRelationShip {
		teacher[id1] = i
		teacher1[i] = id1
		courses := request.TeacherCourseRelationShip[i]
		for j := 0; j < len(courses); j++ {
			c := courses[j]
			m++
			if _, ok := course1[c]; ok {
			} else {
				course1[c] = id2
				course[id2] = c
				id2++
			}
		}
		id1++
	}
	n = id1 + id2
	dinit()
	id1--
	id2--
	for i := range request.TeacherCourseRelationShip {
		courses := request.TeacherCourseRelationShip[i]
		for j := 0; j < len(courses); j++ {
			addEdge(teacher1[i]+1, course1[courses[j]]+1+id1)
		}
	}
	for i := 2; i <= id1+1; i++ {
		addEdge(1, i)
	}
	for j := id1 + 2; j <= n; j++ {
		addEdge(j, n)
	}
	copy1()
	dinic()
	for i := 2; i <= id1+1; i++ {
		for j := head[i]; j != 0; j = e[j].net {
			if e1[j].val > 0 && e[j].val == 0 {
				res.Data[teacher[i-1]] = course[e[j].to-id1-1]
			}
		}
	}
	return res
}
