// Code generated by ent, DO NOT EDIT.

package education

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the education type in the database.
	Label = "education"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInstitution holds the string denoting the institution field in the database.
	FieldInstitution = "institution"
	// FieldDegree holds the string denoting the degree field in the database.
	FieldDegree = "degree"
	// FieldField holds the string denoting the field field in the database.
	FieldField = "field"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldGrade holds the string denoting the grade field in the database.
	FieldGrade = "grade"
	// EdgeResume holds the string denoting the resume edge name in mutations.
	EdgeResume = "resume"
	// Table holds the table name of the education in the database.
	Table = "educations"
	// ResumeTable is the table that holds the resume relation/edge. The primary key declared below.
	ResumeTable = "resume_educations"
	// ResumeInverseTable is the table name for the Resume entity.
	// It exists in this package in order to avoid circular dependency with the "resume" package.
	ResumeInverseTable = "resumes"
)

// Columns holds all SQL columns for education fields.
var Columns = []string{
	FieldID,
	FieldInstitution,
	FieldDegree,
	FieldField,
	FieldStartDate,
	FieldEndDate,
	FieldDescription,
	FieldGrade,
}

var (
	// ResumePrimaryKey and ResumeColumn2 are the table columns denoting the
	// primary key for the resume relation (M2M).
	ResumePrimaryKey = []string{"resume_id", "education_id"}
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
	// InstitutionValidator is a validator for the "institution" field. It is called by the builders before save.
	InstitutionValidator func(string) error
	// DegreeValidator is a validator for the "degree" field. It is called by the builders before save.
	DegreeValidator func(string) error
	// FieldValidator is a validator for the "field" field. It is called by the builders before save.
	FieldValidator func(string) error
)

// OrderOption defines the ordering options for the Education queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByInstitution orders the results by the institution field.
func ByInstitution(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstitution, opts...).ToFunc()
}

// ByDegree orders the results by the degree field.
func ByDegree(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDegree, opts...).ToFunc()
}

// ByField orders the results by the field field.
func ByField(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldField, opts...).ToFunc()
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

// ByGrade orders the results by the grade field.
func ByGrade(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGrade, opts...).ToFunc()
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
