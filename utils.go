package utils

var (
	IsFunction              = isFunction
	IsChan                  = isChan
	IsNil                   = isNil
	GetMapKeys              = getMapKeys
	ShouldAlwaysReturn      = shouldAlwaysReturn
	ShouldEventuallyReturn  = shouldEventuallyReturn
	ShouldEventuallyReceive = shouldEventuallyReceive
	PollFuncReturn          = pollFuncReturn
	VetExtras               = vetExtras
	FixtureGoldenFileName   = fixtureGoldenFileName
	WaitForServerToBeReady  = waitForServerToBeReady
	StartProcess            = startProcess

	IsCompressed = isCompressed
	DeflateData  = deflateData
	InflateData  = inflateData

	WriteFile = writeFile
	ReadFile  = readFile

	SliceContainsString = sliceContainsString

	GetTestEtcd = getTestEtcd
	GetFreePort = getFreePort
)
