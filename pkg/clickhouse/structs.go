package clickhouse

import (
	"time"
)

type ShardBackupType string

const (
	ShardBackupFull   = "full"
	ShardBackupNone   = "none"
	ShardBackupSchema = "schema-only"
)

// Table - ClickHouse table struct
type Table struct {
	// common fields for all `clickhouse-server` versions
	Database string `ch:"database"`
	Name     string `ch:"name"`
	Engine   string `ch:"engine"`
	// fields depends on `clickhouse-server` version
	DataPath         string   `ch:"data_path"`
	DataPaths        []string `ch:"data_paths"`
	UUID             string   `ch:"uuid"`
	CreateTableQuery string   `ch:"create_table_query"`
	TotalBytes       uint64   `ch:"total_bytes"`
	Skip             bool
	BackupType       ShardBackupType
}

// IsSystemTablesFieldPresent - ClickHouse `system.tables` varius field flags
type IsSystemTablesFieldPresent struct {
	IsDataPathPresent         uint64 `ch:"is_data_path_present"`
	IsDataPathsPresent        uint64 `ch:"is_data_paths_present"`
	IsUUIDPresent             uint64 `ch:"is_uuid_present"`
	IsCreateTableQueryPresent uint64 `ch:"is_create_table_query_present"`
	IsTotalBytesPresent       uint64 `ch:"is_total_bytes_present"`
}

type Disk struct {
	Name            string   `ch:"name"`
	Path            string   `ch:"path"`
	Type            string   `ch:"type"`
	FreeSpace       uint64   `ch:"free_space"`
	StoragePolicies []string `ch:"storage_policies"`
	IsBackup        bool
}

// Database - Clickhouse system.databases struct
type Database struct {
	Name   string `ch:"name"`
	Engine string `ch:"engine"`
	Query  string `ch:"query"`
}

// Function - Clickhouse system.functions struct
type Function struct {
	Name        string `ch:"name"`
	CreateQuery string `ch:"create_query"`
}

// Macro - info from system.macros
type Macro struct {
	Macro        string `ch:"macro"`
	Substitution string `ch:"substitution"`
}

// SystemBackups - info from system.backups, universal for all versions
type SystemBackups struct {
	Id                string            `ch:"id"`
	UUID              string            `ch:"uuid"`
	QueryId           string            `ch:"query_id"`
	BackupName        string            `ch:"backup_name"`
	BaseBackupName    string            `ch:"base_backup_name"`
	Name              string            `ch:"name"`
	Status            string            `ch:"status"`
	StatusChangedTime time.Time         `ch:"status_changed_time"`
	Error             string            `ch:"error"`
	Internal          bool              `ch:"internal"`
	StartTime         time.Time         `ch:"start_time"`
	EndTime           time.Time         `ch:"end_time"`
	CompressedSize    uint64            `ch:"compressed_size"`
	UncompressedSize  uint64            `ch:"uncompressed_size"`
	NumFiles          uint64            `ch:"num_files"`
	TotalSize         uint64            `ch:"total_size"`
	NumEntries        uint64            `ch:"num_entries"`
	FilesRead         uint64            `ch:"files_read"`
	BytesRead         uint64            `ch:"bytes_read"`
	ProfileEvents     map[string]uint64 `ch:"ProfileEvents"`
}

// ColumnDataTypes - info from system.parts_columns
type ColumnDataTypes struct {
	Column string   `ch:"column"`
	Types  []string `ch:"uniq_types"`
}

// BackupDataSize - info from system.parts or system.tables when embedded BACKUP statement return zero size
type BackupDataSize struct {
	Size uint64 `ch:"backup_data_size"`
}

type UserDirectory struct {
	Name string `ch:"name"`
}

type RBACObject struct {
	Id   string `ch:"id"`
	Name string `ch:"name"`
}
