package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder provides operations to call the scopedForResource method.
type ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilderInternal instantiates a new ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder and sets the default values.
func NewScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, resource *string)(*ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder) {
    m := &ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/scopedForResource(resource='{resource}')", pathParameters),
    }
    if resource != nil {
        m.BaseRequestBuilder.PathParameters["resource"] = *resource
    }
    return m
}
// NewScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder instantiates a new ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder and sets the default values.
func NewScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function scopedForResource
// Deprecated: This method is obsolete. Use GetAsScopedForResourceWithResourceGetResponse instead.
// returns a ScopedforresourcewithresourceScopedForResourceWithResourceResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilderGetRequestConfiguration)(ScopedforresourcewithresourceScopedForResourceWithResourceResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateScopedforresourcewithresourceScopedForResourceWithResourceResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ScopedforresourcewithresourceScopedForResourceWithResourceResponseable), nil
}
// GetAsScopedForResourceWithResourceGetResponse invoke function scopedForResource
// returns a ScopedforresourcewithresourceScopedForResourceWithResourceGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder) GetAsScopedForResourceWithResourceGetResponse(ctx context.Context, requestConfiguration *ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilderGetRequestConfiguration)(ScopedforresourcewithresourceScopedForResourceWithResourceGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateScopedforresourcewithresourceScopedForResourceWithResourceGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ScopedforresourcewithresourceScopedForResourceWithResourceGetResponseable), nil
}
// ToGetRequestInformation invoke function scopedForResource
// returns a *RequestInformation when successful
func (m *ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder when successful
func (m *ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder) WithUrl(rawUrl string)(*ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder) {
    return NewScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
