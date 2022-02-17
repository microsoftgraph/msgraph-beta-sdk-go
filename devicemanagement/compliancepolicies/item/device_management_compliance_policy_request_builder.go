package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i02451624685ebfdb81d6b1930c79509f0df4294168ed18436aaa09cc42c5c946 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/assignments"
    i0742bff39d39edf9e7fdaa5802f2ceda48231b0df3ad173faa313f5c87241b5a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/assign"
    i15585ed0e13327df9171fd8499ee3db205ffd397a580c4ed84e065a5d220e628 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/settings"
    i1fb378ecdc09aaabf5c55da82fe6def98379ca93c0d2715669e8fa5761d1ed11 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/setscheduledactions"
    iffa40ac9a6da23297b60087932a288a73ea0e6ab171a35dec05137c5d4ed2f62 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/scheduledactionsforrule"
    i1d9a96a6a2f04ecb11762741e92d99faf72d8bfcea65eb2f15cca45f47882da5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/assignments/item"
    i6cd919cf4ce1e5e402de893f996a02ad42aff8ec4ade9971ee45c71e8fbcb47a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/scheduledactionsforrule/item"
    ie4d81e74e3624280f0e519971ef2c7635b27540efa521cec95913734fcecc59f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/settings/item"
)

// DeviceManagementCompliancePolicyRequestBuilder builds and executes requests for operations under \deviceManagement\compliancePolicies\{deviceManagementCompliancePolicy-id}
type DeviceManagementCompliancePolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementCompliancePolicyRequestBuilderDeleteOptions options for Delete
type DeviceManagementCompliancePolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementCompliancePolicyRequestBuilderGetOptions options for Get
type DeviceManagementCompliancePolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementCompliancePolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementCompliancePolicyRequestBuilderGetQueryParameters list of all compliance policies
type DeviceManagementCompliancePolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementCompliancePolicyRequestBuilderPatchOptions options for Patch
type DeviceManagementCompliancePolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementCompliancePolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementCompliancePolicyRequestBuilder) Assign()(*i0742bff39d39edf9e7fdaa5802f2ceda48231b0df3ad173faa313f5c87241b5a.AssignRequestBuilder) {
    return i0742bff39d39edf9e7fdaa5802f2ceda48231b0df3ad173faa313f5c87241b5a.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementCompliancePolicyRequestBuilder) Assignments()(*i02451624685ebfdb81d6b1930c79509f0df4294168ed18436aaa09cc42c5c946.AssignmentsRequestBuilder) {
    return i02451624685ebfdb81d6b1930c79509f0df4294168ed18436aaa09cc42c5c946.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.assignments.item collection
func (m *DeviceManagementCompliancePolicyRequestBuilder) AssignmentsById(id string)(*i1d9a96a6a2f04ecb11762741e92d99faf72d8bfcea65eb2f15cca45f47882da5.DeviceManagementConfigurationPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment_id"] = id
    }
    return i1d9a96a6a2f04ecb11762741e92d99faf72d8bfcea65eb2f15cca45f47882da5.NewDeviceManagementConfigurationPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementCompliancePolicyRequestBuilderInternal instantiates a new DeviceManagementCompliancePolicyRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementCompliancePolicyRequestBuilder) {
    m := &DeviceManagementCompliancePolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementCompliancePolicyRequestBuilder instantiates a new DeviceManagementCompliancePolicyRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementCompliancePolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementCompliancePolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation list of all compliance policies
func (m *DeviceManagementCompliancePolicyRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementCompliancePolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation list of all compliance policies
func (m *DeviceManagementCompliancePolicyRequestBuilder) CreateGetRequestInformation(options *DeviceManagementCompliancePolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation list of all compliance policies
func (m *DeviceManagementCompliancePolicyRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementCompliancePolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete list of all compliance policies
func (m *DeviceManagementCompliancePolicyRequestBuilder) Delete(options *DeviceManagementCompliancePolicyRequestBuilderDeleteOptions)(error) {
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
// Get list of all compliance policies
func (m *DeviceManagementCompliancePolicyRequestBuilder) Get(options *DeviceManagementCompliancePolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementCompliancePolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementCompliancePolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementCompliancePolicy), nil
}
// Patch list of all compliance policies
func (m *DeviceManagementCompliancePolicyRequestBuilder) Patch(options *DeviceManagementCompliancePolicyRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementCompliancePolicyRequestBuilder) ScheduledActionsForRule()(*iffa40ac9a6da23297b60087932a288a73ea0e6ab171a35dec05137c5d4ed2f62.ScheduledActionsForRuleRequestBuilder) {
    return iffa40ac9a6da23297b60087932a288a73ea0e6ab171a35dec05137c5d4ed2f62.NewScheduledActionsForRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionsForRuleById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.scheduledActionsForRule.item collection
func (m *DeviceManagementCompliancePolicyRequestBuilder) ScheduledActionsForRuleById(id string)(*i6cd919cf4ce1e5e402de893f996a02ad42aff8ec4ade9971ee45c71e8fbcb47a.DeviceManagementComplianceScheduledActionForRuleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementComplianceScheduledActionForRule_id"] = id
    }
    return i6cd919cf4ce1e5e402de893f996a02ad42aff8ec4ade9971ee45c71e8fbcb47a.NewDeviceManagementComplianceScheduledActionForRuleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementCompliancePolicyRequestBuilder) SetScheduledActions()(*i1fb378ecdc09aaabf5c55da82fe6def98379ca93c0d2715669e8fa5761d1ed11.SetScheduledActionsRequestBuilder) {
    return i1fb378ecdc09aaabf5c55da82fe6def98379ca93c0d2715669e8fa5761d1ed11.NewSetScheduledActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementCompliancePolicyRequestBuilder) Settings()(*i15585ed0e13327df9171fd8499ee3db205ffd397a580c4ed84e065a5d220e628.SettingsRequestBuilder) {
    return i15585ed0e13327df9171fd8499ee3db205ffd397a580c4ed84e065a5d220e628.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.settings.item collection
func (m *DeviceManagementCompliancePolicyRequestBuilder) SettingsById(id string)(*ie4d81e74e3624280f0e519971ef2c7635b27540efa521cec95913734fcecc59f.DeviceManagementConfigurationSettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting_id"] = id
    }
    return ie4d81e74e3624280f0e519971ef2c7635b27540efa521cec95913734fcecc59f.NewDeviceManagementConfigurationSettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
