package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4fa8edb0c045897c1eedba2d597f16b1fbe1eba86c010bbb3f99622813ad295e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/requests"
    ia0161922114d3b8c7cd781d7f9c4106c03db7e9ef8ea72513b4c54b5e9354ba5 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/operations"
    iac4781d78e0362023abac407b248abb8996da024ec94a9413fc9a3d80755d0ea "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/accessassignments"
    i0580a758f5bfbd03e74223f96f6b1bdb17dd6ffec4a223b1f1fabb8a627a42b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/requests/item"
    i5f3e9a785873705d5204dd08bc17b79ce358f2fd632fbe5eb537c2a5d7fc7ef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/operations/item"
    ic1a2347ea4abb46ab983c9d6ee916f3a9af60bf8f4caf7afae2666a77def3567 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/accessassignments/item"
)

// DelegatedAdminRelationshipItemRequestBuilder builds and executes requests for operations under \tenantRelationships\delegatedAdminRelationships\{delegatedAdminRelationship-id}
type DelegatedAdminRelationshipItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DelegatedAdminRelationshipItemRequestBuilderDeleteOptions options for Delete
type DelegatedAdminRelationshipItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DelegatedAdminRelationshipItemRequestBuilderGetOptions options for Get
type DelegatedAdminRelationshipItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DelegatedAdminRelationshipItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DelegatedAdminRelationshipItemRequestBuilderGetQueryParameters get delegatedAdminRelationships from tenantRelationships
type DelegatedAdminRelationshipItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DelegatedAdminRelationshipItemRequestBuilderPatchOptions options for Patch
type DelegatedAdminRelationshipItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DelegatedAdminRelationship;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DelegatedAdminRelationshipItemRequestBuilder) AccessAssignments()(*iac4781d78e0362023abac407b248abb8996da024ec94a9413fc9a3d80755d0ea.AccessAssignmentsRequestBuilder) {
    return iac4781d78e0362023abac407b248abb8996da024ec94a9413fc9a3d80755d0ea.NewAccessAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.delegatedAdminRelationships.item.accessAssignments.item collection
func (m *DelegatedAdminRelationshipItemRequestBuilder) AccessAssignmentsById(id string)(*ic1a2347ea4abb46ab983c9d6ee916f3a9af60bf8f4caf7afae2666a77def3567.DelegatedAdminAccessAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminAccessAssignment_id"] = id
    }
    return ic1a2347ea4abb46ab983c9d6ee916f3a9af60bf8f4caf7afae2666a77def3567.NewDelegatedAdminAccessAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDelegatedAdminRelationshipItemRequestBuilderInternal instantiates a new DelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DelegatedAdminRelationshipItemRequestBuilder) {
    m := &DelegatedAdminRelationshipItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDelegatedAdminRelationshipItemRequestBuilder instantiates a new DelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DelegatedAdminRelationshipItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedAdminRelationshipItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property delegatedAdminRelationships for tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) CreateDeleteRequestInformation(options *DelegatedAdminRelationshipItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get delegatedAdminRelationships from tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) CreateGetRequestInformation(options *DelegatedAdminRelationshipItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property delegatedAdminRelationships in tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) CreatePatchRequestInformation(options *DelegatedAdminRelationshipItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property delegatedAdminRelationships for tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) Delete(options *DelegatedAdminRelationshipItemRequestBuilderDeleteOptions)(error) {
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
// Get get delegatedAdminRelationships from tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) Get(options *DelegatedAdminRelationshipItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DelegatedAdminRelationship, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDelegatedAdminRelationship() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DelegatedAdminRelationship), nil
}
func (m *DelegatedAdminRelationshipItemRequestBuilder) Operations()(*ia0161922114d3b8c7cd781d7f9c4106c03db7e9ef8ea72513b4c54b5e9354ba5.OperationsRequestBuilder) {
    return ia0161922114d3b8c7cd781d7f9c4106c03db7e9ef8ea72513b4c54b5e9354ba5.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.delegatedAdminRelationships.item.operations.item collection
func (m *DelegatedAdminRelationshipItemRequestBuilder) OperationsById(id string)(*i5f3e9a785873705d5204dd08bc17b79ce358f2fd632fbe5eb537c2a5d7fc7ef6.DelegatedAdminRelationshipOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationshipOperation_id"] = id
    }
    return i5f3e9a785873705d5204dd08bc17b79ce358f2fd632fbe5eb537c2a5d7fc7ef6.NewDelegatedAdminRelationshipOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property delegatedAdminRelationships in tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) Patch(options *DelegatedAdminRelationshipItemRequestBuilderPatchOptions)(error) {
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
func (m *DelegatedAdminRelationshipItemRequestBuilder) Requests()(*i4fa8edb0c045897c1eedba2d597f16b1fbe1eba86c010bbb3f99622813ad295e.RequestsRequestBuilder) {
    return i4fa8edb0c045897c1eedba2d597f16b1fbe1eba86c010bbb3f99622813ad295e.NewRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.delegatedAdminRelationships.item.requests.item collection
func (m *DelegatedAdminRelationshipItemRequestBuilder) RequestsById(id string)(*i0580a758f5bfbd03e74223f96f6b1bdb17dd6ffec4a223b1f1fabb8a627a42b4.DelegatedAdminRelationshipRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationshipRequest_id"] = id
    }
    return i0580a758f5bfbd03e74223f96f6b1bdb17dd6ffec4a223b1f1fabb8a627a42b4.NewDelegatedAdminRelationshipRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
