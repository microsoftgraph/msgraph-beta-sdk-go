package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i67355dbc94c107ed697494362f24074fe629475baf93f5f5554922e54a4817d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/macossoftwareupdateaccountsummaries/item/categorysummaries"
    i407f40cbdcd6e0a8c0a6957a52e770bb996c3efb0c3933a6259c21531a23bf20 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/macossoftwareupdateaccountsummaries/item/categorysummaries/item"
)

// MacOSSoftwareUpdateAccountSummaryItemRequestBuilder provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
type MacOSSoftwareUpdateAccountSummaryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteOptions options for Delete
type MacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetOptions options for Get
type MacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters the MacOS software update account summaries for this account.
type MacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchOptions options for Patch
type MacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateAccountSummaryable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CategorySummaries()(*i67355dbc94c107ed697494362f24074fe629475baf93f5f5554922e54a4817d7.CategorySummariesRequestBuilder) {
    return i67355dbc94c107ed697494362f24074fe629475baf93f5f5554922e54a4817d7.NewCategorySummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategorySummariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.macOSSoftwareUpdateAccountSummaries.item.categorySummaries.item collection
func (m *MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CategorySummariesById(id string)(*i407f40cbdcd6e0a8c0a6957a52e770bb996c3efb0c3933a6259c21531a23bf20.MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["macOSSoftwareUpdateCategorySummary_id"] = id
    }
    return i407f40cbdcd6e0a8c0a6957a52e770bb996c3efb0c3933a6259c21531a23bf20.NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal instantiates a new MacOSSoftwareUpdateAccountSummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    m := &MacOSSoftwareUpdateAccountSummaryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMacOSSoftwareUpdateAccountSummaryItemRequestBuilder instantiates a new MacOSSoftwareUpdateAccountSummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateAccountSummaryItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property macOSSoftwareUpdateAccountSummaries for deviceManagement
func (m *MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CreateDeleteRequestInformation(options *MacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the MacOS software update account summaries for this account.
func (m *MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CreateGetRequestInformation(options *MacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property macOSSoftwareUpdateAccountSummaries in deviceManagement
func (m *MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CreatePatchRequestInformation(options *MacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property macOSSoftwareUpdateAccountSummaries for deviceManagement
func (m *MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Delete(options *MacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the MacOS software update account summaries for this account.
func (m *MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Get(options *MacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateAccountSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateAccountSummaryable), nil
}
// Patch update the navigation property macOSSoftwareUpdateAccountSummaries in deviceManagement
func (m *MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Patch(options *MacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
