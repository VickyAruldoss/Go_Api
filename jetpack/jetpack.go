package jetpack

type Jetpack interface {
	Init()
	StartService()
	Print() string
}

type jetpack struct {
	isToStart bool
}

func NewJetpack() Jetpack {
	return jetpack{}
}
func (jetpack jetpack) Init() {

}
func (jetpack jetpack) Print() string {
	return "OMG it works"
}

func (jetpack jetpack) StartService() {

}
