package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/tasktriggers"
    i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/share"
    i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/restorefactorydefaults"
    i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/getcapabilities"
    ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/shares"
    icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/connectors"
    id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/resetdefaults"
    i1b55b9556d3f0f8504c7f2a35c257328ab30ecf9e03df33df4495177129a2cdf "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/shares/item"
    i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/tasktriggers/item"
    ic1c34d114926cc52d44d68cfff4c584108994c1b4377e24401d8c75538512384 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/connectors/item"
)

// PrinterItemRequestBuilder provides operations to manage the printers property of the microsoft.graph.print entity.
type PrinterItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrinterItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrinterItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrinterItemRequestBuilderGetQueryParameters the list of printers registered in the tenant.
type PrinterItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrinterItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrinterItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrinterItemRequestBuilderGetQueryParameters
}
// PrinterItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrinterItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Connectors provides operations to manage the connectors property of the microsoft.graph.printer entity.
func (m *PrinterItemRequestBuilder) Connectors()(*icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a.ConnectorsRequestBuilder) {
    return icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById provides operations to manage the connectors property of the microsoft.graph.printer entity.
func (m *PrinterItemRequestBuilder) ConnectorsById(id string)(*ic1c34d114926cc52d44d68cfff4c584108994c1b4377e24401d8c75538512384.PrintConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printConnector%2Did"] = id
    }
    return ic1c34d114926cc52d44d68cfff4c584108994c1b4377e24401d8c75538512384.NewPrintConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrinterItemRequestBuilderInternal instantiates a new PrinterItemRequestBuilder and sets the default values.
func NewPrinterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterItemRequestBuilder) {
    m := &PrinterItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printers/{printer%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrinterItemRequestBuilder instantiates a new PrinterItemRequestBuilder and sets the default values.
func NewPrinterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property printers for print
func (m *PrinterItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PrinterItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of printers registered in the tenant.
func (m *PrinterItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrinterItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property printers in print
func (m *PrinterItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable, requestConfiguration *PrinterItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property printers for print
func (m *PrinterItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrinterItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of printers registered in the tenant.
func (m *PrinterItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrinterItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrinterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable), nil
}
// GetCapabilities provides operations to call the getCapabilities method.
func (m *PrinterItemRequestBuilder) GetCapabilities()(*i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.GetCapabilitiesRequestBuilder) {
    return i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.NewGetCapabilitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property printers in print
func (m *PrinterItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable, requestConfiguration *PrinterItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrinterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable), nil
}
// ResetDefaults provides operations to call the resetDefaults method.
func (m *PrinterItemRequestBuilder) ResetDefaults()(*id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.ResetDefaultsRequestBuilder) {
    return id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.NewResetDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreFactoryDefaults provides operations to call the restoreFactoryDefaults method.
func (m *PrinterItemRequestBuilder) RestoreFactoryDefaults()(*i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.RestoreFactoryDefaultsRequestBuilder) {
    return i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.NewRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Share provides operations to manage the share property of the microsoft.graph.printer entity.
func (m *PrinterItemRequestBuilder) Share()(*i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.ShareRequestBuilder) {
    return i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shares provides operations to manage the shares property of the microsoft.graph.printer entity.
func (m *PrinterItemRequestBuilder) Shares()(*ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5.SharesRequestBuilder) {
    return ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharesById provides operations to manage the shares property of the microsoft.graph.printer entity.
func (m *PrinterItemRequestBuilder) SharesById(id string)(*i1b55b9556d3f0f8504c7f2a35c257328ab30ecf9e03df33df4495177129a2cdf.PrinterShareItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printerShare%2Did"] = id
    }
    return i1b55b9556d3f0f8504c7f2a35c257328ab30ecf9e03df33df4495177129a2cdf.NewPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskTriggers provides operations to manage the taskTriggers property of the microsoft.graph.printer entity.
func (m *PrinterItemRequestBuilder) TaskTriggers()(*i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb.TaskTriggersRequestBuilder) {
    return i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb.NewTaskTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskTriggersById provides operations to manage the taskTriggers property of the microsoft.graph.printer entity.
func (m *PrinterItemRequestBuilder) TaskTriggersById(id string)(*i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d.PrintTaskTriggerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printTaskTrigger%2Did"] = id
    }
    return i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d.NewPrintTaskTriggerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
