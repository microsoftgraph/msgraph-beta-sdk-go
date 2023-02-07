package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicesItemMemberOfDirectoryObjectItemRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type DevicesItemMemberOfDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DevicesItemMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type DevicesItemMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicesItemMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicesItemMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicesItemMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewDevicesItemMemberOfDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDevicesItemMemberOfDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemMemberOfDirectoryObjectItemRequestBuilder) {
    m := &DevicesItemMemberOfDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/memberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDevicesItemMemberOfDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDevicesItemMemberOfDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicesItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *DevicesItemMemberOfDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicesItemMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// MicrosoftGraphApplication casts the previous resource to application.
func (m *DevicesItemMemberOfDirectoryObjectItemRequestBuilder) MicrosoftGraphApplication()(*DevicesItemMemberOfItemMicrosoftGraphApplicationRequestBuilder) {
    return NewDevicesItemMemberOfItemMicrosoftGraphApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDevice casts the previous resource to device.
func (m *DevicesItemMemberOfDirectoryObjectItemRequestBuilder) MicrosoftGraphDevice()(*DevicesItemMemberOfItemMicrosoftGraphDeviceRequestBuilder) {
    return NewDevicesItemMemberOfItemMicrosoftGraphDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGroup casts the previous resource to group.
func (m *DevicesItemMemberOfDirectoryObjectItemRequestBuilder) MicrosoftGraphGroup()(*DevicesItemMemberOfItemMicrosoftGraphGroupRequestBuilder) {
    return NewDevicesItemMemberOfItemMicrosoftGraphGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOrgContact casts the previous resource to orgContact.
func (m *DevicesItemMemberOfDirectoryObjectItemRequestBuilder) MicrosoftGraphOrgContact()(*DevicesItemMemberOfItemMicrosoftGraphOrgContactRequestBuilder) {
    return NewDevicesItemMemberOfItemMicrosoftGraphOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *DevicesItemMemberOfDirectoryObjectItemRequestBuilder) MicrosoftGraphServicePrincipal()(*DevicesItemMemberOfItemMicrosoftGraphServicePrincipalRequestBuilder) {
    return NewDevicesItemMemberOfItemMicrosoftGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *DevicesItemMemberOfDirectoryObjectItemRequestBuilder) MicrosoftGraphUser()(*DevicesItemMemberOfItemMicrosoftGraphUserRequestBuilder) {
    return NewDevicesItemMemberOfItemMicrosoftGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *DevicesItemMemberOfDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicesItemMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
