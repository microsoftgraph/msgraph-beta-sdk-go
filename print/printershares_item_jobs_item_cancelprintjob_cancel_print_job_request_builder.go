package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder provides operations to call the cancelPrintJob method.
type PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilderInternal instantiates a new PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder and sets the default values.
func NewPrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) {
    m := &PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/jobs/{printJob%2Did}/cancelPrintJob", pathParameters),
    }
    return m
}
// NewPrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder instantiates a new PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder and sets the default values.
func NewPrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action cancelPrintJob
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) Post(ctx context.Context, requestConfiguration *PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action cancelPrintJob
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder when successful
func (m *PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) WithUrl(rawUrl string)(*PrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) {
    return NewPrintersharesItemJobsItemCancelprintjobCancelPrintJobRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
