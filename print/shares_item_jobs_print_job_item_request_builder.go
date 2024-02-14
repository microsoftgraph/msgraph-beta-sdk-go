package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SharesItemJobsPrintJobItemRequestBuilder provides operations to manage the jobs property of the microsoft.graph.printerBase entity.
type SharesItemJobsPrintJobItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SharesItemJobsPrintJobItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharesItemJobsPrintJobItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SharesItemJobsPrintJobItemRequestBuilderGetQueryParameters get jobs from print
type SharesItemJobsPrintJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SharesItemJobsPrintJobItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharesItemJobsPrintJobItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SharesItemJobsPrintJobItemRequestBuilderGetQueryParameters
}
// SharesItemJobsPrintJobItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharesItemJobsPrintJobItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Abort provides operations to call the abort method.
// returns a *SharesItemJobsItemAbortRequestBuilder when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) Abort()(*SharesItemJobsItemAbortRequestBuilder) {
    return NewSharesItemJobsItemAbortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *SharesItemJobsItemCancelRequestBuilder when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) Cancel()(*SharesItemJobsItemCancelRequestBuilder) {
    return NewSharesItemJobsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CancelPrintJob provides operations to call the cancelPrintJob method.
// returns a *SharesItemJobsItemCancelPrintJobRequestBuilder when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) CancelPrintJob()(*SharesItemJobsItemCancelPrintJobRequestBuilder) {
    return NewSharesItemJobsItemCancelPrintJobRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewSharesItemJobsPrintJobItemRequestBuilderInternal instantiates a new SharesItemJobsPrintJobItemRequestBuilder and sets the default values.
func NewSharesItemJobsPrintJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemJobsPrintJobItemRequestBuilder) {
    m := &SharesItemJobsPrintJobItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/shares/{printerShare%2Did}/jobs/{printJob%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSharesItemJobsPrintJobItemRequestBuilder instantiates a new SharesItemJobsPrintJobItemRequestBuilder and sets the default values.
func NewSharesItemJobsPrintJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemJobsPrintJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharesItemJobsPrintJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property jobs for print
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SharesItemJobsPrintJobItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SharesItemJobsPrintJobItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Documents provides operations to manage the documents property of the microsoft.graph.printJob entity.
// returns a *SharesItemJobsItemDocumentsRequestBuilder when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) Documents()(*SharesItemJobsItemDocumentsRequestBuilder) {
    return NewSharesItemJobsItemDocumentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get jobs from print
// returns a PrintJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SharesItemJobsPrintJobItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SharesItemJobsPrintJobItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable), nil
}
// Patch update the navigation property jobs in print
// returns a PrintJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SharesItemJobsPrintJobItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable, requestConfiguration *SharesItemJobsPrintJobItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable), nil
}
// Redirect provides operations to call the redirect method.
// returns a *SharesItemJobsItemRedirectRequestBuilder when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) Redirect()(*SharesItemJobsItemRedirectRequestBuilder) {
    return NewSharesItemJobsItemRedirectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start provides operations to call the start method.
// returns a *SharesItemJobsItemStartRequestBuilder when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) Start()(*SharesItemJobsItemStartRequestBuilder) {
    return NewSharesItemJobsItemStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StartPrintJob provides operations to call the startPrintJob method.
// returns a *SharesItemJobsItemStartPrintJobRequestBuilder when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) StartPrintJob()(*SharesItemJobsItemStartPrintJobRequestBuilder) {
    return NewSharesItemJobsItemStartPrintJobRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.printJob entity.
// returns a *SharesItemJobsItemTasksRequestBuilder when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) Tasks()(*SharesItemJobsItemTasksRequestBuilder) {
    return NewSharesItemJobsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property jobs for print
// returns a *RequestInformation when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SharesItemJobsPrintJobItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/print/shares/{printerShare%2Did}/jobs/{printJob%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get jobs from print
// returns a *RequestInformation when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SharesItemJobsPrintJobItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property jobs in print
// returns a *RequestInformation when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable, requestConfiguration *SharesItemJobsPrintJobItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/print/shares/{printerShare%2Did}/jobs/{printJob%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SharesItemJobsPrintJobItemRequestBuilder when successful
func (m *SharesItemJobsPrintJobItemRequestBuilder) WithUrl(rawUrl string)(*SharesItemJobsPrintJobItemRequestBuilder) {
    return NewSharesItemJobsPrintJobItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
