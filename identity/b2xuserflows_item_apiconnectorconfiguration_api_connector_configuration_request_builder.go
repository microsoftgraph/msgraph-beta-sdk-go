package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\apiConnectorConfiguration
type B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderGetQueryParameters get the apiConnectorConfiguration property in a b2xIdentityUserFlow to detail the API connectors enabled for the user flow.
type B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderGetQueryParameters
}
// NewB2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderInternal instantiates a new B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder and sets the default values.
func NewB2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) {
    m := &B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/apiConnectorConfiguration{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder instantiates a new B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder and sets the default values.
func NewB2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the apiConnectorConfiguration property in a b2xIdentityUserFlow to detail the API connectors enabled for the user flow.
// returns a UserFlowApiConnectorConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/b2xidentityuserflow-get-apiconnectorconfiguration?view=graph-rest-beta
func (m *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowApiConnectorConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowApiConnectorConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowApiConnectorConfigurationable), nil
}
// PostAttributeCollection provides operations to manage the postAttributeCollection property of the microsoft.graph.userFlowApiConnectorConfiguration entity.
// returns a *B2xuserflowsItemApiconnectorconfigurationPostattributecollectionPostAttributeCollectionRequestBuilder when successful
func (m *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) PostAttributeCollection()(*B2xuserflowsItemApiconnectorconfigurationPostattributecollectionPostAttributeCollectionRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationPostattributecollectionPostAttributeCollectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PostFederationSignup provides operations to manage the postFederationSignup property of the microsoft.graph.userFlowApiConnectorConfiguration entity.
// returns a *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder when successful
func (m *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) PostFederationSignup()(*B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PreTokenIssuance provides operations to manage the preTokenIssuance property of the microsoft.graph.userFlowApiConnectorConfiguration entity.
// returns a *B2xuserflowsItemApiconnectorconfigurationPretokenissuancePreTokenIssuanceRequestBuilder when successful
func (m *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) PreTokenIssuance()(*B2xuserflowsItemApiconnectorconfigurationPretokenissuancePreTokenIssuanceRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationPretokenissuancePreTokenIssuanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the apiConnectorConfiguration property in a b2xIdentityUserFlow to detail the API connectors enabled for the user flow.
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder when successful
func (m *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
