package memberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1f3fce65bf3f10075bfd7eb471be758142d4f91beb7785c02d548360d7ef0754 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/group"
    i2e2bb514dc5ae4b0425ef48e92500fb0f485b6378a51394aac8e373ccaa4970d "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/application"
    i6e13efdd518294df9af61c0b2451ddf1901424f19371b9195123ab7c5e7d5719 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/user"
    ia4b54910583f1060fb151df21e825527b6f2cd44b75a69854e112c6721fadeca "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/device"
    ia7a0a7e27f54c063023b04320604dfa33d4ac71e1fba41f81315235a1f7c2d53 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/count"
    id3df3c4109ba8455f360389f8c9770d79e5f3e4ed3dd9c5ab19264294ca5a829 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/serviceprincipal"
    ied015f8148bf4209c6f8b8afbd8aea13ef3ac3b790aea89515ec9e6434512a38 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/orgcontact"
)

// MemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.servicePrincipal entity.
type MemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MemberOfRequestBuilderGetQueryParameters roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Application()(*i2e2bb514dc5ae4b0425ef48e92500fb0f485b6378a51394aac8e373ccaa4970d.ApplicationRequestBuilder) {
    return i2e2bb514dc5ae4b0425ef48e92500fb0f485b6378a51394aac8e373ccaa4970d.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/memberOf{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
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
func (m *MemberOfRequestBuilder) Count()(*ia7a0a7e27f54c063023b04320604dfa33d4ac71e1fba41f81315235a1f7c2d53.CountRequestBuilder) {
    return ia7a0a7e27f54c063023b04320604dfa33d4ac71e1fba41f81315235a1f7c2d53.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Device()(*ia4b54910583f1060fb151df21e825527b6f2cd44b75a69854e112c6721fadeca.DeviceRequestBuilder) {
    return ia4b54910583f1060fb151df21e825527b6f2cd44b75a69854e112c6721fadeca.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Group()(*i1f3fce65bf3f10075bfd7eb471be758142d4f91beb7785c02d548360d7ef0754.GroupRequestBuilder) {
    return i1f3fce65bf3f10075bfd7eb471be758142d4f91beb7785c02d548360d7ef0754.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MemberOfRequestBuilder) OrgContact()(*ied015f8148bf4209c6f8b8afbd8aea13ef3ac3b790aea89515ec9e6434512a38.OrgContactRequestBuilder) {
    return ied015f8148bf4209c6f8b8afbd8aea13ef3ac3b790aea89515ec9e6434512a38.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MemberOfRequestBuilder) ServicePrincipal()(*id3df3c4109ba8455f360389f8c9770d79e5f3e4ed3dd9c5ab19264294ca5a829.ServicePrincipalRequestBuilder) {
    return id3df3c4109ba8455f360389f8c9770d79e5f3e4ed3dd9c5ab19264294ca5a829.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MemberOfRequestBuilder) User()(*i6e13efdd518294df9af61c0b2451ddf1901424f19371b9195123ab7c5e7d5719.UserRequestBuilder) {
    return i6e13efdd518294df9af61c0b2451ddf1901424f19371b9195123ab7c5e7d5719.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
