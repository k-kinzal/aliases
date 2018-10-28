package version

var (
	version string
)

func GetVersion() string {
	if version == "" {
		return "dev"
	} else {
		return version
	}
}