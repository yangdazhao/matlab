package goMatlab

type MatError int

//go:generate stringer -type MatError -trimprefix mat
const (
	matNoError MatError = iota
	matUnknownError
	matGenericReadError
	matGenericWriteError
	matIndexTooBig
	matFileFormatViolation
	matFailToIdentify
	matBadArgument
	matOutputBadData
	matFullObjectOutputConvert
	matPartObjectOutputConvert
	matFullObjectInputConvert
	matPartObjectInputConvert
	matOperationNotSupported
	matOutOfMemory
	matBadVariableName
	matOperationProhibitedInWriteMode
	matOperationProhibitedInReadMode
	matWriteVariableDoesNotExist
	matReadVariableDoesNotExist
	matFilesystemCouldNotOpen
	matFilesystemCouldNotOpenTemporary
	matFilesystemCouldNotReopen
	matBadOpenMode
	matFilesystemErrorOnClose
)
