package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityProvidersIdentityProviderBaseItemRequestBuilder provides operations to manage the identityProviders property of the microsoft.graph.identityContainer entity.
type IdentityProvidersIdentityProviderBaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentityProvidersIdentityProviderBaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityProvidersIdentityProviderBaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityProvidersIdentityProviderBaseItemRequestBuilderGetQueryParameters get the properties and relationships of the specified identity provider configured in the tenant. Among the types of providers derived from identityProviderBase, you can currently get a socialIdentityProvider or a builtinIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently get a socialIdentityProvider, openIdConnectIdentityProvider, or an appleManagedIdentityProvider resource.
type IdentityProvidersIdentityProviderBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityProvidersIdentityProviderBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityProvidersIdentityProviderBaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityProvidersIdentityProviderBaseItemRequestBuilderGetQueryParameters
}
// IdentityProvidersIdentityProviderBaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityProvidersIdentityProviderBaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal instantiates a new IdentityProviderBaseItemRequestBuilder and sets the default values.
func NewIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    m := &IdentityProvidersIdentityProviderBaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/identityProviders/{identityProviderBase%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewIdentityProvidersIdentityProviderBaseItemRequestBuilder instantiates a new IdentityProviderBaseItemRequestBuilder and sets the default values.
func NewIdentityProvidersIdentityProviderBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an identity provider resource that is of the type specified by the id in the request. Among the types of providers derived from identityProviderBase, you can currently delete a socialIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently delete a socialIdentityProvider, openIdConnectIdentityProvider, or an appleManagedIdentityProvider resource.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityproviderbase-delete?view=graph-rest-1.0
func (m *IdentityProvidersIdentityProviderBaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityProvidersIdentityProviderBaseItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get the properties and relationships of the specified identity provider configured in the tenant. Among the types of providers derived from identityProviderBase, you can currently get a socialIdentityProvider or a builtinIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently get a socialIdentityProvider, openIdConnectIdentityProvider, or an appleManagedIdentityProvider resource.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityproviderbase-get?view=graph-rest-1.0
func (m *IdentityProvidersIdentityProviderBaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityProvidersIdentityProviderBaseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityProviderBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseable), nil
}
// Patch update the properties of the specified identity provider configured in the tenant. Among the types of providers derived from identityProviderBase, you can currently update a socialIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently update a socialIdentityProvider, openIdConnectIdentityProvider, or an appleManagedIdentityProvider resource.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityproviderbase-update?view=graph-rest-1.0
func (m *IdentityProvidersIdentityProviderBaseItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseable, requestConfiguration *IdentityProvidersIdentityProviderBaseItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityProviderBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseable), nil
}
// ToDeleteRequestInformation delete an identity provider resource that is of the type specified by the id in the request. Among the types of providers derived from identityProviderBase, you can currently delete a socialIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently delete a socialIdentityProvider, openIdConnectIdentityProvider, or an appleManagedIdentityProvider resource.
func (m *IdentityProvidersIdentityProviderBaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityProvidersIdentityProviderBaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of the specified identity provider configured in the tenant. Among the types of providers derived from identityProviderBase, you can currently get a socialIdentityProvider or a builtinIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently get a socialIdentityProvider, openIdConnectIdentityProvider, or an appleManagedIdentityProvider resource.
func (m *IdentityProvidersIdentityProviderBaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentityProvidersIdentityProviderBaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of the specified identity provider configured in the tenant. Among the types of providers derived from identityProviderBase, you can currently update a socialIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently update a socialIdentityProvider, openIdConnectIdentityProvider, or an appleManagedIdentityProvider resource.
func (m *IdentityProvidersIdentityProviderBaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseable, requestConfiguration *IdentityProvidersIdentityProviderBaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityProvidersIdentityProviderBaseItemRequestBuilder) WithUrl(rawUrl string)(*IdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    return NewIdentityProvidersIdentityProviderBaseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
