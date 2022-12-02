package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilder provides operations to count the resources in the collection.
type DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewDeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilder) {
    m := &DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceCompliancePolicyStates/$count";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewDeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the number of the resource
func (m *DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "text/plain"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the number of the resource
func (m *DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
