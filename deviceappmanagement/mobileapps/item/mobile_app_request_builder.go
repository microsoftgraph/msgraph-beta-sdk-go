package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i019948c5cc3432cae1aa4d885d81eaee957e9c7ddfea6d161a8a6b12d04d693b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/installsummary"
    i0a3036e7599c27e8daf0ac76114f7781865beca45274e73af8c8395fc2a9665d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses"
    i1e3d61799ac5b1670257d353559a6aa4ec7e545fb82dd83595bf36c04e7e9b1d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/assign"
    i2cc000920189f2c9f846e98b23559d1c3218e590b3b9322853afced5128198d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/iosvppapp"
    i3490d416205b3df42d9b8ee0ff657b79e6262e2c3b79708a8f0fb50f5a1c2499 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses"
    i413f580db54415c14dab7adbde8526f1b741cd16690fd87a8ff30b9108f0ec47 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/updaterelationships"
    i9b6c1e48145f74fcaf231180b0b8a979f941d054e34a28d2e470e01c6a7ee864 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/getrelatedappstateswithuserprincipalnamewithdeviceid"
    iaca75b4f217f7322eba62ee5fd8c052d9514c3394b480376576627995983ab4a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/relationships"
    idd22d981d731ad036f829684abde198fe174082361ffe1a63d853407dca69d18 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/categories"
    ifb2db1cec9c276d6c90084ff24743ae3570fed79bdca00e6ab2191624f49ad72 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/assignments"
    i7578dc338c9dc393bea0c169238874f0e19fedfffc44418ab516b34a4b525138 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item"
    i8ae606a13a79cd675f0f56386d1cfce72468cf3465afc5d50a1e5bfa0e57e603 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/assignments/item"
    ia6a71ed4e86b560dd63fb4ca626455944edeea774053bb5436fcd463ee9a2cc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/relationships/item"
    ifddf3b1a5eae62e5fad4e7d420246006149401e2aacd31254c493f3bd5f95b80 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item"
)

type MobileAppRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type MobileAppRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *MobileAppRequestBuilder) Assign()(*i1e3d61799ac5b1670257d353559a6aa4ec7e545fb82dd83595bf36c04e7e9b1d.AssignRequestBuilder) {
    return i1e3d61799ac5b1670257d353559a6aa4ec7e545fb82dd83595bf36c04e7e9b1d.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) Assignments()(*ifb2db1cec9c276d6c90084ff24743ae3570fed79bdca00e6ab2191624f49ad72.AssignmentsRequestBuilder) {
    return ifb2db1cec9c276d6c90084ff24743ae3570fed79bdca00e6ab2191624f49ad72.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) AssignmentsById(id string)(*i8ae606a13a79cd675f0f56386d1cfce72468cf3465afc5d50a1e5bfa0e57e603.MobileAppAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["mobileAppAssignment_id"] = id
    }
    return i8ae606a13a79cd675f0f56386d1cfce72468cf3465afc5d50a1e5bfa0e57e603.NewMobileAppAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) Categories()(*idd22d981d731ad036f829684abde198fe174082361ffe1a63d853407dca69d18.CategoriesRequestBuilder) {
    return idd22d981d731ad036f829684abde198fe174082361ffe1a63d853407dca69d18.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewMobileAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppRequestBuilder) {
    m := &MobileAppRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/mobileApps/{mobileApp_id}{?select,expand}";
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
func NewMobileAppRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MobileAppRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MobileAppRequestBuilder) CreateGetRequestInformation(q func (value *MobileAppRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(MobileAppRequestBuilderGetQueryParameters)
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
func (m *MobileAppRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MobileAppRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *MobileAppRequestBuilder) DeviceStatuses()(*i3490d416205b3df42d9b8ee0ff657b79e6262e2c3b79708a8f0fb50f5a1c2499.DeviceStatusesRequestBuilder) {
    return i3490d416205b3df42d9b8ee0ff657b79e6262e2c3b79708a8f0fb50f5a1c2499.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) DeviceStatusesById(id string)(*i7578dc338c9dc393bea0c169238874f0e19fedfffc44418ab516b34a4b525138.MobileAppInstallStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["mobileAppInstallStatus_id"] = id
    }
    return i7578dc338c9dc393bea0c169238874f0e19fedfffc44418ab516b34a4b525138.NewMobileAppInstallStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) Get(q func (value *MobileAppRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp, error) {
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
func (m *MobileAppRequestBuilder) GetRelatedAppStatesWithUserPrincipalNameWithDeviceId(userPrincipalName *string, deviceId *string)(*i9b6c1e48145f74fcaf231180b0b8a979f941d054e34a28d2e470e01c6a7ee864.GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    return i9b6c1e48145f74fcaf231180b0b8a979f941d054e34a28d2e470e01c6a7ee864.NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, userPrincipalName, deviceId);
}
func (m *MobileAppRequestBuilder) InstallSummary()(*i019948c5cc3432cae1aa4d885d81eaee957e9c7ddfea6d161a8a6b12d04d693b.InstallSummaryRequestBuilder) {
    return i019948c5cc3432cae1aa4d885d81eaee957e9c7ddfea6d161a8a6b12d04d693b.NewInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) IosVppApp()(*i2cc000920189f2c9f846e98b23559d1c3218e590b3b9322853afced5128198d6.IosVppAppRequestBuilder) {
    return i2cc000920189f2c9f846e98b23559d1c3218e590b3b9322853afced5128198d6.NewIosVppAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *MobileAppRequestBuilder) Relationships()(*iaca75b4f217f7322eba62ee5fd8c052d9514c3394b480376576627995983ab4a.RelationshipsRequestBuilder) {
    return iaca75b4f217f7322eba62ee5fd8c052d9514c3394b480376576627995983ab4a.NewRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) RelationshipsById(id string)(*ia6a71ed4e86b560dd63fb4ca626455944edeea774053bb5436fcd463ee9a2cc2.MobileAppRelationshipRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["mobileAppRelationship_id"] = id
    }
    return ia6a71ed4e86b560dd63fb4ca626455944edeea774053bb5436fcd463ee9a2cc2.NewMobileAppRelationshipRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) UpdateRelationships()(*i413f580db54415c14dab7adbde8526f1b741cd16690fd87a8ff30b9108f0ec47.UpdateRelationshipsRequestBuilder) {
    return i413f580db54415c14dab7adbde8526f1b741cd16690fd87a8ff30b9108f0ec47.NewUpdateRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) UserStatuses()(*i0a3036e7599c27e8daf0ac76114f7781865beca45274e73af8c8395fc2a9665d.UserStatusesRequestBuilder) {
    return i0a3036e7599c27e8daf0ac76114f7781865beca45274e73af8c8395fc2a9665d.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) UserStatusesById(id string)(*ifddf3b1a5eae62e5fad4e7d420246006149401e2aacd31254c493f3bd5f95b80.UserAppInstallStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["userAppInstallStatus_id"] = id
    }
    return ifddf3b1a5eae62e5fad4e7d420246006149401e2aacd31254c493f3bd5f95b80.NewUserAppInstallStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
