package pkg

type Language interface {
	ID() string
	DisplayName() string
	DefaultFileName() string
	Compile() error
	Run() error
}
