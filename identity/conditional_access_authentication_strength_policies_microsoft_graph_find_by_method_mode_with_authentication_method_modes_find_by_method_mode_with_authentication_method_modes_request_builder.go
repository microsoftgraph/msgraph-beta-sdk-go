package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder provides operations to call the findByMethodMode method.
type ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters invoke function findByMethodMode
type ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters struct {
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
// ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters
}
// NewConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal instantiates a new FindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, authenticationMethodModes *string)(*ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    m := &ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/microsoft.graph.findByMethodMode(authenticationMethodModes={authenticationMethodModes}){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
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
// NewConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder instantiates a new FindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function findByMethodMode
func (m *ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesResponseable), nil
}
// ToGetRequestInformation invoke function findByMethodMode
func (m *ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesMicrosoftGraphFindByMethodModeWithAuthenticationMethodModesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
