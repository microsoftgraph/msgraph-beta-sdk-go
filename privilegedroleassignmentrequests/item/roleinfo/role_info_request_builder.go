package roleinfo

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1721b0eaada2dd5efbd4a95951797885d4f2d3cde7447b79ef6372d21be0a6bf "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests/item/roleinfo/summary"
    i3fe6caa29a2409f68e98f1a04bf1f86d4284ac6cc7508637c3461dae6ebb41cc "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests/item/roleinfo/settings"
    i9a7b7ed52442792951b363d33df86a53ea6b293e88308fbeb15e56b47a4a39d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests/item/roleinfo/assignments"
    icf5a99f187b3caa046cd63542f946400cb9a0c7d4870bbafefdae9cab6db72b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests/item/roleinfo/selfdeactivate"
    ie89f568cd53dfcd03f3a71171ccb9a7afeecf2eb06b5e7b0dbecbbaf1fd2bf74 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests/item/roleinfo/selfactivate"
    i5793653b0cfb9d5b0eedc68e4dc192fa8343ad2b1ebc4a9031ffee56bd6ac846 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests/item/roleinfo/assignments/item"
)

// RoleInfoRequestBuilder provides operations to manage the roleInfo property of the microsoft.graph.privilegedRoleAssignmentRequest entity.
type RoleInfoRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleInfoRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleInfoRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleInfoRequestBuilderGetQueryParameters the roleInfo object of the role assignment request.
type RoleInfoRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleInfoRequestBuilderGetQueryParameters
}
// RoleInfoRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleInfoRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments the assignments property
func (m *RoleInfoRequestBuilder) Assignments()(*i9a7b7ed52442792951b363d33df86a53ea6b293e88308fbeb15e56b47a4a39d7.AssignmentsRequestBuilder) {
    return i9a7b7ed52442792951b363d33df86a53ea6b293e88308fbeb15e56b47a4a39d7.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedRoleAssignmentRequests.item.roleInfo.assignments.item collection
func (m *RoleInfoRequestBuilder) AssignmentsById(id string)(*i5793653b0cfb9d5b0eedc68e4dc192fa8343ad2b1ebc4a9031ffee56bd6ac846.PrivilegedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedRoleAssignment%2Did"] = id
    }
    return i5793653b0cfb9d5b0eedc68e4dc192fa8343ad2b1ebc4a9031ffee56bd6ac846.NewPrivilegedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRoleInfoRequestBuilderInternal instantiates a new RoleInfoRequestBuilder and sets the default values.
func NewRoleInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleInfoRequestBuilder) {
    m := &RoleInfoRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoleAssignmentRequests/{privilegedRoleAssignmentRequest%2Did}/roleInfo{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleInfoRequestBuilder instantiates a new RoleInfoRequestBuilder and sets the default values.
func NewRoleInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleInfo for privilegedRoleAssignmentRequests
func (m *RoleInfoRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property roleInfo for privilegedRoleAssignmentRequests
func (m *RoleInfoRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *RoleInfoRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the roleInfo object of the role assignment request.
func (m *RoleInfoRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the roleInfo object of the role assignment request.
func (m *RoleInfoRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RoleInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleInfo in privilegedRoleAssignmentRequests
func (m *RoleInfoRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property roleInfo in privilegedRoleAssignmentRequests
func (m *RoleInfoRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, requestConfiguration *RoleInfoRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property roleInfo for privilegedRoleAssignmentRequests
func (m *RoleInfoRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleInfoRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get the roleInfo object of the role assignment request.
func (m *RoleInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleInfoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable), nil
}
// Patch update the navigation property roleInfo in privilegedRoleAssignmentRequests
func (m *RoleInfoRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, requestConfiguration *RoleInfoRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// SelfActivate the selfActivate property
func (m *RoleInfoRequestBuilder) SelfActivate()(*ie89f568cd53dfcd03f3a71171ccb9a7afeecf2eb06b5e7b0dbecbbaf1fd2bf74.SelfActivateRequestBuilder) {
    return ie89f568cd53dfcd03f3a71171ccb9a7afeecf2eb06b5e7b0dbecbbaf1fd2bf74.NewSelfActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SelfDeactivate the selfDeactivate property
func (m *RoleInfoRequestBuilder) SelfDeactivate()(*icf5a99f187b3caa046cd63542f946400cb9a0c7d4870bbafefdae9cab6db72b6.SelfDeactivateRequestBuilder) {
    return icf5a99f187b3caa046cd63542f946400cb9a0c7d4870bbafefdae9cab6db72b6.NewSelfDeactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *RoleInfoRequestBuilder) Settings()(*i3fe6caa29a2409f68e98f1a04bf1f86d4284ac6cc7508637c3461dae6ebb41cc.SettingsRequestBuilder) {
    return i3fe6caa29a2409f68e98f1a04bf1f86d4284ac6cc7508637c3461dae6ebb41cc.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Summary the summary property
func (m *RoleInfoRequestBuilder) Summary()(*i1721b0eaada2dd5efbd4a95951797885d4f2d3cde7447b79ef6372d21be0a6bf.SummaryRequestBuilder) {
    return i1721b0eaada2dd5efbd4a95951797885d4f2d3cde7447b79ef6372d21be0a6bf.NewSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
