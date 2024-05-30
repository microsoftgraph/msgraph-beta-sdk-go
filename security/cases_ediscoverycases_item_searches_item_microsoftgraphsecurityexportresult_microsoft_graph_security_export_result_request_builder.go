package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder provides operations to call the exportResult method.
type CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder) {
    m := &CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/microsoft.graph.security.exportResult", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder instantiates a new CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action exportResult
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder) Post(ctx context.Context, body CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action exportResult
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultMicrosoftGraphSecurityExportResultRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
