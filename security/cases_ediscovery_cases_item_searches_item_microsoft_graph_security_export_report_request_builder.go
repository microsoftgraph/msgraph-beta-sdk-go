package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder provides operations to call the exportReport method.
type CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilderInternal instantiates a new CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder) {
    m := &CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/microsoft.graph.security.exportReport", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder instantiates a new CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post export an item report from an estimated ediscoverySearch. For details, see Manage a collection estimate.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverysearch-exportreport?view=graph-rest-beta
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder) Post(ctx context.Context, body CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportExportReportPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation export an item report from an estimated ediscoverySearch. For details, see Manage a collection estimate.
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportExportReportPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityExportReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
