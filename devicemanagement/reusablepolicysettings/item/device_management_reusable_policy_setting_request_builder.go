package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1fb6f604a131d7c169ff13705125709e219cf9738fee66eca2a318c9053ec2a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies"
    if265d81c2f0cf9dd01c87ec63f8843a097aa03080289938fb1e32e386d8e965c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/clone"
    i3f11bffb7dba7ee4ad416d2fb7daa8d13e95c1d6616fb962c02e09217e9c6cf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies/item"
)

// DeviceManagementReusablePolicySettingRequestBuilder builds and executes requests for operations under \deviceManagement\reusablePolicySettings\{deviceManagementReusablePolicySetting-id}
type DeviceManagementReusablePolicySettingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementReusablePolicySettingRequestBuilderDeleteOptions options for Delete
type DeviceManagementReusablePolicySettingRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementReusablePolicySettingRequestBuilderGetOptions options for Get
type DeviceManagementReusablePolicySettingRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementReusablePolicySettingRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementReusablePolicySettingRequestBuilderGetQueryParameters list of all reusable settings that can be referred in a policy
type DeviceManagementReusablePolicySettingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementReusablePolicySettingRequestBuilderPatchOptions options for Patch
type DeviceManagementReusablePolicySettingRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReusablePolicySetting;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementReusablePolicySettingRequestBuilder) Clone()(*if265d81c2f0cf9dd01c87ec63f8843a097aa03080289938fb1e32e386d8e965c.CloneRequestBuilder) {
    return if265d81c2f0cf9dd01c87ec63f8843a097aa03080289938fb1e32e386d8e965c.NewCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementReusablePolicySettingRequestBuilderInternal instantiates a new DeviceManagementReusablePolicySettingRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementReusablePolicySettingRequestBuilder) {
    m := &DeviceManagementReusablePolicySettingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementReusablePolicySettingRequestBuilder instantiates a new DeviceManagementReusablePolicySettingRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementReusablePolicySettingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementReusablePolicySettingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation list of all reusable settings that can be referred in a policy
func (m *DeviceManagementReusablePolicySettingRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementReusablePolicySettingRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation list of all reusable settings that can be referred in a policy
func (m *DeviceManagementReusablePolicySettingRequestBuilder) CreateGetRequestInformation(options *DeviceManagementReusablePolicySettingRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation list of all reusable settings that can be referred in a policy
func (m *DeviceManagementReusablePolicySettingRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementReusablePolicySettingRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete list of all reusable settings that can be referred in a policy
func (m *DeviceManagementReusablePolicySettingRequestBuilder) Delete(options *DeviceManagementReusablePolicySettingRequestBuilderDeleteOptions)(error) {
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
// Get list of all reusable settings that can be referred in a policy
func (m *DeviceManagementReusablePolicySettingRequestBuilder) Get(options *DeviceManagementReusablePolicySettingRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReusablePolicySetting, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementReusablePolicySetting() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReusablePolicySetting), nil
}
// Patch list of all reusable settings that can be referred in a policy
func (m *DeviceManagementReusablePolicySettingRequestBuilder) Patch(options *DeviceManagementReusablePolicySettingRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementReusablePolicySettingRequestBuilder) ReferencingConfigurationPolicies()(*i1fb6f604a131d7c169ff13705125709e219cf9738fee66eca2a318c9053ec2a5.ReferencingConfigurationPoliciesRequestBuilder) {
    return i1fb6f604a131d7c169ff13705125709e219cf9738fee66eca2a318c9053ec2a5.NewReferencingConfigurationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReferencingConfigurationPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.reusablePolicySettings.item.referencingConfigurationPolicies.item collection
func (m *DeviceManagementReusablePolicySettingRequestBuilder) ReferencingConfigurationPoliciesById(id string)(*i3f11bffb7dba7ee4ad416d2fb7daa8d13e95c1d6616fb962c02e09217e9c6cf7.DeviceManagementConfigurationPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicy_id"] = id
    }
    return i3f11bffb7dba7ee4ad416d2fb7daa8d13e95c1d6616fb962c02e09217e9c6cf7.NewDeviceManagementConfigurationPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
