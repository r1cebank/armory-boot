// https://github.com/f-secure-foundry/armory-boot
//
// Copyright (c) F-Secure Corporation
// https://foundry.f-secure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package main

import (
	_ "unsafe"

	"github.com/f-secure-foundry/tamago/dma"
)

// Override imx6ul.ramStart, usbarmory.ramSize and dma allocation, as the
// mapped kernel image needs to be within the first 128MiB of RAM.

//go:linkname ramStart runtime.ramStart
var ramStart uint32 = 0x90000000

//go:linkname ramSize runtime.ramSize
var ramSize uint32 = 0x08000000

// DMA region for target kernel boot
var mem *dma.Region

// DMA region for bootloader operation
const (
	dmaStart = 0x98000000
	dmaSize  = 0x08000000
)

// DMA region for target kernel boot
const (
	memoryStart = 0x80000000
	memorySize  = 0x10000000

	kernelOffset = 0x00800000
	paramsOffset = 0x07000000
	initrdOffset = 0x08000000
)

func init() {
	dma.Init(dmaStart, dmaSize)

	mem = &dma.Region{
		Start: memoryStart,
		Size:  memorySize,
	}

	mem.Init()
	mem.Reserve(memorySize, 0)
}
