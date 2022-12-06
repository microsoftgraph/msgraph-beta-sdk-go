package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0146b9e31c3cc4f2ab496bd8b7648bd7422072075b9c3fdb0b6223600352be8a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/children"
    i07fa0b9bcf1ebf1422ecf81f650c09280a537299bd61ea4a889e1779aed5cdb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/createlink"
    i0d7540bf7da4f5a327b935abc70409898a3fd6c1c9c792496574df321f76f182 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/createuploadsession"
    i12f8dcf6e8ea636350dee797d39120b2eff5eb823b749db83370bd17cde4ee26 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/deltawithtoken"
    i1559af655b2f07d6991e747733c0a5dbe16a79c9110b9cd4e95e89d1f533b436 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/analytics"
    i1dda96dcc6ca1d6bba90ca93c12a40d7bd6714ce98037fc360f32671277baa94 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i1ffc874dc11bc59b8f43e08f999b16d5537ed76c0d32ceacaf84e9d01a9a0a54 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/subscriptions"
    i286a92b61d7508a9e3bd4e1607d4dcca000c4c41be5af19c587bbc43e6d8f2e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/content"
    i3434cc42e81799ee683ff05c68279a9af935557da306a30871d4244ec69a586d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/copy"
    i3a42b727ca0a85dd5841ea0caf5e4ddcc0b92314677b92c30cdff812f9d18de6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/checkin"
    i4403a8b36ffdd31c6449e834ff7354472d5bcaf819a2c59390ccca08a944e4e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/thumbnails"
    i62c92f7e8ef283c94ff015c3fe5d9d7f09ae77403203082ac6544c4f367d9676 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/follow"
    i652ad18530a0d67db6af9e17a408a0fe46a970bfe04431202d422c3171c97065 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/restore"
    i68453248b6adc2e0a123184b9f0397803b74547d93a538bca62fb45840ec7c39 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities"
    i69d7ab1245b51d446b0639d6cff194b64c1ff6ee87304184b6c17dd90f5c093c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem"
    i710d1b03133218cf42db16a0f513411325c8307a829644968e565a02e56ae7f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/invite"
    i7ecb3bb901e726c93053e2981950ac4cad1568eec1a8ca7051a9125074a2f106 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/searchwithq"
    i8236d13341080e5a3a7f2865e569b3a9db26acaf5e7254cd2dfb7594f29e3b31 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/validatepermission"
    i8bf71418a40d55c88bd463ae8678f7a80a256e281726ddf4e9c5b367b472be7a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/extractsensitivitylabels"
    i9e601979c83767d12347ff319a26aa16bd8cde9f1cf30b7a89c1c15121909250 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/preview"
    ia10d6bc417675919b129deccbb4a8ff6d0e423f7280ed162a4c41f59cc0e8de4 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/permissions"
    ia379a18993382e2cd165de4f6dc66e23159c7edce7c85491eb79d8d5a828f2bb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/assignsensitivitylabel"
    ib109f1365e8ac7f6b077a01adab1935b4dd6adc390c9e072a8b9a8ef2653e116 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/checkout"
    ibcde7b93c66be3ae4c6307030c9c1fd4912a9617a23339afbb10615f9a735dd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/unfollow"
    iccf725507e0de7b68f548524e89423d3c820bc9d1292124875957bfd3d78a358 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/delta"
    iea15fd2e4983cd4bc7f77c56f9e7d10b714e2d4c76c191417ae77f325254d1fd "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/versions"
    i07d08684839c1c52e4e2f97c0ac854b38d29ec3cfe7f59122208e108260cbb4d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/children/item"
    i23745af441bc7e806a6ea374d93773d7f5b4d51f0b23dbfc2e289f9ed360c56f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/versions/item"
    i8af00cb19bae1d8d547aa285fc2e34cbe5c5c5096ff8de885009b1dcd3611d71 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/permissions/item"
    ib3b0b0b8db03e078a58d5fd2c684b9c2f06f5e77dde80b3c46359e8e3f0480ba "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item"
    ic2204a9b8bb31b272b46aa5d8e8e1802c6b262a7377a2425b7aeb9b634548a33 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/subscriptions/item"
    if9d9f1bd841c39df4ae58bdaf1c987671d7244b7fbc793a0e55c2c3bcb60d75b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/thumbnails/item"
)

// DriveItemItemRequestBuilder provides operations to manage the collection of driveItem entities.
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
// DriveItemItemRequestBuilderGetQueryParameters get entity from workbooks by key
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
func (m *DriveItemItemRequestBuilder) Activities()(*i68453248b6adc2e0a123184b9f0397803b74547d93a538bca62fb45840ec7c39.ActivitiesRequestBuilder) {
    return i68453248b6adc2e0a123184b9f0397803b74547d93a538bca62fb45840ec7c39.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*ib3b0b0b8db03e078a58d5fd2c684b9c2f06f5e77dde80b3c46359e8e3f0480ba.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ib3b0b0b8db03e078a58d5fd2c684b9c2f06f5e77dde80b3c46359e8e3f0480ba.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Analytics()(*i1559af655b2f07d6991e747733c0a5dbe16a79c9110b9cd4e95e89d1f533b436.AnalyticsRequestBuilder) {
    return i1559af655b2f07d6991e747733c0a5dbe16a79c9110b9cd4e95e89d1f533b436.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
func (m *DriveItemItemRequestBuilder) AssignSensitivityLabel()(*ia379a18993382e2cd165de4f6dc66e23159c7edce7c85491eb79d8d5a828f2bb.AssignSensitivityLabelRequestBuilder) {
    return ia379a18993382e2cd165de4f6dc66e23159c7edce7c85491eb79d8d5a828f2bb.NewAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin provides operations to call the checkin method.
func (m *DriveItemItemRequestBuilder) Checkin()(*i3a42b727ca0a85dd5841ea0caf5e4ddcc0b92314677b92c30cdff812f9d18de6.CheckinRequestBuilder) {
    return i3a42b727ca0a85dd5841ea0caf5e4ddcc0b92314677b92c30cdff812f9d18de6.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout provides operations to call the checkout method.
func (m *DriveItemItemRequestBuilder) Checkout()(*ib109f1365e8ac7f6b077a01adab1935b4dd6adc390c9e072a8b9a8ef2653e116.CheckoutRequestBuilder) {
    return ib109f1365e8ac7f6b077a01adab1935b4dd6adc390c9e072a8b9a8ef2653e116.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Children()(*i0146b9e31c3cc4f2ab496bd8b7648bd7422072075b9c3fdb0b6223600352be8a.ChildrenRequestBuilder) {
    return i0146b9e31c3cc4f2ab496bd8b7648bd7422072075b9c3fdb0b6223600352be8a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i07d08684839c1c52e4e2f97c0ac854b38d29ec3cfe7f59122208e108260cbb4d.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return i07d08684839c1c52e4e2f97c0ac854b38d29ec3cfe7f59122208e108260cbb4d.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem%2Did}{?%24select,%24expand}";
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
// Content provides operations to manage the media for the driveItem entity.
func (m *DriveItemItemRequestBuilder) Content()(*i286a92b61d7508a9e3bd4e1607d4dcca000c4c41be5af19c587bbc43e6d8f2e5.ContentRequestBuilder) {
    return i286a92b61d7508a9e3bd4e1607d4dcca000c4c41be5af19c587bbc43e6d8f2e5.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *DriveItemItemRequestBuilder) Copy()(*i3434cc42e81799ee683ff05c68279a9af935557da306a30871d4244ec69a586d.CopyRequestBuilder) {
    return i3434cc42e81799ee683ff05c68279a9af935557da306a30871d4244ec69a586d.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete entity from workbooks
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
// CreateGetRequestInformation get entity from workbooks by key
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
func (m *DriveItemItemRequestBuilder) CreateLink()(*i07fa0b9bcf1ebf1422ecf81f650c09280a537299bd61ea4a889e1779aed5cdb1.CreateLinkRequestBuilder) {
    return i07fa0b9bcf1ebf1422ecf81f650c09280a537299bd61ea4a889e1779aed5cdb1.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update entity in workbooks
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
func (m *DriveItemItemRequestBuilder) CreateUploadSession()(*i0d7540bf7da4f5a327b935abc70409898a3fd6c1c9c792496574df321f76f182.CreateUploadSessionRequestBuilder) {
    return i0d7540bf7da4f5a327b935abc70409898a3fd6c1c9c792496574df321f76f182.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete entity from workbooks
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
func (m *DriveItemItemRequestBuilder) Delta()(*iccf725507e0de7b68f548524e89423d3c820bc9d1292124875957bfd3d78a358.DeltaRequestBuilder) {
    return iccf725507e0de7b68f548524e89423d3c820bc9d1292124875957bfd3d78a358.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *DriveItemItemRequestBuilder) DeltaWithToken(token *string)(*i12f8dcf6e8ea636350dee797d39120b2eff5eb823b749db83370bd17cde4ee26.DeltaWithTokenRequestBuilder) {
    return i12f8dcf6e8ea636350dee797d39120b2eff5eb823b749db83370bd17cde4ee26.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
func (m *DriveItemItemRequestBuilder) ExtractSensitivityLabels()(*i8bf71418a40d55c88bd463ae8678f7a80a256e281726ddf4e9c5b367b472be7a.ExtractSensitivityLabelsRequestBuilder) {
    return i8bf71418a40d55c88bd463ae8678f7a80a256e281726ddf4e9c5b367b472be7a.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow provides operations to call the follow method.
func (m *DriveItemItemRequestBuilder) Follow()(*i62c92f7e8ef283c94ff015c3fe5d9d7f09ae77403203082ac6544c4f367d9676.FollowRequestBuilder) {
    return i62c92f7e8ef283c94ff015c3fe5d9d7f09ae77403203082ac6544c4f367d9676.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from workbooks by key
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
func (m *DriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i1dda96dcc6ca1d6bba90ca93c12a40d7bd6714ce98037fc360f32671277baa94.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i1dda96dcc6ca1d6bba90ca93c12a40d7bd6714ce98037fc360f32671277baa94.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite provides operations to call the invite method.
func (m *DriveItemItemRequestBuilder) Invite()(*i710d1b03133218cf42db16a0f513411325c8307a829644968e565a02e56ae7f6.InviteRequestBuilder) {
    return i710d1b03133218cf42db16a0f513411325c8307a829644968e565a02e56ae7f6.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) ListItem()(*i69d7ab1245b51d446b0639d6cff194b64c1ff6ee87304184b6c17dd90f5c093c.ListItemRequestBuilder) {
    return i69d7ab1245b51d446b0639d6cff194b64c1ff6ee87304184b6c17dd90f5c093c.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in workbooks
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
func (m *DriveItemItemRequestBuilder) Permissions()(*ia10d6bc417675919b129deccbb4a8ff6d0e423f7280ed162a4c41f59cc0e8de4.PermissionsRequestBuilder) {
    return ia10d6bc417675919b129deccbb4a8ff6d0e423f7280ed162a4c41f59cc0e8de4.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i8af00cb19bae1d8d547aa285fc2e34cbe5c5c5096ff8de885009b1dcd3611d71.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i8af00cb19bae1d8d547aa285fc2e34cbe5c5c5096ff8de885009b1dcd3611d71.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview provides operations to call the preview method.
func (m *DriveItemItemRequestBuilder) Preview()(*i9e601979c83767d12347ff319a26aa16bd8cde9f1cf30b7a89c1c15121909250.PreviewRequestBuilder) {
    return i9e601979c83767d12347ff319a26aa16bd8cde9f1cf30b7a89c1c15121909250.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *DriveItemItemRequestBuilder) Restore()(*i652ad18530a0d67db6af9e17a408a0fe46a970bfe04431202d422c3171c97065.RestoreRequestBuilder) {
    return i652ad18530a0d67db6af9e17a408a0fe46a970bfe04431202d422c3171c97065.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemItemRequestBuilder) SearchWithQ(q *string)(*i7ecb3bb901e726c93053e2981950ac4cad1568eec1a8ca7051a9125074a2f106.SearchWithQRequestBuilder) {
    return i7ecb3bb901e726c93053e2981950ac4cad1568eec1a8ca7051a9125074a2f106.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i1ffc874dc11bc59b8f43e08f999b16d5537ed76c0d32ceacaf84e9d01a9a0a54.SubscriptionsRequestBuilder) {
    return i1ffc874dc11bc59b8f43e08f999b16d5537ed76c0d32ceacaf84e9d01a9a0a54.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*ic2204a9b8bb31b272b46aa5d8e8e1802c6b262a7377a2425b7aeb9b634548a33.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return ic2204a9b8bb31b272b46aa5d8e8e1802c6b262a7377a2425b7aeb9b634548a33.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i4403a8b36ffdd31c6449e834ff7354472d5bcaf819a2c59390ccca08a944e4e0.ThumbnailsRequestBuilder) {
    return i4403a8b36ffdd31c6449e834ff7354472d5bcaf819a2c59390ccca08a944e4e0.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*if9d9f1bd841c39df4ae58bdaf1c987671d7244b7fbc793a0e55c2c3bcb60d75b.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return if9d9f1bd841c39df4ae58bdaf1c987671d7244b7fbc793a0e55c2c3bcb60d75b.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow provides operations to call the unfollow method.
func (m *DriveItemItemRequestBuilder) Unfollow()(*ibcde7b93c66be3ae4c6307030c9c1fd4912a9617a23339afbb10615f9a735dd3.UnfollowRequestBuilder) {
    return ibcde7b93c66be3ae4c6307030c9c1fd4912a9617a23339afbb10615f9a735dd3.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *DriveItemItemRequestBuilder) ValidatePermission()(*i8236d13341080e5a3a7f2865e569b3a9db26acaf5e7254cd2dfb7594f29e3b31.ValidatePermissionRequestBuilder) {
    return i8236d13341080e5a3a7f2865e569b3a9db26acaf5e7254cd2dfb7594f29e3b31.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) Versions()(*iea15fd2e4983cd4bc7f77c56f9e7d10b714e2d4c76c191417ae77f325254d1fd.VersionsRequestBuilder) {
    return iea15fd2e4983cd4bc7f77c56f9e7d10b714e2d4c76c191417ae77f325254d1fd.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i23745af441bc7e806a6ea374d93773d7f5b4d51f0b23dbfc2e289f9ed360c56f.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i23745af441bc7e806a6ea374d93773d7f5b4d51f0b23dbfc2e289f9ed360c56f.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
