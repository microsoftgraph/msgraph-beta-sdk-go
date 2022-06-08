package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5cf3995870b2d384d03a75f1f97b58450462e62047e50280b5669fbd6850ba4a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/driveitem"
    i87c6521cb11a7b7f310be90831859523e3346e9943a4abb4a3629f0f454425ba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/fields"
    i9b0e4cc3908d3a9b47f534187288176496b751767f689bc36370e30debe0295a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/versions"
    ia19aa77d004359f91feb0f298caa1315142b0ecc99852cf73095b75cb2fa3245 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/documentsetversions"
    id73225c0893e43d4966050258b7710e87dd648a03be43ac07463794ce393d37a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/createlink"
    ida397936633293e6d5981730e13c9860d09348e2e96500d8fe54d4dfe93931d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/activities"
    ida662ff454dccde72d2055f3c72277d02ea243bcaf9896fe7a31bfd3b268aa3c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/analytics"
    iff7cc5c18fd4b0f8d74b41b15c56c7eefd58c5d986102b4fb0401363b5d38787 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i5a8dc4504b0b76b2f4fb8e97eb1ffd6af40bc3e65e9d3abb09262655d4345313 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/activities/item"
    i762c69f5e82f60e81750b0596721706fb0fc0cc4dff313deef523b89575a3777 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/documentsetversions/item"
    ie4db2192bf63f3da3a9c426cae2d5b7210000b306510f71907256751680fa8c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item/versions/item"
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
func (m *ListItemItemRequestBuilder) Activities()(*ida397936633293e6d5981730e13c9860d09348e2e96500d8fe54d4dfe93931d7.ActivitiesRequestBuilder) {
    return ida397936633293e6d5981730e13c9860d09348e2e96500d8fe54d4dfe93931d7.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.items.item.activities.item collection
func (m *ListItemItemRequestBuilder) ActivitiesById(id string)(*i5a8dc4504b0b76b2f4fb8e97eb1ffd6af40bc3e65e9d3abb09262655d4345313.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i5a8dc4504b0b76b2f4fb8e97eb1ffd6af40bc3e65e9d3abb09262655d4345313.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *ListItemItemRequestBuilder) Analytics()(*ida662ff454dccde72d2055f3c72277d02ea243bcaf9896fe7a31bfd3b268aa3c.AnalyticsRequestBuilder) {
    return ida662ff454dccde72d2055f3c72277d02ea243bcaf9896fe7a31bfd3b268aa3c.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemItemRequestBuilderInternal instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemItemRequestBuilder) {
    m := &ListItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/list/items/{listItem%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property items for users
func (m *ListItemItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property items for users
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
func (m *ListItemItemRequestBuilder) CreateLink()(*id73225c0893e43d4966050258b7710e87dd648a03be43ac07463794ce393d37a.CreateLinkRequestBuilder) {
    return id73225c0893e43d4966050258b7710e87dd648a03be43ac07463794ce393d37a.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in users
func (m *ListItemItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property items in users
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
// Delete delete navigation property items for users
func (m *ListItemItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property items for users
func (m *ListItemItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ListItemItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// DocumentSetVersions the documentSetVersions property
func (m *ListItemItemRequestBuilder) DocumentSetVersions()(*ia19aa77d004359f91feb0f298caa1315142b0ecc99852cf73095b75cb2fa3245.DocumentSetVersionsRequestBuilder) {
    return ia19aa77d004359f91feb0f298caa1315142b0ecc99852cf73095b75cb2fa3245.NewDocumentSetVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DocumentSetVersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.items.item.documentSetVersions.item collection
func (m *ListItemItemRequestBuilder) DocumentSetVersionsById(id string)(*i762c69f5e82f60e81750b0596721706fb0fc0cc4dff313deef523b89575a3777.DocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["documentSetVersion%2Did"] = id
    }
    return i762c69f5e82f60e81750b0596721706fb0fc0cc4dff313deef523b89575a3777.NewDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DriveItem the driveItem property
func (m *ListItemItemRequestBuilder) DriveItem()(*i5cf3995870b2d384d03a75f1f97b58450462e62047e50280b5669fbd6850ba4a.DriveItemRequestBuilder) {
    return i5cf3995870b2d384d03a75f1f97b58450462e62047e50280b5669fbd6850ba4a.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fields the fields property
func (m *ListItemItemRequestBuilder) Fields()(*i87c6521cb11a7b7f310be90831859523e3346e9943a4abb4a3629f0f454425ba.FieldsRequestBuilder) {
    return i87c6521cb11a7b7f310be90831859523e3346e9943a4abb4a3629f0f454425ba.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*iff7cc5c18fd4b0f8d74b41b15c56c7eefd58c5d986102b4fb0401363b5d38787.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return iff7cc5c18fd4b0f8d74b41b15c56c7eefd58c5d986102b4fb0401363b5d38787.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetWithRequestConfigurationAndResponseHandler all items contained in the list.
func (m *ListItemItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ListItemItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// Patch update the navigation property items in users
func (m *ListItemItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property items in users
func (m *ListItemItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Versions the versions property
func (m *ListItemItemRequestBuilder) Versions()(*i9b0e4cc3908d3a9b47f534187288176496b751767f689bc36370e30debe0295a.VersionsRequestBuilder) {
    return i9b0e4cc3908d3a9b47f534187288176496b751767f689bc36370e30debe0295a.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.items.item.versions.item collection
func (m *ListItemItemRequestBuilder) VersionsById(id string)(*ie4db2192bf63f3da3a9c426cae2d5b7210000b306510f71907256751680fa8c8.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion%2Did"] = id
    }
    return ie4db2192bf63f3da3a9c426cae2d5b7210000b306510f71907256751680fa8c8.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
