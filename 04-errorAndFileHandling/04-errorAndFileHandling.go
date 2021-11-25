package _0104_errorAndFileHandling_

import (
	"github.com/alcadeta/lrn_golang/04-errorAndFileHandling/01-errorHandling"
	"github.com/alcadeta/lrn_golang/04-errorAndFileHandling/02-errorsPackage"
	"github.com/alcadeta/lrn_golang/04-errorAndFileHandling/03-deferPanicRecover"
	"github.com/alcadeta/lrn_golang/04-errorAndFileHandling/04-basicFileOperations"
	"github.com/alcadeta/lrn_golang/04-errorAndFileHandling/05-readingFiles"
)

func ErrorAndFileHandling() {
	_010401_errorHandling_.ErrorHandling()
	_010401_errorHandling_.CustomErrors()
	_010401_errorHandling_.HandlingErrors()
	_010402_errorsPackage_.IsFunction()
	_010403_deferPanicRecover_.BasicDeferPanicRecover()
	_010403_deferPanicRecover_.HandlingRuntimeErrors()
	_010404_basicFileOperations_.CreateFile()
	_010404_basicFileOperations_.OpenAndCloseFile()
	_010404_basicFileOperations_.RemoveFile()
	_010404_basicFileOperations_.CopyFile()
	_010404_basicFileOperations_.RenameFile()
	_010404_basicFileOperations_.TruncateFile()
	_010404_basicFileOperations_.FileInfo()
	_010405_readingFiles_.ReadingFilesIntoMemory()
	_010405_readingFiles_.ReadingFilesLineByLine()
}
