package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
type EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderGetQueryParameters get roleAssignmentSchedules from roleManagement
type EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderGetQueryParameters
}
// EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleAssignmentScheduleId provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
// returns a *EnterpriseAppsItemRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder when successful
func (m *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) ByUnifiedRoleAssignmentScheduleId(unifiedRoleAssignmentScheduleId string)(*EnterpriseAppsItemRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleAssignmentScheduleId != "" {
        urlTplParams["unifiedRoleAssignmentSchedule%2Did"] = unifiedRoleAssignmentScheduleId
    }
    return NewEnterpriseAppsItemRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderInternal instantiates a new EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) {
    m := &EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/roleAssignmentSchedules{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder instantiates a new EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EnterpriseAppsItemRoleAssignmentSchedulesCountRequestBuilder when successful
func (m *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) Count()(*EnterpriseAppsItemRoleAssignmentSchedulesCountRequestBuilder) {
    return NewEnterpriseAppsItemRoleAssignmentSchedulesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *EnterpriseAppsItemRoleAssignmentSchedulesFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*EnterpriseAppsItemRoleAssignmentSchedulesFilterByCurrentUserWithOnRequestBuilder) {
    return NewEnterpriseAppsItemRoleAssignmentSchedulesFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get get roleAssignmentSchedules from roleManagement
// returns a UnifiedRoleAssignmentScheduleCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentScheduleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleCollectionResponseable), nil
}
// Post create new navigation property to roleAssignmentSchedules for roleManagement
// returns a UnifiedRoleAssignmentScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable, requestConfiguration *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation get roleAssignmentSchedules from roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to roleAssignmentSchedules for roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable, requestConfiguration *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder when successful
func (m *EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) WithUrl(rawUrl string)(*EnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder) {
    return NewEnterpriseAppsItemRoleAssignmentSchedulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
