package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersItemConnectorsRequestBuilder provides operations to manage the connectors property of the microsoft.graph.printer entity.
type PrintersItemConnectorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersItemConnectorsRequestBuilderGetQueryParameters the connectors that are associated with the printer.
type PrintersItemConnectorsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PrintersItemConnectorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemConnectorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintersItemConnectorsRequestBuilderGetQueryParameters
}
// ByPrintConnectorId provides operations to manage the connectors property of the microsoft.graph.printer entity.
// returns a *PrintersItemConnectorsPrintConnectorItemRequestBuilder when successful
func (m *PrintersItemConnectorsRequestBuilder) ByPrintConnectorId(printConnectorId string)(*PrintersItemConnectorsPrintConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if printConnectorId != "" {
        urlTplParams["printConnector%2Did"] = printConnectorId
    }
    return NewPrintersItemConnectorsPrintConnectorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrintersItemConnectorsRequestBuilderInternal instantiates a new PrintersItemConnectorsRequestBuilder and sets the default values.
func NewPrintersItemConnectorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemConnectorsRequestBuilder) {
    m := &PrintersItemConnectorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printers/{printer%2Did}/connectors{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrintersItemConnectorsRequestBuilder instantiates a new PrintersItemConnectorsRequestBuilder and sets the default values.
func NewPrintersItemConnectorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemConnectorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersItemConnectorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrintersItemConnectorsCountRequestBuilder when successful
func (m *PrintersItemConnectorsRequestBuilder) Count()(*PrintersItemConnectorsCountRequestBuilder) {
    return NewPrintersItemConnectorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the connectors that are associated with the printer.
// returns a PrintConnectorCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersItemConnectorsRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintersItemConnectorsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintConnectorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintConnectorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintConnectorCollectionResponseable), nil
}
// ToGetRequestInformation the connectors that are associated with the printer.
// returns a *RequestInformation when successful
func (m *PrintersItemConnectorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrintersItemConnectorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrintersItemConnectorsRequestBuilder when successful
func (m *PrintersItemConnectorsRequestBuilder) WithUrl(rawUrl string)(*PrintersItemConnectorsRequestBuilder) {
    return NewPrintersItemConnectorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
