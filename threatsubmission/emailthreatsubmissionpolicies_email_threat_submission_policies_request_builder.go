package threatsubmission

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
type EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderGetQueryParameters get emailThreatSubmissionPolicies from threatSubmission
type EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderGetQueryParameters struct {
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
// EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderGetQueryParameters
}
// EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEmailThreatSubmissionPolicyId provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder when successful
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) ByEmailThreatSubmissionPolicyId(emailThreatSubmissionPolicyId string)(*EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if emailThreatSubmissionPolicyId != "" {
        urlTplParams["emailThreatSubmissionPolicy%2Did"] = emailThreatSubmissionPolicyId
    }
    return NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderInternal instantiates a new EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder and sets the default values.
func NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) {
    m := &EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/threatSubmission/emailThreatSubmissionPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder instantiates a new EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder and sets the default values.
func NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EmailthreatsubmissionpoliciesCountRequestBuilder when successful
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) Count()(*EmailthreatsubmissionpoliciesCountRequestBuilder) {
    return NewEmailthreatsubmissionpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get emailThreatSubmissionPolicies from threatSubmission
// returns a EmailThreatSubmissionPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyCollectionResponseable, error) {
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
// Post create new navigation property to emailThreatSubmissionPolicies for threatSubmission
// returns a EmailThreatSubmissionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
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
// ToGetRequestInformation get emailThreatSubmissionPolicies from threatSubmission
// returns a *RequestInformation when successful
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to emailThreatSubmissionPolicies for threatSubmission
// returns a *RequestInformation when successful
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder when successful
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) WithUrl(rawUrl string)(*EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) {
    return NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
