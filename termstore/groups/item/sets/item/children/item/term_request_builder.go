package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i64c233658bc03eb3c31220f413618c0c7b18cc6c3c017a86285263918e0af6ce "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/children/item/children"
    ic6ca3d8ce1cd0b88c740b19e66cf43b82fc76a0cdf592de5590ff11e2aaf404d "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/children/item/set"
    ie315ea8555250c9f65e699d095ae37f42a38e35f7da0d511fbb1248d175e9178 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/children/item/relations"
    i093b36b0596217767c3ff6f3f273143dd3b87f8cfcb379c62c21c3b7b4f7bb97 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/children/item/children/item"
    i6e54352c74f6c6e571034f0947cefb28b599d499e69bd6b12532f5c11ee08a59 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/children/item/relations/item"
)

// TermRequestBuilder builds and executes requests for operations under \termStore\groups\{group-id}\sets\{set-id}\children\{term-id}
type TermRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TermRequestBuilderDeleteOptions options for Delete
type TermRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermRequestBuilderGetOptions options for Get
type TermRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TermRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermRequestBuilderGetQueryParameters children terms of set in term [store].
type TermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TermRequestBuilderPatchOptions options for Patch
type TermRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Term;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TermRequestBuilder) Children()(*i64c233658bc03eb3c31220f413618c0c7b18cc6c3c017a86285263918e0af6ce.ChildrenRequestBuilder) {
    return i64c233658bc03eb3c31220f413618c0c7b18cc6c3c017a86285263918e0af6ce.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.termStore.groups.item.sets.item.children.item.children.item collection
func (m *TermRequestBuilder) ChildrenById(id string)(*i093b36b0596217767c3ff6f3f273143dd3b87f8cfcb379c62c21c3b7b4f7bb97.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id1"] = id
    }
    return i093b36b0596217767c3ff6f3f273143dd3b87f8cfcb379c62c21c3b7b4f7bb97.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermRequestBuilderInternal instantiates a new TermRequestBuilder and sets the default values.
func NewTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermRequestBuilder) {
    m := &TermRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/termStore/groups/{group_id}/sets/{set_id}/children/{term_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermRequestBuilder instantiates a new TermRequestBuilder and sets the default values.
func NewTermRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation children terms of set in term [store].
func (m *TermRequestBuilder) CreateDeleteRequestInformation(options *TermRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation children terms of set in term [store].
func (m *TermRequestBuilder) CreateGetRequestInformation(options *TermRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation children terms of set in term [store].
func (m *TermRequestBuilder) CreatePatchRequestInformation(options *TermRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete children terms of set in term [store].
func (m *TermRequestBuilder) Delete(options *TermRequestBuilderDeleteOptions)(error) {
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
// Get children terms of set in term [store].
func (m *TermRequestBuilder) Get(options *TermRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Term, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTerm() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Term), nil
}
// Patch children terms of set in term [store].
func (m *TermRequestBuilder) Patch(options *TermRequestBuilderPatchOptions)(error) {
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
func (m *TermRequestBuilder) Relations()(*ie315ea8555250c9f65e699d095ae37f42a38e35f7da0d511fbb1248d175e9178.RelationsRequestBuilder) {
    return ie315ea8555250c9f65e699d095ae37f42a38e35f7da0d511fbb1248d175e9178.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.termStore.groups.item.sets.item.children.item.relations.item collection
func (m *TermRequestBuilder) RelationsById(id string)(*i6e54352c74f6c6e571034f0947cefb28b599d499e69bd6b12532f5c11ee08a59.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return i6e54352c74f6c6e571034f0947cefb28b599d499e69bd6b12532f5c11ee08a59.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermRequestBuilder) Set()(*ic6ca3d8ce1cd0b88c740b19e66cf43b82fc76a0cdf592de5590ff11e2aaf404d.SetRequestBuilder) {
    return ic6ca3d8ce1cd0b88c740b19e66cf43b82fc76a0cdf592de5590ff11e2aaf404d.NewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
