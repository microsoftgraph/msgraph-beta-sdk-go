package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder provides operations to call the applyHold method.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    m := &CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/{ediscoveryNoncustodialDataSource%2Did}/microsoft.graph.security.applyHold", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start the process of applying hold on eDiscovery non-custodial data sources. After the operation is created, you can get the status by retrieving the Location parameter from the response headers. The location provides a URL that returns an eDiscoveryHoldOperation object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverynoncustodialdatasource-applyhold?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(error) {
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
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
