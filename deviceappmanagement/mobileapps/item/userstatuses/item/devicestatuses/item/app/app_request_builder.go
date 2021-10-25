package app

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i298beea648eea7ecd6a4c3d720fd467d761b72d167f7ea468e97ce22b751da65 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses/item/app/getrelatedappstateswithuserprincipalnamewithdeviceid"
    i955becffdda044167c61f9cd27395172d0f852d19cd55d86b402f726886b9624 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses/item/app/updaterelationships"
    i9b1b82d0ce5bf6333f34485f2d82f3e3afb9b94aeb733725fb426c899d33545d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses/item/app/assign"
    iac18b6eefd3a1314fa3e1fad7481436233d382c19f6a04d21b24b5e176ed655a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses/item/app/ref"
    idf744410b6e417d148101462c907c090ca20e4a3f3de5ab2aa7283bff887c904 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses/item/app/iosvppapp"
)

type AppRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type AppRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *AppRequestBuilder) Assign()(*i9b1b82d0ce5bf6333f34485f2d82f3e3afb9b94aeb733725fb426c899d33545d.AssignRequestBuilder) {
    return i9b1b82d0ce5bf6333f34485f2d82f3e3afb9b94aeb733725fb426c899d33545d.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppRequestBuilder) {
    m := &AppRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/mobileApps/{mobileApp_id}/userStatuses/{userAppInstallStatus_id}/deviceStatuses/{mobileAppInstallStatus_id}/app{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewAppRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AppRequestBuilder) CreateGetRequestInformation(q func (value *AppRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(AppRequestBuilderGetQueryParameters)
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
func (m *AppRequestBuilder) Get(q func (value *AppRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMobileApp() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp), nil
}
func (m *AppRequestBuilder) GetRelatedAppStatesWithUserPrincipalNameWithDeviceId(userPrincipalName *string, deviceId *string)(*i298beea648eea7ecd6a4c3d720fd467d761b72d167f7ea468e97ce22b751da65.GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    return i298beea648eea7ecd6a4c3d720fd467d761b72d167f7ea468e97ce22b751da65.NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, userPrincipalName, deviceId);
}
func (m *AppRequestBuilder) IosVppApp()(*idf744410b6e417d148101462c907c090ca20e4a3f3de5ab2aa7283bff887c904.IosVppAppRequestBuilder) {
    return idf744410b6e417d148101462c907c090ca20e4a3f3de5ab2aa7283bff887c904.NewIosVppAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AppRequestBuilder) Ref()(*iac18b6eefd3a1314fa3e1fad7481436233d382c19f6a04d21b24b5e176ed655a.RefRequestBuilder) {
    return iac18b6eefd3a1314fa3e1fad7481436233d382c19f6a04d21b24b5e176ed655a.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AppRequestBuilder) UpdateRelationships()(*i955becffdda044167c61f9cd27395172d0f852d19cd55d86b402f726886b9624.UpdateRelationshipsRequestBuilder) {
    return i955becffdda044167c61f9cd27395172d0f852d19cd55d86b402f726886b9624.NewUpdateRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
