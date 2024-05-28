package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder provides operations to call the retrieveScopedPermissions method.
type VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderGetQueryParameters get the permissions and corresponding scope IDs for which the authenticated user has access.
type VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderGetQueryParameters struct {
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
// VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderGetQueryParameters
}
// NewVirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderInternal instantiates a new VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder and sets the default values.
func NewVirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder) {
    m := &VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/retrieveScopedPermissions(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder instantiates a new VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder and sets the default values.
func NewVirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the permissions and corresponding scope IDs for which the authenticated user has access.
// Deprecated: This method is obsolete. Use GetAsRetrieveScopedPermissionsGetResponse instead.
// returns a VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-retrievescopedpermissions?view=graph-rest-beta
func (m *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderGetRequestConfiguration)(VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsResponseable), nil
}
// GetAsRetrieveScopedPermissionsGetResponse get the permissions and corresponding scope IDs for which the authenticated user has access.
// returns a VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-retrievescopedpermissions?view=graph-rest-beta
func (m *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder) GetAsRetrieveScopedPermissionsGetResponse(ctx context.Context, requestConfiguration *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderGetRequestConfiguration)(VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsGetResponseable), nil
}
// ToGetRequestInformation get the permissions and corresponding scope IDs for which the authenticated user has access.
// returns a *RequestInformation when successful
func (m *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder when successful
func (m *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder) {
    return NewVirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
