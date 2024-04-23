package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder provides operations to call the filterByCurrentUser method.
type EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetQueryParameters get a list of the unifiedRoleEligibilitySchedule objects and their properties associated with a particular principal object.
type EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetQueryParameters struct {
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
// EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetQueryParameters
}
// NewEnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderInternal instantiates a new EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, on *string)(*EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder) {
    m := &EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/roleEligibilitySchedules/filterByCurrentUser(on='{on}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if on != nil {
        m.BaseRequestBuilder.PathParameters["on"] = *on
    }
    return m
}
// NewEnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder instantiates a new EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get a list of the unifiedRoleEligibilitySchedule objects and their properties associated with a particular principal object.
// Deprecated: This method is obsolete. Use GetAsFilterByCurrentUserWithOnGetResponse instead.
// returns a EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleeligibilityschedule-filterbycurrentuser?view=graph-rest-beta
func (m *EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnResponseable), nil
}
// GetAsFilterByCurrentUserWithOnGetResponse get a list of the unifiedRoleEligibilitySchedule objects and their properties associated with a particular principal object.
// returns a EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleeligibilityschedule-filterbycurrentuser?view=graph-rest-beta
func (m *EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder) GetAsFilterByCurrentUserWithOnGetResponse(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnGetResponseable), nil
}
// ToGetRequestInformation get a list of the unifiedRoleEligibilitySchedule objects and their properties associated with a particular principal object.
// returns a *RequestInformation when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder) WithUrl(rawUrl string)(*EnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder) {
    return NewEnterpriseAppsItemRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
