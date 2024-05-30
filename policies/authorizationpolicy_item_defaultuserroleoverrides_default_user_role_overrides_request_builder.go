package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder provides operations to manage the defaultUserRoleOverrides property of the microsoft.graph.authorizationPolicy entity.
type AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderGetQueryParameters get defaultUserRoleOverrides from policies
type AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderGetQueryParameters struct {
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
// AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderGetQueryParameters
}
// AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDefaultUserRoleOverrideId provides operations to manage the defaultUserRoleOverrides property of the microsoft.graph.authorizationPolicy entity.
// returns a *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder when successful
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) ByDefaultUserRoleOverrideId(defaultUserRoleOverrideId string)(*AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if defaultUserRoleOverrideId != "" {
        urlTplParams["defaultUserRoleOverride%2Did"] = defaultUserRoleOverrideId
    }
    return NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderInternal instantiates a new AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder and sets the default values.
func NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) {
    m := &AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authorizationPolicy/{authorizationPolicy%2Did}/defaultUserRoleOverrides{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder instantiates a new AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder and sets the default values.
func NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthorizationpolicyItemDefaultuserroleoverridesCountRequestBuilder when successful
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) Count()(*AuthorizationpolicyItemDefaultuserroleoverridesCountRequestBuilder) {
    return NewAuthorizationpolicyItemDefaultuserroleoverridesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get defaultUserRoleOverrides from policies
// returns a DefaultUserRoleOverrideCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDefaultUserRoleOverrideCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideCollectionResponseable), nil
}
// Post create new navigation property to defaultUserRoleOverrides for policies
// returns a DefaultUserRoleOverrideable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideable, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDefaultUserRoleOverrideFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideable), nil
}
// ToGetRequestInformation get defaultUserRoleOverrides from policies
// returns a *RequestInformation when successful
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to defaultUserRoleOverrides for policies
// returns a *RequestInformation when successful
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideable, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder when successful
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) WithUrl(rawUrl string)(*AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder) {
    return NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverridesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
