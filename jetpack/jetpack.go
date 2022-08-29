package jetpack

//go generate mockgen -source=jetpack/jetpack.go -destination= ./mocks/jetpack_mock.go

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
	if jetpack.isToStart {
		return "OMG it works"
	}
	return "OMG it works"
}

func (jetpack jetpack) StartService() {

}
