// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder provides operations to call the applyHold method.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilderInternal instantiates a new CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    m := &CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/{ediscoveryNoncustodialDataSource%2Did}/microsoft.graph.security.applyHold", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder instantiates a new CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start the process of applying hold on eDiscovery non-custodial data sources. After the operation is created, you can get the status by retrieving the Location parameter from the response headers. The location provides a URL that returns an eDiscoveryHoldOperation object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverynoncustodialdatasource-applyhold?view=graph-rest-beta
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation start the process of applying hold on eDiscovery non-custodial data sources. After the operation is created, you can get the status by retrieving the Location parameter from the response headers. The location provides a URL that returns an eDiscoveryHoldOperation object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
