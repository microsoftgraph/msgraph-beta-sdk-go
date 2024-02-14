package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
type ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderGetQueryParameters get a list of the emailThreatSubmissionPolicy objects and their properties.
type ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderGetQueryParameters struct {
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
// ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderGetQueryParameters
}
// ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEmailThreatSubmissionPolicyId provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder when successful
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) ByEmailThreatSubmissionPolicyId(emailThreatSubmissionPolicyId string)(*ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if emailThreatSubmissionPolicyId != "" {
        urlTplParams["emailThreatSubmissionPolicy%2Did"] = emailThreatSubmissionPolicyId
    }
    return NewThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderInternal instantiates a new ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder and sets the default values.
func NewThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) {
    m := &ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatSubmission/emailThreatSubmissionPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder instantiates a new ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder and sets the default values.
func NewThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatSubmissionEmailThreatSubmissionPoliciesCountRequestBuilder when successful
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) Count()(*ThreatSubmissionEmailThreatSubmissionPoliciesCountRequestBuilder) {
    return NewThreatSubmissionEmailThreatSubmissionPoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the emailThreatSubmissionPolicy objects and their properties.
// returns a EmailThreatSubmissionPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-emailthreatsubmissionpolicy-list?view=graph-rest-1.0
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEmailThreatSubmissionPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyCollectionResponseable), nil
}
// Post create new navigation property to emailThreatSubmissionPolicies for security
// returns a EmailThreatSubmissionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEmailThreatSubmissionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable), nil
}
// ToGetRequestInformation get a list of the emailThreatSubmissionPolicy objects and their properties.
// returns a *RequestInformation when successful
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to emailThreatSubmissionPolicies for security
// returns a *RequestInformation when successful
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/security/threatSubmission/emailThreatSubmissionPolicies", m.BaseRequestBuilder.PathParameters)
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
// returns a *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder when successful
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) WithUrl(rawUrl string)(*ThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder) {
    return NewThreatSubmissionEmailThreatSubmissionPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
