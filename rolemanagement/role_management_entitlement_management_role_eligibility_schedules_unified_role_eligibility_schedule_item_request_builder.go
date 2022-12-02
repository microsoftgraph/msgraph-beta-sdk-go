package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
type RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetQueryParameters get roleEligibilitySchedules from roleManagement
type RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetQueryParameters
}
// RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal instantiates a new UnifiedRoleEligibilityScheduleItemRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    m := &RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleEligibilitySchedules/{unifiedRoleEligibilitySchedule%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder instantiates a new UnifiedRoleEligibilityScheduleItemRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleEligibilitySchedules for roleManagement
func (m *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get roleEligibilitySchedules from roleManagement
func (m *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleEligibilitySchedules in roleManagement
func (m *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable, requestConfiguration *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleEligibilitySchedules for roleManagement
func (m *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get roleEligibilitySchedules from roleManagement
func (m *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable), nil
}
// Patch update the navigation property roleEligibilitySchedules in roleManagement
func (m *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable, requestConfiguration *RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable), nil
}
