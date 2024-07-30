package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemFederatedIdentityCredentialsWithNameRequestBuilder provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.application entity.
type ItemFederatedIdentityCredentialsWithNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFederatedIdentityCredentialsWithNameRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFederatedIdentityCredentialsWithNameRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemFederatedIdentityCredentialsWithNameRequestBuilderGetQueryParameters read the properties and relationships of a federatedIdentityCredential object.
type ItemFederatedIdentityCredentialsWithNameRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemFederatedIdentityCredentialsWithNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFederatedIdentityCredentialsWithNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemFederatedIdentityCredentialsWithNameRequestBuilderGetQueryParameters
}
// ItemFederatedIdentityCredentialsWithNameRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFederatedIdentityCredentialsWithNameRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemFederatedIdentityCredentialsWithNameRequestBuilderInternal instantiates a new ItemFederatedIdentityCredentialsWithNameRequestBuilder and sets the default values.
func NewItemFederatedIdentityCredentialsWithNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, name *string)(*ItemFederatedIdentityCredentialsWithNameRequestBuilder) {
    m := &ItemFederatedIdentityCredentialsWithNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/federatedIdentityCredentials(name='{name}'){?%24expand,%24select}", pathParameters),
    }
    if name != nil {
        m.BaseRequestBuilder.PathParameters["name"] = *name
    }
    return m
}
// NewItemFederatedIdentityCredentialsWithNameRequestBuilder instantiates a new ItemFederatedIdentityCredentialsWithNameRequestBuilder and sets the default values.
func NewItemFederatedIdentityCredentialsWithNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFederatedIdentityCredentialsWithNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFederatedIdentityCredentialsWithNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete deletes a federatedIdentityCredential object from an application.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/federatedidentitycredential-delete?view=graph-rest-beta
func (m *ItemFederatedIdentityCredentialsWithNameRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemFederatedIdentityCredentialsWithNameRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a federatedIdentityCredential object.
// returns a FederatedIdentityCredentialable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/federatedidentitycredential-get?view=graph-rest-beta
func (m *ItemFederatedIdentityCredentialsWithNameRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFederatedIdentityCredentialsWithNameRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedIdentityCredentialable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateFederatedIdentityCredentialFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedIdentityCredentialable), nil
}
// Patch create a new federatedIdentityCredential object for an application if it doesn't exist, or update the properties of an existing federatedIdentityCredential object. By configuring a trust relationship between your Microsoft Entra application registration and the identity provider for your compute platform, you can use tokens issued by that platform to authenticate with Microsoft identity platform and call APIs in the Microsoft ecosystem. Maximum of 20 objects can be added to an application.
// returns a FederatedIdentityCredentialable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/federatedidentitycredential-upsert?view=graph-rest-beta
func (m *ItemFederatedIdentityCredentialsWithNameRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedIdentityCredentialable, requestConfiguration *ItemFederatedIdentityCredentialsWithNameRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedIdentityCredentialable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateFederatedIdentityCredentialFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedIdentityCredentialable), nil
}
// ToDeleteRequestInformation deletes a federatedIdentityCredential object from an application.
// returns a *RequestInformation when successful
func (m *ItemFederatedIdentityCredentialsWithNameRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemFederatedIdentityCredentialsWithNameRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a federatedIdentityCredential object.
// returns a *RequestInformation when successful
func (m *ItemFederatedIdentityCredentialsWithNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFederatedIdentityCredentialsWithNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation create a new federatedIdentityCredential object for an application if it doesn't exist, or update the properties of an existing federatedIdentityCredential object. By configuring a trust relationship between your Microsoft Entra application registration and the identity provider for your compute platform, you can use tokens issued by that platform to authenticate with Microsoft identity platform and call APIs in the Microsoft ecosystem. Maximum of 20 objects can be added to an application.
// returns a *RequestInformation when successful
func (m *ItemFederatedIdentityCredentialsWithNameRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedIdentityCredentialable, requestConfiguration *ItemFederatedIdentityCredentialsWithNameRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFederatedIdentityCredentialsWithNameRequestBuilder when successful
func (m *ItemFederatedIdentityCredentialsWithNameRequestBuilder) WithUrl(rawUrl string)(*ItemFederatedIdentityCredentialsWithNameRequestBuilder) {
    return NewItemFederatedIdentityCredentialsWithNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
