package memberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0694eb6bdf5fef36355af16d5569c6d000fe1679769e5447387bd8d9d5623658 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/orgcontact"
    i5a4a618fb54508b7b57d3b1311532d3ab7e9e75f5182ca5c62821edfcaacc998 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/serviceprincipal"
    i5bb42c4b9e80057a3ee8521764dfd56fe962c1f2c52f75158c724925d99f26c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/group"
    i6884b0ae40cf12bddfb97675e0132b59cd0d55fba037c1fd2ed43a523ffa1822 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/device"
    i8ac3eb1b7f4e194d276a0d0d621af0442020699412767bb09a5d5601e1b889fc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/user"
    ib4c7f7a18d56d667bf68eec3f5015bda0262fd6c5b0788331d4a87c1f63d09e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/count"
    ic3a28a341e9326b332ccc7fb5e9235d01b88740d8db5ed03f57ee2976b5d93c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/application"
)

// MemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.user entity.
type MemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MemberOfRequestBuilderGetQueryParameters the groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Application()(*ic3a28a341e9326b332ccc7fb5e9235d01b88740d8db5ed03f57ee2976b5d93c0.ApplicationRequestBuilder) {
    return ic3a28a341e9326b332ccc7fb5e9235d01b88740d8db5ed03f57ee2976b5d93c0.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/memberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
// Count the Count property
func (m *MemberOfRequestBuilder) Count()(*ib4c7f7a18d56d667bf68eec3f5015bda0262fd6c5b0788331d4a87c1f63d09e3.CountRequestBuilder) {
    return ib4c7f7a18d56d667bf68eec3f5015bda0262fd6c5b0788331d4a87c1f63d09e3.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MemberOfRequestBuilder) Device()(*i6884b0ae40cf12bddfb97675e0132b59cd0d55fba037c1fd2ed43a523ffa1822.DeviceRequestBuilder) {
    return i6884b0ae40cf12bddfb97675e0132b59cd0d55fba037c1fd2ed43a523ffa1822.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Group the group property
func (m *MemberOfRequestBuilder) Group()(*i5bb42c4b9e80057a3ee8521764dfd56fe962c1f2c52f75158c724925d99f26c9.GroupRequestBuilder) {
    return i5bb42c4b9e80057a3ee8521764dfd56fe962c1f2c52f75158c724925d99f26c9.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MemberOfRequestBuilder) OrgContact()(*i0694eb6bdf5fef36355af16d5569c6d000fe1679769e5447387bd8d9d5623658.OrgContactRequestBuilder) {
    return i0694eb6bdf5fef36355af16d5569c6d000fe1679769e5447387bd8d9d5623658.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MemberOfRequestBuilder) ServicePrincipal()(*i5a4a618fb54508b7b57d3b1311532d3ab7e9e75f5182ca5c62821edfcaacc998.ServicePrincipalRequestBuilder) {
    return i5a4a618fb54508b7b57d3b1311532d3ab7e9e75f5182ca5c62821edfcaacc998.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MemberOfRequestBuilder) User()(*i8ac3eb1b7f4e194d276a0d0d621af0442020699412767bb09a5d5601e1b889fc.UserRequestBuilder) {
    return i8ac3eb1b7f4e194d276a0d0d621af0442020699412767bb09a5d5601e1b889fc.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
