package filter

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i44c6588cf26da1e68a4634a834ae4938cca8ab8b751a67c8e48ef1a7ef5a8d2a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applyvaluesfilter"
    i5e4856f4bcc593bc95f953bfa49d9e858fefda4d2747fe6f8605c568557c46ae "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/apply"
    i6ad184758517d00a231e073e7037919e70318e2d5f2790e18bc0196e6f789098 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applyiconfilter"
    i7870023e4b4bfbfa091b54f3bb191f42406795f16210a8046c1e0491edbeb58c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applydynamicfilter"
    i7b08a838608247a8cf35a574cb3c468b09eeaf667422e9aea50cd855e992303c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applybottomitemsfilter"
    i92993877462089826266144acdcde4980c2fdd34080f8f3bbaa7d4811b9d4763 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applytopitemsfilter"
    ice061b6e296893401a730d6293266c707fe3012c26e0f244ecbd9d6e792348a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applytoppercentfilter"
    id202dbc6b73bc5a3ad6102c702c24aac9722ada39a7aa2a4cc7c7cd6da3407ef "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applycellcolorfilter"
    id9e15bad6e59b020a99c689019eeee994c6e83ad95eb39ef4c8dd788aa8b0efb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applybottompercentfilter"
    ie5c848ec01ff69ad7a2b44c7b07f673a570dc9b8ab37d7b2b8de3de84b61b083 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applycustomfilter"
    if2add694d98f90db4e6e200c3a76d29f4f834a6a0b171290d344e2d7c9305d4d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applyfontcolorfilter"
    if6196d732d03f2ed7af679f05c76c34f4c76c53e13f020d7aa40178a413ef3ee "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/clear"
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
func (m *FilterRequestBuilder) Apply()(*i5e4856f4bcc593bc95f953bfa49d9e858fefda4d2747fe6f8605c568557c46ae.ApplyRequestBuilder) {
    return i5e4856f4bcc593bc95f953bfa49d9e858fefda4d2747fe6f8605c568557c46ae.NewApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomItemsFilter()(*i7b08a838608247a8cf35a574cb3c468b09eeaf667422e9aea50cd855e992303c.ApplyBottomItemsFilterRequestBuilder) {
    return i7b08a838608247a8cf35a574cb3c468b09eeaf667422e9aea50cd855e992303c.NewApplyBottomItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomPercentFilter()(*id9e15bad6e59b020a99c689019eeee994c6e83ad95eb39ef4c8dd788aa8b0efb.ApplyBottomPercentFilterRequestBuilder) {
    return id9e15bad6e59b020a99c689019eeee994c6e83ad95eb39ef4c8dd788aa8b0efb.NewApplyBottomPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCellColorFilter()(*id202dbc6b73bc5a3ad6102c702c24aac9722ada39a7aa2a4cc7c7cd6da3407ef.ApplyCellColorFilterRequestBuilder) {
    return id202dbc6b73bc5a3ad6102c702c24aac9722ada39a7aa2a4cc7c7cd6da3407ef.NewApplyCellColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCustomFilter()(*ie5c848ec01ff69ad7a2b44c7b07f673a570dc9b8ab37d7b2b8de3de84b61b083.ApplyCustomFilterRequestBuilder) {
    return ie5c848ec01ff69ad7a2b44c7b07f673a570dc9b8ab37d7b2b8de3de84b61b083.NewApplyCustomFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyDynamicFilter()(*i7870023e4b4bfbfa091b54f3bb191f42406795f16210a8046c1e0491edbeb58c.ApplyDynamicFilterRequestBuilder) {
    return i7870023e4b4bfbfa091b54f3bb191f42406795f16210a8046c1e0491edbeb58c.NewApplyDynamicFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyFontColorFilter()(*if2add694d98f90db4e6e200c3a76d29f4f834a6a0b171290d344e2d7c9305d4d.ApplyFontColorFilterRequestBuilder) {
    return if2add694d98f90db4e6e200c3a76d29f4f834a6a0b171290d344e2d7c9305d4d.NewApplyFontColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyIconFilter()(*i6ad184758517d00a231e073e7037919e70318e2d5f2790e18bc0196e6f789098.ApplyIconFilterRequestBuilder) {
    return i6ad184758517d00a231e073e7037919e70318e2d5f2790e18bc0196e6f789098.NewApplyIconFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopItemsFilter()(*i92993877462089826266144acdcde4980c2fdd34080f8f3bbaa7d4811b9d4763.ApplyTopItemsFilterRequestBuilder) {
    return i92993877462089826266144acdcde4980c2fdd34080f8f3bbaa7d4811b9d4763.NewApplyTopItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopPercentFilter()(*ice061b6e296893401a730d6293266c707fe3012c26e0f244ecbd9d6e792348a6.ApplyTopPercentFilterRequestBuilder) {
    return ice061b6e296893401a730d6293266c707fe3012c26e0f244ecbd9d6e792348a6.NewApplyTopPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyValuesFilter()(*i44c6588cf26da1e68a4634a834ae4938cca8ab8b751a67c8e48ef1a7ef5a8d2a.ApplyValuesFilterRequestBuilder) {
    return i44c6588cf26da1e68a4634a834ae4938cca8ab8b751a67c8e48ef1a7ef5a8d2a.NewApplyValuesFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) Clear()(*if6196d732d03f2ed7af679f05c76c34f4c76c53e13f020d7aa40178a413ef3ee.ClearRequestBuilder) {
    return if6196d732d03f2ed7af679f05c76c34f4c76c53e13f020d7aa40178a413ef3ee.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterRequestBuilder) {
    m := &FilterRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/columns/{workbookTableColumn_id}/filter{?select,expand}";
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
