package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities"
    i22f86b00dbed60c4ae66f22a96d25d718083d7ee61dda8addf35f98b8a81a481 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/operations"
    i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes"
    i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/columns"
    i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/drive"
    ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items"
    iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/subscriptions"
    i3b9fc89ed20f5bc6ee2709cbea9da1d72a73e1a237c50904ede0b8286e254701 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/operations/item"
    i574bb51b011ae57be83ae4c1378131adfd141ea2f28ea8d6fb084a7c2a316f9e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/subscriptions/item"
    i84a6362366aa96a7cf24a63ec6fc8d616b93a84df4ef10a2dba84a08f8fae66c "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/columns/item"
    i853471746fafe0cd69c71c0d080ce487ac5cbe90e44c1a8e2f9f2b94a267d3bb "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item"
    i877059848b5a1f38ddd1b25cd60b076bc0b0068b6a845b9627d81740e7eadef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item"
)

// ListItemRequestBuilder provides operations to manage the lists property of the microsoft.graph.site entity.
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ListItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ListItemRequestBuilderGetQueryParameters the collection of lists under this site.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ListItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ListItemRequestBuilderGetQueryParameters
}
// ListItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) Activities()(*i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3.ActivitiesRequestBuilder) {
    return i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns provides operations to manage the columns property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) Columns()(*i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462.ColumnsRequestBuilder) {
    return i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) ColumnsById(id string)(*i84a6362366aa96a7cf24a63ec6fc8d616b93a84df4ef10a2dba84a08f8fae66c.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i84a6362366aa96a7cf24a63ec6fc8d616b93a84df4ef10a2dba84a08f8fae66c.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) ContentTypes()(*i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31.ContentTypesRequestBuilder) {
    return i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById provides operations to manage the contentTypes property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) ContentTypesById(id string)(*i853471746fafe0cd69c71c0d080ce487ac5cbe90e44c1a8e2f9f2b94a267d3bb.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return i853471746fafe0cd69c71c0d080ce487ac5cbe90e44c1a8e2f9f2b94a267d3bb.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property lists for sites
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ListItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of lists under this site.
func (m *ListItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ListItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property lists in sites
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property lists for sites
func (m *ListItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ListItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Drive provides operations to manage the drive property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) Drive()(*i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f.DriveRequestBuilder) {
    return i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of lists under this site.
func (m *ListItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ListItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
// Items provides operations to manage the items property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) Items()(*ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558.ItemsRequestBuilder) {
    return ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) ItemsById(id string)(*i877059848b5a1f38ddd1b25cd60b076bc0b0068b6a845b9627d81740e7eadef6.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem%2Did"] = id
    }
    return i877059848b5a1f38ddd1b25cd60b076bc0b0068b6a845b9627d81740e7eadef6.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) Operations()(*i22f86b00dbed60c4ae66f22a96d25d718083d7ee61dda8addf35f98b8a81a481.OperationsRequestBuilder) {
    return i22f86b00dbed60c4ae66f22a96d25d718083d7ee61dda8addf35f98b8a81a481.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) OperationsById(id string)(*i3b9fc89ed20f5bc6ee2709cbea9da1d72a73e1a237c50904ede0b8286e254701.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return i3b9fc89ed20f5bc6ee2709cbea9da1d72a73e1a237c50904ede0b8286e254701.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property lists in sites
func (m *ListItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) Subscriptions()(*iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40.SubscriptionsRequestBuilder) {
    return iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.list entity.
func (m *ListItemRequestBuilder) SubscriptionsById(id string)(*i574bb51b011ae57be83ae4c1378131adfd141ea2f28ea8d6fb084a7c2a316f9e.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i574bb51b011ae57be83ae4c1378131adfd141ea2f28ea8d6fb084a7c2a316f9e.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
