package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder provides operations to call the getDeploymentAccessKey method.
type IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilderInternal instantiates a new IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder and sets the default values.
func NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder) {
    m := &IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/identities/sensors/microsoft.graph.security.getDeploymentAccessKey()", pathParameters),
    }
    return m
}
// NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder instantiates a new IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder and sets the default values.
func NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getDeploymentAccessKey
// returns a DeploymentAccessKeyTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DeploymentAccessKeyTypeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateDeploymentAccessKeyTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DeploymentAccessKeyTypeable), nil
}
// ToGetRequestInformation invoke function getDeploymentAccessKey
// returns a *RequestInformation when successful
func (m *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder when successful
func (m *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder) WithUrl(rawUrl string)(*IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder) {
    return NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
