package removedevicefirmwareconfigurationinterfacemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
type RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal instantiates a new RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    m := &RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/microsoft.graph.removeDeviceFirmwareConfigurationInterfaceManagement";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder instantiates a new RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation remove device from Device Firmware Configuration Interface management
func (m *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration remove device from Device Firmware Configuration Interface management
func (m *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post remove device from Device Firmware Configuration Interface management
func (m *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) Post(ctx context.Context, requestConfiguration *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
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
