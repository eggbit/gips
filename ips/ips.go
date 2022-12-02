package ips

import (
	"errors"
	"math/big"
	"os"
)

type ipsFile struct {
	data   []byte
	index  int // Holds the current position in the file
	record ipsRecord
}

type ipsRecord struct {
	offset    int // Where to write this record in the ROM
	size      int // The size of this record's data
	rle_size  int
	rle_value int // The value to be copied to the ROM 'rle_size' times.
}

type ipsData struct {
	as_int   int
	as_bytes []byte
}

func New(ips_path string) (*ipsFile, error) {
	data, err := os.ReadFile(ips_path)

	if err != nil {
		return nil, err
	}

	i := ipsFile{
		data:  data,
		index: 5, // Skip the header
	}

	// Validate IPS file
	if string(i.data[:5]) != "PATCH" {
		return nil, errors.New("invalid IPS file")
	}

	return &i, nil
}

func (ips *ipsFile) Apply(rom_path, out_path string) error {
	rom_data, err := os.ReadFile(rom_path)

	if err != nil {
		return err
	}

	new_rom, err := ips.patch(rom_data)

	if err != nil {
		return err
	}

	// Dump the newly patched ROM
	err = os.WriteFile(out_path, new_rom, 0664)

	if err != nil {
		return err
	}

	return nil
}

func (ips *ipsFile) patch(rom_data []byte) ([]byte, error) {
	// Loop through all the records in the IPS file until EOF (0x454f46)
	// TODO: Create new buffer, write new ROM to it instead of modifying the ROM itself
	for string(ips.data[ips.index:ips.index+3]) != "EOF" {
		ips.record.offset = ips.read(3).as_int
		ips.record.size = ips.read(2).as_int

		// RLE handling
		if ips.record.size == 0 {
			ips.record.rle_size = ips.read(2).as_int
			ips.record.rle_value = ips.read(1).as_int

			// Write the changes to the ROM
			for i := 0; i < ips.record.rle_size; i++ {
				rom_data[ips.record.offset+i] = byte(ips.record.rle_value)
			}
		} else {
			// Write the changes to the ROM
			copy(rom_data[ips.record.offset:], ips.read(ips.record.size).as_bytes)
		}
	}

	return rom_data, nil
}

func (ips *ipsFile) read(num_bytes int) ipsData {
	data_block := ips.data[ips.index : ips.index+num_bytes]

	i := ipsData{
		as_int:   int(big.NewInt(0).SetBytes(data_block).Int64()),
		as_bytes: data_block,
	}

	ips.index += num_bytes

	return i
}
