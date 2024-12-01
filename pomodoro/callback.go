package pomodoro

type Callbacks struct {
  OnStart []func(pomodoro *Pomo)
  OnStop []func(pomodoro *Pomo)
  OnTick []func(pomodoro *Pomo)
  OnFinish []func(pomodoro *Pomo)
}

func NewCallbacks() *Callbacks {
  return &Callbacks{}
}

func (c *Callbacks) AddStart(f func(pomodoro *Pomo)) {
  c.OnStart = append(c.OnStart, f)
}

func (c *Callbacks) AddStop(f func(pomodoro *Pomo)) {
  c.OnStop = append(c.OnStop, f)
}

func (c *Callbacks) AddTick(f func(pomodoro *Pomo)) {
  c.OnTick = append(c.OnTick, f)
}

func (c *Callbacks) AddFinish(f func(pomodoro *Pomo)) {
  c.OnFinish = append(c.OnFinish, f)
}

func (c *Callbacks) RunOnStart(pomodoro *Pomo) {
  for _, f := range c.OnStart {
    f(pomodoro)
  }
}

func (c *Callbacks) RunOnStop(pomodoro *Pomo) {
  for _, f := range c.OnStop {
    f(pomodoro)
  }
}

func (c *Callbacks) RunOnTick(pomodoro *Pomo) {
  for _, f := range c.OnTick {
    f(pomodoro)
  }
}

func (c *Callbacks) RunOnFinish(pomodoro *Pomo) {
  for _, f := range c.OnFinish {
    f(pomodoro)
  }
}
