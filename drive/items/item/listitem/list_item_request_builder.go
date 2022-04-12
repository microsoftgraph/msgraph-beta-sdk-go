package listitem

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1540d5090f8acc9592e0736bdf16985e7ef3dffa9734d5d1a24247d637292b31 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i534e346aa6d50cee86e5b41151c44b00a684a3d6e4c8d93a05d29f0ed45c41be "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/versions"
    i54a82f30f0eddf5558df15669ed1a5ed6429eef2342e698bdd676cf9002b5672 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/createlink"
    i8162920159daaa4296ec7c46e858c7baff5831a913b6153680f8ddf32ee02a90 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/documentsetversions"
    i81a639372b141d9e3fbf0f5a4b527ec19be2d25fb40e0bbdc56f434c11e6160f "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/analytics"
    ibc5795d824248dddd06c5789cb5f5901be173d9ab8215fb2f7f11dba7b8c448e "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/driveitem"
    ic39e3bf15c22b7a89fa511211f912d4ebf9f4dddce13c5ddf66c5428656b943a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/activities"
    if6225d9b2bfa5465841da00c027c5be1174eb0a1eed688231a0b5e0cb344ea98 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/fields"
    i20ba21d4a798f4b63ecdb4d5e0846dcd1905681974f40d64c898325f3bc3a5ed "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/activities/item"
    i58c07bc8da8514406a7d3d7fac587f4202185a5aa2ea3744f148373190a829b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/versions/item"
    iaa76a3d705c0a5a0b6cf1af5bfd919cbc4e8eb9cc71341de26b04290e9423f70 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem/documentsetversions/item"
)

// ListItemRequestBuilder provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ListItemRequestBuilderDeleteOptions options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ListItemRequestBuilderGetOptions options for Get
type ListItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ListItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ListItemRequestBuilderGetQueryParameters for drives in SharePoint, the associated document library list item. Read-only. Nullable.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Activities the activities property
func (m *ListItemRequestBuilder) Activities()(*ic39e3bf15c22b7a89fa511211f912d4ebf9f4dddce13c5ddf66c5428656b943a.ActivitiesRequestBuilder) {
    return ic39e3bf15c22b7a89fa511211f912d4ebf9f4dddce13c5ddf66c5428656b943a.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i20ba21d4a798f4b63ecdb4d5e0846dcd1905681974f40d64c898325f3bc3a5ed.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i20ba21d4a798f4b63ecdb4d5e0846dcd1905681974f40d64c898325f3bc3a5ed.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *ListItemRequestBuilder) Analytics()(*i81a639372b141d9e3fbf0f5a4b527ec19be2d25fb40e0bbdc56f434c11e6160f.AnalyticsRequestBuilder) {
    return i81a639372b141d9e3fbf0f5a4b527ec19be2d25fb40e0bbdc56f434c11e6160f.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/items/{driveItem%2Did}/listItem{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property listItem for drive
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateLink the createLink property
func (m *ListItemRequestBuilder) CreateLink()(*i54a82f30f0eddf5558df15669ed1a5ed6429eef2342e698bdd676cf9002b5672.CreateLinkRequestBuilder) {
    return i54a82f30f0eddf5558df15669ed1a5ed6429eef2342e698bdd676cf9002b5672.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property listItem in drive
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(options *ListItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property listItem for drive
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DocumentSetVersions the documentSetVersions property
func (m *ListItemRequestBuilder) DocumentSetVersions()(*i8162920159daaa4296ec7c46e858c7baff5831a913b6153680f8ddf32ee02a90.DocumentSetVersionsRequestBuilder) {
    return i8162920159daaa4296ec7c46e858c7baff5831a913b6153680f8ddf32ee02a90.NewDocumentSetVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DocumentSetVersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item.listItem.documentSetVersions.item collection
func (m *ListItemRequestBuilder) DocumentSetVersionsById(id string)(*iaa76a3d705c0a5a0b6cf1af5bfd919cbc4e8eb9cc71341de26b04290e9423f70.DocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["documentSetVersion%2Did"] = id
    }
    return iaa76a3d705c0a5a0b6cf1af5bfd919cbc4e8eb9cc71341de26b04290e9423f70.NewDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DriveItem the driveItem property
func (m *ListItemRequestBuilder) DriveItem()(*ibc5795d824248dddd06c5789cb5f5901be173d9ab8215fb2f7f11dba7b8c448e.DriveItemRequestBuilder) {
    return ibc5795d824248dddd06c5789cb5f5901be173d9ab8215fb2f7f11dba7b8c448e.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fields the fields property
func (m *ListItemRequestBuilder) Fields()(*if6225d9b2bfa5465841da00c027c5be1174eb0a1eed688231a0b5e0cb344ea98.FieldsRequestBuilder) {
    return if6225d9b2bfa5465841da00c027c5be1174eb0a1eed688231a0b5e0cb344ea98.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i1540d5090f8acc9592e0736bdf16985e7ef3dffa9734d5d1a24247d637292b31.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i1540d5090f8acc9592e0736bdf16985e7ef3dffa9734d5d1a24247d637292b31.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Patch update the navigation property listItem in drive
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Versions the versions property
func (m *ListItemRequestBuilder) Versions()(*i534e346aa6d50cee86e5b41151c44b00a684a3d6e4c8d93a05d29f0ed45c41be.VersionsRequestBuilder) {
    return i534e346aa6d50cee86e5b41151c44b00a684a3d6e4c8d93a05d29f0ed45c41be.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*i58c07bc8da8514406a7d3d7fac587f4202185a5aa2ea3744f148373190a829b6.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion%2Did"] = id
    }
    return i58c07bc8da8514406a7d3d7fac587f4202185a5aa2ea3744f148373190a829b6.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
