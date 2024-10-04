package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentitiesSensorsItemHealthIssuesRequestBuilder provides operations to manage the healthIssues property of the microsoft.graph.security.sensor entity.
type IdentitiesSensorsItemHealthIssuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentitiesSensorsItemHealthIssuesRequestBuilderGetQueryParameters represents potential issues within a customer's Microsoft Defender for Identity configuration that Microsoft Defender for Identity identified related to the sensor.
type IdentitiesSensorsItemHealthIssuesRequestBuilderGetQueryParameters struct {
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
// IdentitiesSensorsItemHealthIssuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitiesSensorsItemHealthIssuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentitiesSensorsItemHealthIssuesRequestBuilderGetQueryParameters
}
// ByHealthIssueId provides operations to manage the healthIssues property of the microsoft.graph.security.sensor entity.
// returns a *IdentitiesSensorsItemHealthIssuesHealthIssueItemRequestBuilder when successful
func (m *IdentitiesSensorsItemHealthIssuesRequestBuilder) ByHealthIssueId(healthIssueId string)(*IdentitiesSensorsItemHealthIssuesHealthIssueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if healthIssueId != "" {
        urlTplParams["healthIssue%2Did"] = healthIssueId
    }
    return NewIdentitiesSensorsItemHealthIssuesHealthIssueItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIdentitiesSensorsItemHealthIssuesRequestBuilderInternal instantiates a new IdentitiesSensorsItemHealthIssuesRequestBuilder and sets the default values.
func NewIdentitiesSensorsItemHealthIssuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsItemHealthIssuesRequestBuilder) {
    m := &IdentitiesSensorsItemHealthIssuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/identities/sensors/{sensor%2Did}/healthIssues{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIdentitiesSensorsItemHealthIssuesRequestBuilder instantiates a new IdentitiesSensorsItemHealthIssuesRequestBuilder and sets the default values.
func NewIdentitiesSensorsItemHealthIssuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsItemHealthIssuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentitiesSensorsItemHealthIssuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IdentitiesSensorsItemHealthIssuesCountRequestBuilder when successful
func (m *IdentitiesSensorsItemHealthIssuesRequestBuilder) Count()(*IdentitiesSensorsItemHealthIssuesCountRequestBuilder) {
    return NewIdentitiesSensorsItemHealthIssuesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents potential issues within a customer's Microsoft Defender for Identity configuration that Microsoft Defender for Identity identified related to the sensor.
// returns a HealthIssueCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentitiesSensorsItemHealthIssuesRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentitiesSensorsItemHealthIssuesRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HealthIssueCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateHealthIssueCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HealthIssueCollectionResponseable), nil
}
// ToGetRequestInformation represents potential issues within a customer's Microsoft Defender for Identity configuration that Microsoft Defender for Identity identified related to the sensor.
// returns a *RequestInformation when successful
func (m *IdentitiesSensorsItemHealthIssuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentitiesSensorsItemHealthIssuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IdentitiesSensorsItemHealthIssuesRequestBuilder when successful
func (m *IdentitiesSensorsItemHealthIssuesRequestBuilder) WithUrl(rawUrl string)(*IdentitiesSensorsItemHealthIssuesRequestBuilder) {
    return NewIdentitiesSensorsItemHealthIssuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
