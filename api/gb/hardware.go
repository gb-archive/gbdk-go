package gb

var (
	P1_REG    = 0xff00 /* Joystick: 1.1.P15.P14.P13.P12.P11.P10 */
	SB_REG    = 0xff01 /* Serial IO data buffer */
	SC_REG    = 0xff02 /* Serial IO control register */
	DIV_REG   = 0xff04 /* Divider register */
	TIMA_REG  = 0xff05 /* Timer counter */
	TMA_REG   = 0xff06 /* Timer modulo */
	TAC_REG   = 0xff07 /* Timer control */
	IF_REG    = 0xff0f /* Interrupt flags: 0.0.0.JOY.SIO.TIM.LCD.VBL */
	NR10_REG  = 0xff10 /* Sound register */
	NR11_REG  = 0xff11 /* Sound register */
	NR12_REG  = 0xff12 /* Sound register */
	NR13_REG  = 0xff13 /* Sound register */
	NR14_REG  = 0xff14 /* Sound register */
	NR21_REG  = 0xff16 /* Sound register */
	NR22_REG  = 0xff17 /* Sound register */
	NR23_REG  = 0xff18 /* Sound register */
	NR24_REG  = 0xff19 /* Sound register */
	NR30_REG  = 0xff1a /* Sound register */
	NR31_REG  = 0xff1b /* Sound register */
	NR32_REG  = 0xff1c /* Sound register */
	NR33_REG  = 0xff1d /* Sound register */
	NR34_REG  = 0xff1e /* Sound register */
	NR41_REG  = 0xff20 /* Sound register */
	NR42_REG  = 0xff21 /* Sound register */
	NR43_REG  = 0xff22 /* Sound register */
	NR44_REG  = 0xff23 /* Sound register */
	NR50_REG  = 0xff24 /* Sound register */
	NR51_REG  = 0xff25 /* Sound register */
	NR52_REG  = 0xff26 /* Sound register */
	LCDC_REG  = 0xff40 /* LCD control */
	STAT_REG  = 0xff41 /* LCD status */
	SCY_REG   = 0xff42 /* Scroll Y */
	SCX_REG   = 0xff43 /* Scroll X */
	LY_REG    = 0xff44 /* LCDC Y-coordinate */
	LYC_REG   = 0xff45 /* LY compare */
	DMA_REG   = 0xff46 /* DMA transfer */
	BGP_REG   = 0xff47 /* BG palette data */
	OBP0_REG  = 0xff48 /* OBJ palette 0 data */
	OBP1_REG  = 0xff49 /* OBJ palette 1 data */
	WY_REG    = 0xff4a /* Window Y coordinate */
	WX_REG    = 0xff4b /* Window X coordinate */
	KEY1_REG  = 0xff4d /* CPU speed */
	VBK_REG   = 0xff4f /* VRAM bank */
	HDMA1_REG = 0xff51 /* DMA control 1 */
	HDMA2_REG = 0xff52 /* DMA control 2 */
	HDMA3_REG = 0xff53 /* DMA control 3 */
	HDMA4_REG = 0xff54 /* DMA control 4 */
	HDMA5_REG = 0xff55 /* DMA control 5 */
	RP_REG    = 0xff56 /* IR port */
	BCPS_REG  = 0xff68 /* BG color palette specification */
	BCPD_REG  = 0xff69 /* BG color palette data */
	OCPS_REG  = 0xff6a /* OBJ color palette specification */
	OCPD_REG  = 0xff6b /* OBJ color palette data */
	SVBK_REG  = 0xff70 /* WRAM bank */
	IE_REG    = 0xffff /* Interrupt enable */
)