package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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

// DeviceConfigurationItemRequestBuilder provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceConfigurationItemRequestBuilderGetQueryParameters the device configurations.
type DeviceConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceConfigurationItemRequestBuilderGetQueryParameters
}
// DeviceConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *DeviceConfigurationItemRequestBuilder) Assign()(*i7e7ab7485c32e3c00675ecc9b604c0fe15e0b309a689ee4edec907b80b4f4bf0.AssignRequestBuilder) {
    return i7e7ab7485c32e3c00675ecc9b604c0fe15e0b309a689ee4edec907b80b4f4bf0.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignedAccessMultiModeProfiles the assignedAccessMultiModeProfiles property
func (m *DeviceConfigurationItemRequestBuilder) AssignedAccessMultiModeProfiles()(*i1f5027d4a0b3340e9b76763ff24f75e1048bb0e622b30753747dac9656c128ac.AssignedAccessMultiModeProfilesRequestBuilder) {
    return i1f5027d4a0b3340e9b76763ff24f75e1048bb0e622b30753747dac9656c128ac.NewAssignedAccessMultiModeProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceConfigurationItemRequestBuilder) Assignments()(*i93863715f80bdcb18e9dcd490d314ff7871f1e5638be84390a5dc4d361302118.AssignmentsRequestBuilder) {
    return i93863715f80bdcb18e9dcd490d314ff7871f1e5638be84390a5dc4d361302118.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.assignments.item collection
func (m *DeviceConfigurationItemRequestBuilder) AssignmentsById(id string)(*iec03053e92d095e4b0c7ff7cef308162ec4e18612ecc279c203b4dc1727f41cf.DeviceConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationAssignment%2Did"] = id
    }
    return iec03053e92d095e4b0c7ff7cef308162ec4e18612ecc279c203b4dc1727f41cf.NewDeviceConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceConfigurationItemRequestBuilderInternal instantiates a new DeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationItemRequestBuilder) {
    m := &DeviceConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceConfigurationItemRequestBuilder instantiates a new DeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the device configurations.
func (m *DeviceConfigurationItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the device configurations.
func (m *DeviceConfigurationItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, requestConfiguration *DeviceConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceConfigurationItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceSettingStateSummaries the deviceSettingStateSummaries property
func (m *DeviceConfigurationItemRequestBuilder) DeviceSettingStateSummaries()(*idf522f9db5866e01531ecd195e3db59649221908dcabfbe0891c3887e2c1e2e5.DeviceSettingStateSummariesRequestBuilder) {
    return idf522f9db5866e01531ecd195e3db59649221908dcabfbe0891c3887e2c1e2e5.NewDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.deviceSettingStateSummaries.item collection
func (m *DeviceConfigurationItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*i77e922bc2634c1b149fd5fbc479252b5a9dce4fcc9a42aa64bb2a5b2192e0580.SettingStateDeviceSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary%2Did"] = id
    }
    return i77e922bc2634c1b149fd5fbc479252b5a9dce4fcc9a42aa64bb2a5b2192e0580.NewSettingStateDeviceSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatuses the deviceStatuses property
func (m *DeviceConfigurationItemRequestBuilder) DeviceStatuses()(*i36d04c87e83326af850beeca83554e31e36b4f1009ebc093f35a9f220fa0a4fe.DeviceStatusesRequestBuilder) {
    return i36d04c87e83326af850beeca83554e31e36b4f1009ebc093f35a9f220fa0a4fe.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.deviceStatuses.item collection
func (m *DeviceConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*i3e34b826795238ed9133fb18aa5cec93b37b6f29016c136e739a889f1b977bfd.DeviceConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationDeviceStatus%2Did"] = id
    }
    return i3e34b826795238ed9133fb18aa5cec93b37b6f29016c136e739a889f1b977bfd.NewDeviceConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatusOverview the deviceStatusOverview property
func (m *DeviceConfigurationItemRequestBuilder) DeviceStatusOverview()(*i2e953dd8de78572f5e28792284839af62e550fdedcfa0f926825b76cf5b55055.DeviceStatusOverviewRequestBuilder) {
    return i2e953dd8de78572f5e28792284839af62e550fdedcfa0f926825b76cf5b55055.NewDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device configurations.
func (m *DeviceConfigurationItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetOmaSettingPlainTextValueWithSecretReferenceValueId provides operations to call the getOmaSettingPlainTextValue method.
func (m *DeviceConfigurationItemRequestBuilder) GetOmaSettingPlainTextValueWithSecretReferenceValueId(secretReferenceValueId *string)(*icb32976dea6dad1b5d5ec1e79c2394eeb3ed55998cb534c9c0129e1ccde5a69a.GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    return icb32976dea6dad1b5d5ec1e79c2394eeb3ed55998cb534c9c0129e1ccde5a69a.NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, secretReferenceValueId);
}
// GetWithRequestConfigurationAndResponseHandler the device configurations.
func (m *DeviceConfigurationItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceConfigurationItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable), nil
}
// GroupAssignments the groupAssignments property
func (m *DeviceConfigurationItemRequestBuilder) GroupAssignments()(*if06d6ff07150b0641b6fde6feeb0dda2440e642bce7d5cbc891b7ac35f6d3a75.GroupAssignmentsRequestBuilder) {
    return if06d6ff07150b0641b6fde6feeb0dda2440e642bce7d5cbc891b7ac35f6d3a75.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.groupAssignments.item collection
func (m *DeviceConfigurationItemRequestBuilder) GroupAssignmentsById(id string)(*i4f57a07bf94a6a03edadd0600f1cacbf130f9cea090055c5e429f5f9239c636e.DeviceConfigurationGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationGroupAssignment%2Did"] = id
    }
    return i4f57a07bf94a6a03edadd0600f1cacbf130f9cea090055c5e429f5f9239c636e.NewDeviceConfigurationGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceConfigurationItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, requestConfiguration *DeviceConfigurationItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UserStatuses the userStatuses property
func (m *DeviceConfigurationItemRequestBuilder) UserStatuses()(*ibb320dd6fcf0348de4b597bd59030122dc0039d1b2595aecef7279311d4c7fd5.UserStatusesRequestBuilder) {
    return ibb320dd6fcf0348de4b597bd59030122dc0039d1b2595aecef7279311d4c7fd5.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceConfigurations.item.userStatuses.item collection
func (m *DeviceConfigurationItemRequestBuilder) UserStatusesById(id string)(*id79155e4004317e808df6e9ae0a7f5ed23f8c89b955537d051b7e20bdfefe0e7.DeviceConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationUserStatus%2Did"] = id
    }
    return id79155e4004317e808df6e9ae0a7f5ed23f8c89b955537d051b7e20bdfefe0e7.NewDeviceConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStatusOverview the userStatusOverview property
func (m *DeviceConfigurationItemRequestBuilder) UserStatusOverview()(*i3e5a9c6abd02192200a577db21edb6d8b50ae8e2d9d5504e95e2c286f60c98a1.UserStatusOverviewRequestBuilder) {
    return i3e5a9c6abd02192200a577db21edb6d8b50ae8e2d9d5504e95e2c286f60c98a1.NewUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsPrivacyAccessControls the windowsPrivacyAccessControls property
func (m *DeviceConfigurationItemRequestBuilder) WindowsPrivacyAccessControls()(*i2b1a7ee7081880857923ce04764b008e8e3a4ab9ca045409530a83ca51bae662.WindowsPrivacyAccessControlsRequestBuilder) {
    return i2b1a7ee7081880857923ce04764b008e8e3a4ab9ca045409530a83ca51bae662.NewWindowsPrivacyAccessControlsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsUpdateForBusinessConfiguration the windowsUpdateForBusinessConfiguration property
func (m *DeviceConfigurationItemRequestBuilder) WindowsUpdateForBusinessConfiguration()(*i209aefe482b229a24f905c3b3df1cfd2ceae762805e5c2a6fe3e7f25cc4eaaca.WindowsUpdateForBusinessConfigurationRequestBuilder) {
    return i209aefe482b229a24f905c3b3df1cfd2ceae762805e5c2a6fe3e7f25cc4eaaca.NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
