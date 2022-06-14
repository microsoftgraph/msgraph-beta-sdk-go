package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i23240d4635f216590f410f77081a54298f8183e1a329b9313bd70c3b9c0398d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/versions"
    i3698548c75c9e39e8fb37d84501a95bbaa309a2be53ae11e8068b466fb837edf "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/follow"
    i61f85fe2c179ac73fc8f411fba8d669196ef105d3cbac9b6cd26b88f3d6e2daa "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/searchwithq"
    i64505e7006add25c659273d4b29738b87096b7ed3bed7899ec8784a67328f67c "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/unfollow"
    i6fc391d816f83b511a377e0649a19546c17517d7156bee27a0b5e396658c369b "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/deltawithtoken"
    i729a9ee6e2e9a291b7ec5a3294f22d77c6f16f6cd791c00f7c4c705a74185103 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/invite"
    i77ee9f4d7f51930527bd106deff5233e33a5832f7e02c7c7dfcf39968e4b2595 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/listitem"
    i78b611da304a58f409193fcb269ce876f2de4c497aeac99809467f5f23bf077e "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i7992c9f9c3d1d5de3cd5daa7d95568f9e4e084d10c63f869beb57e17aea0a27f "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/extractsensitivitylabels"
    i7f80a059caf26cb0557b3b0f4acd7a1b13bb7927d29175db5e068ee693bad276 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/thumbnails"
    i80f5e3a8301b45f22eb0f59585b799701ceb9e98ac351f9e1541012396037ad6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/preview"
    i954ecda5a4250e454578c9bacd0686854e1407cf2feb765f6cbb2e2d4639ecfe "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/content"
    ia4956cd409576e24c4aa76f6d3abd67096a891f666aad2ddb18d169d38f7f254 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/children"
    iae53b43de7e8e82906e25f30ec44b807a656a48e247cc73e93c24ef4cdb9adc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/permissions"
    ib4404c45473f1403377012c51b5b5fbff2ffb0ca563bb9042a3ae81be04d8aa8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/activities"
    ibea3ca638c4c8d93207234447e43ba64990217c9fafd80b6d427e39ccd3aaaad "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/copy"
    ic984a15e7b75dd60cb8f87ca75945db21aecb4b349cbd56d9995006181423109 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/restore"
    icc1d8e152b2f554ccb54fda55a962fac9c1e045537c93edb3b999cafc75977a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/createuploadsession"
    icd7c61588f864663407a53847007d25105b1333bc61df92298519a2c08f367ee "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/validatepermission"
    ie5daf9c8e0450ffc27f9d326f49220da640adf90a49c38a8b09fefbfd680e9c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/delta"
    ie8bdf53c33c7ec4380140c27084ba54054d35b946201613b91a3e1d0b77ed053 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/checkin"
    iece16db43f069ff40a7b7d44a4a7fbc3e549375a87ae6ef54bb6ef94a53d7e17 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/createlink"
    if0998234fbf75af83cedc8709563ce46e8eda3db501048f5b8c0c4d1e60e56d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/checkout"
    if3106bc1fff17deee799295366ec0b5272378755fcc134899eb2437efd3be114 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/analytics"
    if5f6c3ed5cc5b1541e3ba162586eba48333bdf1020c351bf4165c6853c4f0850 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/subscriptions"
    i58f23b76fc4f6634505fdda414ef4371d8aefc19833903fc6bca5ac6e996154f "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/thumbnails/item"
    i5ab78c0964728c9027b740a5691a473d3cbd675dae4ab7fad69cb939e380cf48 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/activities/item"
    i808430ad30f2d332f1f7b6406ba69d7860f148b3363158053e59cb9f665daad0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/permissions/item"
    i863da0d27d3a9ecb60ce962d85288cafcc5b01550621c6f4fee5a0c94b9cfb5f "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/versions/item"
    ic61938f40528e15a355ad1902a63177598cd8fce7acad295259450ae2f9001f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/children/item"
    ie2e93e47590691a8700eb5c1241b3863000367220310a926746fc8a99e259943 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item/subscriptions/item"
)

// DriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type DriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type DriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DriveItemItemRequestBuilderGetQueryParameters
}
// DriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *DriveItemItemRequestBuilder) Activities()(*ib4404c45473f1403377012c51b5b5fbff2ffb0ca563bb9042a3ae81be04d8aa8.ActivitiesRequestBuilder) {
    return ib4404c45473f1403377012c51b5b5fbff2ffb0ca563bb9042a3ae81be04d8aa8.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i5ab78c0964728c9027b740a5691a473d3cbd675dae4ab7fad69cb939e380cf48.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i5ab78c0964728c9027b740a5691a473d3cbd675dae4ab7fad69cb939e380cf48.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *DriveItemItemRequestBuilder) Analytics()(*if3106bc1fff17deee799295366ec0b5272378755fcc134899eb2437efd3be114.AnalyticsRequestBuilder) {
    return if3106bc1fff17deee799295366ec0b5272378755fcc134899eb2437efd3be114.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *DriveItemItemRequestBuilder) Checkin()(*ie8bdf53c33c7ec4380140c27084ba54054d35b946201613b91a3e1d0b77ed053.CheckinRequestBuilder) {
    return ie8bdf53c33c7ec4380140c27084ba54054d35b946201613b91a3e1d0b77ed053.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *DriveItemItemRequestBuilder) Checkout()(*if0998234fbf75af83cedc8709563ce46e8eda3db501048f5b8c0c4d1e60e56d2.CheckoutRequestBuilder) {
    return if0998234fbf75af83cedc8709563ce46e8eda3db501048f5b8c0c4d1e60e56d2.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *DriveItemItemRequestBuilder) Children()(*ia4956cd409576e24c4aa76f6d3abd67096a891f666aad2ddb18d169d38f7f254.ChildrenRequestBuilder) {
    return ia4956cd409576e24c4aa76f6d3abd67096a891f666aad2ddb18d169d38f7f254.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*ic61938f40528e15a355ad1902a63177598cd8fce7acad295259450ae2f9001f4.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return ic61938f40528e15a355ad1902a63177598cd8fce7acad295259450ae2f9001f4.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/items/{driveItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *DriveItemItemRequestBuilder) Content()(*i954ecda5a4250e454578c9bacd0686854e1407cf2feb765f6cbb2e2d4639ecfe.ContentRequestBuilder) {
    return i954ecda5a4250e454578c9bacd0686854e1407cf2feb765f6cbb2e2d4639ecfe.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *DriveItemItemRequestBuilder) Copy()(*ibea3ca638c4c8d93207234447e43ba64990217c9fafd80b6d427e39ccd3aaaad.CopyRequestBuilder) {
    return ibea3ca638c4c8d93207234447e43ba64990217c9fafd80b6d427e39ccd3aaaad.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for drive
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property items for drive
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DriveItemItemRequestBuilder) CreateLink()(*iece16db43f069ff40a7b7d44a4a7fbc3e549375a87ae6ef54bb6ef94a53d7e17.CreateLinkRequestBuilder) {
    return iece16db43f069ff40a7b7d44a4a7fbc3e549375a87ae6ef54bb6ef94a53d7e17.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in drive
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property items in drive
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DriveItemItemRequestBuilder) CreateUploadSession()(*icc1d8e152b2f554ccb54fda55a962fac9c1e045537c93edb3b999cafc75977a4.CreateUploadSessionRequestBuilder) {
    return icc1d8e152b2f554ccb54fda55a962fac9c1e045537c93edb3b999cafc75977a4.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property items for drive
func (m *DriveItemItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property items for drive
func (m *DriveItemItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *DriveItemItemRequestBuilder) Delta()(*ie5daf9c8e0450ffc27f9d326f49220da640adf90a49c38a8b09fefbfd680e9c4.DeltaRequestBuilder) {
    return ie5daf9c8e0450ffc27f9d326f49220da640adf90a49c38a8b09fefbfd680e9c4.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *DriveItemItemRequestBuilder) DeltaWithToken(token *string)(*i6fc391d816f83b511a377e0649a19546c17517d7156bee27a0b5e396658c369b.DeltaWithTokenRequestBuilder) {
    return i6fc391d816f83b511a377e0649a19546c17517d7156bee27a0b5e396658c369b.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels the extractSensitivityLabels property
func (m *DriveItemItemRequestBuilder) ExtractSensitivityLabels()(*i7992c9f9c3d1d5de3cd5daa7d95568f9e4e084d10c63f869beb57e17aea0a27f.ExtractSensitivityLabelsRequestBuilder) {
    return i7992c9f9c3d1d5de3cd5daa7d95568f9e4e084d10c63f869beb57e17aea0a27f.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow the follow property
func (m *DriveItemItemRequestBuilder) Follow()(*i3698548c75c9e39e8fb37d84501a95bbaa309a2be53ae11e8068b466fb837edf.FollowRequestBuilder) {
    return i3698548c75c9e39e8fb37d84501a95bbaa309a2be53ae11e8068b466fb837edf.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *DriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i78b611da304a58f409193fcb269ce876f2de4c497aeac99809467f5f23bf077e.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i78b611da304a58f409193fcb269ce876f2de4c497aeac99809467f5f23bf077e.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetWithRequestConfigurationAndResponseHandler all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *DriveItemItemRequestBuilder) Invite()(*i729a9ee6e2e9a291b7ec5a3294f22d77c6f16f6cd791c00f7c4c705a74185103.InviteRequestBuilder) {
    return i729a9ee6e2e9a291b7ec5a3294f22d77c6f16f6cd791c00f7c4c705a74185103.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *DriveItemItemRequestBuilder) ListItem()(*i77ee9f4d7f51930527bd106deff5233e33a5832f7e02c7c7dfcf39968e4b2595.ListItemRequestBuilder) {
    return i77ee9f4d7f51930527bd106deff5233e33a5832f7e02c7c7dfcf39968e4b2595.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in drive
func (m *DriveItemItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property items in drive
func (m *DriveItemItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*iae53b43de7e8e82906e25f30ec44b807a656a48e247cc73e93c24ef4cdb9adc6.PermissionsRequestBuilder) {
    return iae53b43de7e8e82906e25f30ec44b807a656a48e247cc73e93c24ef4cdb9adc6.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i808430ad30f2d332f1f7b6406ba69d7860f148b3363158053e59cb9f665daad0.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i808430ad30f2d332f1f7b6406ba69d7860f148b3363158053e59cb9f665daad0.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *DriveItemItemRequestBuilder) Preview()(*i80f5e3a8301b45f22eb0f59585b799701ceb9e98ac351f9e1541012396037ad6.PreviewRequestBuilder) {
    return i80f5e3a8301b45f22eb0f59585b799701ceb9e98ac351f9e1541012396037ad6.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *DriveItemItemRequestBuilder) Restore()(*ic984a15e7b75dd60cb8f87ca75945db21aecb4b349cbd56d9995006181423109.RestoreRequestBuilder) {
    return ic984a15e7b75dd60cb8f87ca75945db21aecb4b349cbd56d9995006181423109.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemItemRequestBuilder) SearchWithQ(q *string)(*i61f85fe2c179ac73fc8f411fba8d669196ef105d3cbac9b6cd26b88f3d6e2daa.SearchWithQRequestBuilder) {
    return i61f85fe2c179ac73fc8f411fba8d669196ef105d3cbac9b6cd26b88f3d6e2daa.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *DriveItemItemRequestBuilder) Subscriptions()(*if5f6c3ed5cc5b1541e3ba162586eba48333bdf1020c351bf4165c6853c4f0850.SubscriptionsRequestBuilder) {
    return if5f6c3ed5cc5b1541e3ba162586eba48333bdf1020c351bf4165c6853c4f0850.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*ie2e93e47590691a8700eb5c1241b3863000367220310a926746fc8a99e259943.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return ie2e93e47590691a8700eb5c1241b3863000367220310a926746fc8a99e259943.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i7f80a059caf26cb0557b3b0f4acd7a1b13bb7927d29175db5e068ee693bad276.ThumbnailsRequestBuilder) {
    return i7f80a059caf26cb0557b3b0f4acd7a1b13bb7927d29175db5e068ee693bad276.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i58f23b76fc4f6634505fdda414ef4371d8aefc19833903fc6bca5ac6e996154f.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i58f23b76fc4f6634505fdda414ef4371d8aefc19833903fc6bca5ac6e996154f.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *DriveItemItemRequestBuilder) Unfollow()(*i64505e7006add25c659273d4b29738b87096b7ed3bed7899ec8784a67328f67c.UnfollowRequestBuilder) {
    return i64505e7006add25c659273d4b29738b87096b7ed3bed7899ec8784a67328f67c.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *DriveItemItemRequestBuilder) ValidatePermission()(*icd7c61588f864663407a53847007d25105b1333bc61df92298519a2c08f367ee.ValidatePermissionRequestBuilder) {
    return icd7c61588f864663407a53847007d25105b1333bc61df92298519a2c08f367ee.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *DriveItemItemRequestBuilder) Versions()(*i23240d4635f216590f410f77081a54298f8183e1a329b9313bd70c3b9c0398d2.VersionsRequestBuilder) {
    return i23240d4635f216590f410f77081a54298f8183e1a329b9313bd70c3b9c0398d2.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i863da0d27d3a9ecb60ce962d85288cafcc5b01550621c6f4fee5a0c94b9cfb5f.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i863da0d27d3a9ecb60ce962d85288cafcc5b01550621c6f4fee5a0c94b9cfb5f.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
