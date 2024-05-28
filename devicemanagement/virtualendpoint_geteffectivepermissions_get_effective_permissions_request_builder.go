package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder provides operations to call the getEffectivePermissions method.
type VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderGetQueryParameters get the effective permissions of the currently authenticated user, helping UX hide or disable content that the current user doesn't have access to.
type VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderGetQueryParameters struct {
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
// VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderGetQueryParameters
}
// NewVirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderInternal instantiates a new VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder and sets the default values.
func NewVirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder) {
    m := &VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/getEffectivePermissions(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder instantiates a new VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder and sets the default values.
func NewVirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the effective permissions of the currently authenticated user, helping UX hide or disable content that the current user doesn't have access to.
// Deprecated: This method is obsolete. Use GetAsGetEffectivePermissionsGetResponse instead.
// returns a VirtualendpointGeteffectivepermissionsGetEffectivePermissionsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-geteffectivepermissions?view=graph-rest-beta
func (m *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderGetRequestConfiguration)(VirtualendpointGeteffectivepermissionsGetEffectivePermissionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointGeteffectivepermissionsGetEffectivePermissionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointGeteffectivepermissionsGetEffectivePermissionsResponseable), nil
}
// GetAsGetEffectivePermissionsGetResponse get the effective permissions of the currently authenticated user, helping UX hide or disable content that the current user doesn't have access to.
// returns a VirtualendpointGeteffectivepermissionsGetEffectivePermissionsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-geteffectivepermissions?view=graph-rest-beta
func (m *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder) GetAsGetEffectivePermissionsGetResponse(ctx context.Context, requestConfiguration *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderGetRequestConfiguration)(VirtualendpointGeteffectivepermissionsGetEffectivePermissionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointGeteffectivepermissionsGetEffectivePermissionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointGeteffectivepermissionsGetEffectivePermissionsGetResponseable), nil
}
// ToGetRequestInformation get the effective permissions of the currently authenticated user, helping UX hide or disable content that the current user doesn't have access to.
// returns a *RequestInformation when successful
func (m *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder when successful
func (m *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder) {
    return NewVirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
