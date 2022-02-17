package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i06f1e08316e6d513040d571df20771afa3cc90759cf7871a87710d28106c203c "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms/item/children"
    i3d70c03e1e62301e655e05ff1becef8400eb2551bf3219d1961154b984225eb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms/item/relations"
    i5ff89e5064891af1b042e3e552b31cedfc17419b618b0bf108539b315a42e9e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms/item/set"
    i261bf93a735d97835147d5a6b001ad93504bcec36e8f906b0cebe5b07320adb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms/item/relations/item"
    i8725ca95c2031f0753cea01c7c79e98f0edd64d78fc303b128fc552c601a2453 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms/item/children/item"
)

// TermRequestBuilder builds and executes requests for operations under \termStore\groups\{group-id}\sets\{set-id}\terms\{term-id}
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
// TermRequestBuilderGetQueryParameters all the terms under the set.
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
func (m *TermRequestBuilder) Children()(*i06f1e08316e6d513040d571df20771afa3cc90759cf7871a87710d28106c203c.ChildrenRequestBuilder) {
    return i06f1e08316e6d513040d571df20771afa3cc90759cf7871a87710d28106c203c.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.termStore.groups.item.sets.item.terms.item.children.item collection
func (m *TermRequestBuilder) ChildrenById(id string)(*i8725ca95c2031f0753cea01c7c79e98f0edd64d78fc303b128fc552c601a2453.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id1"] = id
    }
    return i8725ca95c2031f0753cea01c7c79e98f0edd64d78fc303b128fc552c601a2453.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermRequestBuilderInternal instantiates a new TermRequestBuilder and sets the default values.
func NewTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermRequestBuilder) {
    m := &TermRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/termStore/groups/{group_id}/sets/{set_id}/terms/{term_id}{?select,expand}";
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
// CreateDeleteRequestInformation all the terms under the set.
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
// CreateGetRequestInformation all the terms under the set.
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
// CreatePatchRequestInformation all the terms under the set.
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
// Delete all the terms under the set.
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
// Get all the terms under the set.
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
// Patch all the terms under the set.
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
func (m *TermRequestBuilder) Relations()(*i3d70c03e1e62301e655e05ff1becef8400eb2551bf3219d1961154b984225eb1.RelationsRequestBuilder) {
    return i3d70c03e1e62301e655e05ff1becef8400eb2551bf3219d1961154b984225eb1.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.termStore.groups.item.sets.item.terms.item.relations.item collection
func (m *TermRequestBuilder) RelationsById(id string)(*i261bf93a735d97835147d5a6b001ad93504bcec36e8f906b0cebe5b07320adb0.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return i261bf93a735d97835147d5a6b001ad93504bcec36e8f906b0cebe5b07320adb0.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermRequestBuilder) Set()(*i5ff89e5064891af1b042e3e552b31cedfc17419b618b0bf108539b315a42e9e0.SetRequestBuilder) {
    return i5ff89e5064891af1b042e3e552b31cedfc17419b618b0bf108539b315a42e9e0.NewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
