package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2cabb78529abaf7305d60a331ddd55e92756d0de445e4f501407a0a77b794ec1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts"
    i3a815a5e93409d48337d062896c14d9f554d833a002aea7d42ae22a33f331202 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/usedrange"
    i447ad8cdda7b650e77d68669890d2ea6c7d64b9e69b0ab0fa9e7888168f376ac "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/range_escaped"
    i506abd4d70bd1ef1f534fc03014f17697e243f3131a0d1422973220f3a058caf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/cellwithrowwithcolumn"
    i53d1008f6f0011645bea4539f658b20af47a15df7aa7cf8190c438d244bd4d53 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/usedrangewithvaluesonly"
    i6d4e09fd5e7714b78e7a1eebbc4953e0e7ac8fa7e0bdc1b5c394d19c440f5d2b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables"
    i785a75eeb05b449b6524ae8f860fc59253bc96b05900ce3cd6dc82527d3cf860 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/names"
    id310369c610d268ad68e9fd794d3a22bc4d9589b418476f23b54206958f7643e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/protection"
    ie83ae106cb34831278c8d2884975c65018c72b352115f284df5c42b53f48b542 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/rangewithaddress"
    ie8ad4476a699b68d89501af4b887ad8522a2ba9e909f0b73375470b796a35737 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/pivottables"
    iaa8d057031492e203ab61ea5015677620268513945285897a10e9bd84d55bb28 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item"
    ibedd3fa72c42cf4638f9da6dee8aebd631d2bc773bf2ba4ec15f2fae677f4d11 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/names/item"
    ic013416d064ea9cbfb60effff722a9a877789877d6335607271b93bc02a13ee8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/pivottables/item"
    ice0d09f5356f5fa8492439e7281e18171b98af0f539c95a2d02013b0d4c7d50f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item"
)

// WorkbookWorksheetRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}
type WorkbookWorksheetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookWorksheetRequestBuilderDeleteOptions options for Delete
type WorkbookWorksheetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookWorksheetRequestBuilderGetOptions options for Get
type WorkbookWorksheetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookWorksheetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookWorksheetRequestBuilderGetQueryParameters represents a collection of worksheets associated with the workbook. Read-only.
type WorkbookWorksheetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookWorksheetRequestBuilderPatchOptions options for Patch
type WorkbookWorksheetRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CellWithRowWithColumn builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.cell(row={row},column={column})
func (m *WorkbookWorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i506abd4d70bd1ef1f534fc03014f17697e243f3131a0d1422973220f3a058caf.CellWithRowWithColumnRequestBuilder) {
    return i506abd4d70bd1ef1f534fc03014f17697e243f3131a0d1422973220f3a058caf.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookWorksheetRequestBuilder) Charts()(*i2cabb78529abaf7305d60a331ddd55e92756d0de445e4f501407a0a77b794ec1.ChartsRequestBuilder) {
    return i2cabb78529abaf7305d60a331ddd55e92756d0de445e4f501407a0a77b794ec1.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChartsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.charts.item collection
func (m *WorkbookWorksheetRequestBuilder) ChartsById(id string)(*iaa8d057031492e203ab61ea5015677620268513945285897a10e9bd84d55bb28.WorkbookChartRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return iaa8d057031492e203ab61ea5015677620268513945285897a10e9bd84d55bb28.NewWorkbookChartRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWorkbookWorksheetRequestBuilderInternal instantiates a new WorkbookWorksheetRequestBuilder and sets the default values.
func NewWorkbookWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetRequestBuilder) {
    m := &WorkbookWorksheetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookWorksheetRequestBuilder instantiates a new WorkbookWorksheetRequestBuilder and sets the default values.
func NewWorkbookWorksheetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) CreateDeleteRequestInformation(options *WorkbookWorksheetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) CreateGetRequestInformation(options *WorkbookWorksheetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) CreatePatchRequestInformation(options *WorkbookWorksheetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) Delete(options *WorkbookWorksheetRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) Get(options *WorkbookWorksheetRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookWorksheet() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet), nil
}
func (m *WorkbookWorksheetRequestBuilder) Names()(*i785a75eeb05b449b6524ae8f860fc59253bc96b05900ce3cd6dc82527d3cf860.NamesRequestBuilder) {
    return i785a75eeb05b449b6524ae8f860fc59253bc96b05900ce3cd6dc82527d3cf860.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.names.item collection
func (m *WorkbookWorksheetRequestBuilder) NamesById(id string)(*ibedd3fa72c42cf4638f9da6dee8aebd631d2bc773bf2ba4ec15f2fae677f4d11.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id"] = id
    }
    return ibedd3fa72c42cf4638f9da6dee8aebd631d2bc773bf2ba4ec15f2fae677f4d11.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) Patch(options *WorkbookWorksheetRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookWorksheetRequestBuilder) PivotTables()(*ie8ad4476a699b68d89501af4b887ad8522a2ba9e909f0b73375470b796a35737.PivotTablesRequestBuilder) {
    return ie8ad4476a699b68d89501af4b887ad8522a2ba9e909f0b73375470b796a35737.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PivotTablesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.pivotTables.item collection
func (m *WorkbookWorksheetRequestBuilder) PivotTablesById(id string)(*ic013416d064ea9cbfb60effff722a9a877789877d6335607271b93bc02a13ee8.WorkbookPivotTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return ic013416d064ea9cbfb60effff722a9a877789877d6335607271b93bc02a13ee8.NewWorkbookPivotTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) Protection()(*id310369c610d268ad68e9fd794d3a22bc4d9589b418476f23b54206958f7643e.ProtectionRequestBuilder) {
    return id310369c610d268ad68e9fd794d3a22bc4d9589b418476f23b54206958f7643e.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.range()
func (m *WorkbookWorksheetRequestBuilder) Range_escaped()(*i447ad8cdda7b650e77d68669890d2ea6c7d64b9e69b0ab0fa9e7888168f376ac.RangeRequestBuilder) {
    return i447ad8cdda7b650e77d68669890d2ea6c7d64b9e69b0ab0fa9e7888168f376ac.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RangeWithAddress builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.range(address='{address}')
func (m *WorkbookWorksheetRequestBuilder) RangeWithAddress(address *string)(*ie83ae106cb34831278c8d2884975c65018c72b352115f284df5c42b53f48b542.RangeWithAddressRequestBuilder) {
    return ie83ae106cb34831278c8d2884975c65018c72b352115f284df5c42b53f48b542.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorkbookWorksheetRequestBuilder) Tables()(*i6d4e09fd5e7714b78e7a1eebbc4953e0e7ac8fa7e0bdc1b5c394d19c440f5d2b.TablesRequestBuilder) {
    return i6d4e09fd5e7714b78e7a1eebbc4953e0e7ac8fa7e0bdc1b5c394d19c440f5d2b.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TablesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.tables.item collection
func (m *WorkbookWorksheetRequestBuilder) TablesById(id string)(*ice0d09f5356f5fa8492439e7281e18171b98af0f539c95a2d02013b0d4c7d50f.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return ice0d09f5356f5fa8492439e7281e18171b98af0f539c95a2d02013b0d4c7d50f.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsedRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.usedRange()
func (m *WorkbookWorksheetRequestBuilder) UsedRange()(*i3a815a5e93409d48337d062896c14d9f554d833a002aea7d42ae22a33f331202.UsedRangeRequestBuilder) {
    return i3a815a5e93409d48337d062896c14d9f554d833a002aea7d42ae22a33f331202.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedRangeWithValuesOnly builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.usedRange(valuesOnly={valuesOnly})
func (m *WorkbookWorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*i53d1008f6f0011645bea4539f658b20af47a15df7aa7cf8190c438d244bd4d53.UsedRangeWithValuesOnlyRequestBuilder) {
    return i53d1008f6f0011645bea4539f658b20af47a15df7aa7cf8190c438d244bd4d53.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
