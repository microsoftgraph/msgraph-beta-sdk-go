package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i29c75035d87e1e3a958b91c43025f79430cd4c9d3b0cc664ffe065bc407fda28 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/scopedrolemembers"
    i9a39d6011e7f6e52d0463e43f4affbe7fe566deec74cb2cfa89d0164876f73fb "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members"
    ib3392aac80c6073c1b4cb66c1c9a59ef64a733b53d9fbe3bb2fdeea43b1e8ed0 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/extensions"
    i348fc6817b120c2ca4453cfe220f3a31e3dd76b71234d51b3cf59448c038898e "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/extensions/item"
    ibb066ad55db4b7866d16a995539af863ab771176674fded378a94af09016fc1b "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/scopedrolemembers/item"
    id21ccd20a38c8c744ae37c53448fd59c79e0a14764716cf4113d7e8fd6f62db0 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/item"
)

// AdministrativeUnitItemRequestBuilder provides operations to manage the administrativeUnits property of the microsoft.graph.directory entity.
type AdministrativeUnitItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AdministrativeUnitItemRequestBuilderDeleteOptions options for Delete
type AdministrativeUnitItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AdministrativeUnitItemRequestBuilderGetOptions options for Get
type AdministrativeUnitItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AdministrativeUnitItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AdministrativeUnitItemRequestBuilderGetQueryParameters conceptual container for user and group directory objects.
type AdministrativeUnitItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AdministrativeUnitItemRequestBuilderPatchOptions options for Patch
type AdministrativeUnitItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnitable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAdministrativeUnitItemRequestBuilderInternal instantiates a new AdministrativeUnitItemRequestBuilder and sets the default values.
func NewAdministrativeUnitItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitItemRequestBuilder) {
    m := &AdministrativeUnitItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/administrativeUnits/{administrativeUnit_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdministrativeUnitItemRequestBuilder instantiates a new AdministrativeUnitItemRequestBuilder and sets the default values.
func NewAdministrativeUnitItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property administrativeUnits for directory
func (m *AdministrativeUnitItemRequestBuilder) CreateDeleteRequestInformation(options *AdministrativeUnitItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation conceptual container for user and group directory objects.
func (m *AdministrativeUnitItemRequestBuilder) CreateGetRequestInformation(options *AdministrativeUnitItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property administrativeUnits in directory
func (m *AdministrativeUnitItemRequestBuilder) CreatePatchRequestInformation(options *AdministrativeUnitItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property administrativeUnits for directory
func (m *AdministrativeUnitItemRequestBuilder) Delete(options *AdministrativeUnitItemRequestBuilderDeleteOptions)(error) {
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
func (m *AdministrativeUnitItemRequestBuilder) Extensions()(*ib3392aac80c6073c1b4cb66c1c9a59ef64a733b53d9fbe3bb2fdeea43b1e8ed0.ExtensionsRequestBuilder) {
    return ib3392aac80c6073c1b4cb66c1c9a59ef64a733b53d9fbe3bb2fdeea43b1e8ed0.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.administrativeUnits.item.extensions.item collection
func (m *AdministrativeUnitItemRequestBuilder) ExtensionsById(id string)(*i348fc6817b120c2ca4453cfe220f3a31e3dd76b71234d51b3cf59448c038898e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i348fc6817b120c2ca4453cfe220f3a31e3dd76b71234d51b3cf59448c038898e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get conceptual container for user and group directory objects.
func (m *AdministrativeUnitItemRequestBuilder) Get(options *AdministrativeUnitItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnitable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAdministrativeUnitFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnitable), nil
}
func (m *AdministrativeUnitItemRequestBuilder) Members()(*i9a39d6011e7f6e52d0463e43f4affbe7fe566deec74cb2cfa89d0164876f73fb.MembersRequestBuilder) {
    return i9a39d6011e7f6e52d0463e43f4affbe7fe566deec74cb2cfa89d0164876f73fb.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.administrativeUnits.item.members.item collection
func (m *AdministrativeUnitItemRequestBuilder) MembersById(id string)(*id21ccd20a38c8c744ae37c53448fd59c79e0a14764716cf4113d7e8fd6f62db0.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return id21ccd20a38c8c744ae37c53448fd59c79e0a14764716cf4113d7e8fd6f62db0.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property administrativeUnits in directory
func (m *AdministrativeUnitItemRequestBuilder) Patch(options *AdministrativeUnitItemRequestBuilderPatchOptions)(error) {
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
func (m *AdministrativeUnitItemRequestBuilder) ScopedRoleMembers()(*i29c75035d87e1e3a958b91c43025f79430cd4c9d3b0cc664ffe065bc407fda28.ScopedRoleMembersRequestBuilder) {
    return i29c75035d87e1e3a958b91c43025f79430cd4c9d3b0cc664ffe065bc407fda28.NewScopedRoleMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.administrativeUnits.item.scopedRoleMembers.item collection
func (m *AdministrativeUnitItemRequestBuilder) ScopedRoleMembersById(id string)(*ibb066ad55db4b7866d16a995539af863ab771176674fded378a94af09016fc1b.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ibb066ad55db4b7866d16a995539af863ab771176674fded378a94af09016fc1b.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
