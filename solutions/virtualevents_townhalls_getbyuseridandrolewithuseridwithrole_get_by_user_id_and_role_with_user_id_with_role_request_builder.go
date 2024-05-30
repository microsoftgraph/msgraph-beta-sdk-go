package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder provides operations to call the getByUserIdAndRole method.
type VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetQueryParameters invoke function getByUserIdAndRole
type VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetQueryParameters struct {
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
// VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetQueryParameters
}
// NewVirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal instantiates a new VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder and sets the default values.
func NewVirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, role *string, userId *string)(*VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    m := &VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/townhalls/getByUserIdAndRole(userId='{userId}',role='{role}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if role != nil {
        m.BaseRequestBuilder.PathParameters["role"] = *role
    }
    if userId != nil {
        m.BaseRequestBuilder.PathParameters["userId"] = *userId
    }
    return m
}
// NewVirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder instantiates a new VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder and sets the default values.
func NewVirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function getByUserIdAndRole
// Deprecated: This method is obsolete. Use GetAsGetByUserIdAndRoleWithUserIdWithRoleGetResponse instead.
// returns a VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration)(VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleResponseable), nil
}
// GetAsGetByUserIdAndRoleWithUserIdWithRoleGetResponse invoke function getByUserIdAndRole
// returns a VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) GetAsGetByUserIdAndRoleWithUserIdWithRoleGetResponse(ctx context.Context, requestConfiguration *VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration)(VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleGetResponseable), nil
}
// ToGetRequestInformation invoke function getByUserIdAndRole
// returns a *RequestInformation when successful
func (m *VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder when successful
func (m *VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    return NewVirtualeventsTownhallsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
