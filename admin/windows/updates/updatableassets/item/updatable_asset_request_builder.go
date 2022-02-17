package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428d2c5ba8bcc9f861cfc7a3060a1e340984c54ca28a655b2305e6efffd7a1db "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item/removemembers"
    i467add2156077b3c847fe4ad47f96c896d245f282bc0b3f06714966be14fec8e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item/removemembersbyid"
    id908467d80ca881966b187f3af525715c1cf4542c77a3c376e78ab4a4d40d021 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item/addmembersbyid"
    ie0663ce56074dacfdac086ea72d5e5de021ac050bb6e56380a99905fa9551b9c "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item/addmembers"
)

// UpdatableAssetRequestBuilder builds and executes requests for operations under \admin\windows\updates\updatableAssets\{updatableAsset-id}
type UpdatableAssetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdatableAssetRequestBuilderDeleteOptions options for Delete
type UpdatableAssetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UpdatableAssetRequestBuilderGetOptions options for Get
type UpdatableAssetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UpdatableAssetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UpdatableAssetRequestBuilderGetQueryParameters assets registered with the deployment service that can receive updates. Read-only.
type UpdatableAssetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UpdatableAssetRequestBuilderPatchOptions options for Patch
type UpdatableAssetRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UpdatableAssetRequestBuilder) AddMembers()(*ie0663ce56074dacfdac086ea72d5e5de021ac050bb6e56380a99905fa9551b9c.AddMembersRequestBuilder) {
    return ie0663ce56074dacfdac086ea72d5e5de021ac050bb6e56380a99905fa9551b9c.NewAddMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetRequestBuilder) AddMembersById()(*id908467d80ca881966b187f3af525715c1cf4542c77a3c376e78ab4a4d40d021.AddMembersByIdRequestBuilder) {
    return id908467d80ca881966b187f3af525715c1cf4542c77a3c376e78ab4a4d40d021.NewAddMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUpdatableAssetRequestBuilderInternal instantiates a new UpdatableAssetRequestBuilder and sets the default values.
func NewUpdatableAssetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetRequestBuilder) {
    m := &UpdatableAssetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/updatableAssets/{updatableAsset_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdatableAssetRequestBuilder instantiates a new UpdatableAssetRequestBuilder and sets the default values.
func NewUpdatableAssetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatableAssetRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetRequestBuilder) CreateDeleteRequestInformation(options *UpdatableAssetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UpdatableAssetRequestBuilder) CreateGetRequestInformation(options *UpdatableAssetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetRequestBuilder) CreatePatchRequestInformation(options *UpdatableAssetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetRequestBuilder) Delete(options *UpdatableAssetRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetRequestBuilder) Get(options *UpdatableAssetRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUpdatableAsset() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset), nil
}
// Patch assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetRequestBuilder) Patch(options *UpdatableAssetRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *UpdatableAssetRequestBuilder) RemoveMembers()(*i428d2c5ba8bcc9f861cfc7a3060a1e340984c54ca28a655b2305e6efffd7a1db.RemoveMembersRequestBuilder) {
    return i428d2c5ba8bcc9f861cfc7a3060a1e340984c54ca28a655b2305e6efffd7a1db.NewRemoveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetRequestBuilder) RemoveMembersById()(*i467add2156077b3c847fe4ad47f96c896d245f282bc0b3f06714966be14fec8e.RemoveMembersByIdRequestBuilder) {
    return i467add2156077b3c847fe4ad47f96c896d245f282bc0b3f06714966be14fec8e.NewRemoveMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
