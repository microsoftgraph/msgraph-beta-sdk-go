package rolescopetags

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4edbe639b32662e436fc1a2df91f2c1f4f9b92866c08a7d5338db735db338052 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/roleassignments/item/rolescopetags/getrolescopetagsbyid"
    i59b421f7e62c770e57ae33b84588a1af6e28a2f7b0dacd4d0f5e6d5590f30f65 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/roleassignments/item/rolescopetags/hascustomrolescopetag"
    ie1c6af2bb765fdf2e7854c03f501167604144eaec4781e2148d60afb3225dd4f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/roleassignments/item/rolescopetags/ref"
)

// RoleScopeTagsRequestBuilder builds and executes requests for operations under \deviceManagement\roleAssignments\{deviceAndAppManagementRoleAssignment-id}\roleScopeTags
type RoleScopeTagsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleScopeTagsRequestBuilderGetOptions options for Get
type RoleScopeTagsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleScopeTagsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleScopeTagsRequestBuilderGetQueryParameters the set of Role Scope Tags defined on the Role Assignment.
type RoleScopeTagsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewRoleScopeTagsRequestBuilderInternal instantiates a new RoleScopeTagsRequestBuilder and sets the default values.
func NewRoleScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleScopeTagsRequestBuilder) {
    m := &RoleScopeTagsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment_id}/roleScopeTags{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleScopeTagsRequestBuilder instantiates a new RoleScopeTagsRequestBuilder and sets the default values.
func NewRoleScopeTagsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the set of Role Scope Tags defined on the Role Assignment.
func (m *RoleScopeTagsRequestBuilder) CreateGetRequestInformation(options *RoleScopeTagsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the set of Role Scope Tags defined on the Role Assignment.
func (m *RoleScopeTagsRequestBuilder) Get(options *RoleScopeTagsRequestBuilderGetOptions)(*RoleScopeTagsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleScopeTagsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RoleScopeTagsResponse), nil
}
func (m *RoleScopeTagsRequestBuilder) GetRoleScopeTagsById()(*i4edbe639b32662e436fc1a2df91f2c1f4f9b92866c08a7d5338db735db338052.GetRoleScopeTagsByIdRequestBuilder) {
    return i4edbe639b32662e436fc1a2df91f2c1f4f9b92866c08a7d5338db735db338052.NewGetRoleScopeTagsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HasCustomRoleScopeTag builds and executes requests for operations under \deviceManagement\roleAssignments\{deviceAndAppManagementRoleAssignment-id}\roleScopeTags\microsoft.graph.hasCustomRoleScopeTag()
func (m *RoleScopeTagsRequestBuilder) HasCustomRoleScopeTag()(*i59b421f7e62c770e57ae33b84588a1af6e28a2f7b0dacd4d0f5e6d5590f30f65.HasCustomRoleScopeTagRequestBuilder) {
    return i59b421f7e62c770e57ae33b84588a1af6e28a2f7b0dacd4d0f5e6d5590f30f65.NewHasCustomRoleScopeTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RoleScopeTagsRequestBuilder) Ref()(*ie1c6af2bb765fdf2e7854c03f501167604144eaec4781e2148d60afb3225dd4f.RefRequestBuilder) {
    return ie1c6af2bb765fdf2e7854c03f501167604144eaec4781e2148d60afb3225dd4f.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
