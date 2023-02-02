package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder provides operations to call the findByMethodMode method.
type AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters invoke function findByMethodMode
type AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters
}
// NewAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal instantiates a new FindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, authenticationMethodModes *string)(*AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    m := &AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/authenticationStrengthPolicies/microsoft.graph.findByMethodMode(authenticationMethodModes={authenticationMethodModes}){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if authenticationMethodModes != nil {
        urlTplParams["authenticationMethodModes"] = *authenticationMethodModes
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder instantiates a new FindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function findByMethodMode
func (m *AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesResponseable), nil
}
// ToGetRequestInformation invoke function findByMethodMode
func (m *AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
