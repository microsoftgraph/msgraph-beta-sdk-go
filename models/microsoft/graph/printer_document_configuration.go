package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterDocumentConfiguration 
type PrinterDocumentConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    collate *bool;
    // 
    colorMode *PrintColorMode;
    // 
    copies *int32;
    // 
    dpi *int32;
    // 
    duplexMode *PrintDuplexMode;
    // 
    feedDirection *PrinterFeedDirection;
    // 
    feedOrientation *PrinterFeedOrientation;
    // 
    finishings []PrintFinishing;
    // 
    fitPdfToPage *bool;
    // 
    inputBin *string;
    // 
    margin *PrintMargin;
    // 
    mediaSize *string;
    // 
    mediaType *string;
    // 
    multipageLayout *PrintMultipageLayout;
    // 
    orientation *PrintOrientation;
    // 
    outputBin *string;
    // 
    pageRanges []IntegerRange;
    // 
    pagesPerSheet *int32;
    // 
    quality *PrintQuality;
    // 
    scaling *PrintScaling;
}
// NewPrinterDocumentConfiguration instantiates a new printerDocumentConfiguration and sets the default values.
func NewPrinterDocumentConfiguration()(*PrinterDocumentConfiguration) {
    m := &PrinterDocumentConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterDocumentConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCollate gets the collate property value. 
func (m *PrinterDocumentConfiguration) GetCollate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.collate
    }
}
// GetColorMode gets the colorMode property value. 
func (m *PrinterDocumentConfiguration) GetColorMode()(*PrintColorMode) {
    if m == nil {
        return nil
    } else {
        return m.colorMode
    }
}
// GetCopies gets the copies property value. 
func (m *PrinterDocumentConfiguration) GetCopies()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.copies
    }
}
// GetDpi gets the dpi property value. 
func (m *PrinterDocumentConfiguration) GetDpi()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dpi
    }
}
// GetDuplexMode gets the duplexMode property value. 
func (m *PrinterDocumentConfiguration) GetDuplexMode()(*PrintDuplexMode) {
    if m == nil {
        return nil
    } else {
        return m.duplexMode
    }
}
// GetFeedDirection gets the feedDirection property value. 
func (m *PrinterDocumentConfiguration) GetFeedDirection()(*PrinterFeedDirection) {
    if m == nil {
        return nil
    } else {
        return m.feedDirection
    }
}
// GetFeedOrientation gets the feedOrientation property value. 
func (m *PrinterDocumentConfiguration) GetFeedOrientation()(*PrinterFeedOrientation) {
    if m == nil {
        return nil
    } else {
        return m.feedOrientation
    }
}
// GetFinishings gets the finishings property value. 
func (m *PrinterDocumentConfiguration) GetFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.finishings
    }
}
// GetFitPdfToPage gets the fitPdfToPage property value. 
func (m *PrinterDocumentConfiguration) GetFitPdfToPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fitPdfToPage
    }
}
// GetInputBin gets the inputBin property value. 
func (m *PrinterDocumentConfiguration) GetInputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inputBin
    }
}
// GetMargin gets the margin property value. 
func (m *PrinterDocumentConfiguration) GetMargin()(*PrintMargin) {
    if m == nil {
        return nil
    } else {
        return m.margin
    }
}
// GetMediaSize gets the mediaSize property value. 
func (m *PrinterDocumentConfiguration) GetMediaSize()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaSize
    }
}
// GetMediaType gets the mediaType property value. 
func (m *PrinterDocumentConfiguration) GetMediaType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
// GetMultipageLayout gets the multipageLayout property value. 
func (m *PrinterDocumentConfiguration) GetMultipageLayout()(*PrintMultipageLayout) {
    if m == nil {
        return nil
    } else {
        return m.multipageLayout
    }
}
// GetOrientation gets the orientation property value. 
func (m *PrinterDocumentConfiguration) GetOrientation()(*PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.orientation
    }
}
// GetOutputBin gets the outputBin property value. 
func (m *PrinterDocumentConfiguration) GetOutputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputBin
    }
}
// GetPageRanges gets the pageRanges property value. 
func (m *PrinterDocumentConfiguration) GetPageRanges()([]IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.pageRanges
    }
}
// GetPagesPerSheet gets the pagesPerSheet property value. 
func (m *PrinterDocumentConfiguration) GetPagesPerSheet()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pagesPerSheet
    }
}
// GetQuality gets the quality property value. 
func (m *PrinterDocumentConfiguration) GetQuality()(*PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.quality
    }
}
// GetScaling gets the scaling property value. 
func (m *PrinterDocumentConfiguration) GetScaling()(*PrintScaling) {
    if m == nil {
        return nil
    } else {
        return m.scaling
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterDocumentConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["collate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollate(val)
        }
        return nil
    }
    res["colorMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintColorMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColorMode(val.(*PrintColorMode))
        }
        return nil
    }
    res["copies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopies(val)
        }
        return nil
    }
    res["dpi"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDpi(val)
        }
        return nil
    }
    res["duplexMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintDuplexMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuplexMode(val.(*PrintDuplexMode))
        }
        return nil
    }
    res["feedDirection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterFeedDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedDirection(val.(*PrinterFeedDirection))
        }
        return nil
    }
    res["feedOrientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterFeedOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedOrientation(val.(*PrinterFeedOrientation))
        }
        return nil
    }
    res["finishings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintFinishing)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintFinishing, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintFinishing))
            }
            m.SetFinishings(res)
        }
        return nil
    }
    res["fitPdfToPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFitPdfToPage(val)
        }
        return nil
    }
    res["inputBin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInputBin(val)
        }
        return nil
    }
    res["margin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintMargin() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMargin(val.(*PrintMargin))
        }
        return nil
    }
    res["mediaSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaSize(val)
        }
        return nil
    }
    res["mediaType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val)
        }
        return nil
    }
    res["multipageLayout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintMultipageLayout)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultipageLayout(val.(*PrintMultipageLayout))
        }
        return nil
    }
    res["orientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrientation(val.(*PrintOrientation))
        }
        return nil
    }
    res["outputBin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputBin(val)
        }
        return nil
    }
    res["pageRanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntegerRange, len(val))
            for i, v := range val {
                res[i] = *(v.(*IntegerRange))
            }
            m.SetPageRanges(res)
        }
        return nil
    }
    res["pagesPerSheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPagesPerSheet(val)
        }
        return nil
    }
    res["quality"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintQuality)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuality(val.(*PrintQuality))
        }
        return nil
    }
    res["scaling"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintScaling)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScaling(val.(*PrintScaling))
        }
        return nil
    }
    return res
}
func (m *PrinterDocumentConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrinterDocumentConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("collate", m.GetCollate())
        if err != nil {
            return err
        }
    }
    if m.GetColorMode() != nil {
        cast := (*m.GetColorMode()).String()
        err := writer.WriteStringValue("colorMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("copies", m.GetCopies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dpi", m.GetDpi())
        if err != nil {
            return err
        }
    }
    if m.GetDuplexMode() != nil {
        cast := (*m.GetDuplexMode()).String()
        err := writer.WriteStringValue("duplexMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFeedDirection() != nil {
        cast := (*m.GetFeedDirection()).String()
        err := writer.WriteStringValue("feedDirection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFeedOrientation() != nil {
        cast := (*m.GetFeedOrientation()).String()
        err := writer.WriteStringValue("feedOrientation", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFinishings() != nil {
        err := writer.WriteCollectionOfStringValues("finishings", SerializePrintFinishing(m.GetFinishings()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("fitPdfToPage", m.GetFitPdfToPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("inputBin", m.GetInputBin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("margin", m.GetMargin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaSize", m.GetMediaSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    if m.GetMultipageLayout() != nil {
        cast := (*m.GetMultipageLayout()).String()
        err := writer.WriteStringValue("multipageLayout", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrientation() != nil {
        cast := (*m.GetOrientation()).String()
        err := writer.WriteStringValue("orientation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputBin", m.GetOutputBin())
        if err != nil {
            return err
        }
    }
    if m.GetPageRanges() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPageRanges()))
        for i, v := range m.GetPageRanges() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("pageRanges", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pagesPerSheet", m.GetPagesPerSheet())
        if err != nil {
            return err
        }
    }
    if m.GetQuality() != nil {
        cast := (*m.GetQuality()).String()
        err := writer.WriteStringValue("quality", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScaling() != nil {
        cast := (*m.GetScaling()).String()
        err := writer.WriteStringValue("scaling", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterDocumentConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCollate sets the collate property value. 
func (m *PrinterDocumentConfiguration) SetCollate(value *bool)() {
    if m != nil {
        m.collate = value
    }
}
// SetColorMode sets the colorMode property value. 
func (m *PrinterDocumentConfiguration) SetColorMode(value *PrintColorMode)() {
    if m != nil {
        m.colorMode = value
    }
}
// SetCopies sets the copies property value. 
func (m *PrinterDocumentConfiguration) SetCopies(value *int32)() {
    if m != nil {
        m.copies = value
    }
}
// SetDpi sets the dpi property value. 
func (m *PrinterDocumentConfiguration) SetDpi(value *int32)() {
    if m != nil {
        m.dpi = value
    }
}
// SetDuplexMode sets the duplexMode property value. 
func (m *PrinterDocumentConfiguration) SetDuplexMode(value *PrintDuplexMode)() {
    if m != nil {
        m.duplexMode = value
    }
}
// SetFeedDirection sets the feedDirection property value. 
func (m *PrinterDocumentConfiguration) SetFeedDirection(value *PrinterFeedDirection)() {
    if m != nil {
        m.feedDirection = value
    }
}
// SetFeedOrientation sets the feedOrientation property value. 
func (m *PrinterDocumentConfiguration) SetFeedOrientation(value *PrinterFeedOrientation)() {
    if m != nil {
        m.feedOrientation = value
    }
}
// SetFinishings sets the finishings property value. 
func (m *PrinterDocumentConfiguration) SetFinishings(value []PrintFinishing)() {
    if m != nil {
        m.finishings = value
    }
}
// SetFitPdfToPage sets the fitPdfToPage property value. 
func (m *PrinterDocumentConfiguration) SetFitPdfToPage(value *bool)() {
    if m != nil {
        m.fitPdfToPage = value
    }
}
// SetInputBin sets the inputBin property value. 
func (m *PrinterDocumentConfiguration) SetInputBin(value *string)() {
    if m != nil {
        m.inputBin = value
    }
}
// SetMargin sets the margin property value. 
func (m *PrinterDocumentConfiguration) SetMargin(value *PrintMargin)() {
    if m != nil {
        m.margin = value
    }
}
// SetMediaSize sets the mediaSize property value. 
func (m *PrinterDocumentConfiguration) SetMediaSize(value *string)() {
    if m != nil {
        m.mediaSize = value
    }
}
// SetMediaType sets the mediaType property value. 
func (m *PrinterDocumentConfiguration) SetMediaType(value *string)() {
    if m != nil {
        m.mediaType = value
    }
}
// SetMultipageLayout sets the multipageLayout property value. 
func (m *PrinterDocumentConfiguration) SetMultipageLayout(value *PrintMultipageLayout)() {
    if m != nil {
        m.multipageLayout = value
    }
}
// SetOrientation sets the orientation property value. 
func (m *PrinterDocumentConfiguration) SetOrientation(value *PrintOrientation)() {
    if m != nil {
        m.orientation = value
    }
}
// SetOutputBin sets the outputBin property value. 
func (m *PrinterDocumentConfiguration) SetOutputBin(value *string)() {
    if m != nil {
        m.outputBin = value
    }
}
// SetPageRanges sets the pageRanges property value. 
func (m *PrinterDocumentConfiguration) SetPageRanges(value []IntegerRange)() {
    if m != nil {
        m.pageRanges = value
    }
}
// SetPagesPerSheet sets the pagesPerSheet property value. 
func (m *PrinterDocumentConfiguration) SetPagesPerSheet(value *int32)() {
    if m != nil {
        m.pagesPerSheet = value
    }
}
// SetQuality sets the quality property value. 
func (m *PrinterDocumentConfiguration) SetQuality(value *PrintQuality)() {
    if m != nil {
        m.quality = value
    }
}
// SetScaling sets the scaling property value. 
func (m *PrinterDocumentConfiguration) SetScaling(value *PrintScaling)() {
    if m != nil {
        m.scaling = value
    }
}
