package termstore

import (
    i0d491e0ed92613755e9ce87b7bcb61d5ac45e24bb3148376a23cd9790125ae69 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups"
    i8696d51b25dc2d6a8f09c66d2c3a8a89a556550f2c721719ed3f083e08d1e4cc "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8ab8307d36a2288f235f92a810e06d877f6459827fbbf4925375f0c10a27ed06 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item"
    ie85fd90aba6dfc1ffc92628ee3e7c7005a31cc9ec42b25f47257ab099e58a968 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item"
)

// TermStoreRequestBuilder builds and executes requests for operations under \termStore
type TermStoreRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TermStoreRequestBuilderGetOptions options for Get
type TermStoreRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TermStoreRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermStoreRequestBuilderGetQueryParameters get termStore
type TermStoreRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TermStoreRequestBuilderPatchOptions options for Patch
type TermStoreRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Store;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTermStoreRequestBuilderInternal instantiates a new TermStoreRequestBuilder and sets the default values.
func NewTermStoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermStoreRequestBuilder) {
    m := &TermStoreRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/termStore{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermStoreRequestBuilder instantiates a new TermStoreRequestBuilder and sets the default values.
func NewTermStoreRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermStoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermStoreRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get termStore
func (m *TermStoreRequestBuilder) CreateGetRequestInformation(options *TermStoreRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update termStore
func (m *TermStoreRequestBuilder) CreatePatchRequestInformation(options *TermStoreRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get termStore
func (m *TermStoreRequestBuilder) Get(options *TermStoreRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Store, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewStore() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Store), nil
}
func (m *TermStoreRequestBuilder) Groups()(*i0d491e0ed92613755e9ce87b7bcb61d5ac45e24bb3148376a23cd9790125ae69.GroupsRequestBuilder) {
    return i0d491e0ed92613755e9ce87b7bcb61d5ac45e24bb3148376a23cd9790125ae69.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.termStore.groups.item collection
func (m *TermStoreRequestBuilder) GroupsById(id string)(*ie85fd90aba6dfc1ffc92628ee3e7c7005a31cc9ec42b25f47257ab099e58a968.GroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group_id"] = id
    }
    return ie85fd90aba6dfc1ffc92628ee3e7c7005a31cc9ec42b25f47257ab099e58a968.NewGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update termStore
func (m *TermStoreRequestBuilder) Patch(options *TermStoreRequestBuilderPatchOptions)(error) {
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
func (m *TermStoreRequestBuilder) Sets()(*i8696d51b25dc2d6a8f09c66d2c3a8a89a556550f2c721719ed3f083e08d1e4cc.SetsRequestBuilder) {
    return i8696d51b25dc2d6a8f09c66d2c3a8a89a556550f2c721719ed3f083e08d1e4cc.NewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.termStore.sets.item collection
func (m *TermStoreRequestBuilder) SetsById(id string)(*i8ab8307d36a2288f235f92a810e06d877f6459827fbbf4925375f0c10a27ed06.SetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["set_id"] = id
    }
    return i8ab8307d36a2288f235f92a810e06d877f6459827fbbf4925375f0c10a27ed06.NewSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
