package cloudpc

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i21e54f9ee2d47268e20afd63152400c9b4c1f5fccc344e57fac33bd5bafc27e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments"
    i31977c399a73e8dceaa0f68ce2147b8b69cdb6d448081c8df74e04127a6f5b7e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roledefinitions"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id0c942d6747644ac6d906984e3b3d0046ec64b19ed0e042c591f0eda0c0f8957 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/resourcenamespaces"
    i5122da689a05212ccc762ebf1a452cf0621ee3e3ee2ed9c79b323f7d87ab82ba "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments/item"
    i6596aa8a9eb296ce0e4c8af117114f5b675a99379d789079c98d15d2b4b41cb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roledefinitions/item"
    if2fe6d33056a86a8c02bd09d388e1ad97378adc4d03f5d4afe21fd7135b39314 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/resourcenamespaces/item"
)

// CloudPCRequestBuilder builds and executes requests for operations under \roleManagement\cloudPC
type CloudPCRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CloudPCRequestBuilderDeleteOptions options for Delete
type CloudPCRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPCRequestBuilderGetOptions options for Get
type CloudPCRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CloudPCRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPCRequestBuilderGetQueryParameters get cloudPC from roleManagement
type CloudPCRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CloudPCRequestBuilderPatchOptions options for Patch
type CloudPCRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewCloudPCRequestBuilderInternal instantiates a new CloudPCRequestBuilder and sets the default values.
func NewCloudPCRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCRequestBuilder) {
    m := &CloudPCRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/cloudPC{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCloudPCRequestBuilder instantiates a new CloudPCRequestBuilder and sets the default values.
func NewCloudPCRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cloudPC for roleManagement
func (m *CloudPCRequestBuilder) CreateDeleteRequestInformation(options *CloudPCRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get cloudPC from roleManagement
func (m *CloudPCRequestBuilder) CreateGetRequestInformation(options *CloudPCRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property cloudPC in roleManagement
func (m *CloudPCRequestBuilder) CreatePatchRequestInformation(options *CloudPCRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property cloudPC for roleManagement
func (m *CloudPCRequestBuilder) Delete(options *CloudPCRequestBuilderDeleteOptions)(error) {
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
// Get get cloudPC from roleManagement
func (m *CloudPCRequestBuilder) Get(options *CloudPCRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRbacApplicationMultiple() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple), nil
}
// Patch update the navigation property cloudPC in roleManagement
func (m *CloudPCRequestBuilder) Patch(options *CloudPCRequestBuilderPatchOptions)(error) {
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
func (m *CloudPCRequestBuilder) ResourceNamespaces()(*id0c942d6747644ac6d906984e3b3d0046ec64b19ed0e042c591f0eda0c0f8957.ResourceNamespacesRequestBuilder) {
    return id0c942d6747644ac6d906984e3b3d0046ec64b19ed0e042c591f0eda0c0f8957.NewResourceNamespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceNamespacesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.cloudPC.resourceNamespaces.item collection
func (m *CloudPCRequestBuilder) ResourceNamespacesById(id string)(*if2fe6d33056a86a8c02bd09d388e1ad97378adc4d03f5d4afe21fd7135b39314.UnifiedRbacResourceNamespaceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceNamespace_id"] = id
    }
    return if2fe6d33056a86a8c02bd09d388e1ad97378adc4d03f5d4afe21fd7135b39314.NewUnifiedRbacResourceNamespaceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) RoleAssignments()(*i21e54f9ee2d47268e20afd63152400c9b4c1f5fccc344e57fac33bd5bafc27e9.RoleAssignmentsRequestBuilder) {
    return i21e54f9ee2d47268e20afd63152400c9b4c1f5fccc344e57fac33bd5bafc27e9.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.cloudPC.roleAssignments.item collection
func (m *CloudPCRequestBuilder) RoleAssignmentsById(id string)(*i5122da689a05212ccc762ebf1a452cf0621ee3e3ee2ed9c79b323f7d87ab82ba.UnifiedRoleAssignmentMultipleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentMultiple_id"] = id
    }
    return i5122da689a05212ccc762ebf1a452cf0621ee3e3ee2ed9c79b323f7d87ab82ba.NewUnifiedRoleAssignmentMultipleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) RoleDefinitions()(*i31977c399a73e8dceaa0f68ce2147b8b69cdb6d448081c8df74e04127a6f5b7e.RoleDefinitionsRequestBuilder) {
    return i31977c399a73e8dceaa0f68ce2147b8b69cdb6d448081c8df74e04127a6f5b7e.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.cloudPC.roleDefinitions.item collection
func (m *CloudPCRequestBuilder) RoleDefinitionsById(id string)(*i6596aa8a9eb296ce0e4c8af117114f5b675a99379d789079c98d15d2b4b41cb0.UnifiedRoleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition_id"] = id
    }
    return i6596aa8a9eb296ce0e4c8af117114f5b675a99379d789079c98d15d2b4b41cb0.NewUnifiedRoleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
