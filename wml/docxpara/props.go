package docxpara

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/docxrun"
	"github.com/gomutex/godocx/wml/sections"
)

// Numbering Level Associated Paragraph Properties
type ParagraphProp struct {
	// 1. This element specifies the style ID of the paragraph style which shall be used to format the contents of this paragraph.
	Style *ctypes.CTString `xml:"pStyle,omitempty"`

	// 2. Keep Paragraph With Next Paragraph
	KeepNext *ctypes.OnOff `xml:"keepNext,omitempty"`

	// 3. Keep All Lines On One Page
	KeepLines *ctypes.OnOff `xml:"keepLines,omitempty"`

	// 4. Start Paragraph on Next Page
	PageBreakBefore *ctypes.OnOff `xml:"pageBreakBefore,omitempty"`

	// 5. Text Frame Properties
	FrameProp *FrameProp `xml:"framePr,omitempty"`

	// 6. Allow First/Last Line to Display on a Separate Page
	WindowControl *ctypes.OnOff `xml:"widowControl,omitempty"`

	// 7. Numbering Definition Instance Reference
	NumProp *NumProp `xml:"numPr,omitempty"`

	// 8. Suppress Line Numbers for Paragraph
	SuppressLineNmbrs *ctypes.OnOff `xml:"suppressLineNumbers,omitempty"`

	// 9. Paragraph Borders
	Border *ParaBorder `xml:"pBdr,omitempty"`

	// 10. This element specifies the shading applied to the contents of the paragraph.
	Shading *ctypes.Shading `xml:"shd,omitempty"`

	// 11. Set of Custom Tab Stops
	Tabs ctypes.Tabs `xml:"tabs,omitempty"`

	// 12. Suppress Hyphenation for Paragraph
	SuppressAutoHyphens *ctypes.OnOff `xml:"suppressAutoHyphens,omitempty"`

	// 13. Use East Asian Typography Rules for First and Last Character per Line
	Kinsoku *ctypes.OnOff `xml:"kinsoku,omitempty"`

	// 14. Allow Line Breaking At Character Level
	WordWrap *ctypes.OnOff `xml:"wordWrap,omitempty"`

	// 15. Allow Punctuation to Extent Past Text Extents
	OverflowPunct *ctypes.OnOff `xml:"overflowPunct,omitempty"`

	// 16. Compress Punctuation at Start of a Line
	TopLinePunct *ctypes.OnOff `xml:"topLinePunct,omitempty"`

	// 17. Automatically Adjust Spacing of Latin and East Asian Text
	AutoSpaceDE *ctypes.OnOff `xml:"autoSpaceDE,omitempty"`

	// 18. Automatically Adjust Spacing of East Asian Text and Numbers
	AutoSpaceDN *ctypes.OnOff `xml:"autoSpaceDN,omitempty"`

	// 19. Right to Left Paragraph Layout
	Bidi *ctypes.OnOff `xml:"bidi,omitempty"`

	// 20. Automatically Adjust Right Indent When Using Document Grid
	AdjustRightInd *ctypes.OnOff `xml:"adjustRightInd,omitempty"`

	// 21. Use Document Grid Settings for Inter-Line Paragraph Spacing
	SnapToGrid *ctypes.OnOff `xml:"snapToGrid,omitempty"`

	// 22. Spacing Between Lines and Above/Below Paragraph
	Spacing *Spacing `xml:"spacing,omitempty"`

	// 23. Paragraph Indentation
	Indent *Indent `xml:"ind,omitempty"`

	// 24. Ignore Spacing Above and Below When Using Identical Styles
	CtxlSpacing *ctypes.OnOff `xml:"contextualSpacing,omitempty"`

	// 25. Use Left/Right Indents as Inside/Outside Indents
	MirrorIndents *ctypes.OnOff `xml:"mirrorIndents,omitempty"`

	// 26. Prevent Text Frames From Overlapping
	SuppressOverlap *ctypes.OnOff `xml:"suppressOverlap,omitempty"`

	// 27. Paragraph Alignment
	Justification *ctypes.Justification `xml:"jc,omitempty"`

	// 28. Paragraph Text Flow Direction
	TextDirection *ctypes.TextDirection `xml:"textDirection,omitempty"`

	// 29. Vertical Character Alignment on Line
	TextAlignment *ctypes.TextAlign `xml:"textAlignment,omitempty"`

	// 30.Allow Surrounding Paragraphs to Tight Wrap to Text Box Contents
	TextboxTightWrap *ctypes.TextboxTightWrap `xml:"textboxTightWrap,omitempty"`

	// 31. Associated Outline Level
	OutlineLvl *ctypes.DecimalNum `xml:"outlineLvl,omitempty"`

	// 32. Associated HTML div ID
	DivID *ctypes.DecimalNum `xml:"divId,omitempty"`

	// 33. Paragraph Conditional Formatting
	CnfStyle *ctypes.Cnf `xml:"cnfStyle,omitempty"`

	// 34. Run Properties for the Paragraph Mark
	RunProperty *docxrun.RunProperty `xml:"rPr,omitempty"`

	// 35. Section Properties
	SectPr *sections.SectionProp `xml:"sectPr,omitempty"`

	// 36. Revision Information for Paragraph Properties
	PPrChange *PPrChange `xml:"pPrChange,omitempty"`
}

type binElems struct {
	elem    *ctypes.OnOff
	XMLName string
}

func (pp *ParagraphProp) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	elem := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}

	// Opening <w:pPr> element
	if err = e.EncodeToken(elem); err != nil {
		return err
	}

	// 1. PStyle
	if pp.Style != nil {
		if err = pp.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:pStyle"},
		}); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	bElems1 := []binElems{
		{pp.KeepNext, "w:keepNext"},               //2
		{pp.KeepLines, "w:keepLines"},             //3
		{pp.PageBreakBefore, "w:pageBreakBefore"}, //4
	}

	for _, entry := range bElems1 {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling paragraph property `%s`: %w", entry.XMLName, err)
		}
	}

	// 5. FrameProp
	if pp.FrameProp != nil {
		if err = pp.FrameProp.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("FrameProp: %w", err)
		}
	}

	// 6. WindowControl
	if pp.WindowControl != nil {
		if err = pp.WindowControl.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:widowControl"},
		}); err != nil {
			return fmt.Errorf("WindowControl: %w", err)
		}
	}

	// 7. NumProp
	if pp.NumProp != nil {
		if err = pp.NumProp.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("NumberingProperty: %w", err)
		}
	}

	// 8. SuppressLineNmbrs
	if pp.SuppressLineNmbrs != nil {
		if err = pp.SuppressLineNmbrs.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:suppressLineNumbers"},
		}); err != nil {
			return fmt.Errorf("SuppressLineNmbrs: %w", err)
		}
	}

	// 9.Border
	if pp.Border != nil {
		if err = pp.Border.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Border: %w", err)
		}
	}

	// 10. Shading
	if pp.Shading != nil {
		if err = pp.Shading.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:shd"},
		}); err != nil {
			return fmt.Errorf("TextDirection: %w", err)
		}
	}

	// 11. Tabs
	if err = pp.Tabs.MarshalXML(e, xml.StartElement{}); err != nil {
		return fmt.Errorf("Tabs: %w", err)
	}

	bElems2 := []binElems{
		{pp.SuppressAutoHyphens, "w:suppressAutoHyphens"}, //12
		{pp.Kinsoku, "w:kinsoku"},                         //13
		{pp.WordWrap, "w:wordWrap"},                       //4
		{pp.OverflowPunct, "w:overflowPunct"},             //15
		{pp.TopLinePunct, "w:topLinePunct"},               //16
		{pp.AutoSpaceDE, "w:autoSpaceDE"},                 //17
		{pp.AutoSpaceDN, "w:autoSpaceDN"},                 //18
		{pp.Bidi, "w:bidi"},                               //19
		{pp.AdjustRightInd, "w:adjustRightInd"},           //20
		{pp.SnapToGrid, "w:snapToGrid"},                   //21
	}

	for _, entry := range bElems2 {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling paragraph property `%s`: %w", entry.XMLName, err)
		}
	}

	// 22. Spacing
	if pp.Spacing != nil {
		if err = pp.Spacing.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Spacing: %w", err)
		}
	}

	// 23. Indent
	if pp.Indent != nil {
		if err = pp.Indent.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Indent: %w", err)
		}
	}

	bElems3 := []binElems{
		{pp.CtxlSpacing, "w:contextualSpacing"},   //24
		{pp.MirrorIndents, "w:mirrorIndents"},     //25
		{pp.SuppressOverlap, "w:suppressOverlap"}, //26
	}

	for _, entry := range bElems3 {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling paragraph property `%s`: %w", entry.XMLName, err)
		}
	}

	// 27. Justification
	if pp.Justification != nil {
		if err = pp.Justification.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:jc"},
		}); err != nil {
			return fmt.Errorf("Justification: %w", err)
		}
	}

	// 28. TextDirection
	if pp.TextDirection != nil {
		if err = pp.TextDirection.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("TextDirection: %w", err)
		}
	}

	// 29. TextAlignment
	if pp.TextAlignment != nil {
		if err = pp.TextAlignment.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("TextAlignment: %w", err)
		}
	}

	// 30. TextboxTightWrap
	if pp.TextboxTightWrap != nil {
		if err = pp.TextboxTightWrap.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("TextboxTightWrap: %w", err)
		}
	}

	// 31. OutlineLvl
	if pp.OutlineLvl != nil {
		if err = pp.OutlineLvl.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:outlineLvl"},
		}); err != nil {
			return fmt.Errorf("OutlineLvl: %w", err)
		}
	}

	// 32. DivID
	if pp.DivID != nil {
		if err = pp.DivID.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("DivID: %w", err)
		}
	}

	// 33. CnfStyle
	if pp.CnfStyle != nil {
		if err = pp.CnfStyle.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("CnfStyle: %w", err)
		}
	}

	// 34. RunProperty
	if pp.RunProperty != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		if err = e.EncodeElement(pp.RunProperty, propsElement); err != nil {
			return err
		}
	}

	if pp.SectPr != nil {
		if err = pp.SectPr.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("PPrChange: %w", err)
		}
	}

	//36. PPrChange
	if pp.PPrChange != nil {
		if err = pp.PPrChange.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("PPrChange: %w", err)
		}
	}

	return e.EncodeToken(elem.End())
}

// NewParagraphStyle creates a new ParagraphStyle.
func NewParagraphStyle(val string) *ctypes.CTString {
	return &ctypes.CTString{Val: val}
}

// DefaultParagraphStyle creates the default ParagraphStyle with the value "Normal".
func DefaultParagraphStyle() *ctypes.CTString {
	return &ctypes.CTString{Val: "Normal"}
}

func DefaultParaProperty() *ParagraphProp {
	return &ParagraphProp{}
}
