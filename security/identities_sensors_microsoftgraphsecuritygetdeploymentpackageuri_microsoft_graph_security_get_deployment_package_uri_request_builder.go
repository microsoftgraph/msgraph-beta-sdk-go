package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder provides operations to call the getDeploymentPackageUri method.
type IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderInternal instantiates a new IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder and sets the default values.
func NewIdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) {
    m := &IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/identities/sensors/microsoft.graph.security.getDeploymentPackageUri()", pathParameters),
    }
    return m
}
// NewIdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder instantiates a new IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder and sets the default values.
func NewIdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getDeploymentPackageUri
// returns a SensorDeploymentPackageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensorDeploymentPackageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSensorDeploymentPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensorDeploymentPackageable), nil
}
// ToGetRequestInformation invoke function getDeploymentPackageUri
// returns a *RequestInformation when successful
func (m *IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder when successful
func (m *IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) WithUrl(rawUrl string)(*IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) {
    return NewIdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
