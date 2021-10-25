package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0aca550ba17eac0067818817645f66d02a2e7e6318e913a5e0d42924abac01d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/datalabels"
    i26540da21c5f8eb6b9a4afee7461d292b94586e8b133b139818f7cd95c002255 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/setdata"
    i31b926be88a22210ebdbd1d797ceb2374bd98f7356bffee6a84531aca64b035d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/worksheet"
    i47a26552b54256d90e60f5f60f6a5f57a4678a7d006de8586993e4fc090dbb3a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/title"
    ia1ee182da07d37c6045e9b77df600df2d0ee1949c25958f2b9e79a5b33d4066f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/imagewithwidthwithheight"
    ia94c2cc6ebe063ffa72fb4aa6f20a55c594093733f5f07c6bf48a07ab114a488 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/imagewithwidthwithheightwithfittingmode"
    ic69e93b4975cdc4aac5a8223e30cbc2c91ca95dc1178ff4c88247fd1b24355cf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/imagewithwidth"
    ic822cf3c298bb549d096d6cd988c1729b0f8f6f4ac1a0c2d88771eea59634146 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/legend"
    ic8aa1abce179d71029215dd44d7ae4480f84dcdfb5d9017811731ee439071639 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/format"
    idf1102374f7022e929a594345386cdb94c314a1357c852e21bca7ea485eab58c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/image"
    ie100fbbbf10c4a310e3efd238c1c2b067cece060dd80cef3e22792a63b0ded3a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/setposition"
    ie917d564747d430dc1fc7e3fb2c2478d6b626b4a2a7d537bcf1b13209fc033eb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/series"
    if9eae210bfc3bda8f41da0b230fca738b32639f02e0d61b43f9097b06ccb1538 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/axes"
    i2d83e5aeb2446afd5be234a0a33ccb45ba36c3a6e39653924b5f21d7813b084e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/series/item"
)

type WorkbookChartRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorkbookChartRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *WorkbookChartRequestBuilder) Axes()(*if9eae210bfc3bda8f41da0b230fca738b32639f02e0d61b43f9097b06ccb1538.AxesRequestBuilder) {
    return if9eae210bfc3bda8f41da0b230fca738b32639f02e0d61b43f9097b06ccb1538.NewAxesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWorkbookChartRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartRequestBuilder) {
    m := &WorkbookChartRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/charts/{workbookChart_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewWorkbookChartRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookChartRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookChartRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookChartRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookChartRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookChartRequestBuilderGetQueryParameters)
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
func (m *WorkbookChartRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookChartRequestBuilder) DataLabels()(*i0aca550ba17eac0067818817645f66d02a2e7e6318e913a5e0d42924abac01d2.DataLabelsRequestBuilder) {
    return i0aca550ba17eac0067818817645f66d02a2e7e6318e913a5e0d42924abac01d2.NewDataLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorkbookChartRequestBuilder) Format()(*ic8aa1abce179d71029215dd44d7ae4480f84dcdfb5d9017811731ee439071639.FormatRequestBuilder) {
    return ic8aa1abce179d71029215dd44d7ae4480f84dcdfb5d9017811731ee439071639.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Get(q func (value *WorkbookChartRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookChart() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart), nil
}
func (m *WorkbookChartRequestBuilder) Image()(*idf1102374f7022e929a594345386cdb94c314a1357c852e21bca7ea485eab58c.ImageRequestBuilder) {
    return idf1102374f7022e929a594345386cdb94c314a1357c852e21bca7ea485eab58c.NewImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) ImageWithWidth(width *int32)(*ic69e93b4975cdc4aac5a8223e30cbc2c91ca95dc1178ff4c88247fd1b24355cf.ImageWithWidthRequestBuilder) {
    return ic69e93b4975cdc4aac5a8223e30cbc2c91ca95dc1178ff4c88247fd1b24355cf.NewImageWithWidthRequestBuilderInternal(m.pathParameters, m.requestAdapter, width);
}
func (m *WorkbookChartRequestBuilder) ImageWithWidthWithHeight(width *int32, height *int32)(*ia1ee182da07d37c6045e9b77df600df2d0ee1949c25958f2b9e79a5b33d4066f.ImageWithWidthWithHeightRequestBuilder) {
    return ia1ee182da07d37c6045e9b77df600df2d0ee1949c25958f2b9e79a5b33d4066f.NewImageWithWidthWithHeightRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height);
}
func (m *WorkbookChartRequestBuilder) ImageWithWidthWithHeightWithFittingMode(width *int32, height *int32, fittingMode *string)(*ia94c2cc6ebe063ffa72fb4aa6f20a55c594093733f5f07c6bf48a07ab114a488.ImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return ia94c2cc6ebe063ffa72fb4aa6f20a55c594093733f5f07c6bf48a07ab114a488.NewImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height, fittingMode);
}
func (m *WorkbookChartRequestBuilder) Legend()(*ic822cf3c298bb549d096d6cd988c1729b0f8f6f4ac1a0c2d88771eea59634146.LegendRequestBuilder) {
    return ic822cf3c298bb549d096d6cd988c1729b0f8f6f4ac1a0c2d88771eea59634146.NewLegendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorkbookChartRequestBuilder) Series()(*ie917d564747d430dc1fc7e3fb2c2478d6b626b4a2a7d537bcf1b13209fc033eb.SeriesRequestBuilder) {
    return ie917d564747d430dc1fc7e3fb2c2478d6b626b4a2a7d537bcf1b13209fc033eb.NewSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SeriesById(id string)(*i2d83e5aeb2446afd5be234a0a33ccb45ba36c3a6e39653924b5f21d7813b084e.WorkbookChartSeriesRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChartSeries_id"] = id
    }
    return i2d83e5aeb2446afd5be234a0a33ccb45ba36c3a6e39653924b5f21d7813b084e.NewWorkbookChartSeriesRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SetData()(*i26540da21c5f8eb6b9a4afee7461d292b94586e8b133b139818f7cd95c002255.SetDataRequestBuilder) {
    return i26540da21c5f8eb6b9a4afee7461d292b94586e8b133b139818f7cd95c002255.NewSetDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SetPosition()(*ie100fbbbf10c4a310e3efd238c1c2b067cece060dd80cef3e22792a63b0ded3a.SetPositionRequestBuilder) {
    return ie100fbbbf10c4a310e3efd238c1c2b067cece060dd80cef3e22792a63b0ded3a.NewSetPositionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Title()(*i47a26552b54256d90e60f5f60f6a5f57a4678a7d006de8586993e4fc090dbb3a.TitleRequestBuilder) {
    return i47a26552b54256d90e60f5f60f6a5f57a4678a7d006de8586993e4fc090dbb3a.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Worksheet()(*i31b926be88a22210ebdbd1d797ceb2374bd98f7356bffee6a84531aca64b035d.WorksheetRequestBuilder) {
    return i31b926be88a22210ebdbd1d797ceb2374bd98f7356bffee6a84531aca64b035d.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
