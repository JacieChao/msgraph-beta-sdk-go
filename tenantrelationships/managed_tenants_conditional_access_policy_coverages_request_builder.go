package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderGetQueryParameters get a list of the conditionalAccessPolicyCoverage objects and their properties. Use this operation to list of Azure Active Directory conditional access policy coverage across all tenants that are being managed by the multi-tenant management platform. This API is supported in the following national cloud deployments.
type ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderGetQueryParameters
}
// ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByConditionalAccessPolicyCoverageId provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) ByConditionalAccessPolicyCoverageId(conditionalAccessPolicyCoverageId string)(*ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if conditionalAccessPolicyCoverageId != "" {
        urlTplParams["conditionalAccessPolicyCoverage%2Did"] = conditionalAccessPolicyCoverageId
    }
    return NewManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderInternal instantiates a new ConditionalAccessPolicyCoveragesRequestBuilder and sets the default values.
func NewManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) {
    m := &ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/conditionalAccessPolicyCoverages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder instantiates a new ConditionalAccessPolicyCoveragesRequestBuilder and sets the default values.
func NewManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) Count()(*ManagedTenantsConditionalAccessPolicyCoveragesCountRequestBuilder) {
    return NewManagedTenantsConditionalAccessPolicyCoveragesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the conditionalAccessPolicyCoverage objects and their properties. Use this operation to list of Azure Active Directory conditional access policy coverage across all tenants that are being managed by the multi-tenant management platform. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-managedtenant-list-conditionalaccesspolicycoverages?view=graph-rest-1.0
func (m *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateConditionalAccessPolicyCoverageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageCollectionResponseable), nil
}
// Post create new navigation property to conditionalAccessPolicyCoverages for tenantRelationships
func (m *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateConditionalAccessPolicyCoverageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable), nil
}
// ToGetRequestInformation get a list of the conditionalAccessPolicyCoverage objects and their properties. Use this operation to list of Azure Active Directory conditional access policy coverage across all tenants that are being managed by the multi-tenant management platform. This API is supported in the following national cloud deployments.
func (m *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to conditionalAccessPolicyCoverages for tenantRelationships
func (m *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) {
    return NewManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
