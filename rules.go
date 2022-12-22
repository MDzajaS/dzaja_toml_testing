package main

import (
	"fmt"
	"github.com/kkyr/fig"
	"encoding/json"
)

type Rules struct {
	RuleList []Rule `toml:"rule"`
}

type Rule struct {
	FieldLabel 			string 		`toml:"field_label"`
	FieldType 			string 		`toml:"field_type"`
	EncryptionType 		string 		`toml:"encryption_type"`
	Attribute1 			string 		`toml:"attribute1"`
	Attribute2 			string 		`toml:"attribute2"`
	Attribute3 			string 		`toml:"attribute3"`
	Attribute4 			string 		`toml:"attribute4"`
}

func (cfgRules *Rules) Read(filename string) error {
	return fig.Load(cfgRules, fig.File(filename), fig.Tag("toml"), fig.UseEnv(""))
}
	
func main() {
    var cfgRules Rules
	err := cfgRules.Read("rules.toml")
	if err != nil {
		fmt.Println(err)
	} else {
		cfgRulesJSON, _ := json.Marshal(cfgRules)
		fmt.Println(string(cfgRulesJSON))
	}
}