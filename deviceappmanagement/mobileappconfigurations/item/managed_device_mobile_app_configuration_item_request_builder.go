package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatussummary"
    i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses"
    ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/assign"
    ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments"
    if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses"
    if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatussummary"
    i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses/item"
    i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments/item"
    i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses/item"
)

// ManagedDeviceMobileAppConfigurationItemRequestBuilder provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type ManagedDeviceMobileAppConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters the Managed Device Mobile Application Configurations.
type ManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Assign()(*ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69.AssignRequestBuilder) {
    return ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Assignments()(*ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f.AssignmentsRequestBuilder) {
    return ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) AssignmentsById(id string)(*i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2.ManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationAssignment%2Did"] = id
    }
    return i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2.NewManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal instantiates a new ManagedDeviceMobileAppConfigurationItemRequestBuilder and sets the default values.
func NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    m := &ManagedDeviceMobileAppConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceMobileAppConfigurationItemRequestBuilder instantiates a new ManagedDeviceMobileAppConfigurationItemRequestBuilder and sets the default values.
func NewManagedDeviceMobileAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property mobileAppConfigurations for deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property mobileAppConfigurations in deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable, requestConfiguration *ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property mobileAppConfigurations for deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatuses()(*if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450.DeviceStatusesRequestBuilder) {
    return if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888.ManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus%2Did"] = id
    }
    return i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888.NewManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatusSummary provides operations to manage the deviceStatusSummary property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatusSummary()(*i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b.DeviceStatusSummaryRequestBuilder) {
    return i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b.NewDeviceStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable), nil
}
// Patch update the navigation property mobileAppConfigurations in deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable, requestConfiguration *ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable), nil
}
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatuses()(*i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0.UserStatusesRequestBuilder) {
    return i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatusesById(id string)(*i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112.ManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus%2Did"] = id
    }
    return i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112.NewManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStatusSummary provides operations to manage the userStatusSummary property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatusSummary()(*if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4.UserStatusSummaryRequestBuilder) {
    return if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4.NewUserStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
