package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/tasktriggers"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrinterItemRequestBuilderDeleteOptions options for Delete
type PrinterItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrinterItemRequestBuilderGetOptions options for Get
type PrinterItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrinterItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrinterItemRequestBuilderGetQueryParameters the list of printers registered in the tenant.
type PrinterItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrinterItemRequestBuilderPatchOptions options for Patch
type PrinterItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printerable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["printConnector_id"] = id
    }
    return ic1c34d114926cc52d44d68cfff4c584108994c1b4377e24401d8c75538512384.NewPrintConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrinterItemRequestBuilderInternal instantiates a new PrinterItemRequestBuilder and sets the default values.
func NewPrinterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterItemRequestBuilder) {
    m := &PrinterItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printers/{printer_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrinterItemRequestBuilder instantiates a new PrinterItemRequestBuilder and sets the default values.
func NewPrinterItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property printers for print
func (m *PrinterItemRequestBuilder) CreateDeleteRequestInformation(options *PrinterItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of printers registered in the tenant.
func (m *PrinterItemRequestBuilder) CreateGetRequestInformation(options *PrinterItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property printers in print
func (m *PrinterItemRequestBuilder) CreatePatchRequestInformation(options *PrinterItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property printers for print
func (m *PrinterItemRequestBuilder) Delete(options *PrinterItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of printers registered in the tenant.
func (m *PrinterItemRequestBuilder) Get(options *PrinterItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreatePrinterFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printerable), nil
}
// GetCapabilities provides operations to call the getCapabilities method.
func (m *PrinterItemRequestBuilder) GetCapabilities()(*i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.GetCapabilitiesRequestBuilder) {
    return i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.NewGetCapabilitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property printers in print
func (m *PrinterItemRequestBuilder) Patch(options *PrinterItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *PrinterItemRequestBuilder) ResetDefaults()(*id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.ResetDefaultsRequestBuilder) {
    return id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.NewResetDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterItemRequestBuilder) RestoreFactoryDefaults()(*i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.RestoreFactoryDefaultsRequestBuilder) {
    return i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.NewRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterItemRequestBuilder) Share()(*i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.ShareRequestBuilder) {
    return i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["printerShare_id"] = id
    }
    return i1b55b9556d3f0f8504c7f2a35c257328ab30ecf9e03df33df4495177129a2cdf.NewPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["printTaskTrigger_id"] = id
    }
    return i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d.NewPrintTaskTriggerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}