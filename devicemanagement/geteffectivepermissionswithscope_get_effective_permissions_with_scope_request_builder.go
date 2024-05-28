package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder provides operations to call the getEffectivePermissions method.
type GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderGetQueryParameters invoke function getEffectivePermissions
type GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderGetQueryParameters struct {
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
// GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderGetQueryParameters
}
// NewGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderInternal instantiates a new GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder and sets the default values.
func NewGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, scope *string)(*GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder) {
    m := &GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/getEffectivePermissions(scope='{scope}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if scope != nil {
        m.BaseRequestBuilder.PathParameters["scope"] = *scope
    }
    return m
}
// NewGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder instantiates a new GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder and sets the default values.
func NewGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getEffectivePermissions
// Deprecated: This method is obsolete. Use GetAsGetEffectivePermissionsWithScopeGetResponse instead.
// returns a GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderGetRequestConfiguration)(GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeResponseable), nil
}
// GetAsGetEffectivePermissionsWithScopeGetResponse invoke function getEffectivePermissions
// returns a GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder) GetAsGetEffectivePermissionsWithScopeGetResponse(ctx context.Context, requestConfiguration *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderGetRequestConfiguration)(GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeGetResponseable), nil
}
// ToGetRequestInformation invoke function getEffectivePermissions
// returns a *RequestInformation when successful
func (m *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder when successful
func (m *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder) WithUrl(rawUrl string)(*GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder) {
    return NewGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
