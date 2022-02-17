package printer

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i74ff12524d3d4fd6b9fbabd25d54107350d1a8b5c2aeef86db6ca303f299e227 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item/printer/getcapabilities"
    i7f0f0a2791d02c8859c1bdb18376c3d1fbc5ddccdeef6c3b387dc86d5c36aa16 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item/printer/ref"
    ie060ba18b8f783e8b3a84095b6bff459022431e97dabe1e07425541e141500bd "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item/printer/restorefactorydefaults"
    iff88b9b2527f0de1824024f9ea3a7407f6568af9e3281a81bf50054c9279f9e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item/printer/resetdefaults"
)

// PrinterRequestBuilder builds and executes requests for operations under \print\shares\{printerShare-id}\printer
type PrinterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrinterRequestBuilderGetOptions options for Get
type PrinterRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrinterRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrinterRequestBuilderGetQueryParameters the printer that this printer share is related to.
type PrinterRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewPrinterRequestBuilderInternal instantiates a new PrinterRequestBuilder and sets the default values.
func NewPrinterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterRequestBuilder) {
    m := &PrinterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/shares/{printerShare_id}/printer{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrinterRequestBuilder instantiates a new PrinterRequestBuilder and sets the default values.
func NewPrinterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the printer that this printer share is related to.
func (m *PrinterRequestBuilder) CreateGetRequestInformation(options *PrinterRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get the printer that this printer share is related to.
func (m *PrinterRequestBuilder) Get(options *PrinterRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printer, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrinter() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printer), nil
}
// GetCapabilities builds and executes requests for operations under \print\shares\{printerShare-id}\printer\microsoft.graph.getCapabilities()
func (m *PrinterRequestBuilder) GetCapabilities()(*i74ff12524d3d4fd6b9fbabd25d54107350d1a8b5c2aeef86db6ca303f299e227.GetCapabilitiesRequestBuilder) {
    return i74ff12524d3d4fd6b9fbabd25d54107350d1a8b5c2aeef86db6ca303f299e227.NewGetCapabilitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) Ref()(*i7f0f0a2791d02c8859c1bdb18376c3d1fbc5ddccdeef6c3b387dc86d5c36aa16.RefRequestBuilder) {
    return i7f0f0a2791d02c8859c1bdb18376c3d1fbc5ddccdeef6c3b387dc86d5c36aa16.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) ResetDefaults()(*iff88b9b2527f0de1824024f9ea3a7407f6568af9e3281a81bf50054c9279f9e9.ResetDefaultsRequestBuilder) {
    return iff88b9b2527f0de1824024f9ea3a7407f6568af9e3281a81bf50054c9279f9e9.NewResetDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) RestoreFactoryDefaults()(*ie060ba18b8f783e8b3a84095b6bff459022431e97dabe1e07425541e141500bd.RestoreFactoryDefaultsRequestBuilder) {
    return ie060ba18b8f783e8b3a84095b6bff459022431e97dabe1e07425541e141500bd.NewRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
