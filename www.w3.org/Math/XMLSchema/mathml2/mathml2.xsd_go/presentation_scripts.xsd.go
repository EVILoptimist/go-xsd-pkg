//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/Math/XMLSchema/mathml2/presentation/scripts.xsd
package gopkg_WwwW3OrgMathXMLSchemaMathml2Mathml2Xsd

//	This is an XML Schema module for the presentation elements of MathML
//	dealing with subscripts and superscripts.
//	Author: Stéphane Dalmas, INRIA.
import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type XsdGoPkgHasAttr_Subscriptshift_TlengthWithUnit_ struct {
	Subscriptshift TlengthWithUnit `xml:"http://www.w3.org/1998/Math/MathML subscriptshift,attr"`
}

type XsdGoPkgHasAtts_MsubAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

	XsdGoPkgHasAttr_Subscriptshift_TlengthWithUnit_
}

type XsdGoPkgHasAttr_Superscriptshift_TlengthWithUnit_ struct {
	Superscriptshift TlengthWithUnit `xml:"http://www.w3.org/1998/Math/MathML superscriptshift,attr"`
}

type XsdGoPkgHasAtts_MsupAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

	XsdGoPkgHasAttr_Superscriptshift_TlengthWithUnit_
}

type XsdGoPkgHasAtts_MsubsupAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAttr_Accentunder_XsdtBoolean_ struct {
	Accentunder xsdt.Boolean `xml:"http://www.w3.org/1998/Math/MathML accentunder,attr"`
}

type XsdGoPkgHasAtts_MunderAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

	XsdGoPkgHasAttr_Accentunder_XsdtBoolean_
}

type XsdGoPkgHasAtts_MoverAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_MunderoverAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_MmultiscriptsAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

}

type TmsubType struct {
	XsdGoPkgHasGroup_PresentationExprClass

	XsdGoPkgHasAtts_MsubAttlist

}

type TmsupType struct {
	XsdGoPkgHasGroup_PresentationExprClass

	XsdGoPkgHasAtts_MsupAttlist

}

type TmsubsupType struct {
	XsdGoPkgHasGroup_PresentationExprClass

	XsdGoPkgHasAtts_MsubsupAttlist

}

type TmunderType struct {
	XsdGoPkgHasGroup_PresentationExprClass

	XsdGoPkgHasAtts_MunderAttlist

}

type TmoverType struct {
	XsdGoPkgHasGroup_PresentationExprClass

	XsdGoPkgHasAtts_MoverAttlist

}

type TmunderoverType struct {
	XsdGoPkgHasGroup_PresentationExprClass

	XsdGoPkgHasAtts_MunderoverAttlist

}

type TmmultiscriptsType struct {
	XsdGoPkgHasGroup_MmultiscriptsContent

	XsdGoPkgHasAtts_MmultiscriptsAttlist

}

type TnoneType struct {
}

type TmprescriptsType struct {
}

type XsdGoPkgHasElems_Msub struct {
	Msubs []*TmsubType `xml:"http://www.w3.org/1998/Math/MathML msub"`
}

type XsdGoPkgHasElem_Msub struct {
	Msub *TmsubType `xml:"http://www.w3.org/1998/Math/MathML msub"`
}

type XsdGoPkgHasElem_Msup struct {
	Msup *TmsupType `xml:"http://www.w3.org/1998/Math/MathML msup"`
}

type XsdGoPkgHasElems_Msup struct {
	Msups []*TmsupType `xml:"http://www.w3.org/1998/Math/MathML msup"`
}

type XsdGoPkgHasElem_Msubsup struct {
	Msubsup *TmsubsupType `xml:"http://www.w3.org/1998/Math/MathML msubsup"`
}

type XsdGoPkgHasElems_Msubsup struct {
	Msubsups []*TmsubsupType `xml:"http://www.w3.org/1998/Math/MathML msubsup"`
}

type XsdGoPkgHasElem_Munder struct {
	Munder *TmunderType `xml:"http://www.w3.org/1998/Math/MathML munder"`
}

type XsdGoPkgHasElems_Munder struct {
	Munders []*TmunderType `xml:"http://www.w3.org/1998/Math/MathML munder"`
}

type XsdGoPkgHasElems_Mover struct {
	Movers []*TmoverType `xml:"http://www.w3.org/1998/Math/MathML mover"`
}

type XsdGoPkgHasElem_Mover struct {
	Mover *TmoverType `xml:"http://www.w3.org/1998/Math/MathML mover"`
}

type XsdGoPkgHasElems_Munderover struct {
	Munderovers []*TmunderoverType `xml:"http://www.w3.org/1998/Math/MathML munderover"`
}

type XsdGoPkgHasElem_Munderover struct {
	Munderover *TmunderoverType `xml:"http://www.w3.org/1998/Math/MathML munderover"`
}

type XsdGoPkgHasElem_Mmultiscripts struct {
	Mmultiscripts *TmmultiscriptsType `xml:"http://www.w3.org/1998/Math/MathML mmultiscripts"`
}

type XsdGoPkgHasElems_Mmultiscripts struct {
	Mmultiscriptses []*TmmultiscriptsType `xml:"http://www.w3.org/1998/Math/MathML mmultiscripts"`
}

type XsdGoPkgHasElems_None struct {
	Nones []*TnoneType `xml:"http://www.w3.org/1998/Math/MathML none"`
}

type XsdGoPkgHasElem_None struct {
	None *TnoneType `xml:"http://www.w3.org/1998/Math/MathML none"`
}

type XsdGoPkgHasElems_Mprescripts struct {
	Mprescriptses []*TmprescriptsType `xml:"http://www.w3.org/1998/Math/MathML mprescripts"`
}

type XsdGoPkgHasElem_Mprescripts struct {
	Mprescripts *TmprescriptsType `xml:"http://www.w3.org/1998/Math/MathML mprescripts"`
}

type XsdGoPkgHasGroup_PresentationExprOrNoneClass struct {
	XsdGoPkgHasElem_None

	XsdGoPkgHasGroup_PresentationExprClass

}

type XsdGoPkgHasGroup_MmultiscriptsContent struct {
	XsdGoPkgHasGroup_PresentationExprClass

	XsdGoPkgHasGroup_PresentationExprOrNoneClass

	XsdGoPkgHasElem_Mprescripts

}

type XsdGoPkgHasGroup_PresentationScriptClass struct {
	XsdGoPkgHasElem_Msub

	XsdGoPkgHasElem_Msup

	XsdGoPkgHasElem_Msubsup

	XsdGoPkgHasElem_Munder

	XsdGoPkgHasElem_Mover

	XsdGoPkgHasElem_Munderover

	XsdGoPkgHasElem_Mmultiscripts

}
