package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
    i428d2c5ba8bcc9f861cfc7a3060a1e340984c54ca28a655b2305e6efffd7a1db "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item/removemembers"
    i467add2156077b3c847fe4ad47f96c896d245f282bc0b3f06714966be14fec8e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item/removemembersbyid"
    id908467d80ca881966b187f3af525715c1cf4542c77a3c376e78ab4a4d40d021 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item/addmembersbyid"
    ie0663ce56074dacfdac086ea72d5e5de021ac050bb6e56380a99905fa9551b9c "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item/addmembers"
)

// UpdatableAssetItemRequestBuilder provides operations to manage the updatableAssets property of the microsoft.graph.windowsUpdates.updates entity.
type UpdatableAssetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdatableAssetItemRequestBuilderDeleteOptions options for Delete
type UpdatableAssetItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UpdatableAssetItemRequestBuilderGetOptions options for Get
type UpdatableAssetItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UpdatableAssetItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UpdatableAssetItemRequestBuilderGetQueryParameters assets registered with the deployment service that can receive updates. Read-only.
type UpdatableAssetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UpdatableAssetItemRequestBuilderPatchOptions options for Patch
type UpdatableAssetItemRequestBuilderPatchOptions struct {
    // 
    Body ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UpdatableAssetItemRequestBuilder) AddMembers()(*ie0663ce56074dacfdac086ea72d5e5de021ac050bb6e56380a99905fa9551b9c.AddMembersRequestBuilder) {
    return ie0663ce56074dacfdac086ea72d5e5de021ac050bb6e56380a99905fa9551b9c.NewAddMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetItemRequestBuilder) AddMembersById()(*id908467d80ca881966b187f3af525715c1cf4542c77a3c376e78ab4a4d40d021.AddMembersByIdRequestBuilder) {
    return id908467d80ca881966b187f3af525715c1cf4542c77a3c376e78ab4a4d40d021.NewAddMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUpdatableAssetItemRequestBuilderInternal instantiates a new UpdatableAssetItemRequestBuilder and sets the default values.
func NewUpdatableAssetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetItemRequestBuilder) {
    m := &UpdatableAssetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/updatableAssets/{updatableAsset_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdatableAssetItemRequestBuilder instantiates a new UpdatableAssetItemRequestBuilder and sets the default values.
func NewUpdatableAssetItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatableAssetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property updatableAssets for admin
func (m *UpdatableAssetItemRequestBuilder) CreateDeleteRequestInformation(options *UpdatableAssetItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetItemRequestBuilder) CreateGetRequestInformation(options *UpdatableAssetItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property updatableAssets in admin
func (m *UpdatableAssetItemRequestBuilder) CreatePatchRequestInformation(options *UpdatableAssetItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property updatableAssets for admin
func (m *UpdatableAssetItemRequestBuilder) Delete(options *UpdatableAssetItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetItemRequestBuilder) Get(options *UpdatableAssetItemRequestBuilderGetOptions)(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.CreateUpdatableAssetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable), nil
}
// Patch update the navigation property updatableAssets in admin
func (m *UpdatableAssetItemRequestBuilder) Patch(options *UpdatableAssetItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UpdatableAssetItemRequestBuilder) RemoveMembers()(*i428d2c5ba8bcc9f861cfc7a3060a1e340984c54ca28a655b2305e6efffd7a1db.RemoveMembersRequestBuilder) {
    return i428d2c5ba8bcc9f861cfc7a3060a1e340984c54ca28a655b2305e6efffd7a1db.NewRemoveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetItemRequestBuilder) RemoveMembersById()(*i467add2156077b3c847fe4ad47f96c896d245f282bc0b3f06714966be14fec8e.RemoveMembersByIdRequestBuilder) {
    return i467add2156077b3c847fe4ad47f96c896d245f282bc0b3f06714966be14fec8e.NewRemoveMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
