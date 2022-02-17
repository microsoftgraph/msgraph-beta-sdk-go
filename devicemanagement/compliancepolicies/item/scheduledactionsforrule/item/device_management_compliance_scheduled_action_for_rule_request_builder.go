package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i50f3c1db96bb631a5b29e667d93d389d4e7728d45693aaf6947f51821a54ae0e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/scheduledactionsforrule/item/scheduledactionconfigurations"
    ib3a97c3a0c4a95294f8aac578747e997f8e6d9c2c4b26c8c4d20581b8d3e6cda "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/scheduledactionsforrule/item/scheduledactionconfigurations/item"
)

// DeviceManagementComplianceScheduledActionForRuleRequestBuilder builds and executes requests for operations under \deviceManagement\compliancePolicies\{deviceManagementCompliancePolicy-id}\scheduledActionsForRule\{deviceManagementComplianceScheduledActionForRule-id}
type DeviceManagementComplianceScheduledActionForRuleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementComplianceScheduledActionForRuleRequestBuilderDeleteOptions options for Delete
type DeviceManagementComplianceScheduledActionForRuleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementComplianceScheduledActionForRuleRequestBuilderGetOptions options for Get
type DeviceManagementComplianceScheduledActionForRuleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementComplianceScheduledActionForRuleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementComplianceScheduledActionForRuleRequestBuilderGetQueryParameters the list of scheduled action for this rule
type DeviceManagementComplianceScheduledActionForRuleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementComplianceScheduledActionForRuleRequestBuilderPatchOptions options for Patch
type DeviceManagementComplianceScheduledActionForRuleRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementComplianceScheduledActionForRule;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceManagementComplianceScheduledActionForRuleRequestBuilderInternal instantiates a new DeviceManagementComplianceScheduledActionForRuleRequestBuilder and sets the default values.
func NewDeviceManagementComplianceScheduledActionForRuleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementComplianceScheduledActionForRuleRequestBuilder) {
    m := &DeviceManagementComplianceScheduledActionForRuleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy_id}/scheduledActionsForRule/{deviceManagementComplianceScheduledActionForRule_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementComplianceScheduledActionForRuleRequestBuilder instantiates a new DeviceManagementComplianceScheduledActionForRuleRequestBuilder and sets the default values.
func NewDeviceManagementComplianceScheduledActionForRuleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementComplianceScheduledActionForRuleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementComplianceScheduledActionForRuleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of scheduled action for this rule
func (m *DeviceManagementComplianceScheduledActionForRuleRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementComplianceScheduledActionForRuleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of scheduled action for this rule
func (m *DeviceManagementComplianceScheduledActionForRuleRequestBuilder) CreateGetRequestInformation(options *DeviceManagementComplianceScheduledActionForRuleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of scheduled action for this rule
func (m *DeviceManagementComplianceScheduledActionForRuleRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementComplianceScheduledActionForRuleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of scheduled action for this rule
func (m *DeviceManagementComplianceScheduledActionForRuleRequestBuilder) Delete(options *DeviceManagementComplianceScheduledActionForRuleRequestBuilderDeleteOptions)(error) {
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
// Get the list of scheduled action for this rule
func (m *DeviceManagementComplianceScheduledActionForRuleRequestBuilder) Get(options *DeviceManagementComplianceScheduledActionForRuleRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementComplianceScheduledActionForRule, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementComplianceScheduledActionForRule() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementComplianceScheduledActionForRule), nil
}
// Patch the list of scheduled action for this rule
func (m *DeviceManagementComplianceScheduledActionForRuleRequestBuilder) Patch(options *DeviceManagementComplianceScheduledActionForRuleRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementComplianceScheduledActionForRuleRequestBuilder) ScheduledActionConfigurations()(*i50f3c1db96bb631a5b29e667d93d389d4e7728d45693aaf6947f51821a54ae0e.ScheduledActionConfigurationsRequestBuilder) {
    return i50f3c1db96bb631a5b29e667d93d389d4e7728d45693aaf6947f51821a54ae0e.NewScheduledActionConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.scheduledActionsForRule.item.scheduledActionConfigurations.item collection
func (m *DeviceManagementComplianceScheduledActionForRuleRequestBuilder) ScheduledActionConfigurationsById(id string)(*ib3a97c3a0c4a95294f8aac578747e997f8e6d9c2c4b26c8c4d20581b8d3e6cda.DeviceManagementComplianceActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementComplianceActionItem_id"] = id
    }
    return ib3a97c3a0c4a95294f8aac578747e997f8e6d9c2c4b26c8c4d20581b8d3e6cda.NewDeviceManagementComplianceActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
