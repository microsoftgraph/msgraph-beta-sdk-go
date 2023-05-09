package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder provides operations to call the applyHold method.
type EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilderInternal instantiates a new MicrosoftGraphEdiscoveryApplyHoldRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) {
    m := &EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/noncustodialDataSources/{noncustodialDataSource%2Did}/microsoft.graph.ediscovery.applyHold", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder instantiates a new MicrosoftGraphEdiscoveryApplyHoldRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyHold
func (m *EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action applyHold
func (m *EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
