package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i0487123d543cf2f0455b73d36cb4ef6095f250b2893728cfe95a7fe68de3c949 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/principals"
    i32f034c8f89cc3928550c99a5348bf56a4c9d54d81b39aef7a3a2c91e22f1039 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/roledefinition"
    i9fbcca2aad84715b68bcf5e9b22ee6957d1813ee30a1a07b635e14deecf7a4c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/appscopes"
    icce9ff4f56da3b598f62227e100054c17f721d52dc3ccdf101cd0e269e75ecf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/directoryscopes"
    i4c3467330b742e9b38808f445f2b83de61398d9e7eb0c4191bdd72e252c98629 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/appscopes/item"
    i7ef7ffeaa7ffa2b753fb87022fd2626a077098ef248233b1f62a9e1b12db79f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/principals/item"
    ie5c2834d8c0f569f8e6b58184f56ea01d180c0a565481cde156d1ab24bd75da1 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/directoryscopes/item"
)

// UnifiedRoleAssignmentMultipleItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplicationMultiple entity.
type UnifiedRoleAssignmentMultipleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentMultipleItemRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentMultipleItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentMultipleItemRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentMultipleItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters get roleAssignments from roleManagement
type UnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentMultipleItemRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentMultipleItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentMultipleable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopes()(*i9fbcca2aad84715b68bcf5e9b22ee6957d1813ee30a1a07b635e14deecf7a4c8.AppScopesRequestBuilder) {
    return i9fbcca2aad84715b68bcf5e9b22ee6957d1813ee30a1a07b635e14deecf7a4c8.NewAppScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.deviceManagement.roleAssignments.item.appScopes.item collection
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopesById(id string)(*i4c3467330b742e9b38808f445f2b83de61398d9e7eb0c4191bdd72e252c98629.AppScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appScope_id"] = id
    }
    return i4c3467330b742e9b38808f445f2b83de61398d9e7eb0c4191bdd72e252c98629.NewAppScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUnifiedRoleAssignmentMultipleItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentMultipleItemRequestBuilder) {
    m := &UnifiedRoleAssignmentMultipleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement/roleAssignments/{unifiedRoleAssignmentMultiple_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentMultipleItemRequestBuilder instantiates a new UnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentMultipleItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentMultipleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for roleManagement
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentMultipleItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentMultipleItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentMultipleItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) Delete(options *UnifiedRoleAssignmentMultipleItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) DirectoryScopes()(*icce9ff4f56da3b598f62227e100054c17f721d52dc3ccdf101cd0e269e75ecf4.DirectoryScopesRequestBuilder) {
    return icce9ff4f56da3b598f62227e100054c17f721d52dc3ccdf101cd0e269e75ecf4.NewDirectoryScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.deviceManagement.roleAssignments.item.directoryScopes.item collection
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) DirectoryScopesById(id string)(*ie5c2834d8c0f569f8e6b58184f56ea01d180c0a565481cde156d1ab24bd75da1.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ie5c2834d8c0f569f8e6b58184f56ea01d180c0a565481cde156d1ab24bd75da1.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get roleAssignments from roleManagement
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) Get(options *UnifiedRoleAssignmentMultipleItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentMultipleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentMultipleable), nil
}
// Patch update the navigation property roleAssignments in roleManagement
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) Patch(options *UnifiedRoleAssignmentMultipleItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) Principals()(*i0487123d543cf2f0455b73d36cb4ef6095f250b2893728cfe95a7fe68de3c949.PrincipalsRequestBuilder) {
    return i0487123d543cf2f0455b73d36cb4ef6095f250b2893728cfe95a7fe68de3c949.NewPrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrincipalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.deviceManagement.roleAssignments.item.principals.item collection
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) PrincipalsById(id string)(*i7ef7ffeaa7ffa2b753fb87022fd2626a077098ef248233b1f62a9e1b12db79f7.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i7ef7ffeaa7ffa2b753fb87022fd2626a077098ef248233b1f62a9e1b12db79f7.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) RoleDefinition()(*i32f034c8f89cc3928550c99a5348bf56a4c9d54d81b39aef7a3a2c91e22f1039.RoleDefinitionRequestBuilder) {
    return i32f034c8f89cc3928550c99a5348bf56a4c9d54d81b39aef7a3a2c91e22f1039.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
