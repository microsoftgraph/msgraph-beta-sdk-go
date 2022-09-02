package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3494ba260c7e5cb742941c392b00916ff749fe19697d06484c3b7083eabe7c49 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/fields"
    i4b6b4c3281893ceaa0871aa95c8b6c1ae470e33bc991bfd9ff48b8116029e68a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i58c8d16874aac3827dc40a3dbf2328fea48f08417c9c41debf67a58f919c2720 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/driveitem"
    i5ba57fe34d693b4cded862b4e7ada5fe09f96c4be8b524c005dc2feda37966d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/activities"
    ia2b700d43ac5cd289f1eae1bb612360c49638c62dfb54b8c7e8228d6fd206e25 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/analytics"
    ic5c7e5c1201df8e08dd3c31a725a1b6b5c4b8c93b46ad9390ad32b9a00aa10a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/versions"
    icf63e1dabde99d58f91beae56e8ed47fc1a61a3ec747a3742ff506ffd784cb39 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/createlink"
    id7e04dec07271eefedbc9911823f28ba29508ddb380a7a3a501a5f7f888905e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/documentsetversions"
    i5d1a96d4205d20d609fdac7d4097b2ccb3a9aaca71b3808c13b8ce00cbe3ad03 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/documentsetversions/item"
    id15ff120ad8f053cfd1691895b73c567a4d9203be255a2354e8e1c0a99cce095 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/activities/item"
    ie00e09a3d9b974ec973d38ac015be59a09a07678fa7822e52487a59cc84f7610 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/versions/item"
)

// ListItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.list entity.
type ListItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ListItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ListItemItemRequestBuilderGetQueryParameters all items contained in the list.
type ListItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ListItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ListItemItemRequestBuilderGetQueryParameters
}
// ListItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *ListItemItemRequestBuilder) Activities()(*i5ba57fe34d693b4cded862b4e7ada5fe09f96c4be8b524c005dc2feda37966d7.ActivitiesRequestBuilder) {
    return i5ba57fe34d693b4cded862b4e7ada5fe09f96c4be8b524c005dc2feda37966d7.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.items.item.activities.item collection
func (m *ListItemItemRequestBuilder) ActivitiesById(id string)(*id15ff120ad8f053cfd1691895b73c567a4d9203be255a2354e8e1c0a99cce095.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return id15ff120ad8f053cfd1691895b73c567a4d9203be255a2354e8e1c0a99cce095.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *ListItemItemRequestBuilder) Analytics()(*ia2b700d43ac5cd289f1eae1bb612360c49638c62dfb54b8c7e8228d6fd206e25.AnalyticsRequestBuilder) {
    return ia2b700d43ac5cd289f1eae1bb612360c49638c62dfb54b8c7e8228d6fd206e25.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemItemRequestBuilderInternal instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemItemRequestBuilder) {
    m := &ListItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list/items/{listItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemItemRequestBuilder instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property items for drive
func (m *ListItemItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property items for drive
func (m *ListItemItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ListItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all items contained in the list.
func (m *ListItemItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration all items contained in the list.
func (m *ListItemItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ListItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateLink the createLink property
func (m *ListItemItemRequestBuilder) CreateLink()(*icf63e1dabde99d58f91beae56e8ed47fc1a61a3ec747a3742ff506ffd784cb39.CreateLinkRequestBuilder) {
    return icf63e1dabde99d58f91beae56e8ed47fc1a61a3ec747a3742ff506ffd784cb39.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in drive
func (m *ListItemItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property items in drive
func (m *ListItemItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property items for drive
func (m *ListItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DocumentSetVersions the documentSetVersions property
func (m *ListItemItemRequestBuilder) DocumentSetVersions()(*id7e04dec07271eefedbc9911823f28ba29508ddb380a7a3a501a5f7f888905e8.DocumentSetVersionsRequestBuilder) {
    return id7e04dec07271eefedbc9911823f28ba29508ddb380a7a3a501a5f7f888905e8.NewDocumentSetVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DocumentSetVersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.items.item.documentSetVersions.item collection
func (m *ListItemItemRequestBuilder) DocumentSetVersionsById(id string)(*i5d1a96d4205d20d609fdac7d4097b2ccb3a9aaca71b3808c13b8ce00cbe3ad03.DocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["documentSetVersion%2Did"] = id
    }
    return i5d1a96d4205d20d609fdac7d4097b2ccb3a9aaca71b3808c13b8ce00cbe3ad03.NewDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DriveItem the driveItem property
func (m *ListItemItemRequestBuilder) DriveItem()(*i58c8d16874aac3827dc40a3dbf2328fea48f08417c9c41debf67a58f919c2720.DriveItemRequestBuilder) {
    return i58c8d16874aac3827dc40a3dbf2328fea48f08417c9c41debf67a58f919c2720.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fields the fields property
func (m *ListItemItemRequestBuilder) Fields()(*i3494ba260c7e5cb742941c392b00916ff749fe19697d06484c3b7083eabe7c49.FieldsRequestBuilder) {
    return i3494ba260c7e5cb742941c392b00916ff749fe19697d06484c3b7083eabe7c49.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i4b6b4c3281893ceaa0871aa95c8b6c1ae470e33bc991bfd9ff48b8116029e68a.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i4b6b4c3281893ceaa0871aa95c8b6c1ae470e33bc991bfd9ff48b8116029e68a.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Patch update the navigation property items in drive
func (m *ListItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemItemRequestBuilderPatchRequestConfiguration)(error) {
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
// Versions the versions property
func (m *ListItemItemRequestBuilder) Versions()(*ic5c7e5c1201df8e08dd3c31a725a1b6b5c4b8c93b46ad9390ad32b9a00aa10a1.VersionsRequestBuilder) {
    return ic5c7e5c1201df8e08dd3c31a725a1b6b5c4b8c93b46ad9390ad32b9a00aa10a1.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.items.item.versions.item collection
func (m *ListItemItemRequestBuilder) VersionsById(id string)(*ie00e09a3d9b974ec973d38ac015be59a09a07678fa7822e52487a59cc84f7610.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion%2Did"] = id
    }
    return ie00e09a3d9b974ec973d38ac015be59a09a07678fa7822e52487a59cc84f7610.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
