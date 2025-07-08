package logr

type Options struct {
	reqPrompt   string
	respPrompt  string
	dumpingResp bool
}

type Option func(*Options)

func RequestPrompt(prompt string) Option {
	return func(options *Options) {
		options.reqPrompt = prompt
	}
}

func DumpingResponse(prompt string) Option {
	return func(options *Options) {
		options.dumpingResp = true
		options.respPrompt = prompt
	}
}

func getOptions(options ...Option) *Options {
	var option Options
	for _, o := range options {
		o(&option)
	}

	return &option
}
