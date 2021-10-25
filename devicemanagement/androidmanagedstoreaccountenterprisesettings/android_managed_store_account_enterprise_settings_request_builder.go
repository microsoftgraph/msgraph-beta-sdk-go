package androidmanagedstoreaccountenterprisesettings

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/creategoogleplaywebtoken"
    i314c76ed635b4004a342dd0c841e00e8f471d83c6cb71a6344ebcca4508e202a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/completesignup"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5570da22a9e8ab3a3faa339899d50bf58d9d37e83641e9d383b921ee28101901 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/syncapps"
    i573858361497d031356ed7937a66002f0e820977a38c57afb805332721bfbd14 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/unbind"
    i9821f431564656dfb670bc3758898c81b836fe062cdacc51bf1462209f22c694 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/setandroiddeviceownerfullymanagedenrollmentstate"
    ia312e54dde3f35887f50da33001441eb2d48646e6fd640a4c7387523b4934a7e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/requestsignupurl"
    ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/approveapps"
)

type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ApproveApps()(*ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54.ApproveAppsRequestBuilder) {
    return ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54.NewApproveAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CompleteSignup()(*i314c76ed635b4004a342dd0c841e00e8f471d83c6cb71a6344ebcca4508e202a.CompleteSignupRequestBuilder) {
    return i314c76ed635b4004a342dd0c841e00e8f471d83c6cb71a6344ebcca4508e202a.NewCompleteSignupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/androidManagedStoreAccountEnterpriseSettings{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGetRequestInformation(q func (value *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters)
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGooglePlayWebToken()(*i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c.CreateGooglePlayWebTokenRequestBuilder) {
    return i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c.NewCreateGooglePlayWebTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedStoreAccountEnterpriseSettings, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Get(q func (value *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedStoreAccountEnterpriseSettings, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAndroidManagedStoreAccountEnterpriseSettings() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedStoreAccountEnterpriseSettings), nil
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedStoreAccountEnterpriseSettings, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) RequestSignupUrl()(*ia312e54dde3f35887f50da33001441eb2d48646e6fd640a4c7387523b4934a7e.RequestSignupUrlRequestBuilder) {
    return ia312e54dde3f35887f50da33001441eb2d48646e6fd640a4c7387523b4934a7e.NewRequestSignupUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SetAndroidDeviceOwnerFullyManagedEnrollmentState()(*i9821f431564656dfb670bc3758898c81b836fe062cdacc51bf1462209f22c694.SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    return i9821f431564656dfb670bc3758898c81b836fe062cdacc51bf1462209f22c694.NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SyncApps()(*i5570da22a9e8ab3a3faa339899d50bf58d9d37e83641e9d383b921ee28101901.SyncAppsRequestBuilder) {
    return i5570da22a9e8ab3a3faa339899d50bf58d9d37e83641e9d383b921ee28101901.NewSyncAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Unbind()(*i573858361497d031356ed7937a66002f0e820977a38c57afb805332721bfbd14.UnbindRequestBuilder) {
    return i573858361497d031356ed7937a66002f0e820977a38c57afb805332721bfbd14.NewUnbindRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
