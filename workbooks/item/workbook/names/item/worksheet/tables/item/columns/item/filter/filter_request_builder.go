package filter

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0708cb60eb598d4f7e45b8d54b4169444570e94b962f522374b9081152df6284 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applybottomitemsfilter"
    i122b873076d92fce27ca7176482c76cddc11d3b1d95d6d1714a6d3b72a246f1b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applyfontcolorfilter"
    i13554509e84e997bce5cca4d6668ec3212de0460948333e3ce2170953377c2e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applyiconfilter"
    i1faa3dd74ea2405101fe43803ad17c1d02cd6fbdcca84b870a3fc1591457aa87 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applycellcolorfilter"
    i20d62e31d32eaf90370351dc10772c1cec2093bb104814f5da5a3f69bda91d68 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/apply"
    i578e1560d38b6a7568f4bc6348789e9f4e7e2180d078fe38f50735bae9425ce0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applydynamicfilter"
    i588704fba3b4ddd6f027079ff23afd3d1eaaca3e3eb951c4526cc931acdd747f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applycustomfilter"
    id4f94aa07b7bdad74ae42e0aeffeefd5cb9c4ae39427de3a8f51db824df9c864 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/clear"
    ie0941d0f2386b6479dc7542411eef885b3b04ec765398837d4730ce817d3b8b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applyvaluesfilter"
    ie2b1aa51169ce07dae83afb06d9180fe00fe1314881b77eee0c517b0af19f24d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applybottompercentfilter"
    iefe0da864354086f85c5eecf974c2697d991ff7609eb04bebbf6c3594f688687 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applytoppercentfilter"
    ifdc8e9d73933c5f63340ac59b99c168e16a35fce1bd121d75b61e6cc81c7fcf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applytopitemsfilter"
)

type FilterRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type FilterRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *FilterRequestBuilder) Apply()(*i20d62e31d32eaf90370351dc10772c1cec2093bb104814f5da5a3f69bda91d68.ApplyRequestBuilder) {
    return i20d62e31d32eaf90370351dc10772c1cec2093bb104814f5da5a3f69bda91d68.NewApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomItemsFilter()(*i0708cb60eb598d4f7e45b8d54b4169444570e94b962f522374b9081152df6284.ApplyBottomItemsFilterRequestBuilder) {
    return i0708cb60eb598d4f7e45b8d54b4169444570e94b962f522374b9081152df6284.NewApplyBottomItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomPercentFilter()(*ie2b1aa51169ce07dae83afb06d9180fe00fe1314881b77eee0c517b0af19f24d.ApplyBottomPercentFilterRequestBuilder) {
    return ie2b1aa51169ce07dae83afb06d9180fe00fe1314881b77eee0c517b0af19f24d.NewApplyBottomPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCellColorFilter()(*i1faa3dd74ea2405101fe43803ad17c1d02cd6fbdcca84b870a3fc1591457aa87.ApplyCellColorFilterRequestBuilder) {
    return i1faa3dd74ea2405101fe43803ad17c1d02cd6fbdcca84b870a3fc1591457aa87.NewApplyCellColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCustomFilter()(*i588704fba3b4ddd6f027079ff23afd3d1eaaca3e3eb951c4526cc931acdd747f.ApplyCustomFilterRequestBuilder) {
    return i588704fba3b4ddd6f027079ff23afd3d1eaaca3e3eb951c4526cc931acdd747f.NewApplyCustomFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyDynamicFilter()(*i578e1560d38b6a7568f4bc6348789e9f4e7e2180d078fe38f50735bae9425ce0.ApplyDynamicFilterRequestBuilder) {
    return i578e1560d38b6a7568f4bc6348789e9f4e7e2180d078fe38f50735bae9425ce0.NewApplyDynamicFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyFontColorFilter()(*i122b873076d92fce27ca7176482c76cddc11d3b1d95d6d1714a6d3b72a246f1b.ApplyFontColorFilterRequestBuilder) {
    return i122b873076d92fce27ca7176482c76cddc11d3b1d95d6d1714a6d3b72a246f1b.NewApplyFontColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyIconFilter()(*i13554509e84e997bce5cca4d6668ec3212de0460948333e3ce2170953377c2e9.ApplyIconFilterRequestBuilder) {
    return i13554509e84e997bce5cca4d6668ec3212de0460948333e3ce2170953377c2e9.NewApplyIconFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopItemsFilter()(*ifdc8e9d73933c5f63340ac59b99c168e16a35fce1bd121d75b61e6cc81c7fcf9.ApplyTopItemsFilterRequestBuilder) {
    return ifdc8e9d73933c5f63340ac59b99c168e16a35fce1bd121d75b61e6cc81c7fcf9.NewApplyTopItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopPercentFilter()(*iefe0da864354086f85c5eecf974c2697d991ff7609eb04bebbf6c3594f688687.ApplyTopPercentFilterRequestBuilder) {
    return iefe0da864354086f85c5eecf974c2697d991ff7609eb04bebbf6c3594f688687.NewApplyTopPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyValuesFilter()(*ie0941d0f2386b6479dc7542411eef885b3b04ec765398837d4730ce817d3b8b0.ApplyValuesFilterRequestBuilder) {
    return ie0941d0f2386b6479dc7542411eef885b3b04ec765398837d4730ce817d3b8b0.NewApplyValuesFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) Clear()(*id4f94aa07b7bdad74ae42e0aeffeefd5cb9c4ae39427de3a8f51db824df9c864.ClearRequestBuilder) {
    return id4f94aa07b7bdad74ae42e0aeffeefd5cb9c4ae39427de3a8f51db824df9c864.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterRequestBuilder) {
    m := &FilterRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/tables/{workbookTable_id}/columns/{workbookTableColumn_id}/filter{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
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
