package utils

var (
	IsFunction = isFunction
	IsChan     = isChan
	IsNil      = isNil
	GetMapKeys = getMapKeys

	IsCompressed = isCompressed
	DeflateData  = deflateData
	InflateData  = inflateData

	WriteFile = writeFile
	ReadFile  = readFile

	SliceContainsString = sliceContainsString

	GetTestEtcd = getTestEtcd
	GetFreePort = getFreePort
)
