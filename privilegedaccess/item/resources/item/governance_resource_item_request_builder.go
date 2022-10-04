package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GovernanceResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GovernanceResourceItemRequestBuilderGetQueryParameters a collection of resources for the provider.
type GovernanceResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GovernanceResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GovernanceResourceItemRequestBuilderGetQueryParameters
}
// GovernanceResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGovernanceResourceItemRequestBuilderInternal instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewGovernanceResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceResourceItemRequestBuilder) {
    m := &GovernanceResourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceResourceItemRequestBuilder instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewGovernanceResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property resources for privilegedAccess
func (m *GovernanceResourceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a collection of resources for the provider.
func (m *GovernanceResourceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GovernanceResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property resources in privilegedAccess
func (m *GovernanceResourceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *GovernanceResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property resources for privilegedAccess
func (m *GovernanceResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of resources for the provider.
func (m *GovernanceResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GovernanceResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable), nil
}
// Parent the parent property
func (m *GovernanceResourceItemRequestBuilder) Parent()(*ib0aaebf13ea4ba70136f557b80a4927156762d3a0220f350b4c9b1c42b2f0081.ParentRequestBuilder) {
    return ib0aaebf13ea4ba70136f557b80a4927156762d3a0220f350b4c9b1c42b2f0081.NewParentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property resources in privilegedAccess
func (m *GovernanceResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *GovernanceResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable), nil
}
// RoleAssignmentRequests the roleAssignmentRequests property
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
        urlTplParams["governanceRoleAssignmentRequest%2Did"] = id
    }
    return i6a3c41cedb50c4351ee3bc74ef3af69a59935cb40e55ef89b828b0345e418a11.NewGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments the roleAssignments property
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
        urlTplParams["governanceRoleAssignment%2Did"] = id
    }
    return id6d5ee9260c4fa93ba637e63f200cf83913f46edc1e0a5aef9e499f9e0df0962.NewGovernanceRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions the roleDefinitions property
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
        urlTplParams["governanceRoleDefinition%2Did"] = id
    }
    return i673d165f29308b156cf4f29128869e70ff180d4207ec74e37aa01db9557b80b0.NewGovernanceRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleSettings the roleSettings property
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
        urlTplParams["governanceRoleSetting%2Did"] = id
    }
    return i5906c03a7eaf4ae7a45bdb589b8d2bcaff6942303436db9394659bfc3f0f2e2c.NewGovernanceRoleSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
