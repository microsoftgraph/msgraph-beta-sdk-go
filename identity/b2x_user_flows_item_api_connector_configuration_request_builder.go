package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2xUserFlowsItemApiConnectorConfigurationRequestBuilder builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\apiConnectorConfiguration
type B2xUserFlowsItemApiConnectorConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xUserFlowsItemApiConnectorConfigurationRequestBuilderGetQueryParameters configuration for enabling an API connector for use as part of the self-service sign up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
type B2xUserFlowsItemApiConnectorConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xUserFlowsItemApiConnectorConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemApiConnectorConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xUserFlowsItemApiConnectorConfigurationRequestBuilderGetQueryParameters
}
// NewB2xUserFlowsItemApiConnectorConfigurationRequestBuilderInternal instantiates a new ApiConnectorConfigurationRequestBuilder and sets the default values.
func NewB2xUserFlowsItemApiConnectorConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemApiConnectorConfigurationRequestBuilder) {
    m := &B2xUserFlowsItemApiConnectorConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/apiConnectorConfiguration{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewB2xUserFlowsItemApiConnectorConfigurationRequestBuilder instantiates a new ApiConnectorConfigurationRequestBuilder and sets the default values.
func NewB2xUserFlowsItemApiConnectorConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemApiConnectorConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xUserFlowsItemApiConnectorConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get configuration for enabling an API connector for use as part of the self-service sign up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2xUserFlowsItemApiConnectorConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xUserFlowsItemApiConnectorConfigurationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowApiConnectorConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *B2xUserFlowsItemApiConnectorConfigurationRequestBuilder) PostAttributeCollection()(*B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionRequestBuilder) {
    return NewB2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PostFederationSignup provides operations to manage the postFederationSignup property of the microsoft.graph.userFlowApiConnectorConfiguration entity.
func (m *B2xUserFlowsItemApiConnectorConfigurationRequestBuilder) PostFederationSignup()(*B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupRequestBuilder) {
    return NewB2xUserFlowsItemApiConnectorConfigurationPostFederationSignupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PreTokenIssuance provides operations to manage the preTokenIssuance property of the microsoft.graph.userFlowApiConnectorConfiguration entity.
func (m *B2xUserFlowsItemApiConnectorConfigurationRequestBuilder) PreTokenIssuance()(*B2xUserFlowsItemApiConnectorConfigurationPreTokenIssuanceRequestBuilder) {
    return NewB2xUserFlowsItemApiConnectorConfigurationPreTokenIssuanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation configuration for enabling an API connector for use as part of the self-service sign up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2xUserFlowsItemApiConnectorConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xUserFlowsItemApiConnectorConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *B2xUserFlowsItemApiConnectorConfigurationRequestBuilder) WithUrl(rawUrl string)(*B2xUserFlowsItemApiConnectorConfigurationRequestBuilder) {
    return NewB2xUserFlowsItemApiConnectorConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
