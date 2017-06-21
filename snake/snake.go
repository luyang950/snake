package snake

const (
	Up = iota
	Down
	Left
	Right
)

type Snake struct {
	Head *node
	Body []*node
}

type node struct {
	X int
	Y int

	Direction uint

	Pre  *node
	Next *node
}

func (s *Snake) Init() {
	s.Head = &node{
		X:         3,
		Y:         3,
		Direction: Right,
	}
	//s.Head.Next = nil
	s.Head.Next = &node{
		X:         3,
		Y:         2,
		Direction: Right,
		Pre:       s.Head,
	}
	s.Head.Pre = nil
	s.Head.Next.Next = &node{
		X:         3,
		Y:         1,
		Direction: Right,
		Pre:       s.Head.Next,
		Next:      nil,
	}
}

func (s *Snake) HeadUp() {
	s.Head.Direction = Up
}

func (s *Snake) HeadDown() {
	s.Head.Direction = Down
}

func (s *Snake) HeadLeft() {
	s.Head.Direction = Left
}

func (s *Snake) HeadRight() {
	s.Head.Direction = Right
}

func (s *Snake) DirectionName() string {
	switch s.Head.Direction {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}
	return "Right"
}

func (s *Snake) OppoDirection() uint {
	switch s.Head.Direction {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	return Right
}

func (s *Snake) NextStep() (int, int) {
	var x, y int

	switch s.Head.Direction {
	case Up:
		x = s.Head.X - 1
		y = s.Head.Y
	case Down:
		x = s.Head.X + 1
		y = s.Head.Y
	case Left:
		x = s.Head.X
		y = s.Head.Y - 1
	case Right:
		x = s.Head.X
		y = s.Head.Y + 1
	}

	return x, y
}

func (s *Snake) AutoMove() {
	// gen new head
	newHead := &node{
		X:         s.Head.X,
		Y:         s.Head.Y + 1,
		Direction: s.Head.Direction,
		Next:      s.Head,
		Pre:       nil,
	}

	newHead.X, newHead.Y = s.NextStep()

	s.Head = newHead
	s.Head.Next.Pre = s.Head

	cur := s.Head

	// del the last node
	for {
		if cur.Next != nil {
			cur = cur.Next
		} else {
			cur.Pre.Next = nil
			break
		}
	}
}

func (s *Snake) Grow() {
	// gen new head
	newHead := &node{
		X:         s.Head.X,
		Y:         s.Head.Y + 1,
		Direction: s.Head.Direction,
		Next:      s.Head,
		Pre:       nil,
	}

	newHead.X, newHead.Y = s.NextStep()

	s.Head = newHead
	s.Head.Next.Pre = s.Head
}

func (s *Snake) Die(dieSignal chan bool) {
	dieSignal <- true
}
