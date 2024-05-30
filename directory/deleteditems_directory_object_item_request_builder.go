package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeleteditemsDirectoryObjectItemRequestBuilder provides operations to manage the deletedItems property of the microsoft.graph.directory entity.
type DeleteditemsDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeleteditemsDirectoryObjectItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeleteditemsDirectoryObjectItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeleteditemsDirectoryObjectItemRequestBuilderGetQueryParameters recently deleted items. Read-only. Nullable.
type DeleteditemsDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeleteditemsDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeleteditemsDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeleteditemsDirectoryObjectItemRequestBuilderGetQueryParameters
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
// returns a *DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) CheckMemberGroups()(*DeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilder) {
    return NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
// returns a *DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) CheckMemberObjects()(*DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    return NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeleteditemsDirectoryObjectItemRequestBuilderInternal instantiates a new DeleteditemsDirectoryObjectItemRequestBuilder and sets the default values.
func NewDeleteditemsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsDirectoryObjectItemRequestBuilder) {
    m := &DeleteditemsDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/deletedItems/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeleteditemsDirectoryObjectItemRequestBuilder instantiates a new DeleteditemsDirectoryObjectItemRequestBuilder and sets the default values.
func NewDeleteditemsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeleteditemsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deletedItems for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeleteditemsDirectoryObjectItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get recently deleted items. Read-only. Nullable.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeleteditemsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
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
// returns a *DeleteditemsItemGetmembergroupsGetMemberGroupsRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) GetMemberGroups()(*DeleteditemsItemGetmembergroupsGetMemberGroupsRequestBuilder) {
    return NewDeleteditemsItemGetmembergroupsGetMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMemberObjects provides operations to call the getMemberObjects method.
// returns a *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) GetMemberObjects()(*DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    return NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAdministrativeUnit casts the previous resource to administrativeUnit.
// returns a *DeleteditemsItemGraphadministrativeunitGraphAdministrativeUnitRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) GraphAdministrativeUnit()(*DeleteditemsItemGraphadministrativeunitGraphAdministrativeUnitRequestBuilder) {
    return NewDeleteditemsItemGraphadministrativeunitGraphAdministrativeUnitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphApplication casts the previous resource to application.
// returns a *DeleteditemsItemGraphapplicationGraphApplicationRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) GraphApplication()(*DeleteditemsItemGraphapplicationGraphApplicationRequestBuilder) {
    return NewDeleteditemsItemGraphapplicationGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *DeleteditemsItemGraphdeviceGraphDeviceRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) GraphDevice()(*DeleteditemsItemGraphdeviceGraphDeviceRequestBuilder) {
    return NewDeleteditemsItemGraphdeviceGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *DeleteditemsItemGraphgroupGraphGroupRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) GraphGroup()(*DeleteditemsItemGraphgroupGraphGroupRequestBuilder) {
    return NewDeleteditemsItemGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *DeleteditemsItemGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*DeleteditemsItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewDeleteditemsItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *DeleteditemsItemGraphuserGraphUserRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) GraphUser()(*DeleteditemsItemGraphuserGraphUserRequestBuilder) {
    return NewDeleteditemsItemGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *DeleteditemsItemRestoreRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) Restore()(*DeleteditemsItemRestoreRequestBuilder) {
    return NewDeleteditemsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deletedItems for directory
// returns a *RequestInformation when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeleteditemsDirectoryObjectItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation recently deleted items. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeleteditemsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeleteditemsDirectoryObjectItemRequestBuilder when successful
func (m *DeleteditemsDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*DeleteditemsDirectoryObjectItemRequestBuilder) {
    return NewDeleteditemsDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
