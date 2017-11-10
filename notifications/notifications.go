package notifications

// Rate denotes how often a notification is due
type Rate string

// Rate domain values
const (
	EachDay       Rate = "each-day"
	EachWeek      Rate = "each-week"
	EachTwoWeeks  Rate = "each-two-weeks"
	EachMonth     Rate = "each-month"
	EachTwoMonths Rate = "each-two-months"
)

// Weekday denotes the name of the days of the week for notifications
type Weekday string

// Weekday domain values
const (
	Sunday    Weekday = "sunday"
	Monday    Weekday = "monday"
	Tuesday   Weekday = "tuesday"
	Wednesday Weekday = "wednesday"
	Thursday  Weekday = "thursday"
	Friday    Weekday = "friday"
	Saturday  Weekday = "saturday"
)
