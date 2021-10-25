package worksheet

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3cc2b0994e6db1b9b44828f245e0ee096a8fdf21a35e8e57a8d683fbe37a6c66 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/usedrange"
    i4d559a9d2f0944faabdac6b995962e770a2712e24295d8ca1cdbc3e671497200 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables"
    i4e3dc564f15220165545ca9c1a8051d4a51819b13eb458688f66c3e1a96173ad "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/names"
    i4e69ab029b04e85585b291a457f60a62bf92195cf27bc2b4174b057890ea6009 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/protection"
    i5bdd3973ea6dd44c1d10392ef009b87ddc90b512db5ff2af41ddee1a0ebb354a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/pivottables"
    i6a24d88889632583818690aa68e3dfe278c266101b178175d7227e129f24cf05 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/rangewithaddress"
    i7383726aa3e9350c08b2ea8eb7a8dc1dbb515c3d786118b02977ac27f231f649 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/range_escpaped"
    i85a27dcf9694844ca6098cd4783ad70b6d561ca6aac3a5fec6cb4c5fe087fe00 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/usedrangewithvaluesonly"
    i92b646ab946858dd65bf5823c3b08054a63323a9fddd3b73ff1804c7d7be84d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts"
    ic13bd4095b5e1a5c24e414e41a1d0c05d5c2ab8dac0ef1793f202c1b13215d7e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/cellwithrowwithcolumn"
    i243ae082e537c275aab89f17dc8215182845a67c50c901de517acb917c8d7aa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item"
    i49b3386da1523174e42fd24449fffd0274c9d83ffccc997bf68eb0d44ba1d532 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/pivottables/item"
    i612609f592af5542250f428a4e388cb93b22d8513b39183b74caaa634486d4fe "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item"
    if273acbaa7735a0bee5fb7cce891f47f0fb42d1f4d0a580f785a50ff7a961cd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/names/item"
)

type WorksheetRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorksheetRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *WorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*ic13bd4095b5e1a5c24e414e41a1d0c05d5c2ab8dac0ef1793f202c1b13215d7e.CellWithRowWithColumnRequestBuilder) {
    return ic13bd4095b5e1a5c24e414e41a1d0c05d5c2ab8dac0ef1793f202c1b13215d7e.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorksheetRequestBuilder) Charts()(*i92b646ab946858dd65bf5823c3b08054a63323a9fddd3b73ff1804c7d7be84d7.ChartsRequestBuilder) {
    return i92b646ab946858dd65bf5823c3b08054a63323a9fddd3b73ff1804c7d7be84d7.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) ChartsById(id string)(*i612609f592af5542250f428a4e388cb93b22d8513b39183b74caaa634486d4fe.WorkbookChartRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return i612609f592af5542250f428a4e388cb93b22d8513b39183b74caaa634486d4fe.NewWorkbookChartRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    m := &WorksheetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewWorksheetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorksheetRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorksheetRequestBuilder) CreateGetRequestInformation(q func (value *WorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorksheetRequestBuilderGetQueryParameters)
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
func (m *WorksheetRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorksheetRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorksheetRequestBuilder) Get(q func (value *WorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookWorksheet() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet), nil
}
func (m *WorksheetRequestBuilder) Names()(*i4e3dc564f15220165545ca9c1a8051d4a51819b13eb458688f66c3e1a96173ad.NamesRequestBuilder) {
    return i4e3dc564f15220165545ca9c1a8051d4a51819b13eb458688f66c3e1a96173ad.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) NamesById(id string)(*if273acbaa7735a0bee5fb7cce891f47f0fb42d1f4d0a580f785a50ff7a961cd0.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id1"] = id
    }
    return if273acbaa7735a0bee5fb7cce891f47f0fb42d1f4d0a580f785a50ff7a961cd0.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorksheetRequestBuilder) PivotTables()(*i5bdd3973ea6dd44c1d10392ef009b87ddc90b512db5ff2af41ddee1a0ebb354a.PivotTablesRequestBuilder) {
    return i5bdd3973ea6dd44c1d10392ef009b87ddc90b512db5ff2af41ddee1a0ebb354a.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) PivotTablesById(id string)(*i49b3386da1523174e42fd24449fffd0274c9d83ffccc997bf68eb0d44ba1d532.WorkbookPivotTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return i49b3386da1523174e42fd24449fffd0274c9d83ffccc997bf68eb0d44ba1d532.NewWorkbookPivotTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Protection()(*i4e69ab029b04e85585b291a457f60a62bf92195cf27bc2b4174b057890ea6009.ProtectionRequestBuilder) {
    return i4e69ab029b04e85585b291a457f60a62bf92195cf27bc2b4174b057890ea6009.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Range_escpaped()(*i7383726aa3e9350c08b2ea8eb7a8dc1dbb515c3d786118b02977ac27f231f649.RangeRequestBuilder) {
    return i7383726aa3e9350c08b2ea8eb7a8dc1dbb515c3d786118b02977ac27f231f649.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) RangeWithAddress(address *string)(*i6a24d88889632583818690aa68e3dfe278c266101b178175d7227e129f24cf05.RangeWithAddressRequestBuilder) {
    return i6a24d88889632583818690aa68e3dfe278c266101b178175d7227e129f24cf05.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorksheetRequestBuilder) Tables()(*i4d559a9d2f0944faabdac6b995962e770a2712e24295d8ca1cdbc3e671497200.TablesRequestBuilder) {
    return i4d559a9d2f0944faabdac6b995962e770a2712e24295d8ca1cdbc3e671497200.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) TablesById(id string)(*i243ae082e537c275aab89f17dc8215182845a67c50c901de517acb917c8d7aa3.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return i243ae082e537c275aab89f17dc8215182845a67c50c901de517acb917c8d7aa3.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) UsedRange()(*i3cc2b0994e6db1b9b44828f245e0ee096a8fdf21a35e8e57a8d683fbe37a6c66.UsedRangeRequestBuilder) {
    return i3cc2b0994e6db1b9b44828f245e0ee096a8fdf21a35e8e57a8d683fbe37a6c66.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*i85a27dcf9694844ca6098cd4783ad70b6d561ca6aac3a5fec6cb4c5fe087fe00.UsedRangeWithValuesOnlyRequestBuilder) {
    return i85a27dcf9694844ca6098cd4783ad70b6d561ca6aac3a5fec6cb4c5fe087fe00.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
