package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterDocumentConfiguration 
type PrinterDocumentConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The collate property
    collate *bool
    // The colorMode property
    colorMode *PrintColorMode
    // The copies property
    copies *int32
    // The dpi property
    dpi *int32
    // The duplexMode property
    duplexMode *PrintDuplexMode
    // The feedDirection property
    feedDirection *PrinterFeedDirection
    // The feedOrientation property
    feedOrientation *PrinterFeedOrientation
    // The finishings property
    finishings []PrintFinishing
    // The fitPdfToPage property
    fitPdfToPage *bool
    // The inputBin property
    inputBin *string
    // The margin property
    margin PrintMarginable
    // The mediaSize property
    mediaSize *string
    // The mediaType property
    mediaType *string
    // The multipageLayout property
    multipageLayout *PrintMultipageLayout
    // The OdataType property
    odataType *string
    // The orientation property
    orientation *PrintOrientation
    // The outputBin property
    outputBin *string
    // The pageRanges property
    pageRanges []IntegerRangeable
    // The pagesPerSheet property
    pagesPerSheet *int32
    // The quality property
    quality *PrintQuality
    // The scaling property
    scaling *PrintScaling
}
// NewPrinterDocumentConfiguration instantiates a new printerDocumentConfiguration and sets the default values.
func NewPrinterDocumentConfiguration()(*PrinterDocumentConfiguration) {
    m := &PrinterDocumentConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.printerDocumentConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePrinterDocumentConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterDocumentConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterDocumentConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterDocumentConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCollate gets the collate property value. The collate property
func (m *PrinterDocumentConfiguration) GetCollate()(*bool) {
    return m.collate
}
// GetColorMode gets the colorMode property value. The colorMode property
func (m *PrinterDocumentConfiguration) GetColorMode()(*PrintColorMode) {
    return m.colorMode
}
// GetCopies gets the copies property value. The copies property
func (m *PrinterDocumentConfiguration) GetCopies()(*int32) {
    return m.copies
}
// GetDpi gets the dpi property value. The dpi property
func (m *PrinterDocumentConfiguration) GetDpi()(*int32) {
    return m.dpi
}
// GetDuplexMode gets the duplexMode property value. The duplexMode property
func (m *PrinterDocumentConfiguration) GetDuplexMode()(*PrintDuplexMode) {
    return m.duplexMode
}
// GetFeedDirection gets the feedDirection property value. The feedDirection property
func (m *PrinterDocumentConfiguration) GetFeedDirection()(*PrinterFeedDirection) {
    return m.feedDirection
}
// GetFeedOrientation gets the feedOrientation property value. The feedOrientation property
func (m *PrinterDocumentConfiguration) GetFeedOrientation()(*PrinterFeedOrientation) {
    return m.feedOrientation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterDocumentConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["collate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCollate)
    res["colorMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePrintColorMode , m.SetColorMode)
    res["copies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCopies)
    res["dpi"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDpi)
    res["duplexMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePrintDuplexMode , m.SetDuplexMode)
    res["feedDirection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePrinterFeedDirection , m.SetFeedDirection)
    res["feedOrientation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePrinterFeedOrientation , m.SetFeedOrientation)
    res["finishings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParsePrintFinishing , m.SetFinishings)
    res["fitPdfToPage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFitPdfToPage)
    res["inputBin"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInputBin)
    res["margin"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePrintMarginFromDiscriminatorValue , m.SetMargin)
    res["mediaSize"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMediaSize)
    res["mediaType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMediaType)
    res["multipageLayout"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePrintMultipageLayout , m.SetMultipageLayout)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["orientation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePrintOrientation , m.SetOrientation)
    res["outputBin"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOutputBin)
    res["pageRanges"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIntegerRangeFromDiscriminatorValue , m.SetPageRanges)
    res["pagesPerSheet"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPagesPerSheet)
    res["quality"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePrintQuality , m.SetQuality)
    res["scaling"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePrintScaling , m.SetScaling)
    return res
}
// GetFinishings gets the finishings property value. The finishings property
func (m *PrinterDocumentConfiguration) GetFinishings()([]PrintFinishing) {
    return m.finishings
}
// GetFitPdfToPage gets the fitPdfToPage property value. The fitPdfToPage property
func (m *PrinterDocumentConfiguration) GetFitPdfToPage()(*bool) {
    return m.fitPdfToPage
}
// GetInputBin gets the inputBin property value. The inputBin property
func (m *PrinterDocumentConfiguration) GetInputBin()(*string) {
    return m.inputBin
}
// GetMargin gets the margin property value. The margin property
func (m *PrinterDocumentConfiguration) GetMargin()(PrintMarginable) {
    return m.margin
}
// GetMediaSize gets the mediaSize property value. The mediaSize property
func (m *PrinterDocumentConfiguration) GetMediaSize()(*string) {
    return m.mediaSize
}
// GetMediaType gets the mediaType property value. The mediaType property
func (m *PrinterDocumentConfiguration) GetMediaType()(*string) {
    return m.mediaType
}
// GetMultipageLayout gets the multipageLayout property value. The multipageLayout property
func (m *PrinterDocumentConfiguration) GetMultipageLayout()(*PrintMultipageLayout) {
    return m.multipageLayout
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrinterDocumentConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetOrientation gets the orientation property value. The orientation property
func (m *PrinterDocumentConfiguration) GetOrientation()(*PrintOrientation) {
    return m.orientation
}
// GetOutputBin gets the outputBin property value. The outputBin property
func (m *PrinterDocumentConfiguration) GetOutputBin()(*string) {
    return m.outputBin
}
// GetPageRanges gets the pageRanges property value. The pageRanges property
func (m *PrinterDocumentConfiguration) GetPageRanges()([]IntegerRangeable) {
    return m.pageRanges
}
// GetPagesPerSheet gets the pagesPerSheet property value. The pagesPerSheet property
func (m *PrinterDocumentConfiguration) GetPagesPerSheet()(*int32) {
    return m.pagesPerSheet
}
// GetQuality gets the quality property value. The quality property
func (m *PrinterDocumentConfiguration) GetQuality()(*PrintQuality) {
    return m.quality
}
// GetScaling gets the scaling property value. The scaling property
func (m *PrinterDocumentConfiguration) GetScaling()(*PrintScaling) {
    return m.scaling
}
// Serialize serializes information the current object
func (m *PrinterDocumentConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPageRanges())
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
    m.additionalData = value
}
// SetCollate sets the collate property value. The collate property
func (m *PrinterDocumentConfiguration) SetCollate(value *bool)() {
    m.collate = value
}
// SetColorMode sets the colorMode property value. The colorMode property
func (m *PrinterDocumentConfiguration) SetColorMode(value *PrintColorMode)() {
    m.colorMode = value
}
// SetCopies sets the copies property value. The copies property
func (m *PrinterDocumentConfiguration) SetCopies(value *int32)() {
    m.copies = value
}
// SetDpi sets the dpi property value. The dpi property
func (m *PrinterDocumentConfiguration) SetDpi(value *int32)() {
    m.dpi = value
}
// SetDuplexMode sets the duplexMode property value. The duplexMode property
func (m *PrinterDocumentConfiguration) SetDuplexMode(value *PrintDuplexMode)() {
    m.duplexMode = value
}
// SetFeedDirection sets the feedDirection property value. The feedDirection property
func (m *PrinterDocumentConfiguration) SetFeedDirection(value *PrinterFeedDirection)() {
    m.feedDirection = value
}
// SetFeedOrientation sets the feedOrientation property value. The feedOrientation property
func (m *PrinterDocumentConfiguration) SetFeedOrientation(value *PrinterFeedOrientation)() {
    m.feedOrientation = value
}
// SetFinishings sets the finishings property value. The finishings property
func (m *PrinterDocumentConfiguration) SetFinishings(value []PrintFinishing)() {
    m.finishings = value
}
// SetFitPdfToPage sets the fitPdfToPage property value. The fitPdfToPage property
func (m *PrinterDocumentConfiguration) SetFitPdfToPage(value *bool)() {
    m.fitPdfToPage = value
}
// SetInputBin sets the inputBin property value. The inputBin property
func (m *PrinterDocumentConfiguration) SetInputBin(value *string)() {
    m.inputBin = value
}
// SetMargin sets the margin property value. The margin property
func (m *PrinterDocumentConfiguration) SetMargin(value PrintMarginable)() {
    m.margin = value
}
// SetMediaSize sets the mediaSize property value. The mediaSize property
func (m *PrinterDocumentConfiguration) SetMediaSize(value *string)() {
    m.mediaSize = value
}
// SetMediaType sets the mediaType property value. The mediaType property
func (m *PrinterDocumentConfiguration) SetMediaType(value *string)() {
    m.mediaType = value
}
// SetMultipageLayout sets the multipageLayout property value. The multipageLayout property
func (m *PrinterDocumentConfiguration) SetMultipageLayout(value *PrintMultipageLayout)() {
    m.multipageLayout = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrinterDocumentConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrientation sets the orientation property value. The orientation property
func (m *PrinterDocumentConfiguration) SetOrientation(value *PrintOrientation)() {
    m.orientation = value
}
// SetOutputBin sets the outputBin property value. The outputBin property
func (m *PrinterDocumentConfiguration) SetOutputBin(value *string)() {
    m.outputBin = value
}
// SetPageRanges sets the pageRanges property value. The pageRanges property
func (m *PrinterDocumentConfiguration) SetPageRanges(value []IntegerRangeable)() {
    m.pageRanges = value
}
// SetPagesPerSheet sets the pagesPerSheet property value. The pagesPerSheet property
func (m *PrinterDocumentConfiguration) SetPagesPerSheet(value *int32)() {
    m.pagesPerSheet = value
}
// SetQuality sets the quality property value. The quality property
func (m *PrinterDocumentConfiguration) SetQuality(value *PrintQuality)() {
    m.quality = value
}
// SetScaling sets the scaling property value. The scaling property
func (m *PrinterDocumentConfiguration) SetScaling(value *PrintScaling)() {
    m.scaling = value
}
