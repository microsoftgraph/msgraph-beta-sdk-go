package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0487123d543cf2f0455b73d36cb4ef6095f250b2893728cfe95a7fe68de3c949 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/principals"
    i32f034c8f89cc3928550c99a5348bf56a4c9d54d81b39aef7a3a2c91e22f1039 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/roledefinition"
    i9fbcca2aad84715b68bcf5e9b22ee6957d1813ee30a1a07b635e14deecf7a4c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/appscopes"
    icce9ff4f56da3b598f62227e100054c17f721d52dc3ccdf101cd0e269e75ecf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/directoryscopes"
    i4c3467330b742e9b38808f445f2b83de61398d9e7eb0c4191bdd72e252c98629 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/appscopes/item"
)

// UnifiedRoleAssignmentMultipleRequestBuilder builds and executes requests for operations under \roleManagement\deviceManagement\roleAssignments\{unifiedRoleAssignmentMultiple-id}
type UnifiedRoleAssignmentMultipleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentMultipleRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentMultipleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentMultipleRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentMultipleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleAssignmentMultipleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentMultipleRequestBuilderGetQueryParameters get roleAssignments from roleManagement
type UnifiedRoleAssignmentMultipleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentMultipleRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentMultipleRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentMultiple;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) AppScopes()(*i9fbcca2aad84715b68bcf5e9b22ee6957d1813ee30a1a07b635e14deecf7a4c8.AppScopesRequestBuilder) {
    return i9fbcca2aad84715b68bcf5e9b22ee6957d1813ee30a1a07b635e14deecf7a4c8.NewAppScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.deviceManagement.roleAssignments.item.appScopes.item collection
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) AppScopesById(id string)(*i4c3467330b742e9b38808f445f2b83de61398d9e7eb0c4191bdd72e252c98629.AppScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appScope_id"] = id
    }
    return i4c3467330b742e9b38808f445f2b83de61398d9e7eb0c4191bdd72e252c98629.NewAppScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUnifiedRoleAssignmentMultipleRequestBuilderInternal instantiates a new UnifiedRoleAssignmentMultipleRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentMultipleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentMultipleRequestBuilder) {
    m := &UnifiedRoleAssignmentMultipleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement/roleAssignments/{unifiedRoleAssignmentMultiple_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentMultipleRequestBuilder instantiates a new UnifiedRoleAssignmentMultipleRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentMultipleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentMultipleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentMultipleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for roleManagement
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentMultipleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get roleAssignments from roleManagement
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentMultipleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignments in roleManagement
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentMultipleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property roleAssignments for roleManagement
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) Delete(options *UnifiedRoleAssignmentMultipleRequestBuilderDeleteOptions)(error) {
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
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) DirectoryScopes()(*icce9ff4f56da3b598f62227e100054c17f721d52dc3ccdf101cd0e269e75ecf4.DirectoryScopesRequestBuilder) {
    return icce9ff4f56da3b598f62227e100054c17f721d52dc3ccdf101cd0e269e75ecf4.NewDirectoryScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleAssignments from roleManagement
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) Get(options *UnifiedRoleAssignmentMultipleRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentMultiple, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleAssignmentMultiple() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentMultiple), nil
}
// Patch update the navigation property roleAssignments in roleManagement
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) Patch(options *UnifiedRoleAssignmentMultipleRequestBuilderPatchOptions)(error) {
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
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) Principals()(*i0487123d543cf2f0455b73d36cb4ef6095f250b2893728cfe95a7fe68de3c949.PrincipalsRequestBuilder) {
    return i0487123d543cf2f0455b73d36cb4ef6095f250b2893728cfe95a7fe68de3c949.NewPrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentMultipleRequestBuilder) RoleDefinition()(*i32f034c8f89cc3928550c99a5348bf56a4c9d54d81b39aef7a3a2c91e22f1039.RoleDefinitionRequestBuilder) {
    return i32f034c8f89cc3928550c99a5348bf56a4c9d54d81b39aef7a3a2c91e22f1039.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
