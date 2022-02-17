package androidforworksettings

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i280e19d80d7a529d2b9933cf3044f19490e2f99605f9859eeb0710b5d4c6b359 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworksettings/completesignup"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i61c16d732624a212d9246ecf43fd01db8d4bc43debd10e8907b417795c13f8ff "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworksettings/requestsignupurl"
    i97a07e45ebea436c88e03068b5e337121e7e6f6d82dd6f519d86bd92618a0817 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworksettings/syncapps"
    ibc8cdcda9d0880d123b6a234fd0980335960dc7e2f066cbce430ed5b5bd54ef8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworksettings/unbind"
)

// AndroidForWorkSettingsRequestBuilder builds and executes requests for operations under \deviceManagement\androidForWorkSettings
type AndroidForWorkSettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AndroidForWorkSettingsRequestBuilderDeleteOptions options for Delete
type AndroidForWorkSettingsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidForWorkSettingsRequestBuilderGetOptions options for Get
type AndroidForWorkSettingsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AndroidForWorkSettingsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidForWorkSettingsRequestBuilderGetQueryParameters the singleton Android for Work settings entity.
type AndroidForWorkSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AndroidForWorkSettingsRequestBuilderPatchOptions options for Patch
type AndroidForWorkSettingsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidForWorkSettings;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AndroidForWorkSettingsRequestBuilder) CompleteSignup()(*i280e19d80d7a529d2b9933cf3044f19490e2f99605f9859eeb0710b5d4c6b359.CompleteSignupRequestBuilder) {
    return i280e19d80d7a529d2b9933cf3044f19490e2f99605f9859eeb0710b5d4c6b359.NewCompleteSignupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAndroidForWorkSettingsRequestBuilderInternal instantiates a new AndroidForWorkSettingsRequestBuilder and sets the default values.
func NewAndroidForWorkSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidForWorkSettingsRequestBuilder) {
    m := &AndroidForWorkSettingsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidForWorkSettings{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAndroidForWorkSettingsRequestBuilder instantiates a new AndroidForWorkSettingsRequestBuilder and sets the default values.
func NewAndroidForWorkSettingsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidForWorkSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidForWorkSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) CreateDeleteRequestInformation(options *AndroidForWorkSettingsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) CreateGetRequestInformation(options *AndroidForWorkSettingsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
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
// CreatePatchRequestInformation the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) CreatePatchRequestInformation(options *AndroidForWorkSettingsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) Delete(options *AndroidForWorkSettingsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) Get(options *AndroidForWorkSettingsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidForWorkSettings, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAndroidForWorkSettings() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidForWorkSettings), nil
}
// Patch the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) Patch(options *AndroidForWorkSettingsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AndroidForWorkSettingsRequestBuilder) RequestSignupUrl()(*i61c16d732624a212d9246ecf43fd01db8d4bc43debd10e8907b417795c13f8ff.RequestSignupUrlRequestBuilder) {
    return i61c16d732624a212d9246ecf43fd01db8d4bc43debd10e8907b417795c13f8ff.NewRequestSignupUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AndroidForWorkSettingsRequestBuilder) SyncApps()(*i97a07e45ebea436c88e03068b5e337121e7e6f6d82dd6f519d86bd92618a0817.SyncAppsRequestBuilder) {
    return i97a07e45ebea436c88e03068b5e337121e7e6f6d82dd6f519d86bd92618a0817.NewSyncAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AndroidForWorkSettingsRequestBuilder) Unbind()(*ibc8cdcda9d0880d123b6a234fd0980335960dc7e2f066cbce430ed5b5bd54ef8.UnbindRequestBuilder) {
    return ibc8cdcda9d0880d123b6a234fd0980335960dc7e2f066cbce430ed5b5bd54ef8.NewUnbindRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
