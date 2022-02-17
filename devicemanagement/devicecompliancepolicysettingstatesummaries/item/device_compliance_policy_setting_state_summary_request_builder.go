package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i37ad49c0bc6d68e963fdf65f7323a0821e118797dfc71b7300af4b5c7388c030 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicysettingstatesummaries/item/devicecompliancesettingstates"
    i623eaf2a0acafba67f551c159b5f0c09a5a15743477194b11a435ff9831edaa1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicysettingstatesummaries/item/devicecompliancesettingstates/item"
)

// DeviceCompliancePolicySettingStateSummaryRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCompliancePolicySettingStateSummaries\{deviceCompliancePolicySettingStateSummary-id}
type DeviceCompliancePolicySettingStateSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceCompliancePolicySettingStateSummaryRequestBuilderDeleteOptions options for Delete
type DeviceCompliancePolicySettingStateSummaryRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCompliancePolicySettingStateSummaryRequestBuilderGetOptions options for Get
type DeviceCompliancePolicySettingStateSummaryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceCompliancePolicySettingStateSummaryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCompliancePolicySettingStateSummaryRequestBuilderGetQueryParameters the summary states of compliance policy settings for this account.
type DeviceCompliancePolicySettingStateSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceCompliancePolicySettingStateSummaryRequestBuilderPatchOptions options for Patch
type DeviceCompliancePolicySettingStateSummaryRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicySettingStateSummary;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceCompliancePolicySettingStateSummaryRequestBuilderInternal instantiates a new DeviceCompliancePolicySettingStateSummaryRequestBuilder and sets the default values.
func NewDeviceCompliancePolicySettingStateSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicySettingStateSummaryRequestBuilder) {
    m := &DeviceCompliancePolicySettingStateSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCompliancePolicySettingStateSummaryRequestBuilder instantiates a new DeviceCompliancePolicySettingStateSummaryRequestBuilder and sets the default values.
func NewDeviceCompliancePolicySettingStateSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicySettingStateSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePolicySettingStateSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryRequestBuilder) CreateDeleteRequestInformation(options *DeviceCompliancePolicySettingStateSummaryRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryRequestBuilder) CreateGetRequestInformation(options *DeviceCompliancePolicySettingStateSummaryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryRequestBuilder) CreatePatchRequestInformation(options *DeviceCompliancePolicySettingStateSummaryRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryRequestBuilder) Delete(options *DeviceCompliancePolicySettingStateSummaryRequestBuilderDeleteOptions)(error) {
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
func (m *DeviceCompliancePolicySettingStateSummaryRequestBuilder) DeviceComplianceSettingStates()(*i37ad49c0bc6d68e963fdf65f7323a0821e118797dfc71b7300af4b5c7388c030.DeviceComplianceSettingStatesRequestBuilder) {
    return i37ad49c0bc6d68e963fdf65f7323a0821e118797dfc71b7300af4b5c7388c030.NewDeviceComplianceSettingStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceComplianceSettingStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCompliancePolicySettingStateSummaries.item.deviceComplianceSettingStates.item collection
func (m *DeviceCompliancePolicySettingStateSummaryRequestBuilder) DeviceComplianceSettingStatesById(id string)(*i623eaf2a0acafba67f551c159b5f0c09a5a15743477194b11a435ff9831edaa1.DeviceComplianceSettingStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceSettingState_id"] = id
    }
    return i623eaf2a0acafba67f551c159b5f0c09a5a15743477194b11a435ff9831edaa1.NewDeviceComplianceSettingStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryRequestBuilder) Get(options *DeviceCompliancePolicySettingStateSummaryRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicySettingStateSummary, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceCompliancePolicySettingStateSummary() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicySettingStateSummary), nil
}
// Patch the summary states of compliance policy settings for this account.
func (m *DeviceCompliancePolicySettingStateSummaryRequestBuilder) Patch(options *DeviceCompliancePolicySettingStateSummaryRequestBuilderPatchOptions)(error) {
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
