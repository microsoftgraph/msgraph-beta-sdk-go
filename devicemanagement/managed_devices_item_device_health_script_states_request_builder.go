package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
type ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetQueryParameters results of device health scripts that ran for this device. Default is empty list. This property is read-only.
type ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetQueryParameters
}
// NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal instantiates a new DeviceHealthScriptStatesRequestBuilder and sets the default values.
func NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    m := &ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceHealthScriptStates{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilder instantiates a new DeviceHealthScriptStatesRequestBuilder and sets the default values.
func NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) Count()(*ManagedDevicesItemDeviceHealthScriptStatesCountRequestBuilder) {
    return NewManagedDevicesItemDeviceHealthScriptStatesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get results of device health scripts that ran for this device. Default is empty list. This property is read-only.
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptPolicyStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateCollectionResponseable), nil
}
// ToGetRequestInformation results of device health scripts that ran for this device. Default is empty list. This property is read-only.
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithIdWithPolicyIdWithDeviceId provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) WithIdWithPolicyIdWithDeviceId(deviceId *string, id *string, policyId *string)(*ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    return NewManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, deviceId, id, policyId)
}
