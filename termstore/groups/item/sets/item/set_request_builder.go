package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i30e68a2f75c4e30c38b002a2b0f43d8b9324529f09cc859a02fd8d1dae1eb77d "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/relations"
    ibb025a317b0b0e25816729d786d646eba813f51a9af8109606f2c49c82cafd92 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/parentgroup"
    ic9b3fe35558f7761d58013dc735f0a15da655446a5017fdbd92ff01ca57d26af "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/children"
    ifdfd37b1ff263ab28192d6594d1e506db019fcfd499c0da8c7692e5d1cbd9555 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms"
    i423c1de361babb93c98ec43e7cd413b168af956c9a1d9bfdfa989c00950aab25 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/terms/item"
    i7e290f56cfea106212ddd4be05db94843ffd23d18ae81173f668737a2b25e6c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/children/item"
    ia0bd748b4766ea2d50956a69c4675d9c4ca726af0e0f7ecb571880e7120ddf22 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/item/sets/item/relations/item"
)

type SetRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SetRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *SetRequestBuilder) Children()(*ic9b3fe35558f7761d58013dc735f0a15da655446a5017fdbd92ff01ca57d26af.ChildrenRequestBuilder) {
    return ic9b3fe35558f7761d58013dc735f0a15da655446a5017fdbd92ff01ca57d26af.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SetRequestBuilder) ChildrenById(id string)(*i7e290f56cfea106212ddd4be05db94843ffd23d18ae81173f668737a2b25e6c3.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i7e290f56cfea106212ddd4be05db94843ffd23d18ae81173f668737a2b25e6c3.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    m := &SetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/termStore/groups/{group_id}/sets/{set_id}{?select,expand}";
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
func NewSetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SetRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SetRequestBuilder) CreateGetRequestInformation(q func (value *SetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SetRequestBuilderGetQueryParameters)
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
func (m *SetRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Set, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SetRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SetRequestBuilder) Get(q func (value *SetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Set, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSet() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Set), nil
}
func (m *SetRequestBuilder) ParentGroup()(*ibb025a317b0b0e25816729d786d646eba813f51a9af8109606f2c49c82cafd92.ParentGroupRequestBuilder) {
    return ibb025a317b0b0e25816729d786d646eba813f51a9af8109606f2c49c82cafd92.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SetRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Set, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SetRequestBuilder) Relations()(*i30e68a2f75c4e30c38b002a2b0f43d8b9324529f09cc859a02fd8d1dae1eb77d.RelationsRequestBuilder) {
    return i30e68a2f75c4e30c38b002a2b0f43d8b9324529f09cc859a02fd8d1dae1eb77d.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SetRequestBuilder) RelationsById(id string)(*ia0bd748b4766ea2d50956a69c4675d9c4ca726af0e0f7ecb571880e7120ddf22.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return ia0bd748b4766ea2d50956a69c4675d9c4ca726af0e0f7ecb571880e7120ddf22.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SetRequestBuilder) Terms()(*ifdfd37b1ff263ab28192d6594d1e506db019fcfd499c0da8c7692e5d1cbd9555.TermsRequestBuilder) {
    return ifdfd37b1ff263ab28192d6594d1e506db019fcfd499c0da8c7692e5d1cbd9555.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SetRequestBuilder) TermsById(id string)(*i423c1de361babb93c98ec43e7cd413b168af956c9a1d9bfdfa989c00950aab25.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i423c1de361babb93c98ec43e7cd413b168af956c9a1d9bfdfa989c00950aab25.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
