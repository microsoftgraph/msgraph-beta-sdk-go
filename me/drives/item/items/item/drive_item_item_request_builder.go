package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i21e9fcf848cb9c39a2fc3f8761d9c40f31cbca3472bbe9a5cbcdbda4d175a8d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/searchwithq"
    i2279753cbb4e1319aa20ac9a28757bc6d3b4d382bc665bbb1dc4c248f57d5980 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/createlink"
    i3575e154580c5d16893024f44b2f73ea3b26700f3cb7778adb00c19667046a3f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/follow"
    i37ff40e2aa7be9eca0148887400d7b085d8f3de83bf87492286dc92f5a7ac5eb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/analytics"
    i39765dfa1c3eafd523949b6b410f599c0e93968da8276a68b09d6869ce0a7fc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/thumbnails"
    i3dca9f6bf310a5170317883bf7799e0ffa261a2b2542557f2280e1b46e7d54d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/delta"
    i43351f457508aecc9acc29067e433332ebc46c9941ad1ac6e1a5f33f8fcb1e5d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/children"
    i52fdb84788de6f4295e35bd2351366c8c5b985d4d378bc90530d073b39223a0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/subscriptions"
    i5b0a99c5b67d423f45d52d7551bcbc2389478eaa8636e90fc04fde1312b36353 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/content"
    i5f7f4ee0090cc95d75d7c261819f4cc11af67ee4cf962b90e283ca54b09a795b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/permissions"
    i606f41c5fea6c70fd34ccde609068e6bd09206e372d836cc7a3cd47d661d3be8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/listitem"
    i6f3eef764ea1ae21b6822debfa6f6a704cd6435b66bcebeb13a675ec3815bd92 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/versions"
    i76e213ef3af16d5a017724c1129453b8e46917fd1758262825fd2504a111c798 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/unfollow"
    iad57dbed394ee51d4a6e621363c22bc0787cacfea764a2fb293f50016c5ec458 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ib75035b590d4aa86c63fe0d7083badc6525d6470dcc873ca18ebe78c15b61131 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/preview"
    ib83a2fa775431a9271332b767a4c32b249184fbb6dd8404e43700f666927d598 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/copy"
    ibe93737f97f952ab01899b8d2e868778cadcb47a868980aee7049140a22620d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/restore"
    ibf93f730448d1b3d26834dbf6379c101e687bb7c2cc3cc1a3b8b6c8b2f0b3451 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/deltawithtoken"
    ic095af862cc8d74bb2e7b1e3a983f0498f1abc964e1622aa8e1e0df7b3a9e29d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/invite"
    ic1373904794125c73001b993b8b33c6016f56d96174b0260f5b8823ac4e00b57 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/validatepermission"
    icf15246792ca783226cabc2cc3a1ceb518e142ebed900e2c7c1ebecc572b5964 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/createuploadsession"
    id3ae126c5fa04f7016956d96c261d74a03a367f18c87cbd6c0c81c4773856482 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/checkin"
    id68215ba6f53cc548c3873fa209c4815d08e03d6e7263cc0cfcc3ab34ae4d18b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/assignsensitivitylabel"
    ie8ff27e9f76417e88fbdbaff20404d64201e0fb3256421b45e469dce450f5368 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/extractsensitivitylabels"
    if07779cafb32bcb954e8baa2fcdb6283357a647442af0a49aaa2643ddf7a2125 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/activities"
    if911b3a10b9730297fcd733857f877651ed5a7489ee33fe66b62a1e4b4c29263 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/checkout"
    i73ce730b8b40e7e732bae3aa850e526852f4bbcfac2c62b18d895fb8c40cb028 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/children/item"
    i8b02ac8460c4300fef83eeacf0e3a90fa6ad7e02a0ad5f87838abb6d5fb2ce82 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/subscriptions/item"
    ia7e12ef2cde2e1cee53c4ab5522f78d82906572a22510a1b2961ef668cc7a611 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/versions/item"
    ida01d8a4830793c46212039d68a6eefd4dfcee0c4bbfeb15920c2e1eeba59306 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/activities/item"
    ie01afcc24bcb8de9f524e134f72f107f5d9980714d4b3f83b8b0a3d66372a5c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/permissions/item"
    ie8883816406be65444fb3f248da99471dc749778a92112ecb2d943acde1ab3bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/thumbnails/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*if07779cafb32bcb954e8baa2fcdb6283357a647442af0a49aaa2643ddf7a2125.ActivitiesRequestBuilder) {
    return if07779cafb32bcb954e8baa2fcdb6283357a647442af0a49aaa2643ddf7a2125.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*ida01d8a4830793c46212039d68a6eefd4dfcee0c4bbfeb15920c2e1eeba59306.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ida01d8a4830793c46212039d68a6eefd4dfcee0c4bbfeb15920c2e1eeba59306.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *DriveItemItemRequestBuilder) Analytics()(*i37ff40e2aa7be9eca0148887400d7b085d8f3de83bf87492286dc92f5a7ac5eb.AnalyticsRequestBuilder) {
    return i37ff40e2aa7be9eca0148887400d7b085d8f3de83bf87492286dc92f5a7ac5eb.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel the assignSensitivityLabel property
func (m *DriveItemItemRequestBuilder) AssignSensitivityLabel()(*id68215ba6f53cc548c3873fa209c4815d08e03d6e7263cc0cfcc3ab34ae4d18b.AssignSensitivityLabelRequestBuilder) {
    return id68215ba6f53cc548c3873fa209c4815d08e03d6e7263cc0cfcc3ab34ae4d18b.NewAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *DriveItemItemRequestBuilder) Checkin()(*id3ae126c5fa04f7016956d96c261d74a03a367f18c87cbd6c0c81c4773856482.CheckinRequestBuilder) {
    return id3ae126c5fa04f7016956d96c261d74a03a367f18c87cbd6c0c81c4773856482.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *DriveItemItemRequestBuilder) Checkout()(*if911b3a10b9730297fcd733857f877651ed5a7489ee33fe66b62a1e4b4c29263.CheckoutRequestBuilder) {
    return if911b3a10b9730297fcd733857f877651ed5a7489ee33fe66b62a1e4b4c29263.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *DriveItemItemRequestBuilder) Children()(*i43351f457508aecc9acc29067e433332ebc46c9941ad1ac6e1a5f33f8fcb1e5d.ChildrenRequestBuilder) {
    return i43351f457508aecc9acc29067e433332ebc46c9941ad1ac6e1a5f33f8fcb1e5d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i73ce730b8b40e7e732bae3aa850e526852f4bbcfac2c62b18d895fb8c40cb028.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return i73ce730b8b40e7e732bae3aa850e526852f4bbcfac2c62b18d895fb8c40cb028.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive%2Did}/items/{driveItem%2Did}{?%24select,%24expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i5b0a99c5b67d423f45d52d7551bcbc2389478eaa8636e90fc04fde1312b36353.ContentRequestBuilder) {
    return i5b0a99c5b67d423f45d52d7551bcbc2389478eaa8636e90fc04fde1312b36353.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *DriveItemItemRequestBuilder) Copy()(*ib83a2fa775431a9271332b767a4c32b249184fbb6dd8404e43700f666927d598.CopyRequestBuilder) {
    return ib83a2fa775431a9271332b767a4c32b249184fbb6dd8404e43700f666927d598.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for me
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property items for me
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
func (m *DriveItemItemRequestBuilder) CreateLink()(*i2279753cbb4e1319aa20ac9a28757bc6d3b4d382bc665bbb1dc4c248f57d5980.CreateLinkRequestBuilder) {
    return i2279753cbb4e1319aa20ac9a28757bc6d3b4d382bc665bbb1dc4c248f57d5980.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in me
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property items in me
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
func (m *DriveItemItemRequestBuilder) CreateUploadSession()(*icf15246792ca783226cabc2cc3a1ceb518e142ebed900e2c7c1ebecc572b5964.CreateUploadSessionRequestBuilder) {
    return icf15246792ca783226cabc2cc3a1ceb518e142ebed900e2c7c1ebecc572b5964.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property items for me
func (m *DriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *DriveItemItemRequestBuilder) Delta()(*i3dca9f6bf310a5170317883bf7799e0ffa261a2b2542557f2280e1b46e7d54d9.DeltaRequestBuilder) {
    return i3dca9f6bf310a5170317883bf7799e0ffa261a2b2542557f2280e1b46e7d54d9.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *DriveItemItemRequestBuilder) DeltaWithToken(token *string)(*ibf93f730448d1b3d26834dbf6379c101e687bb7c2cc3cc1a3b8b6c8b2f0b3451.DeltaWithTokenRequestBuilder) {
    return ibf93f730448d1b3d26834dbf6379c101e687bb7c2cc3cc1a3b8b6c8b2f0b3451.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels the extractSensitivityLabels property
func (m *DriveItemItemRequestBuilder) ExtractSensitivityLabels()(*ie8ff27e9f76417e88fbdbaff20404d64201e0fb3256421b45e469dce450f5368.ExtractSensitivityLabelsRequestBuilder) {
    return ie8ff27e9f76417e88fbdbaff20404d64201e0fb3256421b45e469dce450f5368.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow the follow property
func (m *DriveItemItemRequestBuilder) Follow()(*i3575e154580c5d16893024f44b2f73ea3b26700f3cb7778adb00c19667046a3f.FollowRequestBuilder) {
    return i3575e154580c5d16893024f44b2f73ea3b26700f3cb7778adb00c19667046a3f.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *DriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*iad57dbed394ee51d4a6e621363c22bc0787cacfea764a2fb293f50016c5ec458.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return iad57dbed394ee51d4a6e621363c22bc0787cacfea764a2fb293f50016c5ec458.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite the invite property
func (m *DriveItemItemRequestBuilder) Invite()(*ic095af862cc8d74bb2e7b1e3a983f0498f1abc964e1622aa8e1e0df7b3a9e29d.InviteRequestBuilder) {
    return ic095af862cc8d74bb2e7b1e3a983f0498f1abc964e1622aa8e1e0df7b3a9e29d.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *DriveItemItemRequestBuilder) ListItem()(*i606f41c5fea6c70fd34ccde609068e6bd09206e372d836cc7a3cd47d661d3be8.ListItemRequestBuilder) {
    return i606f41c5fea6c70fd34ccde609068e6bd09206e372d836cc7a3cd47d661d3be8.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in me
func (m *DriveItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Permissions the permissions property
func (m *DriveItemItemRequestBuilder) Permissions()(*i5f7f4ee0090cc95d75d7c261819f4cc11af67ee4cf962b90e283ca54b09a795b.PermissionsRequestBuilder) {
    return i5f7f4ee0090cc95d75d7c261819f4cc11af67ee4cf962b90e283ca54b09a795b.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*ie01afcc24bcb8de9f524e134f72f107f5d9980714d4b3f83b8b0a3d66372a5c4.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return ie01afcc24bcb8de9f524e134f72f107f5d9980714d4b3f83b8b0a3d66372a5c4.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *DriveItemItemRequestBuilder) Preview()(*ib75035b590d4aa86c63fe0d7083badc6525d6470dcc873ca18ebe78c15b61131.PreviewRequestBuilder) {
    return ib75035b590d4aa86c63fe0d7083badc6525d6470dcc873ca18ebe78c15b61131.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *DriveItemItemRequestBuilder) Restore()(*ibe93737f97f952ab01899b8d2e868778cadcb47a868980aee7049140a22620d3.RestoreRequestBuilder) {
    return ibe93737f97f952ab01899b8d2e868778cadcb47a868980aee7049140a22620d3.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemItemRequestBuilder) SearchWithQ(q *string)(*i21e9fcf848cb9c39a2fc3f8761d9c40f31cbca3472bbe9a5cbcdbda4d175a8d9.SearchWithQRequestBuilder) {
    return i21e9fcf848cb9c39a2fc3f8761d9c40f31cbca3472bbe9a5cbcdbda4d175a8d9.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i52fdb84788de6f4295e35bd2351366c8c5b985d4d378bc90530d073b39223a0c.SubscriptionsRequestBuilder) {
    return i52fdb84788de6f4295e35bd2351366c8c5b985d4d378bc90530d073b39223a0c.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i8b02ac8460c4300fef83eeacf0e3a90fa6ad7e02a0ad5f87838abb6d5fb2ce82.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i8b02ac8460c4300fef83eeacf0e3a90fa6ad7e02a0ad5f87838abb6d5fb2ce82.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i39765dfa1c3eafd523949b6b410f599c0e93968da8276a68b09d6869ce0a7fc8.ThumbnailsRequestBuilder) {
    return i39765dfa1c3eafd523949b6b410f599c0e93968da8276a68b09d6869ce0a7fc8.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ie8883816406be65444fb3f248da99471dc749778a92112ecb2d943acde1ab3bc.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return ie8883816406be65444fb3f248da99471dc749778a92112ecb2d943acde1ab3bc.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *DriveItemItemRequestBuilder) Unfollow()(*i76e213ef3af16d5a017724c1129453b8e46917fd1758262825fd2504a111c798.UnfollowRequestBuilder) {
    return i76e213ef3af16d5a017724c1129453b8e46917fd1758262825fd2504a111c798.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *DriveItemItemRequestBuilder) ValidatePermission()(*ic1373904794125c73001b993b8b33c6016f56d96174b0260f5b8823ac4e00b57.ValidatePermissionRequestBuilder) {
    return ic1373904794125c73001b993b8b33c6016f56d96174b0260f5b8823ac4e00b57.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *DriveItemItemRequestBuilder) Versions()(*i6f3eef764ea1ae21b6822debfa6f6a704cd6435b66bcebeb13a675ec3815bd92.VersionsRequestBuilder) {
    return i6f3eef764ea1ae21b6822debfa6f6a704cd6435b66bcebeb13a675ec3815bd92.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ia7e12ef2cde2e1cee53c4ab5522f78d82906572a22510a1b2961ef668cc7a611.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return ia7e12ef2cde2e1cee53c4ab5522f78d82906572a22510a1b2961ef668cc7a611.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
