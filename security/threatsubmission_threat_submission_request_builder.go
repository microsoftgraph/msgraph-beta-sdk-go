package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatsubmissionThreatSubmissionRequestBuilder provides operations to manage the threatSubmission property of the microsoft.graph.security entity.
type ThreatsubmissionThreatSubmissionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatsubmissionThreatSubmissionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatsubmissionThreatSubmissionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatsubmissionThreatSubmissionRequestBuilderGetQueryParameters a threat submission sent to Microsoft; for example, a suspicious email threat, URL threat, or file threat.
type ThreatsubmissionThreatSubmissionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatsubmissionThreatSubmissionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatsubmissionThreatSubmissionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatsubmissionThreatSubmissionRequestBuilderGetQueryParameters
}
// ThreatsubmissionThreatSubmissionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatsubmissionThreatSubmissionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatsubmissionThreatSubmissionRequestBuilderInternal instantiates a new ThreatsubmissionThreatSubmissionRequestBuilder and sets the default values.
func NewThreatsubmissionThreatSubmissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatsubmissionThreatSubmissionRequestBuilder) {
    m := &ThreatsubmissionThreatSubmissionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatSubmission{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatsubmissionThreatSubmissionRequestBuilder instantiates a new ThreatsubmissionThreatSubmissionRequestBuilder and sets the default values.
func NewThreatsubmissionThreatSubmissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatsubmissionThreatSubmissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatsubmissionThreatSubmissionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property threatSubmission for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatsubmissionThreatSubmissionRequestBuilderDeleteRequestConfiguration)(error) {
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
// EmailThreats provides operations to manage the emailThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *ThreatsubmissionEmailthreatsEmailThreatsRequestBuilder when successful
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) EmailThreats()(*ThreatsubmissionEmailthreatsEmailThreatsRequestBuilder) {
    return NewThreatsubmissionEmailthreatsEmailThreatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EmailThreatSubmissionPolicies provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder when successful
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) EmailThreatSubmissionPolicies()(*ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) {
    return NewThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileThreats provides operations to manage the fileThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *ThreatsubmissionFilethreatsFileThreatsRequestBuilder when successful
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) FileThreats()(*ThreatsubmissionFilethreatsFileThreatsRequestBuilder) {
    return NewThreatsubmissionFilethreatsFileThreatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a threat submission sent to Microsoft; for example, a suspicious email threat, URL threat, or file threat.
// returns a ThreatSubmissionRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatsubmissionThreatSubmissionRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatSubmissionRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable), nil
}
// Patch update the navigation property threatSubmission in security
// returns a ThreatSubmissionRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, requestConfiguration *ThreatsubmissionThreatSubmissionRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatSubmissionRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable), nil
}
// ToDeleteRequestInformation delete navigation property threatSubmission for security
// returns a *RequestInformation when successful
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatsubmissionThreatSubmissionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a threat submission sent to Microsoft; for example, a suspicious email threat, URL threat, or file threat.
// returns a *RequestInformation when successful
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatsubmissionThreatSubmissionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property threatSubmission in security
// returns a *RequestInformation when successful
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, requestConfiguration *ThreatsubmissionThreatSubmissionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UrlThreats provides operations to manage the urlThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *ThreatsubmissionUrlthreatsUrlThreatsRequestBuilder when successful
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) UrlThreats()(*ThreatsubmissionUrlthreatsUrlThreatsRequestBuilder) {
    return NewThreatsubmissionUrlthreatsUrlThreatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ThreatsubmissionThreatSubmissionRequestBuilder when successful
func (m *ThreatsubmissionThreatSubmissionRequestBuilder) WithUrl(rawUrl string)(*ThreatsubmissionThreatSubmissionRequestBuilder) {
    return NewThreatsubmissionThreatSubmissionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
