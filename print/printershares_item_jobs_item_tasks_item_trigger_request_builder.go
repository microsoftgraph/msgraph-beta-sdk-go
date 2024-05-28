package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersharesItemJobsItemTasksItemTriggerRequestBuilder provides operations to manage the trigger property of the microsoft.graph.printTask entity.
type PrintersharesItemJobsItemTasksItemTriggerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersharesItemJobsItemTasksItemTriggerRequestBuilderGetQueryParameters the printTaskTrigger that triggered this task's execution. Read-only.
type PrintersharesItemJobsItemTasksItemTriggerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrintersharesItemJobsItemTasksItemTriggerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersharesItemJobsItemTasksItemTriggerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintersharesItemJobsItemTasksItemTriggerRequestBuilderGetQueryParameters
}
// NewPrintersharesItemJobsItemTasksItemTriggerRequestBuilderInternal instantiates a new PrintersharesItemJobsItemTasksItemTriggerRequestBuilder and sets the default values.
func NewPrintersharesItemJobsItemTasksItemTriggerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemJobsItemTasksItemTriggerRequestBuilder) {
    m := &PrintersharesItemJobsItemTasksItemTriggerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/jobs/{printJob%2Did}/tasks/{printTask%2Did}/trigger{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrintersharesItemJobsItemTasksItemTriggerRequestBuilder instantiates a new PrintersharesItemJobsItemTasksItemTriggerRequestBuilder and sets the default values.
func NewPrintersharesItemJobsItemTasksItemTriggerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemJobsItemTasksItemTriggerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersharesItemJobsItemTasksItemTriggerRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the printTaskTrigger that triggered this task's execution. Read-only.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrintTaskTriggerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersharesItemJobsItemTasksItemTriggerRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintersharesItemJobsItemTasksItemTriggerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintTaskTriggerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintTaskTriggerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintTaskTriggerable), nil
}
// ToGetRequestInformation the printTaskTrigger that triggered this task's execution. Read-only.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *PrintersharesItemJobsItemTasksItemTriggerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrintersharesItemJobsItemTasksItemTriggerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *PrintersharesItemJobsItemTasksItemTriggerRequestBuilder when successful
func (m *PrintersharesItemJobsItemTasksItemTriggerRequestBuilder) WithUrl(rawUrl string)(*PrintersharesItemJobsItemTasksItemTriggerRequestBuilder) {
    return NewPrintersharesItemJobsItemTasksItemTriggerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
