package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia704cb6be080f1ecb092ccb2ba1c744783183d7b569f3801033f58df43a201db "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/macossoftwareupdateaccountsummaries/item/categorysummaries/item/updatestatesummaries"
    iff4193ecc67d59ecea24f90311b1d6953f734fe79c0d8ba60c170a0fa5dd9408 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/macossoftwareupdateaccountsummaries/item/categorysummaries/item/updatestatesummaries/item"
)

// MacOSSoftwareUpdateCategorySummaryItemRequestBuilder builds and executes requests for operations under \deviceManagement\macOSSoftwareUpdateAccountSummaries\{macOSSoftwareUpdateAccountSummary-id}\categorySummaries\{macOSSoftwareUpdateCategorySummary-id}
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteOptions options for Delete
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetOptions options for Get
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters summary of the updates by category.
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchOptions options for Patch
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateCategorySummary;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal instantiates a new MacOSSoftwareUpdateCategorySummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    m := &MacOSSoftwareUpdateCategorySummaryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary_id}/categorySummaries/{macOSSoftwareUpdateCategorySummary_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilder instantiates a new MacOSSoftwareUpdateCategorySummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateDeleteRequestInformation(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateGetRequestInformation(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreatePatchRequestInformation(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Delete(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteOptions)(error) {
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
// Get summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Get(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateCategorySummary, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMacOSSoftwareUpdateCategorySummary() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateCategorySummary), nil
}
// Patch summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Patch(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchOptions)(error) {
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
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) UpdateStateSummaries()(*ia704cb6be080f1ecb092ccb2ba1c744783183d7b569f3801033f58df43a201db.UpdateStateSummariesRequestBuilder) {
    return ia704cb6be080f1ecb092ccb2ba1c744783183d7b569f3801033f58df43a201db.NewUpdateStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.macOSSoftwareUpdateAccountSummaries.item.categorySummaries.item.updateStateSummaries.item collection
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) UpdateStateSummariesById(id string)(*iff4193ecc67d59ecea24f90311b1d6953f734fe79c0d8ba60c170a0fa5dd9408.MacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["macOSSoftwareUpdateStateSummary_id"] = id
    }
    return iff4193ecc67d59ecea24f90311b1d6953f734fe79c0d8ba60c170a0fa5dd9408.NewMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
