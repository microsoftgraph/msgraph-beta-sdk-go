package parentsectiongroup

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i130515351a0e0d20a0cc98ed571d8723202901336f998ff5ff63da1b280c0f29 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item/parentsectiongroup/sectiongroups"
    ib0fe6433726789c12c8dc31920842ac34956dac606b755625fb43874320e9d03 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item/parentsectiongroup/parentnotebook"
    ibfd52f71cfb83f683424817541d7407f04ebee5d188d16ae862c65c796727179 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item/parentsectiongroup/sections"
    ib7d69c572e775b495d698bd5c5f51dd8d3a4785d3b68a6d33d7bf08a368f9776 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item/parentsectiongroup/sectiongroups/item"
    ie975c66f40c0a73b8d44a0e5f7eaa2d5d7ad9ed7396d73c6f4f68cb8b0e4e5d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item/parentsectiongroup/sections/item"
)

type ParentSectionGroupRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ParentSectionGroupRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewParentSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionGroupRequestBuilder) {
    m := &ParentSectionGroupRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/sites/{site_id}/onenote/sections/{onenoteSection_id}/parentSectionGroup{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewParentSectionGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentSectionGroupRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ParentSectionGroupRequestBuilder) CreateGetRequestInformation(q func (value *ParentSectionGroupRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ParentSectionGroupRequestBuilderGetQueryParameters)
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
func (m *ParentSectionGroupRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ParentSectionGroupRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ParentSectionGroupRequestBuilder) Get(q func (value *ParentSectionGroupRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSectionGroup() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup), nil
}
func (m *ParentSectionGroupRequestBuilder) ParentNotebook()(*ib0fe6433726789c12c8dc31920842ac34956dac606b755625fb43874320e9d03.ParentNotebookRequestBuilder) {
    return ib0fe6433726789c12c8dc31920842ac34956dac606b755625fb43874320e9d03.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) ParentSectionGroup()(*ParentSectionGroupRequestBuilder) {
    return NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ParentSectionGroupRequestBuilder) SectionGroups()(*i130515351a0e0d20a0cc98ed571d8723202901336f998ff5ff63da1b280c0f29.SectionGroupsRequestBuilder) {
    return i130515351a0e0d20a0cc98ed571d8723202901336f998ff5ff63da1b280c0f29.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) SectionGroupsById(id string)(*ib7d69c572e775b495d698bd5c5f51dd8d3a4785d3b68a6d33d7bf08a368f9776.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return ib7d69c572e775b495d698bd5c5f51dd8d3a4785d3b68a6d33d7bf08a368f9776.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) Sections()(*ibfd52f71cfb83f683424817541d7407f04ebee5d188d16ae862c65c796727179.SectionsRequestBuilder) {
    return ibfd52f71cfb83f683424817541d7407f04ebee5d188d16ae862c65c796727179.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) SectionsById(id string)(*ie975c66f40c0a73b8d44a0e5f7eaa2d5d7ad9ed7396d73c6f4f68cb8b0e4e5d4.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id1"] = id
    }
    return ie975c66f40c0a73b8d44a0e5f7eaa2d5d7ad9ed7396d73c6f4f68cb8b0e4e5d4.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
