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

// WorkbookWorksheetItemRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}
type WorkbookWorksheetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookWorksheetItemRequestBuilderDeleteOptions options for Delete
type WorkbookWorksheetItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookWorksheetItemRequestBuilderGetOptions options for Get
type WorkbookWorksheetItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookWorksheetItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookWorksheetItemRequestBuilderGetQueryParameters represents a collection of worksheets associated with the workbook. Read-only.
type WorkbookWorksheetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookWorksheetItemRequestBuilderPatchOptions options for Patch
type WorkbookWorksheetItemRequestBuilderPatchOptions struct {
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
func (m *WorkbookWorksheetItemRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i506abd4d70bd1ef1f534fc03014f17697e243f3131a0d1422973220f3a058caf.CellWithRowWithColumnRequestBuilder) {
    return i506abd4d70bd1ef1f534fc03014f17697e243f3131a0d1422973220f3a058caf.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookWorksheetItemRequestBuilder) Charts()(*i2cabb78529abaf7305d60a331ddd55e92756d0de445e4f501407a0a77b794ec1.ChartsRequestBuilder) {
    return i2cabb78529abaf7305d60a331ddd55e92756d0de445e4f501407a0a77b794ec1.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChartsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.charts.item collection
func (m *WorkbookWorksheetItemRequestBuilder) ChartsById(id string)(*iaa8d057031492e203ab61ea5015677620268513945285897a10e9bd84d55bb28.WorkbookChartItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return iaa8d057031492e203ab61ea5015677620268513945285897a10e9bd84d55bb28.NewWorkbookChartItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWorkbookWorksheetItemRequestBuilderInternal instantiates a new WorkbookWorksheetItemRequestBuilder and sets the default values.
func NewWorkbookWorksheetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetItemRequestBuilder) {
    m := &WorkbookWorksheetItemRequestBuilder{
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
// NewWorkbookWorksheetItemRequestBuilder instantiates a new WorkbookWorksheetItemRequestBuilder and sets the default values.
func NewWorkbookWorksheetItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookWorksheetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetItemRequestBuilder) CreateDeleteRequestInformation(options *WorkbookWorksheetItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookWorksheetItemRequestBuilder) CreateGetRequestInformation(options *WorkbookWorksheetItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookWorksheetItemRequestBuilder) CreatePatchRequestInformation(options *WorkbookWorksheetItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookWorksheetItemRequestBuilder) Delete(options *WorkbookWorksheetItemRequestBuilderDeleteOptions)(error) {
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
func (m *WorkbookWorksheetItemRequestBuilder) Get(options *WorkbookWorksheetItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, error) {
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
func (m *WorkbookWorksheetItemRequestBuilder) Names()(*i785a75eeb05b449b6524ae8f860fc59253bc96b05900ce3cd6dc82527d3cf860.NamesRequestBuilder) {
    return i785a75eeb05b449b6524ae8f860fc59253bc96b05900ce3cd6dc82527d3cf860.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.names.item collection
func (m *WorkbookWorksheetItemRequestBuilder) NamesById(id string)(*ibedd3fa72c42cf4638f9da6dee8aebd631d2bc773bf2ba4ec15f2fae677f4d11.WorkbookNamedItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id"] = id
    }
    return ibedd3fa72c42cf4638f9da6dee8aebd631d2bc773bf2ba4ec15f2fae677f4d11.NewWorkbookNamedItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetItemRequestBuilder) Patch(options *WorkbookWorksheetItemRequestBuilderPatchOptions)(error) {
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
func (m *WorkbookWorksheetItemRequestBuilder) PivotTables()(*ie8ad4476a699b68d89501af4b887ad8522a2ba9e909f0b73375470b796a35737.PivotTablesRequestBuilder) {
    return ie8ad4476a699b68d89501af4b887ad8522a2ba9e909f0b73375470b796a35737.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PivotTablesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.pivotTables.item collection
func (m *WorkbookWorksheetItemRequestBuilder) PivotTablesById(id string)(*ic013416d064ea9cbfb60effff722a9a877789877d6335607271b93bc02a13ee8.WorkbookPivotTableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return ic013416d064ea9cbfb60effff722a9a877789877d6335607271b93bc02a13ee8.NewWorkbookPivotTableItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookWorksheetItemRequestBuilder) Protection()(*id310369c610d268ad68e9fd794d3a22bc4d9589b418476f23b54206958f7643e.ProtectionRequestBuilder) {
    return id310369c610d268ad68e9fd794d3a22bc4d9589b418476f23b54206958f7643e.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.range()
func (m *WorkbookWorksheetItemRequestBuilder) Range_escaped()(*i447ad8cdda7b650e77d68669890d2ea6c7d64b9e69b0ab0fa9e7888168f376ac.RangeRequestBuilder) {
    return i447ad8cdda7b650e77d68669890d2ea6c7d64b9e69b0ab0fa9e7888168f376ac.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RangeWithAddress builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.range(address='{address}')
func (m *WorkbookWorksheetItemRequestBuilder) RangeWithAddress(address *string)(*ie83ae106cb34831278c8d2884975c65018c72b352115f284df5c42b53f48b542.RangeWithAddressRequestBuilder) {
    return ie83ae106cb34831278c8d2884975c65018c72b352115f284df5c42b53f48b542.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorkbookWorksheetItemRequestBuilder) Tables()(*i6d4e09fd5e7714b78e7a1eebbc4953e0e7ac8fa7e0bdc1b5c394d19c440f5d2b.TablesRequestBuilder) {
    return i6d4e09fd5e7714b78e7a1eebbc4953e0e7ac8fa7e0bdc1b5c394d19c440f5d2b.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TablesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.tables.item collection
func (m *WorkbookWorksheetItemRequestBuilder) TablesById(id string)(*ice0d09f5356f5fa8492439e7281e18171b98af0f539c95a2d02013b0d4c7d50f.WorkbookTableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return ice0d09f5356f5fa8492439e7281e18171b98af0f539c95a2d02013b0d4c7d50f.NewWorkbookTableItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsedRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.usedRange()
func (m *WorkbookWorksheetItemRequestBuilder) UsedRange()(*i3a815a5e93409d48337d062896c14d9f554d833a002aea7d42ae22a33f331202.UsedRangeRequestBuilder) {
    return i3a815a5e93409d48337d062896c14d9f554d833a002aea7d42ae22a33f331202.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedRangeWithValuesOnly builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.usedRange(valuesOnly={valuesOnly})
func (m *WorkbookWorksheetItemRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*i53d1008f6f0011645bea4539f658b20af47a15df7aa7cf8190c438d244bd4d53.UsedRangeWithValuesOnlyRequestBuilder) {
    return i53d1008f6f0011645bea4539f658b20af47a15df7aa7cf8190c438d244bd4d53.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
