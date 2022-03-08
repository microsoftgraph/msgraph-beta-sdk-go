package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3126bfb1a915f0c85e8910f85b7fbdc19b25d34462fa8ede5a13c297a190602c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicytemplates/item/settingtemplates"
    ifeb77407d3839d95c54629e5156854062d7a5a2e76728f24033dce5b2d483dd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicytemplates/item/settingtemplates/item"
)

// DeviceManagementConfigurationPolicyTemplateItemRequestBuilder provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
type DeviceManagementConfigurationPolicyTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteOptions options for Delete
type DeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetOptions options for Get
type DeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetQueryParameters list of all templates
type DeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchOptions options for Patch
type DeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicyTemplateable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal instantiates a new DeviceManagementConfigurationPolicyTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    m := &DeviceManagementConfigurationPolicyTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementConfigurationPolicyTemplateItemRequestBuilder instantiates a new DeviceManagementConfigurationPolicyTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplateItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property configurationPolicyTemplates for deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation list of all templates
func (m *DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) CreateGetRequestInformation(options *DeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property configurationPolicyTemplates in deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property configurationPolicyTemplates for deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) Delete(options *DeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteOptions)(error) {
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
// Get list of all templates
func (m *DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) Get(options *DeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicyTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicyTemplateable), nil
}
// Patch update the navigation property configurationPolicyTemplates in deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) Patch(options *DeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) SettingTemplates()(*i3126bfb1a915f0c85e8910f85b7fbdc19b25d34462fa8ede5a13c297a190602c.SettingTemplatesRequestBuilder) {
    return i3126bfb1a915f0c85e8910f85b7fbdc19b25d34462fa8ede5a13c297a190602c.NewSettingTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.configurationPolicyTemplates.item.settingTemplates.item collection
func (m *DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) SettingTemplatesById(id string)(*ifeb77407d3839d95c54629e5156854062d7a5a2e76728f24033dce5b2d483dd1.DeviceManagementConfigurationSettingTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingTemplate_id"] = id
    }
    return ifeb77407d3839d95c54629e5156854062d7a5a2e76728f24033dce5b2d483dd1.NewDeviceManagementConfigurationSettingTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
