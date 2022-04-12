package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i69f80c23199c4402937b20220cf1e92e26db31e6a9495384f6071feff5fb666e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/assignments"
    i9e07f7775b89cd0dc8237114af526cce0bc28e5356bbb8234a9c06ded2bd1877 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/settings"
    iab3edfc9e4992848bcd442122584eb8dcfbf0066b96fb9aee67506b106fea968 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/createcopy"
    ibbfd5cc1d108b0a4813146b6ecc9e4823ba67ae5bf93fdd3fb64f99ca4e2d170 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/assign"
    i5e2155a9211be69588396fffd88611feb903891dc97905be90cf3ece71e3c153 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/assignments/item"
    iedbbb6f68a2c0078d0b3190bbd5f0cd4f2d9a2f8875eead0f7a4e33156ffe5ba "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/settings/item"
)

// DeviceManagementConfigurationPolicyItemRequestBuilder provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
type DeviceManagementConfigurationPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementConfigurationPolicyItemRequestBuilderDeleteOptions options for Delete
type DeviceManagementConfigurationPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// DeviceManagementConfigurationPolicyItemRequestBuilderGetOptions options for Get
type DeviceManagementConfigurationPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// DeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters list of all Configuration policies
type DeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementConfigurationPolicyItemRequestBuilderPatchOptions options for Patch
type DeviceManagementConfigurationPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Assign the assign property
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) Assign()(*ibbfd5cc1d108b0a4813146b6ecc9e4823ba67ae5bf93fdd3fb64f99ca4e2d170.AssignRequestBuilder) {
    return ibbfd5cc1d108b0a4813146b6ecc9e4823ba67ae5bf93fdd3fb64f99ca4e2d170.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) Assignments()(*i69f80c23199c4402937b20220cf1e92e26db31e6a9495384f6071feff5fb666e.AssignmentsRequestBuilder) {
    return i69f80c23199c4402937b20220cf1e92e26db31e6a9495384f6071feff5fb666e.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.configurationPolicies.item.assignments.item collection
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) AssignmentsById(id string)(*i5e2155a9211be69588396fffd88611feb903891dc97905be90cf3ece71e3c153.DeviceManagementConfigurationPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment%2Did"] = id
    }
    return i5e2155a9211be69588396fffd88611feb903891dc97905be90cf3ece71e3c153.NewDeviceManagementConfigurationPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementConfigurationPolicyItemRequestBuilderInternal instantiates a new DeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementConfigurationPolicyItemRequestBuilder) {
    m := &DeviceManagementConfigurationPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementConfigurationPolicyItemRequestBuilder instantiates a new DeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateCopy the createCopy property
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) CreateCopy()(*iab3edfc9e4992848bcd442122584eb8dcfbf0066b96fb9aee67506b106fea968.CreateCopyRequestBuilder) {
    return iab3edfc9e4992848bcd442122584eb8dcfbf0066b96fb9aee67506b106fea968.NewCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property configurationPolicies for deviceManagement
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementConfigurationPolicyItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation list of all Configuration policies
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) CreateGetRequestInformation(options *DeviceManagementConfigurationPolicyItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property configurationPolicies in deviceManagement
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementConfigurationPolicyItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property configurationPolicies for deviceManagement
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) Delete(options *DeviceManagementConfigurationPolicyItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of all Configuration policies
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) Get(options *DeviceManagementConfigurationPolicyItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable), nil
}
// Patch update the navigation property configurationPolicies in deviceManagement
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) Patch(options *DeviceManagementConfigurationPolicyItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Settings the settings property
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) Settings()(*i9e07f7775b89cd0dc8237114af526cce0bc28e5356bbb8234a9c06ded2bd1877.SettingsRequestBuilder) {
    return i9e07f7775b89cd0dc8237114af526cce0bc28e5356bbb8234a9c06ded2bd1877.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.configurationPolicies.item.settings.item collection
func (m *DeviceManagementConfigurationPolicyItemRequestBuilder) SettingsById(id string)(*iedbbb6f68a2c0078d0b3190bbd5f0cd4f2d9a2f8875eead0f7a4e33156ffe5ba.DeviceManagementConfigurationSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting%2Did"] = id
    }
    return iedbbb6f68a2c0078d0b3190bbd5f0cd4f2d9a2f8875eead0f7a4e33156ffe5ba.NewDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
