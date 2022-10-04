package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07b471b5bafcb166afc807037d8dca3bdcfce89c630e1f472f35241f937766d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/checkout"
    i217423aac1c727efdab2e5c54952470b0921f5b13d098bdf8b94800ab37d04ff "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/versions"
    i2960dd668f9d968328a1273c3905b10cd1454ba2be4d2d2518d741e39d34cb1a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/permissions"
    i3902be54e62f5079f0ed69218f21f0975e7e00f321907137ffaf2bab27d5ce6b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/extractsensitivitylabels"
    i465b2509f766b56413a989b9b1093c9dd885bd5ad9db93359e73becb56b4c8ed "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/subscriptions"
    i58ed6f131b10b06780d1ea2df19e365e95a719a3fae87e2d329f879f79982e2c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/createlink"
    i6269317abc0b6e5c9803549f001b49914559564d4e0e8925af00cee0a9f962e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/follow"
    i665e64b386f667323a17f15408c03a41c4bbf6c57b33a3bb5c2ef9c2a59a89a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/unfollow"
    i72e51e0fc013425d48885922454ee34dbb8d4c67dd8e586706e0731e06e927eb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/searchwithq"
    i75f959b4e251aa6fc4d58408e8f6ad4d90b04f678c2d5865b2bb00ddc543d577 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/invite"
    i7cd979a1135fa526bb0775b1fe3e6db8a84c98bb45e14a90ed979ce4b9fe1a23 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/preview"
    i81e3ed87008521c326b3ac3c0fc74dbb4bb9dde9ae30e3d99226e6c7ba4eb8f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/thumbnails"
    i92430ec89e7e973775f796d4ac8a8d9780b1bc0d5ad60e7cd57a383e6eef3d5d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/children"
    ia251bf4684ce6c652f9eb3c25934c42fe972ed1dd2719fe5854f1afbc301666a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/listitem"
    iad395222d19c225316722cb38722f8b7da3d5f20975a099108596127caf6c6fd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/delta"
    iaeabdcfc742f6374216f8726dbe42c3ec126484f5bc5c2b32acc5da0596d09ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/content"
    ibb2924ce4ae728b75cf62004f1f8878598fef9529ab8ba8de2f9c554e3591178 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/restore"
    ibb5ac4e4188591c61cd474cbec05f091eaaaba4116317142742d10ea41981219 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/assignsensitivitylabel"
    ic42ef09192fb45a9bc79a029a037222c3ed47782a3d3fc03f7992b62465b21a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/createuploadsession"
    id227a8ee2058982dcd1edd822aaf329d207e2ebbdf0a996a1f0a43952e1e4914 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/checkin"
    id376c2ee4463875800fde0bd7871a64ea129ca4a963393950fa963b6045c2b79 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/deltawithtoken"
    id7af1634662a3667e2519e43a9966c001812bce63e54a87d04dc5ce0a920b5e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/validatepermission"
    id7edf73240a09746f687259da2696c96a0bb4b0b25e7251ea5334efcfc767d7c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/copy"
    idba818df722d655b17638c47b441ec5821743a8896803f37a5aab7f7e02a3c0a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/activities"
    idce2f25461c814df9cbddafba89d9ae2f557d73761223fcda23c80b4f5f3ee97 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/analytics"
    if80a1dc70b76d38f62d13149533b0cbd572868ce071b457f1a7383667be61ef2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i461ab09785c4611bbf76605b4487a9f56eecb53b5d26bfb47eb14a0eadbe1bb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/children/item"
    i46b6ddeff5b488527d3295386cc098d2c87ca02641d9736e959f2cf64b66f45f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/versions/item"
    i55d3cce86d6128683aaeda467695b6cc98199af11165ef0493911461655f52da "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/subscriptions/item"
    i730e2e95d5c36abb56a8162182bfcf1bdb4f3d68c9721d0c03062451b2f30c28 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/thumbnails/item"
    i8b6ef88519f11f71a2ec7d5241d5e3086e28ed71a8d6b76a60af4f6ffdb80e54 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/activities/item"
    id992d48f7a3335ca926bd5d1d61ee50b229fce67be95588ae2b9110f96ad30d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/permissions/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*idba818df722d655b17638c47b441ec5821743a8896803f37a5aab7f7e02a3c0a.ActivitiesRequestBuilder) {
    return idba818df722d655b17638c47b441ec5821743a8896803f37a5aab7f7e02a3c0a.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i8b6ef88519f11f71a2ec7d5241d5e3086e28ed71a8d6b76a60af4f6ffdb80e54.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i8b6ef88519f11f71a2ec7d5241d5e3086e28ed71a8d6b76a60af4f6ffdb80e54.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *DriveItemItemRequestBuilder) Analytics()(*idce2f25461c814df9cbddafba89d9ae2f557d73761223fcda23c80b4f5f3ee97.AnalyticsRequestBuilder) {
    return idce2f25461c814df9cbddafba89d9ae2f557d73761223fcda23c80b4f5f3ee97.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel the assignSensitivityLabel property
func (m *DriveItemItemRequestBuilder) AssignSensitivityLabel()(*ibb5ac4e4188591c61cd474cbec05f091eaaaba4116317142742d10ea41981219.AssignSensitivityLabelRequestBuilder) {
    return ibb5ac4e4188591c61cd474cbec05f091eaaaba4116317142742d10ea41981219.NewAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *DriveItemItemRequestBuilder) Checkin()(*id227a8ee2058982dcd1edd822aaf329d207e2ebbdf0a996a1f0a43952e1e4914.CheckinRequestBuilder) {
    return id227a8ee2058982dcd1edd822aaf329d207e2ebbdf0a996a1f0a43952e1e4914.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *DriveItemItemRequestBuilder) Checkout()(*i07b471b5bafcb166afc807037d8dca3bdcfce89c630e1f472f35241f937766d7.CheckoutRequestBuilder) {
    return i07b471b5bafcb166afc807037d8dca3bdcfce89c630e1f472f35241f937766d7.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *DriveItemItemRequestBuilder) Children()(*i92430ec89e7e973775f796d4ac8a8d9780b1bc0d5ad60e7cd57a383e6eef3d5d.ChildrenRequestBuilder) {
    return i92430ec89e7e973775f796d4ac8a8d9780b1bc0d5ad60e7cd57a383e6eef3d5d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i461ab09785c4611bbf76605b4487a9f56eecb53b5d26bfb47eb14a0eadbe1bb0.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return i461ab09785c4611bbf76605b4487a9f56eecb53b5d26bfb47eb14a0eadbe1bb0.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/drives/{drive%2Did}/items/{driveItem%2Did}{?%24select,%24expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*iaeabdcfc742f6374216f8726dbe42c3ec126484f5bc5c2b32acc5da0596d09ea.ContentRequestBuilder) {
    return iaeabdcfc742f6374216f8726dbe42c3ec126484f5bc5c2b32acc5da0596d09ea.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *DriveItemItemRequestBuilder) Copy()(*id7edf73240a09746f687259da2696c96a0bb4b0b25e7251ea5334efcfc767d7c.CopyRequestBuilder) {
    return id7edf73240a09746f687259da2696c96a0bb4b0b25e7251ea5334efcfc767d7c.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for groups
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DriveItemItemRequestBuilder) CreateLink()(*i58ed6f131b10b06780d1ea2df19e365e95a719a3fae87e2d329f879f79982e2c.CreateLinkRequestBuilder) {
    return i58ed6f131b10b06780d1ea2df19e365e95a719a3fae87e2d329f879f79982e2c.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in groups
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateUploadSession the createUploadSession property
func (m *DriveItemItemRequestBuilder) CreateUploadSession()(*ic42ef09192fb45a9bc79a029a037222c3ed47782a3d3fc03f7992b62465b21a7.CreateUploadSessionRequestBuilder) {
    return ic42ef09192fb45a9bc79a029a037222c3ed47782a3d3fc03f7992b62465b21a7.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property items for groups
func (m *DriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Delta provides operations to call the delta method.
func (m *DriveItemItemRequestBuilder) Delta()(*iad395222d19c225316722cb38722f8b7da3d5f20975a099108596127caf6c6fd.DeltaRequestBuilder) {
    return iad395222d19c225316722cb38722f8b7da3d5f20975a099108596127caf6c6fd.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *DriveItemItemRequestBuilder) DeltaWithToken(token *string)(*id376c2ee4463875800fde0bd7871a64ea129ca4a963393950fa963b6045c2b79.DeltaWithTokenRequestBuilder) {
    return id376c2ee4463875800fde0bd7871a64ea129ca4a963393950fa963b6045c2b79.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels the extractSensitivityLabels property
func (m *DriveItemItemRequestBuilder) ExtractSensitivityLabels()(*i3902be54e62f5079f0ed69218f21f0975e7e00f321907137ffaf2bab27d5ce6b.ExtractSensitivityLabelsRequestBuilder) {
    return i3902be54e62f5079f0ed69218f21f0975e7e00f321907137ffaf2bab27d5ce6b.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow the follow property
func (m *DriveItemItemRequestBuilder) Follow()(*i6269317abc0b6e5c9803549f001b49914559564d4e0e8925af00cee0a9f962e4.FollowRequestBuilder) {
    return i6269317abc0b6e5c9803549f001b49914559564d4e0e8925af00cee0a9f962e4.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *DriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*if80a1dc70b76d38f62d13149533b0cbd572868ce071b457f1a7383667be61ef2.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return if80a1dc70b76d38f62d13149533b0cbd572868ce071b457f1a7383667be61ef2.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite the invite property
func (m *DriveItemItemRequestBuilder) Invite()(*i75f959b4e251aa6fc4d58408e8f6ad4d90b04f678c2d5865b2bb00ddc543d577.InviteRequestBuilder) {
    return i75f959b4e251aa6fc4d58408e8f6ad4d90b04f678c2d5865b2bb00ddc543d577.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *DriveItemItemRequestBuilder) ListItem()(*ia251bf4684ce6c652f9eb3c25934c42fe972ed1dd2719fe5854f1afbc301666a.ListItemRequestBuilder) {
    return ia251bf4684ce6c652f9eb3c25934c42fe972ed1dd2719fe5854f1afbc301666a.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in groups
func (m *DriveItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// Permissions the permissions property
func (m *DriveItemItemRequestBuilder) Permissions()(*i2960dd668f9d968328a1273c3905b10cd1454ba2be4d2d2518d741e39d34cb1a.PermissionsRequestBuilder) {
    return i2960dd668f9d968328a1273c3905b10cd1454ba2be4d2d2518d741e39d34cb1a.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*id992d48f7a3335ca926bd5d1d61ee50b229fce67be95588ae2b9110f96ad30d8.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return id992d48f7a3335ca926bd5d1d61ee50b229fce67be95588ae2b9110f96ad30d8.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *DriveItemItemRequestBuilder) Preview()(*i7cd979a1135fa526bb0775b1fe3e6db8a84c98bb45e14a90ed979ce4b9fe1a23.PreviewRequestBuilder) {
    return i7cd979a1135fa526bb0775b1fe3e6db8a84c98bb45e14a90ed979ce4b9fe1a23.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *DriveItemItemRequestBuilder) Restore()(*ibb2924ce4ae728b75cf62004f1f8878598fef9529ab8ba8de2f9c554e3591178.RestoreRequestBuilder) {
    return ibb2924ce4ae728b75cf62004f1f8878598fef9529ab8ba8de2f9c554e3591178.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemItemRequestBuilder) SearchWithQ(q *string)(*i72e51e0fc013425d48885922454ee34dbb8d4c67dd8e586706e0731e06e927eb.SearchWithQRequestBuilder) {
    return i72e51e0fc013425d48885922454ee34dbb8d4c67dd8e586706e0731e06e927eb.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i465b2509f766b56413a989b9b1093c9dd885bd5ad9db93359e73becb56b4c8ed.SubscriptionsRequestBuilder) {
    return i465b2509f766b56413a989b9b1093c9dd885bd5ad9db93359e73becb56b4c8ed.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i55d3cce86d6128683aaeda467695b6cc98199af11165ef0493911461655f52da.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i55d3cce86d6128683aaeda467695b6cc98199af11165ef0493911461655f52da.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i81e3ed87008521c326b3ac3c0fc74dbb4bb9dde9ae30e3d99226e6c7ba4eb8f5.ThumbnailsRequestBuilder) {
    return i81e3ed87008521c326b3ac3c0fc74dbb4bb9dde9ae30e3d99226e6c7ba4eb8f5.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i730e2e95d5c36abb56a8162182bfcf1bdb4f3d68c9721d0c03062451b2f30c28.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i730e2e95d5c36abb56a8162182bfcf1bdb4f3d68c9721d0c03062451b2f30c28.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *DriveItemItemRequestBuilder) Unfollow()(*i665e64b386f667323a17f15408c03a41c4bbf6c57b33a3bb5c2ef9c2a59a89a8.UnfollowRequestBuilder) {
    return i665e64b386f667323a17f15408c03a41c4bbf6c57b33a3bb5c2ef9c2a59a89a8.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *DriveItemItemRequestBuilder) ValidatePermission()(*id7af1634662a3667e2519e43a9966c001812bce63e54a87d04dc5ce0a920b5e5.ValidatePermissionRequestBuilder) {
    return id7af1634662a3667e2519e43a9966c001812bce63e54a87d04dc5ce0a920b5e5.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *DriveItemItemRequestBuilder) Versions()(*i217423aac1c727efdab2e5c54952470b0921f5b13d098bdf8b94800ab37d04ff.VersionsRequestBuilder) {
    return i217423aac1c727efdab2e5c54952470b0921f5b13d098bdf8b94800ab37d04ff.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i46b6ddeff5b488527d3295386cc098d2c87ca02641d9736e959f2cf64b66f45f.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i46b6ddeff5b488527d3295386cc098d2c87ca02641d9736e959f2cf64b66f45f.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
