package constanta

// ConfigConst is global constanta for config function
type ConfigConst string

// Add function to get value in string for configConstanta
func (cc ConfigConst) String() string {
	return string(cc)
}

// This constanta util for config
const (
	CfgService  ConfigConst = "service"
	CfgDatabase ConfigConst = "database"
	CfgEtc      ConfigConst = "etc"
)
