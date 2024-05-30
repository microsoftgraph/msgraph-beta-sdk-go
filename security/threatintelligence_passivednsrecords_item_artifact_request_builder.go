package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder provides operations to manage the artifact property of the microsoft.graph.security.passiveDnsRecord entity.
type ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligencePassivednsrecordsItemArtifactRequestBuilderGetQueryParameters the artifact related to this passiveDnsRecord entry.
type ThreatintelligencePassivednsrecordsItemArtifactRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligencePassivednsrecordsItemArtifactRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligencePassivednsrecordsItemArtifactRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligencePassivednsrecordsItemArtifactRequestBuilderGetQueryParameters
}
// NewThreatintelligencePassivednsrecordsItemArtifactRequestBuilderInternal instantiates a new ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder and sets the default values.
func NewThreatintelligencePassivednsrecordsItemArtifactRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder) {
    m := &ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/passiveDnsRecords/{passiveDnsRecord%2Did}/artifact{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligencePassivednsrecordsItemArtifactRequestBuilder instantiates a new ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder and sets the default values.
func NewThreatintelligencePassivednsrecordsItemArtifactRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligencePassivednsrecordsItemArtifactRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the artifact related to this passiveDnsRecord entry.
// returns a Artifactable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligencePassivednsrecordsItemArtifactRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Artifactable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateArtifactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Artifactable), nil
}
// ToGetRequestInformation the artifact related to this passiveDnsRecord entry.
// returns a *RequestInformation when successful
func (m *ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligencePassivednsrecordsItemArtifactRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder when successful
func (m *ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder) {
    return NewThreatintelligencePassivednsrecordsItemArtifactRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
