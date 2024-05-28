package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder provides operations to call the findByMethodMode method.
type AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
type AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters struct {
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
// AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters
}
// NewAuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal instantiates a new AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewAuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, authenticationMethodModes *string)(*AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    m := &AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationStrengthPolicies/findByMethodMode(authenticationMethodModes={authenticationMethodModes}){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if authenticationMethodModes != nil {
        m.BaseRequestBuilder.PathParameters["authenticationMethodModes"] = *authenticationMethodModes
    }
    return m
}
// NewAuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder instantiates a new AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewAuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
// Deprecated: This method is obsolete. Use GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse instead.
// returns a AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-findbymethodmode?view=graph-rest-beta
func (m *AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesResponseable), nil
}
// GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
// Deprecated: The findByMethodMode function is deprecated. Please use OData filter query instead. as of 2023-02/FindByMethodModeRemove
// returns a AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-findbymethodmode?view=graph-rest-beta
func (m *AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse(ctx context.Context, requestConfiguration *AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesGetResponseable), nil
}
// ToGetRequestInformation get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
// Deprecated: The findByMethodMode function is deprecated. Please use OData filter query instead. as of 2023-02/FindByMethodModeRemove
// returns a *RequestInformation when successful
func (m *AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The findByMethodMode function is deprecated. Please use OData filter query instead. as of 2023-02/FindByMethodModeRemove
// returns a *AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder when successful
func (m *AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) WithUrl(rawUrl string)(*AuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    return NewAuthenticationstrengthpoliciesFindbymethodmodewithauthenticationmethodmodesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
