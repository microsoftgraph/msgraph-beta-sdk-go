package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
type ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetQueryParameters a list of group members with license errors from this group-based license assignment. Read-only.
type ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal instantiates a new ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    m := &ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder instantiates a new ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a list of group members with license errors from this group-based license assignment. Read-only.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// GraphApplication casts the previous resource to application.
// returns a *ItemMembersWithLicenseErrorsItemGraphApplicationRequestBuilder when successful
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) GraphApplication()(*ItemMembersWithLicenseErrorsItemGraphApplicationRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *ItemMembersWithLicenseErrorsItemGraphDeviceRequestBuilder when successful
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) GraphDevice()(*ItemMembersWithLicenseErrorsItemGraphDeviceRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *ItemMembersWithLicenseErrorsItemGraphGroupRequestBuilder when successful
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) GraphGroup()(*ItemMembersWithLicenseErrorsItemGraphGroupRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *ItemMembersWithLicenseErrorsItemGraphOrgContactRequestBuilder when successful
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) GraphOrgContact()(*ItemMembersWithLicenseErrorsItemGraphOrgContactRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemMembersWithLicenseErrorsItemGraphServicePrincipalRequestBuilder when successful
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemMembersWithLicenseErrorsItemGraphServicePrincipalRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemMembersWithLicenseErrorsItemGraphUserRequestBuilder when successful
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) GraphUser()(*ItemMembersWithLicenseErrorsItemGraphUserRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation a list of group members with license errors from this group-based license assignment. Read-only.
// returns a *RequestInformation when successful
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder when successful
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    return NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
