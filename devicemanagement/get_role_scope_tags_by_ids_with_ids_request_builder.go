package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetRoleScopeTagsByIdsWithIdsRequestBuilder provides operations to call the getRoleScopeTagsByIds method.
type GetRoleScopeTagsByIdsWithIdsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetRoleScopeTagsByIdsWithIdsRequestBuilderGetQueryParameters invoke function getRoleScopeTagsByIds
type GetRoleScopeTagsByIdsWithIdsRequestBuilderGetQueryParameters struct {
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
// GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetQueryParameters
}
// NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal instantiates a new GetRoleScopeTagsByIdsWithIdsRequestBuilder and sets the default values.
func NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, ids *string)(*GetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    m := &GetRoleScopeTagsByIdsWithIdsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/getRoleScopeTagsByIds(ids={ids}){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if ids != nil {
        m.BaseRequestBuilder.PathParameters["ids"] = *ids
    }
    return m
}
// NewGetRoleScopeTagsByIdsWithIdsRequestBuilder instantiates a new GetRoleScopeTagsByIdsWithIdsRequestBuilder and sets the default values.
func NewGetRoleScopeTagsByIdsWithIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getRoleScopeTagsByIds
// Deprecated: This method is obsolete. Use GetAsGetRoleScopeTagsByIdsWithIdsGetResponse instead.
// returns a GetRoleScopeTagsByIdsWithIdsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) Get(ctx context.Context, requestConfiguration *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration)(GetRoleScopeTagsByIdsWithIdsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetRoleScopeTagsByIdsWithIdsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetRoleScopeTagsByIdsWithIdsResponseable), nil
}
// GetAsGetRoleScopeTagsByIdsWithIdsGetResponse invoke function getRoleScopeTagsByIds
// returns a GetRoleScopeTagsByIdsWithIdsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) GetAsGetRoleScopeTagsByIdsWithIdsGetResponse(ctx context.Context, requestConfiguration *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration)(GetRoleScopeTagsByIdsWithIdsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetRoleScopeTagsByIdsWithIdsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetRoleScopeTagsByIdsWithIdsGetResponseable), nil
}
// ToGetRequestInformation invoke function getRoleScopeTagsByIds
// returns a *RequestInformation when successful
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetRoleScopeTagsByIdsWithIdsRequestBuilder when successful
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) WithUrl(rawUrl string)(*GetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    return NewGetRoleScopeTagsByIdsWithIdsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
