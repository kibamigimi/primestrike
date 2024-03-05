package main

type shootnum struct {
	number int
	flag   bool
	time   int
}

func shootnuminit() []*shootnum {
	var sl []*shootnum
	for i := range ballprime {
		sl = append(sl, &shootnum{
			number: ballprime[i],
			flag:   true,
			time:   0,
		})
	}
	return sl

}
func (g *game) set(num int, check bool) {
	for i := range g.shootnum {
		if g.shootnum[i].number == num {
			g.shootnum[i].flag = check
		}
	}
}
func (g *game) check(num int) bool {
	for i := range g.shootnum {
		if g.shootnum[i].number == num {
			return g.shootnum[i].flag
		}
	}
	return true
}
func (g *game) shootnumtimer(num int) {
	for i := range g.shootnum {
		if g.shootnum[i].number == num {
			g.shootnum[i].time = 450
		}
	}
}
func (g *game) shootnumcount() {
	for i := range g.shootnum {
		if g.shootnum[i].flag == false {
			g.shootnum[i].time -= 1
		}
		if g.shootnum[i].time == 0 {
			g.shootnum[i].flag = true
		}
	}
}
func (g *game) shootnumtime(num int) int {
	for i := range g.shootnum {
		if g.shootnum[i].number == num {
			return g.shootnum[i].time
		}
	}
	return 0
}
