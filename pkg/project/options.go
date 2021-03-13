package project

type Option func(p *Project)

func WithCustom(k string, v interface{}) Option { return func(p *Project) { p.Custom[k] = v } }

func WithCustomMap(v map[string]interface{}) Option { return func(p *Project) { p.Custom = v } }
