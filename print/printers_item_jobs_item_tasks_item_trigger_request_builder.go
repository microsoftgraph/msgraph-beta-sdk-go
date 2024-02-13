package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersItemJobsItemTasksItemTriggerRequestBuilder provides operations to manage the trigger property of the microsoft.graph.printTask entity.
type PrintersItemJobsItemTasksItemTriggerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersItemJobsItemTasksItemTriggerRequestBuilderGetQueryParameters the printTaskTrigger that triggered this task's execution. Read-only.
type PrintersItemJobsItemTasksItemTriggerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrintersItemJobsItemTasksItemTriggerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemJobsItemTasksItemTriggerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintersItemJobsItemTasksItemTriggerRequestBuilderGetQueryParameters
}
// NewPrintersItemJobsItemTasksItemTriggerRequestBuilderInternal instantiates a new PrintersItemJobsItemTasksItemTriggerRequestBuilder and sets the default values.
func NewPrintersItemJobsItemTasksItemTriggerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemJobsItemTasksItemTriggerRequestBuilder) {
    m := &PrintersItemJobsItemTasksItemTriggerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printers/{printer%2Did}/jobs/{printJob%2Did}/tasks/{printTask%2Did}/trigger{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrintersItemJobsItemTasksItemTriggerRequestBuilder instantiates a new PrintersItemJobsItemTasksItemTriggerRequestBuilder and sets the default values.
func NewPrintersItemJobsItemTasksItemTriggerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemJobsItemTasksItemTriggerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersItemJobsItemTasksItemTriggerRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the printTaskTrigger that triggered this task's execution. Read-only.
// returns a PrintTaskTriggerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersItemJobsItemTasksItemTriggerRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintersItemJobsItemTasksItemTriggerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintTaskTriggerable, error) {
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
// returns a *RequestInformation when successful
func (m *PrintersItemJobsItemTasksItemTriggerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrintersItemJobsItemTasksItemTriggerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrintersItemJobsItemTasksItemTriggerRequestBuilder when successful
func (m *PrintersItemJobsItemTasksItemTriggerRequestBuilder) WithUrl(rawUrl string)(*PrintersItemJobsItemTasksItemTriggerRequestBuilder) {
    return NewPrintersItemJobsItemTasksItemTriggerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
