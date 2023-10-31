package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder builds and executes requests for operations under \identity\authenticationEventsFlows\{authenticationEventsFlow-id}\graph.externalUsersSelfServiceSignUpEventsFlow\onAttributeCollection
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderGetQueryParameters the configuration for what to invoke when attributes are ready to be collected from the user.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderGetQueryParameters
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderInternal instantiates a new OnAttributeCollectionRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder) {
    m := &AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder instantiates a new OnAttributeCollectionRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the configuration for what to invoke when attributes are ready to be collected from the user.
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnAttributeCollectionHandlerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnAttributeCollectionHandlerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnAttributeCollectionHandlerable), nil
}
// GraphOnAttributeCollectionExternalUsersSelfServiceSignUp casts the previous resource to onAttributeCollectionExternalUsersSelfServiceSignUp.
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder) GraphOnAttributeCollectionExternalUsersSelfServiceSignUp()(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the configuration for what to invoke when attributes are ready to be collected from the user.
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder) WithUrl(rawUrl string)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
