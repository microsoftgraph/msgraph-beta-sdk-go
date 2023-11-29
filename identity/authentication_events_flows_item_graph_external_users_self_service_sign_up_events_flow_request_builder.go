package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder casts the previous resource to externalUsersSelfServiceSignUpEventsFlow.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetQueryParameters get the item of type microsoft.graph.authenticationEventsFlow as microsoft.graph.externalUsersSelfServiceSignUpEventsFlow
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetQueryParameters
}
// Conditions the conditions property
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) Conditions()(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowConditionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderInternal instantiates a new GraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) {
    m := &AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder instantiates a new GraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.authenticationEventsFlow as microsoft.graph.externalUsersSelfServiceSignUpEventsFlow
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExternalUsersSelfServiceSignUpEventsFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExternalUsersSelfServiceSignUpEventsFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExternalUsersSelfServiceSignUpEventsFlowable), nil
}
// OnAttributeCollection the onAttributeCollection property
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) OnAttributeCollection()(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OnAuthenticationMethodLoadStart the onAuthenticationMethodLoadStart property
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) OnAuthenticationMethodLoadStart()(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.authenticationEventsFlow as microsoft.graph.externalUsersSelfServiceSignUpEventsFlow
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) WithUrl(rawUrl string)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
