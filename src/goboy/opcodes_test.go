package GoBoy

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLD_r_r(t *testing.T) {
    gb := &GameBoy{
        Register: &Register{},
    }
    gb.set8Reg(B, 0xfe)
    gb.LD_r_r(0x50) // LD D, B
    assert.Equal(t, gb.get8Reg(B), uint8(0xfe))
    assert.Equal(t, gb.get8Reg(D), uint8(0xfe))
    assert.Equal(t, gb.get16Reg(PC), uint16(0x0001))
}

func TestLD_r_n(t *testing.T) {
    gb := &GameBoy{
        Register: &Register{},
    }
    gb.LD_r_n(0xead) // LD C, 0xad
    assert.Equal(t, gb.get8Reg(C), uint8(0xad)) 
    assert.Equal(t, gb.get16Reg(PC), uint16(0x0002))
}

func TestLD_r_hl(t * testing.T) {
    gb := &GameBoy{
        Register: &Register{},
    }
    gb.mainMemory.rom0[0x1234] = 0x42
    memValue := gb.mainMemory.read(0x1234)
    assert.Equal(t, memValue, uint8(0x42))
    gb.set16Reg(HL, 0x1234)
    gb.LD_r_hl(0x5e)
    assert.Equal(t, gb.get8Reg(E), uint8(0x42))
}
