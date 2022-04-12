package listitem

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i64ae1d84e2e9a0974d14735811a31e62196e49283fde33b9a779b7853ff28b56 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/analytics"
    i6b9b8da539be9a6c273fe2c8dcecdfc26a5b1bf7e8e8f202f6deb0bcd564a9b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/createlink"
    i81fc16a5810338ae5eaf81f26ca306c1f9a1a889acc0a8cea057202d43d91b9b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/documentsetversions"
    i8945b6eb60a65696b4df93ab5efc672172cb0d679196e801cb4e28eeb9fb0865 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/driveitem"
    ibf91eb4a6ce30ec31bd0805887675a599843c7629664636fe8b5ecb48f7480b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/fields"
    id05040ed5254dcf194bcd1bd15a770552ed47016519add1dcd56d104056e6fba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/versions"
    ieb300141379dd260a2cb3c225dfd256b3fb7c13f5a003cbeb60084abd79919a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/activities"
    iec1e64c918db3676b71ca4e4d50d2daa510dbab111824342b8b33bb9b66afe43 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i66dfbff408e09e066422acb8ff194ff02556f7d9c0993b1d77e3d6774a3de68e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/versions/item"
    i7e1159c18c36cae92f859f8bd570e815634d97cc52917e0c8ee60b28e9d5d1d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/documentsetversions/item"
    if2aa3de729f0171c35eaeca993793a6ce89ca9dc030569cef344c187536f093f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem/activities/item"
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
func (m *ListItemRequestBuilder) Activities()(*ieb300141379dd260a2cb3c225dfd256b3fb7c13f5a003cbeb60084abd79919a2.ActivitiesRequestBuilder) {
    return ieb300141379dd260a2cb3c225dfd256b3fb7c13f5a003cbeb60084abd79919a2.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*if2aa3de729f0171c35eaeca993793a6ce89ca9dc030569cef344c187536f093f.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return if2aa3de729f0171c35eaeca993793a6ce89ca9dc030569cef344c187536f093f.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *ListItemRequestBuilder) Analytics()(*i64ae1d84e2e9a0974d14735811a31e62196e49283fde33b9a779b7853ff28b56.AnalyticsRequestBuilder) {
    return i64ae1d84e2e9a0974d14735811a31e62196e49283fde33b9a779b7853ff28b56.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/items/{driveItem%2Did}/listItem{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property listItem for users
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
func (m *ListItemRequestBuilder) CreateLink()(*i6b9b8da539be9a6c273fe2c8dcecdfc26a5b1bf7e8e8f202f6deb0bcd564a9b9.CreateLinkRequestBuilder) {
    return i6b9b8da539be9a6c273fe2c8dcecdfc26a5b1bf7e8e8f202f6deb0bcd564a9b9.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property listItem in users
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
// Delete delete navigation property listItem for users
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
func (m *ListItemRequestBuilder) DocumentSetVersions()(*i81fc16a5810338ae5eaf81f26ca306c1f9a1a889acc0a8cea057202d43d91b9b.DocumentSetVersionsRequestBuilder) {
    return i81fc16a5810338ae5eaf81f26ca306c1f9a1a889acc0a8cea057202d43d91b9b.NewDocumentSetVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DocumentSetVersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item.listItem.documentSetVersions.item collection
func (m *ListItemRequestBuilder) DocumentSetVersionsById(id string)(*i7e1159c18c36cae92f859f8bd570e815634d97cc52917e0c8ee60b28e9d5d1d2.DocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["documentSetVersion%2Did"] = id
    }
    return i7e1159c18c36cae92f859f8bd570e815634d97cc52917e0c8ee60b28e9d5d1d2.NewDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DriveItem the driveItem property
func (m *ListItemRequestBuilder) DriveItem()(*i8945b6eb60a65696b4df93ab5efc672172cb0d679196e801cb4e28eeb9fb0865.DriveItemRequestBuilder) {
    return i8945b6eb60a65696b4df93ab5efc672172cb0d679196e801cb4e28eeb9fb0865.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fields the fields property
func (m *ListItemRequestBuilder) Fields()(*ibf91eb4a6ce30ec31bd0805887675a599843c7629664636fe8b5ecb48f7480b5.FieldsRequestBuilder) {
    return ibf91eb4a6ce30ec31bd0805887675a599843c7629664636fe8b5ecb48f7480b5.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*iec1e64c918db3676b71ca4e4d50d2daa510dbab111824342b8b33bb9b66afe43.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return iec1e64c918db3676b71ca4e4d50d2daa510dbab111824342b8b33bb9b66afe43.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Patch update the navigation property listItem in users
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
func (m *ListItemRequestBuilder) Versions()(*id05040ed5254dcf194bcd1bd15a770552ed47016519add1dcd56d104056e6fba.VersionsRequestBuilder) {
    return id05040ed5254dcf194bcd1bd15a770552ed47016519add1dcd56d104056e6fba.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*i66dfbff408e09e066422acb8ff194ff02556f7d9c0993b1d77e3d6774a3de68e.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion%2Did"] = id
    }
    return i66dfbff408e09e066422acb8ff194ff02556f7d9c0993b1d77e3d6774a3de68e.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
