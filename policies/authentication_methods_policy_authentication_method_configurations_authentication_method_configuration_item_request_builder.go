package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder provides operations to manage the authenticationMethodConfigurations property of the microsoft.graph.authenticationMethodsPolicy entity.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetQueryParameters read the properties and relationships of an externalAuthenticationMethodConfiguration object.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetQueryParameters
}
// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal instantiates a new AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder and sets the default values.
func NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) {
    m := &AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/{authenticationMethodConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder instantiates a new AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder and sets the default values.
func NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an externalAuthenticationMethodConfiguration object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/externalauthenticationmethodconfiguration-delete?view=graph-rest-beta
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an externalAuthenticationMethodConfiguration object.
// returns a AuthenticationMethodConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/externalauthenticationmethodconfiguration-get?view=graph-rest-beta
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationMethodConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodConfigurationable), nil
}
// Patch update the properties of an externalAuthenticationMethodConfiguration object.
// returns a AuthenticationMethodConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/externalauthenticationmethodconfiguration-update?view=graph-rest-beta
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodConfigurationable, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationMethodConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodConfigurationable), nil
}
// ToDeleteRequestInformation delete an externalAuthenticationMethodConfiguration object.
// returns a *RequestInformation when successful
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an externalAuthenticationMethodConfiguration object.
// returns a *RequestInformation when successful
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an externalAuthenticationMethodConfiguration object.
// returns a *RequestInformation when successful
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodConfigurationable, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder when successful
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) {
    return NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
