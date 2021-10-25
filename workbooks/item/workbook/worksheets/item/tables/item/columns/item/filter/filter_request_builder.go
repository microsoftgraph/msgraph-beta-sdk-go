package filter

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0ea8b50d436086366f1769879cee0523aa44a37817d3197aaed6d3cf3106ef88 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applyfontcolorfilter"
    i15870e5322e43d990a69cbe697c73b40eeb33ea1efe5b9e0ca8b98e9cd72ee59 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applytoppercentfilter"
    i1cddcb4acebc500eca65aa379cb89ede1dcbedfb3bec9dd882f230ef16fd380f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/apply"
    i67c164169d88195e5990622742cfb34895bd62a74f585d33409469201798da5c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applycellcolorfilter"
    ia28e741c91d5da42e45c95fb8b102c3e065d65d6080f362f887c5e6afd262bdc "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applyiconfilter"
    ic61bec797d91e15fc8b689b919869786a31c3d528f162eb53bee13fb9bdda57c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/clear"
    icbe146a2abfecbdb199afa6312213458d17ab85ca044eb727ff6e3c7c303b8f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applyvaluesfilter"
    ida9e1f9bc341e752fcddb7d3a81af30bd2b4282d36408c6c26ce3d725d136d3c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applytopitemsfilter"
    idb4dbb70f927f02720b89e45928fa0b3f2346cb0d991f1aa257f09379ace36ac "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applybottompercentfilter"
    idf68ab0e88b09f94c9fae16011578c96a5ca751c0b086530371a49ba92dbb643 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applycustomfilter"
    ie7d6d7a3e649368d5297f1ac3bb03c92a89b09ba096987badb90dfc04c472fda "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applybottomitemsfilter"
    ieb889e057c7c902656edfa8c2c5569d3f8b75912e13feb175725cd9b2a73a4bb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applydynamicfilter"
)

type FilterRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type FilterRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *FilterRequestBuilder) Apply()(*i1cddcb4acebc500eca65aa379cb89ede1dcbedfb3bec9dd882f230ef16fd380f.ApplyRequestBuilder) {
    return i1cddcb4acebc500eca65aa379cb89ede1dcbedfb3bec9dd882f230ef16fd380f.NewApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomItemsFilter()(*ie7d6d7a3e649368d5297f1ac3bb03c92a89b09ba096987badb90dfc04c472fda.ApplyBottomItemsFilterRequestBuilder) {
    return ie7d6d7a3e649368d5297f1ac3bb03c92a89b09ba096987badb90dfc04c472fda.NewApplyBottomItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomPercentFilter()(*idb4dbb70f927f02720b89e45928fa0b3f2346cb0d991f1aa257f09379ace36ac.ApplyBottomPercentFilterRequestBuilder) {
    return idb4dbb70f927f02720b89e45928fa0b3f2346cb0d991f1aa257f09379ace36ac.NewApplyBottomPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCellColorFilter()(*i67c164169d88195e5990622742cfb34895bd62a74f585d33409469201798da5c.ApplyCellColorFilterRequestBuilder) {
    return i67c164169d88195e5990622742cfb34895bd62a74f585d33409469201798da5c.NewApplyCellColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCustomFilter()(*idf68ab0e88b09f94c9fae16011578c96a5ca751c0b086530371a49ba92dbb643.ApplyCustomFilterRequestBuilder) {
    return idf68ab0e88b09f94c9fae16011578c96a5ca751c0b086530371a49ba92dbb643.NewApplyCustomFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyDynamicFilter()(*ieb889e057c7c902656edfa8c2c5569d3f8b75912e13feb175725cd9b2a73a4bb.ApplyDynamicFilterRequestBuilder) {
    return ieb889e057c7c902656edfa8c2c5569d3f8b75912e13feb175725cd9b2a73a4bb.NewApplyDynamicFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyFontColorFilter()(*i0ea8b50d436086366f1769879cee0523aa44a37817d3197aaed6d3cf3106ef88.ApplyFontColorFilterRequestBuilder) {
    return i0ea8b50d436086366f1769879cee0523aa44a37817d3197aaed6d3cf3106ef88.NewApplyFontColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyIconFilter()(*ia28e741c91d5da42e45c95fb8b102c3e065d65d6080f362f887c5e6afd262bdc.ApplyIconFilterRequestBuilder) {
    return ia28e741c91d5da42e45c95fb8b102c3e065d65d6080f362f887c5e6afd262bdc.NewApplyIconFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopItemsFilter()(*ida9e1f9bc341e752fcddb7d3a81af30bd2b4282d36408c6c26ce3d725d136d3c.ApplyTopItemsFilterRequestBuilder) {
    return ida9e1f9bc341e752fcddb7d3a81af30bd2b4282d36408c6c26ce3d725d136d3c.NewApplyTopItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopPercentFilter()(*i15870e5322e43d990a69cbe697c73b40eeb33ea1efe5b9e0ca8b98e9cd72ee59.ApplyTopPercentFilterRequestBuilder) {
    return i15870e5322e43d990a69cbe697c73b40eeb33ea1efe5b9e0ca8b98e9cd72ee59.NewApplyTopPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyValuesFilter()(*icbe146a2abfecbdb199afa6312213458d17ab85ca044eb727ff6e3c7c303b8f3.ApplyValuesFilterRequestBuilder) {
    return icbe146a2abfecbdb199afa6312213458d17ab85ca044eb727ff6e3c7c303b8f3.NewApplyValuesFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) Clear()(*ic61bec797d91e15fc8b689b919869786a31c3d528f162eb53bee13fb9bdda57c.ClearRequestBuilder) {
    return ic61bec797d91e15fc8b689b919869786a31c3d528f162eb53bee13fb9bdda57c.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterRequestBuilder) {
    m := &FilterRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/tables/{workbookTable_id}/columns/{workbookTableColumn_id}/filter{?select,expand}";
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
func NewFilterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilterRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *FilterRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *FilterRequestBuilder) CreateGetRequestInformation(q func (value *FilterRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(FilterRequestBuilderGetQueryParameters)
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
func (m *FilterRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookFilter, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *FilterRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *FilterRequestBuilder) Get(q func (value *FilterRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookFilter, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookFilter() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookFilter), nil
}
func (m *FilterRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookFilter, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
