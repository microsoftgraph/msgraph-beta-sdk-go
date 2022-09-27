package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i0d270c48a37eef745b51dd39fcd6e6c396a49d6e354a0a22c4b38a373b1f8649 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2a6efa35f1da64a13760e57fe6078241cde1fc0716fe6cf741b1913916af2a06 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares"
    i32ff24c6864832ad0f41a0c6482a3f153c54ab6eec76bf30bcdc4a10beb817fc "github.com/microsoftgraph/msgraph-beta-sdk-go/print/services"
    i4c4b7dddb590dec724ed827d1d5c27e935dc2a75d5fc9b1ddb47f1cc82d2f70a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/taskdefinitions"
    i70c2e2fc495120f72326108a5af86cde4684b679d130a86acdd4046390fee262 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports"
    i844ad4158981a129d54aa4a317c17dc8b56ba989d7b7cfa3ffe6aebbb129c245 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/operations"
    i92d1f50b33fa740ede98125a2225a30e704277949de6d5ba06e1e4f209889a99 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers"
    ie0ee21308c5a5876efdcb565ab54dddf497480052514698aa5df8076282afb54 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/connectors"
    i041fa93b1b8c4c6acaa44fa25a4bab5b60619191120dc3943102fdec799fca39 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item"
    i20c82299a583bf997be8654a60766afa3b225e0e8b54ab661cb734c46c2add9b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/operations/item"
    i21ad7a37923a403c8172690e362e15a8a0298e4a82eecbedee9e691ff99bdde3 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item"
    i5ec86a089dc848e3bf34cce2d116379624ef91e1c58d8a32e664e6fb5397c958 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/taskdefinitions/item"
    i6ac2f8e4588dd2304d6a10565c050a97a9cab528db29bf610191796b5477b498 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item"
    i9508344ddcc1d2a6a0ed36e7aeb07d60730f841e1e8d4674d8ae4d356d4932a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/services/item"
    id93d2fbcff2a13af1a3388fed5faf01351e12efec3ce632d2152e37f5ad29d99 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/connectors/item"
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
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintRequestBuilderGetQueryParameters
}
// PrintRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Connectors the connectors property
func (m *PrintRequestBuilder) Connectors()(*ie0ee21308c5a5876efdcb565ab54dddf497480052514698aa5df8076282afb54.ConnectorsRequestBuilder) {
    return ie0ee21308c5a5876efdcb565ab54dddf497480052514698aa5df8076282afb54.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.connectors.item collection
func (m *PrintRequestBuilder) ConnectorsById(id string)(*id93d2fbcff2a13af1a3388fed5faf01351e12efec3ce632d2152e37f5ad29d99.PrintConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printConnector%2Did"] = id
    }
    return id93d2fbcff2a13af1a3388fed5faf01351e12efec3ce632d2152e37f5ad29d99.NewPrintConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *PrintRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get print
func (m *PrintRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PrintRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update print
func (m *PrintRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update print
func (m *PrintRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable, requestConfiguration *PrintRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get print
func (m *PrintRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Operations the operations property
func (m *PrintRequestBuilder) Operations()(*i844ad4158981a129d54aa4a317c17dc8b56ba989d7b7cfa3ffe6aebbb129c245.OperationsRequestBuilder) {
    return i844ad4158981a129d54aa4a317c17dc8b56ba989d7b7cfa3ffe6aebbb129c245.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.operations.item collection
func (m *PrintRequestBuilder) OperationsById(id string)(*i20c82299a583bf997be8654a60766afa3b225e0e8b54ab661cb734c46c2add9b.PrintOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printOperation%2Did"] = id
    }
    return i20c82299a583bf997be8654a60766afa3b225e0e8b54ab661cb734c46c2add9b.NewPrintOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update print
func (m *PrintRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable, requestConfiguration *PrintRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Printers the printers property
func (m *PrintRequestBuilder) Printers()(*i92d1f50b33fa740ede98125a2225a30e704277949de6d5ba06e1e4f209889a99.PrintersRequestBuilder) {
    return i92d1f50b33fa740ede98125a2225a30e704277949de6d5ba06e1e4f209889a99.NewPrintersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrintersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printers.item collection
func (m *PrintRequestBuilder) PrintersById(id string)(*i21ad7a37923a403c8172690e362e15a8a0298e4a82eecbedee9e691ff99bdde3.PrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printer%2Did"] = id
    }
    return i21ad7a37923a403c8172690e362e15a8a0298e4a82eecbedee9e691ff99bdde3.NewPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PrinterShares the printerShares property
func (m *PrintRequestBuilder) PrinterShares()(*i2a6efa35f1da64a13760e57fe6078241cde1fc0716fe6cf741b1913916af2a06.PrinterSharesRequestBuilder) {
    return i2a6efa35f1da64a13760e57fe6078241cde1fc0716fe6cf741b1913916af2a06.NewPrinterSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrinterSharesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printerShares.item collection
func (m *PrintRequestBuilder) PrinterSharesById(id string)(*i6ac2f8e4588dd2304d6a10565c050a97a9cab528db29bf610191796b5477b498.PrinterShareItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printerShare%2Did"] = id
    }
    return i6ac2f8e4588dd2304d6a10565c050a97a9cab528db29bf610191796b5477b498.NewPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reports the reports property
func (m *PrintRequestBuilder) Reports()(*i70c2e2fc495120f72326108a5af86cde4684b679d130a86acdd4046390fee262.ReportsRequestBuilder) {
    return i70c2e2fc495120f72326108a5af86cde4684b679d130a86acdd4046390fee262.NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Services the services property
func (m *PrintRequestBuilder) Services()(*i32ff24c6864832ad0f41a0c6482a3f153c54ab6eec76bf30bcdc4a10beb817fc.ServicesRequestBuilder) {
    return i32ff24c6864832ad0f41a0c6482a3f153c54ab6eec76bf30bcdc4a10beb817fc.NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.services.item collection
func (m *PrintRequestBuilder) ServicesById(id string)(*i9508344ddcc1d2a6a0ed36e7aeb07d60730f841e1e8d4674d8ae4d356d4932a9.PrintServiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printService%2Did"] = id
    }
    return i9508344ddcc1d2a6a0ed36e7aeb07d60730f841e1e8d4674d8ae4d356d4932a9.NewPrintServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Shares the shares property
func (m *PrintRequestBuilder) Shares()(*i0d270c48a37eef745b51dd39fcd6e6c396a49d6e354a0a22c4b38a373b1f8649.SharesRequestBuilder) {
    return i0d270c48a37eef745b51dd39fcd6e6c396a49d6e354a0a22c4b38a373b1f8649.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.shares.item collection
func (m *PrintRequestBuilder) SharesById(id string)(*i041fa93b1b8c4c6acaa44fa25a4bab5b60619191120dc3943102fdec799fca39.PrinterShareItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printerShare%2Did"] = id
    }
    return i041fa93b1b8c4c6acaa44fa25a4bab5b60619191120dc3943102fdec799fca39.NewPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskDefinitions the taskDefinitions property
func (m *PrintRequestBuilder) TaskDefinitions()(*i4c4b7dddb590dec724ed827d1d5c27e935dc2a75d5fc9b1ddb47f1cc82d2f70a.TaskDefinitionsRequestBuilder) {
    return i4c4b7dddb590dec724ed827d1d5c27e935dc2a75d5fc9b1ddb47f1cc82d2f70a.NewTaskDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.taskDefinitions.item collection
func (m *PrintRequestBuilder) TaskDefinitionsById(id string)(*i5ec86a089dc848e3bf34cce2d116379624ef91e1c58d8a32e664e6fb5397c958.PrintTaskDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printTaskDefinition%2Did"] = id
    }
    return i5ec86a089dc848e3bf34cce2d116379624ef91e1c58d8a32e664e6fb5397c958.NewPrintTaskDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
