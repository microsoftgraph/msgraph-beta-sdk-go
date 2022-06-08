package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i199f4fe84420b18a43deaf0667685d67c47afaf49a79254dcc4a317b5afca849 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes"
    i68ae183c09cc78b27dc0a97dc7978e0b869692de627d8e525f322762c2995563 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/activities"
    i7226c2ae89f801e4b64e6b0edd09181cf93ea776caaaeb9088097745135b7c70 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/drive"
    i8550117bc8dba52f0401a4c1cd2a6c36bd55a7d49e62e678bd768efc96140c4d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/operations"
    i8bcaa2dbdfb8a394ff13d6435d39b59476db4815609b2e8151ab695ec81a0265 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/columns"
    i97017b616da5986cf55b57864c53d1b56efbacdd6545a614ade2f0456a8c8e3e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/subscriptions"
    ia21f197c95ba8505493d74f1ef1d48a5400edd8511cb962781f10325527d4d30 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items"
    i781ff56458c20fa1b5ea82ecc4814d2f31e569406c20158e1b3ad4ed52d395bd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/columns/item"
    i8f66e52499614b6eebc4d2df4d0864612a4704ec85bded58c0834e237a19e822 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item"
    i932d70cc9aea88447f5a87d9f6a7d12fcbd584063696775f85bdbcb52661f2ff "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/operations/item"
    ib6346b85256456286c5c40862b2612677f66fb89b5675f3cdb91584973429793 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/subscriptions/item"
    ieb63f6ff8a52c6339b6fd4c8844e2b1f41f1144c131720fadba69f9d07938a2b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item"
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
func (m *ListItemRequestBuilder) Activities()(*i68ae183c09cc78b27dc0a97dc7978e0b869692de627d8e525f322762c2995563.ActivitiesRequestBuilder) {
    return i68ae183c09cc78b27dc0a97dc7978e0b869692de627d8e525f322762c2995563.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *ListItemRequestBuilder) Columns()(*i8bcaa2dbdfb8a394ff13d6435d39b59476db4815609b2e8151ab695ec81a0265.ColumnsRequestBuilder) {
    return i8bcaa2dbdfb8a394ff13d6435d39b59476db4815609b2e8151ab695ec81a0265.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.columns.item collection
func (m *ListItemRequestBuilder) ColumnsById(id string)(*i781ff56458c20fa1b5ea82ecc4814d2f31e569406c20158e1b3ad4ed52d395bd.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i781ff56458c20fa1b5ea82ecc4814d2f31e569406c20158e1b3ad4ed52d395bd.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}{?%24select,%24expand}";
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
func (m *ListItemRequestBuilder) ContentTypes()(*i199f4fe84420b18a43deaf0667685d67c47afaf49a79254dcc4a317b5afca849.ContentTypesRequestBuilder) {
    return i199f4fe84420b18a43deaf0667685d67c47afaf49a79254dcc4a317b5afca849.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item collection
func (m *ListItemRequestBuilder) ContentTypesById(id string)(*i8f66e52499614b6eebc4d2df4d0864612a4704ec85bded58c0834e237a19e822.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return i8f66e52499614b6eebc4d2df4d0864612a4704ec85bded58c0834e237a19e822.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property lists for groups
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property lists for groups
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
// CreateGetRequestInformation the collection of lists under this site.
func (m *ListItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of lists under this site.
func (m *ListItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ListItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property lists in groups
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property lists in groups
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
// Delete delete navigation property lists for groups
func (m *ListItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property lists for groups
func (m *ListItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ListItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *ListItemRequestBuilder) Drive()(*i7226c2ae89f801e4b64e6b0edd09181cf93ea776caaaeb9088097745135b7c70.DriveRequestBuilder) {
    return i7226c2ae89f801e4b64e6b0edd09181cf93ea776caaaeb9088097745135b7c70.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of lists under this site.
func (m *ListItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the collection of lists under this site.
func (m *ListItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ListItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
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
func (m *ListItemRequestBuilder) Items()(*ia21f197c95ba8505493d74f1ef1d48a5400edd8511cb962781f10325527d4d30.ItemsRequestBuilder) {
    return ia21f197c95ba8505493d74f1ef1d48a5400edd8511cb962781f10325527d4d30.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.items.item collection
func (m *ListItemRequestBuilder) ItemsById(id string)(*ieb63f6ff8a52c6339b6fd4c8844e2b1f41f1144c131720fadba69f9d07938a2b.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem%2Did"] = id
    }
    return ieb63f6ff8a52c6339b6fd4c8844e2b1f41f1144c131720fadba69f9d07938a2b.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ListItemRequestBuilder) Operations()(*i8550117bc8dba52f0401a4c1cd2a6c36bd55a7d49e62e678bd768efc96140c4d.OperationsRequestBuilder) {
    return i8550117bc8dba52f0401a4c1cd2a6c36bd55a7d49e62e678bd768efc96140c4d.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.operations.item collection
func (m *ListItemRequestBuilder) OperationsById(id string)(*i932d70cc9aea88447f5a87d9f6a7d12fcbd584063696775f85bdbcb52661f2ff.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return i932d70cc9aea88447f5a87d9f6a7d12fcbd584063696775f85bdbcb52661f2ff.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property lists in groups
func (m *ListItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property lists in groups
func (m *ListItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *ListItemRequestBuilder) Subscriptions()(*i97017b616da5986cf55b57864c53d1b56efbacdd6545a614ade2f0456a8c8e3e.SubscriptionsRequestBuilder) {
    return i97017b616da5986cf55b57864c53d1b56efbacdd6545a614ade2f0456a8c8e3e.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.subscriptions.item collection
func (m *ListItemRequestBuilder) SubscriptionsById(id string)(*ib6346b85256456286c5c40862b2612677f66fb89b5675f3cdb91584973429793.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return ib6346b85256456286c5c40862b2612677f66fb89b5675f3cdb91584973429793.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
