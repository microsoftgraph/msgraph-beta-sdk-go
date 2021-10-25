package updatableassets

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1ad9e4f03fd41b4ff3469626c2221aacd3c3b097757de111b29a11d0e40e5af7 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/enrollassetsbyid"
    i4dfa7ffb4b8c911e44699c4105daaccddf21f9cbf6d740540587c4fac3cb8746 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/unenrollassets"
    i5e033ef8b15c49341c073fc21c659e5769ead616f1066db183a743d44638c27e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/enrollassets"
    i9b7b56068ae594d878d9164fec1e047566feb9a550ec5e453c176a1f0a5f2796 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/unenrollassetsbyid"
)

type UpdatableAssetsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type UpdatableAssetsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escaped []string;
    Skip *int32;
    Top *int32;
}
func NewUpdatableAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetsRequestBuilder) {
    m := &UpdatableAssetsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/admin/windows/updates/updatableAssets{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewUpdatableAssetsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatableAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UpdatableAssetsRequestBuilder) CreateGetRequestInformation(q func (value *UpdatableAssetsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(UpdatableAssetsRequestBuilderGetQueryParameters)
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
func (m *UpdatableAssetsRequestBuilder) CreatePostRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
func (m *UpdatableAssetsRequestBuilder) EnrollAssets()(*i5e033ef8b15c49341c073fc21c659e5769ead616f1066db183a743d44638c27e.EnrollAssetsRequestBuilder) {
    return i5e033ef8b15c49341c073fc21c659e5769ead616f1066db183a743d44638c27e.NewEnrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetsRequestBuilder) EnrollAssetsById()(*i1ad9e4f03fd41b4ff3469626c2221aacd3c3b097757de111b29a11d0e40e5af7.EnrollAssetsByIdRequestBuilder) {
    return i1ad9e4f03fd41b4ff3469626c2221aacd3c3b097757de111b29a11d0e40e5af7.NewEnrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetsRequestBuilder) Get(q func (value *UpdatableAssetsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*UpdatableAssetsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdatableAssetsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*UpdatableAssetsResponse), nil
}
func (m *UpdatableAssetsRequestBuilder) Post(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUpdatableAsset() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset), nil
}
func (m *UpdatableAssetsRequestBuilder) UnenrollAssets()(*i4dfa7ffb4b8c911e44699c4105daaccddf21f9cbf6d740540587c4fac3cb8746.UnenrollAssetsRequestBuilder) {
    return i4dfa7ffb4b8c911e44699c4105daaccddf21f9cbf6d740540587c4fac3cb8746.NewUnenrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetsRequestBuilder) UnenrollAssetsById()(*i9b7b56068ae594d878d9164fec1e047566feb9a550ec5e453c176a1f0a5f2796.UnenrollAssetsByIdRequestBuilder) {
    return i9b7b56068ae594d878d9164fec1e047566feb9a550ec5e453c176a1f0a5f2796.NewUnenrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
