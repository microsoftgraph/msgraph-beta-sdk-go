package root

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0cfb21b89f2be6662549c26a3797133b25184b3dfb5b75f7fcc39e668f1f52ba "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/children"
    i147093ed7991e42030ab95d45b758f6ff0886497bc3f35f9145ecfb5ab4cf972 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/validatepermission"
    i1692e579cd5c54e5cfc53e3dd472fd824f3ca544790efe902be59c28f39c8558 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/invite"
    i194dea4a5cf2ad4f3a404fa45669d2fcef5e75c6adf67e742182321ef35b313a "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/copy"
    i231c4f0071bf428f8c1c635e22b42ce9e07bb8a363bd5923bb2b37fe5f3ed612 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/delta"
    i2ec11f40d612b5807c0011058ee965304ebd8b3611d57b8da75917f4246d2cef "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/follow"
    i36eabd21e28fa780f76254e70543251347fc0637b588934f3852063c63424967 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/subscriptions"
    i3b7b946e0a2645fef9b2c6c8e07d4e9ede3efb19d13dc449ae1de82484c97638 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/searchwithq"
    i538aec6ec1a5dbcbc75dda7f47d60743e51c6eb965f1999aa7e8fa1b049b4fb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/thumbnails"
    i55c2129817d1116256382e18ea19914df03e1a7f25fde5c2e13a0c87abbcd254 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/permissions"
    i6bd353492bc61810ec38ada2077d610730092f7f1abf56a70b7844be07a71479 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/createlink"
    i74b126db4377dcc623e5aa0a5328f7a226211bb0ee96a00eb65c9baa493414c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i75f74a9f470214e89920f233558fd6a17cfa92ca6bf22dfcb4f98c21a1f009c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/content"
    i801d73b826f8fc1f3763aeff6f5e014dff9f25d52eb15ed850c94e347044db1e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/createuploadsession"
    i844a09e2a848a0518b30d6c9486b0fc21ed3d53251643cfcfff453511ccef490 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/versions"
    i8556f9208dc03b23bb1904471ffd2bdfe79b4f5c2b4dfbfcb65bb05f00abab2d "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem"
    i855c2ba178739d024ab839a55ddcbbbf71f4aee97413ab3de3d8695f0d209172 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/assignsensitivitylabel"
    i875aaa204b1a9bea90971933d7319d4e17363029c574b1710c05205fb046fc7a "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/restore"
    i87fdda47cce10fee2ce8b805258a71414e1d86b1e7b7b80b6675663c43d24998 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/checkout"
    i9907f12ddf0e4b4923f27ac4ecaf094fcc41006aa5f3d5ceaa202584f3f03e68 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/unfollow"
    i9cd042c9e46438c7bce87f358e11b3cca0d1ddc296c2c5c1db9328d1d7d50309 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/activities"
    ib9cdff2592b0acec4aa284ce6ebe0a2cf32a5af41f645119e427abf571a30ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/analytics"
    ibe37cf5f5dcf1bc1cfdc42ceb7a59c25e1dc1c699a70c21dea43c4c5097cf445 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/extractsensitivitylabels"
    ic4d262b3aa823ed7c601f1d61666a847b717099d55d610bdcf809ac175c19683 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/deltawithtoken"
    ie7ac4e617b5ab867f78ccf871edfdbdc63b65202a1e640eda695bd75e472ae4c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/preview"
    ie835a3104e62d1bbd236827884460eb5a11917d0263439a648887900a481be94 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/checkin"
    i0ad9c3bbc7ea3c357a65b8cc5f691784c52ce76f23b1eeb54bebe4e2e38fe6f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/versions/item"
    i4f1c46079242ae874121e4eaab792681034ad778ef6010021ce2f297e80d9e7e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/children/item"
    i57df13ad7fef0bc108a1f0f224b2282dbbdcc2899d051c8a902e93cbaa841679 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/thumbnails/item"
    i7c142e142c53d798fd36169ec23884ef699397425bef56513abf167c332f6077 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/subscriptions/item"
    ic63da4ae39cf08296b5759a146dd61e14200041401e8eecb075724a9f1b76d19 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/permissions/item"
    ie50aea2d159f88c27cbd91bde6c1de53b048faf70798c48a700e279c1e5aff78 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/activities/item"
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
// RootRequestBuilderGetQueryParameters retrieve the metadata for a driveItem in a drive by file system path or ID.
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
// Activities provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) Activities()(*i9cd042c9e46438c7bce87f358e11b3cca0d1ddc296c2c5c1db9328d1d7d50309.ActivitiesRequestBuilder) {
    return i9cd042c9e46438c7bce87f358e11b3cca0d1ddc296c2c5c1db9328d1d7d50309.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) ActivitiesById(id string)(*ie50aea2d159f88c27cbd91bde6c1de53b048faf70798c48a700e279c1e5aff78.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ie50aea2d159f88c27cbd91bde6c1de53b048faf70798c48a700e279c1e5aff78.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) Analytics()(*ib9cdff2592b0acec4aa284ce6ebe0a2cf32a5af41f645119e427abf571a30ea1.AnalyticsRequestBuilder) {
    return ib9cdff2592b0acec4aa284ce6ebe0a2cf32a5af41f645119e427abf571a30ea1.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
func (m *RootRequestBuilder) AssignSensitivityLabel()(*i855c2ba178739d024ab839a55ddcbbbf71f4aee97413ab3de3d8695f0d209172.AssignSensitivityLabelRequestBuilder) {
    return i855c2ba178739d024ab839a55ddcbbbf71f4aee97413ab3de3d8695f0d209172.NewAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin provides operations to call the checkin method.
func (m *RootRequestBuilder) Checkin()(*ie835a3104e62d1bbd236827884460eb5a11917d0263439a648887900a481be94.CheckinRequestBuilder) {
    return ie835a3104e62d1bbd236827884460eb5a11917d0263439a648887900a481be94.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout provides operations to call the checkout method.
func (m *RootRequestBuilder) Checkout()(*i87fdda47cce10fee2ce8b805258a71414e1d86b1e7b7b80b6675663c43d24998.CheckoutRequestBuilder) {
    return i87fdda47cce10fee2ce8b805258a71414e1d86b1e7b7b80b6675663c43d24998.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) Children()(*i0cfb21b89f2be6662549c26a3797133b25184b3dfb5b75f7fcc39e668f1f52ba.ChildrenRequestBuilder) {
    return i0cfb21b89f2be6662549c26a3797133b25184b3dfb5b75f7fcc39e668f1f52ba.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) ChildrenById(id string)(*i4f1c46079242ae874121e4eaab792681034ad778ef6010021ce2f297e80d9e7e.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i4f1c46079242ae874121e4eaab792681034ad778ef6010021ce2f297e80d9e7e.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/root{?%24select,%24expand}";
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
// Content provides operations to manage the media for the drive entity.
func (m *RootRequestBuilder) Content()(*i75f74a9f470214e89920f233558fd6a17cfa92ca6bf22dfcb4f98c21a1f009c7.ContentRequestBuilder) {
    return i75f74a9f470214e89920f233558fd6a17cfa92ca6bf22dfcb4f98c21a1f009c7.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *RootRequestBuilder) Copy()(*i194dea4a5cf2ad4f3a404fa45669d2fcef5e75c6adf67e742182321ef35b313a.CopyRequestBuilder) {
    return i194dea4a5cf2ad4f3a404fa45669d2fcef5e75c6adf67e742182321ef35b313a.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for drives
func (m *RootRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the metadata for a driveItem in a drive by file system path or ID.
func (m *RootRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RootRequestBuilder) CreateLink()(*i6bd353492bc61810ec38ada2077d610730092f7f1abf56a70b7844be07a71479.CreateLinkRequestBuilder) {
    return i6bd353492bc61810ec38ada2077d610730092f7f1abf56a70b7844be07a71479.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property root in drives
func (m *RootRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RootRequestBuilder) CreateUploadSession()(*i801d73b826f8fc1f3763aeff6f5e014dff9f25d52eb15ed850c94e347044db1e.CreateUploadSessionRequestBuilder) {
    return i801d73b826f8fc1f3763aeff6f5e014dff9f25d52eb15ed850c94e347044db1e.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for drives
func (m *RootRequestBuilder) Delete(ctx context.Context, requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *RootRequestBuilder) Delta()(*i231c4f0071bf428f8c1c635e22b42ce9e07bb8a363bd5923bb2b37fe5f3ed612.DeltaRequestBuilder) {
    return i231c4f0071bf428f8c1c635e22b42ce9e07bb8a363bd5923bb2b37fe5f3ed612.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *RootRequestBuilder) DeltaWithToken(token *string)(*ic4d262b3aa823ed7c601f1d61666a847b717099d55d610bdcf809ac175c19683.DeltaWithTokenRequestBuilder) {
    return ic4d262b3aa823ed7c601f1d61666a847b717099d55d610bdcf809ac175c19683.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
func (m *RootRequestBuilder) ExtractSensitivityLabels()(*ibe37cf5f5dcf1bc1cfdc42ceb7a59c25e1dc1c699a70c21dea43c4c5097cf445.ExtractSensitivityLabelsRequestBuilder) {
    return ibe37cf5f5dcf1bc1cfdc42ceb7a59c25e1dc1c699a70c21dea43c4c5097cf445.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow provides operations to call the follow method.
func (m *RootRequestBuilder) Follow()(*i2ec11f40d612b5807c0011058ee965304ebd8b3611d57b8da75917f4246d2cef.FollowRequestBuilder) {
    return i2ec11f40d612b5807c0011058ee965304ebd8b3611d57b8da75917f4246d2cef.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the metadata for a driveItem in a drive by file system path or ID.
func (m *RootRequestBuilder) Get(ctx context.Context, requestConfiguration *RootRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *RootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i74b126db4377dcc623e5aa0a5328f7a226211bb0ee96a00eb65c9baa493414c6.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i74b126db4377dcc623e5aa0a5328f7a226211bb0ee96a00eb65c9baa493414c6.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite provides operations to call the invite method.
func (m *RootRequestBuilder) Invite()(*i1692e579cd5c54e5cfc53e3dd472fd824f3ca544790efe902be59c28f39c8558.InviteRequestBuilder) {
    return i1692e579cd5c54e5cfc53e3dd472fd824f3ca544790efe902be59c28f39c8558.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) ListItem()(*i8556f9208dc03b23bb1904471ffd2bdfe79b4f5c2b4dfbfcb65bb05f00abab2d.ListItemRequestBuilder) {
    return i8556f9208dc03b23bb1904471ffd2bdfe79b4f5c2b4dfbfcb65bb05f00abab2d.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in drives
func (m *RootRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *RootRequestBuilder) Permissions()(*i55c2129817d1116256382e18ea19914df03e1a7f25fde5c2e13a0c87abbcd254.PermissionsRequestBuilder) {
    return i55c2129817d1116256382e18ea19914df03e1a7f25fde5c2e13a0c87abbcd254.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) PermissionsById(id string)(*ic63da4ae39cf08296b5759a146dd61e14200041401e8eecb075724a9f1b76d19.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return ic63da4ae39cf08296b5759a146dd61e14200041401e8eecb075724a9f1b76d19.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview provides operations to call the preview method.
func (m *RootRequestBuilder) Preview()(*ie7ac4e617b5ab867f78ccf871edfdbdc63b65202a1e640eda695bd75e472ae4c.PreviewRequestBuilder) {
    return ie7ac4e617b5ab867f78ccf871edfdbdc63b65202a1e640eda695bd75e472ae4c.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *RootRequestBuilder) Restore()(*i875aaa204b1a9bea90971933d7319d4e17363029c574b1710c05205fb046fc7a.RestoreRequestBuilder) {
    return i875aaa204b1a9bea90971933d7319d4e17363029c574b1710c05205fb046fc7a.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *RootRequestBuilder) SearchWithQ(q *string)(*i3b7b946e0a2645fef9b2c6c8e07d4e9ede3efb19d13dc449ae1de82484c97638.SearchWithQRequestBuilder) {
    return i3b7b946e0a2645fef9b2c6c8e07d4e9ede3efb19d13dc449ae1de82484c97638.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) Subscriptions()(*i36eabd21e28fa780f76254e70543251347fc0637b588934f3852063c63424967.SubscriptionsRequestBuilder) {
    return i36eabd21e28fa780f76254e70543251347fc0637b588934f3852063c63424967.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i7c142e142c53d798fd36169ec23884ef699397425bef56513abf167c332f6077.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i7c142e142c53d798fd36169ec23884ef699397425bef56513abf167c332f6077.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) Thumbnails()(*i538aec6ec1a5dbcbc75dda7f47d60743e51c6eb965f1999aa7e8fa1b049b4fb1.ThumbnailsRequestBuilder) {
    return i538aec6ec1a5dbcbc75dda7f47d60743e51c6eb965f1999aa7e8fa1b049b4fb1.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i57df13ad7fef0bc108a1f0f224b2282dbbdcc2899d051c8a902e93cbaa841679.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i57df13ad7fef0bc108a1f0f224b2282dbbdcc2899d051c8a902e93cbaa841679.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow provides operations to call the unfollow method.
func (m *RootRequestBuilder) Unfollow()(*i9907f12ddf0e4b4923f27ac4ecaf094fcc41006aa5f3d5ceaa202584f3f03e68.UnfollowRequestBuilder) {
    return i9907f12ddf0e4b4923f27ac4ecaf094fcc41006aa5f3d5ceaa202584f3f03e68.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *RootRequestBuilder) ValidatePermission()(*i147093ed7991e42030ab95d45b758f6ff0886497bc3f35f9145ecfb5ab4cf972.ValidatePermissionRequestBuilder) {
    return i147093ed7991e42030ab95d45b758f6ff0886497bc3f35f9145ecfb5ab4cf972.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) Versions()(*i844a09e2a848a0518b30d6c9486b0fc21ed3d53251643cfcfff453511ccef490.VersionsRequestBuilder) {
    return i844a09e2a848a0518b30d6c9486b0fc21ed3d53251643cfcfff453511ccef490.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *RootRequestBuilder) VersionsById(id string)(*i0ad9c3bbc7ea3c357a65b8cc5f691784c52ce76f23b1eeb54bebe4e2e38fe6f2.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i0ad9c3bbc7ea3c357a65b8cc5f691784c52ce76f23b1eeb54bebe4e2e38fe6f2.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
