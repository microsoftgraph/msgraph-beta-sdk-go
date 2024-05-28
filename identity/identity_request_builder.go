package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityRequestBuilder provides operations to manage the identityContainer singleton.
type IdentityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentityRequestBuilderGetQueryParameters get identity
type IdentityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityRequestBuilderGetQueryParameters
}
// IdentityRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApiConnectors provides operations to manage the apiConnectors property of the microsoft.graph.identityContainer entity.
// returns a *ApiconnectorsApiConnectorsRequestBuilder when successful
func (m *IdentityRequestBuilder) ApiConnectors()(*ApiconnectorsApiConnectorsRequestBuilder) {
    return NewApiconnectorsApiConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuthenticationEventListeners provides operations to manage the authenticationEventListeners property of the microsoft.graph.identityContainer entity.
// returns a *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder when successful
func (m *IdentityRequestBuilder) AuthenticationEventListeners()(*AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) {
    return NewAuthenticationeventlistenersAuthenticationEventListenersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuthenticationEventsFlows provides operations to manage the authenticationEventsFlows property of the microsoft.graph.identityContainer entity.
// returns a *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder when successful
func (m *IdentityRequestBuilder) AuthenticationEventsFlows()(*AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) {
    return NewAuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// B2cUserFlows provides operations to manage the b2cUserFlows property of the microsoft.graph.identityContainer entity.
// returns a *B2cuserflowsB2cUserFlowsRequestBuilder when successful
func (m *IdentityRequestBuilder) B2cUserFlows()(*B2cuserflowsB2cUserFlowsRequestBuilder) {
    return NewB2cuserflowsB2cUserFlowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// B2xUserFlows provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
// returns a *B2xuserflowsB2xUserFlowsRequestBuilder when successful
func (m *IdentityRequestBuilder) B2xUserFlows()(*B2xuserflowsB2xUserFlowsRequestBuilder) {
    return NewB2xuserflowsB2xUserFlowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConditionalAccess the conditionalAccess property
// returns a *ConditionalaccessConditionalAccessRequestBuilder when successful
func (m *IdentityRequestBuilder) ConditionalAccess()(*ConditionalaccessConditionalAccessRequestBuilder) {
    return NewConditionalaccessConditionalAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIdentityRequestBuilderInternal instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityRequestBuilder) {
    m := &IdentityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIdentityRequestBuilder instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityRequestBuilderInternal(urlParams, requestAdapter)
}
// ContinuousAccessEvaluationPolicy provides operations to manage the continuousAccessEvaluationPolicy property of the microsoft.graph.identityContainer entity.
// returns a *ContinuousaccessevaluationpolicyContinuousAccessEvaluationPolicyRequestBuilder when successful
func (m *IdentityRequestBuilder) ContinuousAccessEvaluationPolicy()(*ContinuousaccessevaluationpolicyContinuousAccessEvaluationPolicyRequestBuilder) {
    return NewContinuousaccessevaluationpolicyContinuousAccessEvaluationPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomAuthenticationExtensions provides operations to manage the customAuthenticationExtensions property of the microsoft.graph.identityContainer entity.
// returns a *CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder when successful
func (m *IdentityRequestBuilder) CustomAuthenticationExtensions()(*CustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilder) {
    return NewCustomauthenticationextensionsCustomAuthenticationExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get identity
// returns a IdentityContainerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentityRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable), nil
}
// IdentityProviders provides operations to manage the identityProviders property of the microsoft.graph.identityContainer entity.
// returns a *IdentityprovidersIdentityProvidersRequestBuilder when successful
func (m *IdentityRequestBuilder) IdentityProviders()(*IdentityprovidersIdentityProvidersRequestBuilder) {
    return NewIdentityprovidersIdentityProvidersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update identity
// returns a IdentityContainerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentityRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, requestConfiguration *IdentityRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable), nil
}
// ToGetRequestInformation get identity
// returns a *RequestInformation when successful
func (m *IdentityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update identity
// returns a *RequestInformation when successful
func (m *IdentityRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, requestConfiguration *IdentityRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// UserFlowAttributes provides operations to manage the userFlowAttributes property of the microsoft.graph.identityContainer entity.
// returns a *UserflowattributesUserFlowAttributesRequestBuilder when successful
func (m *IdentityRequestBuilder) UserFlowAttributes()(*UserflowattributesUserFlowAttributesRequestBuilder) {
    return NewUserflowattributesUserFlowAttributesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserFlows provides operations to manage the userFlows property of the microsoft.graph.identityContainer entity.
// returns a *UserflowsUserFlowsRequestBuilder when successful
func (m *IdentityRequestBuilder) UserFlows()(*UserflowsUserFlowsRequestBuilder) {
    return NewUserflowsUserFlowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IdentityRequestBuilder when successful
func (m *IdentityRequestBuilder) WithUrl(rawUrl string)(*IdentityRequestBuilder) {
    return NewIdentityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
