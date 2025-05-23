// Code generated by ent, DO NOT EDIT.

package experience

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the experience type in the database.
	Label = "experience"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "company"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSkills holds the string denoting the skills field in the database.
	FieldSkills = "skills"
	// FieldCurrent holds the string denoting the current field in the database.
	FieldCurrent = "current"
	// EdgeResume holds the string denoting the resume edge name in mutations.
	EdgeResume = "resume"
	// Table holds the table name of the experience in the database.
	Table = "experiences"
	// ResumeTable is the table that holds the resume relation/edge. The primary key declared below.
	ResumeTable = "resume_experiences"
	// ResumeInverseTable is the table name for the Resume entity.
	// It exists in this package in order to avoid circular dependency with the "resume" package.
	ResumeInverseTable = "resumes"
)

// Columns holds all SQL columns for experience fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldCompany,
	FieldLocation,
	FieldStartDate,
	FieldEndDate,
	FieldDescription,
	FieldSkills,
	FieldCurrent,
}

var (
	// ResumePrimaryKey and ResumeColumn2 are the table columns denoting the
	// primary key for the resume relation (M2M).
	ResumePrimaryKey = []string{"resume_id", "experience_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// CompanyValidator is a validator for the "company" field. It is called by the builders before save.
	CompanyValidator func(string) error
	// DefaultCurrent holds the default value on creation for the "current" field.
	DefaultCurrent bool
)

// OrderOption defines the ordering options for the Experience queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByCompany orders the results by the company field.
func ByCompany(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompany, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// BySkills orders the results by the skills field.
func BySkills(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSkills, opts...).ToFunc()
}

// ByCurrent orders the results by the current field.
func ByCurrent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrent, opts...).ToFunc()
}

// ByResumeCount orders the results by resume count.
func ByResumeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newResumeStep(), opts...)
	}
}

// ByResume orders the results by resume terms.
func ByResume(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResumeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newResumeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResumeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ResumeTable, ResumePrimaryKey...),
	)
}
