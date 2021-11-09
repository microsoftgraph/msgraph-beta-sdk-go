package advancedthreatprotectiononboardingstatesummary

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ice19fdfd521d3baa084d788d924acd2488705b59b4281e0f60085cec8edcab8d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/advancedthreatprotectiononboardingstatesummary/advancedthreatprotectiononboardingdevicesettingstates"
    i8ac982e4babd498ccfdca8159c2c00708b65a8c1d00512bc9fe29edca2a507ca "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/advancedthreatprotectiononboardingstatesummary/advancedthreatprotectiononboardingdevicesettingstates/item"
)

// Builds and executes requests for operations under \deviceManagement\advancedThreatProtectionOnboardingStateSummary
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The summary state of ATP onboarding state for this account.
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdvancedThreatProtectionOnboardingStateSummary;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStates()(*ice19fdfd521d3baa084d788d924acd2488705b59b4281e0f60085cec8edcab8d.AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    return ice19fdfd521d3baa084d788d924acd2488705b59b4281e0f60085cec8edcab8d.NewAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.advancedThreatProtectionOnboardingStateSummary.advancedThreatProtectionOnboardingDeviceSettingStates.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStatesById(id string)(*i8ac982e4babd498ccfdca8159c2c00708b65a8c1d00512bc9fe29edca2a507ca.AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["advancedThreatProtectionOnboardingDeviceSettingState_id"] = id
    }
    return i8ac982e4babd498ccfdca8159c2c00708b65a8c1d00512bc9fe29edca2a507ca.NewAdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    m := &AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// The summary state of ATP onboarding state for this account.
// Parameters:
//  - options : Options for the request
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreateDeleteRequestInformation(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The summary state of ATP onboarding state for this account.
// Parameters:
//  - options : Options for the request
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreateGetRequestInformation(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The summary state of ATP onboarding state for this account.
// Parameters:
//  - options : Options for the request
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreatePatchRequestInformation(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The summary state of ATP onboarding state for this account.
// Parameters:
//  - options : Options for the request
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Delete(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions)(error) {
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
// The summary state of ATP onboarding state for this account.
// Parameters:
//  - options : Options for the request
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Get(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdvancedThreatProtectionOnboardingStateSummary, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAdvancedThreatProtectionOnboardingStateSummary() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdvancedThreatProtectionOnboardingStateSummary), nil
}
// The summary state of ATP onboarding state for this account.
// Parameters:
//  - options : Options for the request
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Patch(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions)(error) {
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
