package sqllint

import (
	"bytes"
	"strings"

	"github.com/pingcap/parser/ast"

	"github.com/erda-project/erda/pkg/swagger/ddlconv"
)

var keywords = map[string]bool{
	"ACCESSIBLE":                    true,
	"ACCOUNT":                       true,
	"ACTION":                        true,
	"ADD":                           true,
	"AFTER":                         true,
	"AGAINST":                       true,
	"AGGREGATE":                     true,
	"ALGORITHM":                     true,
	"ALL":                           true,
	"ALTER":                         true,
	"ALWAYS":                        true,
	"ANALYZE":                       true,
	"AND":                           true,
	"ANY":                           true,
	"AS":                            true,
	"ASC":                           true,
	"ASCII":                         true,
	"ASENSITIVE":                    true,
	"AT":                            true,
	"AUTOEXTEND_SIZE":               true,
	"AUTO_INCREMENT":                true,
	"AVG":                           true,
	"AVG_ROW_LENGTH":                true,
	"BACKUP":                        true,
	"BEFORE":                        true,
	"BEGIN":                         true,
	"BETWEEN":                       true,
	"BIGINT":                        true,
	"BINARY":                        true,
	"BINLOG":                        true,
	"BIT":                           true,
	"BLOB":                          true,
	"BLOCK":                         true,
	"BOOL":                          true,
	"BOOLEAN":                       true,
	"BOTH":                          true,
	"BTREE":                         true,
	"BY":                            true,
	"BYTE":                          true,
	"CACHE":                         true,
	"CALL":                          true,
	"CASCADE":                       true,
	"CASCADED":                      true,
	"CASE":                          true,
	"CATALOG_NAME":                  true,
	"CHAIN":                         true,
	"CHANGE":                        true,
	"CHANGED":                       true,
	"CHANNEL":                       true,
	"CHAR":                          true,
	"CHARACTER":                     true,
	"CHARSET":                       true,
	"CHECK":                         true,
	"CHECKSUM":                      true,
	"CIPHER":                        true,
	"CLASS_ORIGIN":                  true,
	"CLIENT":                        true,
	"CLOSE":                         true,
	"COALESCE":                      true,
	"CODE":                          true,
	"COLLATE":                       true,
	"COLLATION":                     true,
	"COLUMN":                        true,
	"COLUMNS":                       true,
	"COLUMN_FORMAT":                 true,
	"COLUMN_NAME":                   true,
	"COMMENT":                       true,
	"COMMIT":                        true,
	"COMMITTED":                     true,
	"COMPACT":                       true,
	"COMPLETION":                    true,
	"COMPRESSED":                    true,
	"COMPRESSION":                   true,
	"CONCURRENT":                    true,
	"CONDITION":                     true,
	"CONNECTION":                    true,
	"CONSISTENT":                    true,
	"CONSTRAINT":                    true,
	"CONSTRAINT_CATALOG":            true,
	"CONSTRAINT_NAME":               true,
	"CONSTRAINT_SCHEMA":             true,
	"CONTAINS":                      true,
	"CONTEXT":                       true,
	"CONTINUE":                      true,
	"CONVERT":                       true,
	"CPU":                           true,
	"CREATE":                        true,
	"CROSS":                         true,
	"CUBE":                          true,
	"CURRENT":                       true,
	"CURRENT_DATE":                  true,
	"CURRENT_TIME":                  true,
	"CURRENT_TIMESTAMP":             true,
	"CURRENT_USER":                  true,
	"CURSOR":                        true,
	"CURSOR_NAME":                   true,
	"DATA":                          true,
	"DATABASE":                      true,
	"DATABASES":                     true,
	"DATAFILE":                      true,
	"DATE":                          true,
	"DATETIME":                      true,
	"DAY":                           true,
	"DAY_HOUR":                      true,
	"DAY_MICROSECOND":               true,
	"DAY_MINUTE":                    true,
	"DAY_SECOND":                    true,
	"DEALLOCATE":                    true,
	"DEC":                           true,
	"DECIMAL":                       true,
	"DECLARE":                       true,
	"DEFAULT":                       true,
	"DEFAULT_AUTH":                  true,
	"DEFINER":                       true,
	"DELAYED":                       true,
	"DELAY_KEY_WRITE":               true,
	"DELETE":                        true,
	"DESC":                          true,
	"DESCRIBE":                      true,
	"DETERMINISTIC":                 true,
	"DIAGNOSTICS":                   true,
	"DIRECTORY":                     true,
	"DISABLE":                       true,
	"DISCARD":                       true,
	"DISK":                          true,
	"DISTINCT":                      true,
	"DISTINCTROW":                   true,
	"DIV":                           true,
	"DO":                            true,
	"DOUBLE":                        true,
	"DROP":                          true,
	"DUAL":                          true,
	"DUMPFILE":                      true,
	"DUPLICATE":                     true,
	"DYNAMIC":                       true,
	"EACH":                          true,
	"ELSE":                          true,
	"ELSEIF":                        true,
	"ENABLE":                        true,
	"ENCLOSED":                      true,
	"ENCRYPTION":                    true,
	"END":                           true,
	"ENDS":                          true,
	"ENGINE":                        true,
	"ENGINES":                       true,
	"ENUM":                          true,
	"ERROR":                         true,
	"ERRORS":                        true,
	"ESCAPE":                        true,
	"ESCAPED":                       true,
	"EVENT":                         true,
	"EVENTS":                        true,
	"EVERY":                         true,
	"EXCHANGE":                      true,
	"EXECUTE":                       true,
	"EXISTS":                        true,
	"EXIT":                          true,
	"EXPANSION":                     true,
	"EXPIRE":                        true,
	"EXPLAIN":                       true,
	"EXPORT":                        true,
	"EXTENDED":                      true,
	"EXTENT_SIZE":                   true,
	"FALSE":                         true,
	"FAST":                          true,
	"FAULTS":                        true,
	"FETCH":                         true,
	"FIELDS":                        true,
	"FILE":                          true,
	"FILE_BLOCK_SIZE":               true,
	"FILTER":                        true,
	"FIRST":                         true,
	"FIXED":                         true,
	"FLOAT":                         true,
	"FLOAT4":                        true,
	"FLOAT8":                        true,
	"FLUSH":                         true,
	"FOLLOWS":                       true,
	"FOR":                           true,
	"FORCE":                         true,
	"FOREIGN":                       true,
	"FORMAT":                        true,
	"FOUND":                         true,
	"FROM":                          true,
	"FULL":                          true,
	"FULLTEXT":                      true,
	"FUNCTION":                      true,
	"GENERAL":                       true,
	"GENERATED":                     true,
	"GEOMETRY":                      true,
	"GEOMETRYCOLLECTION":            true,
	"GET":                           true,
	"GET_FORMAT":                    true,
	"GLOBAL":                        true,
	"GRANT":                         true,
	"GRANTS":                        true,
	"GROUP":                         true,
	"GROUP_REPLICATION":             true,
	"HANDLER":                       true,
	"HASH":                          true,
	"HAVING":                        true,
	"HELP":                          true,
	"HIGH_PRIORITY":                 true,
	"HOST":                          true,
	"HOSTS":                         true,
	"HOUR":                          true,
	"HOUR_MICROSECOND":              true,
	"HOUR_MINUTE":                   true,
	"HOUR_SECOND":                   true,
	"IDENTIFIED":                    true,
	"IF":                            true,
	"IGNORE":                        true,
	"IGNORE_SERVER_IDS":             true,
	"IMPORT":                        true,
	"IN":                            true,
	"INDEX":                         true,
	"INDEXES":                       true,
	"INFILE":                        true,
	"INITIAL_SIZE":                  true,
	"INNER":                         true,
	"INOUT":                         true,
	"INSENSITIVE":                   true,
	"INSERT":                        true,
	"INSERT_METHOD":                 true,
	"INSTALL":                       true,
	"INSTANCE":                      true,
	"INT":                           true,
	"INT1":                          true,
	"INT2":                          true,
	"INT3":                          true,
	"INT4":                          true,
	"INT8":                          true,
	"INTEGER":                       true,
	"INTERVAL":                      true,
	"INTO":                          true,
	"INVOKER":                       true,
	"IO":                            true,
	"IO_AFTER_GTIDS":                true,
	"IO_BEFORE_GTIDS":               true,
	"IO_THREAD":                     true,
	"IPC":                           true,
	"IS":                            true,
	"ISOLATION":                     true,
	"ISSUER":                        true,
	"ITERATE":                       true,
	"JOIN":                          true,
	"JSON":                          true,
	"KEY":                           true,
	"KEYS":                          true,
	"KEY_BLOCK_SIZE":                true,
	"KILL":                          true,
	"LANGUAGE":                      true,
	"LAST":                          true,
	"LEADING":                       true,
	"LEAVE":                         true,
	"LEAVES":                        true,
	"LEFT":                          true,
	"LESS":                          true,
	"LEVEL":                         true,
	"LIKE":                          true,
	"LIMIT":                         true,
	"LINEAR":                        true,
	"LINES":                         true,
	"LINESTRING":                    true,
	"LIST":                          true,
	"LOAD":                          true,
	"LOCAL":                         true,
	"LOCALTIME":                     true,
	"LOCALTIMESTAMP":                true,
	"LOCK":                          true,
	"LOCKS":                         true,
	"LOGFILE":                       true,
	"LOGS":                          true,
	"LONG":                          true,
	"LONGBLOB":                      true,
	"LONGTEXT":                      true,
	"LOOP":                          true,
	"LOW_PRIORITY":                  true,
	"MASTER":                        true,
	"MASTER_AUTO_POSITION":          true,
	"MASTER_BIND":                   true,
	"MASTER_CONNECT_RETRY":          true,
	"MASTER_DELAY":                  true,
	"MASTER_HEARTBEAT_PERIOD":       true,
	"MASTER_HOST":                   true,
	"MASTER_LOG_FILE":               true,
	"MASTER_LOG_POS":                true,
	"MASTER_PASSWORD":               true,
	"MASTER_PORT":                   true,
	"MASTER_PUBLIC_KEY_PATH":        true,
	"MASTER_RETRY_COUNT":            true,
	"MASTER_SERVER_ID":              true,
	"MASTER_SSL":                    true,
	"MASTER_SSL_CA":                 true,
	"MASTER_SSL_CAPATH":             true,
	"MASTER_SSL_CERT":               true,
	"MASTER_SSL_CIPHER":             true,
	"MASTER_SSL_CRL":                true,
	"MASTER_SSL_CRLPATH":            true,
	"MASTER_SSL_KEY":                true,
	"MASTER_SSL_VERIFY_SERVER_CERT": true,
	"MASTER_TLS_CIPHERSUITES":       true,
	"MASTER_TLS_VERSION":            true,
	"MASTER_USER":                   true,
	"MASTER_ZSTD_COMPRESSION_LEVEL": true,
	"MATCH":                         true,
	"MAXVALUE":                      true,
	"MAX_CONNECTIONS_PER_HOUR":      true,
	"MAX_QUERIES_PER_HOUR":          true,
	"MAX_ROWS":                      true,
	"MAX_SIZE":                      true,
	"MAX_UPDATES_PER_HOUR":          true,
	"MAX_USER_CONNECTIONS":          true,
	"MEDIUM":                        true,
	"MEDIUMBLOB":                    true,
	"MEDIUMINT":                     true,
	"MEDIUMTEXT":                    true,
	"MEMBER":                        true,
	"MEMORY":                        true,
	"MERGE":                         true,
	"MESSAGE_TEXT":                  true,
	"MICROSECOND":                   true,
	"MIDDLEINT":                     true,
	"MIGRATE":                       true,
	"MINUTE":                        true,
	"MINUTE_MICROSECOND":            true,
	"MINUTE_SECOND":                 true,
	"MIN_ROWS":                      true,
	"MOD":                           true,
	"MODE":                          true,
	"MODIFIES":                      true,
	"MODIFY":                        true,
	"MONTH":                         true,
	"MULTILINESTRING":               true,
	"MULTIPOINT":                    true,
	"MULTIPOLYGON":                  true,
	"MUTEX":                         true,
	"MYSQL_ERRNO":                   true,
	"NAME":                          true,
	"NAMES":                         true,
	"NATIONAL":                      true,
	"NATURAL":                       true,
	"NCHAR":                         true,
	"NDB":                           true,
	"NDBCLUSTER":                    true,
	"NEVER":                         true,
	"NEW":                           true,
	"NEXT":                          true,
	"NO":                            true,
	"NODEGROUP":                     true,
	"NONE":                          true,
	"NOT":                           true,
	"NO_WAIT":                       true,
	"NO_WRITE_TO_BINLOG":            true,
	"NULL":                          true,
	"NUMBER":                        true,
	"NUMERIC":                       true,
	"NVARCHAR":                      true,
	"OFFSET":                        true,
	"ON":                            true,
	"ONE":                           true,
	"ONLY":                          true,
	"OPEN":                          true,
	"OPTIMIZE":                      true,
	"OPTIMIZER_COSTS":               true,
	"OPTION":                        true,
	"OPTIONALLY":                    true,
	"OPTIONS":                       true,
	"OR":                            true,
	"ORDER":                         true,
	"OUT":                           true,
	"OUTER":                         true,
	"OUTFILE":                       true,
	"OWNER":                         true,
	"PACK_KEYS":                     true,
	"PAGE":                          true,
	"PARSER":                        true,
	"PARTIAL":                       true,
	"PARTITION":                     true,
	"PARTITIONING":                  true,
	"PARTITIONS":                    true,
	"PASSWORD":                      true,
	"PHASE":                         true,
	"PLUGIN":                        true,
	"PLUGINS":                       true,
	"PLUGIN_DIR":                    true,
	"POINT":                         true,
	"POLYGON":                       true,
	"PORT":                          true,
	"PRECEDES":                      true,
	"PRECISION":                     true,
	"PREPARE":                       true,
	"PRESERVE":                      true,
	"PREV":                          true,
	"PRIMARY":                       true,
	"PRIVILEGES":                    true,
	"PROCEDURE":                     true,
	"PROCESS":                       true,
	"PROCESSLIST":                   true,
	"PROFILE":                       true,
	"PROFILES":                      true,
	"PROXY":                         true,
	"PURGE":                         true,
	"QUARTER":                       true,
	"QUERY":                         true,
	"QUICK":                         true,
	"RANDOM":                        true,
	"RANGE":                         true,
	"RANK":                          true,
	"READ":                          true,
	"READS":                         true,
	"READ_ONLY":                     true,
	"READ_WRITE":                    true,
	"REAL":                          true,
	"REBUILD":                       true,
	"RECOVER":                       true,
	"RECURSIVE":                     true,
	"REDOFILE":                      true,
	"REDO_BUFFER_SIZE":              true,
	"REDUNDANT":                     true,
	"REFERENCE":                     true,
	"REFERENCES":                    true,
	"REGEXP":                        true,
	"RELAY":                         true,
	"RELAYLOG":                      true,
	"RELAY_LOG_FILE":                true,
	"RELAY_LOG_POS":                 true,
	"RELAY_THREAD":                  true,
	"RELEASE":                       true,
	"RELOAD":                        true,
	"REMOTE":                        true,
	"REMOVE":                        true,
	"RENAME":                        true,
	"REORGANIZE":                    true,
	"REPAIR":                        true,
	"REPEAT":                        true,
	"REPEATABLE":                    true,
	"REPLACE":                       true,
	"REPLICATE_DO_DB":               true,
	"REPLICATE_DO_TABLE":            true,
	"REPLICATE_IGNORE_DB":           true,
	"REPLICATE_IGNORE_TABLE":        true,
	"REPLICATE_REWRITE_DB":          true,
	"REPLICATE_WILD_DO_TABLE":       true,
	"REPLICATE_WILD_IGNORE_TABLE":   true,
	"REPLICATION":                   true,
	"REQUIRE":                       true,
	"RESET":                         true,
	"RESIGNAL":                      true,
	"RESTORE":                       true,
	"RESTRICT":                      true,
	"RESUME":                        true,
	"RETURN":                        true,
	"RETURNED_SQLSTATE":             true,
	"RETURNS":                       true,
	"REVERSE":                       true,
	"REVOKE":                        true,
	"RIGHT":                         true,
	"RLIKE":                         true,
	"ROLLBACK":                      true,
	"ROLLUP":                        true,
	"ROTATE":                        true,
	"ROUTINE":                       true,
	"ROW":                           true,
	"ROWS":                          true,
	"ROW_COUNT":                     true,
	"ROW_FORMAT":                    true,
	"RTREE":                         true,
	"SAVEPOINT":                     true,
	"SCHEDULE":                      true,
	"SCHEMA":                        true,
	"SCHEMAS":                       true,
	"SCHEMA_NAME":                   true,
	"SECOND":                        true,
	"SECONDARY":                     true,
	"SECONDARY_ENGINE":              true,
	"SECONDARY_LOAD":                true,
	"SECONDARY_UNLOAD":              true,
	"SECOND_MICROSECOND":            true,
	"SECURITY":                      true,
	"SELECT":                        true,
	"SENSITIVE":                     true,
	"SEPARATOR":                     true,
	"SERIAL":                        true,
	"SERIALIZABLE":                  true,
	"SERVER":                        true,
	"SESSION":                       true,
	"SET":                           true,
	"SHARE":                         true,
	"SHOW":                          true,
	"SHUTDOWN":                      true,
	"SIGNAL":                        true,
	"SIGNED":                        true,
	"SIMPLE":                        true,
	"SKIP":                          true,
	"SLAVE":                         true,
	"SLOW":                          true,
	"SMALLINT":                      true,
	"SNAPSHOT":                      true,
	"SOCKET":                        true,
	"SOME":                          true,
	"SONAME":                        true,
	"SOUNDS":                        true,
	"SOURCE":                        true,
	"SPATIAL":                       true,
	"SPECIFIC":                      true,
	"SQL":                           true,
	"SQLEXCEPTION":                  true,
	"SQLSTATE":                      true,
	"SQLWARNING":                    true,
	"SQL_AFTER_GTIDS":               true,
	"SQL_AFTER_MTS_GAPS":            true,
	"SQL_BEFORE_GTIDS":              true,
	"SQL_BIG_RESULT":                true,
	"SQL_BUFFER_RESULT":             true,
	"SQL_CALC_FOUND_ROWS":           true,
	"SQL_NO_CACHE":                  true,
	"SQL_SMALL_RESULT":              true,
	"SQL_THREAD":                    true,
	"SQL_TSI_DAY":                   true,
	"SQL_TSI_HOUR":                  true,
	"SQL_TSI_MINUTE":                true,
	"SQL_TSI_MONTH":                 true,
	"SQL_TSI_QUARTER":               true,
	"SQL_TSI_SECOND":                true,
	"SQL_TSI_WEEK":                  true,
	"SQL_TSI_YEAR":                  true,
	"SSL":                           true,
	"STACKED":                       true,
	"START":                         true,
	"STARTING":                      true,
	"STARTS":                        true,
	"STATS_AUTO_RECALC":             true,
	"STATS_PERSISTENT":              true,
	"STATS_SAMPLE_PAGES":            true,
	"STATUS":                        true,
	"STOP":                          true,
	"STORAGE":                       true,
	"STORED":                        true,
	"STRAIGHT_JOIN":                 true,
	"STRING":                        true,
	"SUBCLASS_ORIGIN":               true,
	"SUBJECT":                       true,
	"SUBPARTITION":                  true,
	"SUBPARTITIONS":                 true,
	"SUPER":                         true,
	"SUSPEND":                       true,
	"SWAPS":                         true,
	"SWITCHES":                      true,
	"TABLE":                         true,
	"TABLES":                        true,
	"TABLESPACE":                    true,
	"TABLE_CHECKSUM":                true,
	"TABLE_NAME":                    true,
	"TEMPORARY":                     true,
	"TEMPTABLE":                     true,
	"TERMINATED":                    true,
	"TEXT":                          true,
	"THAN":                          true,
	"THEN":                          true,
	"TIME":                          true,
	"TIMESTAMP":                     true,
	"TIMESTAMPADD":                  true,
	"TIMESTAMPDIFF":                 true,
	"TINYBLOB":                      true,
	"TINYINT":                       true,
	"TINYTEXT":                      true,
	"TO":                            true,
	"TRAILING":                      true,
	"TRANSACTION":                   true,
	"TRIGGER":                       true,
	"TRIGGERS":                      true,
	"TRUE":                          true,
	"TRUNCATE":                      true,
	"TYPE":                          true,
	"TYPES":                         true,
	"UNCOMMITTED":                   true,
	"UNDEFINED":                     true,
	"UNDO":                          true,
	"UNDOFILE":                      true,
	"UNDO_BUFFER_SIZE":              true,
	"UNICODE":                       true,
	"UNINSTALL":                     true,
	"UNION":                         true,
	"UNIQUE":                        true,
	"UNKNOWN":                       true,
	"UNLOCK":                        true,
	"UNSIGNED":                      true,
	"UNTIL":                         true,
	"UPDATE":                        true,
	"UPGRADE":                       true,
	"USAGE":                         true,
	"USE":                           true,
	"USER":                          true,
	"USER_RESOURCES":                true,
	"USE_FRM":                       true,
	"USING":                         true,
	"UTC_DATE":                      true,
	"UTC_TIME":                      true,
	"UTC_TIMESTAMP":                 true,
	"VALIDATION":                    true,
	"VALUE":                         true,
	"VALUES":                        true,
	"VARBINARY":                     true,
	"VARCHAR":                       true,
	"VARCHARACTER":                  true,
	"VARIABLES":                     true,
	"VARYING":                       true,
	"VIEW":                          true,
	"VIRTUAL":                       true,
	"WAIT":                          true,
	"WARNINGS":                      true,
	"WEEK":                          true,
	"WEIGHT_STRING":                 true,
	"WHEN":                          true,
	"WHERE":                         true,
	"WHILE":                         true,
	"WITH":                          true,
	"WITHOUT":                       true,
	"WORK":                          true,
	"WRAPPER":                       true,
	"WRITE":                         true,
	"X509":                          true,
	"XA":                            true,
	"XID":                           true,
	"XML":                           true,
	"XOR":                           true,
	"YEAR":                          true,
	"YEAR_MONTH":                    true,
	"ZEROFILL":                      true,
	"ACTIVE":                        true,
	"ADMIN":                         true,
	"ARRAY":                         true,
	"BUCKETS":                       true,
	"CLONE":                         true,
	"COMPONENT":                     true,
	"CUME_DIST":                     true,
	"DEFINITION":                    true,
	"DENSE_RANK":                    true,
	"DESCRIPTION":                   true,
	"EMPTY":                         true,
	"ENFORCED":                      true,
	"EXCEPT":                        true,
	"EXCLUDE":                       true,
	"FAILED_LOGIN_ATTEMPTS":         true,
	"FIRST_VALUE":                   true,
	"FOLLOWING":                     true,
	"GEOMCOLLECTION":                true,
	"GET_MASTER_PUBLIC_KEY":         true,
	"GROUPING":                      true,
	"GROUPS":                        true,
	"HISTOGRAM":                     true,
	"HISTORY":                       true,
	"INACTIVE":                      true,
	"INVISIBLE":                     true,
	"JSON_TABLE":                    true,
	"JSON_VALUE":                    true,
	"LAG":                           true,
	"LAST_VALUE":                    true,
	"LATERAL":                       true,
	"LEAD":                          true,
	"LOCKED":                        true,
	"MASTER_COMPRESSION_ALGORITHMS": true,
	"NESTED":                        true,
	"NETWORK_NAMESPACE":             true,
	"NOWAIT":                        true,
	"NTH_VALUE":                     true,
	"NTILE":                         true,
	"NULLS":                         true,
	"OF":                            true,
	"OFF":                           true,
	"OJ":                            true,
	"OLD":                           true,
	"OPTIONAL":                      true,
	"ORDINALITY":                    true,
	"ORGANIZATION":                  true,
	"OTHERS":                        true,
	"OVER":                          true,
	"PASSWORD_LOCK_TIME":            true,
	"PATH":                          true,
	"PERCENT_RANK":                  true,
	"PERSIST":                       true,
	"PERSIST_ONLY":                  true,
	"PRECEDING":                     true,
	"PRIVILEGE_CHECKS_USER":         true,
	"REQUIRE_ROW_FORMAT":            true,
	"RESOURCE":                      true,
	"RESPECT":                       true,
	"RESTART":                       true,
	"RETAIN":                        true,
	"RETURNING":                     true,
	"REUSE":                         true,
	"ROLE":                          true,
	"ROW_NUMBER":                    true,
	"SRID":                          true,
	"STREAM":                        true,
	"SYSTEM":                        true,
	"THREAD_PRIORITY":               true,
	"TIES":                          true,
	"UNBOUNDED":                     true,
	"VCPU":                          true,
	"VISIBLE":                       true,
	"WINDOW":                        true,
	"ANALYSE":                       true,
	"DES_KEY_FILE":                  true,
	"PARSE_GCOL_EXPR":               true,
	"SQL_CACHE":                     true,
}

type KeywordsLinter struct {
	script Script
	err    error
	text   string
}

func NewKeywordsLinter(script Script) Rule {
	return &KeywordsLinter{script: script}
}

func (l *KeywordsLinter) Enter(in ast.Node) (ast.Node, bool) {
	if l.text == "" || in.Text() != "" {
		l.text = in.Text()
	}

	switch in.(type) {
	case *ast.CreateTableStmt:
		stmt := in.(*ast.CreateTableStmt)
		name := ddlconv.ExtractCreateName(stmt)
		if name == "" {
			return in, false
		}
		if _, ok := keywords[strings.ToUpper(name)]; ok {
			l.err = NewLintError(l.script, l.text, "表名不合法: 不得使用 MySQL 关键字或保留字为表名",
				func(_ []byte) bool {
					return false
				})
			return in, true
		}
	case *ast.ColumnDef:
		col := in.(*ast.ColumnDef)
		name := ddlconv.ExtractColName(col)
		if name == "" {
			return in, false
		}
		if _, ok := keywords[strings.ToUpper(name)]; ok {
			l.err = NewLintError(l.script, l.text, "字段名不合法: 使用了 MySQL 关键字或保留字",
				func(line []byte) bool {
					return bytes.Contains(bytes.ToLower(line), bytes.ToLower([]byte(name)))
				})
			return in, true
		}
	default:
		return in, false
	}

	return in, false
}

func (l *KeywordsLinter) Leave(in ast.Node) (ast.Node, bool) {
	return in, l.err == nil
}

func (l *KeywordsLinter) Error() error {
	return l.err
}
