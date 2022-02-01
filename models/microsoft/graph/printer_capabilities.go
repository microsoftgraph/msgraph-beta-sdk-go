package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterCapabilities 
type PrinterCapabilities struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A list of supported bottom margins(in microns) for the printer.
    bottomMargins []int32;
    // True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise.
    collation *bool;
    // The color modes supported by the printer. Valid values are described in the following table.
    colorModes []PrintColorMode;
    // A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types.
    contentTypes []string;
    // The range of copies per job supported by the printer.
    copiesPerJob *IntegerRange;
    // The list of print resolutions in DPI that are supported by the printer.
    dpis []int32;
    // The list of duplex modes that are supported by the printer. Valid values are described in the following table.
    duplexModes []PrintDuplexMode;
    // 
    feedDirections []PrinterFeedDirection;
    // The list of feed orientations that are supported by the printer.
    feedOrientations []PrinterFeedOrientation;
    // Finishing processes the printer supports for a printed document.
    finishings []PrintFinishing;
    // Supported input bins for the printer.
    inputBins []string;
    // True if color printing is supported by the printer; false otherwise. Read-only.
    isColorPrintingSupported *bool;
    // True if the printer supports printing by page ranges; false otherwise.
    isPageRangeSupported *bool;
    // A list of supported left margins(in microns) for the printer.
    leftMargins []int32;
    // The media (i.e., paper) colors supported by the printer.
    mediaColors []string;
    // The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table.
    mediaSizes []string;
    // The media types supported by the printer.
    mediaTypes []string;
    // The presentation directions supported by the printer. Supported values are described in the following table.
    multipageLayouts []PrintMultipageLayout;
    // The print orientations supported by the printer. Valid values are described in the following table.
    orientations []PrintOrientation;
    // The printer's supported output bins (trays).
    outputBins []string;
    // Supported number of Input Pages to impose upon a single Impression.
    pagesPerSheet []int32;
    // The print qualities supported by the printer.
    qualities []PrintQuality;
    // A list of supported right margins(in microns) for the printer.
    rightMargins []int32;
    // Supported print scalings.
    scalings []PrintScaling;
    // 
    supportedColorConfigurations []PrintColorConfiguration;
    // 
    supportedCopiesPerJob *IntegerRange;
    // 
    supportedDocumentMimeTypes []string;
    // 
    supportedDuplexConfigurations []PrintDuplexConfiguration;
    // 
    supportedFinishings []PrintFinishing;
    // 
    supportedMediaColors []string;
    // 
    supportedMediaSizes []string;
    // 
    supportedMediaTypes []PrintMediaType;
    // 
    supportedOrientations []PrintOrientation;
    // 
    supportedOutputBins []string;
    // 
    supportedPagesPerSheet *IntegerRange;
    // 
    supportedPresentationDirections []PrintPresentationDirection;
    // 
    supportedPrintQualities []PrintQuality;
    // True if the printer supports scaling PDF pages to match the print media size; false otherwise.
    supportsFitPdfToPage *bool;
    // A list of supported top margins(in microns) for the printer.
    topMargins []int32;
}
// NewPrinterCapabilities instantiates a new printerCapabilities and sets the default values.
func NewPrinterCapabilities()(*PrinterCapabilities) {
    m := &PrinterCapabilities{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterCapabilities) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBottomMargins gets the bottomMargins property value. A list of supported bottom margins(in microns) for the printer.
func (m *PrinterCapabilities) GetBottomMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.bottomMargins
    }
}
// GetCollation gets the collation property value. True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise.
func (m *PrinterCapabilities) GetCollation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.collation
    }
}
// GetColorModes gets the colorModes property value. The color modes supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) GetColorModes()([]PrintColorMode) {
    if m == nil {
        return nil
    } else {
        return m.colorModes
    }
}
// GetContentTypes gets the contentTypes property value. A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types.
func (m *PrinterCapabilities) GetContentTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contentTypes
    }
}
// GetCopiesPerJob gets the copiesPerJob property value. The range of copies per job supported by the printer.
func (m *PrinterCapabilities) GetCopiesPerJob()(*IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.copiesPerJob
    }
}
// GetDpis gets the dpis property value. The list of print resolutions in DPI that are supported by the printer.
func (m *PrinterCapabilities) GetDpis()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.dpis
    }
}
// GetDuplexModes gets the duplexModes property value. The list of duplex modes that are supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) GetDuplexModes()([]PrintDuplexMode) {
    if m == nil {
        return nil
    } else {
        return m.duplexModes
    }
}
// GetFeedDirections gets the feedDirections property value. 
func (m *PrinterCapabilities) GetFeedDirections()([]PrinterFeedDirection) {
    if m == nil {
        return nil
    } else {
        return m.feedDirections
    }
}
// GetFeedOrientations gets the feedOrientations property value. The list of feed orientations that are supported by the printer.
func (m *PrinterCapabilities) GetFeedOrientations()([]PrinterFeedOrientation) {
    if m == nil {
        return nil
    } else {
        return m.feedOrientations
    }
}
// GetFinishings gets the finishings property value. Finishing processes the printer supports for a printed document.
func (m *PrinterCapabilities) GetFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.finishings
    }
}
// GetInputBins gets the inputBins property value. Supported input bins for the printer.
func (m *PrinterCapabilities) GetInputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.inputBins
    }
}
// GetIsColorPrintingSupported gets the isColorPrintingSupported property value. True if color printing is supported by the printer; false otherwise. Read-only.
func (m *PrinterCapabilities) GetIsColorPrintingSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isColorPrintingSupported
    }
}
// GetIsPageRangeSupported gets the isPageRangeSupported property value. True if the printer supports printing by page ranges; false otherwise.
func (m *PrinterCapabilities) GetIsPageRangeSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPageRangeSupported
    }
}
// GetLeftMargins gets the leftMargins property value. A list of supported left margins(in microns) for the printer.
func (m *PrinterCapabilities) GetLeftMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.leftMargins
    }
}
// GetMediaColors gets the mediaColors property value. The media (i.e., paper) colors supported by the printer.
func (m *PrinterCapabilities) GetMediaColors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaColors
    }
}
// GetMediaSizes gets the mediaSizes property value. The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table.
func (m *PrinterCapabilities) GetMediaSizes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaSizes
    }
}
// GetMediaTypes gets the mediaTypes property value. The media types supported by the printer.
func (m *PrinterCapabilities) GetMediaTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaTypes
    }
}
// GetMultipageLayouts gets the multipageLayouts property value. The presentation directions supported by the printer. Supported values are described in the following table.
func (m *PrinterCapabilities) GetMultipageLayouts()([]PrintMultipageLayout) {
    if m == nil {
        return nil
    } else {
        return m.multipageLayouts
    }
}
// GetOrientations gets the orientations property value. The print orientations supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) GetOrientations()([]PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.orientations
    }
}
// GetOutputBins gets the outputBins property value. The printer's supported output bins (trays).
func (m *PrinterCapabilities) GetOutputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outputBins
    }
}
// GetPagesPerSheet gets the pagesPerSheet property value. Supported number of Input Pages to impose upon a single Impression.
func (m *PrinterCapabilities) GetPagesPerSheet()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.pagesPerSheet
    }
}
// GetQualities gets the qualities property value. The print qualities supported by the printer.
func (m *PrinterCapabilities) GetQualities()([]PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.qualities
    }
}
// GetRightMargins gets the rightMargins property value. A list of supported right margins(in microns) for the printer.
func (m *PrinterCapabilities) GetRightMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.rightMargins
    }
}
// GetScalings gets the scalings property value. Supported print scalings.
func (m *PrinterCapabilities) GetScalings()([]PrintScaling) {
    if m == nil {
        return nil
    } else {
        return m.scalings
    }
}
// GetSupportedColorConfigurations gets the supportedColorConfigurations property value. 
func (m *PrinterCapabilities) GetSupportedColorConfigurations()([]PrintColorConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.supportedColorConfigurations
    }
}
// GetSupportedCopiesPerJob gets the supportedCopiesPerJob property value. 
func (m *PrinterCapabilities) GetSupportedCopiesPerJob()(*IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.supportedCopiesPerJob
    }
}
// GetSupportedDocumentMimeTypes gets the supportedDocumentMimeTypes property value. 
func (m *PrinterCapabilities) GetSupportedDocumentMimeTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedDocumentMimeTypes
    }
}
// GetSupportedDuplexConfigurations gets the supportedDuplexConfigurations property value. 
func (m *PrinterCapabilities) GetSupportedDuplexConfigurations()([]PrintDuplexConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.supportedDuplexConfigurations
    }
}
// GetSupportedFinishings gets the supportedFinishings property value. 
func (m *PrinterCapabilities) GetSupportedFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.supportedFinishings
    }
}
// GetSupportedMediaColors gets the supportedMediaColors property value. 
func (m *PrinterCapabilities) GetSupportedMediaColors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedMediaColors
    }
}
// GetSupportedMediaSizes gets the supportedMediaSizes property value. 
func (m *PrinterCapabilities) GetSupportedMediaSizes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedMediaSizes
    }
}
// GetSupportedMediaTypes gets the supportedMediaTypes property value. 
func (m *PrinterCapabilities) GetSupportedMediaTypes()([]PrintMediaType) {
    if m == nil {
        return nil
    } else {
        return m.supportedMediaTypes
    }
}
// GetSupportedOrientations gets the supportedOrientations property value. 
func (m *PrinterCapabilities) GetSupportedOrientations()([]PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.supportedOrientations
    }
}
// GetSupportedOutputBins gets the supportedOutputBins property value. 
func (m *PrinterCapabilities) GetSupportedOutputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedOutputBins
    }
}
// GetSupportedPagesPerSheet gets the supportedPagesPerSheet property value. 
func (m *PrinterCapabilities) GetSupportedPagesPerSheet()(*IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.supportedPagesPerSheet
    }
}
// GetSupportedPresentationDirections gets the supportedPresentationDirections property value. 
func (m *PrinterCapabilities) GetSupportedPresentationDirections()([]PrintPresentationDirection) {
    if m == nil {
        return nil
    } else {
        return m.supportedPresentationDirections
    }
}
// GetSupportedPrintQualities gets the supportedPrintQualities property value. 
func (m *PrinterCapabilities) GetSupportedPrintQualities()([]PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.supportedPrintQualities
    }
}
// GetSupportsFitPdfToPage gets the supportsFitPdfToPage property value. True if the printer supports scaling PDF pages to match the print media size; false otherwise.
func (m *PrinterCapabilities) GetSupportsFitPdfToPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supportsFitPdfToPage
    }
}
// GetTopMargins gets the topMargins property value. A list of supported top margins(in microns) for the printer.
func (m *PrinterCapabilities) GetTopMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.topMargins
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterCapabilities) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bottomMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetBottomMargins(res)
        }
        return nil
    }
    res["collation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollation(val)
        }
        return nil
    }
    res["colorModes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintColorMode)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintColorMode, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintColorMode))
            }
            m.SetColorModes(res)
        }
        return nil
    }
    res["contentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetContentTypes(res)
        }
        return nil
    }
    res["copiesPerJob"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopiesPerJob(val.(*IntegerRange))
        }
        return nil
    }
    res["dpis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetDpis(res)
        }
        return nil
    }
    res["duplexModes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintDuplexMode)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintDuplexMode, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintDuplexMode))
            }
            m.SetDuplexModes(res)
        }
        return nil
    }
    res["feedDirections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrinterFeedDirection)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterFeedDirection, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrinterFeedDirection))
            }
            m.SetFeedDirections(res)
        }
        return nil
    }
    res["feedOrientations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrinterFeedOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterFeedOrientation, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrinterFeedOrientation))
            }
            m.SetFeedOrientations(res)
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
    res["inputBins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInputBins(res)
        }
        return nil
    }
    res["isColorPrintingSupported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsColorPrintingSupported(val)
        }
        return nil
    }
    res["isPageRangeSupported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPageRangeSupported(val)
        }
        return nil
    }
    res["leftMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetLeftMargins(res)
        }
        return nil
    }
    res["mediaColors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMediaColors(res)
        }
        return nil
    }
    res["mediaSizes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMediaSizes(res)
        }
        return nil
    }
    res["mediaTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMediaTypes(res)
        }
        return nil
    }
    res["multipageLayouts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintMultipageLayout)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintMultipageLayout, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintMultipageLayout))
            }
            m.SetMultipageLayouts(res)
        }
        return nil
    }
    res["orientations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintOrientation, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintOrientation))
            }
            m.SetOrientations(res)
        }
        return nil
    }
    res["outputBins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOutputBins(res)
        }
        return nil
    }
    res["pagesPerSheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetPagesPerSheet(res)
        }
        return nil
    }
    res["qualities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintQuality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintQuality, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintQuality))
            }
            m.SetQualities(res)
        }
        return nil
    }
    res["rightMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetRightMargins(res)
        }
        return nil
    }
    res["scalings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintScaling)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintScaling, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintScaling))
            }
            m.SetScalings(res)
        }
        return nil
    }
    res["supportedColorConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintColorConfiguration)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintColorConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintColorConfiguration))
            }
            m.SetSupportedColorConfigurations(res)
        }
        return nil
    }
    res["supportedCopiesPerJob"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedCopiesPerJob(val.(*IntegerRange))
        }
        return nil
    }
    res["supportedDocumentMimeTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedDocumentMimeTypes(res)
        }
        return nil
    }
    res["supportedDuplexConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintDuplexConfiguration)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintDuplexConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintDuplexConfiguration))
            }
            m.SetSupportedDuplexConfigurations(res)
        }
        return nil
    }
    res["supportedFinishings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintFinishing)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintFinishing, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintFinishing))
            }
            m.SetSupportedFinishings(res)
        }
        return nil
    }
    res["supportedMediaColors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedMediaColors(res)
        }
        return nil
    }
    res["supportedMediaSizes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedMediaSizes(res)
        }
        return nil
    }
    res["supportedMediaTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintMediaType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintMediaType, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintMediaType))
            }
            m.SetSupportedMediaTypes(res)
        }
        return nil
    }
    res["supportedOrientations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintOrientation, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintOrientation))
            }
            m.SetSupportedOrientations(res)
        }
        return nil
    }
    res["supportedOutputBins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedOutputBins(res)
        }
        return nil
    }
    res["supportedPagesPerSheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedPagesPerSheet(val.(*IntegerRange))
        }
        return nil
    }
    res["supportedPresentationDirections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintPresentationDirection)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintPresentationDirection, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintPresentationDirection))
            }
            m.SetSupportedPresentationDirections(res)
        }
        return nil
    }
    res["supportedPrintQualities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintQuality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintQuality, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintQuality))
            }
            m.SetSupportedPrintQualities(res)
        }
        return nil
    }
    res["supportsFitPdfToPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsFitPdfToPage(val)
        }
        return nil
    }
    res["topMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetTopMargins(res)
        }
        return nil
    }
    return res
}
func (m *PrinterCapabilities) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrinterCapabilities) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetBottomMargins() != nil {
        err := writer.WriteCollectionOfInt32Values("bottomMargins", m.GetBottomMargins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("collation", m.GetCollation())
        if err != nil {
            return err
        }
    }
    if m.GetColorModes() != nil {
        err := writer.WriteCollectionOfStringValues("colorModes", SerializePrintColorMode(m.GetColorModes()))
        if err != nil {
            return err
        }
    }
    if m.GetContentTypes() != nil {
        err := writer.WriteCollectionOfStringValues("contentTypes", m.GetContentTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("copiesPerJob", m.GetCopiesPerJob())
        if err != nil {
            return err
        }
    }
    if m.GetDpis() != nil {
        err := writer.WriteCollectionOfInt32Values("dpis", m.GetDpis())
        if err != nil {
            return err
        }
    }
    if m.GetDuplexModes() != nil {
        err := writer.WriteCollectionOfStringValues("duplexModes", SerializePrintDuplexMode(m.GetDuplexModes()))
        if err != nil {
            return err
        }
    }
    if m.GetFeedDirections() != nil {
        err := writer.WriteCollectionOfStringValues("feedDirections", SerializePrinterFeedDirection(m.GetFeedDirections()))
        if err != nil {
            return err
        }
    }
    if m.GetFeedOrientations() != nil {
        err := writer.WriteCollectionOfStringValues("feedOrientations", SerializePrinterFeedOrientation(m.GetFeedOrientations()))
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
    if m.GetInputBins() != nil {
        err := writer.WriteCollectionOfStringValues("inputBins", m.GetInputBins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isColorPrintingSupported", m.GetIsColorPrintingSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPageRangeSupported", m.GetIsPageRangeSupported())
        if err != nil {
            return err
        }
    }
    if m.GetLeftMargins() != nil {
        err := writer.WriteCollectionOfInt32Values("leftMargins", m.GetLeftMargins())
        if err != nil {
            return err
        }
    }
    if m.GetMediaColors() != nil {
        err := writer.WriteCollectionOfStringValues("mediaColors", m.GetMediaColors())
        if err != nil {
            return err
        }
    }
    if m.GetMediaSizes() != nil {
        err := writer.WriteCollectionOfStringValues("mediaSizes", m.GetMediaSizes())
        if err != nil {
            return err
        }
    }
    if m.GetMediaTypes() != nil {
        err := writer.WriteCollectionOfStringValues("mediaTypes", m.GetMediaTypes())
        if err != nil {
            return err
        }
    }
    if m.GetMultipageLayouts() != nil {
        err := writer.WriteCollectionOfStringValues("multipageLayouts", SerializePrintMultipageLayout(m.GetMultipageLayouts()))
        if err != nil {
            return err
        }
    }
    if m.GetOrientations() != nil {
        err := writer.WriteCollectionOfStringValues("orientations", SerializePrintOrientation(m.GetOrientations()))
        if err != nil {
            return err
        }
    }
    if m.GetOutputBins() != nil {
        err := writer.WriteCollectionOfStringValues("outputBins", m.GetOutputBins())
        if err != nil {
            return err
        }
    }
    if m.GetPagesPerSheet() != nil {
        err := writer.WriteCollectionOfInt32Values("pagesPerSheet", m.GetPagesPerSheet())
        if err != nil {
            return err
        }
    }
    if m.GetQualities() != nil {
        err := writer.WriteCollectionOfStringValues("qualities", SerializePrintQuality(m.GetQualities()))
        if err != nil {
            return err
        }
    }
    if m.GetRightMargins() != nil {
        err := writer.WriteCollectionOfInt32Values("rightMargins", m.GetRightMargins())
        if err != nil {
            return err
        }
    }
    if m.GetScalings() != nil {
        err := writer.WriteCollectionOfStringValues("scalings", SerializePrintScaling(m.GetScalings()))
        if err != nil {
            return err
        }
    }
    if m.GetSupportedColorConfigurations() != nil {
        err := writer.WriteCollectionOfStringValues("supportedColorConfigurations", SerializePrintColorConfiguration(m.GetSupportedColorConfigurations()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("supportedCopiesPerJob", m.GetSupportedCopiesPerJob())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedDocumentMimeTypes() != nil {
        err := writer.WriteCollectionOfStringValues("supportedDocumentMimeTypes", m.GetSupportedDocumentMimeTypes())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedDuplexConfigurations() != nil {
        err := writer.WriteCollectionOfStringValues("supportedDuplexConfigurations", SerializePrintDuplexConfiguration(m.GetSupportedDuplexConfigurations()))
        if err != nil {
            return err
        }
    }
    if m.GetSupportedFinishings() != nil {
        err := writer.WriteCollectionOfStringValues("supportedFinishings", SerializePrintFinishing(m.GetSupportedFinishings()))
        if err != nil {
            return err
        }
    }
    if m.GetSupportedMediaColors() != nil {
        err := writer.WriteCollectionOfStringValues("supportedMediaColors", m.GetSupportedMediaColors())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedMediaSizes() != nil {
        err := writer.WriteCollectionOfStringValues("supportedMediaSizes", m.GetSupportedMediaSizes())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedMediaTypes() != nil {
        err := writer.WriteCollectionOfStringValues("supportedMediaTypes", SerializePrintMediaType(m.GetSupportedMediaTypes()))
        if err != nil {
            return err
        }
    }
    if m.GetSupportedOrientations() != nil {
        err := writer.WriteCollectionOfStringValues("supportedOrientations", SerializePrintOrientation(m.GetSupportedOrientations()))
        if err != nil {
            return err
        }
    }
    if m.GetSupportedOutputBins() != nil {
        err := writer.WriteCollectionOfStringValues("supportedOutputBins", m.GetSupportedOutputBins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("supportedPagesPerSheet", m.GetSupportedPagesPerSheet())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedPresentationDirections() != nil {
        err := writer.WriteCollectionOfStringValues("supportedPresentationDirections", SerializePrintPresentationDirection(m.GetSupportedPresentationDirections()))
        if err != nil {
            return err
        }
    }
    if m.GetSupportedPrintQualities() != nil {
        err := writer.WriteCollectionOfStringValues("supportedPrintQualities", SerializePrintQuality(m.GetSupportedPrintQualities()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("supportsFitPdfToPage", m.GetSupportsFitPdfToPage())
        if err != nil {
            return err
        }
    }
    if m.GetTopMargins() != nil {
        err := writer.WriteCollectionOfInt32Values("topMargins", m.GetTopMargins())
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
func (m *PrinterCapabilities) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBottomMargins sets the bottomMargins property value. A list of supported bottom margins(in microns) for the printer.
func (m *PrinterCapabilities) SetBottomMargins(value []int32)() {
    if m != nil {
        m.bottomMargins = value
    }
}
// SetCollation sets the collation property value. True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise.
func (m *PrinterCapabilities) SetCollation(value *bool)() {
    if m != nil {
        m.collation = value
    }
}
// SetColorModes sets the colorModes property value. The color modes supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) SetColorModes(value []PrintColorMode)() {
    if m != nil {
        m.colorModes = value
    }
}
// SetContentTypes sets the contentTypes property value. A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types.
func (m *PrinterCapabilities) SetContentTypes(value []string)() {
    if m != nil {
        m.contentTypes = value
    }
}
// SetCopiesPerJob sets the copiesPerJob property value. The range of copies per job supported by the printer.
func (m *PrinterCapabilities) SetCopiesPerJob(value *IntegerRange)() {
    if m != nil {
        m.copiesPerJob = value
    }
}
// SetDpis sets the dpis property value. The list of print resolutions in DPI that are supported by the printer.
func (m *PrinterCapabilities) SetDpis(value []int32)() {
    if m != nil {
        m.dpis = value
    }
}
// SetDuplexModes sets the duplexModes property value. The list of duplex modes that are supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) SetDuplexModes(value []PrintDuplexMode)() {
    if m != nil {
        m.duplexModes = value
    }
}
// SetFeedDirections sets the feedDirections property value. 
func (m *PrinterCapabilities) SetFeedDirections(value []PrinterFeedDirection)() {
    if m != nil {
        m.feedDirections = value
    }
}
// SetFeedOrientations sets the feedOrientations property value. The list of feed orientations that are supported by the printer.
func (m *PrinterCapabilities) SetFeedOrientations(value []PrinterFeedOrientation)() {
    if m != nil {
        m.feedOrientations = value
    }
}
// SetFinishings sets the finishings property value. Finishing processes the printer supports for a printed document.
func (m *PrinterCapabilities) SetFinishings(value []PrintFinishing)() {
    if m != nil {
        m.finishings = value
    }
}
// SetInputBins sets the inputBins property value. Supported input bins for the printer.
func (m *PrinterCapabilities) SetInputBins(value []string)() {
    if m != nil {
        m.inputBins = value
    }
}
// SetIsColorPrintingSupported sets the isColorPrintingSupported property value. True if color printing is supported by the printer; false otherwise. Read-only.
func (m *PrinterCapabilities) SetIsColorPrintingSupported(value *bool)() {
    if m != nil {
        m.isColorPrintingSupported = value
    }
}
// SetIsPageRangeSupported sets the isPageRangeSupported property value. True if the printer supports printing by page ranges; false otherwise.
func (m *PrinterCapabilities) SetIsPageRangeSupported(value *bool)() {
    if m != nil {
        m.isPageRangeSupported = value
    }
}
// SetLeftMargins sets the leftMargins property value. A list of supported left margins(in microns) for the printer.
func (m *PrinterCapabilities) SetLeftMargins(value []int32)() {
    if m != nil {
        m.leftMargins = value
    }
}
// SetMediaColors sets the mediaColors property value. The media (i.e., paper) colors supported by the printer.
func (m *PrinterCapabilities) SetMediaColors(value []string)() {
    if m != nil {
        m.mediaColors = value
    }
}
// SetMediaSizes sets the mediaSizes property value. The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table.
func (m *PrinterCapabilities) SetMediaSizes(value []string)() {
    if m != nil {
        m.mediaSizes = value
    }
}
// SetMediaTypes sets the mediaTypes property value. The media types supported by the printer.
func (m *PrinterCapabilities) SetMediaTypes(value []string)() {
    if m != nil {
        m.mediaTypes = value
    }
}
// SetMultipageLayouts sets the multipageLayouts property value. The presentation directions supported by the printer. Supported values are described in the following table.
func (m *PrinterCapabilities) SetMultipageLayouts(value []PrintMultipageLayout)() {
    if m != nil {
        m.multipageLayouts = value
    }
}
// SetOrientations sets the orientations property value. The print orientations supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) SetOrientations(value []PrintOrientation)() {
    if m != nil {
        m.orientations = value
    }
}
// SetOutputBins sets the outputBins property value. The printer's supported output bins (trays).
func (m *PrinterCapabilities) SetOutputBins(value []string)() {
    if m != nil {
        m.outputBins = value
    }
}
// SetPagesPerSheet sets the pagesPerSheet property value. Supported number of Input Pages to impose upon a single Impression.
func (m *PrinterCapabilities) SetPagesPerSheet(value []int32)() {
    if m != nil {
        m.pagesPerSheet = value
    }
}
// SetQualities sets the qualities property value. The print qualities supported by the printer.
func (m *PrinterCapabilities) SetQualities(value []PrintQuality)() {
    if m != nil {
        m.qualities = value
    }
}
// SetRightMargins sets the rightMargins property value. A list of supported right margins(in microns) for the printer.
func (m *PrinterCapabilities) SetRightMargins(value []int32)() {
    if m != nil {
        m.rightMargins = value
    }
}
// SetScalings sets the scalings property value. Supported print scalings.
func (m *PrinterCapabilities) SetScalings(value []PrintScaling)() {
    if m != nil {
        m.scalings = value
    }
}
// SetSupportedColorConfigurations sets the supportedColorConfigurations property value. 
func (m *PrinterCapabilities) SetSupportedColorConfigurations(value []PrintColorConfiguration)() {
    if m != nil {
        m.supportedColorConfigurations = value
    }
}
// SetSupportedCopiesPerJob sets the supportedCopiesPerJob property value. 
func (m *PrinterCapabilities) SetSupportedCopiesPerJob(value *IntegerRange)() {
    if m != nil {
        m.supportedCopiesPerJob = value
    }
}
// SetSupportedDocumentMimeTypes sets the supportedDocumentMimeTypes property value. 
func (m *PrinterCapabilities) SetSupportedDocumentMimeTypes(value []string)() {
    if m != nil {
        m.supportedDocumentMimeTypes = value
    }
}
// SetSupportedDuplexConfigurations sets the supportedDuplexConfigurations property value. 
func (m *PrinterCapabilities) SetSupportedDuplexConfigurations(value []PrintDuplexConfiguration)() {
    if m != nil {
        m.supportedDuplexConfigurations = value
    }
}
// SetSupportedFinishings sets the supportedFinishings property value. 
func (m *PrinterCapabilities) SetSupportedFinishings(value []PrintFinishing)() {
    if m != nil {
        m.supportedFinishings = value
    }
}
// SetSupportedMediaColors sets the supportedMediaColors property value. 
func (m *PrinterCapabilities) SetSupportedMediaColors(value []string)() {
    if m != nil {
        m.supportedMediaColors = value
    }
}
// SetSupportedMediaSizes sets the supportedMediaSizes property value. 
func (m *PrinterCapabilities) SetSupportedMediaSizes(value []string)() {
    if m != nil {
        m.supportedMediaSizes = value
    }
}
// SetSupportedMediaTypes sets the supportedMediaTypes property value. 
func (m *PrinterCapabilities) SetSupportedMediaTypes(value []PrintMediaType)() {
    if m != nil {
        m.supportedMediaTypes = value
    }
}
// SetSupportedOrientations sets the supportedOrientations property value. 
func (m *PrinterCapabilities) SetSupportedOrientations(value []PrintOrientation)() {
    if m != nil {
        m.supportedOrientations = value
    }
}
// SetSupportedOutputBins sets the supportedOutputBins property value. 
func (m *PrinterCapabilities) SetSupportedOutputBins(value []string)() {
    if m != nil {
        m.supportedOutputBins = value
    }
}
// SetSupportedPagesPerSheet sets the supportedPagesPerSheet property value. 
func (m *PrinterCapabilities) SetSupportedPagesPerSheet(value *IntegerRange)() {
    if m != nil {
        m.supportedPagesPerSheet = value
    }
}
// SetSupportedPresentationDirections sets the supportedPresentationDirections property value. 
func (m *PrinterCapabilities) SetSupportedPresentationDirections(value []PrintPresentationDirection)() {
    if m != nil {
        m.supportedPresentationDirections = value
    }
}
// SetSupportedPrintQualities sets the supportedPrintQualities property value. 
func (m *PrinterCapabilities) SetSupportedPrintQualities(value []PrintQuality)() {
    if m != nil {
        m.supportedPrintQualities = value
    }
}
// SetSupportsFitPdfToPage sets the supportsFitPdfToPage property value. True if the printer supports scaling PDF pages to match the print media size; false otherwise.
func (m *PrinterCapabilities) SetSupportsFitPdfToPage(value *bool)() {
    if m != nil {
        m.supportsFitPdfToPage = value
    }
}
// SetTopMargins sets the topMargins property value. A list of supported top margins(in microns) for the printer.
func (m *PrinterCapabilities) SetTopMargins(value []int32)() {
    if m != nil {
        m.topMargins = value
    }
}
