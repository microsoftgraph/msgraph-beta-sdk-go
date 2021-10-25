package administrativeunits

import (
    i277b79a3b313ac8255d78ae90b5e9c683fe411d78cfeb5e1e83d31d38c270e05 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/getbyids"
    ia5a3db4673512570eb5641390e08a082576bca30d7028287114d44432947b385 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/getuserownedobjects"
    ic5b5428ae35afae2a32cefcf30c64f6ab1fbe0e8909656a8a35a1a32357418eb "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/validateproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    if16d334c721262585664df32b0567644d1ab8a2a5c2829d1f520c68c93e0c424 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/delta"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type AdministrativeUnitsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type AdministrativeUnitsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewAdministrativeUnitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitsRequestBuilder) {
    m := &AdministrativeUnitsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/administrativeUnits{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewAdministrativeUnitsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AdministrativeUnitsRequestBuilder) CreateGetRequestInformation(q func (value *AdministrativeUnitsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(AdministrativeUnitsRequestBuilderGetQueryParameters)
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
func (m *AdministrativeUnitsRequestBuilder) CreatePostRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AdministrativeUnitsRequestBuilder) Delta()(*if16d334c721262585664df32b0567644d1ab8a2a5c2829d1f520c68c93e0c424.DeltaRequestBuilder) {
    return if16d334c721262585664df32b0567644d1ab8a2a5c2829d1f520c68c93e0c424.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitsRequestBuilder) Get(q func (value *AdministrativeUnitsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*AdministrativeUnitsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdministrativeUnitsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*AdministrativeUnitsResponse), nil
}
func (m *AdministrativeUnitsRequestBuilder) GetByIds()(*i277b79a3b313ac8255d78ae90b5e9c683fe411d78cfeb5e1e83d31d38c270e05.GetByIdsRequestBuilder) {
    return i277b79a3b313ac8255d78ae90b5e9c683fe411d78cfeb5e1e83d31d38c270e05.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitsRequestBuilder) GetUserOwnedObjects()(*ia5a3db4673512570eb5641390e08a082576bca30d7028287114d44432947b385.GetUserOwnedObjectsRequestBuilder) {
    return ia5a3db4673512570eb5641390e08a082576bca30d7028287114d44432947b385.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitsRequestBuilder) Post(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAdministrativeUnit() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit), nil
}
func (m *AdministrativeUnitsRequestBuilder) ValidateProperties()(*ic5b5428ae35afae2a32cefcf30c64f6ab1fbe0e8909656a8a35a1a32357418eb.ValidatePropertiesRequestBuilder) {
    return ic5b5428ae35afae2a32cefcf30c64f6ab1fbe0e8909656a8a35a1a32357418eb.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
