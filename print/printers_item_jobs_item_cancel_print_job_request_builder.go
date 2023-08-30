package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersItemJobsItemCancelPrintJobRequestBuilder provides operations to call the cancelPrintJob method.
type PrintersItemJobsItemCancelPrintJobRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersItemJobsItemCancelPrintJobRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemJobsItemCancelPrintJobRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrintersItemJobsItemCancelPrintJobRequestBuilderInternal instantiates a new CancelPrintJobRequestBuilder and sets the default values.
func NewPrintersItemJobsItemCancelPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemJobsItemCancelPrintJobRequestBuilder) {
    m := &PrintersItemJobsItemCancelPrintJobRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printers/{printer%2Did}/jobs/{printJob%2Did}/cancelPrintJob", pathParameters),
    }
    return m
}
// NewPrintersItemJobsItemCancelPrintJobRequestBuilder instantiates a new CancelPrintJobRequestBuilder and sets the default values.
func NewPrintersItemJobsItemCancelPrintJobRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemJobsItemCancelPrintJobRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersItemJobsItemCancelPrintJobRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action cancelPrintJob
// Deprecated: The cancelPrintJob API is deprecated and will stop returning data on July 31, 2023. Please use the cancel API instead of this. as of 2023-06/Tasks_And_Plans on 2023-06-13 and will be removed 2023-07-31
func (m *PrintersItemJobsItemCancelPrintJobRequestBuilder) Post(ctx context.Context, requestConfiguration *PrintersItemJobsItemCancelPrintJobRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action cancelPrintJob
// Deprecated: The cancelPrintJob API is deprecated and will stop returning data on July 31, 2023. Please use the cancel API instead of this. as of 2023-06/Tasks_And_Plans on 2023-06-13 and will be removed 2023-07-31
func (m *PrintersItemJobsItemCancelPrintJobRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *PrintersItemJobsItemCancelPrintJobRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The cancelPrintJob API is deprecated and will stop returning data on July 31, 2023. Please use the cancel API instead of this. as of 2023-06/Tasks_And_Plans on 2023-06-13 and will be removed 2023-07-31
func (m *PrintersItemJobsItemCancelPrintJobRequestBuilder) WithUrl(rawUrl string)(*PrintersItemJobsItemCancelPrintJobRequestBuilder) {
    return NewPrintersItemJobsItemCancelPrintJobRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
