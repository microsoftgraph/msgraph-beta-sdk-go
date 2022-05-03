package listitem

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01ecfce418a9bb2b97515777bac0965047afaa52f0f8f627b369926bd24d525e "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/driveitem"
    i096d91ddb6589d42ef9806e45e5ccf06f3102b3e6c266224c8d1a4835338cc76 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/fields"
    i1c39a3ac32f8122cdbd7fe62d1fd584f53c3329f8d3cd2522724528e6f924e5f "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/createlink"
    i594592100db1505a07a2d6fe3d157b5b10059946536a90d2abaabefe5f8ed1c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/documentsetversions"
    i7665139006f04ea7d43ae76069b86ed53d01e8fa9125fd2cd9a68901a0343bce "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/analytics"
    i966f9199b94eedf154890a1d63ffddd7ff4ecb111a25b48bfcdcdf877ab04a05 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/versions"
    id8976ba0dcb9a0f06fe1c66240d2265a2492183a48ab775eb5b2171bc62d65cf "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/activities"
    ie04a4825ecf4f980bf1adee99d0831f108ba64c160181c2b97638eca0316ec94 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    iba2912fd31abac144f7de4dec1764591c687f6c1f1cc33d1937e95e7427034e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/documentsetversions/item"
    ie9386d56c77afaedf91f321d7112111686ace8ac926ba7982fae4d079402b3a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/activities/item"
    if10a5dcd64065862f4c8f8e93e850db2c277b5c0146afe02c14496e8be257f9f "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem/versions/item"
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
func (m *ListItemRequestBuilder) Activities()(*id8976ba0dcb9a0f06fe1c66240d2265a2492183a48ab775eb5b2171bc62d65cf.ActivitiesRequestBuilder) {
    return id8976ba0dcb9a0f06fe1c66240d2265a2492183a48ab775eb5b2171bc62d65cf.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*ie9386d56c77afaedf91f321d7112111686ace8ac926ba7982fae4d079402b3a6.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ie9386d56c77afaedf91f321d7112111686ace8ac926ba7982fae4d079402b3a6.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *ListItemRequestBuilder) Analytics()(*i7665139006f04ea7d43ae76069b86ed53d01e8fa9125fd2cd9a68901a0343bce.AnalyticsRequestBuilder) {
    return i7665139006f04ea7d43ae76069b86ed53d01e8fa9125fd2cd9a68901a0343bce.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/root/listItem{?%24select,%24expand}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property listItem for drive
func (m *ListItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property listItem for drive
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
// CreateGetRequestInformationWithRequestConfiguration for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration for drives in SharePoint, the associated document library list item. Read-only. Nullable.
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
// CreateLink the createLink property
func (m *ListItemRequestBuilder) CreateLink()(*i1c39a3ac32f8122cdbd7fe62d1fd584f53c3329f8d3cd2522724528e6f924e5f.CreateLinkRequestBuilder) {
    return i1c39a3ac32f8122cdbd7fe62d1fd584f53c3329f8d3cd2522724528e6f924e5f.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property listItem in drive
func (m *ListItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property listItem in drive
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
// DeleteWithResponseHandler delete navigation property listItem for drive
func (m *ListItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ListItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property listItem for drive
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
// DocumentSetVersions the documentSetVersions property
func (m *ListItemRequestBuilder) DocumentSetVersions()(*i594592100db1505a07a2d6fe3d157b5b10059946536a90d2abaabefe5f8ed1c2.DocumentSetVersionsRequestBuilder) {
    return i594592100db1505a07a2d6fe3d157b5b10059946536a90d2abaabefe5f8ed1c2.NewDocumentSetVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DocumentSetVersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.listItem.documentSetVersions.item collection
func (m *ListItemRequestBuilder) DocumentSetVersionsById(id string)(*iba2912fd31abac144f7de4dec1764591c687f6c1f1cc33d1937e95e7427034e6.DocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["documentSetVersion%2Did"] = id
    }
    return iba2912fd31abac144f7de4dec1764591c687f6c1f1cc33d1937e95e7427034e6.NewDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DriveItem the driveItem property
func (m *ListItemRequestBuilder) DriveItem()(*i01ecfce418a9bb2b97515777bac0965047afaa52f0f8f627b369926bd24d525e.DriveItemRequestBuilder) {
    return i01ecfce418a9bb2b97515777bac0965047afaa52f0f8f627b369926bd24d525e.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fields the fields property
func (m *ListItemRequestBuilder) Fields()(*i096d91ddb6589d42ef9806e45e5ccf06f3102b3e6c266224c8d1a4835338cc76.FieldsRequestBuilder) {
    return i096d91ddb6589d42ef9806e45e5ccf06f3102b3e6c266224c8d1a4835338cc76.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ie04a4825ecf4f980bf1adee99d0831f108ba64c160181c2b97638eca0316ec94.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ie04a4825ecf4f980bf1adee99d0831f108ba64c160181c2b97638eca0316ec94.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetWithResponseHandler for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ListItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ListItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
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
// PatchWithResponseHandler update the navigation property listItem in drive
func (m *ListItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property listItem in drive
func (m *ListItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *ListItemRequestBuilder) Versions()(*i966f9199b94eedf154890a1d63ffddd7ff4ecb111a25b48bfcdcdf877ab04a05.VersionsRequestBuilder) {
    return i966f9199b94eedf154890a1d63ffddd7ff4ecb111a25b48bfcdcdf877ab04a05.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*if10a5dcd64065862f4c8f8e93e850db2c277b5c0146afe02c14496e8be257f9f.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion%2Did"] = id
    }
    return if10a5dcd64065862f4c8f8e93e850db2c277b5c0146afe02c14496e8be257f9f.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
