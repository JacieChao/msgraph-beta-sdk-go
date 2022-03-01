package customextensionhandlers

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i7218c9142036b1189f7c4b84a2190ace326032139e176ea1987270b72cb3978b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/customextensionhandlers/ref"
)

// CustomExtensionHandlersRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\accessPackageAssignmentPolicies\{accessPackageAssignmentPolicy-id}\customExtensionHandlers
type CustomExtensionHandlersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CustomExtensionHandlersRequestBuilderGetOptions options for Get
type CustomExtensionHandlersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CustomExtensionHandlersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustomExtensionHandlersRequestBuilderGetQueryParameters the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
type CustomExtensionHandlersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewCustomExtensionHandlersRequestBuilderInternal instantiates a new CustomExtensionHandlersRequestBuilder and sets the default values.
func NewCustomExtensionHandlersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomExtensionHandlersRequestBuilder) {
    m := &CustomExtensionHandlersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}/accessPackage/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy_id}/customExtensionHandlers{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustomExtensionHandlersRequestBuilder instantiates a new CustomExtensionHandlersRequestBuilder and sets the default values.
func NewCustomExtensionHandlersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomExtensionHandlersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomExtensionHandlersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
func (m *CustomExtensionHandlersRequestBuilder) CreateGetRequestInformation(options *CustomExtensionHandlersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
func (m *CustomExtensionHandlersRequestBuilder) Get(options *CustomExtensionHandlersRequestBuilderGetOptions)(*CustomExtensionHandlersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomExtensionHandlersResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*CustomExtensionHandlersResponse), nil
}
func (m *CustomExtensionHandlersRequestBuilder) Ref()(*i7218c9142036b1189f7c4b84a2190ace326032139e176ea1987270b72cb3978b.RefRequestBuilder) {
    return i7218c9142036b1189f7c4b84a2190ace326032139e176ea1987270b72cb3978b.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}