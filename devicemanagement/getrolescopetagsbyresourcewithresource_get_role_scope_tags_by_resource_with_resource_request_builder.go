package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder provides operations to call the getRoleScopeTagsByResource method.
type GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderGetQueryParameters invoke function getRoleScopeTagsByResource
type GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderGetQueryParameters struct {
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
// GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderGetQueryParameters
}
// NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal instantiates a new GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder and sets the default values.
func NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, resource *string)(*GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    m := &GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/getRoleScopeTagsByResource(resource='{resource}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if resource != nil {
        m.BaseRequestBuilder.PathParameters["resource"] = *resource
    }
    return m
}
// NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder instantiates a new GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder and sets the default values.
func NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getRoleScopeTagsByResource
// Deprecated: This method is obsolete. Use GetAsGetRoleScopeTagsByResourceWithResourceGetResponse instead.
// returns a GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderGetRequestConfiguration)(GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponseable), nil
}
// GetAsGetRoleScopeTagsByResourceWithResourceGetResponse invoke function getRoleScopeTagsByResource
// returns a GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder) GetAsGetRoleScopeTagsByResourceWithResourceGetResponse(ctx context.Context, requestConfiguration *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderGetRequestConfiguration)(GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponseable), nil
}
// ToGetRequestInformation invoke function getRoleScopeTagsByResource
// returns a *RequestInformation when successful
func (m *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder when successful
func (m *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder) WithUrl(rawUrl string)(*GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    return NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
