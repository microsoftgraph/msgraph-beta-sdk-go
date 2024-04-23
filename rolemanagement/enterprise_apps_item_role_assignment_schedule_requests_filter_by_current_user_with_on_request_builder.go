package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder provides operations to call the filterByCurrentUser method.
type EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetQueryParameters in PIM, retrieve the requests for active role assignments for a particular principal. The principal can be the creator or approver of the unifiedRoleAssignmentScheduleRequest object, or they can be the target of the assignment.
type EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetQueryParameters struct {
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
// EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetQueryParameters
}
// NewEnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderInternal instantiates a new EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, on *string)(*EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder) {
    m := &EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/roleAssignmentScheduleRequests/filterByCurrentUser(on='{on}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if on != nil {
        m.BaseRequestBuilder.PathParameters["on"] = *on
    }
    return m
}
// NewEnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder instantiates a new EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get in PIM, retrieve the requests for active role assignments for a particular principal. The principal can be the creator or approver of the unifiedRoleAssignmentScheduleRequest object, or they can be the target of the assignment.
// Deprecated: This method is obsolete. Use GetAsFilterByCurrentUserWithOnGetResponse instead.
// returns a EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentschedulerequest-filterbycurrentuser?view=graph-rest-beta
func (m *EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponseable), nil
}
// GetAsFilterByCurrentUserWithOnGetResponse in PIM, retrieve the requests for active role assignments for a particular principal. The principal can be the creator or approver of the unifiedRoleAssignmentScheduleRequest object, or they can be the target of the assignment.
// returns a EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentschedulerequest-filterbycurrentuser?view=graph-rest-beta
func (m *EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder) GetAsFilterByCurrentUserWithOnGetResponse(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseable), nil
}
// ToGetRequestInformation in PIM, retrieve the requests for active role assignments for a particular principal. The principal can be the creator or approver of the unifiedRoleAssignmentScheduleRequest object, or they can be the target of the assignment.
// returns a *RequestInformation when successful
func (m *EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder) WithUrl(rawUrl string)(*EnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder) {
    return NewEnterpriseAppsItemRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
