package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ibe61fbd02c16db3e5f015184d8991d363a774ed55310ae904b00e97143389ee8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/groupassignments"
    ide71c0a56dd4ade2d81104d2ecb208e8db8c2910aaea2f231edb2cf623477c53 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/assignments"
    idf88d5dcb4048ac7ff79419fd15c7d58e83a14f8b409ce16869ccf4204ab829f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses"
    i5c99d94fd8ce128189bc114b673df4e56578860abd3b3eadccba77c18952a47d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/groupassignments/item"
    i605e35cc4db71e7fa4db1f4759e05ab605fd4596a0a820b7a556521d22026745 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses/item"
    idae59aa73e78d2aa3e7943286419e2c7b77cd5e337f001c8305c745502c3740f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/assignments/item"
)

// TermsAndConditionsItemRequestBuilder builds and executes requests for operations under \deviceManagement\termsAndConditions\{termsAndConditions-id}
type TermsAndConditionsItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TermsAndConditionsItemRequestBuilderDeleteOptions options for Delete
type TermsAndConditionsItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsAndConditionsItemRequestBuilderGetOptions options for Get
type TermsAndConditionsItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TermsAndConditionsItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsAndConditionsItemRequestBuilderGetQueryParameters the terms and conditions associated with device management of the company.
type TermsAndConditionsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TermsAndConditionsItemRequestBuilderPatchOptions options for Patch
type TermsAndConditionsItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TermsAndConditions;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TermsAndConditionsItemRequestBuilder) AcceptanceStatuses()(*idf88d5dcb4048ac7ff79419fd15c7d58e83a14f8b409ce16869ccf4204ab829f.AcceptanceStatusesRequestBuilder) {
    return idf88d5dcb4048ac7ff79419fd15c7d58e83a14f8b409ce16869ccf4204ab829f.NewAcceptanceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptanceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.termsAndConditions.item.acceptanceStatuses.item collection
func (m *TermsAndConditionsItemRequestBuilder) AcceptanceStatusesById(id string)(*i605e35cc4db71e7fa4db1f4759e05ab605fd4596a0a820b7a556521d22026745.TermsAndConditionsAcceptanceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAcceptanceStatus_id"] = id
    }
    return i605e35cc4db71e7fa4db1f4759e05ab605fd4596a0a820b7a556521d22026745.NewTermsAndConditionsAcceptanceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermsAndConditionsItemRequestBuilder) Assignments()(*ide71c0a56dd4ade2d81104d2ecb208e8db8c2910aaea2f231edb2cf623477c53.AssignmentsRequestBuilder) {
    return ide71c0a56dd4ade2d81104d2ecb208e8db8c2910aaea2f231edb2cf623477c53.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.termsAndConditions.item.assignments.item collection
func (m *TermsAndConditionsItemRequestBuilder) AssignmentsById(id string)(*idae59aa73e78d2aa3e7943286419e2c7b77cd5e337f001c8305c745502c3740f.TermsAndConditionsAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAssignment_id"] = id
    }
    return idae59aa73e78d2aa3e7943286419e2c7b77cd5e337f001c8305c745502c3740f.NewTermsAndConditionsAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermsAndConditionsItemRequestBuilderInternal instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewTermsAndConditionsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsAndConditionsItemRequestBuilder) {
    m := &TermsAndConditionsItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermsAndConditionsItemRequestBuilder instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewTermsAndConditionsItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsAndConditionsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsAndConditionsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) CreateDeleteRequestInformation(options *TermsAndConditionsItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) CreateGetRequestInformation(options *TermsAndConditionsItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) CreatePatchRequestInformation(options *TermsAndConditionsItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) Delete(options *TermsAndConditionsItemRequestBuilderDeleteOptions)(error) {
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
// Get the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) Get(options *TermsAndConditionsItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TermsAndConditions, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTermsAndConditions() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TermsAndConditions), nil
}
func (m *TermsAndConditionsItemRequestBuilder) GroupAssignments()(*ibe61fbd02c16db3e5f015184d8991d363a774ed55310ae904b00e97143389ee8.GroupAssignmentsRequestBuilder) {
    return ibe61fbd02c16db3e5f015184d8991d363a774ed55310ae904b00e97143389ee8.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.termsAndConditions.item.groupAssignments.item collection
func (m *TermsAndConditionsItemRequestBuilder) GroupAssignmentsById(id string)(*i5c99d94fd8ce128189bc114b673df4e56578860abd3b3eadccba77c18952a47d.TermsAndConditionsGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsGroupAssignment_id"] = id
    }
    return i5c99d94fd8ce128189bc114b673df4e56578860abd3b3eadccba77c18952a47d.NewTermsAndConditionsGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) Patch(options *TermsAndConditionsItemRequestBuilderPatchOptions)(error) {
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
