package _0104_errorAndFileHandling_

import (
	errorHandling "github.com/alcadeta/lrn_golang/01-foundations/04-errorAndFileHandling/01-errorHandling"
	errorsPackage "github.com/alcadeta/lrn_golang/01-foundations/04-errorAndFileHandling/02-errorsPackage"
	deferPanicRecover "github.com/alcadeta/lrn_golang/01-foundations/04-errorAndFileHandling/03-deferPanicRecover"
)

func ErrorAndFileHandling() {
	errorHandling.ErrorHandling()
	errorHandling.CustomErrors()
	errorHandling.HandlingErrors()
	errorsPackage.IsFunction()
	deferPanicRecover.BasicDeferPanicRecover()
	deferPanicRecover.HandlingRuntimeErrors()
}
