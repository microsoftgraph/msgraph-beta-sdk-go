package list

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0f8ca929323429752e21f9acb28f80889948ea9ff72848249fca8c74c368589e "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities"
    i1c4ca68087caab22a544a9dc3fb1c29f6076a9264ac284502cbf04a28444aea7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/subscriptions"
    i2ef5930655c892d16078905ec81246034d704193e69907b59b2f14feab329b60 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes"
    id28cbea101cb9521a14cf9640c45800a00477bd62395913072f4e81bf1074012 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/drive"
    idf9a3059d40c90c42068fa8bf9ac9fd67dedf1182980c80d7a384967f9dff8d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items"
    ie43f83ef262031fdb1eb8da2814c4165f5f945383e15adf8746e6a96f3432e23 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/operations"
    if9ff29afb76bbf1ffef5a126d0ba635e14dcd31de06309e9c7c1a69ad2600738 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/columns"
    i327c97fc6449280089d733ba8dee83d9d74247819100cf6629d3355abfea1166 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/operations/item"
    i37771808e5dcfef82cc1e8bc33ad871d583b797abbb1222c5dc6198791d77af2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/subscriptions/item"
    i58ec931c55ab0d0111918ef65feb8b5a53f2f824a19e864696540521faf828f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item"
    i654b96f70860efab65ce04591400b15de588e186e6e7217fc1ca04b9de80627c "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/columns/item"
    i6dd3632bd204e05cac0aa8e9c03666774381075a56cb272b02bc729017380185 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item"
)

// ListRequestBuilder provides operations to manage the list property of the microsoft.graph.drive entity.
type ListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ListRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ListRequestBuilderGetQueryParameters for drives in SharePoint, the underlying document library list. Read-only. Nullable.
type ListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ListRequestBuilderGetQueryParameters
}
// ListRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *ListRequestBuilder) Activities()(*i0f8ca929323429752e21f9acb28f80889948ea9ff72848249fca8c74c368589e.ActivitiesRequestBuilder) {
    return i0f8ca929323429752e21f9acb28f80889948ea9ff72848249fca8c74c368589e.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *ListRequestBuilder) Columns()(*if9ff29afb76bbf1ffef5a126d0ba635e14dcd31de06309e9c7c1a69ad2600738.ColumnsRequestBuilder) {
    return if9ff29afb76bbf1ffef5a126d0ba635e14dcd31de06309e9c7c1a69ad2600738.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*i654b96f70860efab65ce04591400b15de588e186e6e7217fc1ca04b9de80627c.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i654b96f70860efab65ce04591400b15de588e186e6e7217fc1ca04b9de80627c.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListRequestBuilder instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes the contentTypes property
func (m *ListRequestBuilder) ContentTypes()(*i2ef5930655c892d16078905ec81246034d704193e69907b59b2f14feab329b60.ContentTypesRequestBuilder) {
    return i2ef5930655c892d16078905ec81246034d704193e69907b59b2f14feab329b60.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i6dd3632bd204e05cac0aa8e9c03666774381075a56cb272b02bc729017380185.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return i6dd3632bd204e05cac0aa8e9c03666774381075a56cb272b02bc729017380185.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for drive
func (m *ListRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property list for drive
func (m *ListRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ListRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property list in drive
func (m *ListRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property list in drive
func (m *ListRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property list for drive
func (m *ListRequestBuilder) Delete(ctx context.Context, requestConfiguration *ListRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Drive the drive property
func (m *ListRequestBuilder) Drive()(*id28cbea101cb9521a14cf9640c45800a00477bd62395913072f4e81bf1074012.DriveRequestBuilder) {
    return id28cbea101cb9521a14cf9640c45800a00477bd62395913072f4e81bf1074012.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) Get(ctx context.Context, requestConfiguration *ListRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable), nil
}
// Items the items property
func (m *ListRequestBuilder) Items()(*idf9a3059d40c90c42068fa8bf9ac9fd67dedf1182980c80d7a384967f9dff8d5.ItemsRequestBuilder) {
    return idf9a3059d40c90c42068fa8bf9ac9fd67dedf1182980c80d7a384967f9dff8d5.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*i58ec931c55ab0d0111918ef65feb8b5a53f2f824a19e864696540521faf828f4.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem%2Did"] = id
    }
    return i58ec931c55ab0d0111918ef65feb8b5a53f2f824a19e864696540521faf828f4.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ListRequestBuilder) Operations()(*ie43f83ef262031fdb1eb8da2814c4165f5f945383e15adf8746e6a96f3432e23.OperationsRequestBuilder) {
    return ie43f83ef262031fdb1eb8da2814c4165f5f945383e15adf8746e6a96f3432e23.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.operations.item collection
func (m *ListRequestBuilder) OperationsById(id string)(*i327c97fc6449280089d733ba8dee83d9d74247819100cf6629d3355abfea1166.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return i327c97fc6449280089d733ba8dee83d9d74247819100cf6629d3355abfea1166.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in drive
func (m *ListRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Subscriptions the subscriptions property
func (m *ListRequestBuilder) Subscriptions()(*i1c4ca68087caab22a544a9dc3fb1c29f6076a9264ac284502cbf04a28444aea7.SubscriptionsRequestBuilder) {
    return i1c4ca68087caab22a544a9dc3fb1c29f6076a9264ac284502cbf04a28444aea7.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*i37771808e5dcfef82cc1e8bc33ad871d583b797abbb1222c5dc6198791d77af2.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i37771808e5dcfef82cc1e8bc33ad871d583b797abbb1222c5dc6198791d77af2.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
