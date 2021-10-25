package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2cabb78529abaf7305d60a331ddd55e92756d0de445e4f501407a0a77b794ec1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts"
    i3a815a5e93409d48337d062896c14d9f554d833a002aea7d42ae22a33f331202 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/usedrange"
    i4b26040f589583b2a924ea3fbc7a2e9353bd60a36e46660cf74a03f2649b3ae3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/range_escpaped"
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

type WorkbookWorksheetRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorkbookWorksheetRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *WorkbookWorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i506abd4d70bd1ef1f534fc03014f17697e243f3131a0d1422973220f3a058caf.CellWithRowWithColumnRequestBuilder) {
    return i506abd4d70bd1ef1f534fc03014f17697e243f3131a0d1422973220f3a058caf.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookWorksheetRequestBuilder) Charts()(*i2cabb78529abaf7305d60a331ddd55e92756d0de445e4f501407a0a77b794ec1.ChartsRequestBuilder) {
    return i2cabb78529abaf7305d60a331ddd55e92756d0de445e4f501407a0a77b794ec1.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) ChartsById(id string)(*iaa8d057031492e203ab61ea5015677620268513945285897a10e9bd84d55bb28.WorkbookChartRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return iaa8d057031492e203ab61ea5015677620268513945285897a10e9bd84d55bb28.NewWorkbookChartRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWorkbookWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetRequestBuilder) {
    m := &WorkbookWorksheetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}{?select,expand}";
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
func NewWorkbookWorksheetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookWorksheetRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookWorksheetRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookWorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookWorksheetRequestBuilderGetQueryParameters)
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
func (m *WorkbookWorksheetRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookWorksheetRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorkbookWorksheetRequestBuilder) Get(q func (value *WorkbookWorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, error) {
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
func (m *WorkbookWorksheetRequestBuilder) Names()(*i785a75eeb05b449b6524ae8f860fc59253bc96b05900ce3cd6dc82527d3cf860.NamesRequestBuilder) {
    return i785a75eeb05b449b6524ae8f860fc59253bc96b05900ce3cd6dc82527d3cf860.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) NamesById(id string)(*ibedd3fa72c42cf4638f9da6dee8aebd631d2bc773bf2ba4ec15f2fae677f4d11.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id"] = id
    }
    return ibedd3fa72c42cf4638f9da6dee8aebd631d2bc773bf2ba4ec15f2fae677f4d11.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorkbookWorksheetRequestBuilder) PivotTables()(*ie8ad4476a699b68d89501af4b887ad8522a2ba9e909f0b73375470b796a35737.PivotTablesRequestBuilder) {
    return ie8ad4476a699b68d89501af4b887ad8522a2ba9e909f0b73375470b796a35737.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) PivotTablesById(id string)(*ic013416d064ea9cbfb60effff722a9a877789877d6335607271b93bc02a13ee8.WorkbookPivotTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return ic013416d064ea9cbfb60effff722a9a877789877d6335607271b93bc02a13ee8.NewWorkbookPivotTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) Protection()(*id310369c610d268ad68e9fd794d3a22bc4d9589b418476f23b54206958f7643e.ProtectionRequestBuilder) {
    return id310369c610d268ad68e9fd794d3a22bc4d9589b418476f23b54206958f7643e.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) Range_escpaped()(*i4b26040f589583b2a924ea3fbc7a2e9353bd60a36e46660cf74a03f2649b3ae3.RangeRequestBuilder) {
    return i4b26040f589583b2a924ea3fbc7a2e9353bd60a36e46660cf74a03f2649b3ae3.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) RangeWithAddress(address *string)(*ie83ae106cb34831278c8d2884975c65018c72b352115f284df5c42b53f48b542.RangeWithAddressRequestBuilder) {
    return ie83ae106cb34831278c8d2884975c65018c72b352115f284df5c42b53f48b542.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorkbookWorksheetRequestBuilder) Tables()(*i6d4e09fd5e7714b78e7a1eebbc4953e0e7ac8fa7e0bdc1b5c394d19c440f5d2b.TablesRequestBuilder) {
    return i6d4e09fd5e7714b78e7a1eebbc4953e0e7ac8fa7e0bdc1b5c394d19c440f5d2b.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) TablesById(id string)(*ice0d09f5356f5fa8492439e7281e18171b98af0f539c95a2d02013b0d4c7d50f.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return ice0d09f5356f5fa8492439e7281e18171b98af0f539c95a2d02013b0d4c7d50f.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) UsedRange()(*i3a815a5e93409d48337d062896c14d9f554d833a002aea7d42ae22a33f331202.UsedRangeRequestBuilder) {
    return i3a815a5e93409d48337d062896c14d9f554d833a002aea7d42ae22a33f331202.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*i53d1008f6f0011645bea4539f658b20af47a15df7aa7cf8190c438d244bd4d53.UsedRangeWithValuesOnlyRequestBuilder) {
    return i53d1008f6f0011645bea4539f658b20af47a15df7aa7cf8190c438d244bd4d53.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
