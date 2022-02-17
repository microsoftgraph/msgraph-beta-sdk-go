package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04754c5841c9fcefc803d8096b9d648d4a4f8e797cd7d5703b0b2fa021349d4f "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/settings"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i30870af67a7c7e750abaa30f5e84eff1a098db84ecc2ae53c8193d533bc3e771 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/assignments"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ib362bf055682cb60eebaecadb64541d74da0307dbd607faad741d862fab711bf "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/summary"
    ib4b9c3bdee734abec86c1274b61654a23bf76937e3c1b4deb6b3a5baf179dcbe "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/selfdeactivate"
    ifc2ecb6356f4c50497a05265b4b0a311af52bfc4aa6302f476a92d1df6adfafb "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item/selfactivate"
)

// PrivilegedRoleRequestBuilder builds and executes requests for operations under \privilegedRoles\{privilegedRole-id}
type PrivilegedRoleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrivilegedRoleRequestBuilderDeleteOptions options for Delete
type PrivilegedRoleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedRoleRequestBuilderGetOptions options for Get
type PrivilegedRoleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrivilegedRoleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedRoleRequestBuilderGetQueryParameters get entity from privilegedRoles by key
type PrivilegedRoleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrivilegedRoleRequestBuilderPatchOptions options for Patch
type PrivilegedRoleRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRole;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrivilegedRoleRequestBuilder) Assignments()(*i30870af67a7c7e750abaa30f5e84eff1a098db84ecc2ae53c8193d533bc3e771.AssignmentsRequestBuilder) {
    return i30870af67a7c7e750abaa30f5e84eff1a098db84ecc2ae53c8193d533bc3e771.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrivilegedRoleRequestBuilderInternal instantiates a new PrivilegedRoleRequestBuilder and sets the default values.
func NewPrivilegedRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedRoleRequestBuilder) {
    m := &PrivilegedRoleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoles/{privilegedRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedRoleRequestBuilder instantiates a new PrivilegedRoleRequestBuilder and sets the default values.
func NewPrivilegedRoleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedRoleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from privilegedRoles
func (m *PrivilegedRoleRequestBuilder) CreateDeleteRequestInformation(options *PrivilegedRoleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrivilegedRoleRequestBuilder) CreateGetRequestInformation(options *PrivilegedRoleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrivilegedRoleRequestBuilder) CreatePatchRequestInformation(options *PrivilegedRoleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrivilegedRoleRequestBuilder) Delete(options *PrivilegedRoleRequestBuilderDeleteOptions)(error) {
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
// Get get entity from privilegedRoles by key
func (m *PrivilegedRoleRequestBuilder) Get(options *PrivilegedRoleRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRole, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrivilegedRole() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRole), nil
}
// Patch update entity in privilegedRoles
func (m *PrivilegedRoleRequestBuilder) Patch(options *PrivilegedRoleRequestBuilderPatchOptions)(error) {
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
func (m *PrivilegedRoleRequestBuilder) SelfActivate()(*ifc2ecb6356f4c50497a05265b4b0a311af52bfc4aa6302f476a92d1df6adfafb.SelfActivateRequestBuilder) {
    return ifc2ecb6356f4c50497a05265b4b0a311af52bfc4aa6302f476a92d1df6adfafb.NewSelfActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedRoleRequestBuilder) SelfDeactivate()(*ib4b9c3bdee734abec86c1274b61654a23bf76937e3c1b4deb6b3a5baf179dcbe.SelfDeactivateRequestBuilder) {
    return ib4b9c3bdee734abec86c1274b61654a23bf76937e3c1b4deb6b3a5baf179dcbe.NewSelfDeactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedRoleRequestBuilder) Settings()(*i04754c5841c9fcefc803d8096b9d648d4a4f8e797cd7d5703b0b2fa021349d4f.SettingsRequestBuilder) {
    return i04754c5841c9fcefc803d8096b9d648d4a4f8e797cd7d5703b0b2fa021349d4f.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedRoleRequestBuilder) Summary()(*ib362bf055682cb60eebaecadb64541d74da0307dbd607faad741d862fab711bf.SummaryRequestBuilder) {
    return ib362bf055682cb60eebaecadb64541d74da0307dbd607faad741d862fab711bf.NewSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
