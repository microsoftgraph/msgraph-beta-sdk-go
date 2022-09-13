package transitivememberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04b6f6832e6d7041167eb8a90d64106050e903aa2fd84e3bc266ad1a8105f2db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/count"
    i189d5038e54044dcfd3233a6a6e519d4176047555564b18491ba19bb619bae15 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/serviceprincipal"
    i59ef7acea5cb7c08abbc9ef52f9c62b36af5a7e1c3279f161f38f3d2046f84c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/group"
    i5ba8d21eaa9a4f66dc87f668e7ee8c7aaed0ace84d9237630c4cd93396911ec2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/device"
    i97886a363459ee2a387f85fa71c9f3185d01ad80770edd3f46342259ecefb7e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/user"
    ib73e8740bc537f6b05d0301b1acecdc5b199fe8e2b8eaa8c10244f7a90bef263 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/orgcontact"
    ice415fb13b3f9c5b2379d32f1c00b04efcb8c15a9608c336c577b374492bb377 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/application"
)

// TransitiveMemberOfRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
type TransitiveMemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TransitiveMemberOfRequestBuilderGetQueryParameters the groups, including nested groups, and directory roles that a user is a member of. Nullable.
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
func (m *TransitiveMemberOfRequestBuilder) Application()(*ice415fb13b3f9c5b2379d32f1c00b04efcb8c15a9608c336c577b374492bb377.ApplicationRequestBuilder) {
    return ice415fb13b3f9c5b2379d32f1c00b04efcb8c15a9608c336c577b374492bb377.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTransitiveMemberOfRequestBuilderInternal instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewTransitiveMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMemberOfRequestBuilder) {
    m := &TransitiveMemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/transitiveMemberOf{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
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
// Count the Count property
func (m *TransitiveMemberOfRequestBuilder) Count()(*i04b6f6832e6d7041167eb8a90d64106050e903aa2fd84e3bc266ad1a8105f2db.CountRequestBuilder) {
    return i04b6f6832e6d7041167eb8a90d64106050e903aa2fd84e3bc266ad1a8105f2db.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the groups, including nested groups, and directory roles that a user is a member of. Nullable.
func (m *TransitiveMemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the groups, including nested groups, and directory roles that a user is a member of. Nullable.
func (m *TransitiveMemberOfRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TransitiveMemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TransitiveMemberOfRequestBuilder) Device()(*i5ba8d21eaa9a4f66dc87f668e7ee8c7aaed0ace84d9237630c4cd93396911ec2.DeviceRequestBuilder) {
    return i5ba8d21eaa9a4f66dc87f668e7ee8c7aaed0ace84d9237630c4cd93396911ec2.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the groups, including nested groups, and directory roles that a user is a member of. Nullable.
func (m *TransitiveMemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *TransitiveMemberOfRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *TransitiveMemberOfRequestBuilder) Group()(*i59ef7acea5cb7c08abbc9ef52f9c62b36af5a7e1c3279f161f38f3d2046f84c1.GroupRequestBuilder) {
    return i59ef7acea5cb7c08abbc9ef52f9c62b36af5a7e1c3279f161f38f3d2046f84c1.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *TransitiveMemberOfRequestBuilder) OrgContact()(*ib73e8740bc537f6b05d0301b1acecdc5b199fe8e2b8eaa8c10244f7a90bef263.OrgContactRequestBuilder) {
    return ib73e8740bc537f6b05d0301b1acecdc5b199fe8e2b8eaa8c10244f7a90bef263.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *TransitiveMemberOfRequestBuilder) ServicePrincipal()(*i189d5038e54044dcfd3233a6a6e519d4176047555564b18491ba19bb619bae15.ServicePrincipalRequestBuilder) {
    return i189d5038e54044dcfd3233a6a6e519d4176047555564b18491ba19bb619bae15.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *TransitiveMemberOfRequestBuilder) User()(*i97886a363459ee2a387f85fa71c9f3185d01ad80770edd3f46342259ecefb7e7.UserRequestBuilder) {
    return i97886a363459ee2a387f85fa71c9f3185d01ad80770edd3f46342259ecefb7e7.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
