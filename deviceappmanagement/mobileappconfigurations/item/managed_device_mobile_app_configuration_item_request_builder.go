package item

import (
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteOptions options for Delete
type ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderGetOptions options for Get
type ManagedDeviceMobileAppConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters the Managed Device Mobile Application Configurations.
type ManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchOptions options for Patch
type ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Assign the assign property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Assign()(*ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69.AssignRequestBuilder) {
    return ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Assignments()(*ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f.AssignmentsRequestBuilder) {
    return ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.assignments.item collection
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) AssignmentsById(id string)(*i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2.ManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationAssignment_id"] = id
    }
    return i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2.NewManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal instantiates a new ManagedDeviceMobileAppConfigurationItemRequestBuilder and sets the default values.
func NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    m := &ManagedDeviceMobileAppConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration_id}{?select,expand}";
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
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property mobileAppConfigurations in deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property mobileAppConfigurations for deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Delete(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteOptions)(error) {
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
// DeviceStatuses the deviceStatuses property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatuses()(*if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450.DeviceStatusesRequestBuilder) {
    return if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.deviceStatuses.item collection
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888.ManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus_id"] = id
    }
    return i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888.NewManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatusSummary the deviceStatusSummary property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatusSummary()(*i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b.DeviceStatusSummaryRequestBuilder) {
    return i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b.NewDeviceStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Get(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable), nil
}
// Patch update the navigation property mobileAppConfigurations in deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Patch(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchOptions)(error) {
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
// UserStatuses the userStatuses property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatuses()(*i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0.UserStatusesRequestBuilder) {
    return i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.userStatuses.item collection
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatusesById(id string)(*i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112.ManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus_id"] = id
    }
    return i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112.NewManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStatusSummary the userStatusSummary property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatusSummary()(*if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4.UserStatusSummaryRequestBuilder) {
    return if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4.NewUserStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
