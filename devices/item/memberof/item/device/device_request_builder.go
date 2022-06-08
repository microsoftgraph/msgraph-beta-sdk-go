package device

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i15f0ef84e9dd020e0700a6b8541e858a1499ae6f6af3d56605202ed0fc6f8fd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof/item/device/getmembergroups"
    i38a7cca93febf03979f28bf3cb101eb2008f9f8634f17a1b19874f0bfae9fcfe "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof/item/device/checkmemberobjects"
    i939c002158e5f3c191859e5365bf0e256efb6fd3f9556346c6f24baf74a3c457 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof/item/device/checkmembergroups"
    if90cb3036096b5dca3d4bd9d255adc665929cf80ca8ba2e78bbad720c94db9d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof/item/device/restore"
    if9dcbd6e4e8b6c73dc4de0a6f92961478682776da3c8244f6e49c36351603b4c "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof/item/device/getmemberobjects"
)

// DeviceRequestBuilder casts the previous resource to device.
type DeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.device
type DeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceRequestBuilderGetQueryParameters
}
// CheckMemberGroups the checkMemberGroups property
func (m *DeviceRequestBuilder) CheckMemberGroups()(*i939c002158e5f3c191859e5365bf0e256efb6fd3f9556346c6f24baf74a3c457.CheckMemberGroupsRequestBuilder) {
    return i939c002158e5f3c191859e5365bf0e256efb6fd3f9556346c6f24baf74a3c457.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *DeviceRequestBuilder) CheckMemberObjects()(*i38a7cca93febf03979f28bf3cb101eb2008f9f8634f17a1b19874f0bfae9fcfe.CheckMemberObjectsRequestBuilder) {
    return i38a7cca93febf03979f28bf3cb101eb2008f9f8634f17a1b19874f0bfae9fcfe.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.device{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceRequestBuilder instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.device
func (m *DeviceRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the item of type microsoft.graph.directoryObject as microsoft.graph.device
func (m *DeviceRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.device
func (m *DeviceRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *DeviceRequestBuilder) GetMemberGroups()(*i15f0ef84e9dd020e0700a6b8541e858a1499ae6f6af3d56605202ed0fc6f8fd6.GetMemberGroupsRequestBuilder) {
    return i15f0ef84e9dd020e0700a6b8541e858a1499ae6f6af3d56605202ed0fc6f8fd6.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *DeviceRequestBuilder) GetMemberObjects()(*if9dcbd6e4e8b6c73dc4de0a6f92961478682776da3c8244f6e49c36351603b4c.GetMemberObjectsRequestBuilder) {
    return if9dcbd6e4e8b6c73dc4de0a6f92961478682776da3c8244f6e49c36351603b4c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the item of type microsoft.graph.directoryObject as microsoft.graph.device
func (m *DeviceRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable), nil
}
// Restore the restore property
func (m *DeviceRequestBuilder) Restore()(*if90cb3036096b5dca3d4bd9d255adc665929cf80ca8ba2e78bbad720c94db9d2.RestoreRequestBuilder) {
    return if90cb3036096b5dca3d4bd9d255adc665929cf80ca8ba2e78bbad720c94db9d2.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
