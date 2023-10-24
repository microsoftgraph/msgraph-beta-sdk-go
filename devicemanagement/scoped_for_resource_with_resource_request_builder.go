package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ScopedForResourceWithResourceRequestBuilder provides operations to call the scopedForResource method.
type ScopedForResourceWithResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ScopedForResourceWithResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScopedForResourceWithResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewScopedForResourceWithResourceRequestBuilderInternal instantiates a new ScopedForResourceWithResourceRequestBuilder and sets the default values.
func NewScopedForResourceWithResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, resource *string)(*ScopedForResourceWithResourceRequestBuilder) {
    m := &ScopedForResourceWithResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/scopedForResource(resource='{resource}')", pathParameters),
    }
    if resource != nil {
        m.BaseRequestBuilder.PathParameters["resource"] = *resource
    }
    return m
}
// NewScopedForResourceWithResourceRequestBuilder instantiates a new ScopedForResourceWithResourceRequestBuilder and sets the default values.
func NewScopedForResourceWithResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScopedForResourceWithResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScopedForResourceWithResourceRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function scopedForResource
// Deprecated: This method is obsolete. Use GetAsScopedForResourceWithResourceGetResponse instead.
func (m *ScopedForResourceWithResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *ScopedForResourceWithResourceRequestBuilderGetRequestConfiguration)(ScopedForResourceWithResourceResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateScopedForResourceWithResourceResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ScopedForResourceWithResourceResponseable), nil
}
// GetAsScopedForResourceWithResourceGetResponse invoke function scopedForResource
func (m *ScopedForResourceWithResourceRequestBuilder) GetAsScopedForResourceWithResourceGetResponse(ctx context.Context, requestConfiguration *ScopedForResourceWithResourceRequestBuilderGetRequestConfiguration)(ScopedForResourceWithResourceGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateScopedForResourceWithResourceGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ScopedForResourceWithResourceGetResponseable), nil
}
// ToGetRequestInformation invoke function scopedForResource
func (m *ScopedForResourceWithResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ScopedForResourceWithResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ScopedForResourceWithResourceRequestBuilder) WithUrl(rawUrl string)(*ScopedForResourceWithResourceRequestBuilder) {
    return NewScopedForResourceWithResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
