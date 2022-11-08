package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i171cba9a6593eebb15a25bdecb4e7b68992240964dd0dcf94faff9fec6922e8a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/createlink"
    i54a67566247c88e0df811a2ed946590294491d98944f520de597a8bbb6779536 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/documentsetversions"
    i5de2a26d9ca6c076decf6b96a9866a502a36f5b38dd182e17a4ce5c6da907b19 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/versions"
    i6222f486427b16385a7486ef974f1ba5ed889d63aa3d33e0b1b70ded2a5295eb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/activities"
    i8a77e844a33d05eed9cb2980f354e93c663b6bf99f751ac6e06d0bac5861bedc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/fields"
    ib72995c0e040129e0998e017fe97bcb09a69fe4c4f1b94aa27dba02ebb390aed "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/analytics"
    icd1189a97b569563d3109d2245a2eae6333ded6eb34f03bc7a07f9feeb5aa763 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    id1f83b88c5e9118a397fec8d377a56d0843a6059b3b36f13c15bd3dbaebced4b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/driveitem"
    i53d548c1157bc2734cecabf9ff7e0418612a368e2c19523093157ae2ff28323c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/documentsetversions/item"
    i616c47b3c949ac02a16dc2746109684cd8f063d58da7d523b5316e5dcbf3f398 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/versions/item"
    ib68cb8a1eaf3ee22b0407521b30cc967ee9a61e173ac3a5d450c69bfb69cb933 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item/activities/item"
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
// Activities provides operations to manage the activities property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) Activities()(*i6222f486427b16385a7486ef974f1ba5ed889d63aa3d33e0b1b70ded2a5295eb.ActivitiesRequestBuilder) {
    return i6222f486427b16385a7486ef974f1ba5ed889d63aa3d33e0b1b70ded2a5295eb.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) ActivitiesById(id string)(*ib68cb8a1eaf3ee22b0407521b30cc967ee9a61e173ac3a5d450c69bfb69cb933.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ib68cb8a1eaf3ee22b0407521b30cc967ee9a61e173ac3a5d450c69bfb69cb933.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) Analytics()(*ib72995c0e040129e0998e017fe97bcb09a69fe4c4f1b94aa27dba02ebb390aed.AnalyticsRequestBuilder) {
    return ib72995c0e040129e0998e017fe97bcb09a69fe4c4f1b94aa27dba02ebb390aed.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemItemRequestBuilderInternal instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemItemRequestBuilder) {
    m := &ListItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property items for groups
func (m *ListItemItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ListItemItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateLink provides operations to call the createLink method.
func (m *ListItemItemRequestBuilder) CreateLink()(*i171cba9a6593eebb15a25bdecb4e7b68992240964dd0dcf94faff9fec6922e8a.CreateLinkRequestBuilder) {
    return i171cba9a6593eebb15a25bdecb4e7b68992240964dd0dcf94faff9fec6922e8a.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in groups
func (m *ListItemItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property items for groups
func (m *ListItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DocumentSetVersions provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) DocumentSetVersions()(*i54a67566247c88e0df811a2ed946590294491d98944f520de597a8bbb6779536.DocumentSetVersionsRequestBuilder) {
    return i54a67566247c88e0df811a2ed946590294491d98944f520de597a8bbb6779536.NewDocumentSetVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DocumentSetVersionsById provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) DocumentSetVersionsById(id string)(*i53d548c1157bc2734cecabf9ff7e0418612a368e2c19523093157ae2ff28323c.DocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["documentSetVersion%2Did"] = id
    }
    return i53d548c1157bc2734cecabf9ff7e0418612a368e2c19523093157ae2ff28323c.NewDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DriveItem provides operations to manage the driveItem property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) DriveItem()(*id1f83b88c5e9118a397fec8d377a56d0843a6059b3b36f13c15bd3dbaebced4b.DriveItemRequestBuilder) {
    return id1f83b88c5e9118a397fec8d377a56d0843a6059b3b36f13c15bd3dbaebced4b.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fields provides operations to manage the fields property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) Fields()(*i8a77e844a33d05eed9cb2980f354e93c663b6bf99f751ac6e06d0bac5861bedc.FieldsRequestBuilder) {
    return i8a77e844a33d05eed9cb2980f354e93c663b6bf99f751ac6e06d0bac5861bedc.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
func (m *ListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*icd1189a97b569563d3109d2245a2eae6333ded6eb34f03bc7a07f9feeb5aa763.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return icd1189a97b569563d3109d2245a2eae6333ded6eb34f03bc7a07f9feeb5aa763.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Patch update the navigation property items in groups
func (m *ListItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
// Versions provides operations to manage the versions property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) Versions()(*i5de2a26d9ca6c076decf6b96a9866a502a36f5b38dd182e17a4ce5c6da907b19.VersionsRequestBuilder) {
    return i5de2a26d9ca6c076decf6b96a9866a502a36f5b38dd182e17a4ce5c6da907b19.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) VersionsById(id string)(*i616c47b3c949ac02a16dc2746109684cd8f063d58da7d523b5316e5dcbf3f398.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion%2Did"] = id
    }
    return i616c47b3c949ac02a16dc2746109684cd8f063d58da7d523b5316e5dcbf3f398.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
