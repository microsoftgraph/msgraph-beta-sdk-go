package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1ad3f3772b94a4fe9eb4a735934236e8263a9a5e8004b1d911e619b2bfbfe3d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/directoryscope"
    i7fbfdd8acfde58648f2e9802cefa0c4426e7a1189e1de1969339baa0c9777c91 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/principal"
    i8625b9843234464c4fef92338104b70660fbcd1c0a46aa2315dcd2d32e2a8289 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/cancel"
    ia2562bf876dabdf8b8f2d1964328f124aa93d0bdee7008197fc224aa17175c37 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/appscope"
    icc542c1f51bf387e280de865ad6853cc67cc8659b96b1ba75c7efcf4ca21486e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/roledefinition"
    ifa3dab5fb247ec11de2a5930325058c34ca3959c5ac61b988f2682b273c02557 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/targetschedule"
)

// UnifiedRoleEligibilityScheduleRequestItemRequestBuilder provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters get roleEligibilityScheduleRequests from roleManagement
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters
}
// UnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope the appScope property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) AppScope()(*ia2562bf876dabdf8b8f2d1964328f124aa93d0bdee7008197fc224aa17175c37.AppScopeRequestBuilder) {
    return ia2562bf876dabdf8b8f2d1964328f124aa93d0bdee7008197fc224aa17175c37.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Cancel()(*i8625b9843234464c4fef92338104b70660fbcd1c0a46aa2315dcd2d32e2a8289.CancelRequestBuilder) {
    return i8625b9843234464c4fef92338104b70660fbcd1c0a46aa2315dcd2d32e2a8289.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    m := &UnifiedRoleEligibilityScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilder instantiates a new UnifiedRoleEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleEligibilityScheduleRequests for roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property roleEligibilityScheduleRequests for roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get roleEligibilityScheduleRequests from roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get roleEligibilityScheduleRequests from roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleEligibilityScheduleRequests in roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property roleEligibilityScheduleRequests in roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleEligibilityScheduleRequests for roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope the directoryScope property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) DirectoryScope()(*i1ad3f3772b94a4fe9eb4a735934236e8263a9a5e8004b1d911e619b2bfbfe3d9.DirectoryScopeRequestBuilder) {
    return i1ad3f3772b94a4fe9eb4a735934236e8263a9a5e8004b1d911e619b2bfbfe3d9.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleEligibilityScheduleRequests from roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable), nil
}
// Patch update the navigation property roleEligibilityScheduleRequests in roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(error) {
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
// Principal the principal property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Principal()(*i7fbfdd8acfde58648f2e9802cefa0c4426e7a1189e1de1969339baa0c9777c91.PrincipalRequestBuilder) {
    return i7fbfdd8acfde58648f2e9802cefa0c4426e7a1189e1de1969339baa0c9777c91.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition the roleDefinition property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) RoleDefinition()(*icc542c1f51bf387e280de865ad6853cc67cc8659b96b1ba75c7efcf4ca21486e.RoleDefinitionRequestBuilder) {
    return icc542c1f51bf387e280de865ad6853cc67cc8659b96b1ba75c7efcf4ca21486e.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetSchedule the targetSchedule property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) TargetSchedule()(*ifa3dab5fb247ec11de2a5930325058c34ca3959c5ac61b988f2682b273c02557.TargetScheduleRequestBuilder) {
    return ifa3dab5fb247ec11de2a5930325058c34ca3959c5ac61b988f2682b273c02557.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
