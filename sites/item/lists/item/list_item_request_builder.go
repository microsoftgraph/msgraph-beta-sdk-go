package item

import (
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
// Activities the activities property
func (m *ListItemRequestBuilder) Activities()(*i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3.ActivitiesRequestBuilder) {
    return i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *ListItemRequestBuilder) Columns()(*i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462.ColumnsRequestBuilder) {
    return i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.columns.item collection
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
// ContentTypes the contentTypes property
func (m *ListItemRequestBuilder) ContentTypes()(*i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31.ContentTypesRequestBuilder) {
    return i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.contentTypes.item collection
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property lists for sites
func (m *ListItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property lists for sites
func (m *ListItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ListItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the collection of lists under this site.
func (m *ListItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of lists under this site.
func (m *ListItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ListItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property lists in sites
func (m *ListItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property lists in sites
func (m *ListItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property lists for sites
func (m *ListItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ListItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property lists for sites
func (m *ListItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ListItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Drive the drive property
func (m *ListItemRequestBuilder) Drive()(*i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f.DriveRequestBuilder) {
    return i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler the collection of lists under this site.
func (m *ListItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ListItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the collection of lists under this site.
func (m *ListItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ListItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable), nil
}
// Items the items property
func (m *ListItemRequestBuilder) Items()(*ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558.ItemsRequestBuilder) {
    return ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.items.item collection
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
// Operations the operations property
func (m *ListItemRequestBuilder) Operations()(*i22f86b00dbed60c4ae66f22a96d25d718083d7ee61dda8addf35f98b8a81a481.OperationsRequestBuilder) {
    return i22f86b00dbed60c4ae66f22a96d25d718083d7ee61dda8addf35f98b8a81a481.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.operations.item collection
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
// PatchWithResponseHandler update the navigation property lists in sites
func (m *ListItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property lists in sites
func (m *ListItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Subscriptions the subscriptions property
func (m *ListItemRequestBuilder) Subscriptions()(*iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40.SubscriptionsRequestBuilder) {
    return iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.subscriptions.item collection
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
