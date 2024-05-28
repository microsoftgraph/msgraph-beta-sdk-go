package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentitiesSensorsRequestBuilder provides operations to manage the sensors property of the microsoft.graph.security.identityContainer entity.
type IdentitiesSensorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentitiesSensorsRequestBuilderGetQueryParameters get sensors from security
type IdentitiesSensorsRequestBuilderGetQueryParameters struct {
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
// IdentitiesSensorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitiesSensorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentitiesSensorsRequestBuilderGetQueryParameters
}
// IdentitiesSensorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitiesSensorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySensorId provides operations to manage the sensors property of the microsoft.graph.security.identityContainer entity.
// returns a *IdentitiesSensorsSensorItemRequestBuilder when successful
func (m *IdentitiesSensorsRequestBuilder) BySensorId(sensorId string)(*IdentitiesSensorsSensorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sensorId != "" {
        urlTplParams["sensor%2Did"] = sensorId
    }
    return NewIdentitiesSensorsSensorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIdentitiesSensorsRequestBuilderInternal instantiates a new IdentitiesSensorsRequestBuilder and sets the default values.
func NewIdentitiesSensorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsRequestBuilder) {
    m := &IdentitiesSensorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/identities/sensors{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIdentitiesSensorsRequestBuilder instantiates a new IdentitiesSensorsRequestBuilder and sets the default values.
func NewIdentitiesSensorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentitiesSensorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IdentitiesSensorsCountRequestBuilder when successful
func (m *IdentitiesSensorsRequestBuilder) Count()(*IdentitiesSensorsCountRequestBuilder) {
    return NewIdentitiesSensorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get sensors from security
// returns a SensorCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentitiesSensorsRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentitiesSensorsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSensorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensorCollectionResponseable), nil
}
// MicrosoftGraphSecurityGetDeploymentAccessKey provides operations to call the getDeploymentAccessKey method.
// returns a *IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentaccesskeyMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder when successful
func (m *IdentitiesSensorsRequestBuilder) MicrosoftGraphSecurityGetDeploymentAccessKey()(*IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentaccesskeyMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilder) {
    return NewIdentitiesSensorsMicrosoftgraphsecuritygetdeploymentaccesskeyMicrosoftGraphSecurityGetDeploymentAccessKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityGetDeploymentPackageUri provides operations to call the getDeploymentPackageUri method.
// returns a *IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder when successful
func (m *IdentitiesSensorsRequestBuilder) MicrosoftGraphSecurityGetDeploymentPackageUri()(*IdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) {
    return NewIdentitiesSensorsMicrosoftgraphsecuritygetdeploymentpackageuriMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityRegenerateDeploymentAccessKey provides operations to call the regenerateDeploymentAccessKey method.
// returns a *IdentitiesSensorsMicrosoftgraphsecurityregeneratedeploymentaccesskeyMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder when successful
func (m *IdentitiesSensorsRequestBuilder) MicrosoftGraphSecurityRegenerateDeploymentAccessKey()(*IdentitiesSensorsMicrosoftgraphsecurityregeneratedeploymentaccesskeyMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder) {
    return NewIdentitiesSensorsMicrosoftgraphsecurityregeneratedeploymentaccesskeyMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to sensors for security
// returns a Sensorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentitiesSensorsRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Sensorable, requestConfiguration *IdentitiesSensorsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Sensorable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSensorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Sensorable), nil
}
// ToGetRequestInformation get sensors from security
// returns a *RequestInformation when successful
func (m *IdentitiesSensorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentitiesSensorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to sensors for security
// returns a *RequestInformation when successful
func (m *IdentitiesSensorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Sensorable, requestConfiguration *IdentitiesSensorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IdentitiesSensorsRequestBuilder when successful
func (m *IdentitiesSensorsRequestBuilder) WithUrl(rawUrl string)(*IdentitiesSensorsRequestBuilder) {
    return NewIdentitiesSensorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
