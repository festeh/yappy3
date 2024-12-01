package coach

type Callbacks struct {
	OnFocusButtonPress []func(c *Coach)
	OnFocusReceived    []func(c *Coach)
	OnFocusCompleted   []func(c *Coach)
	OnTick             []func(c *Coach)
}

func NewCallbacks() *Callbacks {
	return &Callbacks{}
}

func (c *Callbacks) RunOnFocusButtonPress(coach *Coach) {
	for _, f := range c.OnFocusButtonPress {
		f(coach)
	}
}

func (c *Callbacks) RunOnFocusReceived(coach *Coach) {
	for _, f := range c.OnFocusReceived {
		f(coach)
	}
}
func (c *Callbacks) RunOnFocusCompleted(coach *Coach) {
	for _, f := range c.OnFocusCompleted {
		f(coach)
	}
}

func (c *Callbacks) RunOnTick(coach *Coach) {
	for _, f := range c.OnTick {
		f(coach)
	}
}
