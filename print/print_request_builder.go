package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintRequestBuilder provides operations to manage the print singleton.
type PrintRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrintRequestBuilderGetQueryParameters get print
type PrintRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrintRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintRequestBuilderGetQueryParameters
}
// PrintRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Connectors provides operations to manage the connectors property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Connectors()(*ConnectorsRequestBuilder) {
    return NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById provides operations to manage the connectors property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) ConnectorsById(id string)(*ConnectorsPrintConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printConnector%2Did"] = id
    }
    return NewConnectorsPrintConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrintRequestBuilderInternal instantiates a new PrintRequestBuilder and sets the default values.
func NewPrintRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintRequestBuilder) {
    m := &PrintRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintRequestBuilder instantiates a new PrintRequestBuilder and sets the default values.
func NewPrintRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get print
func (m *PrintRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrintRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// CreatePatchRequestInformation update print
func (m *PrintRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable, requestConfiguration *PrintRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get print
func (m *PrintRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable), nil
}
// Operations provides operations to manage the operations property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Operations()(*OperationsRequestBuilder) {
    return NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) OperationsById(id string)(*OperationsPrintOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printOperation%2Did"] = id
    }
    return NewOperationsPrintOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update print
func (m *PrintRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable, requestConfiguration *PrintRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable), nil
}
// Printers provides operations to manage the printers property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Printers()(*PrintersRequestBuilder) {
    return NewPrintersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrintersById provides operations to manage the printers property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) PrintersById(id string)(*PrintersPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printer%2Did"] = id
    }
    return NewPrintersPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PrinterShares provides operations to manage the printerShares property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) PrinterShares()(*PrinterSharesRequestBuilder) {
    return NewPrinterSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrinterSharesById provides operations to manage the printerShares property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) PrinterSharesById(id string)(*PrinterSharesPrinterShareItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printerShare%2Did"] = id
    }
    return NewPrinterSharesPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reports provides operations to manage the reports property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Reports()(*ReportsRequestBuilder) {
    return NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Services provides operations to manage the services property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Services()(*ServicesRequestBuilder) {
    return NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicesById provides operations to manage the services property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) ServicesById(id string)(*ServicesPrintServiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printService%2Did"] = id
    }
    return NewServicesPrintServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Shares provides operations to manage the shares property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Shares()(*SharesRequestBuilder) {
    return NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharesById provides operations to manage the shares property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) SharesById(id string)(*SharesPrinterShareItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printerShare%2Did"] = id
    }
    return NewSharesPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskDefinitions provides operations to manage the taskDefinitions property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) TaskDefinitions()(*TaskDefinitionsRequestBuilder) {
    return NewTaskDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskDefinitionsById provides operations to manage the taskDefinitions property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) TaskDefinitionsById(id string)(*TaskDefinitionsPrintTaskDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printTaskDefinition%2Did"] = id
    }
    return NewTaskDefinitionsPrintTaskDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
