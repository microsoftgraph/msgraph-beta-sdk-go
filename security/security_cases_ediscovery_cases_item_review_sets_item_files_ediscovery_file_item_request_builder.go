package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder provides operations to manage the files property of the microsoft.graph.security.ediscoveryReviewSet entity.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderGetQueryParameters represents files within the review set.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderGetQueryParameters
}
// SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderInternal instantiates a new EdiscoveryFileItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) {
    m := &SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/files/{ediscoveryFile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder instantiates a new EdiscoveryFileItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the security entity.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) Content()(*SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemContentRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property files for security
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents files within the review set.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property files in security
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryFileable, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Custodian provides operations to manage the custodian property of the microsoft.graph.security.ediscoveryFile entity.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) Custodian()(*SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemCustodianRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemCustodianRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property files for security
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExtractedTextContent provides operations to manage the media for the security entity.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) ExtractedTextContent()(*SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents files within the review set.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryFileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryFileable), nil
}
// Patch update the navigation property files in security
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryFileable, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryFileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryFileable), nil
}
// Tags provides operations to manage the tags property of the microsoft.graph.security.ediscoveryFile entity.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) Tags()(*SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemTagsRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById provides operations to manage the tags property of the microsoft.graph.security.ediscoveryFile entity.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesEdiscoveryFileItemRequestBuilder) TagsById(id string)(*SecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryReviewTag%2Did"] = id
    }
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
