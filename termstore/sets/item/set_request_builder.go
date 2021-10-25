package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i08ca4c5bb1fe453cf948a9337ebc77517c10861b960c120f6a7bdf151e23091e "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/children"
    ib4bf2188d376329fec86f06954eedc5c031d9700786bc5e49c5d813183216a29 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/terms"
    ibdb574ef03f0eec0fc47939c5b968e4c54cff1b5a2960cb579f607fb551d04c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/relations"
    icc7c8df3b795dc69afaa4bb6bdd0758e60da8070cbb794629c98ab73eaeb21f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/parentgroup"
    i1fdd91eaeb77b5232aaa6637b51c8335e644143a6bec1447e4639d345a19ce09 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/terms/item"
    ia9e638bcb5f0cac6ab821e69fff2fbadde36cbb4ae8162280da0e79f201144d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/children/item"
    icad4fc6304f6da07b148cc11a131410e87779371a7fe5d2470d5b412287bc38d "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/relations/item"
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
func (m *SetRequestBuilder) Children()(*i08ca4c5bb1fe453cf948a9337ebc77517c10861b960c120f6a7bdf151e23091e.ChildrenRequestBuilder) {
    return i08ca4c5bb1fe453cf948a9337ebc77517c10861b960c120f6a7bdf151e23091e.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SetRequestBuilder) ChildrenById(id string)(*ia9e638bcb5f0cac6ab821e69fff2fbadde36cbb4ae8162280da0e79f201144d2.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return ia9e638bcb5f0cac6ab821e69fff2fbadde36cbb4ae8162280da0e79f201144d2.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    m := &SetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/termStore/sets/{set_id}{?select,expand}";
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
func (m *SetRequestBuilder) ParentGroup()(*icc7c8df3b795dc69afaa4bb6bdd0758e60da8070cbb794629c98ab73eaeb21f6.ParentGroupRequestBuilder) {
    return icc7c8df3b795dc69afaa4bb6bdd0758e60da8070cbb794629c98ab73eaeb21f6.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *SetRequestBuilder) Relations()(*ibdb574ef03f0eec0fc47939c5b968e4c54cff1b5a2960cb579f607fb551d04c0.RelationsRequestBuilder) {
    return ibdb574ef03f0eec0fc47939c5b968e4c54cff1b5a2960cb579f607fb551d04c0.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SetRequestBuilder) RelationsById(id string)(*icad4fc6304f6da07b148cc11a131410e87779371a7fe5d2470d5b412287bc38d.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return icad4fc6304f6da07b148cc11a131410e87779371a7fe5d2470d5b412287bc38d.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SetRequestBuilder) Terms()(*ib4bf2188d376329fec86f06954eedc5c031d9700786bc5e49c5d813183216a29.TermsRequestBuilder) {
    return ib4bf2188d376329fec86f06954eedc5c031d9700786bc5e49c5d813183216a29.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SetRequestBuilder) TermsById(id string)(*i1fdd91eaeb77b5232aaa6637b51c8335e644143a6bec1447e4639d345a19ce09.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i1fdd91eaeb77b5232aaa6637b51c8335e644143a6bec1447e4639d345a19ce09.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
