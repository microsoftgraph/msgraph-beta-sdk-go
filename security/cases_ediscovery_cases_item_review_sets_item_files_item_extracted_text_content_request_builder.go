package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder provides operations to manage the media for the security entity.
type CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderInternal instantiates a new CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/files/{ediscoveryFile%2Did}/extractedTextContent", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder instantiates a new CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get extractedTextContent for the navigation property files from security
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewset-list-files?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put update extractedTextContent for the navigation property files in security
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get extractedTextContent for the navigation property files from security
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update extractedTextContent for the navigation property files in security
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder) {
    return NewCasesEdiscoveryCasesItemReviewSetsItemFilesItemExtractedTextContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
