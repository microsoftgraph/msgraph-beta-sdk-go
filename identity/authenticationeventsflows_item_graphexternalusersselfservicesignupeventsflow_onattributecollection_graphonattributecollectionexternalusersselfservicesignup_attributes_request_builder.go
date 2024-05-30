package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder provides operations to manage the attributes property of the microsoft.graph.onAttributeCollectionExternalUsersSelfServiceSignUp entity.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderGetQueryParameters get an identityUserFlowAttribute collection associated with an external identities self-service sign up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object. These are the attributes that are collected from the user during the authentication experience that's defined by the user flow.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderGetQueryParameters struct {
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
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderGetQueryParameters
}
// ByIdentityUserFlowAttributeId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.authenticationEventsFlows.item.graphExternalUsersSelfServiceSignUpEventsFlow.onAttributeCollection.graphOnAttributeCollectionExternalUsersSelfServiceSignUp.attributes.item collection
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesIdentityUserFlowAttributeItemRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder) ByIdentityUserFlowAttributeId(identityUserFlowAttributeId string)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesIdentityUserFlowAttributeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if identityUserFlowAttributeId != "" {
        urlTplParams["identityUserFlowAttribute%2Did"] = identityUserFlowAttributeId
    }
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesIdentityUserFlowAttributeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderInternal instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder) {
    m := &AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/graph.onAttributeCollectionExternalUsersSelfServiceSignUp/attributes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesCountRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder) Count()(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesCountRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get an identityUserFlowAttribute collection associated with an external identities self-service sign up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object. These are the attributes that are collected from the user during the authentication experience that's defined by the user flow.
// returns a IdentityUserFlowAttributeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onattributecollectionexternalusersselfservicesignup-list-attributes?view=graph-rest-beta
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityUserFlowAttributeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeCollectionResponseable), nil
}
// Ref provides operations to manage the collection of identityContainer entities.
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder) Ref()(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get an identityUserFlowAttribute collection associated with an external identities self-service sign up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object. These are the attributes that are collected from the user during the authentication experience that's defined by the user flow.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
