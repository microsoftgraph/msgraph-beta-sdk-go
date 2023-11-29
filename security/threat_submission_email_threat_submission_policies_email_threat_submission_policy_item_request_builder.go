package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
type ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters read the properties and relationships of an emailThreatSubmissionPolicy object.
type ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters
}
// ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal instantiates a new EmailThreatSubmissionPolicyItemRequestBuilder and sets the default values.
func NewThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    m := &ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatSubmission/emailThreatSubmissionPolicies/{emailThreatSubmissionPolicy%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder instantiates a new EmailThreatSubmissionPolicyItemRequestBuilder and sets the default values.
func NewThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property emailThreatSubmissionPolicies for security
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of an emailThreatSubmissionPolicy object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-emailthreatsubmissionpolicy-get?view=graph-rest-1.0
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an emailThreatSubmissionPolicy object.
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    return NewThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
