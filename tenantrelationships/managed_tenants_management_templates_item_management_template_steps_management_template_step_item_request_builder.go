package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managementTemplate entity.
type ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderGetQueryParameters get managementTemplateSteps from tenantRelationships
type ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderGetQueryParameters
}
// NewManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderInternal instantiates a new ManagementTemplateStepItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder) {
    m := &ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplates/{managementTemplate%2Did}/managementTemplateSteps/{managementTemplateStep%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder instantiates a new ManagementTemplateStepItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get managementTemplateSteps from tenantRelationships
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepable), nil
}
// ToGetRequestInformation get managementTemplateSteps from tenantRelationships
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder) {
    return NewManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
