package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i599f577b7cc2b4c7ec4cac4c88c7a0ac038b8dd35e1dd9006d8902495994d7a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item/printer"
    i83d84a5570e033ee2154c9058ba94b73e02d2477c15f7f62129d78cf94bf25e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item/allowedusers"
    ia38aafee6e99eed8fcd1a421defe593236a7b1a6baa17cc46600d14ef087a65c "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item/allowedgroups"
)

// PrinterShareItemRequestBuilder builds and executes requests for operations under \print\shares\{printerShare-id}
type PrinterShareItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrinterShareItemRequestBuilderDeleteOptions options for Delete
type PrinterShareItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrinterShareItemRequestBuilderGetOptions options for Get
type PrinterShareItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrinterShareItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrinterShareItemRequestBuilderGetQueryParameters the list of printer shares registered in the tenant.
type PrinterShareItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrinterShareItemRequestBuilderPatchOptions options for Patch
type PrinterShareItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrinterShare;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrinterShareItemRequestBuilder) AllowedGroups()(*ia38aafee6e99eed8fcd1a421defe593236a7b1a6baa17cc46600d14ef087a65c.AllowedGroupsRequestBuilder) {
    return ia38aafee6e99eed8fcd1a421defe593236a7b1a6baa17cc46600d14ef087a65c.NewAllowedGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterShareItemRequestBuilder) AllowedUsers()(*i83d84a5570e033ee2154c9058ba94b73e02d2477c15f7f62129d78cf94bf25e1.AllowedUsersRequestBuilder) {
    return i83d84a5570e033ee2154c9058ba94b73e02d2477c15f7f62129d78cf94bf25e1.NewAllowedUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrinterShareItemRequestBuilderInternal instantiates a new PrinterShareItemRequestBuilder and sets the default values.
func NewPrinterShareItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterShareItemRequestBuilder) {
    m := &PrinterShareItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/shares/{printerShare_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrinterShareItemRequestBuilder instantiates a new PrinterShareItemRequestBuilder and sets the default values.
func NewPrinterShareItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterShareItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterShareItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of printer shares registered in the tenant.
func (m *PrinterShareItemRequestBuilder) CreateDeleteRequestInformation(options *PrinterShareItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of printer shares registered in the tenant.
func (m *PrinterShareItemRequestBuilder) CreateGetRequestInformation(options *PrinterShareItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of printer shares registered in the tenant.
func (m *PrinterShareItemRequestBuilder) CreatePatchRequestInformation(options *PrinterShareItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of printer shares registered in the tenant.
func (m *PrinterShareItemRequestBuilder) Delete(options *PrinterShareItemRequestBuilderDeleteOptions)(error) {
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
// Get the list of printer shares registered in the tenant.
func (m *PrinterShareItemRequestBuilder) Get(options *PrinterShareItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrinterShare, error) {
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
// Patch the list of printer shares registered in the tenant.
func (m *PrinterShareItemRequestBuilder) Patch(options *PrinterShareItemRequestBuilderPatchOptions)(error) {
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
func (m *PrinterShareItemRequestBuilder) Printer()(*i599f577b7cc2b4c7ec4cac4c88c7a0ac038b8dd35e1dd9006d8902495994d7a6.PrinterRequestBuilder) {
    return i599f577b7cc2b4c7ec4cac4c88c7a0ac038b8dd35e1dd9006d8902495994d7a6.NewPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
