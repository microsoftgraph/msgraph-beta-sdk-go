package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i047b93bf153214244fb14440d6aeffd28955ea21a303b9d4e0d52ba1e332e59f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/application"
    i7d31074ae7bc6b45526bd1a342929505c009d89110e65c65b006dfe4a5c7fdc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/serviceprincipal"
    i831e60cc7c2552c2fee09dc4f244884f41b0224e9df644857e9a38cd915a14fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/device"
    ic3bd7c901b7a989b1f5fe879c4966614417ee0a68126b04cb7d2bfe39a69d59d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/orgcontact"
    ic8ab70b92cd629ec38e0f94f2bcd6c31e8cb54dc3e8bcae23a377199eb7f3c02 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user"
    ie721b0ebba863e1d11d7fd038683df2ef450d44ab01d5fe07cd21d745dab3627 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/group"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderGetQueryParameters get transitiveMemberOf from me
type DirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application the application property
func (m *DirectoryObjectItemRequestBuilder) Application()(*i047b93bf153214244fb14440d6aeffd28955ea21a303b9d4e0d52ba1e332e59f.ApplicationRequestBuilder) {
    return i047b93bf153214244fb14440d6aeffd28955ea21a303b9d4e0d52ba1e332e59f.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/transitiveMemberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get transitiveMemberOf from me
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get transitiveMemberOf from me
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Device the device property
func (m *DirectoryObjectItemRequestBuilder) Device()(*i831e60cc7c2552c2fee09dc4f244884f41b0224e9df644857e9a38cd915a14fc.DeviceRequestBuilder) {
    return i831e60cc7c2552c2fee09dc4f244884f41b0224e9df644857e9a38cd915a14fc.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get transitiveMemberOf from me
func (m *DirectoryObjectItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get transitiveMemberOf from me
func (m *DirectoryObjectItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Group the group property
func (m *DirectoryObjectItemRequestBuilder) Group()(*ie721b0ebba863e1d11d7fd038683df2ef450d44ab01d5fe07cd21d745dab3627.GroupRequestBuilder) {
    return ie721b0ebba863e1d11d7fd038683df2ef450d44ab01d5fe07cd21d745dab3627.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*ic3bd7c901b7a989b1f5fe879c4966614417ee0a68126b04cb7d2bfe39a69d59d.OrgContactRequestBuilder) {
    return ic3bd7c901b7a989b1f5fe879c4966614417ee0a68126b04cb7d2bfe39a69d59d.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i7d31074ae7bc6b45526bd1a342929505c009d89110e65c65b006dfe4a5c7fdc2.ServicePrincipalRequestBuilder) {
    return i7d31074ae7bc6b45526bd1a342929505c009d89110e65c65b006dfe4a5c7fdc2.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*ic8ab70b92cd629ec38e0f94f2bcd6c31e8cb54dc3e8bcae23a377199eb7f3c02.UserRequestBuilder) {
    return ic8ab70b92cd629ec38e0f94f2bcd6c31e8cb54dc3e8bcae23a377199eb7f3c02.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
