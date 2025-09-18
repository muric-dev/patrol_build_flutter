package build_constants

const (
	Version      = "FLUTTER_BUILD_PATROL_VERSION"
	Platform     = "FLUTTER_BUILD_PATROL_PLATFORM"
	Target       = "FLUTTER_BUILD_PATROL_TARGET" // comma-separated
	BuildType    = "FLUTTER_BUILD_PATROL_BUILD_TYPE"
	Tags         = "FLUTTER_BUILD_PATROL_TAGS"          // optional, comma-separated
	ExcludedTags = "FLUTTER_BUILD_PATROL_EXCLUDED_TAGS" // optional, comma-separated
	IsVerbose    = "FLUTTER_BUILD_PATROL_IS_VERBOSE"
	IsCovered    = "FLUTTER_BUILD_PATROL_IS_COVERED"
	FilePath     = "TARGET_DIRECTORY_PATH"
)
