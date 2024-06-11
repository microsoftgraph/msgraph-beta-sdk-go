package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder provides operations to call the activate method.
type EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilderInternal instantiates a new EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder) {
    m := &EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians/{custodian%2Did}/microsoft.graph.ediscovery.activate", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder instantiates a new EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post activate a custodian that was released from a case. This method makes the custodian part of the case again. For details, see Manage custodians in an Advanced eDiscovery case.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-custodian-activate?view=graph-rest-beta
func (m *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation activate a custodian that was released from a case. This method makes the custodian part of the case again. For details, see Manage custodians in an Advanced eDiscovery case.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryactivateMicrosoftGraphEdiscoveryActivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
