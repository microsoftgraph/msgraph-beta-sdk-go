package root

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i05a211ac2ae44e6870193e603f0a31e541a2c25da2283a17fd0f213aa1fa50c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/delta"
    i09a126f4908f567485fc32bb158a1a4b20211b117ad08b9bfc7d26eb9f5451eb "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/content"
    i178794378d42ed55d9a9c4d65a1c4ab9a7ab74207d5d62e6adb5559738a4c284 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/copy"
    i184aef01445a07bacf7eb34262da7f276c260c40192ac823afb65cc4a7633f1a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/analytics"
    i265217f7ce26cff53b8799c28b035dca1d399639549db4ddd4e81e87f024981e "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/searchwithq"
    i272c0d28544c6b3914ed3fc93ba9b0d318c3df27974d42852f9fbf0ffa3c1ee4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/restore"
    i2b695595caf897cadb50aebc5a0f557a1cc4ad81052a1cac0c20769137b6c449 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/permissions"
    i665bd46e542584fbb1e64cb252507badc4e3d89280116beb0fcb383141639176 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/preview"
    i99dec5917f189f07b941f0af454e20d24e6a05de83be2ef9aa75bd0bb19ccfab "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/checkin"
    ia051d8653049bc5b4ed721f87ef74e6a6f98703294cae8c9cb26c3b75bb3800a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/extractsensitivitylabels"
    ia37b7adeb36356a67b48f2ea7350e35f5c25696a3527d14f974953d783db6caf "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/activities"
    iad9096207b89af0d541fff142c32b69ad17329661f64af59e5ec79321f8a805e "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/subscriptions"
    ib631e14e5c19c4b96df9154c3d72c540acde929be4dbc5294297164599670824 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/invite"
    ib695e935397298b0896bd5bea3f2317818e631b3dea7090cab84dbbadfddade9 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem"
    ib965af1bc5d047aa74c28788584d9b32704be43399f8313a3194d7052b965b12 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/deltawithtoken"
    ibe4c9dbbb4f55f58fde5a1c1a98e87d374d3f80083bfeb6dc4bfcad2b85c8247 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/children"
    ibf04211e5706a8fa2d6037c385f02c3728741c1bda1f14d4fdb076dc797038e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/thumbnails"
    ibf512cfc2c74f88666f6023f5100d3b29acc6a21e01d7156d19c1c47188cccb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/createuploadsession"
    id35c5c48e4cfcd45320d6e762c7f1349ee9a888e17734a8bdf11f4ef1829e600 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    idb9f67444070759a508e9262858eae1949bb39aa84a47051c7b959ddabec7cda "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/follow"
    ie93feb8f67d9498ed9f77ba2768a3e525e43366767c59b7b8d9d8bdbdafad113 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/checkout"
    iea718b4670a92f7fa645e61304ac21dd5e1759f54b3afa9d4ea12f9bbd3632be "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/versions"
    if01be42df7c6a9fc042d8e03026d635402a0c162088983a47b3e405c8e050327 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/createlink"
    if184c618a6b0a698758a9107b3c1dc43545c3a34a84207e6fa6aa397a4e45ab3 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/validatepermission"
    ifec12f84262dc5552c729f368d89953c143d715e56e87ff850654f8d48016dd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/unfollow"
    i3a6d4d1d922aca6539049acf1be4e4ceb74e33d1e4ef5fdbacc65fdfc3aa17a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/children/item"
    i81e065e135c5e13716fa269e31607dd05236975d7f5c49a9a00d5f54c15273e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/permissions/item"
    i8e59f342fa199bace9f2a29bd494ef5b5eef2ba7d4ffe16ffe63cc6571047eff "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/thumbnails/item"
    iafd51b75b6c2dfedbe569009d68f20b2a000ae6cbaf10b0895d4ce579d42bbe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/activities/item"
    ib6b7d756f8f85b8dd43fd8babecc7cacaa631bf86762453c1fa457a19e388b6b "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/versions/item"
    ifb9212e266c7381fa46df1045f3a949ee177f2cb7ddd4cd50628cbe0b09f0fa6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/subscriptions/item"
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
func (m *RootRequestBuilder) Activities()(*ia37b7adeb36356a67b48f2ea7350e35f5c25696a3527d14f974953d783db6caf.ActivitiesRequestBuilder) {
    return ia37b7adeb36356a67b48f2ea7350e35f5c25696a3527d14f974953d783db6caf.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*iafd51b75b6c2dfedbe569009d68f20b2a000ae6cbaf10b0895d4ce579d42bbe5.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return iafd51b75b6c2dfedbe569009d68f20b2a000ae6cbaf10b0895d4ce579d42bbe5.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*i184aef01445a07bacf7eb34262da7f276c260c40192ac823afb65cc4a7633f1a.AnalyticsRequestBuilder) {
    return i184aef01445a07bacf7eb34262da7f276c260c40192ac823afb65cc4a7633f1a.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *RootRequestBuilder) Checkin()(*i99dec5917f189f07b941f0af454e20d24e6a05de83be2ef9aa75bd0bb19ccfab.CheckinRequestBuilder) {
    return i99dec5917f189f07b941f0af454e20d24e6a05de83be2ef9aa75bd0bb19ccfab.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *RootRequestBuilder) Checkout()(*ie93feb8f67d9498ed9f77ba2768a3e525e43366767c59b7b8d9d8bdbdafad113.CheckoutRequestBuilder) {
    return ie93feb8f67d9498ed9f77ba2768a3e525e43366767c59b7b8d9d8bdbdafad113.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*ibe4c9dbbb4f55f58fde5a1c1a98e87d374d3f80083bfeb6dc4bfcad2b85c8247.ChildrenRequestBuilder) {
    return ibe4c9dbbb4f55f58fde5a1c1a98e87d374d3f80083bfeb6dc4bfcad2b85c8247.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i3a6d4d1d922aca6539049acf1be4e4ceb74e33d1e4ef5fdbacc65fdfc3aa17a4.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i3a6d4d1d922aca6539049acf1be4e4ceb74e33d1e4ef5fdbacc65fdfc3aa17a4.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/root{?%24select,%24expand}";
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
func (m *RootRequestBuilder) Content()(*i09a126f4908f567485fc32bb158a1a4b20211b117ad08b9bfc7d26eb9f5451eb.ContentRequestBuilder) {
    return i09a126f4908f567485fc32bb158a1a4b20211b117ad08b9bfc7d26eb9f5451eb.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *RootRequestBuilder) Copy()(*i178794378d42ed55d9a9c4d65a1c4ab9a7ab74207d5d62e6adb5559738a4c284.CopyRequestBuilder) {
    return i178794378d42ed55d9a9c4d65a1c4ab9a7ab74207d5d62e6adb5559738a4c284.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for drive
func (m *RootRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for drive
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
func (m *RootRequestBuilder) CreateLink()(*if01be42df7c6a9fc042d8e03026d635402a0c162088983a47b3e405c8e050327.CreateLinkRequestBuilder) {
    return if01be42df7c6a9fc042d8e03026d635402a0c162088983a47b3e405c8e050327.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property root in drive
func (m *RootRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in drive
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
func (m *RootRequestBuilder) CreateUploadSession()(*ibf512cfc2c74f88666f6023f5100d3b29acc6a21e01d7156d19c1c47188cccb6.CreateUploadSessionRequestBuilder) {
    return ibf512cfc2c74f88666f6023f5100d3b29acc6a21e01d7156d19c1c47188cccb6.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for drive
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
func (m *RootRequestBuilder) Delta()(*i05a211ac2ae44e6870193e603f0a31e541a2c25da2283a17fd0f213aa1fa50c9.DeltaRequestBuilder) {
    return i05a211ac2ae44e6870193e603f0a31e541a2c25da2283a17fd0f213aa1fa50c9.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *RootRequestBuilder) DeltaWithToken(token *string)(*ib965af1bc5d047aa74c28788584d9b32704be43399f8313a3194d7052b965b12.DeltaWithTokenRequestBuilder) {
    return ib965af1bc5d047aa74c28788584d9b32704be43399f8313a3194d7052b965b12.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels the extractSensitivityLabels property
func (m *RootRequestBuilder) ExtractSensitivityLabels()(*ia051d8653049bc5b4ed721f87ef74e6a6f98703294cae8c9cb26c3b75bb3800a.ExtractSensitivityLabelsRequestBuilder) {
    return ia051d8653049bc5b4ed721f87ef74e6a6f98703294cae8c9cb26c3b75bb3800a.NewExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow the follow property
func (m *RootRequestBuilder) Follow()(*idb9f67444070759a508e9262858eae1949bb39aa84a47051c7b959ddabec7cda.FollowRequestBuilder) {
    return idb9f67444070759a508e9262858eae1949bb39aa84a47051c7b959ddabec7cda.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the root folder of the drive. Read-only.
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
func (m *RootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*id35c5c48e4cfcd45320d6e762c7f1349ee9a888e17734a8bdf11f4ef1829e600.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return id35c5c48e4cfcd45320d6e762c7f1349ee9a888e17734a8bdf11f4ef1829e600.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite the invite property
func (m *RootRequestBuilder) Invite()(*ib631e14e5c19c4b96df9154c3d72c540acde929be4dbc5294297164599670824.InviteRequestBuilder) {
    return ib631e14e5c19c4b96df9154c3d72c540acde929be4dbc5294297164599670824.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *RootRequestBuilder) ListItem()(*ib695e935397298b0896bd5bea3f2317818e631b3dea7090cab84dbbadfddade9.ListItemRequestBuilder) {
    return ib695e935397298b0896bd5bea3f2317818e631b3dea7090cab84dbbadfddade9.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in drive
func (m *RootRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(error) {
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
func (m *RootRequestBuilder) Permissions()(*i2b695595caf897cadb50aebc5a0f557a1cc4ad81052a1cac0c20769137b6c449.PermissionsRequestBuilder) {
    return i2b695595caf897cadb50aebc5a0f557a1cc4ad81052a1cac0c20769137b6c449.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i81e065e135c5e13716fa269e31607dd05236975d7f5c49a9a00d5f54c15273e8.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i81e065e135c5e13716fa269e31607dd05236975d7f5c49a9a00d5f54c15273e8.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *RootRequestBuilder) Preview()(*i665bd46e542584fbb1e64cb252507badc4e3d89280116beb0fcb383141639176.PreviewRequestBuilder) {
    return i665bd46e542584fbb1e64cb252507badc4e3d89280116beb0fcb383141639176.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *RootRequestBuilder) Restore()(*i272c0d28544c6b3914ed3fc93ba9b0d318c3df27974d42852f9fbf0ffa3c1ee4.RestoreRequestBuilder) {
    return i272c0d28544c6b3914ed3fc93ba9b0d318c3df27974d42852f9fbf0ffa3c1ee4.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *RootRequestBuilder) SearchWithQ(q *string)(*i265217f7ce26cff53b8799c28b035dca1d399639549db4ddd4e81e87f024981e.SearchWithQRequestBuilder) {
    return i265217f7ce26cff53b8799c28b035dca1d399639549db4ddd4e81e87f024981e.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*iad9096207b89af0d541fff142c32b69ad17329661f64af59e5ec79321f8a805e.SubscriptionsRequestBuilder) {
    return iad9096207b89af0d541fff142c32b69ad17329661f64af59e5ec79321f8a805e.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*ifb9212e266c7381fa46df1045f3a949ee177f2cb7ddd4cd50628cbe0b09f0fa6.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return ifb9212e266c7381fa46df1045f3a949ee177f2cb7ddd4cd50628cbe0b09f0fa6.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*ibf04211e5706a8fa2d6037c385f02c3728741c1bda1f14d4fdb076dc797038e0.ThumbnailsRequestBuilder) {
    return ibf04211e5706a8fa2d6037c385f02c3728741c1bda1f14d4fdb076dc797038e0.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i8e59f342fa199bace9f2a29bd494ef5b5eef2ba7d4ffe16ffe63cc6571047eff.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i8e59f342fa199bace9f2a29bd494ef5b5eef2ba7d4ffe16ffe63cc6571047eff.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *RootRequestBuilder) Unfollow()(*ifec12f84262dc5552c729f368d89953c143d715e56e87ff850654f8d48016dd2.UnfollowRequestBuilder) {
    return ifec12f84262dc5552c729f368d89953c143d715e56e87ff850654f8d48016dd2.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *RootRequestBuilder) ValidatePermission()(*if184c618a6b0a698758a9107b3c1dc43545c3a34a84207e6fa6aa397a4e45ab3.ValidatePermissionRequestBuilder) {
    return if184c618a6b0a698758a9107b3c1dc43545c3a34a84207e6fa6aa397a4e45ab3.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*iea718b4670a92f7fa645e61304ac21dd5e1759f54b3afa9d4ea12f9bbd3632be.VersionsRequestBuilder) {
    return iea718b4670a92f7fa645e61304ac21dd5e1759f54b3afa9d4ea12f9bbd3632be.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*ib6b7d756f8f85b8dd43fd8babecc7cacaa631bf86762453c1fa457a19e388b6b.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return ib6b7d756f8f85b8dd43fd8babecc7cacaa631bf86762453c1fa457a19e388b6b.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
