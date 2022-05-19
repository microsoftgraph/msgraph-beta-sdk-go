package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3159f1c2c541350a1a7f0eaefed1347bf08c2363c20af892006aad187437e5cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user"
    i6a37b4f1a711df897cc9f075aac3304fdd5cd2fe8ed29f7435e12744a84e9988 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/orgcontact"
    i9be85b2c79e098646235933868db11702f35703119688df18d5951d8462c90bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/serviceprincipal"
    ia6c76b9433e399bd8415b72b374362d17c4a4903f68374eb723bdbc7e6db046b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/application"
    ic91d890f9e01d27cbc10ab2d481a029c3a277120f04340a2ac34616319f9edfe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/device"
    id791222e31187e8693f2ade5f1169f95066d8351d37a4e029bd29f09d31d942d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/group"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
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
func (m *DirectoryObjectItemRequestBuilder) Application()(*ia6c76b9433e399bd8415b72b374362d17c4a4903f68374eb723bdbc7e6db046b.ApplicationRequestBuilder) {
    return ia6c76b9433e399bd8415b72b374362d17c4a4903f68374eb723bdbc7e6db046b.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DirectoryObjectItemRequestBuilder) Device()(*ic91d890f9e01d27cbc10ab2d481a029c3a277120f04340a2ac34616319f9edfe.DeviceRequestBuilder) {
    return ic91d890f9e01d27cbc10ab2d481a029c3a277120f04340a2ac34616319f9edfe.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
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
func (m *DirectoryObjectItemRequestBuilder) Group()(*id791222e31187e8693f2ade5f1169f95066d8351d37a4e029bd29f09d31d942d.GroupRequestBuilder) {
    return id791222e31187e8693f2ade5f1169f95066d8351d37a4e029bd29f09d31d942d.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*i6a37b4f1a711df897cc9f075aac3304fdd5cd2fe8ed29f7435e12744a84e9988.OrgContactRequestBuilder) {
    return i6a37b4f1a711df897cc9f075aac3304fdd5cd2fe8ed29f7435e12744a84e9988.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i9be85b2c79e098646235933868db11702f35703119688df18d5951d8462c90bc.ServicePrincipalRequestBuilder) {
    return i9be85b2c79e098646235933868db11702f35703119688df18d5951d8462c90bc.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i3159f1c2c541350a1a7f0eaefed1347bf08c2363c20af892006aad187437e5cd.UserRequestBuilder) {
    return i3159f1c2c541350a1a7f0eaefed1347bf08c2363c20af892006aad187437e5cd.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
