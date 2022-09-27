package root

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04bf9604718e0ae60b87e876a89e06624e62df01683c4f5c9d5888f0d85fa980 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/subscriptions"
    i05036fda23d42f66b0c931bf9c67aa77ca384d14e8dbe2c3f399234ec0e22f25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/thumbnails"
    i06c6ee8ec64989acb1a4c8f3cc3175ab9032615ae3ae8349a8a2c7988937e34d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/restore"
    i0afbd01b1126459376f38d5eb91442ba90b79e313fd2c07d9af0a5a617f2e90d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i30385fc07f753f4b85cdfb168df8603e7c08d523142ff6c77a5b413c1710435e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/validatepermission"
    i4286905b4bfca832e7d9c2015910cd6fb49ad2984c7de3f020589e726166b425 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/createuploadsession"
    i463bf2b6037ac8ed973cdc06b5b4e53a3edcb7174bd9bc94af25df6c54e91710 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/searchwithq"
    i4b3b3b26723de4cf0e2c0f82d3e74dbf1625a8d698b04dc8b3e351ea7bcd83af "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/createlink"
    i66954b603176ccff8af39f76b6aa2c8e6a4fa565a57aad07db99bfd34edb6820 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/invite"
    i67e4eb3c6d101d79ba06ed09a3844590642bdfdffbdee1ba9f29c99aff2eddea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/permissions"
    i6e388f09306deb8394779e5ed6b9a49627d1c6844c110c3b02368eeea4fa3252 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/deltawithtoken"
    i73d2513943183cb86cb44c927bf9ec731b074a759959852f6226167de5219d1a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/children"
    i7fac145e81d00bc3f8a046fdaeafa552b45b0dd728af40a99ea91384816ef4f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/unfollow"
    i85b42abbd4b054e587ffb69ce6d17cb83c885dc5d4f9f3c38a38c86b281f1f47 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/preview"
    i9c6e8a7fe1c9cc5b7e0ea6c10c1af6f88e7214ebfec50e64d9cecd6189cf62a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/analytics"
    ia25ae513c8a06c69f8e9a2af04dc1fdb56a445d26c487329094eb799760f235c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/listitem"
    ia783022e23d4a61455349d260caeaf7f827f148e0b229a292c2448b110933f0f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/activities"
    ib2a50031c8be26632ffac1f72d440d39608fce3dc0fc86ec7c84acbe8d8fbe43 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/assignsensitivitylabel"
    ib3375e0c67634a9e831a5be1208f27d49f468fbfd49b1b01ae47388c423cdae2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/delta"
    id52349d17c2abfaeae34ad4140c685947c41537424ad15ae45eab17d6b18b558 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/checkin"
    id5e391f2587d620e59fa3ee116774d47aa5530d731d5043f38e72e0e74a7a020 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/copy"
    id884479eaa004023d5bec8409bb89bb8dfc37c406b7fd09c2e61fed26ad48f6d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/follow"
    idad51f65dbf6409214a5262c790e55b2927ab950c32f2c9496643423250afbb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/extractsensitivitylabels"
    ifc35f41a199aa72c9e4dc64c26fc81bbcecfe30b0a508fcf312d679044a2a291 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/checkout"
    ifcde4c248d0fb16b7a3592bf94897136073d710d9aad05ea6da584f65e2631fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/content"
    iff014ae58e15af2c05d5f6a5712f20fbb33698c83bed3dfbac294a0e5ad0c6ca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/versions"
    i048cac4ee0d8d62b3c2b23511eb9186c6e2fdede9968406b07e88caf4df27ef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/activities/item"
    i6246ed490a551373dc53d9eb146e31953d1de9b86dc8fddc38ff9fd7cac5c8f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/versions/item"
    i794a610a5528f5f6d3c429be3e0f48213381e2de9f7f3c9f949ebf5497110518 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/permissions/item"
    i7e84b7d2d7188c19f7cdd6dc65d61b7bb5b964d0cca98d42f9cb67fe0ccedb77 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/subscriptions/item"
    id78a84f29c1764444a930c32f704f8900d71328fe66c44a8c731acd2089749ae "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/thumbnails/item"
    ie83cf697fa749a07f5a8e20f45c1986602567c1c297e5c1127d6e513588956a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/children/item"
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
func (m *RootRequestBuilder) Activities()(*ia783022e23d4a61455349d260caeaf7f827f148e0b229a292c2448b110933f0f.ActivitiesRequestBuilder) {
    return ia783022e23d4a61455349d260caeaf7f827f148e0b229a292c2448b110933f0f.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*i048cac4ee0d8d62b3c2b23511eb9186c6e2fdede9968406b07e88caf4df27ef6.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i048cac4ee0d8d62b3c2b23511eb9186c6e2fdede9968406b07e88caf4df27ef6.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*i9c6e8a7fe1c9cc5b7e0ea6c10c1af6f88e7214ebfec50e64d9cecd6189cf62a9.AnalyticsRequestBuilder) {
    return i9c6e8a7fe1c9cc5b7e0ea6c10c1af6f88e7214ebfec50e64d9cecd6189cf62a9.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel the assignSensitivityLabel property
func (m *RootRequestBuilder) AssignSensitivityLabel()(*ib2a50031c8be26632ffac1f72d440d39608fce3dc0fc86ec7c84acbe8d8fbe43.AssignSensitivityLabelRequestBuilder) {
    return ib2a50031c8be26632ffac1f72d440d39608fce3dc0fc86ec7c84acbe8d8fbe43.NewAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *RootRequestBuilder) Checkin()(*id52349d17c2abfaeae34ad4140c685947c41537424ad15ae45eab17d6b18b558.CheckinRequestBuilder) {
    return id52349d17c2abfaeae34ad4140c685947c41537424ad15ae45eab17d6b18b558.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *RootRequestBuilder) Checkout()(*ifc35f41a199aa72c9e4dc64c26fc81bbcecfe30b0a508fcf312d679044a2a291.CheckoutRequestBuilder) {
    return ifc35f41a199aa72c9e4dc64c26fc81bbcecfe30b0a508fcf312d679044a2a291.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i73d2513943183cb86cb44c927bf9ec731b074a759959852f6226167de5219d1a.ChildrenRequestBuilder) {
    return i73d2513943183cb86cb44c927bf9ec731b074a759959852f6226167de5219d1a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*ie83cf697fa749a07f5a8e20f45c1986602567c1c297e5c1127d6e513588956a7.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ie83cf697fa749a07f5a8e20f45c1986602567c1c297e5c1127d6e513588956a7.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive%2Did}/root{?%24select,%24expand}";
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
func (m *RootRequestBuilder) Content()(*ifcde4c248d0fb16b7a3592bf94897136073d710d9aad05ea6da584f65e2631fa.ContentRequestBuilder) {
    return ifcde4c248d0fb16b7a3592bf94897136073d710d9aad05ea6da584f65e2631fa.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *RootRequestBuilder) Copy()(*id5e391f2587d620e59fa3ee116774d47aa5530d731d5043f38e72e0e74a7a020.CopyRequestBuilder) {
    return id5e391f2587d620e59fa3ee116774d47aa5530d731d5043f38e72e0e74a7a020.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for me
func (m *RootRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for me
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
func (m *RootRequestBuilder) CreateLink()(*i4b3b3b26723de4cf0e2c0f82d3e74dbf1625a8d698b04dc8b3e351ea7bcd83af.CreateLinkRequestBuilder) {
    return i4b3b3b26723de4cf0e2c0f82d3e74dbf1625a8d698b04dc8b3e351ea7bcd83af.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property root in me
func (m *RootRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in me
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
func (m *RootRequestBuilder) CreateUploadSession()(*i4286905b4bfca832e7d9c2015910cd6fb49ad2984c7de3f020589e726166b425.CreateUploadSessionRequestBuilder) {
    return i4286905b4bfca832e7d9c2015910cd6fb49ad2984c7de3f020589e726166b425.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for me
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
func (m *RootRequestBuilder) Delta()(*ib3375e0c67634a9e831a5be1208f27d49f468fbfd49b1b01ae47388c423cdae2.DeltaRequestBuilder) {
    return ib3375e0c67634a9e831a5be1208f27d49f468fbfd49b1b01ae47388c423cdae2.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *RootRequestBuilder) DeltaWithToken(token *string)(*i6e388f09306deb8394779e5ed6b9a49627d1c6844c110c3b02368eeea4fa3252.DeltaWithTokenRequestBuilder) {
    return i6e388f09306deb8394779e5ed6b9a49627d1c6844c110c3b02368eeea4fa3252.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels the extractSensitivityLabels property
func (m *RootRequestBuilder) ExtractSensitivityLabels()(*idad51f65dbf6409214a5262c790e55b2927ab950c32f2c9496643423250afbb5.ExtractSensitivityLabelsRequestBuilder) {
    return idad51f65dbf6409214a5262c790e55b2927ab950c32f2c9496643423250afbb5.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow the follow property
func (m *RootRequestBuilder) Follow()(*id884479eaa004023d5bec8409bb89bb8dfc37c406b7fd09c2e61fed26ad48f6d.FollowRequestBuilder) {
    return id884479eaa004023d5bec8409bb89bb8dfc37c406b7fd09c2e61fed26ad48f6d.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *RootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i0afbd01b1126459376f38d5eb91442ba90b79e313fd2c07d9af0a5a617f2e90d.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i0afbd01b1126459376f38d5eb91442ba90b79e313fd2c07d9af0a5a617f2e90d.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite the invite property
func (m *RootRequestBuilder) Invite()(*i66954b603176ccff8af39f76b6aa2c8e6a4fa565a57aad07db99bfd34edb6820.InviteRequestBuilder) {
    return i66954b603176ccff8af39f76b6aa2c8e6a4fa565a57aad07db99bfd34edb6820.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *RootRequestBuilder) ListItem()(*ia25ae513c8a06c69f8e9a2af04dc1fdb56a445d26c487329094eb799760f235c.ListItemRequestBuilder) {
    return ia25ae513c8a06c69f8e9a2af04dc1fdb56a445d26c487329094eb799760f235c.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in me
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
func (m *RootRequestBuilder) Permissions()(*i67e4eb3c6d101d79ba06ed09a3844590642bdfdffbdee1ba9f29c99aff2eddea.PermissionsRequestBuilder) {
    return i67e4eb3c6d101d79ba06ed09a3844590642bdfdffbdee1ba9f29c99aff2eddea.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i794a610a5528f5f6d3c429be3e0f48213381e2de9f7f3c9f949ebf5497110518.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i794a610a5528f5f6d3c429be3e0f48213381e2de9f7f3c9f949ebf5497110518.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *RootRequestBuilder) Preview()(*i85b42abbd4b054e587ffb69ce6d17cb83c885dc5d4f9f3c38a38c86b281f1f47.PreviewRequestBuilder) {
    return i85b42abbd4b054e587ffb69ce6d17cb83c885dc5d4f9f3c38a38c86b281f1f47.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *RootRequestBuilder) Restore()(*i06c6ee8ec64989acb1a4c8f3cc3175ab9032615ae3ae8349a8a2c7988937e34d.RestoreRequestBuilder) {
    return i06c6ee8ec64989acb1a4c8f3cc3175ab9032615ae3ae8349a8a2c7988937e34d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *RootRequestBuilder) SearchWithQ(q *string)(*i463bf2b6037ac8ed973cdc06b5b4e53a3edcb7174bd9bc94af25df6c54e91710.SearchWithQRequestBuilder) {
    return i463bf2b6037ac8ed973cdc06b5b4e53a3edcb7174bd9bc94af25df6c54e91710.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*i04bf9604718e0ae60b87e876a89e06624e62df01683c4f5c9d5888f0d85fa980.SubscriptionsRequestBuilder) {
    return i04bf9604718e0ae60b87e876a89e06624e62df01683c4f5c9d5888f0d85fa980.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i7e84b7d2d7188c19f7cdd6dc65d61b7bb5b964d0cca98d42f9cb67fe0ccedb77.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i7e84b7d2d7188c19f7cdd6dc65d61b7bb5b964d0cca98d42f9cb67fe0ccedb77.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i05036fda23d42f66b0c931bf9c67aa77ca384d14e8dbe2c3f399234ec0e22f25.ThumbnailsRequestBuilder) {
    return i05036fda23d42f66b0c931bf9c67aa77ca384d14e8dbe2c3f399234ec0e22f25.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*id78a84f29c1764444a930c32f704f8900d71328fe66c44a8c731acd2089749ae.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return id78a84f29c1764444a930c32f704f8900d71328fe66c44a8c731acd2089749ae.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *RootRequestBuilder) Unfollow()(*i7fac145e81d00bc3f8a046fdaeafa552b45b0dd728af40a99ea91384816ef4f9.UnfollowRequestBuilder) {
    return i7fac145e81d00bc3f8a046fdaeafa552b45b0dd728af40a99ea91384816ef4f9.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *RootRequestBuilder) ValidatePermission()(*i30385fc07f753f4b85cdfb168df8603e7c08d523142ff6c77a5b413c1710435e.ValidatePermissionRequestBuilder) {
    return i30385fc07f753f4b85cdfb168df8603e7c08d523142ff6c77a5b413c1710435e.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*iff014ae58e15af2c05d5f6a5712f20fbb33698c83bed3dfbac294a0e5ad0c6ca.VersionsRequestBuilder) {
    return iff014ae58e15af2c05d5f6a5712f20fbb33698c83bed3dfbac294a0e5ad0c6ca.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i6246ed490a551373dc53d9eb146e31953d1de9b86dc8fddc38ff9fd7cac5c8f4.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i6246ed490a551373dc53d9eb146e31953d1de9b86dc8fddc38ff9fd7cac5c8f4.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
