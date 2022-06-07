package memberof

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i083a719c480c798f02abc8cf77c6c4aeb8f1424ed5234ab3cb891a0968409fcd "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/device"
    i2209aa480b92340afe33b32c142fc00bf6e92cb8396441a1d563040e13c191f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/group"
    i22298fe19f96482d122feedfa4c5c9391b3df82150d10361d002839ed9cc8c6d "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/serviceprincipal"
    i2a7736943402ac3a4fdfd8f575bc33537164de1032b59f7efb7944a81c5a66c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/orgcontact"
    i9efe8513b832969f3314978cdbe5e794f3b18e7433f6dce969fb945d5f9be171 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/user"
    ie0751a2cf81dc015535b102dd95ea602954133134d8c06b43e3b08363b67eb09 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/application"
    ie8e6b6346baf9a1a0d43144636d549977a2e6235a015be22bf37a6bb38872e6a "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/count"
)

// MemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.orgContact entity.
type MemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MemberOfRequestBuilderGetQueryParameters groups that this contact is a member of. Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Application()(*ie0751a2cf81dc015535b102dd95ea602954133134d8c06b43e3b08363b67eb09.ApplicationRequestBuilder) {
    return ie0751a2cf81dc015535b102dd95ea602954133134d8c06b43e3b08363b67eb09.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact%2Did}/memberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
func (m *MemberOfRequestBuilder) Count()(*ie8e6b6346baf9a1a0d43144636d549977a2e6235a015be22bf37a6bb38872e6a.CountRequestBuilder) {
    return ie8e6b6346baf9a1a0d43144636d549977a2e6235a015be22bf37a6bb38872e6a.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation groups that this contact is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration groups that this contact is a member of. Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Device()(*i083a719c480c798f02abc8cf77c6c4aeb8f1424ed5234ab3cb891a0968409fcd.DeviceRequestBuilder) {
    return i083a719c480c798f02abc8cf77c6c4aeb8f1424ed5234ab3cb891a0968409fcd.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups that this contact is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler groups that this contact is a member of. Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Group()(*i2209aa480b92340afe33b32c142fc00bf6e92cb8396441a1d563040e13c191f7.GroupRequestBuilder) {
    return i2209aa480b92340afe33b32c142fc00bf6e92cb8396441a1d563040e13c191f7.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MemberOfRequestBuilder) OrgContact()(*i2a7736943402ac3a4fdfd8f575bc33537164de1032b59f7efb7944a81c5a66c0.OrgContactRequestBuilder) {
    return i2a7736943402ac3a4fdfd8f575bc33537164de1032b59f7efb7944a81c5a66c0.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MemberOfRequestBuilder) ServicePrincipal()(*i22298fe19f96482d122feedfa4c5c9391b3df82150d10361d002839ed9cc8c6d.ServicePrincipalRequestBuilder) {
    return i22298fe19f96482d122feedfa4c5c9391b3df82150d10361d002839ed9cc8c6d.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MemberOfRequestBuilder) User()(*i9efe8513b832969f3314978cdbe5e794f3b18e7433f6dce969fb945d5f9be171.UserRequestBuilder) {
    return i9efe8513b832969f3314978cdbe5e794f3b18e7433f6dce969fb945d5f9be171.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
