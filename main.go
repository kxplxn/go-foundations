package main

import (
	functionsAndMethods "github.com/alcadeta/lrn_golang/01-foundations/03-functionsAndObjectOrientation/01-functionsAndMethods"
	objectOrientation "github.com/alcadeta/lrn_golang/01-foundations/03-functionsAndObjectOrientation/02-objectOrientation"
	parameters "github.com/alcadeta/lrn_golang/01-foundations/03-functionsAndObjectOrientation/04-parameters"
	returnValues "github.com/alcadeta/lrn_golang/01-foundations/03-functionsAndObjectOrientation/05-returnValues"
	variadicFunctions "github.com/alcadeta/lrn_golang/01-foundations/03-functionsAndObjectOrientation/06-variadicFunctions"
	structs "github.com/alcadeta/lrn_golang/01-foundations/03-functionsAndObjectOrientation/07-structs"
	methods "github.com/alcadeta/lrn_golang/01-foundations/03-functionsAndObjectOrientation/08-methods"
)

func functionsAndObjectOrientation() {
	functionsAndMethods.Functions()
	functionsAndMethods.Methods()
	objectOrientation.BasicObjects()
	objectOrientation.ObjectComposition()
	objectOrientation.AnonymousFields()
	objectOrientation.Interfaces()
	parameters.BasicParameters()
	parameters.SliceParameters()
	parameters.PointerArguments()
	returnValues.SingleReturnValue()
	returnValues.MultipleReturnValues()
	returnValues.NamedReturnValues()
	variadicFunctions.ThePackOperator()
	variadicFunctions.TheUnpackOperator()
	structs.StructDefinition()
	structs.StructLiterals()
	structs.StructReferences()
	structs.StructPointers()
	methods.DeclaringAndCallingMethods()
}

func main() {
	functionsAndObjectOrientation()
}
