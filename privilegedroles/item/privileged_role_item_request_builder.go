package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04754c5841c9fcefc803d8096b9d648d4a4f8e797cd7d5703b0b2fa021349d4f "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/settings"
    i30870af67a7c7e750abaa30f5e84eff1a098db84ecc2ae53c8193d533bc3e771 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/assignments"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ib362bf055682cb60eebaecadb64541d74da0307dbd607faad741d862fab711bf "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/summary"
    ib4b9c3bdee734abec86c1274b61654a23bf76937e3c1b4deb6b3a5baf179dcbe "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/selfdeactivate"
    ifc2ecb6356f4c50497a05265b4b0a311af52bfc4aa6302f476a92d1df6adfafb "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/selfactivate"
    i0aae3989cf245a50ac6c762e9583dab23911e23dcb918057ff75fd93264ed4ca "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/assignments/item"
)

// PrivilegedRoleItemRequestBuilder provides operations to manage the collection of privilegedRole entities.
type PrivilegedRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrivilegedRoleItemRequestBuilderDeleteOptions options for Delete
type PrivilegedRoleItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedRoleItemRequestBuilderGetOptions options for Get
type PrivilegedRoleItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrivilegedRoleItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedRoleItemRequestBuilderGetQueryParameters get entity from privilegedRoles by key
type PrivilegedRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrivilegedRoleItemRequestBuilderPatchOptions options for Patch
type PrivilegedRoleItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrivilegedRoleItemRequestBuilder) Assignments()(*i30870af67a7c7e750abaa30f5e84eff1a098db84ecc2ae53c8193d533bc3e771.AssignmentsRequestBuilder) {
    return i30870af67a7c7e750abaa30f5e84eff1a098db84ecc2ae53c8193d533bc3e771.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedRoles.item.assignments.item collection
func (m *PrivilegedRoleItemRequestBuilder) AssignmentsById(id string)(*i0aae3989cf245a50ac6c762e9583dab23911e23dcb918057ff75fd93264ed4ca.PrivilegedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedRoleAssignment_id"] = id
    }
    return i0aae3989cf245a50ac6c762e9583dab23911e23dcb918057ff75fd93264ed4ca.NewPrivilegedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrivilegedRoleItemRequestBuilderInternal instantiates a new PrivilegedRoleItemRequestBuilder and sets the default values.
func NewPrivilegedRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedRoleItemRequestBuilder) {
    m := &PrivilegedRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoles/{privilegedRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedRoleItemRequestBuilder instantiates a new PrivilegedRoleItemRequestBuilder and sets the default values.
func NewPrivilegedRoleItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from privilegedRoles
func (m *PrivilegedRoleItemRequestBuilder) CreateDeleteRequestInformation(options *PrivilegedRoleItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from privilegedRoles by key
func (m *PrivilegedRoleItemRequestBuilder) CreateGetRequestInformation(options *PrivilegedRoleItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in privilegedRoles
func (m *PrivilegedRoleItemRequestBuilder) CreatePatchRequestInformation(options *PrivilegedRoleItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from privilegedRoles
func (m *PrivilegedRoleItemRequestBuilder) Delete(options *PrivilegedRoleItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from privilegedRoles by key
func (m *PrivilegedRoleItemRequestBuilder) Get(options *PrivilegedRoleItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreatePrivilegedRoleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleable), nil
}
// Patch update entity in privilegedRoles
func (m *PrivilegedRoleItemRequestBuilder) Patch(options *PrivilegedRoleItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *PrivilegedRoleItemRequestBuilder) SelfActivate()(*ifc2ecb6356f4c50497a05265b4b0a311af52bfc4aa6302f476a92d1df6adfafb.SelfActivateRequestBuilder) {
    return ifc2ecb6356f4c50497a05265b4b0a311af52bfc4aa6302f476a92d1df6adfafb.NewSelfActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedRoleItemRequestBuilder) SelfDeactivate()(*ib4b9c3bdee734abec86c1274b61654a23bf76937e3c1b4deb6b3a5baf179dcbe.SelfDeactivateRequestBuilder) {
    return ib4b9c3bdee734abec86c1274b61654a23bf76937e3c1b4deb6b3a5baf179dcbe.NewSelfDeactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedRoleItemRequestBuilder) Settings()(*i04754c5841c9fcefc803d8096b9d648d4a4f8e797cd7d5703b0b2fa021349d4f.SettingsRequestBuilder) {
    return i04754c5841c9fcefc803d8096b9d648d4a4f8e797cd7d5703b0b2fa021349d4f.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedRoleItemRequestBuilder) Summary()(*ib362bf055682cb60eebaecadb64541d74da0307dbd607faad741d862fab711bf.SummaryRequestBuilder) {
    return ib362bf055682cb60eebaecadb64541d74da0307dbd607faad741d862fab711bf.NewSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
