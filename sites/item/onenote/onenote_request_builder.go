package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    iaa664a82ae9fffb3b2512033e7c1b2d12af53c401334b545cafc70c8e036b62d "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages"
    ibfb1096a4e9af5ebda9b7adc8df22c6b1585d4223f132d3d8d5d7acccef4f64b "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/operations"
    ic900129d3b82e90100ca8cbca83dabb8d47a3215c619636a2207492943c84948 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups"
    idfec715ee87619edf143beb50b903e0f6e13db5fce25ebb7d844873504b2a198 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/notebooks"
    ied650362e2081744707f87dc144030142278c0ea808a1bfa9ca9138067ae2f93 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections"
    ieefe6ead6bc64ab8fc788a0eae1bbb4b1da67d79b5756271266acffdf2c1a2c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/resources"
    i162a80b6b9a759989691d4208d5edb74a978dccfe02475d9bfdf09e730886189 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/resources/item"
    i2d1af23696cae1167a969e48e382916eb69ee04dddd2d084cac2bbd758295ef3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/notebooks/item"
    i5911c0075c684f98fcf33c2566f164c22aa659ca4d9a0c2366a6a7078d6bf4aa "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item"
    i69abb64a0d2c33207d2a4dc4d47ee4ff558142af3af7154bc8f0630b6683dcf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item"
    i8962b2ca94b13c5fa106579d43d7572e68c3c0db2309a183bacf4c839def4c72 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/operations/item"
    i8e1701b9e2b09e4f98248fe7d70ff436c05c4d12eca781e8170d8a0a2c3164d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups/item"
)

type OnenoteRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OnenoteRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/sites/{site_id}/onenote{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(q func (value *OnenoteRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OnenoteRequestBuilderGetQueryParameters)
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
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OnenoteRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OnenoteRequestBuilder) Get(q func (value *OnenoteRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenote() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote), nil
}
func (m *OnenoteRequestBuilder) Notebooks()(*idfec715ee87619edf143beb50b903e0f6e13db5fce25ebb7d844873504b2a198.NotebooksRequestBuilder) {
    return idfec715ee87619edf143beb50b903e0f6e13db5fce25ebb7d844873504b2a198.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i2d1af23696cae1167a969e48e382916eb69ee04dddd2d084cac2bbd758295ef3.NotebookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return i2d1af23696cae1167a969e48e382916eb69ee04dddd2d084cac2bbd758295ef3.NewNotebookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*ibfb1096a4e9af5ebda9b7adc8df22c6b1585d4223f132d3d8d5d7acccef4f64b.OperationsRequestBuilder) {
    return ibfb1096a4e9af5ebda9b7adc8df22c6b1585d4223f132d3d8d5d7acccef4f64b.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i8962b2ca94b13c5fa106579d43d7572e68c3c0db2309a183bacf4c839def4c72.OnenoteOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i8962b2ca94b13c5fa106579d43d7572e68c3c0db2309a183bacf4c839def4c72.NewOnenoteOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*iaa664a82ae9fffb3b2512033e7c1b2d12af53c401334b545cafc70c8e036b62d.PagesRequestBuilder) {
    return iaa664a82ae9fffb3b2512033e7c1b2d12af53c401334b545cafc70c8e036b62d.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) PagesById(id string)(*i69abb64a0d2c33207d2a4dc4d47ee4ff558142af3af7154bc8f0630b6683dcf7.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i69abb64a0d2c33207d2a4dc4d47ee4ff558142af3af7154bc8f0630b6683dcf7.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OnenoteRequestBuilder) Resources()(*ieefe6ead6bc64ab8fc788a0eae1bbb4b1da67d79b5756271266acffdf2c1a2c3.ResourcesRequestBuilder) {
    return ieefe6ead6bc64ab8fc788a0eae1bbb4b1da67d79b5756271266acffdf2c1a2c3.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*i162a80b6b9a759989691d4208d5edb74a978dccfe02475d9bfdf09e730886189.OnenoteResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return i162a80b6b9a759989691d4208d5edb74a978dccfe02475d9bfdf09e730886189.NewOnenoteResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*ic900129d3b82e90100ca8cbca83dabb8d47a3215c619636a2207492943c84948.SectionGroupsRequestBuilder) {
    return ic900129d3b82e90100ca8cbca83dabb8d47a3215c619636a2207492943c84948.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i8e1701b9e2b09e4f98248fe7d70ff436c05c4d12eca781e8170d8a0a2c3164d5.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i8e1701b9e2b09e4f98248fe7d70ff436c05c4d12eca781e8170d8a0a2c3164d5.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*ied650362e2081744707f87dc144030142278c0ea808a1bfa9ca9138067ae2f93.SectionsRequestBuilder) {
    return ied650362e2081744707f87dc144030142278c0ea808a1bfa9ca9138067ae2f93.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i5911c0075c684f98fcf33c2566f164c22aa659ca4d9a0c2366a6a7078d6bf4aa.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i5911c0075c684f98fcf33c2566f164c22aa659ca4d9a0c2366a6a7078d6bf4aa.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
