package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/userstatuses"
    i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/groupassignments"
    ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assign"
    ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assignments"
    if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/devicestatuses"
    i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/devicestatuses/item"
    i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/groupassignments/item"
    i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/userstatuses/item"
    ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assignments/item"
)

// IosLobAppProvisioningConfigurationItemRequestBuilder provides operations to manage the iosLobAppProvisioningConfigurations property of the microsoft.graph.deviceAppManagement entity.
type IosLobAppProvisioningConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters the IOS Lob App Provisioning Configurations.
type IosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters
}
// IosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Assign()(*ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85.AssignRequestBuilder) {
    return ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Assignments()(*ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567.AssignmentsRequestBuilder) {
    return ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.assignments.item collection
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) AssignmentsById(id string)(*ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0.IosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosLobAppProvisioningConfigurationAssignment%2Did"] = id
    }
    return ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0.NewIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIosLobAppProvisioningConfigurationItemRequestBuilderInternal instantiates a new IosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosLobAppProvisioningConfigurationItemRequestBuilder) {
    m := &IosLobAppProvisioningConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIosLobAppProvisioningConfigurationItemRequestBuilder instantiates a new IosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosLobAppProvisioningConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosLobAppProvisioningConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *IosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *IosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *IosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceStatuses the deviceStatuses property
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatuses()(*if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5.DeviceStatusesRequestBuilder) {
    return if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.deviceStatuses.item collection
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8.ManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus%2Did"] = id
    }
    return i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8.NewManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable), nil
}
// GroupAssignments the groupAssignments property
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignments()(*i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0.GroupAssignmentsRequestBuilder) {
    return i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.groupAssignments.item collection
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignmentsById(id string)(*i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19.MobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppProvisioningConfigGroupAssignment%2Did"] = id
    }
    return i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19.NewMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *IosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UserStatuses the userStatuses property
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) UserStatuses()(*i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929.UserStatusesRequestBuilder) {
    return i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.userStatuses.item collection
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) UserStatusesById(id string)(*i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d.ManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus%2Did"] = id
    }
    return i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d.NewManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
