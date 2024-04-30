package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder provides operations to call the getByUserIdAndRole method.
type VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetQueryParameters get a list of virtualEventWebinar objects where the specified user is either the organizer or a coorganizer.
type VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetQueryParameters struct {
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
// VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetQueryParameters
}
// NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal instantiates a new VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, role *string, userId *string)(*VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    m := &VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/getByUserIdAndRole(userId='{userId}',role='{role}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if role != nil {
        m.BaseRequestBuilder.PathParameters["role"] = *role
    }
    if userId != nil {
        m.BaseRequestBuilder.PathParameters["userId"] = *userId
    }
    return m
}
// NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder instantiates a new VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get get a list of virtualEventWebinar objects where the specified user is either the organizer or a coorganizer.
// Deprecated: This method is obsolete. Use GetAsGetByUserIdAndRoleWithUserIdWithRoleGetResponse instead.
// returns a VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventwebinar-getbyuseridandrole?view=graph-rest-beta
func (m *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration)(VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleResponseable), nil
}
// GetAsGetByUserIdAndRoleWithUserIdWithRoleGetResponse get a list of virtualEventWebinar objects where the specified user is either the organizer or a coorganizer.
// returns a VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventwebinar-getbyuseridandrole?view=graph-rest-beta
func (m *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) GetAsGetByUserIdAndRoleWithUserIdWithRoleGetResponse(ctx context.Context, requestConfiguration *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration)(VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleGetResponseable), nil
}
// ToGetRequestInformation get a list of virtualEventWebinar objects where the specified user is either the organizer or a coorganizer.
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder when successful
func (m *VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    return NewVirtualEventsWebinarsGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
