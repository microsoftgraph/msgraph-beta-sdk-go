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

// DeviceManagementCompliancePolicyItemRequestBuilder builds and executes requests for operations under \deviceManagement\compliancePolicies\{deviceManagementCompliancePolicy-id}
type DeviceManagementCompliancePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementCompliancePolicyItemRequestBuilderDeleteOptions options for Delete
type DeviceManagementCompliancePolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementCompliancePolicyItemRequestBuilderGetOptions options for Get
type DeviceManagementCompliancePolicyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementCompliancePolicyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementCompliancePolicyItemRequestBuilderGetQueryParameters list of all compliance policies
type DeviceManagementCompliancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementCompliancePolicyItemRequestBuilderPatchOptions options for Patch
type DeviceManagementCompliancePolicyItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementCompliancePolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Assign()(*i0742bff39d39edf9e7fdaa5802f2ceda48231b0df3ad173faa313f5c87241b5a.AssignRequestBuilder) {
    return i0742bff39d39edf9e7fdaa5802f2ceda48231b0df3ad173faa313f5c87241b5a.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Assignments()(*i02451624685ebfdb81d6b1930c79509f0df4294168ed18436aaa09cc42c5c946.AssignmentsRequestBuilder) {
    return i02451624685ebfdb81d6b1930c79509f0df4294168ed18436aaa09cc42c5c946.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.assignments.item collection
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) AssignmentsById(id string)(*i1d9a96a6a2f04ecb11762741e92d99faf72d8bfcea65eb2f15cca45f47882da5.DeviceManagementConfigurationPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment_id"] = id
    }
    return i1d9a96a6a2f04ecb11762741e92d99faf72d8bfcea65eb2f15cca45f47882da5.NewDeviceManagementConfigurationPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementCompliancePolicyItemRequestBuilderInternal instantiates a new DeviceManagementCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementCompliancePolicyItemRequestBuilder) {
    m := &DeviceManagementCompliancePolicyItemRequestBuilder{
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
// NewDeviceManagementCompliancePolicyItemRequestBuilder instantiates a new DeviceManagementCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePolicyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementCompliancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementCompliancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation list of all compliance policies
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementCompliancePolicyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) CreateGetRequestInformation(options *DeviceManagementCompliancePolicyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementCompliancePolicyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Delete(options *DeviceManagementCompliancePolicyItemRequestBuilderDeleteOptions)(error) {
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
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Get(options *DeviceManagementCompliancePolicyItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementCompliancePolicy, error) {
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
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Patch(options *DeviceManagementCompliancePolicyItemRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) ScheduledActionsForRule()(*iffa40ac9a6da23297b60087932a288a73ea0e6ab171a35dec05137c5d4ed2f62.ScheduledActionsForRuleRequestBuilder) {
    return iffa40ac9a6da23297b60087932a288a73ea0e6ab171a35dec05137c5d4ed2f62.NewScheduledActionsForRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionsForRuleById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.scheduledActionsForRule.item collection
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) ScheduledActionsForRuleById(id string)(*i6cd919cf4ce1e5e402de893f996a02ad42aff8ec4ade9971ee45c71e8fbcb47a.DeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementComplianceScheduledActionForRule_id"] = id
    }
    return i6cd919cf4ce1e5e402de893f996a02ad42aff8ec4ade9971ee45c71e8fbcb47a.NewDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) SetScheduledActions()(*i1fb378ecdc09aaabf5c55da82fe6def98379ca93c0d2715669e8fa5761d1ed11.SetScheduledActionsRequestBuilder) {
    return i1fb378ecdc09aaabf5c55da82fe6def98379ca93c0d2715669e8fa5761d1ed11.NewSetScheduledActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Settings()(*i15585ed0e13327df9171fd8499ee3db205ffd397a580c4ed84e065a5d220e628.SettingsRequestBuilder) {
    return i15585ed0e13327df9171fd8499ee3db205ffd397a580c4ed84e065a5d220e628.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.settings.item collection
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) SettingsById(id string)(*ie4d81e74e3624280f0e519971ef2c7635b27540efa521cec95913734fcecc59f.DeviceManagementConfigurationSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting_id"] = id
    }
    return ie4d81e74e3624280f0e519971ef2c7635b27540efa521cec95913734fcecc59f.NewDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
