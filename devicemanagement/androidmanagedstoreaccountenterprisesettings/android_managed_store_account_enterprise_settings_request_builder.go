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

// Builds and executes requests for operations under \deviceManagement\androidManagedStoreAccountEnterpriseSettings
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The singleton Android managed store account enterprise settings entity.
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedStoreAccountEnterpriseSettings;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ApproveApps()(*ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54.ApproveAppsRequestBuilder) {
    return ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54.NewApproveAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CompleteSignup()(*i314c76ed635b4004a342dd0c841e00e8f471d83c6cb71a6344ebcca4508e202a.CompleteSignupRequestBuilder) {
    return i314c76ed635b4004a342dd0c841e00e8f471d83c6cb71a6344ebcca4508e202a.NewCompleteSignupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
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
// Instantiates a new AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// The singleton Android managed store account enterprise settings entity.
// Parameters:
//  - options : Options for the request
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateDeleteRequestInformation(options *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The singleton Android managed store account enterprise settings entity.
// Parameters:
//  - options : Options for the request
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGetRequestInformation(options *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGooglePlayWebToken()(*i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c.CreateGooglePlayWebTokenRequestBuilder) {
    return i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c.NewCreateGooglePlayWebTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The singleton Android managed store account enterprise settings entity.
// Parameters:
//  - options : Options for the request
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreatePatchRequestInformation(options *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The singleton Android managed store account enterprise settings entity.
// Parameters:
//  - options : Options for the request
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Delete(options *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// The singleton Android managed store account enterprise settings entity.
// Parameters:
//  - options : Options for the request
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Get(options *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedStoreAccountEnterpriseSettings, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAndroidManagedStoreAccountEnterpriseSettings() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedStoreAccountEnterpriseSettings), nil
}
// The singleton Android managed store account enterprise settings entity.
// Parameters:
//  - options : Options for the request
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Patch(options *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
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
