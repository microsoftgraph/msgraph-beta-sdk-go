package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9375b68a6cf8c4f83797797ec44b3879b94345db57d0fb5ea574f869a3495641 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item/removemembers"
    ibdd955c15af1ffd72a16c3447461288c34a9983679f832b64197435160037fc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item/removemembersbyid"
    id976fda7cf3837c62cb909768a45cf653bcaefbdbba660d43bc7d7490545d920 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item/addmembersbyid"
    if09247f50a62b73060b83df7a38cb520f78f6d06ce735570e8585a8632ca9542 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item/addmembers"
)

// UpdatableAssetItemRequestBuilder provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
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
// UpdatableAssetItemRequestBuilderGetQueryParameters specifies the assets to exclude from the audience.
type UpdatableAssetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UpdatableAssetItemRequestBuilderPatchOptions options for Patch
type UpdatableAssetItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAssetable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UpdatableAssetItemRequestBuilder) AddMembers()(*if09247f50a62b73060b83df7a38cb520f78f6d06ce735570e8585a8632ca9542.AddMembersRequestBuilder) {
    return if09247f50a62b73060b83df7a38cb520f78f6d06ce735570e8585a8632ca9542.NewAddMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetItemRequestBuilder) AddMembersById()(*id976fda7cf3837c62cb909768a45cf653bcaefbdbba660d43bc7d7490545d920.AddMembersByIdRequestBuilder) {
    return id976fda7cf3837c62cb909768a45cf653bcaefbdbba660d43bc7d7490545d920.NewAddMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUpdatableAssetItemRequestBuilderInternal instantiates a new UpdatableAssetItemRequestBuilder and sets the default values.
func NewUpdatableAssetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetItemRequestBuilder) {
    m := &UpdatableAssetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment_id}/audience/exclusions/{updatableAsset_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exclusions for admin
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
// CreateGetRequestInformation specifies the assets to exclude from the audience.
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
// CreatePatchRequestInformation update the navigation property exclusions in admin
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
// Delete delete navigation property exclusions for admin
func (m *UpdatableAssetItemRequestBuilder) Delete(options *UpdatableAssetItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get specifies the assets to exclude from the audience.
func (m *UpdatableAssetItemRequestBuilder) Get(options *UpdatableAssetItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAssetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUpdatableAssetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAssetable), nil
}
// Patch update the navigation property exclusions in admin
func (m *UpdatableAssetItemRequestBuilder) Patch(options *UpdatableAssetItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UpdatableAssetItemRequestBuilder) RemoveMembers()(*i9375b68a6cf8c4f83797797ec44b3879b94345db57d0fb5ea574f869a3495641.RemoveMembersRequestBuilder) {
    return i9375b68a6cf8c4f83797797ec44b3879b94345db57d0fb5ea574f869a3495641.NewRemoveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetItemRequestBuilder) RemoveMembersById()(*ibdd955c15af1ffd72a16c3447461288c34a9983679f832b64197435160037fc8.RemoveMembersByIdRequestBuilder) {
    return ibdd955c15af1ffd72a16c3447461288c34a9983679f832b64197435160037fc8.NewRemoveMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
