package worksheet

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2aa4e02febc1186d80ce9939a3f7c89865378103c6b317d9546cdd66e1992e69 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/worksheet/rangewithaddress"
    i67c6157cb06b04034d73a01041f45e3dddf9787a90576cbb539df467db265383 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/worksheet/usedrangewithvaluesonly"
    i68b86d1e655383b84d3e8b3ceebdaedf702ab5d7b893741dbde7dec2481118cc "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/worksheet/range_escaped"
    iace1ecce9e08ddb41acf75a30f8b0a07bb714657ef5a1fb81f1cfe76663a658f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/worksheet/cellwithrowwithcolumn"
    id93757ae6834a11ce8bb362be039578a2bd9d19e8bd02b56f88a679c9b335e02 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/worksheet/usedrange"
)

// WorksheetRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\worksheet
type WorksheetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorksheetRequestBuilderDeleteOptions options for Delete
type WorksheetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorksheetRequestBuilderGetOptions options for Get
type WorksheetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorksheetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorksheetRequestBuilderGetQueryParameters the worksheet containing the current table. Read-only.
type WorksheetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorksheetRequestBuilderPatchOptions options for Patch
type WorksheetRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CellWithRowWithColumn builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\worksheet\microsoft.graph.cell(row={row},column={column})
func (m *WorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*iace1ecce9e08ddb41acf75a30f8b0a07bb714657ef5a1fb81f1cfe76663a658f.CellWithRowWithColumnRequestBuilder) {
    return iace1ecce9e08ddb41acf75a30f8b0a07bb714657ef5a1fb81f1cfe76663a658f.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
// NewWorksheetRequestBuilderInternal instantiates a new WorksheetRequestBuilder and sets the default values.
func NewWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    m := &WorksheetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/tables/{workbookTable_id}/worksheet{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorksheetRequestBuilder instantiates a new WorksheetRequestBuilder and sets the default values.
func NewWorksheetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the worksheet containing the current table. Read-only.
func (m *WorksheetRequestBuilder) CreateDeleteRequestInformation(options *WorksheetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the worksheet containing the current table. Read-only.
func (m *WorksheetRequestBuilder) CreateGetRequestInformation(options *WorksheetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the worksheet containing the current table. Read-only.
func (m *WorksheetRequestBuilder) CreatePatchRequestInformation(options *WorksheetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the worksheet containing the current table. Read-only.
func (m *WorksheetRequestBuilder) Delete(options *WorksheetRequestBuilderDeleteOptions)(error) {
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
// Get the worksheet containing the current table. Read-only.
func (m *WorksheetRequestBuilder) Get(options *WorksheetRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, error) {
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
// Patch the worksheet containing the current table. Read-only.
func (m *WorksheetRequestBuilder) Patch(options *WorksheetRequestBuilderPatchOptions)(error) {
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
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\worksheet\microsoft.graph.range()
func (m *WorksheetRequestBuilder) Range_escaped()(*i68b86d1e655383b84d3e8b3ceebdaedf702ab5d7b893741dbde7dec2481118cc.RangeRequestBuilder) {
    return i68b86d1e655383b84d3e8b3ceebdaedf702ab5d7b893741dbde7dec2481118cc.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RangeWithAddress builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\worksheet\microsoft.graph.range(address='{address}')
func (m *WorksheetRequestBuilder) RangeWithAddress(address *string)(*i2aa4e02febc1186d80ce9939a3f7c89865378103c6b317d9546cdd66e1992e69.RangeWithAddressRequestBuilder) {
    return i2aa4e02febc1186d80ce9939a3f7c89865378103c6b317d9546cdd66e1992e69.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
// UsedRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\worksheet\microsoft.graph.usedRange()
func (m *WorksheetRequestBuilder) UsedRange()(*id93757ae6834a11ce8bb362be039578a2bd9d19e8bd02b56f88a679c9b335e02.UsedRangeRequestBuilder) {
    return id93757ae6834a11ce8bb362be039578a2bd9d19e8bd02b56f88a679c9b335e02.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedRangeWithValuesOnly builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\worksheet\microsoft.graph.usedRange(valuesOnly={valuesOnly})
func (m *WorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*i67c6157cb06b04034d73a01041f45e3dddf9787a90576cbb539df467db265383.UsedRangeWithValuesOnlyRequestBuilder) {
    return i67c6157cb06b04034d73a01041f45e3dddf9787a90576cbb539df467db265383.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
