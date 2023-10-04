package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersItemTaskTriggersItemDefinitionRequestBuilder provides operations to manage the definition property of the microsoft.graph.printTaskTrigger entity.
type PrintersItemTaskTriggersItemDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersItemTaskTriggersItemDefinitionRequestBuilderGetQueryParameters an abstract definition that is used to create a printTask when triggered by a print event. Read-only.
type PrintersItemTaskTriggersItemDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrintersItemTaskTriggersItemDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemTaskTriggersItemDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintersItemTaskTriggersItemDefinitionRequestBuilderGetQueryParameters
}
// NewPrintersItemTaskTriggersItemDefinitionRequestBuilderInternal instantiates a new DefinitionRequestBuilder and sets the default values.
func NewPrintersItemTaskTriggersItemDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemTaskTriggersItemDefinitionRequestBuilder) {
    m := &PrintersItemTaskTriggersItemDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printers/{printer%2Did}/taskTriggers/{printTaskTrigger%2Did}/definition{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewPrintersItemTaskTriggersItemDefinitionRequestBuilder instantiates a new DefinitionRequestBuilder and sets the default values.
func NewPrintersItemTaskTriggersItemDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemTaskTriggersItemDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersItemTaskTriggersItemDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get an abstract definition that is used to create a printTask when triggered by a print event. Read-only.
func (m *PrintersItemTaskTriggersItemDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintersItemTaskTriggersItemDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintTaskDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintTaskDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintTaskDefinitionable), nil
}
// ToGetRequestInformation an abstract definition that is used to create a printTask when triggered by a print event. Read-only.
func (m *PrintersItemTaskTriggersItemDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrintersItemTaskTriggersItemDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PrintersItemTaskTriggersItemDefinitionRequestBuilder) WithUrl(rawUrl string)(*PrintersItemTaskTriggersItemDefinitionRequestBuilder) {
    return NewPrintersItemTaskTriggersItemDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
