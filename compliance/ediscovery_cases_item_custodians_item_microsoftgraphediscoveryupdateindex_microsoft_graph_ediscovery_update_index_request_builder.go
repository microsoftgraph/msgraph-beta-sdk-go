package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder provides operations to call the updateIndex method.
type EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderInternal instantiates a new EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) {
    m := &EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians/{custodian%2Did}/microsoft.graph.ediscovery.updateIndex", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder instantiates a new EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateIndex
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action updateIndex
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
