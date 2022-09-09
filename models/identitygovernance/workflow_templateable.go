package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// WorkflowTemplateable 
type WorkflowTemplateable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*LifecycleWorkflowCategory)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExecutionConditions()(WorkflowExecutionConditionsable)
    GetTasks()([]Taskable)
    SetCategory(value *LifecycleWorkflowCategory)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExecutionConditions(value WorkflowExecutionConditionsable)()
    SetTasks(value []Taskable)()
}