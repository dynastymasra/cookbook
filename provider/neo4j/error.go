package neo4j

const (
	ErrorRecordNotFound    = "result contains no records"
	ErrorRecordMoreThanOne = "result contains more than one record"
)

// IsNotFound check Neo4J single query result if not found
func IsNotFound(err error) bool {
	if err != nil {
		return err.Error() == ErrorRecordNotFound
	}
	return false
}

// IsNotSingleResult check Neo4J query result more than one
func IsNotSingleResult(err error) bool {
	if err != nil {
		return err.Error() == ErrorRecordMoreThanOne
	}
	return false
}
