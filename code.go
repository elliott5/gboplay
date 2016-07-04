package gboplay

import "fmt"

type Playtime int

func (p Playtime) String() string {
	return fmt.Sprintf("{ play.Playtime value %d }", p)
}
