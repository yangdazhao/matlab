// Code generated by "stringer -type MatError -trimprefix mat"; DO NOT EDIT.

package goMatlab

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[matNoError-0]
	_ = x[matUnknownError-1]
	_ = x[matGenericReadError-2]
	_ = x[matGenericWriteError-3]
	_ = x[matIndexTooBig-4]
	_ = x[matFileFormatViolation-5]
	_ = x[matFailToIdentify-6]
	_ = x[matBadArgument-7]
	_ = x[matOutputBadData-8]
	_ = x[matFullObjectOutputConvert-9]
	_ = x[matPartObjectOutputConvert-10]
	_ = x[matFullObjectInputConvert-11]
	_ = x[matPartObjectInputConvert-12]
	_ = x[matOperationNotSupported-13]
	_ = x[matOutOfMemory-14]
	_ = x[matBadVariableName-15]
	_ = x[matOperationProhibitedInWriteMode-16]
	_ = x[matOperationProhibitedInReadMode-17]
	_ = x[matWriteVariableDoesNotExist-18]
	_ = x[matReadVariableDoesNotExist-19]
	_ = x[matFilesystemCouldNotOpen-20]
	_ = x[matFilesystemCouldNotOpenTemporary-21]
	_ = x[matFilesystemCouldNotReopen-22]
	_ = x[matBadOpenMode-23]
	_ = x[matFilesystemErrorOnClose-24]
}

const _MatError_name = "NoErrorUnknownErrorGenericReadErrorGenericWriteErrorIndexTooBigFileFormatViolationFailToIdentifyBadArgumentOutputBadDataFullObjectOutputConvertPartObjectOutputConvertFullObjectInputConvertPartObjectInputConvertOperationNotSupportedOutOfMemoryBadVariableNameOperationProhibitedInWriteModeOperationProhibitedInReadModeWriteVariableDoesNotExistReadVariableDoesNotExistFilesystemCouldNotOpenFilesystemCouldNotOpenTemporaryFilesystemCouldNotReopenBadOpenModeFilesystemErrorOnClose"

var _MatError_index = [...]uint16{0, 7, 19, 35, 52, 63, 82, 96, 107, 120, 143, 166, 188, 210, 231, 242, 257, 287, 316, 341, 365, 387, 418, 442, 453, 475}

func (i MatError) String() string {
	if i < 0 || i >= MatError(len(_MatError_index)-1) {
		return "MatError(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MatError_name[_MatError_index[i]:_MatError_index[i+1]]
}
