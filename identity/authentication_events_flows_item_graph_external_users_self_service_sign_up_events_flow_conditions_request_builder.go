package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder builds and executes requests for operations under \identity\authenticationEventsFlows\{authenticationEventsFlow-id}\graph.externalUsersSelfServiceSignUpEventsFlow\conditions
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderGetQueryParameters the conditions representing the context of the authentication request that will be used to decide whether the events policy will be invoked.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderGetQueryParameters
}
// Applications the applications property
// returns a *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsApplicationsRequestBuilder when successful
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder) Applications()(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsApplicationsRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderInternal instantiates a new AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder) {
    m := &AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/conditions{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder instantiates a new AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the conditions representing the context of the authentication request that will be used to decide whether the events policy will be invoked.
// returns a AuthenticationConditionsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConditionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationConditionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConditionsable), nil
}
// ToGetRequestInformation the conditions representing the context of the authentication request that will be used to decide whether the events policy will be invoked.
// returns a *RequestInformation when successful
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder when successful
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
