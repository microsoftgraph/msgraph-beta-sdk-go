package printer

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia53388d107d18b7b6bb184c54c9838943c88235b68da03e9af1b4e749213daf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/printer/restorefactorydefaults"
    ia9b6d39e72654f16838c410d763d10cd685546e047c1c1c5a67e6795b5e8d139 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/printer/ref"
    ib474284d9436e6f8551d47f035e654aaf3e942509725bfec074dd0021c828bd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/printer/resetdefaults"
    ic395d901ea47de5e1da0e5f094da0bf591159f680b75575381d6eb9712c3ab17 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/printer/getcapabilities"
)

// PrinterRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\printer
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
    m.urlTemplate = "{+baseurl}/print/printerShares/{printerShare_id}/printer{?select,expand}";
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
// GetCapabilities builds and executes requests for operations under \print\printerShares\{printerShare-id}\printer\microsoft.graph.getCapabilities()
func (m *PrinterRequestBuilder) GetCapabilities()(*ic395d901ea47de5e1da0e5f094da0bf591159f680b75575381d6eb9712c3ab17.GetCapabilitiesRequestBuilder) {
    return ic395d901ea47de5e1da0e5f094da0bf591159f680b75575381d6eb9712c3ab17.NewGetCapabilitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) Ref()(*ia9b6d39e72654f16838c410d763d10cd685546e047c1c1c5a67e6795b5e8d139.RefRequestBuilder) {
    return ia9b6d39e72654f16838c410d763d10cd685546e047c1c1c5a67e6795b5e8d139.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) ResetDefaults()(*ib474284d9436e6f8551d47f035e654aaf3e942509725bfec074dd0021c828bd5.ResetDefaultsRequestBuilder) {
    return ib474284d9436e6f8551d47f035e654aaf3e942509725bfec074dd0021c828bd5.NewResetDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) RestoreFactoryDefaults()(*ia53388d107d18b7b6bb184c54c9838943c88235b68da03e9af1b4e749213daf5.RestoreFactoryDefaultsRequestBuilder) {
    return ia53388d107d18b7b6bb184c54c9838943c88235b68da03e9af1b4e749213daf5.NewRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
