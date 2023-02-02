package threatsubmission

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatSubmissionRequestBuilder provides operations to manage the threatSubmissionRoot singleton.
type ThreatSubmissionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ThreatSubmissionRequestBuilderGetQueryParameters get threatSubmission
type ThreatSubmissionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatSubmissionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatSubmissionRequestBuilderGetQueryParameters
}
// ThreatSubmissionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatSubmissionRequestBuilderInternal instantiates a new ThreatSubmissionRequestBuilder and sets the default values.
func NewThreatSubmissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionRequestBuilder) {
    m := &ThreatSubmissionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/threatSubmission{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewThreatSubmissionRequestBuilder instantiates a new ThreatSubmissionRequestBuilder and sets the default values.
func NewThreatSubmissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatSubmissionRequestBuilderInternal(urlParams, requestAdapter)
}
// EmailThreats provides operations to manage the emailThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
func (m *ThreatSubmissionRequestBuilder) EmailThreats()(*EmailThreatsRequestBuilder) {
    return NewEmailThreatsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EmailThreatsById provides operations to manage the emailThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
func (m *ThreatSubmissionRequestBuilder) EmailThreatsById(id string)(*EmailThreatsEmailThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEmailThreatsEmailThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// EmailThreatSubmissionPolicies provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
func (m *ThreatSubmissionRequestBuilder) EmailThreatSubmissionPolicies()(*EmailThreatSubmissionPoliciesRequestBuilder) {
    return NewEmailThreatSubmissionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EmailThreatSubmissionPoliciesById provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
func (m *ThreatSubmissionRequestBuilder) EmailThreatSubmissionPoliciesById(id string)(*EmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// FileThreats provides operations to manage the fileThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
func (m *ThreatSubmissionRequestBuilder) FileThreats()(*FileThreatsRequestBuilder) {
    return NewFileThreatsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FileThreatsById provides operations to manage the fileThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
func (m *ThreatSubmissionRequestBuilder) FileThreatsById(id string)(*FileThreatsFileThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewFileThreatsFileThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Get get threatSubmission
func (m *ThreatSubmissionRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatSubmissionRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatSubmissionRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable), nil
}
// Patch update threatSubmission
func (m *ThreatSubmissionRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, requestConfiguration *ThreatSubmissionRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatSubmissionRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable), nil
}
// ToGetRequestInformation get threatSubmission
func (m *ThreatSubmissionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatSubmissionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update threatSubmission
func (m *ThreatSubmissionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, requestConfiguration *ThreatSubmissionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UrlThreats provides operations to manage the urlThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
func (m *ThreatSubmissionRequestBuilder) UrlThreats()(*UrlThreatsRequestBuilder) {
    return NewUrlThreatsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UrlThreatsById provides operations to manage the urlThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
func (m *ThreatSubmissionRequestBuilder) UrlThreatsById(id string)(*UrlThreatsUrlThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUrlThreatsUrlThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
