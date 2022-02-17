package worksheet

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0309dc107a085057e6d232ddb11b3e4e5d17cc5c603848f4cfaf3cbd7381c9c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/worksheet/usedrange"
    i39ae8911fd8fc184736bed4875aebc0b6662cfacf2818b8c5633a337d17c024d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/worksheet/range_escaped"
    i8ad92e4abb57f63ac30cec5ab367502a6086fd8018044d9641dfdd769da701a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/worksheet/cellwithrowwithcolumn"
    ia5226808d2d6f63c23ce71b748a78e5a3594cd307f8d72fe7527a0530f8c217d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/worksheet/usedrangewithvaluesonly"
    id09b8e4f489c03439293f901405c953a7f2baba5f8592172276c8b5cb0f98731 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/worksheet/rangewithaddress"
)

// WorksheetRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\worksheet
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
// WorksheetRequestBuilderGetQueryParameters the worksheet containing the current chart. Read-only.
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
// CellWithRowWithColumn builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\worksheet\microsoft.graph.cell(row={row},column={column})
func (m *WorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i8ad92e4abb57f63ac30cec5ab367502a6086fd8018044d9641dfdd769da701a6.CellWithRowWithColumnRequestBuilder) {
    return i8ad92e4abb57f63ac30cec5ab367502a6086fd8018044d9641dfdd769da701a6.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
// NewWorksheetRequestBuilderInternal instantiates a new WorksheetRequestBuilder and sets the default values.
func NewWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    m := &WorksheetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/charts/{workbookChart_id}/worksheet{?select,expand}";
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
// CreateDeleteRequestInformation the worksheet containing the current chart. Read-only.
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
// CreateGetRequestInformation the worksheet containing the current chart. Read-only.
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
// CreatePatchRequestInformation the worksheet containing the current chart. Read-only.
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
// Delete the worksheet containing the current chart. Read-only.
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
// Get the worksheet containing the current chart. Read-only.
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
// Patch the worksheet containing the current chart. Read-only.
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
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\worksheet\microsoft.graph.range()
func (m *WorksheetRequestBuilder) Range_escaped()(*i39ae8911fd8fc184736bed4875aebc0b6662cfacf2818b8c5633a337d17c024d.RangeRequestBuilder) {
    return i39ae8911fd8fc184736bed4875aebc0b6662cfacf2818b8c5633a337d17c024d.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RangeWithAddress builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\worksheet\microsoft.graph.range(address='{address}')
func (m *WorksheetRequestBuilder) RangeWithAddress(address *string)(*id09b8e4f489c03439293f901405c953a7f2baba5f8592172276c8b5cb0f98731.RangeWithAddressRequestBuilder) {
    return id09b8e4f489c03439293f901405c953a7f2baba5f8592172276c8b5cb0f98731.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
// UsedRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\worksheet\microsoft.graph.usedRange()
func (m *WorksheetRequestBuilder) UsedRange()(*i0309dc107a085057e6d232ddb11b3e4e5d17cc5c603848f4cfaf3cbd7381c9c7.UsedRangeRequestBuilder) {
    return i0309dc107a085057e6d232ddb11b3e4e5d17cc5c603848f4cfaf3cbd7381c9c7.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedRangeWithValuesOnly builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\worksheet\microsoft.graph.usedRange(valuesOnly={valuesOnly})
func (m *WorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ia5226808d2d6f63c23ce71b748a78e5a3594cd307f8d72fe7527a0530f8c217d.UsedRangeWithValuesOnlyRequestBuilder) {
    return ia5226808d2d6f63c23ce71b748a78e5a3594cd307f8d72fe7527a0530f8c217d.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
