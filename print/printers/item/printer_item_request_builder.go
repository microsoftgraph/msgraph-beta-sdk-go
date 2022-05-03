package item

import (
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
// Connectors the connectors property
func (m *PrinterItemRequestBuilder) Connectors()(*icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a.ConnectorsRequestBuilder) {
    return icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printers.item.connectors.item collection
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
func (m *PrinterItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property printers for print
func (m *PrinterItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *PrinterItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PrinterItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the list of printers registered in the tenant.
func (m *PrinterItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PrinterItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
func (m *PrinterItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property printers in print
func (m *PrinterItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable, requestConfiguration *PrinterItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property printers for print
func (m *PrinterItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property printers for print
func (m *PrinterItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *PrinterItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of printers registered in the tenant.
func (m *PrinterItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetCapabilities provides operations to call the getCapabilities method.
func (m *PrinterItemRequestBuilder) GetCapabilities()(*i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.GetCapabilitiesRequestBuilder) {
    return i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.NewGetCapabilitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler the list of printers registered in the tenant.
func (m *PrinterItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *PrinterItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrinterFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable), nil
}
// Patch update the navigation property printers in print
func (m *PrinterItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property printers in print
func (m *PrinterItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Printerable, requestConfiguration *PrinterItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ResetDefaults the resetDefaults property
func (m *PrinterItemRequestBuilder) ResetDefaults()(*id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.ResetDefaultsRequestBuilder) {
    return id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.NewResetDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreFactoryDefaults the restoreFactoryDefaults property
func (m *PrinterItemRequestBuilder) RestoreFactoryDefaults()(*i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.RestoreFactoryDefaultsRequestBuilder) {
    return i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.NewRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Share the share property
func (m *PrinterItemRequestBuilder) Share()(*i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.ShareRequestBuilder) {
    return i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shares the shares property
func (m *PrinterItemRequestBuilder) Shares()(*ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5.SharesRequestBuilder) {
    return ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printers.item.shares.item collection
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
// TaskTriggers the taskTriggers property
func (m *PrinterItemRequestBuilder) TaskTriggers()(*i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb.TaskTriggersRequestBuilder) {
    return i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb.NewTaskTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskTriggersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printers.item.taskTriggers.item collection
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
