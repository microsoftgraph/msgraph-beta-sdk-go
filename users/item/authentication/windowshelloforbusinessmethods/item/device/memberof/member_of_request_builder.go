package memberof

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0309ccb34a26daabb22023701478f4dace2f4ba5bcbcf7d112642f4050b8a97b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/user"
    i1909c424a216ccb7adc982f4d4fe9f7cc0c3053f467e4a8b35c49d4e63c9dde4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/application"
    i1dd2102acb0db88703ea8658b540f6912d87d4486f959cd3c4b7bcb8199c9949 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/group"
    i8dee3eeca41084baa8104711f9a823ea52c7ee682440dd1ec1489b45d2729c44 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/device"
    ib13ae580feb2923855b0daed88bdd65fb503fd2d6a92ea5cdc1e9e4d09227ccd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/serviceprincipal"
    ib766882f7007a63b11874524b5483a1fa9fc52b505c974b7676e4c7e8e6ed7fe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/orgcontact"
    ibd92db897ec4d20226cad1cae20289f5c057e54a3352b8100c29a17acaebf60a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/count"
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
func (m *MemberOfRequestBuilder) Application()(*i1909c424a216ccb7adc982f4d4fe9f7cc0c3053f467e4a8b35c49d4e63c9dde4.ApplicationRequestBuilder) {
    return i1909c424a216ccb7adc982f4d4fe9f7cc0c3053f467e4a8b35c49d4e63c9dde4.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/memberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
func (m *MemberOfRequestBuilder) Count()(*ibd92db897ec4d20226cad1cae20289f5c057e54a3352b8100c29a17acaebf60a.CountRequestBuilder) {
    return ibd92db897ec4d20226cad1cae20289f5c057e54a3352b8100c29a17acaebf60a.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MemberOfRequestBuilder) Device()(*i8dee3eeca41084baa8104711f9a823ea52c7ee682440dd1ec1489b45d2729c44.DeviceRequestBuilder) {
    return i8dee3eeca41084baa8104711f9a823ea52c7ee682440dd1ec1489b45d2729c44.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MemberOfRequestBuilder) Group()(*i1dd2102acb0db88703ea8658b540f6912d87d4486f959cd3c4b7bcb8199c9949.GroupRequestBuilder) {
    return i1dd2102acb0db88703ea8658b540f6912d87d4486f959cd3c4b7bcb8199c9949.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MemberOfRequestBuilder) OrgContact()(*ib766882f7007a63b11874524b5483a1fa9fc52b505c974b7676e4c7e8e6ed7fe.OrgContactRequestBuilder) {
    return ib766882f7007a63b11874524b5483a1fa9fc52b505c974b7676e4c7e8e6ed7fe.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MemberOfRequestBuilder) ServicePrincipal()(*ib13ae580feb2923855b0daed88bdd65fb503fd2d6a92ea5cdc1e9e4d09227ccd.ServicePrincipalRequestBuilder) {
    return ib13ae580feb2923855b0daed88bdd65fb503fd2d6a92ea5cdc1e9e4d09227ccd.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MemberOfRequestBuilder) User()(*i0309ccb34a26daabb22023701478f4dace2f4ba5bcbcf7d112642f4050b8a97b.UserRequestBuilder) {
    return i0309ccb34a26daabb22023701478f4dace2f4ba5bcbcf7d112642f4050b8a97b.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
