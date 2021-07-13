package define

const (
	C_EMPTY = iota
	C_ZHENG
	C_QIN
	C_QI
	C_LU
	C_WU
	C_YUE
	C_CHU
	C_RONG
	C_DI
	C_SONG
)

const (
	B_C_JUN = iota
	B_C_ZU
	B_C_CHEN
	B_C_CE
	B_C_MENG
	B_C_DU
	B_C_DI
	B_C_QI
)

const (
	DESTINY_NUM  = 1
	BASE_TAX     = 3
	TAX_PER_TILE = 1
)

const (
	NOT_EXEC_BLAKC_OR_WHITE = iota
	BLACK_STAGE
	WHITE_STAGE
)

const (
	STAGE_XINGZHENG = iota
	STAGE_DIAOQIAN
	STAGE_XIUZHENG
)

// 武器代码
const (
	WEAPON_BU = iota
	WEAPON_GONG
	WEAPON_QI
	WEAPON_CHE
)

// 武器价格
const (
	PRICE_WEAPON_BU   = 3
	PRICE_WEAPON_GONG = 4
	PRICE_WEAPON_QI   = 5
	PRICE_WEAPON_CHE  = 6
)

const (
	M_EMPTY = iota
	M_SHAN
	M_SHUI
	M_DAMO
	M_NONGTIAN
)

const (
	W_C_BU   = 1
	W_C_GONG = 2
	W_C_QI   = 3
	W_C_CHE  = 4
)

//士气价格
var PRICE_SHIQI []int = []int{1, 2, 3, 4, 5}
