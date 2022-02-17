package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ibf107fd196437aaede650d26155860f2bbbbd3d7235c31cb4f80608cefdeddd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduledactionsforrule/item/scheduledactionconfigurations"
    i30c81441182809790d7deb6c0da1307530c3dbb5567c662d63e456dddcb16f5a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduledactionsforrule/item/scheduledactionconfigurations/item"
)

// DeviceComplianceScheduledActionForRuleRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCompliancePolicies\{deviceCompliancePolicy-id}\scheduledActionsForRule\{deviceComplianceScheduledActionForRule-id}
type DeviceComplianceScheduledActionForRuleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceComplianceScheduledActionForRuleRequestBuilderDeleteOptions options for Delete
type DeviceComplianceScheduledActionForRuleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceComplianceScheduledActionForRuleRequestBuilderGetOptions options for Get
type DeviceComplianceScheduledActionForRuleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceComplianceScheduledActionForRuleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceComplianceScheduledActionForRuleRequestBuilderGetQueryParameters the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
type DeviceComplianceScheduledActionForRuleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceComplianceScheduledActionForRuleRequestBuilderPatchOptions options for Patch
type DeviceComplianceScheduledActionForRuleRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScheduledActionForRule;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceComplianceScheduledActionForRuleRequestBuilderInternal instantiates a new DeviceComplianceScheduledActionForRuleRequestBuilder and sets the default values.
func NewDeviceComplianceScheduledActionForRuleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceComplianceScheduledActionForRuleRequestBuilder) {
    m := &DeviceComplianceScheduledActionForRuleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy_id}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceComplianceScheduledActionForRuleRequestBuilder instantiates a new DeviceComplianceScheduledActionForRuleRequestBuilder and sets the default values.
func NewDeviceComplianceScheduledActionForRuleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceComplianceScheduledActionForRuleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceComplianceScheduledActionForRuleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceComplianceScheduledActionForRuleRequestBuilder) CreateDeleteRequestInformation(options *DeviceComplianceScheduledActionForRuleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceComplianceScheduledActionForRuleRequestBuilder) CreateGetRequestInformation(options *DeviceComplianceScheduledActionForRuleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceComplianceScheduledActionForRuleRequestBuilder) CreatePatchRequestInformation(options *DeviceComplianceScheduledActionForRuleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceComplianceScheduledActionForRuleRequestBuilder) Delete(options *DeviceComplianceScheduledActionForRuleRequestBuilderDeleteOptions)(error) {
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
// Get the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceComplianceScheduledActionForRuleRequestBuilder) Get(options *DeviceComplianceScheduledActionForRuleRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScheduledActionForRule, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceComplianceScheduledActionForRule() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScheduledActionForRule), nil
}
// Patch the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceComplianceScheduledActionForRuleRequestBuilder) Patch(options *DeviceComplianceScheduledActionForRuleRequestBuilderPatchOptions)(error) {
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
func (m *DeviceComplianceScheduledActionForRuleRequestBuilder) ScheduledActionConfigurations()(*ibf107fd196437aaede650d26155860f2bbbbd3d7235c31cb4f80608cefdeddd9.ScheduledActionConfigurationsRequestBuilder) {
    return ibf107fd196437aaede650d26155860f2bbbbd3d7235c31cb4f80608cefdeddd9.NewScheduledActionConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCompliancePolicies.item.scheduledActionsForRule.item.scheduledActionConfigurations.item collection
func (m *DeviceComplianceScheduledActionForRuleRequestBuilder) ScheduledActionConfigurationsById(id string)(*i30c81441182809790d7deb6c0da1307530c3dbb5567c662d63e456dddcb16f5a.DeviceComplianceActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceActionItem_id"] = id
    }
    return i30c81441182809790d7deb6c0da1307530c3dbb5567c662d63e456dddcb16f5a.NewDeviceComplianceActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
