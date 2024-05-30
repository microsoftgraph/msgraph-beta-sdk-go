package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder provides operations to call the cancelPrintJob method.
type PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilderInternal instantiates a new PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder and sets the default values.
func NewPrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) {
    m := &PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printers/{printer%2Did}/jobs/{printJob%2Did}/cancelPrintJob", pathParameters),
    }
    return m
}
// NewPrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder instantiates a new PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder and sets the default values.
func NewPrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action cancelPrintJob
// Deprecated: The cancelPrintJob API is deprecated and will stop returning data on July 31, 2023. Please use the cancel API instead of this. as of 2023-06/Tasks_And_Plans
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) Post(ctx context.Context, requestConfiguration *PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilderPostRequestConfiguration)(error) {
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
// Deprecated: The cancelPrintJob API is deprecated and will stop returning data on July 31, 2023. Please use the cancel API instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The cancelPrintJob API is deprecated and will stop returning data on July 31, 2023. Please use the cancel API instead of this. as of 2023-06/Tasks_And_Plans
// returns a *PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder when successful
func (m *PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) WithUrl(rawUrl string)(*PrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder) {
    return NewPrintersItemJobsItemCancelprintjobCancelPrintJobRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
