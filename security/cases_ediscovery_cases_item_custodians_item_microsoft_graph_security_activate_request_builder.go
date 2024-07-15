package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder provides operations to call the activate method.
type CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilderInternal instantiates a new CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder) {
    m := &CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/microsoft.graph.security.activate", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder instantiates a new CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post activate a custodian that has been released from a case to make them part of the case again. For details, see Manage custodians in an eDiscovery (Premium) case.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-activate?view=graph-rest-beta
func (m *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation activate a custodian that has been released from a case to make them part of the case again. For details, see Manage custodians in an eDiscovery (Premium) case.
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
