package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i21357f81dd67d1f048d315ac01f454d176950dd54ea7804555e588cd3eff8c41 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/termstore"
    i30365f7545fb0bfa84213c73d2831a20632673b88b79579fecb0fbdf47714c92 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/getbypathwithpath"
    i39ed30b8a944c4df2cb1c7d7719a6e16eed74f779bf8c7f46e7ec330dc32490d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/drives"
    i48525d077819c844d3854d1d3ed447d8615dae48edb552ae208e6dadac2913df "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes"
    i48ea2a845dc13dbdbaaf728c5005ec957407a596f9bdb13c97ff6a9c23bb80c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/permissions"
    i5e304ff4e8f7389e05b8435725a95166eb67d5d1e116cc3b18d200fbd7bdaa7e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/analytics"
    i70ffe9e61c8e968ae0dc889d6f9ca1d946f32f80c595c1b4cfc51668c4ae18f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/drive"
    i71869306bbdbc613872f0a8f42f33a73181332142694f31c6ca16a38d27b5022 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/columns"
    iaff8ef3298b8a50022c41f0588c95ecabe270d388a402d0debfb273bb76823d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote"
    ib3a7728f14558a828657cc0b2022eb065932c8bdf52b373d5242cfcb39017929 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/getapplicablecontenttypesforlistwithlistid"
    ibc4bf0f224e96734ea6c3e1ccc8784d21b398ae8904db6f204237b6482921a0c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/operations"
    ic74a7f8d1e57a8da50add8729c911139b5ce66432f0fd2005e45756231bf7e06 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists"
    ida491665689c206cec14e94c4ba2d1a53e5513fecd63f74d6f168640b02f0598 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/pages"
    ie4755f3ab2b3be7e0faaa8f804ac25658f750fa0888f86c64b1acb46498e6d15 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/sites"
    ie550bbbe668a53a0aac5fe9ee9f3ac9fc9d9d3c06204956f7c3d7dd3a17445a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/externalcolumns"
    ief47e71fa531377bc96a38300a71ed5155d7aa8f72f080267a5c03dc19037638 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ifb02c18e8f70a2a669af59896c7c40e851e9d0ff30f6df756f8d8eeafa69a2a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/items"
    i053597ae2fca36945f3da0e944bf07318cb94bde79534f0ef588eac99861007f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/externalcolumns/item"
    i161351685aba3b02962a6a41ccbc3125edc0df60596b473175fadb0ca4ee50c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/pages/item"
    i2ad58809366fe51c914338bc998230f2c52e936efb52e3da9dc1795da6d2e60e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item"
    i727152b614c462054b3f5cbde245c93660e74322a2e4614e0bf0455749ffe99d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/items/item"
    i73e8a92b70513f051da5d1a665c9c5c97d07dcfc0bc14436b3e9ddae0a733601 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item"
    i7dbcdb87846307ec83020a5f7bc018b8694724bb7641788349fd036a81ce8c46 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/operations/item"
    i8d7b2e54750fab899bdfdc413ec9096c45dd35a331acfc965e0139c931df339e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/permissions/item"
    ic4ea50c9f752d2b4c86c26f61c9fe954fa57cf84715361c40caf7028754c21d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/drives/item"
    id9138e79961f7d02552ce3b5941d270ed2f0a8a46497542d2459bd97dccbc06b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/sites/item"
    if0253c120b3b3a099cb60e9326d7b7be9d1749a2813c7fa0cae43bf0e67deb80 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/columns/item"
)

// SiteItemRequestBuilder provides operations to manage the sites property of the microsoft.graph.group entity.
type SiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SiteItemRequestBuilderGetQueryParameters the list of SharePoint sites in this group. Access the default site with /sites/root.
type SiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SiteItemRequestBuilderGetQueryParameters
}
// SiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics the analytics property
func (m *SiteItemRequestBuilder) Analytics()(*i5e304ff4e8f7389e05b8435725a95166eb67d5d1e116cc3b18d200fbd7bdaa7e.AnalyticsRequestBuilder) {
    return i5e304ff4e8f7389e05b8435725a95166eb67d5d1e116cc3b18d200fbd7bdaa7e.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *SiteItemRequestBuilder) Columns()(*i71869306bbdbc613872f0a8f42f33a73181332142694f31c6ca16a38d27b5022.ColumnsRequestBuilder) {
    return i71869306bbdbc613872f0a8f42f33a73181332142694f31c6ca16a38d27b5022.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.columns.item collection
func (m *SiteItemRequestBuilder) ColumnsById(id string)(*if0253c120b3b3a099cb60e9326d7b7be9d1749a2813c7fa0cae43bf0e67deb80.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return if0253c120b3b3a099cb60e9326d7b7be9d1749a2813c7fa0cae43bf0e67deb80.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SiteItemRequestBuilder) {
    m := &SiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes the contentTypes property
func (m *SiteItemRequestBuilder) ContentTypes()(*i48525d077819c844d3854d1d3ed447d8615dae48edb552ae208e6dadac2913df.ContentTypesRequestBuilder) {
    return i48525d077819c844d3854d1d3ed447d8615dae48edb552ae208e6dadac2913df.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.contentTypes.item collection
func (m *SiteItemRequestBuilder) ContentTypesById(id string)(*i73e8a92b70513f051da5d1a665c9c5c97d07dcfc0bc14436b3e9ddae0a733601.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return i73e8a92b70513f051da5d1a665c9c5c97d07dcfc0bc14436b3e9ddae0a733601.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SiteItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sites in groups
func (m *SiteItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, requestConfiguration *SiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Drive the drive property
func (m *SiteItemRequestBuilder) Drive()(*i70ffe9e61c8e968ae0dc889d6f9ca1d946f32f80c595c1b4cfc51668c4ae18f3.DriveRequestBuilder) {
    return i70ffe9e61c8e968ae0dc889d6f9ca1d946f32f80c595c1b4cfc51668c4ae18f3.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives the drives property
func (m *SiteItemRequestBuilder) Drives()(*i39ed30b8a944c4df2cb1c7d7719a6e16eed74f779bf8c7f46e7ec330dc32490d.DrivesRequestBuilder) {
    return i39ed30b8a944c4df2cb1c7d7719a6e16eed74f779bf8c7f46e7ec330dc32490d.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.drives.item collection
func (m *SiteItemRequestBuilder) DrivesById(id string)(*ic4ea50c9f752d2b4c86c26f61c9fe954fa57cf84715361c40caf7028754c21d7.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return ic4ea50c9f752d2b4c86c26f61c9fe954fa57cf84715361c40caf7028754c21d7.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalColumns the externalColumns property
func (m *SiteItemRequestBuilder) ExternalColumns()(*ie550bbbe668a53a0aac5fe9ee9f3ac9fc9d9d3c06204956f7c3d7dd3a17445a8.ExternalColumnsRequestBuilder) {
    return ie550bbbe668a53a0aac5fe9ee9f3ac9fc9d9d3c06204956f7c3d7dd3a17445a8.NewExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.externalColumns.item collection
func (m *SiteItemRequestBuilder) ExternalColumnsById(id string)(*i053597ae2fca36945f3da0e944bf07318cb94bde79534f0ef588eac99861007f.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i053597ae2fca36945f3da0e944bf07318cb94bde79534f0ef588eac99861007f.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SiteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *SiteItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ief47e71fa531377bc96a38300a71ed5155d7aa8f72f080267a5c03dc19037638.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ief47e71fa531377bc96a38300a71ed5155d7aa8f72f080267a5c03dc19037638.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
func (m *SiteItemRequestBuilder) GetApplicableContentTypesForListWithListId(listId *string)(*ib3a7728f14558a828657cc0b2022eb065932c8bdf52b373d5242cfcb39017929.GetApplicableContentTypesForListWithListIdRequestBuilder) {
    return ib3a7728f14558a828657cc0b2022eb065932c8bdf52b373d5242cfcb39017929.NewGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, listId);
}
// GetByPathWithPath provides operations to call the getByPath method.
func (m *SiteItemRequestBuilder) GetByPathWithPath(path *string)(*i30365f7545fb0bfa84213c73d2831a20632673b88b79579fecb0fbdf47714c92.GetByPathWithPathRequestBuilder) {
    return i30365f7545fb0bfa84213c73d2831a20632673b88b79579fecb0fbdf47714c92.NewGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path);
}
// Items the items property
func (m *SiteItemRequestBuilder) Items()(*ifb02c18e8f70a2a669af59896c7c40e851e9d0ff30f6df756f8d8eeafa69a2a7.ItemsRequestBuilder) {
    return ifb02c18e8f70a2a669af59896c7c40e851e9d0ff30f6df756f8d8eeafa69a2a7.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.items.item collection
func (m *SiteItemRequestBuilder) ItemsById(id string)(*i727152b614c462054b3f5cbde245c93660e74322a2e4614e0bf0455749ffe99d.BaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem%2Did"] = id
    }
    return i727152b614c462054b3f5cbde245c93660e74322a2e4614e0bf0455749ffe99d.NewBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Lists the lists property
func (m *SiteItemRequestBuilder) Lists()(*ic74a7f8d1e57a8da50add8729c911139b5ce66432f0fd2005e45756231bf7e06.ListsRequestBuilder) {
    return ic74a7f8d1e57a8da50add8729c911139b5ce66432f0fd2005e45756231bf7e06.NewListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item collection
func (m *SiteItemRequestBuilder) ListsById(id string)(*i2ad58809366fe51c914338bc998230f2c52e936efb52e3da9dc1795da6d2e60e.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list%2Did"] = id
    }
    return i2ad58809366fe51c914338bc998230f2c52e936efb52e3da9dc1795da6d2e60e.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote the onenote property
func (m *SiteItemRequestBuilder) Onenote()(*iaff8ef3298b8a50022c41f0588c95ecabe270d388a402d0debfb273bb76823d1.OnenoteRequestBuilder) {
    return iaff8ef3298b8a50022c41f0588c95ecabe270d388a402d0debfb273bb76823d1.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations the operations property
func (m *SiteItemRequestBuilder) Operations()(*ibc4bf0f224e96734ea6c3e1ccc8784d21b398ae8904db6f204237b6482921a0c.OperationsRequestBuilder) {
    return ibc4bf0f224e96734ea6c3e1ccc8784d21b398ae8904db6f204237b6482921a0c.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.operations.item collection
func (m *SiteItemRequestBuilder) OperationsById(id string)(*i7dbcdb87846307ec83020a5f7bc018b8694724bb7641788349fd036a81ce8c46.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return i7dbcdb87846307ec83020a5f7bc018b8694724bb7641788349fd036a81ce8c46.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *SiteItemRequestBuilder) Pages()(*ida491665689c206cec14e94c4ba2d1a53e5513fecd63f74d6f168640b02f0598.PagesRequestBuilder) {
    return ida491665689c206cec14e94c4ba2d1a53e5513fecd63f74d6f168640b02f0598.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.pages.item collection
func (m *SiteItemRequestBuilder) PagesById(id string)(*i161351685aba3b02962a6a41ccbc3125edc0df60596b473175fadb0ca4ee50c2.SitePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sitePage%2Did"] = id
    }
    return i161351685aba3b02962a6a41ccbc3125edc0df60596b473175fadb0ca4ee50c2.NewSitePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property sites in groups
func (m *SiteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, requestConfiguration *SiteItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable), nil
}
// Permissions the permissions property
func (m *SiteItemRequestBuilder) Permissions()(*i48ea2a845dc13dbdbaaf728c5005ec957407a596f9bdb13c97ff6a9c23bb80c6.PermissionsRequestBuilder) {
    return i48ea2a845dc13dbdbaaf728c5005ec957407a596f9bdb13c97ff6a9c23bb80c6.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.permissions.item collection
func (m *SiteItemRequestBuilder) PermissionsById(id string)(*i8d7b2e54750fab899bdfdc413ec9096c45dd35a331acfc965e0139c931df339e.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i8d7b2e54750fab899bdfdc413ec9096c45dd35a331acfc965e0139c931df339e.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites the sites property
func (m *SiteItemRequestBuilder) Sites()(*ie4755f3ab2b3be7e0faaa8f804ac25658f750fa0888f86c64b1acb46498e6d15.SitesRequestBuilder) {
    return ie4755f3ab2b3be7e0faaa8f804ac25658f750fa0888f86c64b1acb46498e6d15.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.sites.item collection
func (m *SiteItemRequestBuilder) SitesById(id string)(*id9138e79961f7d02552ce3b5941d270ed2f0a8a46497542d2459bd97dccbc06b.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did1"] = id
    }
    return id9138e79961f7d02552ce3b5941d270ed2f0a8a46497542d2459bd97dccbc06b.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TermStore the termStore property
func (m *SiteItemRequestBuilder) TermStore()(*i21357f81dd67d1f048d315ac01f454d176950dd54ea7804555e588cd3eff8c41.TermStoreRequestBuilder) {
    return i21357f81dd67d1f048d315ac01f454d176950dd54ea7804555e588cd3eff8c41.NewTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
