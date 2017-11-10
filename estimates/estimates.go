package estimates

// Outcome denotes the outcome of an estimation
type Outcome string

// Outcome domain values
const (
	Pending  Outcome = "pending"
	Positive Outcome = "positive"
	Negative Outcome = "negative"
)
