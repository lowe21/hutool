package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic(err)
	}
}
