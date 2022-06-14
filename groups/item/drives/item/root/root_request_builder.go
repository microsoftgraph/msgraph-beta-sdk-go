package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0038b1d7c66ec87b720bb6f335bf260d01e756277ebb6253a9675ba71f125eeb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/thumbnails"
    i1e72d9b1b1a1d6c4e4ddb398f2c540cd5edd38cec114ed2a943a836ec666b886 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/preview"
    i2177573a33e937fa3e73804230d11e61195f46ba418b0cddbebe322e7456096a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/unfollow"
    i3ee022e362928301f95e98536126ae08be259ddb3d74fcc35882ec3d6f209cc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/content"
    i4f30060526a7ab017ec8b892eecc7382ed0c0c87f397eaa72449097df8531dce "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i5015a7d5a11f5438f05a340dde81e67dedd395f7d76b74acd28bda09eb60dd25 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/follow"
    i59e8bbfbd3f60bae567a8408523b60e2da6b7989b7e672dd97a3c767fd417f23 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/subscriptions"
    i5cadbf0e0cbe14e1247b46965fa621de54bf7eada89dec2f5a9b59c9477b1d58 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/searchwithq"
    i5ce3ec157cfddb136d0d349fb3cecea590a20e228fd29509b7b1497c1f04a2ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/listitem"
    i6838c65611d5c9690c205f0f506fadd9464e9b59415fc54bb3252f16603ef7db "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/checkout"
    i6ec1534d2a31fed146899b9d06252a57f3682a9d852217105bf67ef0be0d3e10 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/children"
    i78b760e588e7f0ad16b38496752a9e7e32b4d6121b9f376e0ce370e0d7838080 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/delta"
    i7b22c004b007f707fffb61f62aebd243618b82a7690df4f7676cfbfc95e7c8dd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/createlink"
    i7b30ff9f95b373c3e13607c33154d64f6e61f7672c1b9a986371754a359201c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/restore"
    i82154dd857d1f4147284ffd2c005f1cb11a2e6dba3e0d970db99c39f04c4ad3d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/activities"
    i8e06844303b316afa4bb273c79b3822aa8583f64316c2c08e3e3fa351d53c86e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/versions"
    i8eda8ffa1856e2550facf64d6ac2e6fd80c97db380da53825976359a2b82d66a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/copy"
    i96582fcc546146b23a12588326b0654e1cf912fe39660d39b83acb0df3350855 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/invite"
    ia31cf1768b70a6ea501154cc3e28967ccbcd8ddef4a8a7427614d3636d2897d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/checkin"
    ib3995c47851f422d62b4d3949e9b41ce690968919b5fbc57da073da93893317f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/deltawithtoken"
    id97719380bf38dd33a04dc42f078677496b6a3e8aab5f3da8752ffd389d429e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/analytics"
    idb035671a49671c9c3860c6e20e520df498ee814eaa53d85eba7bc03cdc024ed "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/permissions"
    ie6d3d6837faea48e9ede130461bc86ade18514ff7aca746f0f0ce7070db4637d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/validatepermission"
    ief8b40ca45102a40c5032c1741da38649d2295edc766e637c42a3ad0827ecfb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/createuploadsession"
    if9ef0f06284f96f058632b7442153b7e621c3cc4c8368233e8f1365e447995da "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/extractsensitivitylabels"
    i34c79bd133f2c52bf82779507f18904f43c5de372db9b004de5a4f83505809bf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/children/item"
    i532e3abc360ac6f1fa4f890bd2c94e559c2d68325145fc45dba0b117dcc9b68e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/versions/item"
    i7d0a17f40953bcd7090a60e2809adcb724cfcd3db5d588342f80400d3b778d2c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/subscriptions/item"
    i83b0f9503a085d5076598d259da4b52da18ce21d577b876440a3993997518ab9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/permissions/item"
    ib24f9a2a2766fa9430ff8f21a5a0491c4493d94f0d2bef85b623278ad14433ef "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/thumbnails/item"
    ie651bb323681d0ad0e4318a76c4128cd5d1eac694c815fdcf3388b964126ff0d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/activities/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RootRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RootRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RootRequestBuilderGetQueryParameters
}
// RootRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *RootRequestBuilder) Activities()(*i82154dd857d1f4147284ffd2c005f1cb11a2e6dba3e0d970db99c39f04c4ad3d.ActivitiesRequestBuilder) {
    return i82154dd857d1f4147284ffd2c005f1cb11a2e6dba3e0d970db99c39f04c4ad3d.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*ie651bb323681d0ad0e4318a76c4128cd5d1eac694c815fdcf3388b964126ff0d.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ie651bb323681d0ad0e4318a76c4128cd5d1eac694c815fdcf3388b964126ff0d.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*id97719380bf38dd33a04dc42f078677496b6a3e8aab5f3da8752ffd389d429e8.AnalyticsRequestBuilder) {
    return id97719380bf38dd33a04dc42f078677496b6a3e8aab5f3da8752ffd389d429e8.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *RootRequestBuilder) Checkin()(*ia31cf1768b70a6ea501154cc3e28967ccbcd8ddef4a8a7427614d3636d2897d9.CheckinRequestBuilder) {
    return ia31cf1768b70a6ea501154cc3e28967ccbcd8ddef4a8a7427614d3636d2897d9.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *RootRequestBuilder) Checkout()(*i6838c65611d5c9690c205f0f506fadd9464e9b59415fc54bb3252f16603ef7db.CheckoutRequestBuilder) {
    return i6838c65611d5c9690c205f0f506fadd9464e9b59415fc54bb3252f16603ef7db.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i6ec1534d2a31fed146899b9d06252a57f3682a9d852217105bf67ef0be0d3e10.ChildrenRequestBuilder) {
    return i6ec1534d2a31fed146899b9d06252a57f3682a9d852217105bf67ef0be0d3e10.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i34c79bd133f2c52bf82779507f18904f43c5de372db9b004de5a4f83505809bf.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i34c79bd133f2c52bf82779507f18904f43c5de372db9b004de5a4f83505809bf.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *RootRequestBuilder) Content()(*i3ee022e362928301f95e98536126ae08be259ddb3d74fcc35882ec3d6f209cc7.ContentRequestBuilder) {
    return i3ee022e362928301f95e98536126ae08be259ddb3d74fcc35882ec3d6f209cc7.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *RootRequestBuilder) Copy()(*i8eda8ffa1856e2550facf64d6ac2e6fd80c97db380da53825976359a2b82d66a.CopyRequestBuilder) {
    return i8eda8ffa1856e2550facf64d6ac2e6fd80c97db380da53825976359a2b82d66a.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for groups
func (m *RootRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for groups
func (m *RootRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RootRequestBuilder) CreateLink()(*i7b22c004b007f707fffb61f62aebd243618b82a7690df4f7676cfbfc95e7c8dd.CreateLinkRequestBuilder) {
    return i7b22c004b007f707fffb61f62aebd243618b82a7690df4f7676cfbfc95e7c8dd.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property root in groups
func (m *RootRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in groups
func (m *RootRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateUploadSession the createUploadSession property
func (m *RootRequestBuilder) CreateUploadSession()(*ief8b40ca45102a40c5032c1741da38649d2295edc766e637c42a3ad0827ecfb5.CreateUploadSessionRequestBuilder) {
    return ief8b40ca45102a40c5032c1741da38649d2295edc766e637c42a3ad0827ecfb5.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for groups
func (m *RootRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property root for groups
func (m *RootRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Delta provides operations to call the delta method.
func (m *RootRequestBuilder) Delta()(*i78b760e588e7f0ad16b38496752a9e7e32b4d6121b9f376e0ce370e0d7838080.DeltaRequestBuilder) {
    return i78b760e588e7f0ad16b38496752a9e7e32b4d6121b9f376e0ce370e0d7838080.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *RootRequestBuilder) DeltaWithToken(token *string)(*ib3995c47851f422d62b4d3949e9b41ce690968919b5fbc57da073da93893317f.DeltaWithTokenRequestBuilder) {
    return ib3995c47851f422d62b4d3949e9b41ce690968919b5fbc57da073da93893317f.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels the extractSensitivityLabels property
func (m *RootRequestBuilder) ExtractSensitivityLabels()(*if9ef0f06284f96f058632b7442153b7e621c3cc4c8368233e8f1365e447995da.ExtractSensitivityLabelsRequestBuilder) {
    return if9ef0f06284f96f058632b7442153b7e621c3cc4c8368233e8f1365e447995da.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow the follow property
func (m *RootRequestBuilder) Follow()(*i5015a7d5a11f5438f05a340dde81e67dedd395f7d76b74acd28bda09eb60dd25.FollowRequestBuilder) {
    return i5015a7d5a11f5438f05a340dde81e67dedd395f7d76b74acd28bda09eb60dd25.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *RootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i4f30060526a7ab017ec8b892eecc7382ed0c0c87f397eaa72449097df8531dce.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i4f30060526a7ab017ec8b892eecc7382ed0c0c87f397eaa72449097df8531dce.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetWithRequestConfigurationAndResponseHandler the root folder of the drive. Read-only.
func (m *RootRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RootRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// Invite the invite property
func (m *RootRequestBuilder) Invite()(*i96582fcc546146b23a12588326b0654e1cf912fe39660d39b83acb0df3350855.InviteRequestBuilder) {
    return i96582fcc546146b23a12588326b0654e1cf912fe39660d39b83acb0df3350855.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *RootRequestBuilder) ListItem()(*i5ce3ec157cfddb136d0d349fb3cecea590a20e228fd29509b7b1497c1f04a2ea.ListItemRequestBuilder) {
    return i5ce3ec157cfddb136d0d349fb3cecea590a20e228fd29509b7b1497c1f04a2ea.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in groups
func (m *RootRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property root in groups
func (m *RootRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Permissions the permissions property
func (m *RootRequestBuilder) Permissions()(*idb035671a49671c9c3860c6e20e520df498ee814eaa53d85eba7bc03cdc024ed.PermissionsRequestBuilder) {
    return idb035671a49671c9c3860c6e20e520df498ee814eaa53d85eba7bc03cdc024ed.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i83b0f9503a085d5076598d259da4b52da18ce21d577b876440a3993997518ab9.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i83b0f9503a085d5076598d259da4b52da18ce21d577b876440a3993997518ab9.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *RootRequestBuilder) Preview()(*i1e72d9b1b1a1d6c4e4ddb398f2c540cd5edd38cec114ed2a943a836ec666b886.PreviewRequestBuilder) {
    return i1e72d9b1b1a1d6c4e4ddb398f2c540cd5edd38cec114ed2a943a836ec666b886.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *RootRequestBuilder) Restore()(*i7b30ff9f95b373c3e13607c33154d64f6e61f7672c1b9a986371754a359201c7.RestoreRequestBuilder) {
    return i7b30ff9f95b373c3e13607c33154d64f6e61f7672c1b9a986371754a359201c7.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *RootRequestBuilder) SearchWithQ(q *string)(*i5cadbf0e0cbe14e1247b46965fa621de54bf7eada89dec2f5a9b59c9477b1d58.SearchWithQRequestBuilder) {
    return i5cadbf0e0cbe14e1247b46965fa621de54bf7eada89dec2f5a9b59c9477b1d58.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*i59e8bbfbd3f60bae567a8408523b60e2da6b7989b7e672dd97a3c767fd417f23.SubscriptionsRequestBuilder) {
    return i59e8bbfbd3f60bae567a8408523b60e2da6b7989b7e672dd97a3c767fd417f23.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i7d0a17f40953bcd7090a60e2809adcb724cfcd3db5d588342f80400d3b778d2c.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i7d0a17f40953bcd7090a60e2809adcb724cfcd3db5d588342f80400d3b778d2c.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i0038b1d7c66ec87b720bb6f335bf260d01e756277ebb6253a9675ba71f125eeb.ThumbnailsRequestBuilder) {
    return i0038b1d7c66ec87b720bb6f335bf260d01e756277ebb6253a9675ba71f125eeb.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*ib24f9a2a2766fa9430ff8f21a5a0491c4493d94f0d2bef85b623278ad14433ef.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return ib24f9a2a2766fa9430ff8f21a5a0491c4493d94f0d2bef85b623278ad14433ef.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *RootRequestBuilder) Unfollow()(*i2177573a33e937fa3e73804230d11e61195f46ba418b0cddbebe322e7456096a.UnfollowRequestBuilder) {
    return i2177573a33e937fa3e73804230d11e61195f46ba418b0cddbebe322e7456096a.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *RootRequestBuilder) ValidatePermission()(*ie6d3d6837faea48e9ede130461bc86ade18514ff7aca746f0f0ce7070db4637d.ValidatePermissionRequestBuilder) {
    return ie6d3d6837faea48e9ede130461bc86ade18514ff7aca746f0f0ce7070db4637d.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*i8e06844303b316afa4bb273c79b3822aa8583f64316c2c08e3e3fa351d53c86e.VersionsRequestBuilder) {
    return i8e06844303b316afa4bb273c79b3822aa8583f64316c2c08e3e3fa351d53c86e.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i532e3abc360ac6f1fa4f890bd2c94e559c2d68325145fc45dba0b117dcc9b68e.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i532e3abc360ac6f1fa4f890bd2c94e559c2d68325145fc45dba0b117dcc9b68e.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
