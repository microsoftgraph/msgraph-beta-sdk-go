package privilegedsignupstatus

import (
    i505eff958ea7946e725489038a9fbe85c297626b944e12d9c6f68697f1a65258 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/issignedup"
    i576c2492be4604b88fea75e47108c8bd8db61a3de84b266733d2d7849bec46eb "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/cansignup"
    ib823f830e3ea0500a214b973cf170985ed37b35d907e2b81e06df8aaa8f66256 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/completesetup"
    ibe660ff6aa6da80e38b282109151294b9409fe46a692bb37484dd8a1705b5a94 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/signup"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type PrivilegedSignupStatusRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PrivilegedSignupStatusRequestBuilderGetQueryParameters struct {
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
func (m *PrivilegedSignupStatusRequestBuilder) CanSignUp()(*i576c2492be4604b88fea75e47108c8bd8db61a3de84b266733d2d7849bec46eb.CanSignUpRequestBuilder) {
    return i576c2492be4604b88fea75e47108c8bd8db61a3de84b266733d2d7849bec46eb.NewCanSignUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedSignupStatusRequestBuilder) CompleteSetup()(*ib823f830e3ea0500a214b973cf170985ed37b35d907e2b81e06df8aaa8f66256.CompleteSetupRequestBuilder) {
    return ib823f830e3ea0500a214b973cf170985ed37b35d907e2b81e06df8aaa8f66256.NewCompleteSetupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewPrivilegedSignupStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedSignupStatusRequestBuilder) {
    m := &PrivilegedSignupStatusRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/privilegedSignupStatus{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewPrivilegedSignupStatusRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedSignupStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedSignupStatusRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrivilegedSignupStatusRequestBuilder) CreateGetRequestInformation(q func (value *PrivilegedSignupStatusRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PrivilegedSignupStatusRequestBuilderGetQueryParameters)
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
func (m *PrivilegedSignupStatusRequestBuilder) CreatePostRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedSignupStatus, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrivilegedSignupStatusRequestBuilder) Get(q func (value *PrivilegedSignupStatusRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*PrivilegedSignupStatusResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedSignupStatusResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*PrivilegedSignupStatusResponse), nil
}
func (m *PrivilegedSignupStatusRequestBuilder) IsSignedUp()(*i505eff958ea7946e725489038a9fbe85c297626b944e12d9c6f68697f1a65258.IsSignedUpRequestBuilder) {
    return i505eff958ea7946e725489038a9fbe85c297626b944e12d9c6f68697f1a65258.NewIsSignedUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedSignupStatusRequestBuilder) Post(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedSignupStatus, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedSignupStatus, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrivilegedSignupStatus() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedSignupStatus), nil
}
func (m *PrivilegedSignupStatusRequestBuilder) SignUp()(*ibe660ff6aa6da80e38b282109151294b9409fe46a692bb37484dd8a1705b5a94.SignUpRequestBuilder) {
    return ibe660ff6aa6da80e38b282109151294b9409fe46a692bb37484dd8a1705b5a94.NewSignUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
