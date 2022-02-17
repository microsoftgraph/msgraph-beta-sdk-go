package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0d5087c3ad32d953f8943c67d689d4c5195c4fb05be95af0629f34f07a60a238 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies/item/createcopy"
    i4d107188a1663176c0b62795195e592630fe2a9313d4bd98f7794f92e3613a93 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies/item/assign"
    icd6db9d110e9b5d6c4c3b85e477b0e48543f90eba93b007ff12a0ab751c123f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies/item/settings"
    id78dbb2875ade726a0e23fd4722ee6f2932822f82b0a577ae6a7f1b5c867b301 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies/item/assignments"
    i6b80f55ae6e72817db04d0c200abcf17df216061f041ed1f058dcca442672332 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies/item/settings/item"
    i7dbe63647e96f08c5b9d73aa6d410307a762d7c0ec3381e98a0284c4ac977af6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies/item/assignments/item"
)

// DeviceManagementConfigurationPolicyRequestBuilder builds and executes requests for operations under \deviceManagement\reusablePolicySettings\{deviceManagementReusablePolicySetting-id}\referencingConfigurationPolicies\{deviceManagementConfigurationPolicy-id}
type DeviceManagementConfigurationPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementConfigurationPolicyRequestBuilderDeleteOptions options for Delete
type DeviceManagementConfigurationPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationPolicyRequestBuilderGetOptions options for Get
type DeviceManagementConfigurationPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementConfigurationPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationPolicyRequestBuilderGetQueryParameters configuration policies referencing the current reusable setting. This property is read-only.
type DeviceManagementConfigurationPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementConfigurationPolicyRequestBuilderPatchOptions options for Patch
type DeviceManagementConfigurationPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Assign()(*i4d107188a1663176c0b62795195e592630fe2a9313d4bd98f7794f92e3613a93.AssignRequestBuilder) {
    return i4d107188a1663176c0b62795195e592630fe2a9313d4bd98f7794f92e3613a93.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Assignments()(*id78dbb2875ade726a0e23fd4722ee6f2932822f82b0a577ae6a7f1b5c867b301.AssignmentsRequestBuilder) {
    return id78dbb2875ade726a0e23fd4722ee6f2932822f82b0a577ae6a7f1b5c867b301.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.reusablePolicySettings.item.referencingConfigurationPolicies.item.assignments.item collection
func (m *DeviceManagementConfigurationPolicyRequestBuilder) AssignmentsById(id string)(*i7dbe63647e96f08c5b9d73aa6d410307a762d7c0ec3381e98a0284c4ac977af6.DeviceManagementConfigurationPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment_id"] = id
    }
    return i7dbe63647e96f08c5b9d73aa6d410307a762d7c0ec3381e98a0284c4ac977af6.NewDeviceManagementConfigurationPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementConfigurationPolicyRequestBuilderInternal instantiates a new DeviceManagementConfigurationPolicyRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationPolicyRequestBuilder) {
    m := &DeviceManagementConfigurationPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting_id}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementConfigurationPolicyRequestBuilder instantiates a new DeviceManagementConfigurationPolicyRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementConfigurationPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceManagementConfigurationPolicyRequestBuilder) CreateCopy()(*i0d5087c3ad32d953f8943c67d689d4c5195c4fb05be95af0629f34f07a60a238.CreateCopyRequestBuilder) {
    return i0d5087c3ad32d953f8943c67d689d4c5195c4fb05be95af0629f34f07a60a238.NewCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementConfigurationPolicyRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementConfigurationPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementConfigurationPolicyRequestBuilder) CreateGetRequestInformation(options *DeviceManagementConfigurationPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementConfigurationPolicyRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementConfigurationPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Delete(options *DeviceManagementConfigurationPolicyRequestBuilderDeleteOptions)(error) {
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
// Get configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Get(options *DeviceManagementConfigurationPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementConfigurationPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicy), nil
}
// Patch configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Patch(options *DeviceManagementConfigurationPolicyRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Settings()(*icd6db9d110e9b5d6c4c3b85e477b0e48543f90eba93b007ff12a0ab751c123f4.SettingsRequestBuilder) {
    return icd6db9d110e9b5d6c4c3b85e477b0e48543f90eba93b007ff12a0ab751c123f4.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.reusablePolicySettings.item.referencingConfigurationPolicies.item.settings.item collection
func (m *DeviceManagementConfigurationPolicyRequestBuilder) SettingsById(id string)(*i6b80f55ae6e72817db04d0c200abcf17df216061f041ed1f058dcca442672332.DeviceManagementConfigurationSettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting_id"] = id
    }
    return i6b80f55ae6e72817db04d0c200abcf17df216061f041ed1f058dcca442672332.NewDeviceManagementConfigurationSettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
