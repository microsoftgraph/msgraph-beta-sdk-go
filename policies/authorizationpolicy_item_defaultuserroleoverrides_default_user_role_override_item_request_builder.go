package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder provides operations to manage the defaultUserRoleOverrides property of the microsoft.graph.authorizationPolicy entity.
type AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderGetQueryParameters get defaultUserRoleOverrides from policies
type AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderGetQueryParameters
}
// AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderInternal instantiates a new AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder and sets the default values.
func NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) {
    m := &AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authorizationPolicy/{authorizationPolicy%2Did}/defaultUserRoleOverrides/{defaultUserRoleOverride%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder instantiates a new AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder and sets the default values.
func NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property defaultUserRoleOverrides for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get defaultUserRoleOverrides from policies
// returns a DefaultUserRoleOverrideable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property defaultUserRoleOverrides in policies
// returns a DefaultUserRoleOverrideable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideable, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property defaultUserRoleOverrides for policies
// returns a *RequestInformation when successful
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get defaultUserRoleOverrides from policies
// returns a *RequestInformation when successful
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property defaultUserRoleOverrides in policies
// returns a *RequestInformation when successful
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DefaultUserRoleOverrideable, requestConfiguration *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder when successful
func (m *AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) WithUrl(rawUrl string)(*AuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder) {
    return NewAuthorizationpolicyItemDefaultuserroleoverridesDefaultUserRoleOverrideItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
