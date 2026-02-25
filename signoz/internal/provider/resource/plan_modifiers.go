package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// useUnknownOnUpdate is a plan modifier that sets the value to unknown
// during updates. Use this for computed fields that can change independently
// of Terraform, like alert firing state or update timestamps.
//
// On create: retains the prior behavior (unknown).
// On update: forces unknown so Terraform accepts any value returned by the API.
type useUnknownOnUpdate struct{}

func (m useUnknownOnUpdate) Description(_ context.Context) string {
	return "Sets the value to unknown during updates to prevent inconsistency errors for volatile computed fields."
}

func (m useUnknownOnUpdate) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useUnknownOnUpdate) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// If creating (no prior state), leave as unknown (default behavior).
	if req.State.Raw.IsNull() {
		return
	}

	// On update, force unknown so Terraform accepts any value the API returns.
	// This is necessary for fields like state, update_at, and update_by which
	// change on every API call regardless of what Terraform planned.
	resp.PlanValue = types.StringUnknown()
}
