package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder provides operations to call the applyHold method.
type EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilderInternal instantiates a new EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) {
    m := &EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/noncustodialDataSources/{noncustodialDataSource%2Did}/microsoft.graph.ediscovery.applyHold", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder instantiates a new EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyHold
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action applyHold
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
