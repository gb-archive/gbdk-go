package asset

import (
	"github.com/Akatsuki-py/gbdk-go/api/gb"
	"github.com/Akatsuki-py/gbdk-go/api/macro"
)

var BGTileBank = macro.Define(0)

var BGTile = []gb.UINT8{
	0xFF, 0xFF, 0xFB, 0xFB, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0x7F, 0x7F, 0xF7, 0xF7, 0xFF, 0xFF,
}
