package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatsubmissionFilethreatsFileThreatsRequestBuilder provides operations to manage the fileThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
type ThreatsubmissionFilethreatsFileThreatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatsubmissionFilethreatsFileThreatsRequestBuilderGetQueryParameters get a list of the fileThreatSubmission objects and their properties.
type ThreatsubmissionFilethreatsFileThreatsRequestBuilderGetQueryParameters struct {
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
// ThreatsubmissionFilethreatsFileThreatsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatsubmissionFilethreatsFileThreatsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatsubmissionFilethreatsFileThreatsRequestBuilderGetQueryParameters
}
// ThreatsubmissionFilethreatsFileThreatsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatsubmissionFilethreatsFileThreatsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByFileThreatSubmissionId provides operations to manage the fileThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *ThreatsubmissionFilethreatsFileThreatSubmissionItemRequestBuilder when successful
func (m *ThreatsubmissionFilethreatsFileThreatsRequestBuilder) ByFileThreatSubmissionId(fileThreatSubmissionId string)(*ThreatsubmissionFilethreatsFileThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if fileThreatSubmissionId != "" {
        urlTplParams["fileThreatSubmission%2Did"] = fileThreatSubmissionId
    }
    return NewThreatsubmissionFilethreatsFileThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatsubmissionFilethreatsFileThreatsRequestBuilderInternal instantiates a new ThreatsubmissionFilethreatsFileThreatsRequestBuilder and sets the default values.
func NewThreatsubmissionFilethreatsFileThreatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatsubmissionFilethreatsFileThreatsRequestBuilder) {
    m := &ThreatsubmissionFilethreatsFileThreatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatSubmission/fileThreats{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatsubmissionFilethreatsFileThreatsRequestBuilder instantiates a new ThreatsubmissionFilethreatsFileThreatsRequestBuilder and sets the default values.
func NewThreatsubmissionFilethreatsFileThreatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatsubmissionFilethreatsFileThreatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatsubmissionFilethreatsFileThreatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatsubmissionFilethreatsCountRequestBuilder when successful
func (m *ThreatsubmissionFilethreatsFileThreatsRequestBuilder) Count()(*ThreatsubmissionFilethreatsCountRequestBuilder) {
    return NewThreatsubmissionFilethreatsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the fileThreatSubmission objects and their properties.
// returns a FileThreatSubmissionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-filethreatsubmission-list?view=graph-rest-beta
func (m *ThreatsubmissionFilethreatsFileThreatsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatsubmissionFilethreatsFileThreatsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FileThreatSubmissionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateFileThreatSubmissionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FileThreatSubmissionCollectionResponseable), nil
}
// Post create new navigation property to fileThreats for security
// returns a FileThreatSubmissionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatsubmissionFilethreatsFileThreatsRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FileThreatSubmissionable, requestConfiguration *ThreatsubmissionFilethreatsFileThreatsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FileThreatSubmissionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateFileThreatSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FileThreatSubmissionable), nil
}
// ToGetRequestInformation get a list of the fileThreatSubmission objects and their properties.
// returns a *RequestInformation when successful
func (m *ThreatsubmissionFilethreatsFileThreatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatsubmissionFilethreatsFileThreatsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to fileThreats for security
// returns a *RequestInformation when successful
func (m *ThreatsubmissionFilethreatsFileThreatsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FileThreatSubmissionable, requestConfiguration *ThreatsubmissionFilethreatsFileThreatsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatsubmissionFilethreatsFileThreatsRequestBuilder when successful
func (m *ThreatsubmissionFilethreatsFileThreatsRequestBuilder) WithUrl(rawUrl string)(*ThreatsubmissionFilethreatsFileThreatsRequestBuilder) {
    return NewThreatsubmissionFilethreatsFileThreatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
