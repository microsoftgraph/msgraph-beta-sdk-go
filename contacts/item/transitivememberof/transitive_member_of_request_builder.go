package transitivememberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2988e4d0f87e2e0a36221d845642ffebed085ca169a5a52043f3b12c6988a6d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/serviceprincipal"
    i51522ead6b40ea99c1077a3877a31b3d503f0807502d73fd210b02eac0a3ad1b "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/count"
    i58b86c2aaf88b65364ff48cab7c11c23ff3937f1add1c04bd96adba7ca4b3b81 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/orgcontact"
    ic568ac37a6169a61d2597a664dcef2b9e59824d043aaac735eca63c7b51b2d6f "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/group"
    ic96408c1b8ab4178352359aae814c42447a66b772695cc0e4c764a46f103003a "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/application"
    ieab82ae5a4439eb7e187b0793a427d195974bda2f0596d2884be28a67427d61e "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/user"
    ifb75a1d434d2124552cfd900c5015333e1244cb34c248e10deaad102f0d2889e "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/device"
)

// TransitiveMemberOfRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.orgContact entity.
type TransitiveMemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TransitiveMemberOfRequestBuilderGetQueryParameters get transitiveMemberOf from contacts
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
func (m *TransitiveMemberOfRequestBuilder) Application()(*ic96408c1b8ab4178352359aae814c42447a66b772695cc0e4c764a46f103003a.ApplicationRequestBuilder) {
    return ic96408c1b8ab4178352359aae814c42447a66b772695cc0e4c764a46f103003a.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTransitiveMemberOfRequestBuilderInternal instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewTransitiveMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMemberOfRequestBuilder) {
    m := &TransitiveMemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact%2Did}/transitiveMemberOf{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
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
func (m *TransitiveMemberOfRequestBuilder) Count()(*i51522ead6b40ea99c1077a3877a31b3d503f0807502d73fd210b02eac0a3ad1b.CountRequestBuilder) {
    return i51522ead6b40ea99c1077a3877a31b3d503f0807502d73fd210b02eac0a3ad1b.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get transitiveMemberOf from contacts
func (m *TransitiveMemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get transitiveMemberOf from contacts
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
func (m *TransitiveMemberOfRequestBuilder) Device()(*ifb75a1d434d2124552cfd900c5015333e1244cb34c248e10deaad102f0d2889e.DeviceRequestBuilder) {
    return ifb75a1d434d2124552cfd900c5015333e1244cb34c248e10deaad102f0d2889e.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get transitiveMemberOf from contacts
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
func (m *TransitiveMemberOfRequestBuilder) Group()(*ic568ac37a6169a61d2597a664dcef2b9e59824d043aaac735eca63c7b51b2d6f.GroupRequestBuilder) {
    return ic568ac37a6169a61d2597a664dcef2b9e59824d043aaac735eca63c7b51b2d6f.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *TransitiveMemberOfRequestBuilder) OrgContact()(*i58b86c2aaf88b65364ff48cab7c11c23ff3937f1add1c04bd96adba7ca4b3b81.OrgContactRequestBuilder) {
    return i58b86c2aaf88b65364ff48cab7c11c23ff3937f1add1c04bd96adba7ca4b3b81.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *TransitiveMemberOfRequestBuilder) ServicePrincipal()(*i2988e4d0f87e2e0a36221d845642ffebed085ca169a5a52043f3b12c6988a6d0.ServicePrincipalRequestBuilder) {
    return i2988e4d0f87e2e0a36221d845642ffebed085ca169a5a52043f3b12c6988a6d0.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *TransitiveMemberOfRequestBuilder) User()(*ieab82ae5a4439eb7e187b0793a427d195974bda2f0596d2884be28a67427d61e.UserRequestBuilder) {
    return ieab82ae5a4439eb7e187b0793a427d195974bda2f0596d2884be28a67427d61e.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
