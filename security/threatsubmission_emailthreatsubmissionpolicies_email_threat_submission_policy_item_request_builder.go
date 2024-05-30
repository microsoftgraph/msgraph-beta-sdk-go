package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
type ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters read the properties and relationships of an emailThreatSubmissionPolicy object.
type ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters
}
// ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal instantiates a new ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder and sets the default values.
func NewThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    m := &ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatSubmission/emailThreatSubmissionPolicies/{emailThreatSubmissionPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder instantiates a new ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder and sets the default values.
func NewThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property emailThreatSubmissionPolicies for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an emailThreatSubmissionPolicy object.
// returns a EmailThreatSubmissionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-emailthreatsubmissionpolicy-get?view=graph-rest-beta
func (m *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property emailThreatSubmissionPolicies in security
// returns a EmailThreatSubmissionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property emailThreatSubmissionPolicies for security
// returns a *RequestInformation when successful
func (m *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an emailThreatSubmissionPolicy object.
// returns a *RequestInformation when successful
func (m *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property emailThreatSubmissionPolicies in security
// returns a *RequestInformation when successful
func (m *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder when successful
func (m *ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    return NewThreatsubmissionEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
