package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i34a5e9b6cb103e1786cc9bc0844b427d915cfd9f19ccaf56b5f0f1bab78f824d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/printer"
    i8a20ab890b50d7d004d72ec88971b0682a50e050460da1b743a65b594f56cb56 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/allowedusers"
    id1d036967566825b0af9120492b03a4ec20e8592dbc04e3ef967f4a9ea9fd161 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/allowedgroups"
)

// PrinterShareRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}
type PrinterShareRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrinterShareRequestBuilderDeleteOptions options for Delete
type PrinterShareRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrinterShareRequestBuilderGetOptions options for Get
type PrinterShareRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrinterShareRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrinterShareRequestBuilderGetQueryParameters get printerShares from print
type PrinterShareRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrinterShareRequestBuilderPatchOptions options for Patch
type PrinterShareRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrinterShare;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrinterShareRequestBuilder) AllowedGroups()(*id1d036967566825b0af9120492b03a4ec20e8592dbc04e3ef967f4a9ea9fd161.AllowedGroupsRequestBuilder) {
    return id1d036967566825b0af9120492b03a4ec20e8592dbc04e3ef967f4a9ea9fd161.NewAllowedGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterShareRequestBuilder) AllowedUsers()(*i8a20ab890b50d7d004d72ec88971b0682a50e050460da1b743a65b594f56cb56.AllowedUsersRequestBuilder) {
    return i8a20ab890b50d7d004d72ec88971b0682a50e050460da1b743a65b594f56cb56.NewAllowedUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrinterShareRequestBuilderInternal instantiates a new PrinterShareRequestBuilder and sets the default values.
func NewPrinterShareRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterShareRequestBuilder) {
    m := &PrinterShareRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printerShares/{printerShare_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrinterShareRequestBuilder instantiates a new PrinterShareRequestBuilder and sets the default values.
func NewPrinterShareRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterShareRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterShareRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property printerShares for print
func (m *PrinterShareRequestBuilder) CreateDeleteRequestInformation(options *PrinterShareRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation get printerShares from print
func (m *PrinterShareRequestBuilder) CreateGetRequestInformation(options *PrinterShareRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property printerShares in print
func (m *PrinterShareRequestBuilder) CreatePatchRequestInformation(options *PrinterShareRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Delete delete navigation property printerShares for print
func (m *PrinterShareRequestBuilder) Delete(options *PrinterShareRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get printerShares from print
func (m *PrinterShareRequestBuilder) Get(options *PrinterShareRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrinterShare, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrinterShare() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrinterShare), nil
}
// Patch update the navigation property printerShares in print
func (m *PrinterShareRequestBuilder) Patch(options *PrinterShareRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *PrinterShareRequestBuilder) Printer()(*i34a5e9b6cb103e1786cc9bc0844b427d915cfd9f19ccaf56b5f0f1bab78f824d.PrinterRequestBuilder) {
    return i34a5e9b6cb103e1786cc9bc0844b427d915cfd9f19ccaf56b5f0f1bab78f824d.NewPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
