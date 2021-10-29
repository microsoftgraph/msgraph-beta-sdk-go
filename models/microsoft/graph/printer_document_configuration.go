package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new printerDocumentConfiguration and sets the default values.
func NewPrinterDocumentConfiguration()(*PrinterDocumentConfiguration) {
    m := &PrinterDocumentConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterDocumentConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the collate property value. 
func (m *PrinterDocumentConfiguration) GetCollate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.collate
    }
}
// Gets the colorMode property value. 
func (m *PrinterDocumentConfiguration) GetColorMode()(*PrintColorMode) {
    if m == nil {
        return nil
    } else {
        return m.colorMode
    }
}
// Gets the copies property value. 
func (m *PrinterDocumentConfiguration) GetCopies()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.copies
    }
}
// Gets the dpi property value. 
func (m *PrinterDocumentConfiguration) GetDpi()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dpi
    }
}
// Gets the duplexMode property value. 
func (m *PrinterDocumentConfiguration) GetDuplexMode()(*PrintDuplexMode) {
    if m == nil {
        return nil
    } else {
        return m.duplexMode
    }
}
// Gets the feedDirection property value. 
func (m *PrinterDocumentConfiguration) GetFeedDirection()(*PrinterFeedDirection) {
    if m == nil {
        return nil
    } else {
        return m.feedDirection
    }
}
// Gets the feedOrientation property value. 
func (m *PrinterDocumentConfiguration) GetFeedOrientation()(*PrinterFeedOrientation) {
    if m == nil {
        return nil
    } else {
        return m.feedOrientation
    }
}
// Gets the finishings property value. 
func (m *PrinterDocumentConfiguration) GetFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.finishings
    }
}
// Gets the fitPdfToPage property value. 
func (m *PrinterDocumentConfiguration) GetFitPdfToPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fitPdfToPage
    }
}
// Gets the inputBin property value. 
func (m *PrinterDocumentConfiguration) GetInputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inputBin
    }
}
// Gets the margin property value. 
func (m *PrinterDocumentConfiguration) GetMargin()(*PrintMargin) {
    if m == nil {
        return nil
    } else {
        return m.margin
    }
}
// Gets the mediaSize property value. 
func (m *PrinterDocumentConfiguration) GetMediaSize()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaSize
    }
}
// Gets the mediaType property value. 
func (m *PrinterDocumentConfiguration) GetMediaType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
// Gets the multipageLayout property value. 
func (m *PrinterDocumentConfiguration) GetMultipageLayout()(*PrintMultipageLayout) {
    if m == nil {
        return nil
    } else {
        return m.multipageLayout
    }
}
// Gets the orientation property value. 
func (m *PrinterDocumentConfiguration) GetOrientation()(*PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.orientation
    }
}
// Gets the outputBin property value. 
func (m *PrinterDocumentConfiguration) GetOutputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputBin
    }
}
// Gets the pageRanges property value. 
func (m *PrinterDocumentConfiguration) GetPageRanges()([]IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.pageRanges
    }
}
// Gets the pagesPerSheet property value. 
func (m *PrinterDocumentConfiguration) GetPagesPerSheet()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pagesPerSheet
    }
}
// Gets the quality property value. 
func (m *PrinterDocumentConfiguration) GetQuality()(*PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.quality
    }
}
// Gets the scaling property value. 
func (m *PrinterDocumentConfiguration) GetScaling()(*PrintScaling) {
    if m == nil {
        return nil
    } else {
        return m.scaling
    }
}
// The deserialization information for the current model
func (m *PrinterDocumentConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["collate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCollate(val)
        return nil
    }
    res["colorMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintColorMode)
        if err != nil {
            return err
        }
        cast := val.(PrintColorMode)
        m.SetColorMode(&cast)
        return nil
    }
    res["copies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCopies(val)
        return nil
    }
    res["dpi"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDpi(val)
        return nil
    }
    res["duplexMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintDuplexMode)
        if err != nil {
            return err
        }
        cast := val.(PrintDuplexMode)
        m.SetDuplexMode(&cast)
        return nil
    }
    res["feedDirection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterFeedDirection)
        if err != nil {
            return err
        }
        cast := val.(PrinterFeedDirection)
        m.SetFeedDirection(&cast)
        return nil
    }
    res["feedOrientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterFeedOrientation)
        if err != nil {
            return err
        }
        cast := val.(PrinterFeedOrientation)
        m.SetFeedOrientation(&cast)
        return nil
    }
    res["finishings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintFinishing)
        if err != nil {
            return err
        }
        res := make([]PrintFinishing, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintFinishing))
        }
        m.SetFinishings(res)
        return nil
    }
    res["fitPdfToPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFitPdfToPage(val)
        return nil
    }
    res["inputBin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInputBin(val)
        return nil
    }
    res["margin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintMargin() })
        if err != nil {
            return err
        }
        m.SetMargin(val.(*PrintMargin))
        return nil
    }
    res["mediaSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMediaSize(val)
        return nil
    }
    res["mediaType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMediaType(val)
        return nil
    }
    res["multipageLayout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintMultipageLayout)
        if err != nil {
            return err
        }
        cast := val.(PrintMultipageLayout)
        m.SetMultipageLayout(&cast)
        return nil
    }
    res["orientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintOrientation)
        if err != nil {
            return err
        }
        cast := val.(PrintOrientation)
        m.SetOrientation(&cast)
        return nil
    }
    res["outputBin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOutputBin(val)
        return nil
    }
    res["pageRanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        res := make([]IntegerRange, len(val))
        for i, v := range val {
            res[i] = *(v.(*IntegerRange))
        }
        m.SetPageRanges(res)
        return nil
    }
    res["pagesPerSheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPagesPerSheet(val)
        return nil
    }
    res["quality"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintQuality)
        if err != nil {
            return err
        }
        cast := val.(PrintQuality)
        m.SetQuality(&cast)
        return nil
    }
    res["scaling"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintScaling)
        if err != nil {
            return err
        }
        cast := val.(PrintScaling)
        m.SetScaling(&cast)
        return nil
    }
    return res
}
func (m *PrinterDocumentConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrinterDocumentConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("collate", m.GetCollate())
        if err != nil {
            return err
        }
    }
    if m.GetColorMode() != nil {
        cast := m.GetColorMode().String()
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
        cast := m.GetDuplexMode().String()
        err := writer.WriteStringValue("duplexMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFeedDirection() != nil {
        cast := m.GetFeedDirection().String()
        err := writer.WriteStringValue("feedDirection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFeedOrientation() != nil {
        cast := m.GetFeedOrientation().String()
        err := writer.WriteStringValue("feedOrientation", &cast)
        if err != nil {
            return err
        }
    }
    {
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
        cast := m.GetMultipageLayout().String()
        err := writer.WriteStringValue("multipageLayout", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrientation() != nil {
        cast := m.GetOrientation().String()
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
    {
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
        cast := m.GetQuality().String()
        err := writer.WriteStringValue("quality", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScaling() != nil {
        cast := m.GetScaling().String()
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PrinterDocumentConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the collate property value. 
// Parameters:
//  - value : Value to set for the collate property.
func (m *PrinterDocumentConfiguration) SetCollate(value *bool)() {
    m.collate = value
}
// Sets the colorMode property value. 
// Parameters:
//  - value : Value to set for the colorMode property.
func (m *PrinterDocumentConfiguration) SetColorMode(value *PrintColorMode)() {
    m.colorMode = value
}
// Sets the copies property value. 
// Parameters:
//  - value : Value to set for the copies property.
func (m *PrinterDocumentConfiguration) SetCopies(value *int32)() {
    m.copies = value
}
// Sets the dpi property value. 
// Parameters:
//  - value : Value to set for the dpi property.
func (m *PrinterDocumentConfiguration) SetDpi(value *int32)() {
    m.dpi = value
}
// Sets the duplexMode property value. 
// Parameters:
//  - value : Value to set for the duplexMode property.
func (m *PrinterDocumentConfiguration) SetDuplexMode(value *PrintDuplexMode)() {
    m.duplexMode = value
}
// Sets the feedDirection property value. 
// Parameters:
//  - value : Value to set for the feedDirection property.
func (m *PrinterDocumentConfiguration) SetFeedDirection(value *PrinterFeedDirection)() {
    m.feedDirection = value
}
// Sets the feedOrientation property value. 
// Parameters:
//  - value : Value to set for the feedOrientation property.
func (m *PrinterDocumentConfiguration) SetFeedOrientation(value *PrinterFeedOrientation)() {
    m.feedOrientation = value
}
// Sets the finishings property value. 
// Parameters:
//  - value : Value to set for the finishings property.
func (m *PrinterDocumentConfiguration) SetFinishings(value []PrintFinishing)() {
    m.finishings = value
}
// Sets the fitPdfToPage property value. 
// Parameters:
//  - value : Value to set for the fitPdfToPage property.
func (m *PrinterDocumentConfiguration) SetFitPdfToPage(value *bool)() {
    m.fitPdfToPage = value
}
// Sets the inputBin property value. 
// Parameters:
//  - value : Value to set for the inputBin property.
func (m *PrinterDocumentConfiguration) SetInputBin(value *string)() {
    m.inputBin = value
}
// Sets the margin property value. 
// Parameters:
//  - value : Value to set for the margin property.
func (m *PrinterDocumentConfiguration) SetMargin(value *PrintMargin)() {
    m.margin = value
}
// Sets the mediaSize property value. 
// Parameters:
//  - value : Value to set for the mediaSize property.
func (m *PrinterDocumentConfiguration) SetMediaSize(value *string)() {
    m.mediaSize = value
}
// Sets the mediaType property value. 
// Parameters:
//  - value : Value to set for the mediaType property.
func (m *PrinterDocumentConfiguration) SetMediaType(value *string)() {
    m.mediaType = value
}
// Sets the multipageLayout property value. 
// Parameters:
//  - value : Value to set for the multipageLayout property.
func (m *PrinterDocumentConfiguration) SetMultipageLayout(value *PrintMultipageLayout)() {
    m.multipageLayout = value
}
// Sets the orientation property value. 
// Parameters:
//  - value : Value to set for the orientation property.
func (m *PrinterDocumentConfiguration) SetOrientation(value *PrintOrientation)() {
    m.orientation = value
}
// Sets the outputBin property value. 
// Parameters:
//  - value : Value to set for the outputBin property.
func (m *PrinterDocumentConfiguration) SetOutputBin(value *string)() {
    m.outputBin = value
}
// Sets the pageRanges property value. 
// Parameters:
//  - value : Value to set for the pageRanges property.
func (m *PrinterDocumentConfiguration) SetPageRanges(value []IntegerRange)() {
    m.pageRanges = value
}
// Sets the pagesPerSheet property value. 
// Parameters:
//  - value : Value to set for the pagesPerSheet property.
func (m *PrinterDocumentConfiguration) SetPagesPerSheet(value *int32)() {
    m.pagesPerSheet = value
}
// Sets the quality property value. 
// Parameters:
//  - value : Value to set for the quality property.
func (m *PrinterDocumentConfiguration) SetQuality(value *PrintQuality)() {
    m.quality = value
}
// Sets the scaling property value. 
// Parameters:
//  - value : Value to set for the scaling property.
func (m *PrinterDocumentConfiguration) SetScaling(value *PrintScaling)() {
    m.scaling = value
}
