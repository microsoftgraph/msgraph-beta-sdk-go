// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder provides operations to call the findByMethodMode method.
type AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
type AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters struct {
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
// AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters
}
// NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal instantiates a new AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, authenticationMethodModes *string)(*AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    m := &AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationStrengthPolicies/findByMethodMode(authenticationMethodModes={authenticationMethodModes}){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if authenticationMethodModes != nil {
        m.BaseRequestBuilder.PathParameters["authenticationMethodModes"] = *authenticationMethodModes
    }
    return m
}
// NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder instantiates a new AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
// Deprecated: This method is obsolete. Use GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse instead.
// returns a AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-findbymethodmode?view=graph-rest-beta
func (m *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesResponseable), nil
}
// GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
// Deprecated: The findByMethodMode function is deprecated. Please use OData filter query instead. as of 2023-02/FindByMethodModeRemove on 2023-02-01 and will be removed 2023-03-31
// returns a AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-findbymethodmode?view=graph-rest-beta
func (m *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable), nil
}
// ToGetRequestInformation get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
// Deprecated: The findByMethodMode function is deprecated. Please use OData filter query instead. as of 2023-02/FindByMethodModeRemove on 2023-02-01 and will be removed 2023-03-31
// returns a *RequestInformation when successful
func (m *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The findByMethodMode function is deprecated. Please use OData filter query instead. as of 2023-02/FindByMethodModeRemove on 2023-02-01 and will be removed 2023-03-31
// returns a *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder when successful
func (m *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) WithUrl(rawUrl string)(*AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    return NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
