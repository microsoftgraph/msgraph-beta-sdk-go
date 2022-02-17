package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i114eb222f1988e9aa4e6ceab0dfa7bc1d169caae2302e09152051c4056697fe1 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms/item/relations/item/fromterm"
    i4ba309f3af9caaf0d5cb51e5b4a0d41ac2421a52548e3388924aefb487c1ed5b "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms/item/relations/item/toterm"
    icaa86f1ed1f02df176ff856607bc76793ab15ed45b44c35c175177694a8347b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms/item/relations/item/set"
)

// RelationRequestBuilder builds and executes requests for operations under \termStore\groups\{group-id}\sets\{set-id}\terms\{term-id}\relations\{relation-id}
type RelationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RelationRequestBuilderDeleteOptions options for Delete
type RelationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RelationRequestBuilderGetOptions options for Get
type RelationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RelationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RelationRequestBuilderGetQueryParameters to indicate which terms are related to the current term as either pinned or reused.
type RelationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RelationRequestBuilderPatchOptions options for Patch
type RelationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Relation;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRelationRequestBuilderInternal instantiates a new RelationRequestBuilder and sets the default values.
func NewRelationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RelationRequestBuilder) {
    m := &RelationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/termStore/groups/{group_id}/sets/{set_id}/terms/{term_id}/relations/{relation_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRelationRequestBuilder instantiates a new RelationRequestBuilder and sets the default values.
func NewRelationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RelationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRelationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation to indicate which terms are related to the current term as either pinned or reused.
func (m *RelationRequestBuilder) CreateDeleteRequestInformation(options *RelationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation to indicate which terms are related to the current term as either pinned or reused.
func (m *RelationRequestBuilder) CreateGetRequestInformation(options *RelationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation to indicate which terms are related to the current term as either pinned or reused.
func (m *RelationRequestBuilder) CreatePatchRequestInformation(options *RelationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete to indicate which terms are related to the current term as either pinned or reused.
func (m *RelationRequestBuilder) Delete(options *RelationRequestBuilderDeleteOptions)(error) {
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
func (m *RelationRequestBuilder) FromTerm()(*i114eb222f1988e9aa4e6ceab0dfa7bc1d169caae2302e09152051c4056697fe1.FromTermRequestBuilder) {
    return i114eb222f1988e9aa4e6ceab0dfa7bc1d169caae2302e09152051c4056697fe1.NewFromTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get to indicate which terms are related to the current term as either pinned or reused.
func (m *RelationRequestBuilder) Get(options *RelationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Relation, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRelation() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Relation), nil
}
// Patch to indicate which terms are related to the current term as either pinned or reused.
func (m *RelationRequestBuilder) Patch(options *RelationRequestBuilderPatchOptions)(error) {
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
func (m *RelationRequestBuilder) Set()(*icaa86f1ed1f02df176ff856607bc76793ab15ed45b44c35c175177694a8347b4.SetRequestBuilder) {
    return icaa86f1ed1f02df176ff856607bc76793ab15ed45b44c35c175177694a8347b4.NewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RelationRequestBuilder) ToTerm()(*i4ba309f3af9caaf0d5cb51e5b4a0d41ac2421a52548e3388924aefb487c1ed5b.ToTermRequestBuilder) {
    return i4ba309f3af9caaf0d5cb51e5b4a0d41ac2421a52548e3388924aefb487c1ed5b.NewToTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
