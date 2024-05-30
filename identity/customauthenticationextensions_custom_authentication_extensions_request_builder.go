package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder provides operations to manage the customAuthenticationExtensions property of the microsoft.graph.identityContainer entity.
type CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderGetQueryParameters get a list of the customAuthenticationExtension objects and their properties. The following derived types are supported.
type CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderGetQueryParameters struct {
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
// CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderGetQueryParameters
}
// CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCustomAuthenticationExtensionId provides operations to manage the customAuthenticationExtensions property of the microsoft.graph.identityContainer entity.
// returns a *CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) ByCustomAuthenticationExtensionId(customAuthenticationExtensionId string)(*CustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if customAuthenticationExtensionId != "" {
        urlTplParams["customAuthenticationExtension%2Did"] = customAuthenticationExtensionId
    }
    return NewCustomauthenticationextensionsCustomAuthenticationExtensionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderInternal instantiates a new CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder and sets the default values.
func NewCustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) {
    m := &CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/customAuthenticationExtensions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder instantiates a new CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder and sets the default values.
func NewCustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CustomauthenticationextensionsCountRequestBuilder when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) Count()(*CustomauthenticationextensionsCountRequestBuilder) {
    return NewCustomauthenticationextensionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the customAuthenticationExtension objects and their properties. The following derived types are supported.
// returns a CustomAuthenticationExtensionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitycontainer-list-customauthenticationextensions?view=graph-rest-beta
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) Get(ctx context.Context, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomAuthenticationExtensionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionCollectionResponseable), nil
}
// Post create a new customAuthenticationExtension object. The following derived types are currently supported.
// returns a CustomAuthenticationExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitycontainer-post-customauthenticationextensions?view=graph-rest-beta
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionable, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get a list of the customAuthenticationExtension objects and their properties. The following derived types are supported.
// returns a *RequestInformation when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new customAuthenticationExtension object. The following derived types are currently supported.
// returns a *RequestInformation when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAuthenticationExtensionable, requestConfiguration *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ValidateAuthenticationConfiguration provides operations to call the validateAuthenticationConfiguration method.
// returns a *CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) ValidateAuthenticationConfiguration()(*CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) {
    return NewCustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder when successful
func (m *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) WithUrl(rawUrl string)(*CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) {
    return NewCustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
