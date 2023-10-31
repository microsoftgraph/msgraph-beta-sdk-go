package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder casts the previous resource to externalUsersSelfServiceSignUpEventsFlow.
type AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetQueryParameters get the items of type microsoft.graph.externalUsersSelfServiceSignUpEventsFlow in the microsoft.graph.authenticationEventsFlow collection
type AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetQueryParameters struct {
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
// AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetQueryParameters
}
// NewAuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderInternal instantiates a new GraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) {
    m := &AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/graph.externalUsersSelfServiceSignUpEventsFlow{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder instantiates a new GraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) Count()(*AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowCountRequestBuilder) {
    return NewAuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.externalUsersSelfServiceSignUpEventsFlow in the microsoft.graph.authenticationEventsFlow collection
func (m *AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExternalUsersSelfServiceSignUpEventsFlowCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExternalUsersSelfServiceSignUpEventsFlowCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExternalUsersSelfServiceSignUpEventsFlowCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.externalUsersSelfServiceSignUpEventsFlow in the microsoft.graph.authenticationEventsFlow collection
func (m *AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) WithUrl(rawUrl string)(*AuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) {
    return NewAuthenticationEventsFlowsGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
