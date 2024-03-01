package constant

const (
	//Error Handler
	ErrorBindingJson  = "ErrorBindingJson"
	ErrorMarshalJson  = "ErrorMarshalJson"
	ErrorFormatingCsv = "Pastikan format CSV yang di gunakan sesuai"

	MaximalFileSizeCSV = 5097152

	//log
	LogTimePrefix      = " [Time] = "
	LogMethodPrefix    = "[Method] = "
	LogPathPrefix      = " [Path] = "
	LogIPPrefix        = " [IP] = "
	LogStartPrefix     = "[Start]"
	LogStopPrefix      = "[Stop]"
	LogResponsePrefix  = " [Response] = "
	LogErrorPrefix     = " [Error] = "
	DateYYYYMMDDHHIISS = "2006-01-02 15:04:05"
	LogRequestPrefix   = " [Request] ="

	//variable
	TipeAksesKanwil   = "kanwil"
	TipeAksesPusat    = "pusat"
	TipeAksesArea     = "area"
	TipeAksesCabang   = "cabang"
	TipeAksesOutlet   = "outlet"
	ExpectedDelimeter = ','
)
