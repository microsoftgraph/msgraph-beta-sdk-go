package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
type SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters get emailThreatSubmissionPolicies from security
type SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetQueryParameters
}
// SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal instantiates a new EmailThreatSubmissionPolicyItemRequestBuilder and sets the default values.
func NewSecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    m := &SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/threatSubmission/emailThreatSubmissionPolicies/{emailThreatSubmissionPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder instantiates a new EmailThreatSubmissionPolicyItemRequestBuilder and sets the default values.
func NewSecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property emailThreatSubmissionPolicies for security
func (m *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get emailThreatSubmissionPolicies from security
func (m *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property emailThreatSubmissionPolicies in security
func (m *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property emailThreatSubmissionPolicies for security
func (m *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get emailThreatSubmissionPolicies from security
func (m *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEmailThreatSubmissionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable), nil
}
// Patch update the navigation property emailThreatSubmissionPolicies in security
func (m *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, requestConfiguration *SecurityThreatSubmissionEmailThreatSubmissionPoliciesEmailThreatSubmissionPolicyItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEmailThreatSubmissionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EmailThreatSubmissionPolicyable), nil
}
