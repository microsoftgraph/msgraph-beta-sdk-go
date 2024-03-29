package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
type AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderGetQueryParameters get the authenticationCombinationConfiguration objects for an authentication strength policy. The objects can be of one or more of the following derived types:* fido2combinationConfigurations* x509certificatecombinationconfiguration authenticationCombinationConfiguration objects are supported only for custom authentication strengths.
type AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderGetQueryParameters struct {
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
// AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderGetQueryParameters
}
// AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationCombinationConfigurationId provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
// returns a *AuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder when successful
func (m *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) ByAuthenticationCombinationConfigurationId(authenticationCombinationConfigurationId string)(*AuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationCombinationConfigurationId != "" {
        urlTplParams["authenticationCombinationConfiguration%2Did"] = authenticationCombinationConfigurationId
    }
    return NewAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderInternal instantiates a new AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder and sets the default values.
func NewAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) {
    m := &AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationStrengthPolicies/{authenticationStrengthPolicy%2Did}/combinationConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder instantiates a new AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder and sets the default values.
func NewAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationStrengthPoliciesItemCombinationConfigurationsCountRequestBuilder when successful
func (m *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) Count()(*AuthenticationStrengthPoliciesItemCombinationConfigurationsCountRequestBuilder) {
    return NewAuthenticationStrengthPoliciesItemCombinationConfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the authenticationCombinationConfiguration objects for an authentication strength policy. The objects can be of one or more of the following derived types:* fido2combinationConfigurations* x509certificatecombinationconfiguration authenticationCombinationConfiguration objects are supported only for custom authentication strengths.
// returns a AuthenticationCombinationConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-list-combinationconfigurations?view=graph-rest-1.0
func (m *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationCombinationConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationCombinationConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationCombinationConfigurationCollectionResponseable), nil
}
// Post create a new authenticationCombinationConfiguration object which can be of one of the following derived types:* fido2combinationConfigurations* x509certificatecombinationconfiguration
// returns a AuthenticationCombinationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-post-combinationconfigurations?view=graph-rest-1.0
func (m *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationCombinationConfigurationable, requestConfiguration *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationCombinationConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationCombinationConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationCombinationConfigurationable), nil
}
// ToGetRequestInformation get the authenticationCombinationConfiguration objects for an authentication strength policy. The objects can be of one or more of the following derived types:* fido2combinationConfigurations* x509certificatecombinationconfiguration authenticationCombinationConfiguration objects are supported only for custom authentication strengths.
// returns a *RequestInformation when successful
func (m *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new authenticationCombinationConfiguration object which can be of one of the following derived types:* fido2combinationConfigurations* x509certificatecombinationconfiguration
// returns a *RequestInformation when successful
func (m *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationCombinationConfigurationable, requestConfiguration *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/policies/authenticationStrengthPolicies/{authenticationStrengthPolicy%2Did}/combinationConfigurations", m.BaseRequestBuilder.PathParameters)
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
// returns a *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder when successful
func (m *AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) {
    return NewAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
