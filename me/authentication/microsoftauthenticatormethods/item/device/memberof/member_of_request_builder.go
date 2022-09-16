package memberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i015344863768ead96fedb9f34b54abf64f681672bd6954de1931345c6ac108b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/count"
    i4f14a0ad70f60b68dd47e0f2807e8b050cf58abe49152522dd31eb447163872c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/device"
    i7bebbafd12bf8ff88e449f6a265ae1e0ba4c3f3caf335cbed8ee96d4d7128e15 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/application"
    ia85baa61e03ef346bd4c50e65bf78cc7bfbd2e6f7b5e53382bcc3ffb394162c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/serviceprincipal"
    ib0d247c9bcb76936e4203cc3ca689e4ba12fdbca8144be5cdc0e789ecf04b936 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/group"
    id1abf58fce4c0de863e4445cf43ec4b51e276cccd78d8d5886fd1b23492517b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/orgcontact"
    if165dbc832b608b829dd54f13d72503280c1acb86c238252fd91cbfaa756ed62 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/user"
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
func (m *MemberOfRequestBuilder) Application()(*i7bebbafd12bf8ff88e449f6a265ae1e0ba4c3f3caf335cbed8ee96d4d7128e15.ApplicationRequestBuilder) {
    return i7bebbafd12bf8ff88e449f6a265ae1e0ba4c3f3caf335cbed8ee96d4d7128e15.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
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
func (m *MemberOfRequestBuilder) Count()(*i015344863768ead96fedb9f34b54abf64f681672bd6954de1931345c6ac108b0.CountRequestBuilder) {
    return i015344863768ead96fedb9f34b54abf64f681672bd6954de1931345c6ac108b0.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MemberOfRequestBuilder) Device()(*i4f14a0ad70f60b68dd47e0f2807e8b050cf58abe49152522dd31eb447163872c.DeviceRequestBuilder) {
    return i4f14a0ad70f60b68dd47e0f2807e8b050cf58abe49152522dd31eb447163872c.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *MemberOfRequestBuilder) Group()(*ib0d247c9bcb76936e4203cc3ca689e4ba12fdbca8144be5cdc0e789ecf04b936.GroupRequestBuilder) {
    return ib0d247c9bcb76936e4203cc3ca689e4ba12fdbca8144be5cdc0e789ecf04b936.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MemberOfRequestBuilder) OrgContact()(*id1abf58fce4c0de863e4445cf43ec4b51e276cccd78d8d5886fd1b23492517b8.OrgContactRequestBuilder) {
    return id1abf58fce4c0de863e4445cf43ec4b51e276cccd78d8d5886fd1b23492517b8.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MemberOfRequestBuilder) ServicePrincipal()(*ia85baa61e03ef346bd4c50e65bf78cc7bfbd2e6f7b5e53382bcc3ffb394162c1.ServicePrincipalRequestBuilder) {
    return ia85baa61e03ef346bd4c50e65bf78cc7bfbd2e6f7b5e53382bcc3ffb394162c1.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MemberOfRequestBuilder) User()(*if165dbc832b608b829dd54f13d72503280c1acb86c238252fd91cbfaa756ed62.UserRequestBuilder) {
    return if165dbc832b608b829dd54f13d72503280c1acb86c238252fd91cbfaa756ed62.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
