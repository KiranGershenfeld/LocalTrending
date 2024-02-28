package jobs

type Job interface {
	Run() error
	Fail() error
}
