package threatsubmission

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
type EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters get emailThreatSubmissionPolicies from threatSubmission
type EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters
}
// EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal instantiates a new EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder and sets the default values.
func NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    m := &EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/threatSubmission/emailThreatSubmissionPolicies/{emailThreatSubmissionPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder instantiates a new EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder and sets the default values.
func NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property emailThreatSubmissionPolicies for threatSubmission
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get emailThreatSubmissionPolicies from threatSubmission
// returns a EmailThreatSubmissionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
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
// Patch update the navigation property emailThreatSubmissionPolicies in threatSubmission
// returns a EmailThreatSubmissionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
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
// ToDeleteRequestInformation delete navigation property emailThreatSubmissionPolicies for threatSubmission
// returns a *RequestInformation when successful
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get emailThreatSubmissionPolicies from threatSubmission
// returns a *RequestInformation when successful
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property emailThreatSubmissionPolicies in threatSubmission
// returns a *RequestInformation when successful
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder when successful
func (m *EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) WithUrl(rawUrl string)(*EmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    return NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
