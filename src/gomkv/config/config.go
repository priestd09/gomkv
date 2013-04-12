package config

const (
	DEFAULT_PROFILE = "High Profile"
	DEFAULT_PREFIX  = ""
)

type GomkvConfig struct {
	Profile       string
	Prefix        string
	EpisodeOffset int
	SeasonOffset  int
	Episodic      bool
	AacOnly       bool
	Mobile        bool
	M4v           bool
	EnableSubs    bool
}
