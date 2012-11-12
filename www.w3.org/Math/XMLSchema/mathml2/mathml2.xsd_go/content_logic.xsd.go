//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/Math/XMLSchema/mathml2/content/logic.xsd
package gopkg_WwwW3OrgMathXMLSchemaMathml2Mathml2Xsd

//	This is an XML Schema module for the logic operators of content MathML.
//	Author: Stéphane Dalmas, INRIA.
import (
)

type TLogicType struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasElem_And struct {
	And TElementaryFunctionsType `xml:"http://www.w3.org/1998/Math/MathML and"`
}

type XsdGoPkgHasElems_And struct {
	Ands []TElementaryFunctionsType `xml:"http://www.w3.org/1998/Math/MathML and"`
}

type XsdGoPkgHasElems_Or struct {
	Ors []*TLogicType `xml:"http://www.w3.org/1998/Math/MathML or"`
}

type XsdGoPkgHasElem_Or struct {
	Or *TLogicType `xml:"http://www.w3.org/1998/Math/MathML or"`
}

type XsdGoPkgHasElems_Xor struct {
	Xors []*TLogicType `xml:"http://www.w3.org/1998/Math/MathML xor"`
}

type XsdGoPkgHasElem_Xor struct {
	Xor *TLogicType `xml:"http://www.w3.org/1998/Math/MathML xor"`
}

type XsdGoPkgHasElem_Not struct {
	Not *TLogicType `xml:"http://www.w3.org/1998/Math/MathML not"`
}

type XsdGoPkgHasElems_Not struct {
	Nots []*TLogicType `xml:"http://www.w3.org/1998/Math/MathML not"`
}

type XsdGoPkgHasElems_Exists struct {
	Existses []*TLogicType `xml:"http://www.w3.org/1998/Math/MathML exists"`
}

type XsdGoPkgHasElem_Exists struct {
	Exists *TLogicType `xml:"http://www.w3.org/1998/Math/MathML exists"`
}

type XsdGoPkgHasElem_Forall struct {
	Forall *TLogicType `xml:"http://www.w3.org/1998/Math/MathML forall"`
}

type XsdGoPkgHasElems_Forall struct {
	Foralls []*TLogicType `xml:"http://www.w3.org/1998/Math/MathML forall"`
}

type XsdGoPkgHasElem_Implies struct {
	Implies *TLogicType `xml:"http://www.w3.org/1998/Math/MathML implies"`
}

type XsdGoPkgHasElems_Implies struct {
	Implieses []*TLogicType `xml:"http://www.w3.org/1998/Math/MathML implies"`
}

type XsdGoPkgHasGroup_ContentLogicClass struct {
	XsdGoPkgHasElem_And

	XsdGoPkgHasElem_Or

	XsdGoPkgHasElem_Xor

	XsdGoPkgHasElem_Not

	XsdGoPkgHasElem_Exists

	XsdGoPkgHasElem_Forall

	XsdGoPkgHasElem_Implies

}
