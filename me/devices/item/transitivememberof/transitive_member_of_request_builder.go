package transitivememberof

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i10acb4cee6b9488ae73a93a21517f4af31aecbd223500dab9b2c95e6e7bf4dfa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/orgcontact"
    i20358954c676fa47b5af4375ef0e74ccb099b8ee0eb19cfa71c0a7d56105e6a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/count"
    i50a444adb6af12fa4ff8957856f1a622546cb24ab8e5086c6b62f31d5e000ea8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/user"
    id67e0b0983394bce5e2e6e59d7c1a4a97610e4ee3af2a4d8f47ead0716b901b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/serviceprincipal"
    idca213c52879dcb8df60c5aaf0f889ca62621d56bba27e6ca04965001f474d1f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/application"
    ie04fd15fec81612459a039018d89c8a543498079c9f825ae33a9bf1ae3220904 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/group"
    ifebbe562e4b9ae0c380a6c4ebce74861abba8d3deb1feb0edbc4cab6b2cbe578 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/device"
)

// TransitiveMemberOfRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
type TransitiveMemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TransitiveMemberOfRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
type TransitiveMemberOfRequestBuilderGetQueryParameters struct {
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
// TransitiveMemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransitiveMemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TransitiveMemberOfRequestBuilderGetQueryParameters
}
// Application the application property
func (m *TransitiveMemberOfRequestBuilder) Application()(*idca213c52879dcb8df60c5aaf0f889ca62621d56bba27e6ca04965001f474d1f.ApplicationRequestBuilder) {
    return idca213c52879dcb8df60c5aaf0f889ca62621d56bba27e6ca04965001f474d1f.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTransitiveMemberOfRequestBuilderInternal instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewTransitiveMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMemberOfRequestBuilder) {
    m := &TransitiveMemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/transitiveMemberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTransitiveMemberOfRequestBuilder instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewTransitiveMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransitiveMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *TransitiveMemberOfRequestBuilder) Count()(*i20358954c676fa47b5af4375ef0e74ccb099b8ee0eb19cfa71c0a7d56105e6a5.CountRequestBuilder) {
    return i20358954c676fa47b5af4375ef0e74ccb099b8ee0eb19cfa71c0a7d56105e6a5.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *TransitiveMemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *TransitiveMemberOfRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TransitiveMemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device the device property
func (m *TransitiveMemberOfRequestBuilder) Device()(*ifebbe562e4b9ae0c380a6c4ebce74861abba8d3deb1feb0edbc4cab6b2cbe578.DeviceRequestBuilder) {
    return ifebbe562e4b9ae0c380a6c4ebce74861abba8d3deb1feb0edbc4cab6b2cbe578.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *TransitiveMemberOfRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *TransitiveMemberOfRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *TransitiveMemberOfRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Group the group property
func (m *TransitiveMemberOfRequestBuilder) Group()(*ie04fd15fec81612459a039018d89c8a543498079c9f825ae33a9bf1ae3220904.GroupRequestBuilder) {
    return ie04fd15fec81612459a039018d89c8a543498079c9f825ae33a9bf1ae3220904.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *TransitiveMemberOfRequestBuilder) OrgContact()(*i10acb4cee6b9488ae73a93a21517f4af31aecbd223500dab9b2c95e6e7bf4dfa.OrgContactRequestBuilder) {
    return i10acb4cee6b9488ae73a93a21517f4af31aecbd223500dab9b2c95e6e7bf4dfa.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *TransitiveMemberOfRequestBuilder) ServicePrincipal()(*id67e0b0983394bce5e2e6e59d7c1a4a97610e4ee3af2a4d8f47ead0716b901b6.ServicePrincipalRequestBuilder) {
    return id67e0b0983394bce5e2e6e59d7c1a4a97610e4ee3af2a4d8f47ead0716b901b6.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *TransitiveMemberOfRequestBuilder) User()(*i50a444adb6af12fa4ff8957856f1a622546cb24ab8e5086c6b62f31d5e000ea8.UserRequestBuilder) {
    return i50a444adb6af12fa4ff8957856f1a622546cb24ab8e5086c6b62f31d5e000ea8.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
