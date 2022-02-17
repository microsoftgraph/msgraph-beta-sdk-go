package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9375b68a6cf8c4f83797797ec44b3879b94345db57d0fb5ea574f869a3495641 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item/removemembers"
    ibdd955c15af1ffd72a16c3447461288c34a9983679f832b64197435160037fc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item/removemembersbyid"
    id976fda7cf3837c62cb909768a45cf653bcaefbdbba660d43bc7d7490545d920 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item/addmembersbyid"
    if09247f50a62b73060b83df7a38cb520f78f6d06ce735570e8585a8632ca9542 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item/addmembers"
)

// UpdatableAssetRequestBuilder builds and executes requests for operations under \admin\windows\updates\deployments\{deployment-id}\audience\exclusions\{updatableAsset-id}
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
// UpdatableAssetRequestBuilderGetQueryParameters specifies the assets to exclude from the audience.
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
func (m *UpdatableAssetRequestBuilder) AddMembers()(*if09247f50a62b73060b83df7a38cb520f78f6d06ce735570e8585a8632ca9542.AddMembersRequestBuilder) {
    return if09247f50a62b73060b83df7a38cb520f78f6d06ce735570e8585a8632ca9542.NewAddMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetRequestBuilder) AddMembersById()(*id976fda7cf3837c62cb909768a45cf653bcaefbdbba660d43bc7d7490545d920.AddMembersByIdRequestBuilder) {
    return id976fda7cf3837c62cb909768a45cf653bcaefbdbba660d43bc7d7490545d920.NewAddMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUpdatableAssetRequestBuilderInternal instantiates a new UpdatableAssetRequestBuilder and sets the default values.
func NewUpdatableAssetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetRequestBuilder) {
    m := &UpdatableAssetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment_id}/audience/exclusions/{updatableAsset_id}{?select,expand}";
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
// CreateDeleteRequestInformation specifies the assets to exclude from the audience.
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
// CreateGetRequestInformation specifies the assets to exclude from the audience.
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
// CreatePatchRequestInformation specifies the assets to exclude from the audience.
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
// Delete specifies the assets to exclude from the audience.
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
// Get specifies the assets to exclude from the audience.
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
// Patch specifies the assets to exclude from the audience.
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
func (m *UpdatableAssetRequestBuilder) RemoveMembers()(*i9375b68a6cf8c4f83797797ec44b3879b94345db57d0fb5ea574f869a3495641.RemoveMembersRequestBuilder) {
    return i9375b68a6cf8c4f83797797ec44b3879b94345db57d0fb5ea574f869a3495641.NewRemoveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetRequestBuilder) RemoveMembersById()(*ibdd955c15af1ffd72a16c3447461288c34a9983679f832b64197435160037fc8.RemoveMembersByIdRequestBuilder) {
    return ibdd955c15af1ffd72a16c3447461288c34a9983679f832b64197435160037fc8.NewRemoveMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
