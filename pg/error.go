package pg

import (
	"github.com/fobus1289/ufa_shared/http"
	"github.com/jackc/pgconn"
)

const (
	// Class 00 — Successful Completion
	ERR_SUCCESSFUL_COMPLETION = "successful_completion" // 00000

	// Class 01 — Warning
	ERR_WARNING                           = "warning"                               // 01000
	ERR_DYNAMIC_RESULT_SETS_RETURNED      = "dynamic_result_sets_returned"          // 0100C
	ERR_IMPLICIT_ZERO_BIT_PADDING         = "implicit_zero_bit_padding"             // 01008
	ERR_NULL_VALUE_ELIMINATED_IN_SET_FUNC = "null_value_eliminated_in_set_function" // 01003
	ERR_PRIVILEGE_NOT_GRANTED             = "privilege_not_granted"                 // 01007
	ERR_PRIVILEGE_NOT_REVOKED             = "privilege_not_revoked"                 // 01006
	ERR_STRING_DATA_RIGHT_TRUNCATION      = "string_data_right_truncation"          // 01004
	ERR_DEPRECATED_FEATURE                = "deprecated_feature"                    // 01P01

	// Class 02 — No Data (this is also a warning class per the SQL standard)
	ERR_NO_DATA                                    = "no_data"                                    // 02000
	ERR_NO_ADDITIONAL_DYNAMIC_RESULT_SETS_RETURNED = "no_additional_dynamic_result_sets_returned" // 02001

	// Class 03 — SQL Statement Not Yet Complete
	ERR_SQL_STATEMENT_NOT_YET_COMPLETE = "sql_statement_not_yet_complete" // 03000

	// Class 08 — Connection Exception
	ERR_CONNECTION_EXCEPTION                                = "connection_exception"                                // 08000
	ERR_CONNECTION_DOES_NOT_EXIST                           = "connection_does_not_exist"                           // 08003
	ERR_CONNECTION_FAILURE                                  = "connection_failure"                                  // 08006
	ERR_SQL_CLIENT_UNABLE_TO_ESTABLISH_SQL_CONNECTION       = "sql_client_unable_to_establish_sql_connection"       // 08001
	ERR_SQL_SERVER_REJECTED_ESTABLISHMENT_OF_SQL_CONNECTION = "sql_server_rejected_establishment_of_sql_connection" // 08004
	ERR_TRANSACTION_RESOLUTION_UNKNOWN                      = "transaction_resolution_unknown"                      // 08007
	ERR_PROTOCOL_VIOLATION                                  = "protocol_violation"                                  // 08P01

	// Class 09 — Triggered Action Exception
	ERR_TRIGGERED_ACTION_EXCEPTION = "triggered_action_exception" // 09000

	// Class 0A — Feature Not Supported
	ERR_FEATURE_NOT_SUPPORTED = "feature_not_supported" // 0A000

	// Class 0B — Invalid Transaction Initiation
	ERR_INVALID_TRANSACTION_INITIATION = "invalid_transaction_initiation" // 0B000

	// Class 0F — Locator Exception
	ERR_LOCATOR_EXCEPTION             = "locator_exception"             // 0F000
	ERR_INVALID_LOCATOR_SPECIFICATION = "invalid_locator_specification" // 0F001

	// Class 0L — Invalid Grantor
	ERR_INVALID_GRANTOR         = "invalid_grantor"         // 0L000
	ERR_INVALID_GRANT_OPERATION = "invalid_grant_operation" // 0LP01

	// Class 0P — Invalid Role Specification
	ERR_INVALID_ROLE_SPECIFICATION = "invalid_role_specification" // 0P000

	// Class 0Z — Diagnostics Exception
	ERR_DIAGNOSTICS_EXCEPTION                               = "diagnostics_exception"                               // 0Z000
	ERR_STACKED_DIAGNOSTICS_ACCESSED_WITHOUT_ACTIVE_HANDLER = "stacked_diagnostics_accessed_without_active_handler" // 0Z002

	// Class 20 — Case Not Found
	ERR_CASE_NOT_FOUND = "case_not_found" // 20000

	// Class 21 — Cardinality Violation
	ERR_CARDINALITY_VIOLATION = "cardinality_violation" // 21000

	// Class 22 — Data Exception
	ERR_DATA_EXCEPTION                              = "data_exception"                              // 22000
	ERR_ARRAY_SUBSCRIPT_ERROR                       = "array_subscript_error"                       // 2202E
	ERR_CHARACTER_NOT_IN_REPERTOIRE                 = "character_not_in_repertoire"                 // 22021
	ERR_DATETIME_FIELD_OVERFLOW                     = "datetime_field_overflow"                     // 22008
	ERR_DIVISION_BY_ZERO                            = "division_by_zero"                            // 22012
	ERR_ERROR_IN_ASSIGNMENT                         = "error_in_assignment"                         // 22005
	ERR_ESCAPE_CHARACTER_CONFLICT                   = "escape_character_conflict"                   // 2200B
	ERR_INDICATOR_OVERFLOW                          = "indicator_overflow"                          // 22022
	ERR_INTERVAL_FIELD_OVERFLOW                     = "interval_field_overflow"                     // 22015
	ERR_INVALID_ARGUMENT_FOR_LOGARITHM              = "invalid_argument_for_logarithm"              // 2201E
	ERR_INVALID_ARGUMENT_FOR_NTILE_FUNCTION         = "invalid_argument_for_ntile_function"         // 22014
	ERR_INVALID_ARGUMENT_FOR_NTH_VALUE_FUNCTION     = "invalid_argument_for_nth_value_function"     // 22016
	ERR_INVALID_ARGUMENT_FOR_POWER_FUNCTION         = "invalid_argument_for_power_function"         // 2201F
	ERR_INVALID_ARGUMENT_FOR_WIDTH_BUCKET_FUNCTION  = "invalid_argument_for_width_bucket_function"  // 2201G
	ERR_INVALID_CHARACTER_VALUE_FOR_CAST            = "invalid_character_value_for_cast"            // 22018
	ERR_INVALID_DATETIME_FORMAT                     = "invalid_datetime_format"                     // 22007
	ERR_INVALID_ESCAPE_CHARACTER                    = "invalid_escape_character"                    // 22019
	ERR_INVALID_ESCAPE_OCTET                        = "invalid_escape_octet"                        // 2200D
	ERR_INVALID_ESCAPE_SEQUENCE                     = "invalid_escape_sequence"                     // 22025
	ERR_NONSTANDARD_USE_OF_ESCAPE_CHARACTER         = "nonstandard_use_of_escape_character"         // 22P06
	ERR_INVALID_INDICATOR_PARAMETER_VALUE           = "invalid_indicator_parameter_value"           // 22010
	ERR_INVALID_PARAMETER_VALUE                     = "invalid_parameter_value"                     // 22023
	ERR_INVALID_PRECEDING_OR_FOLLOWING_SIZE         = "invalid_preceding_or_following_size"         // 22013
	ERR_INVALID_REGULAR_EXPRESSION                  = "invalid_regular_expression"                  // 2201B
	ERR_INVALID_ROW_COUNT_IN_LIMIT_CLAUSE           = "invalid_row_count_in_limit_clause"           // 2201W
	ERR_INVALID_ROW_COUNT_IN_RESULT_OFFSET_CLAUSE   = "invalid_row_count_in_result_offset_clause"   // 2201X
	ERR_INVALID_TABLE_SAMPLE_ARGUMENT               = "invalid_table_sample_argument"               // 2202H
	ERR_INVALID_TABLE_SAMPLE_REPEAT                 = "invalid_table_sample_repeat"                 // 2202G
	ERR_INVALID_TIME_ZONE_DISPLACEMENT_VALUE        = "invalid_time_zone_displacement_value"        // 22009
	ERR_INVALID_USE_OF_ESCAPE_CHARACTER             = "invalid_use_of_escape_character"             // 2200C
	ERR_MOST_SPECIFIC_TYPE_MISMATCH                 = "most_specific_type_mismatch"                 // 2200G
	ERR_NULL_VALUE_NOT_ALLOWED_DATA_EXCEPTION       = "null_value_not_allowed_data_exception"       // 22004
	ERR_NULL_VALUE_NO_INDICATOR_PARAMETER           = "null_value_no_indicator_parameter"           // 22002
	ERR_NUMERIC_VALUE_OUT_OF_RANGE                  = "numeric_value_out_of_range"                  // 22003
	ERR_SEQUENCE_GENERATOR_LIMIT_EXCEEDED           = "sequence_generator_limit_exceeded"           // 2200H
	ERR_STRING_DATA_LENGTH_MISMATCH                 = "string_data_length_mismatch"                 // 22026
	ERR_STRING_DATA_RIGHT_TRUNCATION_DATA_EXCEPTION = "string_data_right_truncation_data_exception" // 22001
	ERR_SUBSTRING_ERROR                             = "substring_error"                             // 22011
	ERR_TRIM_ERROR                                  = "trim_error"                                  // 22027
	ERR_UNTERMINATED_C_STRING                       = "unterminated_c_string"                       // 22024
	ERR_ZERO_LENGTH_CHARACTER_STRING                = "zero_length_character_string"                // 2200F
	ERR_FLOATING_POINT_EXCEPTION                    = "floating_point_exception"                    // 22P01
	ERR_INVALID_TEXT_REPRESENTATION                 = "invalid_text_representation"                 // 22P02
	ERR_INVALID_BINARY_REPRESENTATION               = "invalid_binary_representation"               // 22P03
	ERR_BAD_COPY_FILE_FORMAT                        = "bad_copy_file_format"                        // 22P04
	ERR_UNTRANSLATABLE_CHARACTER                    = "untranslatable_character"                    // 22P05
	ERR_NOT_AN_XML_DOCUMENT                         = "not_an_xml_document"                         // 2200L
	ERR_INVALID_XML_DOCUMENT                        = "invalid_xml_document"                        // 2200M
	ERR_INVALID_XML_CONTENT                         = "invalid_xml_content"                         // 2200N
	ERR_INVALID_XML_COMMENT                         = "invalid_xml_comment"                         // 2200S
	ERR_INVALID_XML_PROCESSING_INSTRUCTION          = "invalid_xml_processing_instruction"          // 2200T

	// Class 23 — Integrity Constraint Violation
	ERR_INTEGRITY_CONSTRAINT_VIOLATION = "integrity_constraint_violation" // 23000
	ERR_RESTRICT_VIOLATION             = "restrict_violation"             // 23001
	ERR_NOT_NULL_VIOLATION             = "not_null_violation"             // 23502
	ERR_FOREIGN_KEY_VIOLATION          = "foreign_key_violation"          // 23503
	ERR_UNIQUE_VIOLATION               = "unique_violation"               // 23505
	ERR_CHECK_VIOLATION                = "check_violation"                // 23514
	ERR_EXCLUSION_VIOLATION            = "exclusion_violation"            // 23P01

	// Class 24 — Invalid Cursor State
	ERR_INVALID_CURSOR_STATE = "invalid_cursor_state" // 24000

	// Class 25 — Invalid Transaction State
	ERR_INVALID_TRANSACTION_STATE                             = "invalid_transaction_state"                             // 25000
	ERR_ACTIVE_SQL_TRANSACTION                                = "active_sql_transaction"                                // 25001
	ERR_BRANCH_TRANSACTION_ALREADY_ACTIVE                     = "branch_transaction_already_active"                     // 25002
	ERR_HELD_CURSOR_REQUIRES_SAME_ISOLATION_LEVEL             = "held_cursor_requires_same_isolation_level"             // 25008
	ERR_INAPPROPRIATE_ACCESS_MODE_FOR_BRANCH_TRANSACTION      = "inappropriate_access_mode_for_branch_transaction"      // 25003
	ERR_INAPPROPRIATE_ISOLATION_LEVEL_FOR_BRANCH_TRANSACTION  = "inappropriate_isolation_level_for_branch_transaction"  // 25004
	ERR_NO_ACTIVE_SQL_TRANSACTION                             = "no_active_sql_transaction"                             // 25P01
	ERR_READ_ONLY_SQL_TRANSACTION                             = "read_only_sql_transaction"                             // 25006
	ERR_SCHEMA_AND_DATA_STATEMENT_MIXING_NOT_SUPPORTED        = "schema_and_data_statement_mixing_not_supported"        // 25007
	ERR_NO_ACTIVE_SQL_TRANSACTION_FOR_BRANCH_TRANSACTION      = "no_active_sql_transaction_for_branch_transaction"      // 25005
	ERR_READ_ONLY_SQL_TRANSACTION_STATEMENT_ISSUED            = "read_only_sql_transaction_statement_issued"            // 25011
	ERR_READ_ONLY_SQL_TRANSACTION_STATEMENT_ISSUED_IN_FG_ONLY = "read_only_sql_transaction_statement_issued_in_fg_only" // 25012

	// Class 26 — Invalid SQL Statement Name
	ERR_INVALID_SQL_STATEMENT_NAME = "invalid_sql_statement_name" // 26000

	// Class 27 — Triggered Data Change Violation
	ERR_TRIGGERED_DATA_CHANGE_VIOLATION = "triggered_data_change_violation" // 27000

	// Class 28 — Invalid Authorization Specification
	ERR_INVALID_AUTHORIZATION_SPECIFICATION = "invalid_authorization_specification" // 28000
	ERR_INVALID_PASSWORD                    = "invalid_password"                    // 28P01

	// Class 2B — Dependent Privilege Descriptors Still Exist
	ERR_DEPENDENT_PRIVILEGE_DESCRIPTORS_STILL_EXIST = "dependent_privilege_descriptors_still_exist" // 2B000

	// Class 2D — Invalid Transaction Termination
	ERR_INVALID_TRANSACTION_TERMINATION = "invalid_transaction_termination" // 2D000

	// Class 2F — SQL Routine Exception
	ERR_SQL_ROUTINE_EXCEPTION                 = "sql_routine_exception"                 // 2F000
	ERR_FUNCTION_EXECUTED_NO_RETURN_STATEMENT = "function_executed_no_return_statement" // 2F005
	ERR_MODIFYING_SQL_DATA_NOT_PERMITTED      = "modifying_sql_data_not_permitted"      // 2F002
	ERR_PROHIBITED_SQL_STATEMENT_ATTEMPTED    = "prohibited_sql_statement_attempted"    // 2F003
	ERR_READING_SQL_DATA_NOT_PERMITTED        = "reading_sql_data_not_permitted"        // 2F004

	// Class 34 — Invalid Cursor Name
	ERR_INVALID_CURSOR_NAME = "invalid_cursor_name" // 34000

	// Class 38 — External Routine Exception
	ERR_EXTERNAL_ROUTINE_EXCEPTION                          = "external_routine_exception"                          // 38000
	ERR_CONTAINING_SQL_NOT_PERMITTED                        = "containing_sql_not_permitted"                        // 38001
	ERR_MODIFYING_SQL_DATA_NOT_PERMITTED_EXTERNAL_ROUTINE   = "modifying_sql_data_not_permitted_external_routine"   // 38002
	ERR_PROHIBITED_SQL_STATEMENT_ATTEMPTED_EXTERNAL_ROUTINE = "prohibited_sql_statement_attempted_external_routine" // 38003
	ERR_READING_SQL_DATA_NOT_PERMITTED_EXTERNAL_ROUTINE     = "reading_sql_data_not_permitted_external_routine"     // 38004

	// Class 39 — External Routine Invocation Exception
	ERR_EXTERNAL_ROUTINE_INVOCATION_EXCEPTION                        = "external_routine_invocation_exception"                        // 39000
	ERR_INVALID_SQLSTATE_RETURNED                                    = "invalid_sqlstate_returned"                                    // 39001
	ERR_NULL_VALUE_NOT_ALLOWED_EXTERNAL_ROUTINE_INVOCATION_EXCEPTION = "null_value_not_allowed_external_routine_invocation_exception" // 39004
	ERR_TRIGGER_PROTOCOL_VIOLATED                                    = "trigger_protocol_violated"                                    // 39P01
	ERR_SRF_PROTOCOL_VIOLATED                                        = "srf_protocol_violated"                                        // 39P02
	ERR_EVENT_TRIGGER_PROTOCOL_VIOLATED                              = "event_trigger_protocol_violated"                              // 39P03

	// Class 3B — Savepoint Exception
	ERR_SAVEPOINT_EXCEPTION             = "savepoint_exception"             // 3B000
	ERR_INVALID_SAVEPOINT_SPECIFICATION = "invalid_savepoint_specification" // 3B001

	// Class 3D — Invalid Catalog Name
	ERR_INVALID_CATALOG_NAME = "invalid_catalog_name" // 3D000

	// Class 3F — Invalid Schema Name
	ERR_INVALID_SCHEMA_NAME = "invalid_schema_name" // 3F000

	// Class 40 — Transaction Rollback
	ERR_TRANSACTION_ROLLBACK                       = "transaction_rollback"                       // 40000
	ERR_TRANSACTION_INTEGRITY_CONSTRAINT_VIOLATION = "transaction_integrity_constraint_violation" // 40002
	ERR_SERIALIZATION_FAILURE                      = "serialization_failure"                      // 40001
	ERR_STATEMENT_COMPLETION_UNKNOWN               = "statement_completion_unknown"               // 40003
	ERR_DEADLOCK_DETECTED                          = "deadlock_detected"                          // 40P01

	// Class 42 — Syntax Error or Access Rule Violation
	ERR_SYNTAX_ERROR_OR_ACCESS_RULE_VIOLATION = "syntax_error_or_access_rule_violation" // 42000
	ERR_SYNTAX_ERROR                          = "syntax_error"                          // 42601
	ERR_INSUFFICIENT_PRIVILEGE                = "insufficient_privilege"                // 42501
	ERR_CANNOT_COERCE                         = "cannot_coerce"                         // 42846
	ERR_GROUPING_ERROR                        = "grouping_error"                        // 42803
	ERR_WINDOWING_ERROR                       = "windowing_error"                       // 42P20
	ERR_INVALID_RECURSION                     = "invalid_recursion"                     // 42P19
	ERR_INVALID_FOREIGN_KEY                   = "invalid_foreign_key"                   // 42830
	ERR_INVALID_NAME                          = "invalid_name"                          // 42602
	ERR_NAME_TOO_LONG                         = "name_too_long"                         // 42622
	ERR_RESERVED_NAME                         = "reserved_name"                         // 42939
	ERR_DATATYPE_MISMATCH                     = "datatype_mismatch"                     // 42804
	ERR_INDETERMINATE_DATATYPE                = "indeterminate_datatype"                // 42P18
	ERR_COLLATION_MISMATCH                    = "collation_mismatch"                    // 42P21
	ERR_INDETERMINATE_COLLATION               = "indeterminate_collation"               // 42P22
	ERR_WRONG_OBJECT_TYPE                     = "wrong_object_type"                     // 42809
	ERR_GENERATED_ALWAYS                      = "generated_always"                      // 428C9
	ERR_UNDEFINED_COLUMN                      = "undefined_column"                      // 42703
	ERR_UNDEFINED_FUNCTION                    = "undefined_function"                    // 42883
	ERR_UNDEFINED_TABLE                       = "undefined_table"                       // 42P01
	ERR_UNDEFINED_PARAMETER                   = "undefined_parameter"                   // 42P02
	ERR_UNDEFINED_OBJECT                      = "undefined_object"                      // 42704
	ERR_DUPLICATE_COLUMN                      = "duplicate_column"                      // 42701
	ERR_DUPLICATE_CURSOR                      = "duplicate_cursor"                      // 42P03
	ERR_DUPLICATE_DATABASE                    = "duplicate_database"                    // 42P04
	ERR_DUPLICATE_FUNCTION                    = "duplicate_function"                    // 42723
	ERR_DUPLICATE_PREPARED_STATEMENT          = "duplicate_prepared_statement"          // 42P05
	ERR_DUPLICATE_SCHEMA                      = "duplicate_schema"                      // 42P06
	ERR_DUPLICATE_TABLE                       = "duplicate_table"                       // 42P07
	ERR_DUPLICATE_ALIAS                       = "duplicate_alias"                       // 42712
	ERR_DUPLICATE_OBJECT                      = "duplicate_object"                      // 42710
	ERR_AMBIGUOUS_COLUMN                      = "ambiguous_column"                      // 42702
	ERR_AMBIGUOUS_FUNCTION                    = "ambiguous_function"                    // 42725
	ERR_AMBIGUOUS_PARAMETER                   = "ambiguous_parameter"                   // 42P08
	ERR_AMBIGUOUS_ALIAS                       = "ambiguous_alias"                       // 42P09
	ERR_INVALID_COLUMN_REFERENCE              = "invalid_column_reference"              // 42P10
	ERR_INVALID_COLUMN_DEFINITION             = "invalid_column_definition"             // 42611
	ERR_INVALID_CURSOR_DEFINITION             = "invalid_cursor_definition"             // 42P11
	ERR_INVALID_DATABASE_DEFINITION           = "invalid_database_definition"           // 42P12
	ERR_INVALID_FUNCTION_DEFINITION           = "invalid_function_definition"           // 42P13
	ERR_INVALID_PREPARED_STATEMENT_DEFINITION = "invalid_prepared_statement_definition" // 42P14
	ERR_INVALID_SCHEMA_DEFINITION             = "invalid_schema_definition"             // 42P15
	ERR_INVALID_TABLE_DEFINITION              = "invalid_table_definition"              // 42P16
	ERR_INVALID_OBJECT_DEFINITION             = "invalid_object_definition"             // 42P17

	// Class 44 — WITH CHECK OPTION Violation
	ERR_WITH_CHECK_OPTION_VIOLATION = "with_check_option_violation" // 44000

	// Class 53 — Insufficient Resources
	ERR_INSUFFICIENT_RESOURCES       = "insufficient_resources"       // 53000
	ERR_DISK_FULL                    = "disk_full"                    // 53100
	ERR_OUT_OF_MEMORY                = "out_of_memory"                // 53200
	ERR_TOO_MANY_CONNECTIONS         = "too_many_connections"         // 53300
	ERR_CONFIGURATION_LIMIT_EXCEEDED = "configuration_limit_exceeded" // 53400

	// Class 54 — Program Limit Exceeded
	ERR_PROGRAM_LIMIT_EXCEEDED = "program_limit_exceeded" // 54000
	ERR_STATEMENT_TOO_COMPLEX  = "statement_too_complex"  // 54001
	ERR_TOO_MANY_COLUMNS       = "too_many_columns"       // 54011
	ERR_TOO_MANY_ARGUMENTS     = "too_many_arguments"     // 54023

	// Class 55 — Object Not In Prerequisite State
	ERR_OBJECT_NOT_IN_PREREQUISITE_STATE = "object_not_in_prerequisite_state" // 55000
	ERR_OBJECT_IN_USE                    = "object_in_use"                    // 55006
	ERR_CANT_CHANGE_RUNTIME_PARAM        = "cant_change_runtime_param"        // 55P02
	ERR_LOCK_NOT_AVAILABLE               = "lock_not_available"               // 55P03
	ERR_UNSAFE_NEW_ENUM_VALUE_USAGE      = "unsafe_new_enum_value_usage"      // 55P04

	// Class 57 — Operator Intervention
	ERR_OPERATOR_INTERVENTION = "operator_intervention" // 57000
	ERR_QUERY_CANCELED        = "query_canceled"        // 57014
	ERR_ADMIN_SHUTDOWN        = "admin_shutdown"        // 57P01
	ERR_CRASH_SHUTDOWN        = "crash_shutdown"        // 57P02
	ERR_CANNOT_CONNECT_NOW    = "cannot_connect_now"    // 57P03
	ERR_DATABASE_DROPPED      = "database_dropped"      // 57P04
	ERR_IDLE_SESSION_TIMEOUT  = "idle_session_timeout"  // 57P05

	// Class 58 — System Error (errors external to PostgreSQL itself)
	ERR_SYSTEM_ERROR   = "system_error"   // 58000
	ERR_IO_ERROR       = "io_error"       // 58030
	ERR_UNDEFINED_FILE = "undefined_file" // 58P01
	ERR_DUPLICATE_FILE = "duplicate_file" // 58P02

	// Class 72 — Snapshot Failure
	ERR_SNAPSHOT_TOO_OLD = "snapshot_too_old" // 72000

	// Class F0 — Configuration File Error
	ERR_CONFIG_FILE_ERROR = "config_file_error" // F0000
	ERR_LOCK_FILE_EXISTS  = "lock_file_exists"  // F0001

	// Class HV — Foreign Data Wrapper Error (SQL/MED)
	ERR_FDW_ERROR                                  = "fdw_error"                                  // HV000
	ERR_FDW_COLUMN_NAME_NOT_FOUND                  = "fdw_column_name_not_found"                  // HV005
	ERR_FDW_DYNAMIC_PARAMETER_VALUE_NEEDED         = "fdw_dynamic_parameter_value_needed"         // HV002
	ERR_FDW_FUNCTION_SEQUENCE_ERROR                = "fdw_function_sequence_error"                // HV010
	ERR_FDW_INCONSISTENT_DESCRIPTOR_INFORMATION    = "fdw_inconsistent_descriptor_information"    // HV021
	ERR_FDW_INVALID_ATTRIBUTE_VALUE                = "fdw_invalid_attribute_value"                // HV024
	ERR_FDW_INVALID_COLUMN_NAME                    = "fdw_invalid_column_name"                    // HV007
	ERR_FDW_INVALID_COLUMN_NUMBER                  = "fdw_invalid_column_number"                  // HV008
	ERR_FDW_INVALID_DATA_TYPE                      = "fdw_invalid_data_type"                      // HV004
	ERR_FDW_INVALID_DATA_TYPE_DESCRIPTORS          = "fdw_invalid_data_type_descriptors"          // HV006
	ERR_FDW_INVALID_DESCRIPTOR_FIELD_IDENTIFIER    = "fdw_invalid_descriptor_field_identifier"    // HV091
	ERR_FDW_INVALID_HANDLE                         = "fdw_invalid_handle"                         // HV00B
	ERR_FDW_INVALID_OPTION_INDEX                   = "fdw_invalid_option_index"                   // HV00C
	ERR_FDW_INVALID_OPTION_NAME                    = "fdw_invalid_option_name"                    // HV00D
	ERR_FDW_INVALID_STRING_LENGTH_OR_BUFFER_LENGTH = "fdw_invalid_string_length_or_buffer_length" // HV090
	ERR_FDW_INVALID_STRING_FORMAT                  = "fdw_invalid_string_format"                  // HV00A
	ERR_FDW_INVALID_USE_OF_NULL_POINTER            = "fdw_invalid_use_of_null_pointer"            // HV009
	ERR_FDW_TOO_MANY_HANDLES                       = "fdw_too_many_handles"                       // HV014
	ERR_FDW_OUT_OF_MEMORY                          = "fdw_out_of_memory"                          // HV001
	ERR_FDW_NO_SCHEMAS                             = "fdw_no_schemas"                             // HV00P
	ERR_FDW_OPTION_NAME_NOT_FOUND                  = "fdw_option_name_not_found"                  // HV00J
	ERR_FDW_REPLY_HANDLE                           = "fdw_reply_handle"                           // HV00K
	ERR_FDW_SCHEMA_NOT_FOUND                       = "fdw_schema_not_found"                       // HV00Q
	ERR_FDW_TABLE_NOT_FOUND                        = "fdw_table_not_found"                        // HV00R
	ERR_FDW_UNABLE_TO_CREATE_EXECUTION             = "fdw_unable_to_create_execution"             // HV00L
	ERR_FDW_UNABLE_TO_CREATE_REPLY                 = "fdw_unable_to_create_reply"                 // HV00M
	ERR_FDW_UNABLE_TO_ESTABLISH_CONNECTION         = "fdw_unable_to_establish_connection"         // HV00N

	// Class P0 — PL/pgSQL Error
	ERR_PLPGSQL_ERROR   = "plpgsql_error"   // P0000
	ERR_RAISE_EXCEPTION = "raise_exception" // P0001
	ERR_NO_DATA_FOUND   = "no_data_found"   // P0002
	ERR_TOO_MANY_ROWS   = "too_many_rows"   // P0003
	ERR_ASSERT_FAILURE  = "assert_failure"  // P0004

	// Class XX — Internal Error
	ERR_INTERNAL_ERROR  = "internal_error"  // XX000
	ERR_DATA_CORRUPTED  = "data_corrupted"  // XX001
	ERR_INDEX_CORRUPTED = "index_corrupted" // XX002
)

var reposeErrors = map[string]any{
	"00000": http.ErrorResponse(ERR_SUCCESSFUL_COMPLETION),
	"42701": http.ErrorResponse(ERR_DUPLICATE_COLUMN),
	"XX000": http.ErrorResponse(ERR_INTERNAL_ERROR),
	"P0002": http.ErrorResponse(ERR_NO_DATA_FOUND),
	"54023": http.ErrorResponse(ERR_TOO_MANY_ARGUMENTS),
	"54011": http.ErrorResponse(ERR_TOO_MANY_COLUMNS),
}

func Error(err any) any {

	pgErr, ok := err.(*pgconn.PgError)

	if !ok {
		return err
	}

	key := pgErr.SQLState()

	code := reposeErrors[key]

	if code == nil {
		code = http.ErrorResponse(key)
	}

	return code
}
