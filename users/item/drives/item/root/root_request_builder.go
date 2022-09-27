package root

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i093a79137daeb9ae088062ba13cde869f70e91fa581db36ec919e460c4ff49ec "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/validatepermission"
    i0ff1ae699bc68e404d3d3aab6d6cacecbb1340f0ec95eb4c2db97c606fc95ce3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/createlink"
    i348cfa40d80f2bd346049a604b61e3bbf99a86eaaac1524d02c1cc02ffd86851 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/permissions"
    i4b7ee927e3c2dc792ca09cddc3d0981a89253866d9f177e70356375494e5bd07 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/activities"
    i50e5dbbc7adce54d30748bc54c16ad058df719691e3dd262034b95e19684d59d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/analytics"
    i521b27574ef6b5fe4a6bfd7b853c014c561e97c0cf1244f787e95138a1e57d39 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/invite"
    i658e4b44555250a3d896740bfd9a82dcd8af8ce52910c3bc230569041605a43d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/subscriptions"
    i6c6dcdf98e4d0069fa4979b0f011a559e1255d805e8ca9c12c7b2c8ae4ce66f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/assignsensitivitylabel"
    i786356f560a4a8d426056e8340ecb57ddae9d12d69890ac34fb69a73b8469e38 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/checkin"
    i7b7ef26a39790f7760efc19094f8a199de1cdbd8ffb8f10d4309d54d9d571530 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/thumbnails"
    i8f75eaec56c7ab5659de30f31953e0926f050150805e53382e5637b9579e4ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/listitem"
    ia940bf4537c9d33ed7aa53527ac61586e93207869cea2aabb22432e33816f7f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/extractsensitivitylabels"
    iac52cab0ab0db188a977bc6178beb827af0c6e099c6a5593a0503e97ca41c6a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/content"
    iaf7ea098efa9f34e9359f14fd6c49319154abe7ce01aeface9f41b6c1c45ee06 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/delta"
    ib5242641354f8b0c9538dc8a05a47b72c3c4178ffe6922887a00dcfc6004095b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/preview"
    ib780e6381ef2d81efabd373745325e68fdd0ac722022578f0aa585b2eb4b00f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/follow"
    ibbc735c44cd014f89146ec819a553ab21c81f7dbe3601079a44dc3f7ebbcfd36 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/searchwithq"
    ibff0b4d33209b315966bf59f99fe1478d7383d45cbb30624ed8c4be22ddcff85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/children"
    ic8991ce87c50af1fdc1187aca9a2c6acd2b2455433a9d39b29e2dc16cde46841 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/checkout"
    ice0ad55746dc70f1efd68c3351d89a1e35dc9a9c58400789e60a62350ec73187 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/createuploadsession"
    id0a597ca323d0506ecf70966fdba5a96909866d7e055edbb39b5c49f62163176 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    id16fa65ecd844813064c510a149b0aa2ff07a8365cbf4662c4c47e63e886b026 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/deltawithtoken"
    id6956dd05a2e0b17e1dfc6210de6a3bb1847b08af5b5da0e9a0e5c799556c506 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/restore"
    id8d5a319b88e388b070ce15327e1b6ba8751257d8c613f3e363927fa1df9b30b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/versions"
    idf8f03936462ae9905c25d598e62b0e5e6f0dbef87b4583d760dd607c490190a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/unfollow"
    ief82b7badd289ac6b0ea22e27535261ae9607a67cc8aec6946595dc4fe3a8749 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/copy"
    i470c54a9bee9a8ac27ba6d46057802b98035f174bfac53a07d4fb571fa8b4fa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/thumbnails/item"
    i4873d5650fb447ccd691622a251102eac60f90f73cee1bc40ea7b8a45fe10739 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/activities/item"
    i7c223ca62e6930463b88db6f01fae288e09d61d0e65f94e1bf9b2d7560c3a068 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/children/item"
    i7d632e25daa2136a74aff02a73d2cfc6ab090eca471ab318343c51239e330360 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/versions/item"
    i8aa83421f1959464fd1940ca85cb28228718e593552c3ae1909ed33b6cf13631 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/subscriptions/item"
    ibf85f2aba96b5b04085e8e4e6bc260e595f161ebb0aa75654b03be699f981bd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/permissions/item"
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
// Activities the activities property
func (m *RootRequestBuilder) Activities()(*i4b7ee927e3c2dc792ca09cddc3d0981a89253866d9f177e70356375494e5bd07.ActivitiesRequestBuilder) {
    return i4b7ee927e3c2dc792ca09cddc3d0981a89253866d9f177e70356375494e5bd07.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*i4873d5650fb447ccd691622a251102eac60f90f73cee1bc40ea7b8a45fe10739.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i4873d5650fb447ccd691622a251102eac60f90f73cee1bc40ea7b8a45fe10739.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*i50e5dbbc7adce54d30748bc54c16ad058df719691e3dd262034b95e19684d59d.AnalyticsRequestBuilder) {
    return i50e5dbbc7adce54d30748bc54c16ad058df719691e3dd262034b95e19684d59d.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel the assignSensitivityLabel property
func (m *RootRequestBuilder) AssignSensitivityLabel()(*i6c6dcdf98e4d0069fa4979b0f011a559e1255d805e8ca9c12c7b2c8ae4ce66f7.AssignSensitivityLabelRequestBuilder) {
    return i6c6dcdf98e4d0069fa4979b0f011a559e1255d805e8ca9c12c7b2c8ae4ce66f7.NewAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *RootRequestBuilder) Checkin()(*i786356f560a4a8d426056e8340ecb57ddae9d12d69890ac34fb69a73b8469e38.CheckinRequestBuilder) {
    return i786356f560a4a8d426056e8340ecb57ddae9d12d69890ac34fb69a73b8469e38.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *RootRequestBuilder) Checkout()(*ic8991ce87c50af1fdc1187aca9a2c6acd2b2455433a9d39b29e2dc16cde46841.CheckoutRequestBuilder) {
    return ic8991ce87c50af1fdc1187aca9a2c6acd2b2455433a9d39b29e2dc16cde46841.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*ibff0b4d33209b315966bf59f99fe1478d7383d45cbb30624ed8c4be22ddcff85.ChildrenRequestBuilder) {
    return ibff0b4d33209b315966bf59f99fe1478d7383d45cbb30624ed8c4be22ddcff85.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i7c223ca62e6930463b88db6f01fae288e09d61d0e65f94e1bf9b2d7560c3a068.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i7c223ca62e6930463b88db6f01fae288e09d61d0e65f94e1bf9b2d7560c3a068.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/root{?%24select,%24expand}";
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
func (m *RootRequestBuilder) Content()(*iac52cab0ab0db188a977bc6178beb827af0c6e099c6a5593a0503e97ca41c6a9.ContentRequestBuilder) {
    return iac52cab0ab0db188a977bc6178beb827af0c6e099c6a5593a0503e97ca41c6a9.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *RootRequestBuilder) Copy()(*ief82b7badd289ac6b0ea22e27535261ae9607a67cc8aec6946595dc4fe3a8749.CopyRequestBuilder) {
    return ief82b7badd289ac6b0ea22e27535261ae9607a67cc8aec6946595dc4fe3a8749.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for users
func (m *RootRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for users
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
// CreateGetRequestInformation retrieve the metadata for a driveItem in a drive by file system path or ID.
func (m *RootRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieve the metadata for a driveItem in a drive by file system path or ID.
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
func (m *RootRequestBuilder) CreateLink()(*i0ff1ae699bc68e404d3d3aab6d6cacecbb1340f0ec95eb4c2db97c606fc95ce3.CreateLinkRequestBuilder) {
    return i0ff1ae699bc68e404d3d3aab6d6cacecbb1340f0ec95eb4c2db97c606fc95ce3.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property root in users
func (m *RootRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in users
func (m *RootRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateUploadSession the createUploadSession property
func (m *RootRequestBuilder) CreateUploadSession()(*ice0ad55746dc70f1efd68c3351d89a1e35dc9a9c58400789e60a62350ec73187.CreateUploadSessionRequestBuilder) {
    return ice0ad55746dc70f1efd68c3351d89a1e35dc9a9c58400789e60a62350ec73187.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for users
func (m *RootRequestBuilder) Delete(ctx context.Context, requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *RootRequestBuilder) Delta()(*iaf7ea098efa9f34e9359f14fd6c49319154abe7ce01aeface9f41b6c1c45ee06.DeltaRequestBuilder) {
    return iaf7ea098efa9f34e9359f14fd6c49319154abe7ce01aeface9f41b6c1c45ee06.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *RootRequestBuilder) DeltaWithToken(token *string)(*id16fa65ecd844813064c510a149b0aa2ff07a8365cbf4662c4c47e63e886b026.DeltaWithTokenRequestBuilder) {
    return id16fa65ecd844813064c510a149b0aa2ff07a8365cbf4662c4c47e63e886b026.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels the extractSensitivityLabels property
func (m *RootRequestBuilder) ExtractSensitivityLabels()(*ia940bf4537c9d33ed7aa53527ac61586e93207869cea2aabb22432e33816f7f1.ExtractSensitivityLabelsRequestBuilder) {
    return ia940bf4537c9d33ed7aa53527ac61586e93207869cea2aabb22432e33816f7f1.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow the follow property
func (m *RootRequestBuilder) Follow()(*ib780e6381ef2d81efabd373745325e68fdd0ac722022578f0aa585b2eb4b00f1.FollowRequestBuilder) {
    return ib780e6381ef2d81efabd373745325e68fdd0ac722022578f0aa585b2eb4b00f1.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the metadata for a driveItem in a drive by file system path or ID.
func (m *RootRequestBuilder) Get(ctx context.Context, requestConfiguration *RootRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *RootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*id0a597ca323d0506ecf70966fdba5a96909866d7e055edbb39b5c49f62163176.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return id0a597ca323d0506ecf70966fdba5a96909866d7e055edbb39b5c49f62163176.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite the invite property
func (m *RootRequestBuilder) Invite()(*i521b27574ef6b5fe4a6bfd7b853c014c561e97c0cf1244f787e95138a1e57d39.InviteRequestBuilder) {
    return i521b27574ef6b5fe4a6bfd7b853c014c561e97c0cf1244f787e95138a1e57d39.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *RootRequestBuilder) ListItem()(*i8f75eaec56c7ab5659de30f31953e0926f050150805e53382e5637b9579e4ea1.ListItemRequestBuilder) {
    return i8f75eaec56c7ab5659de30f31953e0926f050150805e53382e5637b9579e4ea1.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in users
func (m *RootRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
func (m *RootRequestBuilder) Permissions()(*i348cfa40d80f2bd346049a604b61e3bbf99a86eaaac1524d02c1cc02ffd86851.PermissionsRequestBuilder) {
    return i348cfa40d80f2bd346049a604b61e3bbf99a86eaaac1524d02c1cc02ffd86851.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*ibf85f2aba96b5b04085e8e4e6bc260e595f161ebb0aa75654b03be699f981bd6.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return ibf85f2aba96b5b04085e8e4e6bc260e595f161ebb0aa75654b03be699f981bd6.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *RootRequestBuilder) Preview()(*ib5242641354f8b0c9538dc8a05a47b72c3c4178ffe6922887a00dcfc6004095b.PreviewRequestBuilder) {
    return ib5242641354f8b0c9538dc8a05a47b72c3c4178ffe6922887a00dcfc6004095b.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *RootRequestBuilder) Restore()(*id6956dd05a2e0b17e1dfc6210de6a3bb1847b08af5b5da0e9a0e5c799556c506.RestoreRequestBuilder) {
    return id6956dd05a2e0b17e1dfc6210de6a3bb1847b08af5b5da0e9a0e5c799556c506.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *RootRequestBuilder) SearchWithQ(q *string)(*ibbc735c44cd014f89146ec819a553ab21c81f7dbe3601079a44dc3f7ebbcfd36.SearchWithQRequestBuilder) {
    return ibbc735c44cd014f89146ec819a553ab21c81f7dbe3601079a44dc3f7ebbcfd36.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*i658e4b44555250a3d896740bfd9a82dcd8af8ce52910c3bc230569041605a43d.SubscriptionsRequestBuilder) {
    return i658e4b44555250a3d896740bfd9a82dcd8af8ce52910c3bc230569041605a43d.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i8aa83421f1959464fd1940ca85cb28228718e593552c3ae1909ed33b6cf13631.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i8aa83421f1959464fd1940ca85cb28228718e593552c3ae1909ed33b6cf13631.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i7b7ef26a39790f7760efc19094f8a199de1cdbd8ffb8f10d4309d54d9d571530.ThumbnailsRequestBuilder) {
    return i7b7ef26a39790f7760efc19094f8a199de1cdbd8ffb8f10d4309d54d9d571530.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i470c54a9bee9a8ac27ba6d46057802b98035f174bfac53a07d4fb571fa8b4fa9.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i470c54a9bee9a8ac27ba6d46057802b98035f174bfac53a07d4fb571fa8b4fa9.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *RootRequestBuilder) Unfollow()(*idf8f03936462ae9905c25d598e62b0e5e6f0dbef87b4583d760dd607c490190a.UnfollowRequestBuilder) {
    return idf8f03936462ae9905c25d598e62b0e5e6f0dbef87b4583d760dd607c490190a.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *RootRequestBuilder) ValidatePermission()(*i093a79137daeb9ae088062ba13cde869f70e91fa581db36ec919e460c4ff49ec.ValidatePermissionRequestBuilder) {
    return i093a79137daeb9ae088062ba13cde869f70e91fa581db36ec919e460c4ff49ec.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*id8d5a319b88e388b070ce15327e1b6ba8751257d8c613f3e363927fa1df9b30b.VersionsRequestBuilder) {
    return id8d5a319b88e388b070ce15327e1b6ba8751257d8c613f3e363927fa1df9b30b.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i7d632e25daa2136a74aff02a73d2cfc6ab090eca471ab318343c51239e330360.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i7d632e25daa2136a74aff02a73d2cfc6ab090eca471ab318343c51239e330360.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
