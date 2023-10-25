package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder provides operations to call the reopen method.
type EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilderInternal instantiates a new MicrosoftGraphEdiscoveryReopenRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder) {
    m := &EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/microsoft.graph.ediscovery.reopen", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder instantiates a new MicrosoftGraphEdiscoveryReopenRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reopen an eDiscovery case that was closed. For details, see Reopen a closed case. This API is available in the following national cloud deployments.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-case-reopen?view=graph-rest-1.0
func (m *EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation reopen an eDiscovery case that was closed. For details, see Reopen a closed case. This API is available in the following national cloud deployments.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder) {
    return NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
