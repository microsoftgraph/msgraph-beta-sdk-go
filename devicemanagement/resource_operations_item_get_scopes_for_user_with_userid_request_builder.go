package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder provides operations to call the getScopesForUser method.
type ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ResourceOperationsItemGetScopesForUserWithUseridRequestBuilderGetQueryParameters invoke function getScopesForUser
type ResourceOperationsItemGetScopesForUserWithUseridRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ResourceOperationsItemGetScopesForUserWithUseridRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceOperationsItemGetScopesForUserWithUseridRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ResourceOperationsItemGetScopesForUserWithUseridRequestBuilderGetQueryParameters
}
// NewResourceOperationsItemGetScopesForUserWithUseridRequestBuilderInternal instantiates a new GetScopesForUserWithUseridRequestBuilder and sets the default values.
func NewResourceOperationsItemGetScopesForUserWithUseridRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, userid *string)(*ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder) {
    m := &ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/resourceOperations/{resourceOperation%2Did}/getScopesForUser(userid='{userid}'){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    if userid != nil {
        m.BaseRequestBuilder.PathParameters["userid"] = *userid
    }
    return m
}
// NewResourceOperationsItemGetScopesForUserWithUseridRequestBuilder instantiates a new GetScopesForUserWithUseridRequestBuilder and sets the default values.
func NewResourceOperationsItemGetScopesForUserWithUseridRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceOperationsItemGetScopesForUserWithUseridRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getScopesForUser
// Deprecated: This method is obsolete. Use GetAsGetScopesForUserWithUseridGetResponse instead.
func (m *ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder) Get(ctx context.Context, requestConfiguration *ResourceOperationsItemGetScopesForUserWithUseridRequestBuilderGetRequestConfiguration)(ResourceOperationsItemGetScopesForUserWithUseridResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateResourceOperationsItemGetScopesForUserWithUseridResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ResourceOperationsItemGetScopesForUserWithUseridResponseable), nil
}
// GetAsGetScopesForUserWithUseridGetResponse invoke function getScopesForUser
func (m *ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder) GetAsGetScopesForUserWithUseridGetResponse(ctx context.Context, requestConfiguration *ResourceOperationsItemGetScopesForUserWithUseridRequestBuilderGetRequestConfiguration)(ResourceOperationsItemGetScopesForUserWithUseridGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateResourceOperationsItemGetScopesForUserWithUseridGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ResourceOperationsItemGetScopesForUserWithUseridGetResponseable), nil
}
// ToGetRequestInformation invoke function getScopesForUser
func (m *ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ResourceOperationsItemGetScopesForUserWithUseridRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder) WithUrl(rawUrl string)(*ResourceOperationsItemGetScopesForUserWithUseridRequestBuilder) {
    return NewResourceOperationsItemGetScopesForUserWithUseridRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
