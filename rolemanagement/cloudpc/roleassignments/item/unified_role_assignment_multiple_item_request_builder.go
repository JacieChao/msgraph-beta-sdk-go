package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2424edf24ccaf6954cd497652d997c5f0fbd9fa915451718fa8cd09e6afa86ea "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments/item/roledefinition"
    i3ea80963a870d6c2da85b7564b54a59b5083cdecd3fa05a1c910c987a6f5a6a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments/item/principals"
    i82b47c6d1d1d3da67a4ce306857a8170b9626633b5d4daf7bbcc6ea00da8dfca "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments/item/directoryscopes"
    ia3c1ca878d611a589d25bead520613141a62abf7f329cc799c6ff5f3a2e0862c "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments/item/appscopes"
    ib2d17989f9484a745a0618303d9de9c94a6e8769ad7a3bc73bb9c5231ff6cc36 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments/item/appscopes/item"
)

// UnifiedRoleAssignmentMultipleItemRequestBuilder builds and executes requests for operations under \roleManagement\cloudPC\roleAssignments\{unifiedRoleAssignmentMultiple-id}
type UnifiedRoleAssignmentMultipleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentMultipleItemRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentMultipleItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentMultipleItemRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentMultipleItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters get roleAssignments from roleManagement
type UnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentMultipleItemRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentMultipleItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentMultiple;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopes()(*ia3c1ca878d611a589d25bead520613141a62abf7f329cc799c6ff5f3a2e0862c.AppScopesRequestBuilder) {
    return ia3c1ca878d611a589d25bead520613141a62abf7f329cc799c6ff5f3a2e0862c.NewAppScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.cloudPC.roleAssignments.item.appScopes.item collection
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopesById(id string)(*ib2d17989f9484a745a0618303d9de9c94a6e8769ad7a3bc73bb9c5231ff6cc36.AppScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appScope_id"] = id
    }
    return ib2d17989f9484a745a0618303d9de9c94a6e8769ad7a3bc73bb9c5231ff6cc36.NewAppScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUnifiedRoleAssignmentMultipleItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentMultipleItemRequestBuilder) {
    m := &UnifiedRoleAssignmentMultipleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/cloudPC/roleAssignments/{unifiedRoleAssignmentMultiple_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentMultipleItemRequestBuilder instantiates a new UnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentMultipleItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentMultipleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for roleManagement
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentMultipleItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get roleAssignments from roleManagement
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentMultipleItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignments in roleManagement
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentMultipleItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property roleAssignments for roleManagement
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) Delete(options *UnifiedRoleAssignmentMultipleItemRequestBuilderDeleteOptions)(error) {
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
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) DirectoryScopes()(*i82b47c6d1d1d3da67a4ce306857a8170b9626633b5d4daf7bbcc6ea00da8dfca.DirectoryScopesRequestBuilder) {
    return i82b47c6d1d1d3da67a4ce306857a8170b9626633b5d4daf7bbcc6ea00da8dfca.NewDirectoryScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleAssignments from roleManagement
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) Get(options *UnifiedRoleAssignmentMultipleItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentMultiple, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleAssignmentMultiple() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentMultiple), nil
}
// Patch update the navigation property roleAssignments in roleManagement
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) Patch(options *UnifiedRoleAssignmentMultipleItemRequestBuilderPatchOptions)(error) {
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
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) Principals()(*i3ea80963a870d6c2da85b7564b54a59b5083cdecd3fa05a1c910c987a6f5a6a4.PrincipalsRequestBuilder) {
    return i3ea80963a870d6c2da85b7564b54a59b5083cdecd3fa05a1c910c987a6f5a6a4.NewPrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentMultipleItemRequestBuilder) RoleDefinition()(*i2424edf24ccaf6954cd497652d997c5f0fbd9fa915451718fa8cd09e6afa86ea.RoleDefinitionRequestBuilder) {
    return i2424edf24ccaf6954cd497652d997c5f0fbd9fa915451718fa8cd09e6afa86ea.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}