package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal instantiates a new SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder and sets the default values.
func NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    m := &DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/microsoft.graph.setAndroidDeviceOwnerFullyManagedEnrollmentState";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder instantiates a new SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder and sets the default values.
func NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyable, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyable, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
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
