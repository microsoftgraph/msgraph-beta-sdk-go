package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedItemsDirectoryObjectItemRequestBuilder provides operations to manage the deletedItems property of the microsoft.graph.directory entity.
type DeletedItemsDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedItemsDirectoryObjectItemRequestBuilderGetQueryParameters recently deleted items. Read-only. Nullable.
type DeletedItemsDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeletedItemsDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedItemsDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedItemsDirectoryObjectItemRequestBuilderGetQueryParameters
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
// returns a *DeletedItemsItemCheckMemberGroupsRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) CheckMemberGroups()(*DeletedItemsItemCheckMemberGroupsRequestBuilder) {
    return NewDeletedItemsItemCheckMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
// returns a *DeletedItemsItemCheckMemberObjectsRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) CheckMemberObjects()(*DeletedItemsItemCheckMemberObjectsRequestBuilder) {
    return NewDeletedItemsItemCheckMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeletedItemsDirectoryObjectItemRequestBuilderInternal instantiates a new DeletedItemsDirectoryObjectItemRequestBuilder and sets the default values.
func NewDeletedItemsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedItemsDirectoryObjectItemRequestBuilder) {
    m := &DeletedItemsDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/deletedItems/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeletedItemsDirectoryObjectItemRequestBuilder instantiates a new DeletedItemsDirectoryObjectItemRequestBuilder and sets the default values.
func NewDeletedItemsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedItemsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedItemsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get recently deleted items. Read-only. Nullable.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedItemsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
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
// GetMemberGroups provides operations to call the getMemberGroups method.
// returns a *DeletedItemsItemGetMemberGroupsRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) GetMemberGroups()(*DeletedItemsItemGetMemberGroupsRequestBuilder) {
    return NewDeletedItemsItemGetMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMemberObjects provides operations to call the getMemberObjects method.
// returns a *DeletedItemsItemGetMemberObjectsRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) GetMemberObjects()(*DeletedItemsItemGetMemberObjectsRequestBuilder) {
    return NewDeletedItemsItemGetMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAdministrativeUnit casts the previous resource to administrativeUnit.
// returns a *DeletedItemsItemGraphAdministrativeUnitRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) GraphAdministrativeUnit()(*DeletedItemsItemGraphAdministrativeUnitRequestBuilder) {
    return NewDeletedItemsItemGraphAdministrativeUnitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphApplication casts the previous resource to application.
// returns a *DeletedItemsItemGraphApplicationRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) GraphApplication()(*DeletedItemsItemGraphApplicationRequestBuilder) {
    return NewDeletedItemsItemGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *DeletedItemsItemGraphDeviceRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) GraphDevice()(*DeletedItemsItemGraphDeviceRequestBuilder) {
    return NewDeletedItemsItemGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *DeletedItemsItemGraphGroupRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) GraphGroup()(*DeletedItemsItemGraphGroupRequestBuilder) {
    return NewDeletedItemsItemGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *DeletedItemsItemGraphServicePrincipalRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*DeletedItemsItemGraphServicePrincipalRequestBuilder) {
    return NewDeletedItemsItemGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *DeletedItemsItemGraphUserRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) GraphUser()(*DeletedItemsItemGraphUserRequestBuilder) {
    return NewDeletedItemsItemGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *DeletedItemsItemRestoreRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) Restore()(*DeletedItemsItemRestoreRequestBuilder) {
    return NewDeletedItemsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation recently deleted items. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedItemsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedItemsDirectoryObjectItemRequestBuilder when successful
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*DeletedItemsDirectoryObjectItemRequestBuilder) {
    return NewDeletedItemsDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
