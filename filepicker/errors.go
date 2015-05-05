package filepicker

import "strconv"

// FPError represents an error which could be returned from filepicker.io service.
type FPError int

// Error satisfies builtin.error interface. It prints the error string with
// the reason of failure.
func (e FPError) Error() string {
	var prefix = `filepicker: ` + strconv.Itoa(int(e)) + ` - `
	if msg, ok := errmsgs[e]; ok {
		return prefix + msg
	}
	return prefix + `connection error`
}

const (
	ErrOtherDomainsCantRead  FPError = 113
	ErrRequestWebsiteFailed  FPError = 114
	ErrFileNotFound          FPError = 115
	ErrGeneralReadError      FPError = 118
	ErrCannotFindWriteBlob   FPError = 121
	ErrWriteUrlUnreachable   FPError = 122
	ErrCannotFindInputFile   FPError = 141
	ErrFileCannotBeConverted FPError = 142
	ErrConvertUnknownFailure FPError = 143
	ErrFileStoreUnreachable  FPError = 151
	ErrRemoteUrlUnreachable  FPError = 152
	ErrStatFileNotFound      FPError = 161
	ErrStatFetchingMetadata  FPError = 162
	ErrRmFileCannotBeFound   FPError = 171
	ErrRmStoreUnreachable    FPError = 172
	ErrBadParameters         FPError = 400
	ErrInvalidRequest        FPError = 403
)

var errmsgs = map[FPError]string{
	ErrOtherDomainsCantRead:  `requested website does not allow other domains to read data`,
	ErrRequestWebsiteFailed:  `requested website  had an error`,
	ErrFileNotFound:          `file not found`,
	ErrGeneralReadError:      `general read error`,
	ErrCannotFindWriteBlob:   `Blob to write to could not be found`,
	ErrWriteUrlUnreachable:   `remote URL could not be reached`,
	ErrCannotFindInputFile:   `inputted file cannot be found.`,
	ErrFileCannotBeConverted: `file cannot be converted with the requested parameters.`,
	ErrConvertUnknownFailure: `unknown error when converting the file.`,
	ErrFileStoreUnreachable:  `file store could not be reached`,
	ErrRemoteUrlUnreachable:  `remote URL could not be reached`,
	ErrStatFileNotFound:      `fetching metadata of non existing file`,
	ErrStatFetchingMetadata:  `error fetching metadata`,
	ErrRmFileCannotBeFound:   `file cannot be found, and may have already been deleted`,
	ErrRmStoreUnreachable:    `underlying content store could not be reached`,
	ErrBadParameters:         `bad parameters were passed to the server`,
	ErrInvalidRequest:        `invalid request`,
}
