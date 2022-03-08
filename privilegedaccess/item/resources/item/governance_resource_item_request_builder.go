package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2e9006b5b06b12fded8deaf8708ce23ea064941fe87e92263a0b0bcd91c37adc "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item/rolesettings"
    i3605d902f92a6bea8a258127b9515b3d3c546eb9d658395d666d58b4d0c0f90e "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item/roledefinitions"
    ib0aaebf13ea4ba70136f557b80a4927156762d3a0220f350b4c9b1c42b2f0081 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item/parent"
    ib542076fdb1e7df91be849b5efcd3faf53db7fd9695526930016466566214990 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item/roleassignments"
    icd973db06be9431b2115931b7c6e087161e5e617bafd8a80b393a773795c6fa4 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item/roleassignmentrequests"
    i5906c03a7eaf4ae7a45bdb589b8d2bcaff6942303436db9394659bfc3f0f2e2c "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item/rolesettings/item"
    i673d165f29308b156cf4f29128869e70ff180d4207ec74e37aa01db9557b80b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item/roledefinitions/item"
    i6a3c41cedb50c4351ee3bc74ef3af69a59935cb40e55ef89b828b0345e418a11 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item/roleassignmentrequests/item"
    id6d5ee9260c4fa93ba637e63f200cf83913f46edc1e0a5aef9e499f9e0df0962 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item/roleassignments/item"
)

// GovernanceResourceItemRequestBuilder provides operations to manage the resources property of the microsoft.graph.privilegedAccess entity.
type GovernanceResourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GovernanceResourceItemRequestBuilderDeleteOptions options for Delete
type GovernanceResourceItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceResourceItemRequestBuilderGetOptions options for Get
type GovernanceResourceItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GovernanceResourceItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceResourceItemRequestBuilderGetQueryParameters a collection of resources for the provider.
type GovernanceResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GovernanceResourceItemRequestBuilderPatchOptions options for Patch
type GovernanceResourceItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceResourceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGovernanceResourceItemRequestBuilderInternal instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewGovernanceResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceResourceItemRequestBuilder) {
    m := &GovernanceResourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess_id}/resources/{governanceResource_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceResourceItemRequestBuilder instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewGovernanceResourceItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property resources for privilegedAccess
func (m *GovernanceResourceItemRequestBuilder) CreateDeleteRequestInformation(options *GovernanceResourceItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of resources for the provider.
func (m *GovernanceResourceItemRequestBuilder) CreateGetRequestInformation(options *GovernanceResourceItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property resources in privilegedAccess
func (m *GovernanceResourceItemRequestBuilder) CreatePatchRequestInformation(options *GovernanceResourceItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property resources for privilegedAccess
func (m *GovernanceResourceItemRequestBuilder) Delete(options *GovernanceResourceItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of resources for the provider.
func (m *GovernanceResourceItemRequestBuilder) Get(options *GovernanceResourceItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceResourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateGovernanceResourceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceResourceable), nil
}
func (m *GovernanceResourceItemRequestBuilder) Parent()(*ib0aaebf13ea4ba70136f557b80a4927156762d3a0220f350b4c9b1c42b2f0081.ParentRequestBuilder) {
    return ib0aaebf13ea4ba70136f557b80a4927156762d3a0220f350b4c9b1c42b2f0081.NewParentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property resources in privilegedAccess
func (m *GovernanceResourceItemRequestBuilder) Patch(options *GovernanceResourceItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *GovernanceResourceItemRequestBuilder) RoleAssignmentRequests()(*icd973db06be9431b2115931b7c6e087161e5e617bafd8a80b393a773795c6fa4.RoleAssignmentRequestsRequestBuilder) {
    return icd973db06be9431b2115931b7c6e087161e5e617bafd8a80b393a773795c6fa4.NewRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item.resources.item.roleAssignmentRequests.item collection
func (m *GovernanceResourceItemRequestBuilder) RoleAssignmentRequestsById(id string)(*i6a3c41cedb50c4351ee3bc74ef3af69a59935cb40e55ef89b828b0345e418a11.GovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest_id"] = id
    }
    return i6a3c41cedb50c4351ee3bc74ef3af69a59935cb40e55ef89b828b0345e418a11.NewGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GovernanceResourceItemRequestBuilder) RoleAssignments()(*ib542076fdb1e7df91be849b5efcd3faf53db7fd9695526930016466566214990.RoleAssignmentsRequestBuilder) {
    return ib542076fdb1e7df91be849b5efcd3faf53db7fd9695526930016466566214990.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item.resources.item.roleAssignments.item collection
func (m *GovernanceResourceItemRequestBuilder) RoleAssignmentsById(id string)(*id6d5ee9260c4fa93ba637e63f200cf83913f46edc1e0a5aef9e499f9e0df0962.GovernanceRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment_id"] = id
    }
    return id6d5ee9260c4fa93ba637e63f200cf83913f46edc1e0a5aef9e499f9e0df0962.NewGovernanceRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GovernanceResourceItemRequestBuilder) RoleDefinitions()(*i3605d902f92a6bea8a258127b9515b3d3c546eb9d658395d666d58b4d0c0f90e.RoleDefinitionsRequestBuilder) {
    return i3605d902f92a6bea8a258127b9515b3d3c546eb9d658395d666d58b4d0c0f90e.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item.resources.item.roleDefinitions.item collection
func (m *GovernanceResourceItemRequestBuilder) RoleDefinitionsById(id string)(*i673d165f29308b156cf4f29128869e70ff180d4207ec74e37aa01db9557b80b0.GovernanceRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition_id"] = id
    }
    return i673d165f29308b156cf4f29128869e70ff180d4207ec74e37aa01db9557b80b0.NewGovernanceRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GovernanceResourceItemRequestBuilder) RoleSettings()(*i2e9006b5b06b12fded8deaf8708ce23ea064941fe87e92263a0b0bcd91c37adc.RoleSettingsRequestBuilder) {
    return i2e9006b5b06b12fded8deaf8708ce23ea064941fe87e92263a0b0bcd91c37adc.NewRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item.resources.item.roleSettings.item collection
func (m *GovernanceResourceItemRequestBuilder) RoleSettingsById(id string)(*i5906c03a7eaf4ae7a45bdb589b8d2bcaff6942303436db9394659bfc3f0f2e2c.GovernanceRoleSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleSetting_id"] = id
    }
    return i5906c03a7eaf4ae7a45bdb589b8d2bcaff6942303436db9394659bfc3f0f2e2c.NewGovernanceRoleSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
