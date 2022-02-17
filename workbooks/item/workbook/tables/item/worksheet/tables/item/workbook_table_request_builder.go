package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i03e8315740e50741be28e0de5b0dbc979248935d5a52b8bd57ae1a03bb0cc9ce "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/range_escaped"
    i32ae25ae9d9d6e1f97d393a3512d4664345ae284435f62e14907150840fb4f23 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/totalrowrange"
    i33f608b1a5b1f40cbedd25ca23610dc70efd3ef30bdf83ca0364a5bcc6219b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/databodyrange"
    i576ef957c75800d01681f5c99bb2e8bdf4192b3bc6a0aeb0e21eeb43b3c558af "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/headerrowrange"
    i62b751f5318254f7603879bab6c4ed76db74fc7e7c7b9e719a6a79312b0b88f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/reapplyfilters"
    i8aa38c528205520959488d8f2a7e023f19e7f8a0bf2f36873d6c50ab2150f8c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/clearfilters"
    ia19841dfd8dac02a16ef964e9ae689716da5ba326c0089ac50187e4ce5fa6e7d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/converttorange"
)

// WorkbookTableRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}
type WorkbookTableRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookTableRequestBuilderDeleteOptions options for Delete
type WorkbookTableRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookTableRequestBuilderGetOptions options for Get
type WorkbookTableRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookTableRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookTableRequestBuilderGetQueryParameters collection of tables that are part of the worksheet. Read-only.
type WorkbookTableRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookTableRequestBuilderPatchOptions options for Patch
type WorkbookTableRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WorkbookTableRequestBuilder) ClearFilters()(*i8aa38c528205520959488d8f2a7e023f19e7f8a0bf2f36873d6c50ab2150f8c3.ClearFiltersRequestBuilder) {
    return i8aa38c528205520959488d8f2a7e023f19e7f8a0bf2f36873d6c50ab2150f8c3.NewClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookTableRequestBuilderInternal instantiates a new WorkbookTableRequestBuilder and sets the default values.
func NewWorkbookTableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    m := &WorkbookTableRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/worksheet/tables/{workbookTable_id1}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookTableRequestBuilder instantiates a new WorkbookTableRequestBuilder and sets the default values.
func NewWorkbookTableRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookTableRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookTableRequestBuilder) ConvertToRange()(*ia19841dfd8dac02a16ef964e9ae689716da5ba326c0089ac50187e4ce5fa6e7d.ConvertToRangeRequestBuilder) {
    return ia19841dfd8dac02a16ef964e9ae689716da5ba326c0089ac50187e4ce5fa6e7d.NewConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookTableRequestBuilder) CreateDeleteRequestInformation(options *WorkbookTableRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookTableRequestBuilder) CreateGetRequestInformation(options *WorkbookTableRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookTableRequestBuilder) CreatePatchRequestInformation(options *WorkbookTableRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// DataBodyRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableRequestBuilder) DataBodyRange()(*i33f608b1a5b1f40cbedd25ca23610dc70efd3ef30bdf83ca0364a5bcc6219b3c.DataBodyRangeRequestBuilder) {
    return i33f608b1a5b1f40cbedd25ca23610dc70efd3ef30bdf83ca0364a5bcc6219b3c.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookTableRequestBuilder) Delete(options *WorkbookTableRequestBuilderDeleteOptions)(error) {
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
// Get collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookTableRequestBuilder) Get(options *WorkbookTableRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookTable() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable), nil
}
// HeaderRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}\microsoft.graph.headerRowRange()
func (m *WorkbookTableRequestBuilder) HeaderRowRange()(*i576ef957c75800d01681f5c99bb2e8bdf4192b3bc6a0aeb0e21eeb43b3c558af.HeaderRowRangeRequestBuilder) {
    return i576ef957c75800d01681f5c99bb2e8bdf4192b3bc6a0aeb0e21eeb43b3c558af.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookTableRequestBuilder) Patch(options *WorkbookTableRequestBuilderPatchOptions)(error) {
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
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}\microsoft.graph.range()
func (m *WorkbookTableRequestBuilder) Range_escaped()(*i03e8315740e50741be28e0de5b0dbc979248935d5a52b8bd57ae1a03bb0cc9ce.RangeRequestBuilder) {
    return i03e8315740e50741be28e0de5b0dbc979248935d5a52b8bd57ae1a03bb0cc9ce.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ReapplyFilters()(*i62b751f5318254f7603879bab6c4ed76db74fc7e7c7b9e719a6a79312b0b88f1.ReapplyFiltersRequestBuilder) {
    return i62b751f5318254f7603879bab6c4ed76db74fc7e7c7b9e719a6a79312b0b88f1.NewReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TotalRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}\microsoft.graph.totalRowRange()
func (m *WorkbookTableRequestBuilder) TotalRowRange()(*i32ae25ae9d9d6e1f97d393a3512d4664345ae284435f62e14907150840fb4f23.TotalRowRangeRequestBuilder) {
    return i32ae25ae9d9d6e1f97d393a3512d4664345ae284435f62e14907150840fb4f23.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
