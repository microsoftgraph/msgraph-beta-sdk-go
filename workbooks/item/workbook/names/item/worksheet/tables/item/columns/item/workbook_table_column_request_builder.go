package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1362cbc0cab6ed9bd7fa91638adc4e3f60f4aafc46d441a8e9c395dbdd8b40d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/databodyrange"
    i5baac89d317e6c6826be1fd842f5bfe712bd3b4fec439c5d6bb3b6957193e309 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/headerrowrange"
    ia861fa8bb974eeaec16e6a8897e48002ba16a1dc14bb92da3eccd624e9be4301 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/totalrowrange"
    ic4e0951a61f51f6e21fe3d5ea26d68d429d8f4d6245e3636253a65b0d1af2f3c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter"
    ifcb19eb044149b972a3eb701d6f39cf2ce9bfde706662be1cc9193bea8446921 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/range_escaped"
)

// WorkbookTableColumnRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\columns\{workbookTableColumn-id}
type WorkbookTableColumnRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookTableColumnRequestBuilderDeleteOptions options for Delete
type WorkbookTableColumnRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookTableColumnRequestBuilderGetOptions options for Get
type WorkbookTableColumnRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookTableColumnRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookTableColumnRequestBuilderGetQueryParameters represents a collection of all the columns in the table. Read-only.
type WorkbookTableColumnRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookTableColumnRequestBuilderPatchOptions options for Patch
type WorkbookTableColumnRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableColumn;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWorkbookTableColumnRequestBuilderInternal instantiates a new WorkbookTableColumnRequestBuilder and sets the default values.
func NewWorkbookTableColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableColumnRequestBuilder) {
    m := &WorkbookTableColumnRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/tables/{workbookTable_id}/columns/{workbookTableColumn_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookTableColumnRequestBuilder instantiates a new WorkbookTableColumnRequestBuilder and sets the default values.
func NewWorkbookTableColumnRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableColumnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookTableColumnRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) CreateDeleteRequestInformation(options *WorkbookTableColumnRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) CreateGetRequestInformation(options *WorkbookTableColumnRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) CreatePatchRequestInformation(options *WorkbookTableColumnRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// DataBodyRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableColumnRequestBuilder) DataBodyRange()(*i1362cbc0cab6ed9bd7fa91638adc4e3f60f4aafc46d441a8e9c395dbdd8b40d8.DataBodyRangeRequestBuilder) {
    return i1362cbc0cab6ed9bd7fa91638adc4e3f60f4aafc46d441a8e9c395dbdd8b40d8.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) Delete(options *WorkbookTableColumnRequestBuilderDeleteOptions)(error) {
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
func (m *WorkbookTableColumnRequestBuilder) Filter()(*ic4e0951a61f51f6e21fe3d5ea26d68d429d8f4d6245e3636253a65b0d1af2f3c.FilterRequestBuilder) {
    return ic4e0951a61f51f6e21fe3d5ea26d68d429d8f4d6245e3636253a65b0d1af2f3c.NewFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) Get(options *WorkbookTableColumnRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableColumn, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookTableColumn() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableColumn), nil
}
// HeaderRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableColumnRequestBuilder) HeaderRowRange()(*i5baac89d317e6c6826be1fd842f5bfe712bd3b4fec439c5d6bb3b6957193e309.HeaderRowRangeRequestBuilder) {
    return i5baac89d317e6c6826be1fd842f5bfe712bd3b4fec439c5d6bb3b6957193e309.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) Patch(options *WorkbookTableColumnRequestBuilderPatchOptions)(error) {
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
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.range()
func (m *WorkbookTableColumnRequestBuilder) Range_escaped()(*ifcb19eb044149b972a3eb701d6f39cf2ce9bfde706662be1cc9193bea8446921.RangeRequestBuilder) {
    return ifcb19eb044149b972a3eb701d6f39cf2ce9bfde706662be1cc9193bea8446921.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TotalRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableColumnRequestBuilder) TotalRowRange()(*ia861fa8bb974eeaec16e6a8897e48002ba16a1dc14bb92da3eccd624e9be4301.TotalRowRangeRequestBuilder) {
    return ia861fa8bb974eeaec16e6a8897e48002ba16a1dc14bb92da3eccd624e9be4301.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
