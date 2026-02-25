package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// useStateForUnknownIncludingNull is a plan modifier that copies the prior
// state value to the plan even when the state value is null. The built-in
// UseStateForUnknown() skips when state is null, which causes Optional+Computed
// fields (like evaluation) to show as "(known after apply)" when they're null
// in state and not set in config.
type useStateForUnknownIncludingNull struct{}

func (m useStateForUnknownIncludingNull) Description(_ context.Context) string {
	return "Uses the prior state value (including null) for unknown plan values."
}

func (m useStateForUnknownIncludingNull) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useStateForUnknownIncludingNull) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// On create (no prior state), keep unknown — let the provider compute.
	if req.State.Raw.IsNull() {
		return
	}

	// Only act when the plan value is unknown.
	if !resp.PlanValue.IsUnknown() {
		return
	}

	// Copy state value to plan, even if it's null.
	resp.PlanValue = req.StateValue
}

// objectUseStateForUnknownIncludingNull is the Object variant of the above.
type objectUseStateForUnknownIncludingNull struct{}

func (m objectUseStateForUnknownIncludingNull) Description(_ context.Context) string {
	return "Uses the prior state value (including null) for unknown plan values."
}

func (m objectUseStateForUnknownIncludingNull) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m objectUseStateForUnknownIncludingNull) PlanModifyObject(_ context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	// On create (no prior state), keep unknown — let the provider compute.
	if req.State.Raw.IsNull() {
		return
	}

	// Only act when the plan value is unknown.
	if !resp.PlanValue.IsUnknown() {
		return
	}

	// Copy state value to plan, even if it's null.
	resp.PlanValue = req.StateValue
}

// Compile-time interface checks.
var (
	_ planmodifier.String = useStateForUnknownIncludingNull{}
	_ planmodifier.Object = objectUseStateForUnknownIncludingNull{}
)
