package webconfig

type ConfigErrors map[string]error

func (this *ConfigErrors) Len() int { // {{{
	return len(*this)
} // }}}
