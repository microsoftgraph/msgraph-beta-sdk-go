package members

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1176955d1572a0a51dfbfc85782b35d96e113417a9a8b9b76a5b852c0f30ef85 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/serviceprincipal"
    i22ccab6c8244ad62a125b878f072e40c80d16b64165a8030d6b6a7788b9fd2e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/application"
    i293474319befec6e6483824f688dd04e17f1e2546d3f2cf12e7dc631925459d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/ref"
    i44f580ab4b7eb0ac15e005c1f03f3e2e3cb6dd57d22b24c3d26938a2154869e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/device"
    i4f3b62050e2e927a03b259d63b9a1a64603bf4e5f72c5a8ec226dbf804874b3b "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/group"
    i5edaf39c5f0daab121a8ef93fee75b5c9ca20d596f5fab3035ce146595fb56be "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/orgcontact"
    i9f0c4d0fb69ef7c8344bc8df9ec10a52d41f4361f0bc3f85f83bd488b8b67a03 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/user"
    if906d1f4f4fc01c30aef844a7051cd8e14d0075a6236da5400a11dc87e7ad7c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/count"
)

// MembersRequestBuilder provides operations to manage the members property of the microsoft.graph.administrativeUnit entity.
type MembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MembersRequestBuilderGetQueryParameters users and groups that are members of this administrative unit. Supports $expand.
type MembersRequestBuilderGetQueryParameters struct {
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
// MembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MembersRequestBuilderGetQueryParameters
}
// Application the application property
func (m *MembersRequestBuilder) Application()(*i22ccab6c8244ad62a125b878f072e40c80d16b64165a8030d6b6a7788b9fd2e0.ApplicationRequestBuilder) {
    return i22ccab6c8244ad62a125b878f072e40c80d16b64165a8030d6b6a7788b9fd2e0.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersRequestBuilder) {
    m := &MembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/administrativeUnits/{administrativeUnit%2Did}/members{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *MembersRequestBuilder) Count()(*if906d1f4f4fc01c30aef844a7051cd8e14d0075a6236da5400a11dc87e7ad7c2.CountRequestBuilder) {
    return if906d1f4f4fc01c30aef844a7051cd8e14d0075a6236da5400a11dc87e7ad7c2.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation users and groups that are members of this administrative unit. Supports $expand.
func (m *MembersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration users and groups that are members of this administrative unit. Supports $expand.
func (m *MembersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MembersRequestBuilder) Device()(*i44f580ab4b7eb0ac15e005c1f03f3e2e3cb6dd57d22b24c3d26938a2154869e2.DeviceRequestBuilder) {
    return i44f580ab4b7eb0ac15e005c1f03f3e2e3cb6dd57d22b24c3d26938a2154869e2.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get users and groups that are members of this administrative unit. Supports $expand.
func (m *MembersRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler users and groups that are members of this administrative unit. Supports $expand.
func (m *MembersRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *MembersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *MembersRequestBuilder) Group()(*i4f3b62050e2e927a03b259d63b9a1a64603bf4e5f72c5a8ec226dbf804874b3b.GroupRequestBuilder) {
    return i4f3b62050e2e927a03b259d63b9a1a64603bf4e5f72c5a8ec226dbf804874b3b.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MembersRequestBuilder) OrgContact()(*i5edaf39c5f0daab121a8ef93fee75b5c9ca20d596f5fab3035ce146595fb56be.OrgContactRequestBuilder) {
    return i5edaf39c5f0daab121a8ef93fee75b5c9ca20d596f5fab3035ce146595fb56be.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the ref property
func (m *MembersRequestBuilder) Ref()(*i293474319befec6e6483824f688dd04e17f1e2546d3f2cf12e7dc631925459d6.RefRequestBuilder) {
    return i293474319befec6e6483824f688dd04e17f1e2546d3f2cf12e7dc631925459d6.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MembersRequestBuilder) ServicePrincipal()(*i1176955d1572a0a51dfbfc85782b35d96e113417a9a8b9b76a5b852c0f30ef85.ServicePrincipalRequestBuilder) {
    return i1176955d1572a0a51dfbfc85782b35d96e113417a9a8b9b76a5b852c0f30ef85.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MembersRequestBuilder) User()(*i9f0c4d0fb69ef7c8344bc8df9ec10a52d41f4361f0bc3f85f83bd488b8b67a03.UserRequestBuilder) {
    return i9f0c4d0fb69ef7c8344bc8df9ec10a52d41f4361f0bc3f85f83bd488b8b67a03.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
