package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder provides operations to manage the userSources property of the microsoft.graph.security.ediscoveryHoldPolicy entity.
type SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderGetQueryParameters data sources that represent Exchange mailboxes.
type SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderGetQueryParameters
}
// SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderInternal instantiates a new UserSourceItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) {
    m := &SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/legalHolds/{ediscoveryHoldPolicy%2Did}/userSources/{userSource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder instantiates a new UserSourceItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userSources for security
func (m *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation data sources that represent Exchange mailboxes.
func (m *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userSources in security
func (m *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable, requestConfiguration *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property userSources for security
func (m *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get data sources that represent Exchange mailboxes.
func (m *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateUserSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable), nil
}
// Patch update the navigation property userSources in security
func (m *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable, requestConfiguration *SecurityCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateUserSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable), nil
}
