package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1f5027d4a0b3340e9b76763ff24f75e1048bb0e622b30753747dac9656c128ac "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/assignedaccessmultimodeprofiles"
    i209aefe482b229a24f905c3b3df1cfd2ceae762805e5c2a6fe3e7f25cc4eaaca "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/windowsupdateforbusinessconfiguration"
    i2b1a7ee7081880857923ce04764b008e8e3a4ab9ca045409530a83ca51bae662 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/windowsprivacyaccesscontrols"
    i2e953dd8de78572f5e28792284839af62e550fdedcfa0f926825b76cf5b55055 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/devicestatusoverview"
    i36d04c87e83326af850beeca83554e31e36b4f1009ebc093f35a9f220fa0a4fe "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/devicestatuses"
    i3e5a9c6abd02192200a577db21edb6d8b50ae8e2d9d5504e95e2c286f60c98a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/userstatusoverview"
    i7e7ab7485c32e3c00675ecc9b604c0fe15e0b309a689ee4edec907b80b4f4bf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/assign"
    i93863715f80bdcb18e9dcd490d314ff7871f1e5638be84390a5dc4d361302118 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/assignments"
    ibb320dd6fcf0348de4b597bd59030122dc0039d1b2595aecef7279311d4c7fd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/userstatuses"
    icb32976dea6dad1b5d5ec1e79c2394eeb3ed55998cb534c9c0129e1ccde5a69a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/getomasettingplaintextvaluewithsecretreferencevalueid"
    idf522f9db5866e01531ecd195e3db59649221908dcabfbe0891c3887e2c1e2e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/devicesettingstatesummaries"
    if06d6ff07150b0641b6fde6feeb0dda2440e642bce7d5cbc891b7ac35f6d3a75 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments"
    i3e34b826795238ed9133fb18aa5cec93b37b6f29016c136e739a889f1b977bfd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/devicestatuses/item"
    i4f57a07bf94a6a03edadd0600f1cacbf130f9cea090055c5e429f5f9239c636e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments/item"
    i77e922bc2634c1b149fd5fbc479252b5a9dce4fcc9a42aa64bb2a5b2192e0580 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/devicesettingstatesummaries/item"
    id79155e4004317e808df6e9ae0a7f5ed23f8c89b955537d051b7e20bdfefe0e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/userstatuses/item"
    iec03053e92d095e4b0c7ff7cef308162ec4e18612ecc279c203b4dc1727f41cf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/assignments/item"
)

// DeviceConfigurationRequestBuilder builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}
type DeviceConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceConfigurationRequestBuilderDeleteOptions options for Delete
type DeviceConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceConfigurationRequestBuilderGetOptions options for Get
type DeviceConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceConfigurationRequestBuilderGetQueryParameters the device configurations.
type DeviceConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceConfigurationRequestBuilderPatchOptions options for Patch
type DeviceConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceConfigurationRequestBuilder) Assign()(*i7e7ab7485c32e3c00675ecc9b604c0fe15e0b309a689ee4edec907b80b4f4bf0.AssignRequestBuilder) {
    return i7e7ab7485c32e3c00675ecc9b604c0fe15e0b309a689ee4edec907b80b4f4bf0.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) AssignedAccessMultiModeProfiles()(*i1f5027d4a0b3340e9b76763ff24f75e1048bb0e622b30753747dac9656c128ac.AssignedAccessMultiModeProfilesRequestBuilder) {
    return i1f5027d4a0b3340e9b76763ff24f75e1048bb0e622b30753747dac9656c128ac.NewAssignedAccessMultiModeProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) Assignments()(*i93863715f80bdcb18e9dcd490d314ff7871f1e5638be84390a5dc4d361302118.AssignmentsRequestBuilder) {
    return i93863715f80bdcb18e9dcd490d314ff7871f1e5638be84390a5dc4d361302118.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.assignments.item collection
func (m *DeviceConfigurationRequestBuilder) AssignmentsById(id string)(*iec03053e92d095e4b0c7ff7cef308162ec4e18612ecc279c203b4dc1727f41cf.DeviceConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationAssignment_id"] = id
    }
    return iec03053e92d095e4b0c7ff7cef308162ec4e18612ecc279c203b4dc1727f41cf.NewDeviceConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceConfigurationRequestBuilderInternal instantiates a new DeviceConfigurationRequestBuilder and sets the default values.
func NewDeviceConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationRequestBuilder) {
    m := &DeviceConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceConfigurationRequestBuilder instantiates a new DeviceConfigurationRequestBuilder and sets the default values.
func NewDeviceConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the device configurations.
func (m *DeviceConfigurationRequestBuilder) CreateDeleteRequestInformation(options *DeviceConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the device configurations.
func (m *DeviceConfigurationRequestBuilder) CreateGetRequestInformation(options *DeviceConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the device configurations.
func (m *DeviceConfigurationRequestBuilder) CreatePatchRequestInformation(options *DeviceConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the device configurations.
func (m *DeviceConfigurationRequestBuilder) Delete(options *DeviceConfigurationRequestBuilderDeleteOptions)(error) {
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
func (m *DeviceConfigurationRequestBuilder) DeviceSettingStateSummaries()(*idf522f9db5866e01531ecd195e3db59649221908dcabfbe0891c3887e2c1e2e5.DeviceSettingStateSummariesRequestBuilder) {
    return idf522f9db5866e01531ecd195e3db59649221908dcabfbe0891c3887e2c1e2e5.NewDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.deviceSettingStateSummaries.item collection
func (m *DeviceConfigurationRequestBuilder) DeviceSettingStateSummariesById(id string)(*i77e922bc2634c1b149fd5fbc479252b5a9dce4fcc9a42aa64bb2a5b2192e0580.SettingStateDeviceSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary_id"] = id
    }
    return i77e922bc2634c1b149fd5fbc479252b5a9dce4fcc9a42aa64bb2a5b2192e0580.NewSettingStateDeviceSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) DeviceStatuses()(*i36d04c87e83326af850beeca83554e31e36b4f1009ebc093f35a9f220fa0a4fe.DeviceStatusesRequestBuilder) {
    return i36d04c87e83326af850beeca83554e31e36b4f1009ebc093f35a9f220fa0a4fe.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.deviceStatuses.item collection
func (m *DeviceConfigurationRequestBuilder) DeviceStatusesById(id string)(*i3e34b826795238ed9133fb18aa5cec93b37b6f29016c136e739a889f1b977bfd.DeviceConfigurationDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationDeviceStatus_id"] = id
    }
    return i3e34b826795238ed9133fb18aa5cec93b37b6f29016c136e739a889f1b977bfd.NewDeviceConfigurationDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) DeviceStatusOverview()(*i2e953dd8de78572f5e28792284839af62e550fdedcfa0f926825b76cf5b55055.DeviceStatusOverviewRequestBuilder) {
    return i2e953dd8de78572f5e28792284839af62e550fdedcfa0f926825b76cf5b55055.NewDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device configurations.
func (m *DeviceConfigurationRequestBuilder) Get(options *DeviceConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceConfiguration() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceConfiguration), nil
}
// GetOmaSettingPlainTextValueWithSecretReferenceValueId builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}\microsoft.graph.getOmaSettingPlainTextValue(secretReferenceValueId='{secretReferenceValueId}')
func (m *DeviceConfigurationRequestBuilder) GetOmaSettingPlainTextValueWithSecretReferenceValueId(secretReferenceValueId *string)(*icb32976dea6dad1b5d5ec1e79c2394eeb3ed55998cb534c9c0129e1ccde5a69a.GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    return icb32976dea6dad1b5d5ec1e79c2394eeb3ed55998cb534c9c0129e1ccde5a69a.NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, secretReferenceValueId);
}
func (m *DeviceConfigurationRequestBuilder) GroupAssignments()(*if06d6ff07150b0641b6fde6feeb0dda2440e642bce7d5cbc891b7ac35f6d3a75.GroupAssignmentsRequestBuilder) {
    return if06d6ff07150b0641b6fde6feeb0dda2440e642bce7d5cbc891b7ac35f6d3a75.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.groupAssignments.item collection
func (m *DeviceConfigurationRequestBuilder) GroupAssignmentsById(id string)(*i4f57a07bf94a6a03edadd0600f1cacbf130f9cea090055c5e429f5f9239c636e.DeviceConfigurationGroupAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationGroupAssignment_id"] = id
    }
    return i4f57a07bf94a6a03edadd0600f1cacbf130f9cea090055c5e429f5f9239c636e.NewDeviceConfigurationGroupAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the device configurations.
func (m *DeviceConfigurationRequestBuilder) Patch(options *DeviceConfigurationRequestBuilderPatchOptions)(error) {
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
func (m *DeviceConfigurationRequestBuilder) UserStatuses()(*ibb320dd6fcf0348de4b597bd59030122dc0039d1b2595aecef7279311d4c7fd5.UserStatusesRequestBuilder) {
    return ibb320dd6fcf0348de4b597bd59030122dc0039d1b2595aecef7279311d4c7fd5.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.userStatuses.item collection
func (m *DeviceConfigurationRequestBuilder) UserStatusesById(id string)(*id79155e4004317e808df6e9ae0a7f5ed23f8c89b955537d051b7e20bdfefe0e7.DeviceConfigurationUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationUserStatus_id"] = id
    }
    return id79155e4004317e808df6e9ae0a7f5ed23f8c89b955537d051b7e20bdfefe0e7.NewDeviceConfigurationUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) UserStatusOverview()(*i3e5a9c6abd02192200a577db21edb6d8b50ae8e2d9d5504e95e2c286f60c98a1.UserStatusOverviewRequestBuilder) {
    return i3e5a9c6abd02192200a577db21edb6d8b50ae8e2d9d5504e95e2c286f60c98a1.NewUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) WindowsPrivacyAccessControls()(*i2b1a7ee7081880857923ce04764b008e8e3a4ab9ca045409530a83ca51bae662.WindowsPrivacyAccessControlsRequestBuilder) {
    return i2b1a7ee7081880857923ce04764b008e8e3a4ab9ca045409530a83ca51bae662.NewWindowsPrivacyAccessControlsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) WindowsUpdateForBusinessConfiguration()(*i209aefe482b229a24f905c3b3df1cfd2ceae762805e5c2a6fe3e7f25cc4eaaca.WindowsUpdateForBusinessConfigurationRequestBuilder) {
    return i209aefe482b229a24f905c3b3df1cfd2ceae762805e5c2a6fe3e7f25cc4eaaca.NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
