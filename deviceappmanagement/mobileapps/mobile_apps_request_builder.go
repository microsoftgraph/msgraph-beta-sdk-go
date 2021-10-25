package mobileapps

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/haspayloadlinks"
    i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/getmobileappcountwithstatus"
    i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/gettopmobileappswithstatuswithcount"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/validatexml"
)

type MobileAppsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type MobileAppsRequestBuilderGetQueryParameters struct {
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
func NewMobileAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppsRequestBuilder) {
    m := &MobileAppsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/mobileApps{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewMobileAppsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MobileAppsRequestBuilder) CreateGetRequestInformation(q func (value *MobileAppsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(MobileAppsRequestBuilderGetQueryParameters)
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
func (m *MobileAppsRequestBuilder) CreatePostRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MobileAppsRequestBuilder) Get(q func (value *MobileAppsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*MobileAppsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*MobileAppsResponse), nil
}
func (m *MobileAppsRequestBuilder) GetMobileAppCountWithStatus(status *string)(*i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98.GetMobileAppCountWithStatusRequestBuilder) {
    return i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98.NewGetMobileAppCountWithStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter, status);
}
func (m *MobileAppsRequestBuilder) GetTopMobileAppsWithStatusWithCount(status *string, count *int64)(*i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200.GetTopMobileAppsWithStatusWithCountRequestBuilder) {
    return i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200.NewGetTopMobileAppsWithStatusWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, status, count);
}
func (m *MobileAppsRequestBuilder) HasPayloadLinks()(*i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944.HasPayloadLinksRequestBuilder) {
    return i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944.NewHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppsRequestBuilder) Post(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMobileApp() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp), nil
}
func (m *MobileAppsRequestBuilder) ValidateXml()(*id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe.ValidateXmlRequestBuilder) {
    return id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe.NewValidateXmlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
