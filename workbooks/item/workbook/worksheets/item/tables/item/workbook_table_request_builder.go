package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i16928683cb77a9cab228d1c15741977a9c69ce42f5beca4d51b01686c3183df5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns"
    i290776d40de3e0417a597844398218739052bc3ab0de2474aa8033b7f4416996 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/sort"
    i527a5eaa6f2822882c957ec83b97b66b96f9406d429a6188c010c781863c5210 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/headerrowrange"
    i57feb25878488a3a2ba08d67a30e44df113e3f3d88e89be3cbeca57beddc7e35 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/range_escaped"
    i58a3264aa4edaf82ceb001b3eeff2b1c74892f596e898f97c7539c9995624c5a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/reapplyfilters"
    i791a12f5f3fe0fb8ce19992be08700bb6ad6a29be1b253bbed59714e636ccbe7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/converttorange"
    i7d5fbf49355209347913c698e527321689f085d88a5633d89cd395d46f7a676d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/totalrowrange"
    ia5040aa8b6bbd0905fa3dd40063ad26c26c63d84a40d143e3a9e85375534167f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/clearfilters"
    idab150ee0f159d9e6317aaa0342a137b89b5f9ef956b159fe50950c8ff998608 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/worksheet"
    ieb9658e86bfc7d5d8fa952719520400465702f4ef64328f785ffc81c51d6d3a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/rows"
    ife82c54b58f44f9d5198dfbaf386e863e4ab89f65f4a036f7a49f4d28b267b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/databodyrange"
    i315bca7f8ff97d9a48ce543dfa3bd30d52b8535ddcac982392e0ccbe28505e65 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item"
    ib84ec0df39d8fe799726c480b40091ae6ee109e781d54635593a8d7dd8f22c2d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/rows/item"
)

type WorkbookTableRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorkbookTableRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *WorkbookTableRequestBuilder) ClearFilters()(*ia5040aa8b6bbd0905fa3dd40063ad26c26c63d84a40d143e3a9e85375534167f.ClearFiltersRequestBuilder) {
    return ia5040aa8b6bbd0905fa3dd40063ad26c26c63d84a40d143e3a9e85375534167f.NewClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Columns()(*i16928683cb77a9cab228d1c15741977a9c69ce42f5beca4d51b01686c3183df5.ColumnsRequestBuilder) {
    return i16928683cb77a9cab228d1c15741977a9c69ce42f5beca4d51b01686c3183df5.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ColumnsById(id string)(*i315bca7f8ff97d9a48ce543dfa3bd30d52b8535ddcac982392e0ccbe28505e65.WorkbookTableColumnRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableColumn_id"] = id
    }
    return i315bca7f8ff97d9a48ce543dfa3bd30d52b8535ddcac982392e0ccbe28505e65.NewWorkbookTableColumnRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWorkbookTableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    m := &WorkbookTableRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/tables/{workbookTable_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewWorkbookTableRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookTableRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookTableRequestBuilder) ConvertToRange()(*i791a12f5f3fe0fb8ce19992be08700bb6ad6a29be1b253bbed59714e636ccbe7.ConvertToRangeRequestBuilder) {
    return i791a12f5f3fe0fb8ce19992be08700bb6ad6a29be1b253bbed59714e636ccbe7.NewConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *WorkbookTableRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookTableRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookTableRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *WorkbookTableRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *WorkbookTableRequestBuilder) DataBodyRange()(*ife82c54b58f44f9d5198dfbaf386e863e4ab89f65f4a036f7a49f4d28b267b3c.DataBodyRangeRequestBuilder) {
    return ife82c54b58f44f9d5198dfbaf386e863e4ab89f65f4a036f7a49f4d28b267b3c.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookTableRequestBuilder) Get(q func (value *WorkbookTableRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookTable() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable), nil
}
func (m *WorkbookTableRequestBuilder) HeaderRowRange()(*i527a5eaa6f2822882c957ec83b97b66b96f9406d429a6188c010c781863c5210.HeaderRowRangeRequestBuilder) {
    return i527a5eaa6f2822882c957ec83b97b66b96f9406d429a6188c010c781863c5210.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookTableRequestBuilder) Range_escaped()(*i57feb25878488a3a2ba08d67a30e44df113e3f3d88e89be3cbeca57beddc7e35.RangeRequestBuilder) {
    return i57feb25878488a3a2ba08d67a30e44df113e3f3d88e89be3cbeca57beddc7e35.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ReapplyFilters()(*i58a3264aa4edaf82ceb001b3eeff2b1c74892f596e898f97c7539c9995624c5a.ReapplyFiltersRequestBuilder) {
    return i58a3264aa4edaf82ceb001b3eeff2b1c74892f596e898f97c7539c9995624c5a.NewReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Rows()(*ieb9658e86bfc7d5d8fa952719520400465702f4ef64328f785ffc81c51d6d3a1.RowsRequestBuilder) {
    return ieb9658e86bfc7d5d8fa952719520400465702f4ef64328f785ffc81c51d6d3a1.NewRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) RowsById(id string)(*ib84ec0df39d8fe799726c480b40091ae6ee109e781d54635593a8d7dd8f22c2d.WorkbookTableRowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableRow_id"] = id
    }
    return ib84ec0df39d8fe799726c480b40091ae6ee109e781d54635593a8d7dd8f22c2d.NewWorkbookTableRowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Sort()(*i290776d40de3e0417a597844398218739052bc3ab0de2474aa8033b7f4416996.SortRequestBuilder) {
    return i290776d40de3e0417a597844398218739052bc3ab0de2474aa8033b7f4416996.NewSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) TotalRowRange()(*i7d5fbf49355209347913c698e527321689f085d88a5633d89cd395d46f7a676d.TotalRowRangeRequestBuilder) {
    return i7d5fbf49355209347913c698e527321689f085d88a5633d89cd395d46f7a676d.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Worksheet()(*idab150ee0f159d9e6317aaa0342a137b89b5f9ef956b159fe50950c8ff998608.WorksheetRequestBuilder) {
    return idab150ee0f159d9e6317aaa0342a137b89b5f9ef956b159fe50950c8ff998608.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
