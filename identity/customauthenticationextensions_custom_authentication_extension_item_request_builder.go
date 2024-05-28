package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder provides operations to manage the customAuthenticationExtensions property of the microsoft.graph.identityContainer entity.
type CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderGetQueryParameters read the properties and relationships of an authenticationEventListener object. The @odata.type property in the response object indicates the type of the authenticationEventListener object.
type CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderGetQueryParameters
}
// CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderInternal instantiates a new CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder and sets the default values.
func NewCustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) {
    m := &CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/customAuthenticationExtensions/{customAuthenticationExtension%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder instantiates a new CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder and sets the default values.
func NewCustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a customAuthenticationExtension object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/customauthenticationextension-delete?view=graph-rest-beta
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an authenticationEventListener object. The @odata.type property in the response object indicates the type of the authenticationEventListener object.
// returns a CustomAuthenticationExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationeventlistener-get?view=graph-rest-beta
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomAuthenticationExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionable), nil
}
// Patch update the navigation property customAuthenticationExtensions in identity
// returns a CustomAuthenticationExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionable, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomAuthenticationExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionable), nil
}
// ToDeleteRequestInformation delete a customAuthenticationExtension object.
// returns a *RequestInformation when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an authenticationEventListener object. The @odata.type property in the response object indicates the type of the authenticationEventListener object.
// returns a *RequestInformation when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property customAuthenticationExtensions in identity
// returns a *RequestInformation when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionable, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ValidateAuthenticationConfiguration provides operations to call the validateAuthenticationConfiguration method.
// returns a *CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) ValidateAuthenticationConfiguration()(*CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) {
    return NewCustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) WithUrl(rawUrl string)(*CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) {
    return NewCustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
