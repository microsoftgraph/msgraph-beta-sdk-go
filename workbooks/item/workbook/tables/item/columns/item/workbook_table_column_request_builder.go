package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0cbb31e408bccdef9c2b90150f9b32761053d47747c98476004719ae57aeda8a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/range_escaped"
    i1b3dc0be7b958b596fa2c94ca04a4dc4ecb8b6b0c600f1cbe0b8f3da032e1f2b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter"
    i1ecd7845fb939da1c6df676225e72fd72ebb1530bae854675598ef1d7d7cfbda "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/databodyrange"
    i88eedb45dd017baf320026f8cc8a30ea7ef664b43fa3fe9c73f7f57ae2fd3133 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/totalrowrange"
    if0d993aaca17c5a3533457a2c51a2b6b255a9dede015ecadedc6434d1e962ff2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/headerrowrange"
)

// WorkbookTableColumnRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}
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
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/columns/{workbookTableColumn_id}{?select,expand}";
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
// DataBodyRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableColumnRequestBuilder) DataBodyRange()(*i1ecd7845fb939da1c6df676225e72fd72ebb1530bae854675598ef1d7d7cfbda.DataBodyRangeRequestBuilder) {
    return i1ecd7845fb939da1c6df676225e72fd72ebb1530bae854675598ef1d7d7cfbda.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *WorkbookTableColumnRequestBuilder) Filter()(*i1b3dc0be7b958b596fa2c94ca04a4dc4ecb8b6b0c600f1cbe0b8f3da032e1f2b.FilterRequestBuilder) {
    return i1b3dc0be7b958b596fa2c94ca04a4dc4ecb8b6b0c600f1cbe0b8f3da032e1f2b.NewFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// HeaderRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableColumnRequestBuilder) HeaderRowRange()(*if0d993aaca17c5a3533457a2c51a2b6b255a9dede015ecadedc6434d1e962ff2.HeaderRowRangeRequestBuilder) {
    return if0d993aaca17c5a3533457a2c51a2b6b255a9dede015ecadedc6434d1e962ff2.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.range()
func (m *WorkbookTableColumnRequestBuilder) Range_escaped()(*i0cbb31e408bccdef9c2b90150f9b32761053d47747c98476004719ae57aeda8a.RangeRequestBuilder) {
    return i0cbb31e408bccdef9c2b90150f9b32761053d47747c98476004719ae57aeda8a.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TotalRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableColumnRequestBuilder) TotalRowRange()(*i88eedb45dd017baf320026f8cc8a30ea7ef664b43fa3fe9c73f7f57ae2fd3133.TotalRowRangeRequestBuilder) {
    return i88eedb45dd017baf320026f8cc8a30ea7ef664b43fa3fe9c73f7f57ae2fd3133.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
