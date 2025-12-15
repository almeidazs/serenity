package types

type Config struct {
	Schema        string     `json:"$schema"`
	Naming        *RuleGroup `json:"naming"`
	Complexity    *RuleGroup `json:"complexity"`
	BestPractices *RuleGroup `json:"bestPractices"`
	ErrorHandling *RuleGroup `json:"errorHandling"`
	Imports       *RuleGroup `json:"imports"`
	Exclude       []string   `json:"exclude"`
}

type RuleGroup struct {
	Enabled bool           `json:"enabled"`
	Rules   map[string]any `json:"rules"`
}

