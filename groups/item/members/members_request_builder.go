package members

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2715231aa1b97ebc0e299d1c80b5aa685b14ad2cbadc52b883e53096df40d25d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/group"
    i5b707f4d1a3b6bb92ef009b5a12f4aa2f906e2a505066d14fad393f89f2d88b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/orgcontact"
    i93eb95fc5cc21c1b7ff87b7fc9a1ccaac4340ca6d76b73e5c766740fae74fa58 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/device"
    ia4ba8805ab0ada495b5c5fbb898da91d5e48daac76a44a52a6a3d8a1abfe0c8e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/ref"
    ice345922f6e79bd547a80a813a84ebd47224ff22ce94579cf0aeee901a020457 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/serviceprincipal"
    idc6ec6a1bedd6935f4390a1b2911a367f6435ef357cd42c0c7fa2134416bacb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/count"
    ie29d88f3895d7a38a02ccfe9f35c806f95afb3d0eb9bea6d555f522c74a11049 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/application"
    if0c842f2c2a020efe10aa63f93518dcf5ddfb94911ef5e1d22fe197c5a4cb68d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/user"
)

// MembersRequestBuilder provides operations to manage the members property of the microsoft.graph.group entity.
type MembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MembersRequestBuilderGetQueryParameters direct members of this group, who can be users, devices, other groups, or service principals. Supports the List members, Add member, and Remove member operations. Nullable. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=members($select=id,userPrincipalName,displayName).
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
func (m *MembersRequestBuilder) Application()(*ie29d88f3895d7a38a02ccfe9f35c806f95afb3d0eb9bea6d555f522c74a11049.ApplicationRequestBuilder) {
    return ie29d88f3895d7a38a02ccfe9f35c806f95afb3d0eb9bea6d555f522c74a11049.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersRequestBuilder) {
    m := &MembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/members{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
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
// Count the Count property
func (m *MembersRequestBuilder) Count()(*idc6ec6a1bedd6935f4390a1b2911a367f6435ef357cd42c0c7fa2134416bacb1.CountRequestBuilder) {
    return idc6ec6a1bedd6935f4390a1b2911a367f6435ef357cd42c0c7fa2134416bacb1.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation direct members of this group, who can be users, devices, other groups, or service principals. Supports the List members, Add member, and Remove member operations. Nullable. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=members($select=id,userPrincipalName,displayName).
func (m *MembersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration direct members of this group, who can be users, devices, other groups, or service principals. Supports the List members, Add member, and Remove member operations. Nullable. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=members($select=id,userPrincipalName,displayName).
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
func (m *MembersRequestBuilder) Device()(*i93eb95fc5cc21c1b7ff87b7fc9a1ccaac4340ca6d76b73e5c766740fae74fa58.DeviceRequestBuilder) {
    return i93eb95fc5cc21c1b7ff87b7fc9a1ccaac4340ca6d76b73e5c766740fae74fa58.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get direct members of this group, who can be users, devices, other groups, or service principals. Supports the List members, Add member, and Remove member operations. Nullable. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=members($select=id,userPrincipalName,displayName).
func (m *MembersRequestBuilder) Get(ctx context.Context, requestConfiguration *MembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *MembersRequestBuilder) Group()(*i2715231aa1b97ebc0e299d1c80b5aa685b14ad2cbadc52b883e53096df40d25d.GroupRequestBuilder) {
    return i2715231aa1b97ebc0e299d1c80b5aa685b14ad2cbadc52b883e53096df40d25d.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MembersRequestBuilder) OrgContact()(*i5b707f4d1a3b6bb92ef009b5a12f4aa2f906e2a505066d14fad393f89f2d88b1.OrgContactRequestBuilder) {
    return i5b707f4d1a3b6bb92ef009b5a12f4aa2f906e2a505066d14fad393f89f2d88b1.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the Ref property
func (m *MembersRequestBuilder) Ref()(*ia4ba8805ab0ada495b5c5fbb898da91d5e48daac76a44a52a6a3d8a1abfe0c8e.RefRequestBuilder) {
    return ia4ba8805ab0ada495b5c5fbb898da91d5e48daac76a44a52a6a3d8a1abfe0c8e.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MembersRequestBuilder) ServicePrincipal()(*ice345922f6e79bd547a80a813a84ebd47224ff22ce94579cf0aeee901a020457.ServicePrincipalRequestBuilder) {
    return ice345922f6e79bd547a80a813a84ebd47224ff22ce94579cf0aeee901a020457.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MembersRequestBuilder) User()(*if0c842f2c2a020efe10aa63f93518dcf5ddfb94911ef5e1d22fe197c5a4cb68d.UserRequestBuilder) {
    return if0c842f2c2a020efe10aa63f93518dcf5ddfb94911ef5e1d22fe197c5a4cb68d.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
