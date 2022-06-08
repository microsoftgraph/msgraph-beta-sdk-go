package memberof

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i11a03e3cd98a41d3133212e0d0e4354d795246d78c893fe5e4502e98ac9f1dc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/application"
    i23a8075beacdbb88ce482d88c5419722ea6a6af2f9160878d5b78798da261a47 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/count"
    i495d0d66801998afef5a955db7d8f2b29e91e7df866b6501adee923096040482 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/group"
    i7f72ee123f0438a066bade583cd96e6d4ffb9b127e7e4f0bc5fca74f5a853022 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/orgcontact"
    i892072d337523cccecef17df2c4f70368f7cbf1db75294bece7ac3015872f95b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/user"
    ic277aa2d812fbb84bd1664747a9bad09ac7e4b33b53696dfdcdf4ac97a35d8e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/serviceprincipal"
    ife10a2335d1b83ee542e278f84cd13d36c36638c56f2956cc43bc5e4dab42b94 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/device"
)

// MemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type MemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MemberOfRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type MemberOfRequestBuilderGetQueryParameters struct {
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
// MemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MemberOfRequestBuilderGetQueryParameters
}
// Application the application property
func (m *MemberOfRequestBuilder) Application()(*i11a03e3cd98a41d3133212e0d0e4354d795246d78c893fe5e4502e98ac9f1dc6.ApplicationRequestBuilder) {
    return i11a03e3cd98a41d3133212e0d0e4354d795246d78c893fe5e4502e98ac9f1dc6.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMemberOfRequestBuilder instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *MemberOfRequestBuilder) Count()(*i23a8075beacdbb88ce482d88c5419722ea6a6af2f9160878d5b78798da261a47.CountRequestBuilder) {
    return i23a8075beacdbb88ce482d88c5419722ea6a6af2f9160878d5b78798da261a47.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MemberOfRequestBuilder) Device()(*ife10a2335d1b83ee542e278f84cd13d36c36638c56f2956cc43bc5e4dab42b94.DeviceRequestBuilder) {
    return ife10a2335d1b83ee542e278f84cd13d36c36638c56f2956cc43bc5e4dab42b94.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *MemberOfRequestBuilder) Group()(*i495d0d66801998afef5a955db7d8f2b29e91e7df866b6501adee923096040482.GroupRequestBuilder) {
    return i495d0d66801998afef5a955db7d8f2b29e91e7df866b6501adee923096040482.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MemberOfRequestBuilder) OrgContact()(*i7f72ee123f0438a066bade583cd96e6d4ffb9b127e7e4f0bc5fca74f5a853022.OrgContactRequestBuilder) {
    return i7f72ee123f0438a066bade583cd96e6d4ffb9b127e7e4f0bc5fca74f5a853022.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MemberOfRequestBuilder) ServicePrincipal()(*ic277aa2d812fbb84bd1664747a9bad09ac7e4b33b53696dfdcdf4ac97a35d8e2.ServicePrincipalRequestBuilder) {
    return ic277aa2d812fbb84bd1664747a9bad09ac7e4b33b53696dfdcdf4ac97a35d8e2.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MemberOfRequestBuilder) User()(*i892072d337523cccecef17df2c4f70368f7cbf1db75294bece7ac3015872f95b.UserRequestBuilder) {
    return i892072d337523cccecef17df2c4f70368f7cbf1db75294bece7ac3015872f95b.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
