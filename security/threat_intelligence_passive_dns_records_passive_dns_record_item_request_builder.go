package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder provides operations to manage the passiveDnsRecords property of the microsoft.graph.security.threatIntelligence entity.
type ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderGetQueryParameters read the properties and relationships of a passiveDnsRecord object.
type ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderGetQueryParameters
}
// ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Artifact provides operations to manage the artifact property of the microsoft.graph.security.passiveDnsRecord entity.
// returns a *ThreatIntelligencePassiveDnsRecordsItemArtifactRequestBuilder when successful
func (m *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) Artifact()(*ThreatIntelligencePassiveDnsRecordsItemArtifactRequestBuilder) {
    return NewThreatIntelligencePassiveDnsRecordsItemArtifactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderInternal instantiates a new ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder and sets the default values.
func NewThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) {
    m := &ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/passiveDnsRecords/{passiveDnsRecord%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder instantiates a new ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder and sets the default values.
func NewThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property passiveDnsRecords for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of a passiveDnsRecord object.
// returns a PassiveDnsRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-passivednsrecord-get?view=graph-rest-1.0
func (m *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PassiveDnsRecordable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreatePassiveDnsRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PassiveDnsRecordable), nil
}
// ParentHost provides operations to manage the parentHost property of the microsoft.graph.security.passiveDnsRecord entity.
// returns a *ThreatIntelligencePassiveDnsRecordsItemParentHostRequestBuilder when successful
func (m *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) ParentHost()(*ThreatIntelligencePassiveDnsRecordsItemParentHostRequestBuilder) {
    return NewThreatIntelligencePassiveDnsRecordsItemParentHostRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property passiveDnsRecords in security
// returns a PassiveDnsRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PassiveDnsRecordable, requestConfiguration *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PassiveDnsRecordable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreatePassiveDnsRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PassiveDnsRecordable), nil
}
// ToDeleteRequestInformation delete navigation property passiveDnsRecords for security
// returns a *RequestInformation when successful
func (m *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/security/threatIntelligence/passiveDnsRecords/{passiveDnsRecord%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a passiveDnsRecord object.
// returns a *RequestInformation when successful
func (m *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property passiveDnsRecords in security
// returns a *RequestInformation when successful
func (m *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PassiveDnsRecordable, requestConfiguration *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/security/threatIntelligence/passiveDnsRecords/{passiveDnsRecord%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder when successful
func (m *ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) WithUrl(rawUrl string)(*ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) {
    return NewThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
