package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder provides operations to manage the targetSchedule property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
type EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderGetQueryParameters the schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand.
type EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderGetQueryParameters
}
// NewEnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderInternal instantiates a new TargetScheduleRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) {
    m := &EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest%2Did}/targetSchedule{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder instantiates a new TargetScheduleRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand.
func (m *EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable), nil
}
// ToGetRequestInformation the schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand.
func (m *EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) WithUrl(rawUrl string)(*EnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) {
    return NewEnterpriseAppsItemRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
