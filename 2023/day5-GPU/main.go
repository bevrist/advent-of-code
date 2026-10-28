package main

import (
	"fmt"

	"github.com/MasterOfBinary/go-opencl/opencl"
)

const (
	deviceType = opencl.DeviceTypeAll

	arrSize = 4

	programCode = `
int constant lookup[4] = {2, 4, 7, 1};
//int constant lookup2[4] = {2, 4, 7, 1};

kernel void kern(global float* out)
{
	size_t i = get_global_id(0);

	int count = i;
	for (int j=0; j<2; j++){
		count = count+j;
	}
	// count = count - 10;
	out[i] = count;
}
`
)

func main() {
	platforms, err := opencl.GetPlatforms()
	if err != nil {
		panic(err)
	}

	foundDevice := false

	// var platform opencl.Platform
	var device opencl.Device
	var name string
	for _, curPlatform := range platforms {
		err = curPlatform.GetInfo(opencl.PlatformName, &name)
		if err != nil {
			panic(err)
		}

		var devices []opencl.Device
		devices, err = curPlatform.GetDevices(deviceType)
		if err != nil {
			panic(err)
		}

		// Use the first available device
		if len(devices) > 0 && !foundDevice {
			var available bool
			err = devices[0].GetInfo(opencl.DeviceAvailable, &available)
			if err == nil && available {
				// platform = curPlatform
				device = devices[0]
				foundDevice = true
			}
		}

		// version := curPlatform.GetVersion()
		// fmt.Printf("Name: %v, devices: %v, version: %v\n", name, len(devices), version)
	}

	if !foundDevice {
		panic("No device found")
	}

	var context opencl.Context
	context, err = device.CreateContext()
	if err != nil {
		panic(err)
	}
	defer context.Release()

	var commandQueue opencl.CommandQueue
	commandQueue, err = context.CreateCommandQueue(device)
	if err != nil {
		panic(err)
	}
	defer commandQueue.Release()

	var program opencl.Program
	program, err = context.CreateProgramWithSource(programCode)
	if err != nil {
		panic(err)
	}
	defer program.Release()

	var log string
	err = program.Build(device, &log)
	if err != nil {
		fmt.Println(log)
		panic(err)
	}

	kernel, err := program.CreateKernel("kern")
	if err != nil {
		panic(err)
	}
	defer kernel.Release()

	buffer, err := context.CreateBuffer([]opencl.MemFlags{opencl.MemWriteOnly}, arrSize*4)
	if err != nil {
		panic(err)
	}
	defer buffer.Release()

	err = kernel.SetArg(0, buffer.Size(), &buffer)
	if err != nil {
		panic(err)
	}

	err = commandQueue.EnqueueNDRangeKernel(kernel, 1, []uint64{arrSize})
	if err != nil {
		panic(err)
	}

	commandQueue.Flush()
	commandQueue.Finish()

	data := make([]float32, arrSize)

	err = commandQueue.EnqueueReadBuffer(buffer, true, data)
	if err != nil {
		panic(err)
	}

	// printHeader("Output")
	fmt.Print("out: ")
	for _, item := range data {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}
