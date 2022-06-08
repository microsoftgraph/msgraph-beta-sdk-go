package listitem

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06dd18c22efea84624b18e1cc94c78983cc1c5bc99fd517c83daaa90ab4cb811 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i529eecf9a39472817b91cb99c0f2664c2fe83a61324142474841226d8cab6226 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/versions"
    i6fb339b032e7fa56506baff5189093b70298ffdaaf46c968119a0d926b13735f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/fields"
    i8c553aa1c8c4c2b81c0062f66e2868ab7452581bcd6757919ee9bb5a787db993 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/analytics"
    iaebdcfc54ab2edc6fe5105d1afafedf26835f7f668ac4b066860aa351101b1ad "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/activities"
    ic74031409af3859d635eedfd882b293877cb73a97e93264c90268bb5f3431d11 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/createlink"
    id7e57a35682e3412603d19488ff182fd358ae2b3be0be56ca2372e5b7cb1e1dd "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/driveitem"
    if45677a9d14c4f567148644e377a19ad27f80c951e3a29319a2f960337478a0b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/documentsetversions"
    i1fff618239a15c0e5436ecca8942e53d53dbdbd2ab6b75b3c031d5090bc95ba5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/documentsetversions/item"
    i4eec53871899220cfb963e8222bb0e0d672fb91c668739a3b13cc94d4a9b14b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/versions/item"
    ibe08c884ca571ea57b684ee03e3420c41ac87ffe76b7f866dfac0c1419574108 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem/activities/item"
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
// ListItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ListItemRequestBuilderGetQueryParameters for drives in SharePoint, the associated document library list item. Read-only. Nullable.
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
func (m *ListItemRequestBuilder) Activities()(*iaebdcfc54ab2edc6fe5105d1afafedf26835f7f668ac4b066860aa351101b1ad.ActivitiesRequestBuilder) {
    return iaebdcfc54ab2edc6fe5105d1afafedf26835f7f668ac4b066860aa351101b1ad.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*ibe08c884ca571ea57b684ee03e3420c41ac87ffe76b7f866dfac0c1419574108.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ibe08c884ca571ea57b684ee03e3420c41ac87ffe76b7f866dfac0c1419574108.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *ListItemRequestBuilder) Analytics()(*i8c553aa1c8c4c2b81c0062f66e2868ab7452581bcd6757919ee9bb5a787db993.AnalyticsRequestBuilder) {
    return i8c553aa1c8c4c2b81c0062f66e2868ab7452581bcd6757919ee9bb5a787db993.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem%2Did}/listItem{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property listItem for workbooks
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property listItem for workbooks
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
// CreateGetRequestInformation for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration for drives in SharePoint, the associated document library list item. Read-only. Nullable.
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
// CreateLink the createLink property
func (m *ListItemRequestBuilder) CreateLink()(*ic74031409af3859d635eedfd882b293877cb73a97e93264c90268bb5f3431d11.CreateLinkRequestBuilder) {
    return ic74031409af3859d635eedfd882b293877cb73a97e93264c90268bb5f3431d11.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property listItem in workbooks
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property listItem in workbooks
func (m *ListItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property listItem for workbooks
func (m *ListItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property listItem for workbooks
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
// DocumentSetVersions the documentSetVersions property
func (m *ListItemRequestBuilder) DocumentSetVersions()(*if45677a9d14c4f567148644e377a19ad27f80c951e3a29319a2f960337478a0b.DocumentSetVersionsRequestBuilder) {
    return if45677a9d14c4f567148644e377a19ad27f80c951e3a29319a2f960337478a0b.NewDocumentSetVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DocumentSetVersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.listItem.documentSetVersions.item collection
func (m *ListItemRequestBuilder) DocumentSetVersionsById(id string)(*i1fff618239a15c0e5436ecca8942e53d53dbdbd2ab6b75b3c031d5090bc95ba5.DocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["documentSetVersion%2Did"] = id
    }
    return i1fff618239a15c0e5436ecca8942e53d53dbdbd2ab6b75b3c031d5090bc95ba5.NewDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DriveItem the driveItem property
func (m *ListItemRequestBuilder) DriveItem()(*id7e57a35682e3412603d19488ff182fd358ae2b3be0be56ca2372e5b7cb1e1dd.DriveItemRequestBuilder) {
    return id7e57a35682e3412603d19488ff182fd358ae2b3be0be56ca2372e5b7cb1e1dd.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fields the fields property
func (m *ListItemRequestBuilder) Fields()(*i6fb339b032e7fa56506baff5189093b70298ffdaaf46c968119a0d926b13735f.FieldsRequestBuilder) {
    return i6fb339b032e7fa56506baff5189093b70298ffdaaf46c968119a0d926b13735f.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i06dd18c22efea84624b18e1cc94c78983cc1c5bc99fd517c83daaa90ab4cb811.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i06dd18c22efea84624b18e1cc94c78983cc1c5bc99fd517c83daaa90ab4cb811.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetWithRequestConfigurationAndResponseHandler for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ListItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
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
// Patch update the navigation property listItem in workbooks
func (m *ListItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property listItem in workbooks
func (m *ListItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *ListItemRequestBuilder) Versions()(*i529eecf9a39472817b91cb99c0f2664c2fe83a61324142474841226d8cab6226.VersionsRequestBuilder) {
    return i529eecf9a39472817b91cb99c0f2664c2fe83a61324142474841226d8cab6226.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*i4eec53871899220cfb963e8222bb0e0d672fb91c668739a3b13cc94d4a9b14b2.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion%2Did"] = id
    }
    return i4eec53871899220cfb963e8222bb0e0d672fb91c668739a3b13cc94d4a9b14b2.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
