package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesSiteItemRequestBuilder provides operations to manage the sites property of the microsoft.graph.group entity.
type ItemSitesSiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesSiteItemRequestBuilderGetQueryParameters the list of SharePoint sites in this group. Access the default site with /sites/root.
type ItemSitesSiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesSiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesSiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesSiteItemRequestBuilderGetQueryParameters
}
// ItemSitesSiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesSiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Analytics()(*ItemSitesItemAnalyticsRequestBuilder) {
    return NewItemSitesItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Columns()(*ItemSitesItemColumnsRequestBuilder) {
    return NewItemSitesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ColumnsById(id string)(*ItemSitesItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// NewItemSitesSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewItemSitesSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, siteId *string)(*ItemSitesSiteItemRequestBuilder) {
    m := &ItemSitesSiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if siteId != nil {
        urlTplParams["site%2Did"] = *siteId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewItemSitesSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesSiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesSiteItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ContentTypes()(*ItemSitesItemContentTypesRequestBuilder) {
    return NewItemSitesItemContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContentTypesById provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ContentTypesById(id string)(*ItemSitesItemContentTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemContentTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Drive provides operations to manage the drive property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Drive()(*ItemSitesItemDriveRequestBuilder) {
    return NewItemSitesItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Drives()(*ItemSitesItemDrivesRequestBuilder) {
    return NewItemSitesItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) DrivesById(id string)(*ItemSitesItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ExternalColumns()(*ItemSitesItemExternalColumnsRequestBuilder) {
    return NewItemSitesItemExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExternalColumnsById provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ExternalColumnsById(id string)(*ItemSitesItemExternalColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemExternalColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Get the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *ItemSitesSiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesSiteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable), nil
}
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) InformationProtection()(*ItemSitesItemInformationProtectionRequestBuilder) {
    return NewItemSitesItemInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Items provides operations to manage the items property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Items()(*ItemSitesItemItemsRequestBuilder) {
    return NewItemSitesItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ItemsById provides operations to manage the items property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ItemsById(id string)(*ItemSitesItemItemsBaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemItemsBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Lists()(*ItemSitesItemListsRequestBuilder) {
    return NewItemSitesItemListsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ListsById provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ListsById(id string)(*ItemSitesItemListsListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemListsListItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemSitesSiteItemRequestBuilder) MicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemSitesItemMicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemSitesItemMicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime)
}
// MicrosoftGraphGetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
func (m *ItemSitesSiteItemRequestBuilder) MicrosoftGraphGetApplicableContentTypesForListWithListId(listId *string)(*ItemSitesItemMicrosoftGraphGetApplicableContentTypesForListWithListIdGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewItemSitesItemMicrosoftGraphGetApplicableContentTypesForListWithListIdGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, listId)
}
// MicrosoftGraphGetByPathWithPath provides operations to call the getByPath method.
func (m *ItemSitesSiteItemRequestBuilder) MicrosoftGraphGetByPathWithPath(path *string)(*ItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilder) {
    return NewItemSitesItemMicrosoftGraphGetByPathWithPathGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Onenote()(*ItemSitesItemOnenoteRequestBuilder) {
    return NewItemSitesItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Operations()(*ItemSitesItemOperationsRequestBuilder) {
    return NewItemSitesItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) OperationsById(id string)(*ItemSitesItemOperationsRichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemOperationsRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Pages provides operations to manage the pages property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Pages()(*ItemSitesItemPagesRequestBuilder) {
    return NewItemSitesItemPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PagesById provides operations to manage the pages property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) PagesById(id string)(*ItemSitesItemPagesSitePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemPagesSitePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Patch update the navigation property sites in groups
func (m *ItemSitesSiteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, requestConfiguration *ItemSitesSiteItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Permissions()(*ItemSitesItemPermissionsRequestBuilder) {
    return NewItemSitesItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) PermissionsById(id string)(*ItemSitesItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Sites()(*ItemSitesItemSitesRequestBuilder) {
    return NewItemSitesItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SitesById provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) SitesById(id string)(*ItemSitesItemSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) TermStore()(*ItemSitesItemTermStoreRequestBuilder) {
    return NewItemSitesItemTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *ItemSitesSiteItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesSiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sites in groups
func (m *ItemSitesSiteItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, requestConfiguration *ItemSitesSiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
