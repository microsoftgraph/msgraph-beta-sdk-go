package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0162d60b270184ff2da0d229a6627422de6b3376b9ec1ca110a43407f8e33cdc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/invite"
    i03fda11e671eb36bee13e25e3e9b7862084dc785a14f093a05e62df80681796f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/content"
    i23f743e3e59d3542f75db30e3b39223f6f1c3f485139a340010e208a5d7f562b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/versions"
    i301be4eddbf1da3b8030fcab9029d5633867d2e39cca78a60413de3ee7b95c9c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/extractsensitivitylabels"
    i4f8e007051c45c635dc78e9437790431f5f43e6c3ca6f3e74074ca9d2edaa711 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/delta"
    i59f438e50c6a80a1653b614801cfeb12f2b820296bbce653018840e1f55f5d16 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/validatepermission"
    i5e0caf385fd0a72673862609cf32c4402e9eba33ef56846ecb3c64b17416e3b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/checkin"
    i5efcbe4b461c812355323caa9dc418905255f803cf1cf6260dcc892a8b9ed564 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/children"
    i6823f4a9450f5e2aa2bf9bf6496d95df699a41e10d6b1d2835c8421e1d753dc9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/assignsensitivitylabel"
    i6c289201e7f2e775278d4b92176cf3c09061495aef9ede11695d4e3154732e84 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/preview"
    i6f5104209c23bee9190bacc5292755c47101c4e311b94f75247635a6b88cb4cc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/checkout"
    i6fe8622edd27093c27a424c6ec65a9999da88f9aac0c0a17d00d0e4ecf7b5316 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/deltawithtoken"
    i7021ce989d3be4c224e26d0dbaa6136089065fa4c7c08b6f8d132f6d6588c963 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/analytics"
    i77758e00d72451ad1b4205c84cefed0be6e4f6ac7d828e7a81067a9b36e5826c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/createlink"
    i914d35933988d4207bb270b7495822ddb2cdff917df7876985c635fc1775a412 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/thumbnails"
    ia3bb93bc9befb9c2bcebd5daca5f25eb48c4c0564a67fc189f5a53eea0bdf60c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem"
    icf1fa6f1e0e31fd36a304be74aac9ab1ab95af067f125091e8362942b1b55449 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    id46ba11cea86eb4bf383db15698a2f61b22837a4471d1e4c005d929b629d0f0a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/createuploadsession"
    id8c4b260bf14a68a4eb2e0e6d105e7150363bd59af872fd759d47beabb4949d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/restore"
    idbe8275856d8fedd2fa248418909c5b36129b3c0b7154d2c961811b3cf35e829 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/activities"
    idf0b80a60e82ebe95e49f8a72b4fdaff6d5011f039b985bdaa81f830d809be12 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/permissions"
    ie5653038f868a3cbbaecd1f0765e4a47d8fc0c86da6fb48e843e543bf5f195eb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/searchwithq"
    ie9bb17dc21d71a2c13f1e1eead7c3d82a3e62b61d690414c415a2b68d9de5531 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/copy"
    ieca79f66e2aea726ad8626cf7a82ca108697b92513af5d6bb1b70b3abf83b17a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/follow"
    if15a9ba7729fdebc3c57a8ca44847546466efb149c4df186ff7d8cf6395e1c35 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/subscriptions"
    if6192a1f1dcd3cb12f0afd42452485333668962820979acf9446d6fd49850f23 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/unfollow"
    i4842cc54fffeac38d470c14c899c0d0e0fddefe6bc8020310bfa6c3a7477bc1f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/subscriptions/item"
    i7310b5803cdc10efcb7c5025f842e7b6c7175de4f87b0ff20422795e648d9a26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/activities/item"
    i9a2829926ca071ddf884f456479b8ccc48890a1b9bcfaec84bf1efe1d866494d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/versions/item"
    i9d26c5ed3e1dc52c55976335bc31cd3e38d3224ac6417512d4910d03c1776584 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/thumbnails/item"
    ide9d84689ef4976f2fb708443e6cf9f693bf4dc8a4301d5ab6caf47d2aec7e07 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/permissions/item"
    iefaf12d1fdb776bb7078a229349c10b5181934746d31a87a8236cbed30664968 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/children/item"
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
// Activities provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Activities()(*idbe8275856d8fedd2fa248418909c5b36129b3c0b7154d2c961811b3cf35e829.ActivitiesRequestBuilder) {
    return idbe8275856d8fedd2fa248418909c5b36129b3c0b7154d2c961811b3cf35e829.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i7310b5803cdc10efcb7c5025f842e7b6c7175de4f87b0ff20422795e648d9a26.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i7310b5803cdc10efcb7c5025f842e7b6c7175de4f87b0ff20422795e648d9a26.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Analytics()(*i7021ce989d3be4c224e26d0dbaa6136089065fa4c7c08b6f8d132f6d6588c963.AnalyticsRequestBuilder) {
    return i7021ce989d3be4c224e26d0dbaa6136089065fa4c7c08b6f8d132f6d6588c963.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
func (m *DriveItemItemRequestBuilder) AssignSensitivityLabel()(*i6823f4a9450f5e2aa2bf9bf6496d95df699a41e10d6b1d2835c8421e1d753dc9.AssignSensitivityLabelRequestBuilder) {
    return i6823f4a9450f5e2aa2bf9bf6496d95df699a41e10d6b1d2835c8421e1d753dc9.NewAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin provides operations to call the checkin method.
func (m *DriveItemItemRequestBuilder) Checkin()(*i5e0caf385fd0a72673862609cf32c4402e9eba33ef56846ecb3c64b17416e3b0.CheckinRequestBuilder) {
    return i5e0caf385fd0a72673862609cf32c4402e9eba33ef56846ecb3c64b17416e3b0.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout provides operations to call the checkout method.
func (m *DriveItemItemRequestBuilder) Checkout()(*i6f5104209c23bee9190bacc5292755c47101c4e311b94f75247635a6b88cb4cc.CheckoutRequestBuilder) {
    return i6f5104209c23bee9190bacc5292755c47101c4e311b94f75247635a6b88cb4cc.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Children()(*i5efcbe4b461c812355323caa9dc418905255f803cf1cf6260dcc892a8b9ed564.ChildrenRequestBuilder) {
    return i5efcbe4b461c812355323caa9dc418905255f803cf1cf6260dcc892a8b9ed564.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*iefaf12d1fdb776bb7078a229349c10b5181934746d31a87a8236cbed30664968.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return iefaf12d1fdb776bb7078a229349c10b5181934746d31a87a8236cbed30664968.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/items/{driveItem%2Did}{?%24select,%24expand}";
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
// Content provides operations to manage the media for the user entity.
func (m *DriveItemItemRequestBuilder) Content()(*i03fda11e671eb36bee13e25e3e9b7862084dc785a14f093a05e62df80681796f.ContentRequestBuilder) {
    return i03fda11e671eb36bee13e25e3e9b7862084dc785a14f093a05e62df80681796f.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *DriveItemItemRequestBuilder) Copy()(*ie9bb17dc21d71a2c13f1e1eead7c3d82a3e62b61d690414c415a2b68d9de5531.CopyRequestBuilder) {
    return ie9bb17dc21d71a2c13f1e1eead7c3d82a3e62b61d690414c415a2b68d9de5531.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for users
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
// CreateLink provides operations to call the createLink method.
func (m *DriveItemItemRequestBuilder) CreateLink()(*i77758e00d72451ad1b4205c84cefed0be6e4f6ac7d828e7a81067a9b36e5826c.CreateLinkRequestBuilder) {
    return i77758e00d72451ad1b4205c84cefed0be6e4f6ac7d828e7a81067a9b36e5826c.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in users
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
// CreateUploadSession provides operations to call the createUploadSession method.
func (m *DriveItemItemRequestBuilder) CreateUploadSession()(*id46ba11cea86eb4bf383db15698a2f61b22837a4471d1e4c005d929b629d0f0a.CreateUploadSessionRequestBuilder) {
    return id46ba11cea86eb4bf383db15698a2f61b22837a4471d1e4c005d929b629d0f0a.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property items for users
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
func (m *DriveItemItemRequestBuilder) Delta()(*i4f8e007051c45c635dc78e9437790431f5f43e6c3ca6f3e74074ca9d2edaa711.DeltaRequestBuilder) {
    return i4f8e007051c45c635dc78e9437790431f5f43e6c3ca6f3e74074ca9d2edaa711.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *DriveItemItemRequestBuilder) DeltaWithToken(token *string)(*i6fe8622edd27093c27a424c6ec65a9999da88f9aac0c0a17d00d0e4ecf7b5316.DeltaWithTokenRequestBuilder) {
    return i6fe8622edd27093c27a424c6ec65a9999da88f9aac0c0a17d00d0e4ecf7b5316.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
func (m *DriveItemItemRequestBuilder) ExtractSensitivityLabels()(*i301be4eddbf1da3b8030fcab9029d5633867d2e39cca78a60413de3ee7b95c9c.ExtractSensitivityLabelsRequestBuilder) {
    return i301be4eddbf1da3b8030fcab9029d5633867d2e39cca78a60413de3ee7b95c9c.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow provides operations to call the follow method.
func (m *DriveItemItemRequestBuilder) Follow()(*ieca79f66e2aea726ad8626cf7a82ca108697b92513af5d6bb1b70b3abf83b17a.FollowRequestBuilder) {
    return ieca79f66e2aea726ad8626cf7a82ca108697b92513af5d6bb1b70b3abf83b17a.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *DriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*icf1fa6f1e0e31fd36a304be74aac9ab1ab95af067f125091e8362942b1b55449.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return icf1fa6f1e0e31fd36a304be74aac9ab1ab95af067f125091e8362942b1b55449.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite provides operations to call the invite method.
func (m *DriveItemItemRequestBuilder) Invite()(*i0162d60b270184ff2da0d229a6627422de6b3376b9ec1ca110a43407f8e33cdc.InviteRequestBuilder) {
    return i0162d60b270184ff2da0d229a6627422de6b3376b9ec1ca110a43407f8e33cdc.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) ListItem()(*ia3bb93bc9befb9c2bcebd5daca5f25eb48c4c0564a67fc189f5a53eea0bdf60c.ListItemRequestBuilder) {
    return ia3bb93bc9befb9c2bcebd5daca5f25eb48c4c0564a67fc189f5a53eea0bdf60c.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in users
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
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Permissions()(*idf0b80a60e82ebe95e49f8a72b4fdaff6d5011f039b985bdaa81f830d809be12.PermissionsRequestBuilder) {
    return idf0b80a60e82ebe95e49f8a72b4fdaff6d5011f039b985bdaa81f830d809be12.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*ide9d84689ef4976f2fb708443e6cf9f693bf4dc8a4301d5ab6caf47d2aec7e07.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return ide9d84689ef4976f2fb708443e6cf9f693bf4dc8a4301d5ab6caf47d2aec7e07.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview provides operations to call the preview method.
func (m *DriveItemItemRequestBuilder) Preview()(*i6c289201e7f2e775278d4b92176cf3c09061495aef9ede11695d4e3154732e84.PreviewRequestBuilder) {
    return i6c289201e7f2e775278d4b92176cf3c09061495aef9ede11695d4e3154732e84.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *DriveItemItemRequestBuilder) Restore()(*id8c4b260bf14a68a4eb2e0e6d105e7150363bd59af872fd759d47beabb4949d6.RestoreRequestBuilder) {
    return id8c4b260bf14a68a4eb2e0e6d105e7150363bd59af872fd759d47beabb4949d6.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemItemRequestBuilder) SearchWithQ(q *string)(*ie5653038f868a3cbbaecd1f0765e4a47d8fc0c86da6fb48e843e543bf5f195eb.SearchWithQRequestBuilder) {
    return ie5653038f868a3cbbaecd1f0765e4a47d8fc0c86da6fb48e843e543bf5f195eb.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Subscriptions()(*if15a9ba7729fdebc3c57a8ca44847546466efb149c4df186ff7d8cf6395e1c35.SubscriptionsRequestBuilder) {
    return if15a9ba7729fdebc3c57a8ca44847546466efb149c4df186ff7d8cf6395e1c35.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i4842cc54fffeac38d470c14c899c0d0e0fddefe6bc8020310bfa6c3a7477bc1f.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i4842cc54fffeac38d470c14c899c0d0e0fddefe6bc8020310bfa6c3a7477bc1f.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i914d35933988d4207bb270b7495822ddb2cdff917df7876985c635fc1775a412.ThumbnailsRequestBuilder) {
    return i914d35933988d4207bb270b7495822ddb2cdff917df7876985c635fc1775a412.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i9d26c5ed3e1dc52c55976335bc31cd3e38d3224ac6417512d4910d03c1776584.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i9d26c5ed3e1dc52c55976335bc31cd3e38d3224ac6417512d4910d03c1776584.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow provides operations to call the unfollow method.
func (m *DriveItemItemRequestBuilder) Unfollow()(*if6192a1f1dcd3cb12f0afd42452485333668962820979acf9446d6fd49850f23.UnfollowRequestBuilder) {
    return if6192a1f1dcd3cb12f0afd42452485333668962820979acf9446d6fd49850f23.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *DriveItemItemRequestBuilder) ValidatePermission()(*i59f438e50c6a80a1653b614801cfeb12f2b820296bbce653018840e1f55f5d16.ValidatePermissionRequestBuilder) {
    return i59f438e50c6a80a1653b614801cfeb12f2b820296bbce653018840e1f55f5d16.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Versions()(*i23f743e3e59d3542f75db30e3b39223f6f1c3f485139a340010e208a5d7f562b.VersionsRequestBuilder) {
    return i23f743e3e59d3542f75db30e3b39223f6f1c3f485139a340010e208a5d7f562b.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i9a2829926ca071ddf884f456479b8ccc48890a1b9bcfaec84bf1efe1d866494d.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i9a2829926ca071ddf884f456479b8ccc48890a1b9bcfaec84bf1efe1d866494d.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
