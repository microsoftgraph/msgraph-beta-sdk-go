package tiindicators

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5fd228a8811b3ee0e602ce99917149a6ef997b242fa74579dc28eb98aeff4610 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/updatetiindicators"
    i84faad3a80813a60a94e3416bfa220426811428a0f17f7a125a8a8389a71931a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/deletetiindicatorsbyexternalid"
    id9554bf1b52f05e494c0e5962b41bb42f06710f09ff2286b921260ce9ef0c431 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/submittiindicators"
    ie38c9a1118b781a8cc486e08009768123690676d39af853ebdf06098c5110f09 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/deletetiindicators"
)

type TiIndicatorsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type TiIndicatorsRequestBuilderGetQueryParameters struct {
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
func NewTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TiIndicatorsRequestBuilder) {
    m := &TiIndicatorsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/security/tiIndicators{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewTiIndicatorsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TiIndicatorsRequestBuilder) CreateGetRequestInformation(q func (value *TiIndicatorsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TiIndicatorsRequestBuilderGetQueryParameters)
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
func (m *TiIndicatorsRequestBuilder) CreatePostRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TiIndicator, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TiIndicatorsRequestBuilder) DeleteTiIndicators()(*ie38c9a1118b781a8cc486e08009768123690676d39af853ebdf06098c5110f09.DeleteTiIndicatorsRequestBuilder) {
    return ie38c9a1118b781a8cc486e08009768123690676d39af853ebdf06098c5110f09.NewDeleteTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TiIndicatorsRequestBuilder) DeleteTiIndicatorsByExternalId()(*i84faad3a80813a60a94e3416bfa220426811428a0f17f7a125a8a8389a71931a.DeleteTiIndicatorsByExternalIdRequestBuilder) {
    return i84faad3a80813a60a94e3416bfa220426811428a0f17f7a125a8a8389a71931a.NewDeleteTiIndicatorsByExternalIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TiIndicatorsRequestBuilder) Get(q func (value *TiIndicatorsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*TiIndicatorsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTiIndicatorsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*TiIndicatorsResponse), nil
}
func (m *TiIndicatorsRequestBuilder) Post(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TiIndicator, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TiIndicator, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTiIndicator() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TiIndicator), nil
}
func (m *TiIndicatorsRequestBuilder) SubmitTiIndicators()(*id9554bf1b52f05e494c0e5962b41bb42f06710f09ff2286b921260ce9ef0c431.SubmitTiIndicatorsRequestBuilder) {
    return id9554bf1b52f05e494c0e5962b41bb42f06710f09ff2286b921260ce9ef0c431.NewSubmitTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TiIndicatorsRequestBuilder) UpdateTiIndicators()(*i5fd228a8811b3ee0e602ce99917149a6ef997b242fa74579dc28eb98aeff4610.UpdateTiIndicatorsRequestBuilder) {
    return i5fd228a8811b3ee0e602ce99917149a6ef997b242fa74579dc28eb98aeff4610.NewUpdateTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
