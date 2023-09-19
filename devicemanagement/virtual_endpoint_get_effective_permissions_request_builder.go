package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointGetEffectivePermissionsRequestBuilder provides operations to call the getEffectivePermissions method.
type VirtualEndpointGetEffectivePermissionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointGetEffectivePermissionsRequestBuilderGetQueryParameters invoke function getEffectivePermissions
type VirtualEndpointGetEffectivePermissionsRequestBuilderGetQueryParameters struct {
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
// VirtualEndpointGetEffectivePermissionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointGetEffectivePermissionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointGetEffectivePermissionsRequestBuilderGetQueryParameters
}
// NewVirtualEndpointGetEffectivePermissionsRequestBuilderInternal instantiates a new GetEffectivePermissionsRequestBuilder and sets the default values.
func NewVirtualEndpointGetEffectivePermissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointGetEffectivePermissionsRequestBuilder) {
    m := &VirtualEndpointGetEffectivePermissionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/getEffectivePermissions(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewVirtualEndpointGetEffectivePermissionsRequestBuilder instantiates a new GetEffectivePermissionsRequestBuilder and sets the default values.
func NewVirtualEndpointGetEffectivePermissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointGetEffectivePermissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointGetEffectivePermissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getEffectivePermissions
func (m *VirtualEndpointGetEffectivePermissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointGetEffectivePermissionsRequestBuilderGetRequestConfiguration)(VirtualEndpointGetEffectivePermissionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointGetEffectivePermissionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointGetEffectivePermissionsResponseable), nil
}
// ToGetRequestInformation invoke function getEffectivePermissions
func (m *VirtualEndpointGetEffectivePermissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointGetEffectivePermissionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointGetEffectivePermissionsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointGetEffectivePermissionsRequestBuilder) {
    return NewVirtualEndpointGetEffectivePermissionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
