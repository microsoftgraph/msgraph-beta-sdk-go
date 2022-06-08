package device

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i22874e26517cff84536ae66f9686685f7d78b6bd7a946c0cd2d6ef2f8176e751 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof/item/device/getmembergroups"
    i2dba819dbcb85b5f029d9bc1f9bc0ab0d8b70c7a754cb87244326843ca2995b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof/item/device/checkmemberobjects"
    i67a4ae9241d453f560a4049397027690c3d0e595cee0d15c3ffa884dc35e6919 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof/item/device/restore"
    i775f5aa7d4b1d4a8e157bc516eeef0a784579c8abd01551ae6e6e5eb965f8933 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof/item/device/getmemberobjects"
    ic41446b6f2f08d411bf1580ca91227657eb89947c6359e46c6f61bef5532ed87 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof/item/device/checkmembergroups"
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
func (m *DeviceRequestBuilder) CheckMemberGroups()(*ic41446b6f2f08d411bf1580ca91227657eb89947c6359e46c6f61bef5532ed87.CheckMemberGroupsRequestBuilder) {
    return ic41446b6f2f08d411bf1580ca91227657eb89947c6359e46c6f61bef5532ed87.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *DeviceRequestBuilder) CheckMemberObjects()(*i2dba819dbcb85b5f029d9bc1f9bc0ab0d8b70c7a754cb87244326843ca2995b7.CheckMemberObjectsRequestBuilder) {
    return i2dba819dbcb85b5f029d9bc1f9bc0ab0d8b70c7a754cb87244326843ca2995b7.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.device{?%24select,%24expand}";
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
func (m *DeviceRequestBuilder) GetMemberGroups()(*i22874e26517cff84536ae66f9686685f7d78b6bd7a946c0cd2d6ef2f8176e751.GetMemberGroupsRequestBuilder) {
    return i22874e26517cff84536ae66f9686685f7d78b6bd7a946c0cd2d6ef2f8176e751.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *DeviceRequestBuilder) GetMemberObjects()(*i775f5aa7d4b1d4a8e157bc516eeef0a784579c8abd01551ae6e6e5eb965f8933.GetMemberObjectsRequestBuilder) {
    return i775f5aa7d4b1d4a8e157bc516eeef0a784579c8abd01551ae6e6e5eb965f8933.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *DeviceRequestBuilder) Restore()(*i67a4ae9241d453f560a4049397027690c3d0e595cee0d15c3ffa884dc35e6919.RestoreRequestBuilder) {
    return i67a4ae9241d453f560a4049397027690c3d0e595cee0d15c3ffa884dc35e6919.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
