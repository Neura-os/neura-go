package neura

type Actor struct {
	Type       string                 `json:"type"`
	ID         string                 `json:"id"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

type Resource struct {
	Type       string                 `json:"type"`
	ID         string                 `json:"id"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

type DecisionRequest struct {
	Intent        string                 `json:"intent"`
	Actor         Actor                  `json:"actor"`
	Resource      Resource               `json:"resource"`
	Context       map[string]interface{} `json:"context,omitempty"`
	CorrelationID string                 `json:"correlation_id,omitempty"`
}

type DecisionResponse struct {
	ID        string                 `json:"id"`
	Outcome   string                 `json:"outcome"`
	Class     string                 `json:"class,omitempty"`
	Evidence  *DecisionEvidence      `json:"evidence,omitempty"`
	Reason    string                 `json:"reason"`
	Policies  []string               `json:"policies"`
	Timestamp string                 `json:"timestamp"`
	AuditRef  string                 `json:"audit_ref"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type DecisionEvidence struct {
	DEO struct {
		DecisionID        string   `json:"decision_id"`
		Timestamp         string   `json:"timestamp"`
		Actor             string   `json:"actor"`
		AllowedByPolicy   string   `json:"allowed_by_policy,omitempty"`
		PolicyHash        string   `json:"policy_hash,omitempty"`
		MemorySnapshotRef string   `json:"memory_snapshot_ref,omitempty"`
		RiskScore         float64  `json:"risk_score"`
		ToolPermissions   []string `json:"tool_permissions,omitempty"`
		Confidence        float64  `json:"confidence"`
		Override          bool     `json:"override"`
	} `json:"deo"`
}

type ValidationResponse struct {
	Valid            bool   `json:"valid"`
	PredictedOutcome string `json:"predicted_outcome,omitempty"`
	Error            string `json:"error,omitempty"`
}

type MemoryRequest struct {
	Content    string                 `json:"content"`
	Type       string                 `json:"type"` // "episodic", "semantic"
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	IdentityID string                 `json:"identity_id,omitempty"`
}

type MemorySearchRequest struct {
	Query      string `json:"query"`
	Limit      int    `json:"limit,omitempty"`
	IdentityID string `json:"identity_id,omitempty"`
}

type MemoryResponse struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type AuthRequest struct {
	OrgID       string   `json:"org_id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions,omitempty"`
}

type AuthResponse struct {
	APIKey    string `json:"api_key"`
	Secret    string `json:"secret"`
	ExpiresAt string `json:"expires_at,omitempty"`
	Message   string `json:"message"`
}
