package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1fb6f604a131d7c169ff13705125709e219cf9738fee66eca2a318c9053ec2a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies"
    if265d81c2f0cf9dd01c87ec63f8843a097aa03080289938fb1e32e386d8e965c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/clone"
    i3f11bffb7dba7ee4ad416d2fb7daa8d13e95c1d6616fb962c02e09217e9c6cf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies/item"
)

// DeviceManagementReusablePolicySettingItemRequestBuilder provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
type DeviceManagementReusablePolicySettingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementReusablePolicySettingItemRequestBuilderDeleteOptions options for Delete
type DeviceManagementReusablePolicySettingItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementReusablePolicySettingItemRequestBuilderGetOptions options for Get
type DeviceManagementReusablePolicySettingItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementReusablePolicySettingItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementReusablePolicySettingItemRequestBuilderGetQueryParameters list of all reusable settings that can be referred in a policy
type DeviceManagementReusablePolicySettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementReusablePolicySettingItemRequestBuilderPatchOptions options for Patch
type DeviceManagementReusablePolicySettingItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReusablePolicySettingable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementReusablePolicySettingItemRequestBuilder) Clone()(*if265d81c2f0cf9dd01c87ec63f8843a097aa03080289938fb1e32e386d8e965c.CloneRequestBuilder) {
    return if265d81c2f0cf9dd01c87ec63f8843a097aa03080289938fb1e32e386d8e965c.NewCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementReusablePolicySettingItemRequestBuilderInternal instantiates a new DeviceManagementReusablePolicySettingItemRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementReusablePolicySettingItemRequestBuilder) {
    m := &DeviceManagementReusablePolicySettingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementReusablePolicySettingItemRequestBuilder instantiates a new DeviceManagementReusablePolicySettingItemRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementReusablePolicySettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementReusablePolicySettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reusablePolicySettings for deviceManagement
func (m *DeviceManagementReusablePolicySettingItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementReusablePolicySettingItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementReusablePolicySettingItemRequestBuilder) CreateGetRequestInformation(options *DeviceManagementReusablePolicySettingItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property reusablePolicySettings in deviceManagement
func (m *DeviceManagementReusablePolicySettingItemRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementReusablePolicySettingItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property reusablePolicySettings for deviceManagement
func (m *DeviceManagementReusablePolicySettingItemRequestBuilder) Delete(options *DeviceManagementReusablePolicySettingItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of all reusable settings that can be referred in a policy
func (m *DeviceManagementReusablePolicySettingItemRequestBuilder) Get(options *DeviceManagementReusablePolicySettingItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReusablePolicySettingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementReusablePolicySettingable), nil
}
// Patch update the navigation property reusablePolicySettings in deviceManagement
func (m *DeviceManagementReusablePolicySettingItemRequestBuilder) Patch(options *DeviceManagementReusablePolicySettingItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceManagementReusablePolicySettingItemRequestBuilder) ReferencingConfigurationPolicies()(*i1fb6f604a131d7c169ff13705125709e219cf9738fee66eca2a318c9053ec2a5.ReferencingConfigurationPoliciesRequestBuilder) {
    return i1fb6f604a131d7c169ff13705125709e219cf9738fee66eca2a318c9053ec2a5.NewReferencingConfigurationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReferencingConfigurationPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.reusablePolicySettings.item.referencingConfigurationPolicies.item collection
func (m *DeviceManagementReusablePolicySettingItemRequestBuilder) ReferencingConfigurationPoliciesById(id string)(*i3f11bffb7dba7ee4ad416d2fb7daa8d13e95c1d6616fb962c02e09217e9c6cf7.DeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicy_id"] = id
    }
    return i3f11bffb7dba7ee4ad416d2fb7daa8d13e95c1d6616fb962c02e09217e9c6cf7.NewDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
