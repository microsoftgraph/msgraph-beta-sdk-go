package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3e76f3b694b2892b3559305dd73a32d8e8037daa8dfa3ed58c3dbe4ce9c458b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/sort"
    i50aec57cf26af7ac706e2ca57bb86ae58f9386583f9c6c389255c7e7b5d9cd3d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/reapplyfilters"
    i60fb9ac6ba5d60f25faed4ae4631ee0fbd3fa80d6bfe4c2f3094cf9d4915cf9b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/clearfilters"
    i6c37a4555679a908182aa3b897857e1a736ea412af7cc59d82636dffce75be27 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/range_escaped"
    i9037314378a6d8a47f599ff5a814639cb40d42e904c840797e1d4d61f484adc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/converttorange"
    i906f1a36fbc3b92bef4711dc68af87cef1d5a363f436d2f95b56f94b90ae115d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/headerrowrange"
    i9d0ee725d6e4543c865c576ed69cb4522119358727e8883a237aad8de98b9ec8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/databodyrange"
    ib8ad0582ac6d5eb3fa199ccb2b91d177ccd45e3e2d6ade21982782cdd8473397 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns"
    idc3b68989b0c7a42d3c8172d0c048c7b6e0b4730d40a5aa60160180a67ea88ea "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/worksheet"
    ie3cdda7496a5b91e5d616b7ecbe82dc4efe9a2b348a5fd5966334bd00b5b567e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/totalrowrange"
    ifd3e20ee0c127a275b04eabd7711e7c5d1ed8343cb9a55181b246906885a4c14 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/rows"
    i5b99478514fc578dc975ffd5c62f32cd2dc2aeda4f651eb15e98d94ee0c38aae "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/rows/item"
    if5cfda93cad0d5440d4a9f28ff3f91b09852f8ebbb9f5f3bad3ad640e342b69b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item"
)

// WorkbookTableRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}
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
func (m *WorkbookTableRequestBuilder) ClearFilters()(*i60fb9ac6ba5d60f25faed4ae4631ee0fbd3fa80d6bfe4c2f3094cf9d4915cf9b.ClearFiltersRequestBuilder) {
    return i60fb9ac6ba5d60f25faed4ae4631ee0fbd3fa80d6bfe4c2f3094cf9d4915cf9b.NewClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Columns()(*ib8ad0582ac6d5eb3fa199ccb2b91d177ccd45e3e2d6ade21982782cdd8473397.ColumnsRequestBuilder) {
    return ib8ad0582ac6d5eb3fa199ccb2b91d177ccd45e3e2d6ade21982782cdd8473397.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.names.item.worksheet.tables.item.columns.item collection
func (m *WorkbookTableRequestBuilder) ColumnsById(id string)(*if5cfda93cad0d5440d4a9f28ff3f91b09852f8ebbb9f5f3bad3ad640e342b69b.WorkbookTableColumnRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableColumn_id"] = id
    }
    return if5cfda93cad0d5440d4a9f28ff3f91b09852f8ebbb9f5f3bad3ad640e342b69b.NewWorkbookTableColumnRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWorkbookTableRequestBuilderInternal instantiates a new WorkbookTableRequestBuilder and sets the default values.
func NewWorkbookTableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    m := &WorkbookTableRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/tables/{workbookTable_id}{?select,expand}";
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
func (m *WorkbookTableRequestBuilder) ConvertToRange()(*i9037314378a6d8a47f599ff5a814639cb40d42e904c840797e1d4d61f484adc5.ConvertToRangeRequestBuilder) {
    return i9037314378a6d8a47f599ff5a814639cb40d42e904c840797e1d4d61f484adc5.NewConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// DataBodyRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableRequestBuilder) DataBodyRange()(*i9d0ee725d6e4543c865c576ed69cb4522119358727e8883a237aad8de98b9ec8.DataBodyRangeRequestBuilder) {
    return i9d0ee725d6e4543c865c576ed69cb4522119358727e8883a237aad8de98b9ec8.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// HeaderRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableRequestBuilder) HeaderRowRange()(*i906f1a36fbc3b92bef4711dc68af87cef1d5a363f436d2f95b56f94b90ae115d.HeaderRowRangeRequestBuilder) {
    return i906f1a36fbc3b92bef4711dc68af87cef1d5a363f436d2f95b56f94b90ae115d.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\microsoft.graph.range()
func (m *WorkbookTableRequestBuilder) Range_escaped()(*i6c37a4555679a908182aa3b897857e1a736ea412af7cc59d82636dffce75be27.RangeRequestBuilder) {
    return i6c37a4555679a908182aa3b897857e1a736ea412af7cc59d82636dffce75be27.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ReapplyFilters()(*i50aec57cf26af7ac706e2ca57bb86ae58f9386583f9c6c389255c7e7b5d9cd3d.ReapplyFiltersRequestBuilder) {
    return i50aec57cf26af7ac706e2ca57bb86ae58f9386583f9c6c389255c7e7b5d9cd3d.NewReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Rows()(*ifd3e20ee0c127a275b04eabd7711e7c5d1ed8343cb9a55181b246906885a4c14.RowsRequestBuilder) {
    return ifd3e20ee0c127a275b04eabd7711e7c5d1ed8343cb9a55181b246906885a4c14.NewRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RowsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.names.item.worksheet.tables.item.rows.item collection
func (m *WorkbookTableRequestBuilder) RowsById(id string)(*i5b99478514fc578dc975ffd5c62f32cd2dc2aeda4f651eb15e98d94ee0c38aae.WorkbookTableRowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableRow_id"] = id
    }
    return i5b99478514fc578dc975ffd5c62f32cd2dc2aeda4f651eb15e98d94ee0c38aae.NewWorkbookTableRowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Sort()(*i3e76f3b694b2892b3559305dd73a32d8e8037daa8dfa3ed58c3dbe4ce9c458b9.SortRequestBuilder) {
    return i3e76f3b694b2892b3559305dd73a32d8e8037daa8dfa3ed58c3dbe4ce9c458b9.NewSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TotalRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableRequestBuilder) TotalRowRange()(*ie3cdda7496a5b91e5d616b7ecbe82dc4efe9a2b348a5fd5966334bd00b5b567e.TotalRowRangeRequestBuilder) {
    return ie3cdda7496a5b91e5d616b7ecbe82dc4efe9a2b348a5fd5966334bd00b5b567e.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Worksheet()(*idc3b68989b0c7a42d3c8172d0c048c7b6e0b4730d40a5aa60160180a67ea88ea.WorksheetRequestBuilder) {
    return idc3b68989b0c7a42d3c8172d0c048c7b6e0b4730d40a5aa60160180a67ea88ea.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
